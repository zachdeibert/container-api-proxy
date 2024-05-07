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
package model

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/zachdeibert/container-api-proxy/common"
)

type Api struct {
	Version common.Version
}

func NewApi(version common.Version) Api {
	filename := path.Join(DownloadCache, fmt.Sprintf("%s.yaml", version))
	f, err := os.Open(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
		if f, err = os.Create(filename); err != nil {
			panic(err)
		}
		defer f.Close()
		url := fmt.Sprintf("https://docs.docker.com/reference/engine/%s.yaml", version)
		fmt.Printf("Downloading %s...\n", url)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		if _, err = io.Copy(f, res.Body); err != nil {
			panic(err)
		}
		if _, err = f.Seek(0, io.SeekStart); err != nil {
			panic(err)
		}
	} else {
		defer f.Close()
	}
	return Api{Version: version}
}
