// Copyright (c) 2017-2018, Jan Cajthaml <jan.cajthaml@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	event := make(chan []byte)

	var wg sync.WaitGroup

	sink := NewSink(event)
	processor := NewProcessor(event)

	wg.Add(1)
	go func() {
		defer wg.Done()
		processor.Run(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sink.Run(ctx)
	}()

	cancelSignal := make(chan os.Signal, 1)
	signal.Notify(cancelSignal, os.Interrupt)
	defer func() {
		signal.Stop(cancelSignal)
		cancel()
	}()

	go func() {
		select {
		case <-cancelSignal:
			cancel()
		case <-ctx.Done():
		}
	}()

	wg.Wait()
}
