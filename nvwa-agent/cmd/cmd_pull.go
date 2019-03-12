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
	"github.com/nvwa-io/nvwa-io/nvwa-agent/lang"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/libs/logger"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/logic"
	"github.com/nvwa-io/nvwa-io/nvwa-agent/option"
	"github.com/spf13/cobra"
	"os"
)

// pull sub command
func addPullSubCommand(command *cobra.Command, opt *option.Option) {
	cmd := &cobra.Command{
		Use:   lang.I("agent.cmd.pull.use"),
		Short: lang.I("agent.cmd.pull.short"),
		Run: func(cmd *cobra.Command, args []string) {
			runSubCommand(cmd, args, opt, func(command *cobra.Command, args []string, opt *option.Option) {
				err := logic.NewPkgLogic().Option(opt).Pull()
				if err != nil {
					logger.Errorf("Failed to pull app=%s, version package=%s, err=%s", opt.App, opt.Pull.VersionPkg, err.Error())
					os.Exit(1)
				}

				logger.Infof("Pull app=%s, version package=%s succeeded ... ok", opt.App, opt.Pull.VersionPkg)
			})
		},
	}

	cmd.Flags().StringVar(&opt.Pull.VersionPkg, "version-pkg", "", lang.I("agent.cmd.pull.arg.version-pkg"))
	command.AddCommand(cmd)
}
