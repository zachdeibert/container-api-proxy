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
package common

import (
	"errors"
	"fmt"
)

type Version struct {
	Major int
	Minor int
}

func NewVersion(major int, minor int) Version {
	return Version{Major: major, Minor: minor}
}

func (v Version) Next() Version {
	if v.Major != 1 {
		panic(errors.New("only major version 1 is currently supported"))
	}
	return NewVersion(v.Major, v.Minor+1)
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%02d", v.Major, v.Minor)
}
