// Copyright (c) 2016-2018, Jan Cajthaml <jan.cajthaml@gmail.com>
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

func NewSink() *sink {
	return &sink{event: make(chan string)}
}

func (r *sink) Run(addr *net.UDPAddr) {
	inputStreamConn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println("Fatal: %v", err)
		os.Exit(1)
	}

	defer inputStreamConn.Close()

	buf := make([]byte, BUF_SIZE)

forever:
	n, _, err := inputStreamConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("Error: %v", err)
		goto forever
	}

	r.event <- string(buf[0:n])

	goto forever
}
