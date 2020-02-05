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
	"fmt"
	"regexp"
	"strings"
)

// types c|g|ms suffix (count, gauce, timer)
var pattern = regexp.MustCompile(`(?P<name>.*?):(?P<value>.*?)\|(?P<type>[a-z])(?:\|#(?P<tags>.*))?`)

type Processor struct {
	in <-chan []byte
}

func NewProcessor(in <-chan []byte) Processor {
	return Processor{
		in: in,
	}
}

func (r *Processor) ProcessEvents() {
	for {
		select {
		case event := <-r.in:
			if len(event) == 0 {
				return
			}
			processEvent(event)
		default:
		}
	}
}

func (r *Processor) Run(ctx context.Context) {
	go r.ProcessEvents()
	<-ctx.Done()
}

func processEvent(event []byte) {
	parsed := string(event)
	match := pattern.FindStringSubmatch(parsed)
	if len(match) < 3 {
		fmt.Println("Invalid event", parsed)
		return
	}
	fmt.Println(strings.TrimSuffix(parsed, "\n"))
}
