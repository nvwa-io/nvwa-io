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

package oss_client

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"io/ioutil"
	"strings"
)

func C() *Client {
	return new(Client)
}

type OssClientConfig struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
	Bucket       string
}

type Client struct {
	c *OssClientConfig
}

func (t *Client) Config(c *OssClientConfig) *Client {
	t.c = c
	return t
}

// get oss bucket
func (t *Client) bucket() (*oss.Bucket, error) {
	client, err := oss.New(t.c.Endpoint, t.c.AccessKey, t.c.AccessSecret)
	if err != nil {
		logger.Errorf("New oss instance failed: %s", err.Error())
		return nil, err
	}

	bucket, err := client.Bucket(t.c.Bucket)
	if err != nil {
		logger.Errorf("Put Object From File failed: %s", err.Error())
	}

	return bucket, err
}

// get file md5 from oss object ETag
func (t *Client) GetObjectMd5(key string) (string, error) {
	bucket, err := t.bucket()
	if err != nil {
		return "", err
	}

	header, err := bucket.GetObjectMeta(key)
	if err != nil {
		logger.Errorf("Failed to GetObjectMeta: %s", err.Error())
		return "", err
	}

	if _, ok := header["ETag"]; !ok {
		info := "Header lack of field ETag"
		logger.Errorf(info)
		return "", errors.New(info)
	}

	return strings.ToLower(header.Get("ETag")), nil
}

// upload local file to oss
// http://t.Endpoint/key
func (t *Client) Put(filePath, key string) error {
	bucket, err := t.bucket()
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(key, filePath)
	if err != nil {
		logger.Errorf("Put Object From File failed: %s", err.Error())
	}

	return err
}

// get file content
func (t *Client) Get(key string) ([]byte, error) {
	bucket, err := t.bucket()
	if err != nil {
		return nil, err
	}

	reader, err := bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

// delete oss file
func (t *Client) Delete(key string) error {
	bucket, err := t.bucket()
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(key)
	if err != nil {
		logger.Errorf("Delete Object fail: %s", err.Error())
		return err
	}

	return err
}
