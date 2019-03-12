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
	"fmt"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
	"time"
)

const (
	HOURS_AGO = 3600 * 48
)

var DefaultBuildWorker = new(BuildWorker)

type BuildWorker struct{}

// Deal build tasks
func (t *BuildWorker) Deal() {
	for {
		t.DealOnce()
	}
}

func (t *BuildWorker) DealOnce() {
	// 1. Get all tasks which are waiting to build
	list, err := DefaultBuildSvr.GetWaitToBuilds(HOURS_AGO)
	if err != nil {
		logger.Errorf("Failed to GetWaitToBuilds, err=%s", err.Error())
		time.Sleep(time.Second * 30)
		return
	}

	if len(list) == 0 {
		time.Sleep(3)
		return
	}

	// 2. Start deal every task
	for _, v := range list {
		go DefaultBuildSvr.DealBuild(&v)
	}
}

// notify which builds are finished (success or failed)
func (t *BuildWorker) NotifyBuildEnd() {
	for {
		t.NotifyBuildEndOnce()
	}
}

func (t *BuildWorker) NotifyBuildEndOnce() {
	// 1. look up build list which are finished and has not been notify
	list, err := DefaultBuildSvr.GetWaitToNotifyList()
	if err != nil {
		time.Sleep(time.Second * 30)
		logger.Errorf("Failed to GetWaitToNotifyList, err=%s", err.Error())
		return
	}

	if len(list) == 0 {
		time.Sleep(3)
		return
	}

	// 2. notify user who create this build
	for _, v := range list {
		// @TODO trans to human language
		go DefaultNotifySvr.CommitNotifyByUid(v.Uid, fmt.Sprintf("Build %d finished.", v.Id))
	}

}
