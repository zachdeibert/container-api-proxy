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
	"strings"
)

type Transport interface {
	Dial(ctx context.Context, address net.Addr) (net.Conn, error)
	Listen(address net.Addr) (net.Listener, error)
	Parse(address string) (addr net.Addr, prefix string, err error)
}

var transports map[string]Transport = map[string]Transport{"tcp": newTcpTransport()}

func GetTransport(address string) (Transport, error) {
	if parts := strings.SplitN(address, "://", 2); len(parts) != 2 {
		return nil, fmt.Errorf("invalid address '%s'", address)
	} else if transport, ok := transports[parts[0]]; !ok {
		return nil, fmt.Errorf("invalid protocol '%s'", parts[0])
	} else {
		return transport, nil
	}
}
