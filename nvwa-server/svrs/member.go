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
	"github.com/go-ozzo/ozzo-dbx"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

var DefaultMemberSvr = new(MemberSvr)

type MemberSvr struct {
}

func (t *MemberSvr) IsExist(projectId, uid int64) (bool, error) {
	_, err := t.GetByProjectIdUid(projectId, uid)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func (t *MemberSvr) GetByProjectIdUid(projectId, uid int64) (*MemberEntity, error) {
	m := new(MemberEntity)
	err := DefaultMemberDao.GetOneByExp(dbx.HashExp{
		"project_id": projectId,
		"uid":        uid,
	}, m)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to GetByProjectIdUid, err=%s", err.Error())
		}

		return nil, err
	}

	return m, nil
}

func (t *MemberSvr) GetProjectPermAuditByProjectIdsUid(projectIds []int64, uid int64) (map[int64]bool, error) {
	// user's audit permission
	members, err := DefaultMemberSvr.GetByProjectIdsAndUid(projectIds, uid)
	if err != nil {
		return nil, err
	}
	projectRoleIds := make([]interface{}, 0)
	for _, v := range members {
		projectRoleIds = append(projectRoleIds, v.ProjectRoleId)
	}
	roleAuditList := make([]ProjectRolePermEntity, 0)
	err = GetDb().Select("*").From(DefaultProjectRolePermDao.Table()).
		Where(dbx.HashExp{
			"project_role_id": projectRoleIds,
			"perm":            PERM_ENV_AUDIT,
			"enabled":         ENABLED,
		}).All(&roleAuditList)
	if err != nil {
		return nil, err
	}

	mapRoleIdPermAudit := map[int64]ProjectRolePermEntity{}
	for _, v := range roleAuditList {
		mapRoleIdPermAudit[v.ProjectRoleId] = v
	}

	// whether current uid has permission to audit
	projectIdPermAudit := map[int64]bool{}
	for _, m := range members {
		if _, ok := mapRoleIdPermAudit[m.ProjectRoleId]; ok {
			projectIdPermAudit[m.ProjectId] = true
		} else {
			projectIdPermAudit[m.ProjectId] = false
		}
	}

	return projectIdPermAudit, nil
}

func (t *MemberSvr) GetByProjectIdsAndUid(projectIds []int64, uid int64) ([]MemberEntity, error) {
	intfProjectIds := make([]interface{}, 0)
	for _, v := range projectIds {
		intfProjectIds = append(intfProjectIds, v)
	}

	list := make([]MemberEntity, 0)
	err := GetDb().Select("*").From(DefaultMemberDao.Table()).
		Where(dbx.HashExp{
			"project_id": intfProjectIds,
			"uid":        uid,
		}).All(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t *MemberSvr) Add(projectId, uid, projectRoleId int64) (int64, error) {
	id, err := DefaultMemberDao.Create(dbx.Params{
		"project_id":      projectId,
		"uid":             uid,
		"project_role_id": projectRoleId,
	})

	if err != nil {
		logger.Errorf("Failed to add member, err=%s", err.Error())
		return 0, err
	}

	return id, nil
}

func (t *MemberSvr) GetAllByProjectId(projectId int64) ([]MemberEntity, error) {
	list := make([]MemberEntity, 0)
	err := DefaultMemberDao.GetAllByExp(dbx.HashExp{
		"project_id": projectId,
	}, &list)
	if err != nil {
		logger.Errorf("Failed to GetAllByExp, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}

// get all members with member related info, e.g: user's detail and project role info
func (t *MemberSvr) GetAllDetailByProjectId(projectId int64) ([]map[string]interface{}, error) {
	list, err := t.GetAllByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	// get users
	uids := make([]int64, 0)
	for _, m := range list {
		uids = append(uids, m.Uid)
	}

	users, err := DefaultUserSvr.GetByUids(uids)
	if err != nil {
		return nil, err
	}

	// get project roles
	projectRoles, err := DefaultProjectRoleSvr.ListMapAll()
	if err != nil {
		return nil, err
	}

	res := make([]map[string]interface{}, 0)
	for _, m := range list {
		item := map[string]interface{}{
			"member":       m,
			"user":         users[m.Uid],
			"project_role": projectRoles[m.ProjectRoleId],
		}
		res = append(res, item)
	}

	return res, nil
}

func (t *MemberSvr) DeleteById(id int64) error {
	_, err := DefaultMemberDao.DeleteById(id)
	if err != nil {
		logger.Errorf("Failed to DeleteById, id=%d, err=%s", id, err.Error())
		return err
	}

	return nil
}

func (t *MemberSvr) GetProjectIdsByUid(uid int64) ([]int64, error) {
	list, err := t.GetAllByUid(uid)
	if err != nil {
		return nil, err
	}

	projectIds := make([]int64, 0)
	for _, v := range list {
		projectIds = append(projectIds, v.ProjectId)
	}

	return projectIds, nil
}

func (t *MemberSvr) GetAllByUid(uid int64) ([]MemberEntity, error) {
	list := make([]MemberEntity, 0)
	err := DefaultMemberDao.GetAllByExp(dbx.HashExp{
		"uid": uid,
	}, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t *MemberSvr) UpdateById(id int64, params dbx.Params) error {
	_, err := DefaultMemberDao.UpdateById(id, params)
	if err != nil {
		logger.Errorf("Failed to update member id=%d, err=%s", id, err.Error())
		return err
	}

	return nil
}

// BasePermission, e.g: Tester
// [project]
// 1. View project detail
// [app]
// 1. View apps / environments / clusters / deployments
// 2. Deploy apps
func (t *MemberSvr) HasBasePermission(uid, projectId int64) bool {

	return true
}

// MediumPermission, e.g: Developer
// [project]
// 1. View project detail
// [app]
// 1. View apps / environments / clusters / deployments
// 2. Edit apps / environments / clusters / deployments
// 3. Add apps / environments / clusters / deployments
// 4. Deploy apps
func (t *MemberSvr) HasMediumPermission(uid, projectId int64) bool {

	return true
}

// MediumPermission, e.g: Developer
// [project]
// 1. View project detail
// 2. Modify project info
// [app]
// 1. View apps / environments / clusters / deployments
// 2. Edit apps / environments / clusters / deployments
// 3. Add apps / environments / clusters / deployments
// 4. Delete apps / environments / clusters / deployments
// 5. Deploy apps
// 6. Config whether the app's environment need to be audited
// 7. Deployment audit
func (t *MemberSvr) HasAllPermission(uid, projectId int64) bool {

	return true
}
