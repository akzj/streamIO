// Copyright 2020-2026 The streamIO Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sstore

import (
	"bufio"
	"encoding/binary"
	"github.com/akzj/streamIO/pkg/ref"
	"github.com/akzj/streamIO/pkg/sstore/pb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"hash/crc32"
	"io"
	"os"
	"sync"
)

type segment struct {
	*ref.RefCount
	filename string
	f        *os.File
	meta     *pb.SegmentMeta
	l        *sync.RWMutex
	delete   bool
}

func newSegment() *segment {
	sm := &segment{
		filename: "",
		f:        nil,
		meta: &pb.SegmentMeta{
			From:        &pb.Version{},
			To:          &pb.Version{},
			CreateTS:    0,
			SectionOffsets: map[int64]*pb.SectionOffset{},
		},
		l:      new(sync.RWMutex),
		delete: false,
	}
	sm.RefCount = ref.NewRefCount(0, func() {
		_ = sm.close()
	})
	return sm
}

func openSegment(filename string) (*segment, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()
	segment := newSegment()
	segment.filename = filename
	segment.f = f
	//seek to Read meta length
	if _, err := f.Seek(-4, io.SeekEnd); err != nil {
		return nil, errors.WithStack(err)
	}
	//Read meta length
	var headerLen int32
	if err := binary.Read(f, binary.BigEndian, &headerLen); err != nil {
		return nil, errors.WithStack(err)
	}
	//seek to Read meta
	if _, err := f.Seek(-int64(headerLen)-4, io.SeekEnd); err != nil {
		return nil, errors.WithStack(err)
	}
	data := make([]byte, headerLen)
	n, err := f.Read(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if n != len(data) {
		return nil, errors.WithMessage(io.ErrUnexpectedEOF, "Read segment head failed")
	}
	if err := proto.Unmarshal(data, segment.meta); err != nil {
		return nil, errors.WithStack(err)
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, errors.WithStack(err)
	}
	f = nil
	return segment, nil
}
func (s *segment) FromVersion() *pb.Version {
	return s.meta.From
}

func (s *segment) offsetInfo(streamID int64) (*pb.SectionOffset, error) {
	indexInfo, ok := s.meta.SectionOffsets[streamID]
	if ok == false {
		return indexInfo, ErrNoSectionOffset
	}
	return indexInfo, nil
}

func (s *segment) Reader(streamID int64) *sectionReader {
	info, ok := s.meta.SectionOffsets[streamID]
	if !ok {
		return nil
	}
	return &sectionReader{
		sectionOffset: info,
		r:             io.NewSectionReader(s.f, info.Offset, info.End-info.Begin),
	}
}

func (s *segment) deleteOnClose(delete bool) error {
	s.l.Lock()
	defer s.l.Unlock()
	if s.f == nil {
		return errors.Errorf("segment is close")
	}
	s.delete = delete
	return nil
}

func (s *segment) close() error {
	s.l.Lock()
	defer s.l.Unlock()
	if s.f == nil {
		return nil
	}
	if err := s.f.Close(); err != nil {
		return err
	}
	if s.delete {
		if err := os.Remove(s.filename); err != nil {
			return errors.WithStack(err)
		}
	}
	s.f = nil
	return nil
}

func flushStreamTable(filename string, table *streamTable) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	seg := newSegment()
	seg.f = f
	seg.filename = filename
	var Offset int64
	writer := bufio.NewWriterSize(seg.f, 1024*1024)
	for streamID, mStream := range table.streams {
		hash := crc32.NewIEEE()
		mWriter := io.MultiWriter(writer, hash)
		n, err := mStream.WriteTo(mWriter)
		if err != nil {
			return err
		}
		seg.meta.SectionOffsets[streamID] = &pb.SectionOffset{
			StreamID: streamID,
			Offset:   Offset,
			CRC:      hash.Sum32(),
			Begin:    mStream.begin,
			End:      mStream.end,
		}
		Offset += int64(n)
	}
	seg.meta.From = table.from
	seg.meta.To = table.to
	seg.meta.CreateTS = table.CreateTS.Unix()
	data, err := proto.Marshal(seg.meta)
	if err != nil {
		return err
	}
	if _, err := writer.Write(data); err != nil {
		return err
	}
	if err := binary.Write(writer, binary.BigEndian, int32(len(data))); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	if err := seg.close(); err != nil {
		return err
	}
	return nil
}
