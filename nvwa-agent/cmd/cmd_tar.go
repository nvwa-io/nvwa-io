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

// tar sub command
func addTarSubCommand(command *cobra.Command, opt *option.Option) {
	cmd := &cobra.Command{
		Use:   lang.I("agent.cmd.tar.use"),
		Short: lang.I("agent.cmd.tar.short"),
		Run: func(cmd *cobra.Command, args []string) {
			runSubCommand(cmd, args, opt, func(command *cobra.Command, args []string, opt *option.Option) {
				pkgName, err := logic.NewPkgLogic().Option(opt).Tar(opt.Tar.BuildId, opt.Tar.Files, opt.Tar.Excludes)
				if err != nil {
					logger.Errorf(err.Error())
					os.Exit(1)
				}

				// set pkgName to env, so other process can get packed version package name from env in the same shell
				os.Setenv("PACKED_VERSION_PACKAGE", pkgName)
				logger.Infof("Tar %s succeeded ...ok", pkgName)
				os.Exit(0)
			})
		},
	}

	cmd.Flags().Int64Var(&opt.Tar.BuildId, "build-id", 0, lang.I("agent.cmd.tar.arg.build-id"))
	cmd.Flags().StringSliceVar(&opt.Tar.Files, "files", nil, lang.I("agent.cmd.tar.arg.files"))
	cmd.Flags().StringSliceVar(&opt.Tar.Excludes, "excludes", nil, lang.I("agent.cmd.tar.arg.excludes"))
	command.AddCommand(cmd)
}
