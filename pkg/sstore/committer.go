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
	"github.com/akzj/streamIO/pkg/block-queue"
	log "github.com/sirupsen/logrus"
)

type committer struct {
	store            *Store
	streamTable      *streamTable
	queue            *block_queue.QueueWithContext
	flushQueue       *block_queue.QueueWithContext
	indexUpdateQueue *block_queue.QueueWithContext
}

func newCommitter(
	store *Store,
	queue *block_queue.QueueWithContext,
	flushQueue *block_queue.QueueWithContext,
	updateIndexQueue *block_queue.QueueWithContext) *committer {

	return &committer{
		store:            store,
		streamTable:      newStreamTable(store.endMap, store.options.BlockSize, 128),
		queue:            queue,
		flushQueue:       flushQueue,
		indexUpdateQueue: updateIndexQueue,
	}
}

func (c *committer) flush() {
	mStreamMap := c.streamTable
	c.streamTable = newStreamTable(mStreamMap.endMap, c.store.options.BlockSize,
		len(c.streamTable.streams))

	c.store.appendStreamTable(mStreamMap)

	if err := c.flushQueue.Push(flushSegment{streamTable: mStreamMap, callback: func(filename string, err error) {
		if err != nil {
			log.Fatal(err)
		}
		if err := c.store.flushCallback(filename); err != nil {
			log.Fatal(err)
		}
	}}); err != nil {
		log.Fatal(err)
	}
}

func (c *committer) processLoop() {
	var mStreams = make([]*stream, 0, 128)
	for {
		items, err := c.queue.PopAll(nil)
		if err != nil {
			log.Warn(err)
			return
		}
		var notifies = make([]interface{}, 0, len(items))
		for _, item := range items {
			if c.streamTable.size >= c.store.options.MaxMStreamTableSize {
				c.flush()
			}
			request := item.(*WriteEntry)
			mStream, err := c.streamTable.appendEntry(request.Entry, &request.end)
			if err != nil {
				request.err = err
				continue
			}
			if mStream != nil {
				mStreams = append(mStreams, mStream)
			}
			notifies = append(notifies, notify{
				streamID: request.Entry.StreamID,
				end:      request.end,
			})
		}
		update := updateSectionTable{
			notifies:  notifies,
			callbacks: items,
		}
		if len(mStreams) != 0 {
			update.mStreams = mStreams
			mStreams = make([]*stream, 0, 128)
		}
		if err := c.indexUpdateQueue.Push(update); err != nil {
			log.Fatal(err)
		}
	}
}
