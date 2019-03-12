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

var DefaultClusterSvr = new(ClusterSvr)

type ClusterSvr struct{}

func (t *ClusterSvr) IsExist(appId int64, name string, excludeId ...int64) (bool, error) {
	p, err := t.GetByName(appId, name)
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

func (t *ClusterSvr) GetByName(appId int64, name string) (*ClusterEntity, error) {
	entity := new(ClusterEntity)
	err := DefaultClusterDao.GetOneByExp(dbx.HashExp{
		"app_id": appId,
		"name":   name,
	}, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (t *ClusterSvr) CreateByEntity(entity *ClusterEntity) (int64, error) {
	id, err := DefaultClusterDao.Create(entity)
	if err != nil {
		logger.Errorf("Failed to CreateByEntity, entity=%v, err=%s", entity, err.Error())
		return 0, err
	}

	return id, nil
}

func (t *ClusterSvr) GetById(id int64) (*ClusterEntity, error) {
	entity := new(ClusterEntity)
	err := DefaultClusterDao.GetById(id, entity)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to GetById, err=%s", err.Error())
		}

		return nil, err
	}

	return entity, nil
}

func (t *ClusterSvr) ListByEnvId(envId int64) ([]ClusterEntity, error) {
	list := make([]ClusterEntity, 0)
	err := DefaultClusterDao.GetAllByExp(dbx.HashExp{
		"env_id": envId,
	}, &list)

	if err != nil {
		logger.Errorf("Failed to ListByEnvId, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}

func (t *ClusterSvr) GetByEnvIds(envIds []int64) (map[int64][]ClusterEntity, error) {
	list := make([]ClusterEntity, 0)
	err := DefaultClusterDao.GetAllByFieldInt64("env_id", envIds, &list)
	if err != nil {
		logger.Errorf("Failed to Get cluster of envIds=%v, err=%s", envIds, err.Error())
		return nil, err
	}

	res := make(map[int64][]ClusterEntity)
	for _, v := range list {
		if _, ok := res[v.EnvId]; !ok {
			res[v.EnvId] = make([]ClusterEntity, 0)
		}

		res[v.EnvId] = append(res[v.EnvId], v)
	}

	// full fill data
	for _, id := range envIds {
		if _, ok := res[id]; !ok {
			res[id] = make([]ClusterEntity, 0)
		}
	}

	return res, nil
}

func (t *ClusterSvr) GetByIds(ids []int64) (map[int64]ClusterEntity, error) {
	list := make([]ClusterEntity, 0)
	err := DefaultClusterDao.GetAllByIdsInt64(ids, &list)
	if err != nil {
		return nil, err
	}

	res := make(map[int64]ClusterEntity)
	for _, v := range list {
		res[v.Id] = v
	}

	return res, nil
}
