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
	"fmt"
	"net"
	"os"
)

// NewSink creates new UDP sink with channel for complete events
func NewSink() Sink {
	return Sink{
		event: make(chan []byte),
	}
}

// Run starts UDP sink with UPD stream source
func (r *Sink) Run(addr *net.UDPAddr) {
	inputStreamConn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer inputStreamConn.Close()

	buf := make([]byte, BufferSize)

	for {
		n, _, err := inputStreamConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		r.event <- buf[0:n]
	}
}
