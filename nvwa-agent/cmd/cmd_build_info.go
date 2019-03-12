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

// build sub command
func addBuildInfoSubCommand(command *cobra.Command, opt *option.Option) {
	cmd := &cobra.Command{
		Use:   lang.I("agent.cmd.build-info.use"),
		Short: lang.I("agent.cmd.build-info.short"),
		Run: func(cmd *cobra.Command, args []string) {
			runSubCommand(cmd, args, opt, func(command *cobra.Command, args []string, opt *option.Option) {
				status := option.BuildStatus2Int(opt.BuildInfo.Status)
				err := logic.NewApiLogic().Option(opt).UpdateBuildInfo(opt.BuildInfo.BuildId, map[string]interface{}{
					"status":            status,
					"branch":            opt.JenkinsEnv.GitBranch,
					"commit_id":         opt.JenkinsEnv.GitCommit,
					"jenkins_build_num": opt.JenkinsEnv.BuildNumber,
				})

				if err != nil {
					logger.Errorf("Failed to Update build info, err=%s", err.Error())
					os.Exit(1)
				}

				os.Exit(0)
			})
		},
	}

	cmd.Flags().Int64Var(&opt.BuildInfo.BuildId, "build-id", 0, lang.I("agent.cmd.build-info.arg.build-id"))
	cmd.Flags().StringVar(&opt.BuildInfo.Status, "status", "", lang.I("agent.cmd.build-info.arg.status"))
	command.AddCommand(cmd)
}
