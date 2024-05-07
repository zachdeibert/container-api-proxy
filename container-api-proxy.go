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
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/zachdeibert/container-api-proxy/common/transport"
	"github.com/zachdeibert/container-api-proxy/server"
)

type listenAddresses []string

func (la listenAddresses) String() string {
	return strings.Join(la, ", ")
}

func (la *listenAddresses) Set(value string) error {
	if trans, err := transport.GetTransport(value); err != nil {
		return err
	} else if _, _, err := trans.Parse(value); err != nil {
		return err
	}
	*la = append(*la, value)
	return nil
}

var ListenAddresses listenAddresses

func main() {
	msg := strings.Builder{}
	msg.WriteString("Starting up container API proxy with arguments:")
	for _, arg := range os.Args {
		msg.WriteString(fmt.Sprintf(" \"%s\"", arg))
	}
	log.Print(msg.String())
	flag.Var(&ListenAddresses, "H", "addresses to listen on")
	flag.Var(&ListenAddresses, "host", "addresses to listen on")
	flag.Parse()
	servers := []*server.ApiServer{}
	for _, addr := range ListenAddresses {
		srv := server.NewApiServer()
		if err := srv.Start(addr); err != nil {
			log.Printf("Unable to listen on %s: %s", addr, err)
		} else {
			servers = append(servers, srv)
		}
	}
	if len(servers) == 0 {
		log.Print("No listeners were successfully configured.  Please specify a valid listener configuration with -H or --host.")
	} else {
		log.Println("Container API proxy initialization complete.")
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-c
	}
	log.Println("Shutting down container API proxy...")
	for _, srv := range servers {
		if err := srv.Stop(); err != nil {
			log.Printf("Unable to shut down listener: %s", err)
		}
	}
	log.Println("Container API proxy shutdown complete.")
}
