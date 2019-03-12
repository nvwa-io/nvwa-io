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

package svrs

import (
	"github.com/go-ozzo/ozzo-dbx"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
)

var DefaultJobStepSvr = new(JobStepSvr)

type JobStepSvr struct{}

func (t *JobStepSvr) CreateByJob(job *JobEntity, step int) (int64, error) {
	id, err := DefaultJobStepDao.CreateByMap(dbx.Params{
		"job_id":        job.Id,
		"app_id":        job.AppId,
		"deployment_id": job.DeploymentId,
		"step":          step,
		"status":        JOB_STEP_STATUS_DEALING,
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *JobStepSvr) UpdateByJobIdStep(jobId int64, step int, params dbx.Params) error {
	_, err := GetDb().Update(DefaultJobStepDao.Table(), params, dbx.HashExp{
		"job_id": jobId,
		"step":   step,
	}).Execute()
	if err != nil {
		return err
	}

	return nil
}

func (t *JobStepSvr) GetByJobId(jobId int64) ([]JobStepEntity, error) {
	list := make([]JobStepEntity, 0)
	err := DefaultJobStepDao.GetAllByExp(dbx.HashExp{
		"job_id": jobId,
	}, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
