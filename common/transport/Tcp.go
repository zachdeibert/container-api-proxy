// Copyright 2024 Zach Deibert
//
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
package transport

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type tcpTransport struct {
}

func newTcpTransport() *tcpTransport {
	return &tcpTransport{}
}

func (trans tcpTransport) Dial(ctx context.Context, address net.Addr) (net.Conn, error) {
	if addr, err := trans.parseAddr(address); err != nil {
		return nil, err
	} else {
		return (&net.Dialer{}).DialContext(ctx, address.Network(), fmt.Sprintf("%s:%d", addr.IP.String(), addr.Port))
	}
}

func (trans tcpTransport) Listen(address net.Addr) (net.Listener, error) {
	if addr, err := trans.parseAddr(address); err != nil {
		return nil, err
	} else {
		return net.ListenTCP(address.Network(), &addr)
	}
}

func (trans tcpTransport) Parse(address string) (addr net.Addr, prefix string, err error) {
	if !strings.HasPrefix(address, "tcp://") {
		return nil, "", fmt.Errorf("address '%s' is not a TCP address", address)
	}
	address = address[len("tcp://"):]
	parts := strings.SplitN(address, "/", 2)
	addr = NewAddr("tcp", parts[0])
	if len(parts) == 2 {
		prefix = "/" + parts[1]
	}
	if _, err := trans.parseAddr(addr); err != nil {
		return nil, "", err
	}
	return
}

func (trans tcpTransport) parseAddr(address net.Addr) (net.TCPAddr, error) {
	parts := strings.SplitN(address.String(), ":", 2)
	addr := net.TCPAddr{}
	if addr.IP = net.ParseIP(parts[0]); addr.IP == nil {
		if ips, err := net.LookupIP(parts[0]); err != nil {
			return addr, err
		} else if len(ips) > 0 {
			addr.IP = ips[0]
		} else {
			return addr, fmt.Errorf("unable to resolve host '%s'", parts[0])
		}
	}
	if len(parts) == 2 {
		if port, err := strconv.ParseInt(parts[1], 10, 16); err != nil {
			return addr, err
		} else {
			addr.Port = int(port)
		}
	} else {
		addr.Port = 2375
	}
	return addr, nil
}
