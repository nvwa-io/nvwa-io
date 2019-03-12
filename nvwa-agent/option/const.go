// Copyright 2019 - now The https://github.com/nvwa-io/nvwa-io Authors
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

package option

import "log"

var (
	buildStatusMap = map[string]int{
		"building":      20,
		"build-success": 30,
		"build-failed":  40,
		"pack-success":  50,
		"pack-failed":   60,
		"push-success":  70,
		"push-failed":   80,
	}
)

// trans string status to nvwa-server build status (int)
func BuildStatus2Int(status string) int {
	if _, ok := buildStatusMap[status]; !ok {
		log.Fatal("Invalid build status: " + status)
	}

	return buildStatusMap[status]
}
