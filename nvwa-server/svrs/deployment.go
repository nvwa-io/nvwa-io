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
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"
	"strings"
)

var DefaultDeploymentSvr = new(DeploymentSvr)

type DeploymentSvr struct{}

// create deployment and init deployment jobs
func (t *DeploymentSvr) CreateAndInitJob(entity *DeploymentEntity) (int64, error) {
	// 1.1 full fill entity
	err := t.fullFillEntity(entity)
	if err != nil {
		return 0, err
	}

	// 1.2 prepare deployment params
	b, _ := json.Marshal(entity)
	deploymentParams := dbx.Params{}
	err = json.Unmarshal(b, &deploymentParams)
	if err != nil {
		logger.Errorf("Failed to unmashal entity to dbx.Params{}, err=%s", err.Error())
		return 0, err
	}
	deploymentParams["ctime"] = libs.GetNow()

	// 2.1 transaction to create deployment and jobs
	tx, err := GetDb().Begin()
	if err != nil {
		logger.Errorf("Failed to Begin transaction, err=%s", err.Error())
		return 0, err
	}

	// 2.1.1 insert deployment
	res, err := tx.Insert(DefaultDeploymentDao.Table(), deploymentParams).Execute()
	if err != nil {
		tx.Rollback()
		logger.Errorf("Failed to Insert deployment, err=%s", err.Error())
		return 0, err
	}

	deploymentId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		logger.Errorf("Failed to get deployment id, err=%s", err.Error())
		return 0, err
	}

	// 2.2 init deployment jobs
	clusters, err := DefaultClusterSvr.ListByEnvId(entity.EnvId)
	if err != nil {
		return 0, err
	}

	for _, v := range clusters {
		_, err := tx.Insert(DefaultJobDao.Table(), dbx.Params{
			"deployment_id": deploymentId,
			"app_id":        entity.AppId,
			"env_id":        entity.EnvId,
			"cluster_id":    v.Id,
			"all_hosts":     v.Hosts,
			"deploy_hosts":  v.Hosts,
			"exclude_hosts": "",
			"status":        JOB_STATUS_CREATED,
			"ctime":         libs.GetNow(),
		}).Execute()
		if err != nil {
			logger.Errorf("Failed to init job, err=%s", err.Error())
			tx.Rollback()
			return 0, err
		}
	}

	// 3.1 commit transaction
	err = tx.Commit()
	if err != nil {
		logger.Errorf("Failed to commit create deployment, err=%s", err.Error())
		return 0, err
	}

	return deploymentId, nil
}

func (t *DeploymentSvr) fullFillEntity(entity *DeploymentEntity) error {
	// 1.1 get pkg info
	pkg, err := DefaultPkgSvr.GetById(entity.PkgId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NoFound(err, lang.I("pkg.not.found"))
		}
		return err
	}

	// 1.2 get env
	env, err := DefaultEnvSvr.GetById(entity.EnvId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errs.NoFound(err, lang.I("env.not.found"))
		}
		return err
	}

	// 1.3 get cluster by envId
	clusters, err := DefaultClusterSvr.ListByEnvId(entity.EnvId)
	if err != nil {
		return err
	}

	clusterIds := make([]string, 0)
	clusterHosts := make(map[int64][]string)
	for _, v := range clusters {
		clusterIds = append(clusterIds, fmt.Sprintf("%d", v.Id))

		if v.Hosts == "" {
			return errors.New(lang.I("cluster.hosts.not.empty") + ":" + v.Name)
		}
		clusterHosts[v.Id] = strings.Split(v.Hosts, ",")
	}
	b, _ := json.Marshal(clusterHosts)

	// 2. full fill deployment fields
	entity.Pkg = pkg.Name
	entity.ClusterHosts = string(b)
	if len(strings.Split(entity.ClusterIds, ",")) != len(clusterIds) {
		entity.IsAllCluster = false
	} else {
		entity.IsAllCluster = true
	}

	entity.Branch = pkg.Branch
	entity.CommitId = pkg.CommitId
	entity.IsAutoDeploy = env.IsAutoDeploy
	entity.IsNeedAudit = env.IsNeedAudit
	entity.Status = DEPLOYMENT_STATUS_CREATED

	return nil
}

func (t *DeploymentSvr) ListPageDetailByAppIds(appIds []int64, page, pagesize int) ([]map[string]interface{}, error) {
	// 1.1 deployment list
	list := make([]DeploymentEntity, 0)
	err := DefaultDeploymentDao.GetDefaultPageList(&list, dbx.HashExp{
		"app_id": libs.SliceInt64ToSliceIntf(appIds),
	}, page, pagesize, true)
	if err != nil {
		return nil, err
	}

	// 2.1 deployment jobs
	deploymentIds := make([]int64, 0)
	for _, v := range list {
		deploymentIds = append(deploymentIds, v.Id)
	}

	jobs, err := DefaultJobSvr.GetByDeploymentIds(deploymentIds)
	if err != nil {
		return nil, err
	}

	// 2.2 get job cluster
	clusterIds := make([]int64, 0)
	for _, v := range jobs {
		for _, j := range v {
			clusterIds = append(clusterIds, j.ClusterId)
		}
	}
	clusters, err := DefaultClusterSvr.GetByIds(clusterIds)
	if err != nil {
		return nil, err
	}

	// 3.1 get apps
	apps, err := DefaultAppSvr.GetByIds(appIds)
	if err != nil {
		return nil, err
	}

	// 4.1 format return list
	//[{
	//    "app":AppEntity,
	//    "deployment":DeploymentEntity,
	//    "jobs":[{
	//        "job": JobEntity,
	//        "cluster": ClusterEntity
	//    }]
	//}]
	res := make([]map[string]interface{}, 0)
	for _, v := range list {
		tmpJobs := make([]map[string]interface{}, 0)
		for _, j := range jobs[v.Id] {
			tmpJobs = append(tmpJobs, map[string]interface{}{
				"job":     j,
				"cluster": clusters[j.ClusterId],
			})
		}

		tmp := map[string]interface{}{
			"app":        apps[v.AppId],
			"deployment": v,
			"jobs":       tmpJobs,
		}

		res = append(res, tmp)
	}

	return res, nil
}

func (t *DeploymentSvr) CountByAppIds(appIds []int64) (int, error) {
	return DefaultDeploymentDao.GetDefaultTotal(dbx.HashExp{
		"app_id": libs.SliceInt64ToSliceIntf(appIds),
	})
}

// @TODO try to update deployment status
func (t *DeploymentSvr) TryUpdateDeploymentStatus(id int64) error {

	return nil
}

func (t *DeploymentSvr) GetByIds(ids []int64) (map[int64]DeploymentEntity, error) {
	list := make([]DeploymentEntity, 0)
	err := DefaultDeploymentDao.GetAllByIdsInt64(ids, &list)
	if err != nil {
		return nil, err
	}

	data := make(map[int64]DeploymentEntity)
	for _, v := range list {
		data[v.Id] = v
	}

	return data, nil
}
