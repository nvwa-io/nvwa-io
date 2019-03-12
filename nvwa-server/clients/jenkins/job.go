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

package jenkins

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
	"net/http"
	"strconv"
	"strings"
)

const (
	PATH_CREAT                = "/createItem/"
	PATH_DELETE               = "/job/%s/doDelete"                   // /job/{name}/deDelete
	PATH_UPDATE               = "/job/%s/config.xml"                 // appName
	PATH_VIEW_JOB             = "/job/%s/"                           // appName
	PATH_BUILD_WITHPARAMETERS = "/job/%s/buildWithParameters"        // appName
	PATH_LOG_TEXT             = "/job/%s/%d/logText/progressiveText" // appName, jenkins job num
	PATH_JOB_JSON             = "/job/%s/%d/api/json"                // appName, jenkins job num
)

func NewJob(c *Client) *Job {
	return &Job{
		c: c,
	}
}

type (
	Job struct {
		c *Client
	}

	requestParam struct {
		Method string
		Url    string
		Header map[string]string
		Query  map[string]string
		Body   interface{}
	}
)

func (t *Job) request(rp requestParam) (*resty.Response, error) {
	jenkinsCrumb, err := t.c.CrumbIssuer()
	if err != nil {
		return nil, err
	}

	// jenkins crsf header
	if rp.Header == nil {
		rp.Header = make(map[string]string)
	}
	rp.Header[jenkinsCrumb[0]] = jenkinsCrumb[1]

	// init resty request
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))
	req := resty.R().
		SetHeaders(rp.Header).
		SetQueryParams(rp.Query).
		SetBody(rp.Body).
		SetBasicAuth(t.c.Username, t.c.Password)

	// send http request
	var res *resty.Response
	switch rp.Method {
	case http.MethodGet:
		res, err = req.Get(rp.Url)

	case http.MethodPost:
		res, err = req.Post(rp.Url)

	case http.MethodHead:
		res, err = req.Head(rp.Url)

	default:
		return nil, errors.New(fmt.Sprintf("Http method not supported: %s.", rp.Method))
	}

	return res, err
}

func (t *Job) Create(name, xmlConfig string) error {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), PATH_CREAT)
	res, err := t.request(requestParam{
		Method: http.MethodPost,
		Url:    url,
		Header: map[string]string{
			"Content-Type": "application/xml",
		},
		Query: map[string]string{
			"name": name,
		},
		Body: xmlConfig,
	})
	if err != nil {
		return err
	}

	if res.StatusCode() < 200 || res.StatusCode() > 299 {
		info := fmt.Sprintf("Failed to create %s, status code: %d, response: %s", name, res.StatusCode(), string(res.Body()))
		return errors.New(info)
	}

	return nil
}

func (t *Job) Update(name, xmlConfig string) error {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), fmt.Sprintf(PATH_UPDATE, name))
	res, err := t.request(requestParam{
		Method: http.MethodPost,
		Url:    url,
		Header: map[string]string{
			"Content-Type": "application/xml",
		},
		Body: xmlConfig,
	})
	if err != nil {
		return err
	}

	if res.StatusCode() < 200 || res.StatusCode() > 299 {
		info := fmt.Sprintf("Failed to update %s, status code: %d, response: %s", name, res.StatusCode(), string(res.Body()))
		return errors.New(info)
	}

	return nil
}

func (t *Job) Delete(name string) error {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), fmt.Sprintf(PATH_DELETE, name))
	res, err := t.request(requestParam{
		Method: http.MethodPost,
		Url:    url,
	})
	if err != nil {
		return err
	}

	if res.StatusCode() < 200 || res.StatusCode() > 299 {
		info := fmt.Sprintf("Failed to delete %s, status code: %d, response: %s", name, res.StatusCode(), string(res.Body()))
		return errors.New(info)
	}

	return nil
}

func (t *Job) IsExist(name string) (bool, error) {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), fmt.Sprintf(PATH_VIEW_JOB, name))
	res, err := t.request(requestParam{
		Method: http.MethodGet,
		Url:    url,
	})
	if err != nil {
		return false, err
	}

	if res.StatusCode() >= 200 && res.StatusCode() <= 299 {
		return true, nil
	} else if res.StatusCode() == 404 {
		return false, nil
	} else {
		info := fmt.Sprintf("Failed to get %s, status code: %d, response: %s", name, res.StatusCode(), string(res.Body()))
		return false, errors.New(info)
	}
}

// buildParams = {"BUILD_ID": "1", "BUILD_BRANCH": "master"}
func (t *Job) BuildWithParameters(name string, buildParams map[string]string) error {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), fmt.Sprintf(PATH_BUILD_WITHPARAMETERS, name))
	res, err := t.request(requestParam{
		Method: http.MethodPost,
		Url:    url,
		Query:  buildParams,
	})
	if err != nil {
		return err
	}
	if res.StatusCode() < 200 || res.StatusCode() > 299 {
		info := fmt.Sprintf("Failed to BuildWithParameters %s, status code: %d, response: %s", name, res.StatusCode(), string(res.Body()))
		return errors.New(info)
	}

	return nil
}

func (t *Job) GetLogSize(name string, buildNum int) (int, error) {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), fmt.Sprintf(PATH_LOG_TEXT, name, buildNum))
	res, err := t.request(requestParam{
		Method: http.MethodHead,
		Url:    url,
	})
	if err != nil {
		return 0, err
	}
	logSize, err := strconv.Atoi(res.Header().Get("X-Text-Size"))
	if err != nil {
		return 0, err
	}

	return logSize, nil
}

// get log of the job
// jenkins job name, jenkins job build number, (limit[0]:start, limit[1]logSize)
func (t *Job) GetTextLog(name string, buildNum int, limit ...int) (string, error) {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), fmt.Sprintf(PATH_LOG_TEXT, name, buildNum))
	start := 0
	switch len(limit) {
	case 2:
		// recalculate start
		totalBytes, _ := t.GetLogSize(name, buildNum)
		if totalBytes > 0 &&
			limit[1] < totalBytes {
			start = totalBytes - limit[1]
		}
	case 1:
		start = limit[0]
	default:
		start = 0
	}

	res, err := t.request(requestParam{
		Method: http.MethodGet,
		Url:    url,
		Query: map[string]string{
			"start": strconv.Itoa(start),
		},
	})
	if err != nil {
		return "", err
	}

	return string(res.Body()), nil
}

// get job status
func (t *Job) GetStatus(name string, buildNum int) (string, bool, error) {
	url := fmt.Sprintf("%s%s", strings.TrimRight(t.c.Domain, "/"), fmt.Sprintf(PATH_JOB_JSON, name, buildNum))

	res, err := t.request(requestParam{
		Method: http.MethodGet,
		Url:    url,
		Query: map[string]string{
			"tree": "result,building",
		},
	})
	if err != nil {
		return "", true, err
	}

	if res.StatusCode() < 200 || res.StatusCode() > 299 {
		return "", true, errors.New(fmt.Sprintf("Failed to request %s, status code=%d, response=%s", url, res.StatusCode(), string(res.Body())))
	}

	type statusResp struct {
		Building bool   `json:"building"`
		Result   string `json:"result"`
	}
	statusR := new(statusResp)
	err = json.Unmarshal(res.Body(), statusR)
	if err != nil {
		return "", true, errors.New(fmt.Sprintf("Failed to json decode response, url=%s, err=%s, body=%s", url, err.Error(), string(res.Body())))
	}

	return statusR.Result, statusR.Building, nil
}
