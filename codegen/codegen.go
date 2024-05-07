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
	"github.com/zachdeibert/container-api-proxy/codegen/model"
	"github.com/zachdeibert/container-api-proxy/common"
)

var (
	MinimumVersion = common.NewVersion(1, 25)
	MaximumVersion = common.NewVersion(1, 45)
)

func main() {
	version := MinimumVersion
	for {
		model.NewApi(version)
		if version == MaximumVersion {
			break
		}
		version = version.Next()
	}
}
