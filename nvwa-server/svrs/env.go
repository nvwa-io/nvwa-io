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
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

var DefaultEnvSvr = new(EnvSvr)

type EnvSvr struct{}

func (t *EnvSvr) IsExist(appId int64, name string, excludeId ...int64) (bool, error) {
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

func (t *EnvSvr) GetByName(appId int64, name string) (*EnvEntity, error) {
	entity := new(EnvEntity)
	err := DefaultEnvDao.GetOneByExp(dbx.HashExp{
		"app_id": appId,
		"name":   name,
	}, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (t *EnvSvr) CreateAndInitDefaultCluster(entity *EnvEntity) (int64, error) {
	tx, err := GetDb().Begin()
	res, err := tx.Insert(DefaultEnvDao.Table(), dbx.Params{
		"uid":             entity.Uid,
		"app_id":          entity.AppId,
		"name":            entity.Name,
		"permit_branches": entity.PermitBranches,
		"is_need_audit":   entity.IsNeedAudit,
		"cmd_env":         entity.CmdEnv,
		"ctime":           libs.GetNow(),
	}).Execute()
	if err != nil {
		logger.Errorf("Failed to insert env=%v, err=%s", *entity, err.Error())
		tx.Rollback()
		return 0, err
	}

	envId, err := res.LastInsertId()
	if err != nil {
		logger.Errorf("Failed to get lastInsertId, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Insert(DefaultClusterDao.Table(), dbx.Params{
		"app_id": entity.AppId,
		"env_id": envId,
		"uid":    entity.Uid,
		"name":   lang.I("cluster.default"),
		"hosts":  "",
		"ctime":  libs.GetNow(),
	}).Execute()
	if err != nil {
		logger.Errorf("Failed to create default cluster, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		logger.Errorf("Failed to commit insert env and default cluster, err=%s", err.Error())
		tx.Rollback()
		return 0, err
	}

	return envId, nil
}

func (t *EnvSvr) CreateByEntity(entity *EnvEntity) (int64, error) {
	id, err := DefaultEnvDao.Create(entity)
	if err != nil {
		logger.Errorf("Failed to CreateByEntity, entity=%v, err=%s", entity, err.Error())
		return 0, err
	}

	return id, nil
}

func (t *EnvSvr) GetById(id int64) (*EnvEntity, error) {
	entity := new(EnvEntity)
	err := DefaultEnvDao.GetById(id, entity)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to GetById, err=%s", err.Error())
		}

		return nil, err
	}

	return entity, nil
}

func (t *EnvSvr) GetByIds(ids []int64) (map[int64]EnvEntity, error) {
	list := make([]EnvEntity, 0)
	err := DefaultEnvDao.GetAllByIdsInt64(ids, &list)
	if err != nil {
		return nil, err
	}

	res := make(map[int64]EnvEntity)
	for _, v := range list {
		res[v.Id] = v
	}

	return res, nil
}

func (t *EnvSvr) ListAllByAppId(appId int64) ([]EnvEntity, error) {
	list := make([]EnvEntity, 0)
	err := DefaultEnvDao.GetAllByExp(dbx.HashExp{
		"app_id": appId,
	}, &list)

	if err != nil {
		logger.Errorf("Failed to ListAllByAppId, err=%s", err.Error())
		return nil, err
	}

	return list, nil
}

func (t *EnvSvr) GetByAppIds(appIds []int64) (map[int64][]EnvEntity, error) {
	list := make([]EnvEntity, 0)
	err := DefaultEnvDao.GetAllByFieldInt64("app_id", appIds, &list)
	if err != nil {
		return nil, err
	}

	// group by appId
	res := make(map[int64][]EnvEntity)
	for _, v := range list {
		if _, ok := res[v.AppId]; !ok {
			res[v.AppId] = make([]EnvEntity, 0)
		}

		res[v.AppId] = append(res[v.AppId], v)
	}

	return res, nil
}
