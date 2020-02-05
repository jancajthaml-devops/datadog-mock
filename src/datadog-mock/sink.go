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
	"net"
	"os"
	"strings"
	"syscall"
)

type Sink struct {
	out chan<- []byte
}

// NewSink creates new UDP sink with channel for complete events
func NewSink(out chan<- []byte) Sink {
	return Sink{
		out: out,
	}
}

func (r *Sink) ReadFromUDP(conn *net.UDPConn) {
	buf := make([]byte, 0xffff)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		switch {
		case (n == 0 && err == nil):
			return
		case err != nil && strings.Contains(err.Error(), "use of closed network connection"):
			return
		case err == syscall.EAGAIN || err == syscall.EWOULDBLOCK:
			continue
		case n == 0:
			continue
		case err != nil:
			fmt.Println(err)
			continue
		default:
		}
		var result = make([]byte, n)
		copy(result, buf[0:n])
		r.out <- result
	}
}

// Run starts UDP sink with UPD stream source
func (r *Sink) Run(ctx context.Context) {
	addr, err := net.ResolveUDPAddr("udp", ":8125")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	go r.ReadFromUDP(conn)

	<-ctx.Done()
}
