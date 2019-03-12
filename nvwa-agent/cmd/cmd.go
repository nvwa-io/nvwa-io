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

package cmd

import (
	"fmt"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/option"
	"github.com/spf13/cobra"
)

const (
	DEFAULT_TIMEOUT = 10 * 60
)

func NewAgentCmd() *cobra.Command {
	opt := option.NewOption()
	cmd := &cobra.Command{
		Use:  lang.I("agent.name"),
		Long: lang.I("agent.long"),
		//Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(lang.I("agent.no.operation"))
			cmd.Usage()
		},
	}

	// global args
	cmd.PersistentFlags().StringVarP(&opt.ConfFile, "config", "c", "/data/nvwa/nvwa-agent/etc/config.toml", lang.I("agent.cmd.config"))
	cmd.PersistentFlags().StringVar(&opt.App, "app", "", lang.I("agent.cmd.app"))

	// add sub commands
	addNotifySubCommand(cmd, opt)
	addPushSubCommand(cmd, opt)
	addPullSubCommand(cmd, opt)
	addBuildInfoSubCommand(cmd, opt)
	addTarSubCommand(cmd, opt)
	return cmd
}

// run sub command
func runSubCommand(command *cobra.Command, args []string, opt *option.Option,
	run func(command *cobra.Command, args []string, opt *option.Option)) {
	opt.Config()
	run(command, args, opt)
}
