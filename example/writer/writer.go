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

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/akzj/streamIO/client"
	"sync"
	"time"
)

func main() {
	var metaServer string
	var data string
	flag.StringVar(&metaServer, "ms", "127.0.0.1:5000", "--ms [ip:port]")
	flag.StringVar(&data, "data", "hello"+time.Now().String(), "--data []")
	flag.Parse()

	ctx := context.Background()
	msClient, err := client.NewMetaServiceClient(ctx, metaServer)
	if err != nil {
		panic(err)
	}
	client := client.NewClient(msClient)

	infoItem, err := client.GetOrCreateStreamInfoItem(ctx, "hello")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(infoItem)

	session, err := client.NewStreamSession(ctx, 1, infoItem)
	if err != nil {
		panic(err.Error())
	}

	writer, err := session.NewWriter()
	if err != nil {
		panic(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	writer.WriteWithCb([]byte("hello world"), func(err error) {
		wg.Done()
		if err != nil {
			panic(err.Error())
		}
	})
	wg.Wait()
}
