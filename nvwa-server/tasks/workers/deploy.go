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

package workers

import (
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
	"time"
)

var DefaultDeployWorker = new(DeployWorker)

type DeployWorker struct{}

func (t *DeployWorker) Deal() {
	for {
		t.DealOnce()
	}
}

func (t *DeployWorker) DealOnce() {
	// 1. Get all tasks which are waiting to build
	list, err := DefaultJobSvr.GetWaitToJobs(HOURS_AGO)
	if err != nil {
		logger.Errorf("Failed to GetWaitToJobs, err=%s", err.Error())
		time.Sleep(time.Second * 30)
		return
	}

	if len(list) == 0 {
		time.Sleep(3)
		return
	}

	// 2. Start deal jobs
	for _, v := range list {
		go DefaultJobSvr.DealJob(&v)
	}
}
