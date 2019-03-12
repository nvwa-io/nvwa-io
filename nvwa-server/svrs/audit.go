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
	"errors"
	"github.com/go-ozzo/ozzo-dbx"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"strconv"
	"time"
)

var DefaultAuditSvr = new(AuditSvr)

type AuditSvr struct{}

func (t *AuditSvr) GetByUid(uid int64, page, pagesize int) ([]AuditEntity, error) {
	list := make([]AuditEntity, 0)
	err := DefaultAuditDao.GetDefaultPageList(&list, dbx.HashExp{"uid": uid}, page, pagesize, true)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t *AuditSvr) GetCountByUid(uid int64) (int, error) {
	total, err := DefaultAuditDao.GetDefaultTotal(dbx.HashExp{
		"uid": uid,
	})
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (t *AuditSvr) GetByProjectIdsAndStatus(projectIds []int64, status []int, page, pagesize int, datetime ...string) ([]AuditEntity, error) {
	list := make([]AuditEntity, 0)
	query := t.formatQuery("*", projectIds, status, datetime...)

	offset := (page - 1) * pagesize
	query.Offset(int64(offset)).Limit(int64(pagesize))
	err := query.All(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t *AuditSvr) FullFillList(list []AuditEntity, uid int64) ([]map[string]interface{}, error) {
	if len(list) == 0 {
		return []map[string]interface{}{}, nil
	}

	deploymentIds := make([]int64, 0)
	projectIds := make([]int64, 0)
	appIds := make([]int64, 0)
	uids := make([]int64, 0)
	envIds := make([]int64, 0)

	for _, v := range list {
		deploymentIds = append(deploymentIds, v.DeploymentId)
		projectIds = append(projectIds, v.ProjectId)
		appIds = append(appIds, v.AppId)
		envIds = append(envIds, v.EnvId)
		uids = append(uids, v.Uid, v.AuditUid)
	}

	deployments, err := DefaultDeploymentSvr.GetByIds(deploymentIds)
	if err != nil {
		return nil, err
	}
	projects, err := DefaultProjectSvr.GetByIds(projectIds)
	if err != nil {
		return nil, err
	}

	apps, err := DefaultAppSvr.GetByIds(appIds)
	if err != nil {
		return nil, err
	}

	users, err := DefaultUserSvr.GetByUids(uids)
	if err != nil {
		return nil, err
	}

	envs, err := DefaultEnvSvr.GetByIds(envIds)
	if err != nil {
		return nil, err
	}

	// whether current user has permission to audit
	projectIdPermAudit, err := DefaultMemberSvr.GetProjectPermAuditByProjectIdsUid(projectIds, uid)
	if err != nil {
		return nil, err
	}

	data := make([]map[string]interface{}, 0)
	for _, v := range list {
		tmp := map[string]interface{}{
			"audit":          v,
			"deployment":     deployments[v.DeploymentId],
			"project":        projects[v.ProjectId],
			"app":            apps[v.AppId],
			"user":           users[v.Uid],
			"audit_user":     users[v.AuditUid],
			"env":            envs[v.EnvId],
			"has_perm_audit": false,
		}
		if _, ok := projects[v.ProjectId]; ok {
			if _, hasPerm := projectIdPermAudit[v.ProjectId]; hasPerm {
				tmp["has_perm_audit"] = projectIdPermAudit[v.ProjectId]
			}
		}
		data = append(data, tmp)
	}

	return data, nil
}

func (t *AuditSvr) formatQuery(selectStr string, projectIds []int64, status []int, datetime ...string) *dbx.SelectQuery {
	intfProjectids := make([]interface{}, 0)
	intfStatus := make([]interface{}, 0)

	for _, v := range projectIds {
		intfProjectids = append(intfProjectids, v)
	}
	for _, v := range status {
		intfStatus = append(intfStatus, v)
	}

	query := GetDb().Select(selectStr).From(DefaultAuditDao.Table()).
		Where(dbx.HashExp{
			"project_id": intfProjectids,
			"status":     intfStatus,
			"enabled":    ENABLED,
		})
	if len(datetime) > 0 {
		query.AndWhere(dbx.NewExp("ctime > {:datetime}", dbx.Params{"datetime": datetime[0]}))
	} else {
		// if not set time limit, default: 1 month
		query.AndWhere(dbx.NewExp("ctime > {:datetime}", dbx.Params{"datetime": libs.Date("Y-m-d H:i:s", time.Now().Unix()-30*24*3600)}))
	}

	return query
}

func (t *AuditSvr) GetCountByProjectIdsAndStatus(projectIds []int64, status []int, datetime ...string) (int, error) {
	query := t.formatQuery("count(*) as total", projectIds, status, datetime...)
	var ret dbx.NullStringMap
	err := query.One(&ret)
	if err != nil {
		return 0, err
	}

	if !ret["total"].Valid {
		return 0, errors.New("Failed to count(*)")
	}

	return strconv.Atoi(ret["total"].String)
}

func (t *AuditSvr) UpdateStatusById(id int64, status int, auditUid int64) error {
	// get audit
	audit := new(AuditEntity)
	err := DefaultAuditDao.GetById(id, audit)
	if err != nil {
		return err
	}

	// get deployment
	deployment := new(DeploymentEntity)
	err = DefaultDeploymentDao.GetById(audit.DeploymentId, deployment)
	if err != nil {
		return err
	}

	// update audit status and deployment status
	tx, err := GetDb().Begin()
	if err != nil {
		return err
	}

	_, err = tx.Update(DefaultAuditDao.Table(), dbx.Params{
		"status":    status,
		"audit_uid": auditUid,
	}, dbx.HashExp{"id": id}).Execute()
	if err != nil {
		tx.Rollback()
		logger.Errorf("Failed to update audit status, err=%s", err.Error())
		return err
	}

	if deployment.Status == DEPLOYMENT_STATUS_WAIT_AUDIT ||
		deployment.Status == DEPLOYMENT_STATUS_CREATED {
		_, err = tx.Update(DefaultDeploymentDao.Table(), dbx.Params{
			"status": status,
		}, dbx.HashExp{"id": deployment.Id}).Execute()
		if err != nil {
			tx.Rollback()
			logger.Errorf("Failed to update deployment status, err=%s", err.Error())
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		logger.Errorf("Failed to update audit and deployment status, err=%s", err.Error())
		return err
	}

	return nil
}
