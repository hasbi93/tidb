// Copyright 2023 PingCAP, Inc.
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

package util

import (
	"fmt"
	"regexp"

	"github.com/pingcap/tidb/build"
)

// shouldRun checks whether a `file` should be analyzed in the specific pass
func shouldRun(passName string, fileName string) bool {
	config, ok := build.NogoConfig[passName]
	if !ok {
		return true
	}

	if config.OnlyFiles != nil {
		for _, f := range config.OnlyFiles {
			matched, err := regexp.MatchString(f, fileName)
			if err != nil {
				panic(fmt.Sprintf("regex is wrong: %s", f))
			}

			if matched {
				return true
			}
		}

		return false
	}

	if config.ExcludeFiles != nil {
		for f := range config.ExcludeFiles {
			matched, err := regexp.MatchString(f, fileName)
			if err != nil {
				panic(fmt.Sprintf("regex is wrong: %s", f))
			}

			if matched {
				return false
			}
		}

		return true
	}

	return true
}
