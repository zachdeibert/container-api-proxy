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
package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/zachdeibert/container-api-proxy/common/transport"
)

type ApiServer struct {
	mux        *http.ServeMux
	server     http.Server
	serveError error
}

func NewApiServer() *ApiServer {
	return &ApiServer{mux: http.NewServeMux(), serveError: nil, server: http.Server{}}
}

func (server *ApiServer) Start(address string) error {
	trans, err := transport.GetTransport(address)
	if err != nil {
		return err
	}
	addr, prefix, err := trans.Parse(address)
	if err != nil {
		return err
	}
	listener, err := trans.Listen(addr)
	if err != nil {
		return err
	}
	if prefix != "" {
		mux := http.NewServeMux()
		mux.Handle(prefix, server.mux)
		server.server.Handler = mux
	} else {
		server.server.Handler = server.mux
	}
	go func() {
		server.serveError = server.server.Serve(listener)
	}()
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		conn, err := trans.Dial(ctx, addr)
		cancel()
		if server.serveError != nil {
			return server.serveError
		} else if err == nil {
			conn.Close()
			log.Printf("Listener on %s initialized.", address)
			return nil
		}
	}
}

func (server *ApiServer) Stop() error {
	if err := server.server.Shutdown(context.Background()); err != nil {
		return err
	}
	if server.serveError != http.ErrServerClosed {
		return server.serveError
	}
	return nil
}
