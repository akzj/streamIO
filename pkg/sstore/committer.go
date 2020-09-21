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
	queue *block_queue.QueueWithContext

	mutableMStreamMap *mStreamTable

	flushQueue *block_queue.QueueWithContext

	indexUpdateQueue *block_queue.QueueWithContext

	store *Store
}

func newCommitter(
	store *Store,
	queue *block_queue.QueueWithContext,
	flushQueue *block_queue.QueueWithContext,
	updateIndexQueue *block_queue.QueueWithContext) *committer {

	return &committer{
		queue:             queue,
		mutableMStreamMap: newMStreamTable(store.endMap, store.options.BlockSize, 128),
		flushQueue:        flushQueue,
		indexUpdateQueue:  updateIndexQueue,
		store:             store,
	}
}

//commitSegmentFile receive segment file from master,and append segment to local store

func (c *committer) flush() {
	mStreamMap := c.mutableMStreamMap
	c.mutableMStreamMap = newMStreamTable(mStreamMap.endMap, c.store.options.BlockSize,
		len(c.mutableMStreamMap.mStreams))

	c.store.appendMStreamTable(mStreamMap)

	if err := c.flushQueue.Push(flushSegment{mStreamTable: mStreamMap, callback: func(filename string, err error) {
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

/*func (c *committer) processLoop() {
	for {
		items, err := c.queue.PopAll(nil)
		if err != nil {
			log.Warn(err)
			return
		}
		for _, item := range items {
			if c.mutableMStreamMap.mSize >= c.store.options.MaxMStreamTableSize {
				c.flush()
			}
			request := item.(*WriteEntryRequest)
			stream, end := c.mutableMStreamMap.appendEntry(request.Entry)
			if end == -1 {
				request.err = ErrOffset
				continue
			}
			request.end = end
			if stream != nil {
				c.store.indexTable.update(stream)
			}
			err := c.store.streamWatcher.queue.Push(notify{
				streamID: request.Entry.StreamID,
				end:      end,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		c.store.callbackWorker.queue.PushMany(items)
	}
}*/


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
			if c.mutableMStreamMap.mSize >= c.store.options.MaxMStreamTableSize {
				c.flush()
			}
			request := item.(*WriteEntryRequest)
			mStream, end := c.mutableMStreamMap.appendEntry(request.Entry)
			if end == -1 {
				request.err = ErrOffset
				continue
			}
			request.end = end
			if mStream != nil {
				mStreams = append(mStreams, mStream)
			}
			notifies = append(notifies, notify{
				streamID: request.Entry.StreamID,
				end:      end,
			})
		}
		update := updateIndexTable{
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

