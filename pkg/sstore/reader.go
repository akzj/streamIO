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
	"github.com/pkg/errors"
	"io"
)

type reader struct {
	offset   int64
	streamID int64
	index    *offsetIndex
	endMap   *int64LockMap
}

func newReader(streamID int64, index *offsetIndex, endMap *int64LockMap) *reader {
	return &reader{
		offset:   0,
		streamID: streamID,
		index:    index,
		endMap:   endMap,
	}
}

func (r *reader) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, ErrWhence
	case io.SeekStart:
	case io.SeekCurrent:
		offset += r.offset
	case io.SeekEnd:
		limit, ok := r.endMap.get(r.streamID)
		if !ok {
			panic("no find stream end")
		}
		offset += limit
	}
	begin, ok := r.index.begin()
	if ok && offset < begin {
		return 0, ErrOffset
	}
	r.offset = offset
	return offset, nil
}

func (r *reader) Read(p []byte) (int, error) {
	buf := p
	var size int
	for len(buf) > 0 {
		item, err := r.index.find(r.offset)
		if err != nil {
			return 0, err
		}
		if item.stream != nil {
			n, err := item.stream.ReadAt(buf, r.offset)
			if err != nil {
				if err == io.EOF {
					if size == 0 {
						return 0, err
					}
					return size, nil
				}
				return size, err
			}
			buf = buf[n:]
			size += n
			r.offset += int64(n)
		} else if item.segment != nil {
			if item.segment.IncRef() < 0 {
				return size, errors.WithStack(ErrOffset)
			}
			n, err := item.segment.Reader(r.streamID).ReadAt(buf, r.offset)
			item.segment.DecRef()
			if err != nil {
				if err == io.EOF {
					if size == 0 {
						return size, err
					}
					return size, nil
				} else {
					return size, err
				}
			}
			buf = buf[n:]
			size += n
			r.offset += int64(n)
		}
	}
	return size, nil
}
