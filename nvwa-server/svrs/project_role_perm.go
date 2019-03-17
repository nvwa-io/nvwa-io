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
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

var DefaultProjectRolePermSvr = new(ProjectRolePermSvr)

type ProjectRolePermSvr struct {
}

func (t *ProjectRolePermSvr) Create(projectRoleId int64, perm string) (int64, error) {
	id, err := DefaultProjectRolePermDao.CreateByMap(dbx.Params{
		"project_role_id": projectRoleId,
		"perm":            perm,
	})
	if err != nil {
		logger.Errorf("Failed to CreateByMap, err=%s", err.Error())
		return 0, err
	}

	return id, nil
}

func (t *ProjectRolePermSvr) IsExist(projectRoleId int64, perm string) (bool, error) {
	_, err := t.GetByProjectRoleIdPerm(projectRoleId, perm)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}

	return true, nil
}

func (t *ProjectRolePermSvr) GetByProjectRoleIdPerm(projectRoleId int64, perm string) (*ProjectRolePermEntity, error) {
	entity := new(ProjectRolePermEntity)
	err := DefaultProjectRolePermDao.GetOneByExp(dbx.HashExp{
		"project_role_id": projectRoleId,
		"perm":            perm,
	}, entity)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to GetByProjectRoleIdPerm, err=%s", err.Error())
		}

		return nil, err
	}

	return entity, nil
}

func (t *ProjectRolePermSvr) GetById(id int64) (*ProjectRolePermEntity, error) {
	p := new(ProjectRolePermEntity)
	err := DefaultProjectRolePermDao.GetById(id, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProjectRolePermSvr) ListAll(projectRoleId ...int64) ([]ProjectRolePermEntity, error) {
	var list = make([]ProjectRolePermEntity, 0)

	exp := dbx.HashExp{}
	if len(projectRoleId) > 0 {
		exp["project_role_id"] = projectRoleId[0]
	}

	err := DefaultProjectRolePermDao.GetAllByExp(exp, &list)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to ListAll, err=%s", err.Error())
		}

		return nil, err
	}

	return list, nil
}

func (t *ProjectRolePermSvr) DeleteById(id int64) error {
	_, err := DefaultProjectRolePermDao.DeleteById(id)
	if err != nil {
		logger.Errorf("Failed to DeleteById, err=%s", err.Error())
		return err
	}

	return nil
}

// Batch update project_role_id perms
// transaction operations:
// 1. delete all old records
// 2. insert all new records
func (t *ProjectRolePermSvr) BatchUpdate(projectRoleId int64, projectRoleName string, perms []string) error {
	tx, err := GetDb().Begin()
	if err != nil {
		return err
	}

	_, err = tx.Update(DefaultProjectRoleDao.Table(), dbx.Params{
		"name": projectRoleName,
	}, dbx.HashExp{"id": projectRoleId}).Execute()
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Delete(DefaultProjectRolePermDao.Table(), dbx.HashExp{"project_role_id": projectRoleId}).Execute()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, perm := range perms {
		_, err = tx.Insert(DefaultProjectRolePermDao.Table(), dbx.Params{
			"project_role_id": projectRoleId,
			"perm":            perm,
			"ctime":           libs.GetNow(),
		}).Execute()

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (t *ProjectRolePermSvr) GetAllProjectRoleIdMap() (map[int64][]ProjectRolePermEntity, error) {
	list := make([]ProjectRolePermEntity, 0)
	err := DefaultProjectRolePermDao.GetAllByExp(dbx.HashExp{}, &list)
	if err != nil {
		return nil, err
	}

	data := make(map[int64][]ProjectRolePermEntity)
	for _, v := range list {
		if _, ok := data[v.ProjectRoleId]; !ok {
			data[v.ProjectRoleId] = make([]ProjectRolePermEntity, 0)
		}

		data[v.ProjectRoleId] = append(data[v.ProjectRoleId], v)
	}

	return data, nil
}

// create project role and bind role permissions with transaction
func (t *ProjectRolePermSvr) CreateProjectRoleAndBindPerms(projectRoleName string, perms []string) error {
	tx, err := GetDb().Begin()
	if err != nil {
		return err
	}

	res, err := tx.Insert(DefaultProjectRoleDao.Table(), dbx.Params{
		"name":  projectRoleName,
		"ctime": libs.GetNow(),
	}).Execute()
	if err != nil {
		tx.Rollback()
		return err
	}
	projectRoleId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, perm := range perms {
		_, err = tx.Insert(DefaultProjectRolePermDao.Table(), dbx.Params{
			"project_role_id": projectRoleId,
			"perm":            perm,
			"ctime":           libs.GetNow(),
		}).Execute()

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
