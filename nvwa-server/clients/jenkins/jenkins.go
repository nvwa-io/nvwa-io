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
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
)

func C() *Client {
	return new(Client)
}

type Client struct {
	// e.g: http://localhost:8080
	Domain string

	// jenkins basic auth username and password
	Username string
	Password string
}

func (t *Client) Config(domain, username, password string) *Client {
	t.Domain = domain
	t.Username = username
	t.Password = password
	return t
}

// Jenkins-Crumb:9dc7abb9785492b95f781a4de469e600
// for jenkins CSRF Protection
func (t *Client) CrumbIssuer() ([]string, error) {
	url := fmt.Sprintf(strings.TrimRight(t.Domain, "/") + "/crumbIssuer/api/xml?xpath=concat(//crumbRequestField,\":\",//crumb)")
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))
	res, err := resty.R().SetBasicAuth(t.Username, t.Password).Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() >= 200 && res.StatusCode() <= 299 {
		headerCrumb := strings.Split(string(res.Body()), ":")
		if len(headerCrumb) != 2 {
			return nil, errors.New("Invalid jenkins crumb: " + string(res.Body()))
		}
		return headerCrumb, nil
	} else if res.StatusCode() == 404 {
		// CSRF Protection
		// jenkins has unchecked "Prevent Cross Site Request Forgery exploits"
		return []string{"NoCrumbIssuer", "NoCrumbIssuer"}, nil
	} else {
		// HTTP request failed
		info := fmt.Sprintf("Failed to get CrumbIssuer, status code: %d, output: %s",
			res.StatusCode(),
			string(res.Body()))

		return nil, errors.New(info)
	}
}

func (t *Client) Job() *Job {
	return NewJob(t)
}
