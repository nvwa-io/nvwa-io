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

package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nvwa-io/nvwa-io/nvwa-server/daos"
	"github.com/nvwa-io/nvwa-io/nvwa-server/entities/vo"
	"github.com/nvwa-io/nvwa-io/nvwa-server/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/errs"
	. "github.com/nvwa-io/nvwa-io/nvwa-server/svrs"
)

type ClusterController struct {
	BaseAuthController
}

// @Title Create cluster
// @router / [post]
func (t *ClusterController) Create() {
	// json decode request
	req := new(vo.ReqCluster)
	err := t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// validate request params
	err = req.Valid()
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// check whether cluster exist
	ok, err := DefaultClusterSvr.IsExist(req.Cluster.AppId, req.Cluster.Name)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_CLUSTER_EXIST, req.Cluster.Name)
		return
	}

	// insert cluster
	req.Cluster.Uid = t.uid()
	id, err := DefaultClusterSvr.CreateByEntity(&req.Cluster)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"id": id,
	})
}

// @Title Update cluster
// @router /:cluster_id [put]
func (t *ClusterController) Update() {
	id, err := t.GetInt64(":cluster_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":cluster_id")
		return
	}

	// json decode request
	req := new(vo.ReqCluster)
	err = t.ReadRequestJson(&req)
	if err != nil {
		t.FailJson(errs.ERR_PARAM, err.Error())
		return
	}

	// validate request params
	valid := validation.Validation{}
	valid.Required(req.Cluster.Name, "name").Message(lang.I("cluster.name.not.empty"))
	valid.Required(req.Cluster.Hosts, "hosts").Message(lang.I("cluster.hosts.not.empty"))
	if valid.HasErrors() {
		t.FailJson(errs.ERR_PARAM, valid.Errors[0].Error())
		return
	}

	// check whether cluster exist
	ok, err := DefaultClusterSvr.IsExist(req.Cluster.AppId, req.Cluster.Name, id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}
	if ok {
		t.FailJson(errs.ERR_CLUSTER_EXIST, req.Cluster.Name)
		return
	}

	// update cluster
	_, err = daos.DefaultClusterDao.UpdateById(id, dbx.Params{
		"name":  req.Cluster.Name,
		"hosts": req.Cluster.Hosts,
	})

	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}

// @Title Get cluster
// @router /:cluster_id [get]
func (t *ClusterController) Detail() {
	id, err := t.GetInt64(":cluster_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":cluster_id")
		return
	}

	entity, err := DefaultClusterSvr.GetById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"cluster": entity,
	})
}

// @Title Get env clusters
// @router /env/:env_id [get]
func (t *ClusterController) List() {
	envId, err := t.GetInt64(":env_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":env_id")
		return
	}

	list, err := DefaultClusterSvr.ListByEnvId(envId)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson(RespData{
		"list": list,
	})
}

// @Title Delete cluster
// @router /:cluster_id [delete]
func (t *ClusterController) Delete() {
	id, err := t.GetInt64(":cluster_id")
	if err != nil {
		t.FailJson(errs.ERR_PARAM, ":cluster_id")
		return
	}

	_, err = daos.DefaultClusterDao.DeleteById(id)
	if err != nil {
		t.FailJson(errs.ERR_OPERATE, err.Error())
		return
	}

	t.SuccJson()
}
