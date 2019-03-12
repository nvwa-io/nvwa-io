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
	"fmt"
	"github.com/go-ozzo/ozzo-dbx"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

var DefaultUserSvr = new(UserSvr)

type UserSvr struct {
}

func (t *UserSvr) GetByUsername(username string) (*UserEntity, error) {
	u := new(UserEntity)
	err := DefaultUserDao.GetOneByExp(dbx.HashExp{
		"username": username,
	}, u)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to GetByUsername, username=%s, err=%s", username, err.Error())
		}

		return nil, err
	}

	return u, nil
}

func (t *UserSvr) GetById(id int64) (*UserEntity, error) {
	u := new(UserEntity)
	err := DefaultUserDao.GetById(id, u)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to GetById, uid=%d, err=%s", id, err.Error())
		}

		return nil, err
	}

	return u, nil
}

func (t *UserSvr) GetByUsernamePassword(username, password string) (*UserEntity, error) {
	u := new(UserEntity)
	err := DefaultUserDao.GetOneByExp(dbx.HashExp{
		"username": username,
		"password": t.encodePassword(password),
	}, u)

	if err != nil {
		if err != sql.ErrNoRows {
			logger.Errorf("Failed to GetByUsernamePassword, username=%s, err=%s", username, err.Error())
		}

		return nil, err
	}

	return u, nil
}

func (t *UserSvr) Register(username, email, password string) (int64, error) {
	id, err := DefaultUserDao.CreateByMap(dbx.Params{
		"username":     username,
		"display_name": username,
		"email":        email,
		"password":     t.encodePassword(password),
		"role":         ROLE_USER,
	})
	if err != nil {
		logger.Errorf("Failed to insert user, err=%s", err.Error())
		return 0, err
	}

	return id, nil
}

func (t *UserSvr) encodePassword(password string) string {
	return libs.Md5Str(fmt.Sprintf("%s%s", password, "@@nvwa-io&&pwd@@"))
}

// Get by uid, and return map[uid]UserEntity
func (t *UserSvr) GetByUids(uids []int64) (map[int64]UserEntity, error) {
	users := make([]UserEntity, 0, len(uids))
	err := DefaultUserDao.GetAllByIdsInt64(uids, &users)
	if err != nil {
		logger.Errorf("Failed to GetByUids, err=%s", err.Error())
		return nil, err
	}

	res := make(map[int64]UserEntity)

	for _, u := range users {
		res[u.Id] = u
	}

	return res, nil
}

func (t *UserSvr) GetAll() ([]UserEntity, error) {
	list := make([]UserEntity, 0)
	err := DefaultUserDao.GetAllByExp(dbx.HashExp{"enabled": ENABLED}, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
