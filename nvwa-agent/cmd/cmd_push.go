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

// push sub command
func addPushSubCommand(command *cobra.Command, opt *option.Option) {
	cmd := &cobra.Command{
		Use:   lang.I("agent.cmd.push.use"),
		Short: lang.I("agent.cmd.push.short"),
		Run: func(cmd *cobra.Command, args []string) {
			runSubCommand(cmd, args, opt, func(command *cobra.Command, args []string, opt *option.Option) {
				err := logic.NewPkgLogic().Option(opt).Push()
				if err != nil {
					logger.Errorf("Failed to push app=%s, version package=%s, err=%s", opt.App, opt.Push.VersionPkg, err.Error())
					os.Exit(1)
				}

				logger.Infof("Push app=%s, version package=%s succeeded ... ok", opt.App, opt.Push.VersionPkg)
			})
		},
	}

	cmd.Flags().Int64Var(&opt.Push.BuildId, "build-id", 0, lang.I("agent.cmd.push.arg.build-id"))
	cmd.Flags().StringVar(&opt.Push.VersionPkg, "version-pkg", "", lang.I("agent.cmd.push.arg.version-pkg"))
	command.AddCommand(cmd)
}
