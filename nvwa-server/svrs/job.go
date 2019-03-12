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
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"time"
)

var DefaultJobSvr = new(JobSvr)

type JobSvr struct{}

func (t *JobSvr) GetByDeploymentIds(deploymentIds []int64) (map[int64][]JobEntity, error) {
	list := make([]JobEntity, 0)
	err := DefaultJobDao.GetAllByFieldInt64("deployment_id", deploymentIds, &list)
	if err != nil {
		return nil, err
	}

	// group by deploymentId
	res := make(map[int64][]JobEntity)
	for _, v := range list {
		if _, ok := res[v.DeploymentId]; !ok {
			res[v.DeploymentId] = make([]JobEntity, 0)
		}

		res[v.DeploymentId] = append(res[v.DeploymentId], v)
	}

	return res, nil
}

// get deploy tasks created nearly n seconds and status = DEPLOYMENT_STATUS_CREATED
func (t *JobSvr) GetWaitToJobs(seconds int64) ([]JobEntity, error) {
	list := make([]JobEntity, 0)

	limitDatetime := libs.Date("Y-m-d H:i:s", time.Now().Unix()-seconds)
	err := GetDb().Select("*").From(DefaultJobDao.Table()).
		Where(dbx.NewExp(fmt.Sprintf("ctime > '%s'", limitDatetime))).
		AndWhere(dbx.HashExp{"status": JOB_STATUS_READY}).
		OrderBy("id ASC").
		All(&list)
	if err != nil {
		logger.Errorf("Failed to GetWaiToJobs, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}

// deal deploy jobs which status = JOB_STATUS_READY
func (t *JobSvr) DealJob(job *JobEntity) {
	// 1.1 update job status to JOB_STATUS_DEALING
	_, err := DefaultJobDao.UpdateById(job.Id, dbx.Params{
		"status": JOB_STATUS_DEALING,
	})
	if err != nil {
		logger.Errorf("Failed to update Job status: %s", "Dealing")
		return
	}

	app, err := DefaultAppSvr.GetById(job.AppId)
	if err != nil {
		logger.Errorf("Failed to get app, id=%d, err=%s", job.AppId, err.Error())
		return
	}

	deploy := new(DeploymentEntity)
	err = DefaultDeploymentDao.GetById(job.DeploymentId, deploy)
	if err != nil {
		logger.Errorf("Failed to get deployment, deploymentId=%d, err=%s", job.DeploymentId, err.Error())
		return
	}

	// 2.1 start to deal deploy job
	err = NewJobFlow(deploy, job, app).Do()

	// 3.1 update job and deployment status
	jobStatus := JOB_STATUS_SUCC
	if err != nil {
		jobStatus = JOB_STATUS_FAILED
	}
	_, err = DefaultJobDao.UpdateById(job.Id, dbx.Params{
		"status": jobStatus,
	})
	if err != nil {
		logger.Errorf("Failed to update Job status: %s", "Dealing")
		return
	}

	// 4.1 try to update deployment status
	err = DefaultDeploymentSvr.TryUpdateDeploymentStatus(job.DeploymentId)
	if err != nil {
		logger.Errorf("Failed to UpdateDeploymentStatus, err=%s", err.Error())
	}
}

func (t *JobSvr) Start(jobId int64) error {
	_, err := DefaultJobDao.UpdateById(jobId, dbx.Params{
		"status": JOB_STATUS_READY,
	})
	return err
}
