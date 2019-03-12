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
	"github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

var (
	DefaultSystemSvr = new(SystemSvr)
	system           *entities.SystemEntity
)

type SystemSvr struct{}

// @TODO reset system while change system configuration
func (t *SystemSvr) Get() *entities.SystemEntity {
	if system != nil {
		return system
	}

	system = new(entities.SystemEntity)
	err := daos.DefaultSystemDao.GetOneByExp(dbx.HashExp{
		"enabled": daos.ENABLED,
	}, system)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Errorf("No System configuration in table system.")
		}

		panic(err.Error())
	}

	return system
}
