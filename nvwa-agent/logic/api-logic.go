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

package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/libs/logger"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/option"
	"gopkg.in/resty.v1"
)

func NewApiLogic() *ApiLogic {
	resty.SetRetryCount(3)
	resty.FlexibleRedirectPolicy(3)
	return new(ApiLogic)
}

type ApiLogic struct {
	opt *option.Option
}

func (t *ApiLogic) Option(opt *option.Option) *ApiLogic {
	t.opt = opt
	return t
}

func (t *ApiLogic) checkApiResponse(res *resty.Response) error {
	if res.StatusCode() < 200 || res.StatusCode() > 299 {
		info := fmt.Sprintf("Reqeust failed, status code: %d, response: %s", res.StatusCode(), string(res.Body()))
		return errors.New(info)
	}

	// parse response data
	type response struct {
		Code int
		Msg  string
	}

	resp := new(response)
	err := json.Unmarshal(res.Body(), resp)
	if err != nil {
		info := fmt.Sprintf("Failed to decode reponse, err=%s, body=%s", err.Error(), string(res.Body()))
		logger.Errorf(info)
		return errors.New(info)
	}

	if resp.Code != 200 {
		info := fmt.Sprintf("Failed, code=%d, msg=%s", resp.Code, resp.Msg)
		logger.Errorf(info)
		return errors.New(info)
	}

	return nil
}

// update build info from nvwa-agent
func (t *ApiLogic) UpdateBuildInfo(buildId int64, params map[string]interface{}) error {
	res, err := resty.R().SetQueryParams(map[string]string{
		"build_id": fmt.Sprintf("%d", buildId),
		"hostname": libs.Hostname(),
	}).SetBody(params).Put(t.opt.Cfg.NvwaApi.BuildInfo)
	if err != nil {
		info := fmt.Sprintf("Request UpdateBuildInfo: %s failed, err=%s", t.opt.Cfg.NvwaApi.BuildInfo, err.Error())
		logger.Errorf(info)
		return errors.New(info)
	}

	return t.checkApiResponse(res)
}

// update version package to nvwa-server
func (t *ApiLogic) UploadPackage(buildId int64, pkgName, pkgFile string) error {
	res, err := resty.R().SetQueryParams(map[string]string{
		"build_id": fmt.Sprintf("%d", buildId),
		"hostname": libs.Hostname(),
	}).SetFile("pkgFile", pkgFile).Post(t.opt.Cfg.NvwaApi.UploadPackage)
	if err != nil {
		info := fmt.Sprintf("Request UploadPackage: %s failed, err=%s", t.opt.Cfg.NvwaApi.UploadPackage, err.Error())
		logger.Errorf(info)
		return errors.New(info)
	}

	return t.checkApiResponse(res)
}

// send notify to nvwa-server with message
func (t *ApiLogic) Notify() error {
	body := map[string]interface{}{
		"message": t.opt.Notify.Message,
		"users":   t.opt.Notify.NotifyUser,
	}
	res, err := resty.R().SetQueryParams(map[string]string{
		"build_id": fmt.Sprintf("%d", t.opt.Notify.BuildId),
		"hostname": libs.Hostname(),
	}).SetBody(body).Post(t.opt.Cfg.NvwaApi.Notify)
	if err != nil {
		info := fmt.Sprintf("Request Notify: %s failed, err=%s", t.opt.Cfg.NvwaApi.Notify, err.Error())
		logger.Errorf(info)
		return errors.New(info)
	}

	return t.checkApiResponse(res)
}
