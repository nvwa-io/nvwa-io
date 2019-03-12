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
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

var DefaultPkgSvr = new(PkgSvr)

type PkgSvr struct{}

func (t *PkgSvr) GetById(id int64) (*PkgEntity, error) {
	entity := new(PkgEntity)
	err := DefaultPkgDao.GetById(id, entity)
	if err != nil {
		logger.Errorf("Failed to GetById, err=%s", err.Error())
		return nil, err
	}

	return entity, nil
}

func (t *PkgSvr) GetLatestListByEnvId(appId int64, branch []string, storageType string, limit ...int) ([]PkgEntity, error) {
	cond := dbx.HashExp{
		"app_id":       appId,
		"storage_type": storageType,
		"enabled":      ENABLED,
	}

	// branch filter
	if len(branch) > 0 {
		inBranches := []interface{}{}
		for _, v := range branch {
			inBranches = append(inBranches, v)
		}
		cond["branch"] = inBranches
	}

	num := 20
	if len(limit) > 0 {
		num = limit[0]
	}

	list := make([]PkgEntity, 0)
	err := GetDb().Select("*").From(DefaultPkgDao.Table()).
		Where(cond).
		OrderBy("id DESC").
		Limit(int64(num)).
		All(&list)
	if err != nil {
		logger.Errorf("Failed to get latest pkg list, appId=%d, err=%s", appId, err.Error())
		return nil, err
	}

	return list, nil
}
