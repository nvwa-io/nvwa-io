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

var DefaultProjectRoleSvr = new(ProjectRoleSvr)

type ProjectRoleSvr struct {
}

func (t *ProjectRoleSvr) Create(name string) (int64, error) {
	id, err := DefaultProjectRoleDao.CreateByMap(dbx.Params{
		"name": name,
	})
	if err != nil {
		logger.Errorf("Failed to CreateByMap, err=%s", err.Error())
		return 0, err
	}

	return id, nil
}

func (t *ProjectRoleSvr) GetByName(name string) (*ProjectRoleEntity, error) {
	p := new(ProjectRoleEntity)
	err := DefaultProjectRoleDao.GetOneByExp(dbx.HashExp{
		"name": name,
	}, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProjectRoleSvr) IsExist(name string, excludeId ...int64) (bool, error) {
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

func (t *ProjectRoleSvr) GetById(id int64) (*ProjectRoleEntity, error) {
	p := new(ProjectRoleEntity)
	err := DefaultProjectRoleDao.GetById(id, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProjectRoleSvr) ListAll() ([]ProjectRoleEntity, error) {
	var list = make([]ProjectRoleEntity, 0)
	err := DefaultProjectRoleDao.GetAllByExp(dbx.HashExp{}, &list)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to ListAllByUid, err=%s", err.Error())
		}

		return nil, err
	}

	return list, nil
}

// list all project roles with map[project_role_id]ProjectRoleEntity
func (t *ProjectRoleSvr) ListMapAll() (map[int64]ProjectRoleEntity, error) {
	list, err := t.ListAll()
	if err != nil {
		return nil, err
	}

	res := make(map[int64]ProjectRoleEntity)
	for _, r := range list {
		res[r.Id] = r
	}

	return res, nil
}

func (t *ProjectRoleSvr) DeleteById(id int64) error {
	_, err := DefaultProjectRoleDao.DeleteById(id)
	if err != nil {
		logger.Errorf("Failed to DeleteById, err=%s", err.Error())
		return err
	}

	return nil
}
