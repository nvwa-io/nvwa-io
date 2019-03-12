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

package ansible

import (
	"errors"
	"fmt"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs"
	"github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
	"os"
	"strings"
	"time"
)

const (
	ANSIBLE_CONCURRENCY        = 10
	ANSIBLE_CONNECTION_TIMEOUT = 10
	ANSIBLE_SSH_ARGS           = "ANSIBLE_SSH_ARGS='-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false'"
)

func C() *Client {
	return new(Client)
}

type Client struct{}

// get local ansible version
func (t *Client) Version() ([]byte, error) {
	cmd := "ansible --version"
	return libs.CmdExecShellDefault(cmd)
}

// execute shell
// hosts: ["192.168.20.12","192.168.20.11:22","192.168.20.10:444"]
func (t *Client) ExecShell(user, shell string, hosts []string, timeout int) ([]byte, string, error) {
	// gen inventory file
	inventoryFile, err := t.genInventoryFile(user, hosts)
	if err != nil {
		logger.Errorf("Failed to genInventoryFile, err=%s", err.Error())
		return nil, "", err
	}
	defer os.Remove(inventoryFile)

	// format ansible cmd
	cmd := fmt.Sprintf("%s ansible all -m shell -a %s -f %d -i %s -T %d",
		ANSIBLE_SSH_ARGS,
		libs.EscapeShellArg(shell+" warn=False"),
		ANSIBLE_CONCURRENCY,
		inventoryFile,
		ANSIBLE_CONNECTION_TIMEOUT,
	)

	// exec
	output, err := libs.CmdExecShell(cmd, timeout)
	if err != nil {
		logger.Errorf("Failed to execute: %s , err=%s", cmd, err.Error())
		return output, cmd, err
	}

	return output, cmd, nil
}

// copy file from src to remote host dest
// hosts: ["192.168.20.12","192.168.20.11:22","192.168.20.10:444"]
func (t *Client) CopyFile(user, src, dest string, hosts []string, timeout ...int) ([]byte, string, error) {
	// gen inventory file
	inventoryFile, err := t.genInventoryFile(user, hosts)
	if err != nil {
		logger.Errorf("Failed to genInventoryFile, err=%s", err.Error())
		return nil, "", err
	}
	defer os.Remove(inventoryFile)

	// format ansible cmd
	cmd := fmt.Sprintf("%s ansible all -m copy -u %s -a %s -f %d -i %s -T %d",
		ANSIBLE_SSH_ARGS,
		user,
		libs.EscapeShellArg(fmt.Sprintf("src=%s dest=%s warn=False", src, dest)),
		ANSIBLE_CONCURRENCY,
		inventoryFile,
		ANSIBLE_CONNECTION_TIMEOUT,
	)

	// exec copy file with timeout, default 3600 seconds
	vart := 3600
	if len(timeout) > 0 {
		vart = timeout[0]
	}
	output, err := libs.CmdExecShell(cmd, vart)
	if err != nil {
		logger.Errorf("Failed to execute: %s , err=%s", cmd, err.Error())
		return output, cmd, err
	}

	return output, cmd, nil
}

// generate temporary ansible inventory file
// user: ssh user
// hosts: ["192.168.20.12","192.168.20.11:22","192.168.20.10:444"]
func (t *Client) genInventoryFile(user string, hosts []string) (string, error) {
	if len(hosts) == 0 {
		return "", errors.New("Hosts can't be empty")
	}

	// temporary inventory file path
	tmpRoot := "/tmp/nvwa"
	_, err := libs.CmdExecShellDefault(fmt.Sprintf("mkdir -p %s", tmpRoot))
	if err != nil {
		err = errors.New(err.Error())
		logger.Errorf("Failed to genInventoryFile, err=%s", err.Error())
		return "", err
	}

	// format inventory file contents
	tmpFile := fmt.Sprintf(tmpRoot+"/tmp_ansible_i_%d", time.Now().UnixNano())
	contents := make([]string, 0)
	for _, v := range hosts {
		if v == "" {
			continue
		}

		arr := strings.Split(v, ":")
		ip := arr[0]
		port := "22"
		if len(arr) > 1 {
			port = arr[1]
		}
		contents = append(contents,
			fmt.Sprintf("%s ansible_ssh_user=%s ansible_ssh_port=%s", ip, user, port),
		)
	}

	if len(contents) == 0 {
		return "", errors.New("Hosts can't be empty, failed to format ansible inventory file.")
	}

	// write contents to temporary file
	data := strings.Join(contents, "\n")
	logger.Debugf("[ansible inventory file]")
	logger.Debugf("%s", contents)
	err = libs.FilePutContent(tmpFile, data)
	if err != nil {
		return "", err
	}

	return tmpFile, nil
}
