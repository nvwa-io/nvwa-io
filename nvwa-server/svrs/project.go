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

var DefaultProjectSvr = new(ProjectSvr)

type ProjectSvr struct {
}

// Create project and member relation with transaction
func (t *ProjectSvr) Create(uid int64, name, description string) (int64, error) {
	tx, err := GetDb().Begin()
	if err != nil {
		return 0, err
	}

	// create project
	res, err := tx.Insert(DefaultProjectDao.Table(), dbx.Params{
		"uid":         uid,
		"name":        name,
		"description": description,
		"ctime":       libs.GetNow(),
	}).Execute()
	if err != nil {
		logger.Errorf("Failed to create project, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Errorf("Failed to get LastInterId, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}

	// add member relation
	_, err = tx.Insert(DefaultMemberDao.Table(), dbx.Params{
		"project_id":      id,
		"uid":             uid,
		"project_role_id": DefaultSystemSvr.Get().DefaultProjectRoleId,
		"ctime":           libs.GetNow(),
	}).Execute()
	if err != nil {
		logger.Errorf("Failed to add project member relation, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}

	// commit operations
	err = tx.Commit()
	if err != nil {
		logger.Errorf("Failed to commit create project, err=%s", err.Error())
		return 0, err
	}

	return id, err
}

func (t *ProjectSvr) GetByName(name string) (*ProjectEntity, error) {
	p := new(ProjectEntity)
	err := DefaultProjectDao.GetOneByExp(dbx.HashExp{
		"name": name,
	}, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProjectSvr) GetById(id int64) (*ProjectEntity, error) {
	p := new(ProjectEntity)
	err := DefaultProjectDao.GetById(id, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProjectSvr) GetByIds(ids []int64) (map[int64]ProjectEntity, error) {
	list := make([]ProjectEntity, 0)
	err := DefaultProjectDao.GetAllByIdsInt64(ids, &list)
	if err != nil {
		return nil, err
	}

	data := make(map[int64]ProjectEntity)
	for _, v := range list {
		data[v.Id] = v
	}

	return data, nil
}

func (t *ProjectSvr) IsExist(name string, excludeId ...int64) (bool, error) {
	p, err := t.GetByName(name)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}

		return false, nil
	}

	if len(excludeId) > 0 && p.Id == excludeId[0] {
		return false, nil
	}

	return true, nil
}

func (t *ProjectSvr) ListAllByUid(uid int64) ([]ProjectEntity, error) {
	var list = make([]ProjectEntity, 0)
	err := DefaultProjectDao.GetAllByExp(dbx.HashExp{
		"uid": uid,
	}, &list)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to ListAllByUid, err=%s", err.Error())
		}
		return nil, err
	}

	return list, nil
}

func (t *ProjectSvr) DeleteById(id int64) error {
	_, err := DefaultProjectDao.DeleteById(id)
	if err != nil {
		logger.Errorf("Failed to DeleteById, err=%s", err.Error())
		return err
	}

	return nil
}
