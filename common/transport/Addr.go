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

type Addr struct {
	network string
	string  string
}

func NewAddr(network string, string string) Addr {
	return Addr{network: network, string: string}
}

func (addr Addr) Network() string {
	return addr.network
}

func (addr Addr) String() string {
	return addr.string
}
