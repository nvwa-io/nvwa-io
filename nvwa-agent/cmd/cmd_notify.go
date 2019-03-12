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

// notify sub command
func addNotifySubCommand(command *cobra.Command, opt *option.Option) {
	cmd := &cobra.Command{
		Use:   lang.I("agent.cmd.notify.use"),
		Short: lang.I("agent.cmd.notify.short"),
		Run: func(cmd *cobra.Command, args []string) {
			runSubCommand(cmd, args, opt, func(command *cobra.Command, args []string, opt *option.Option) {
				err := logic.NewApiLogic().Option(opt).Notify()
				if err != nil {
					logger.Errorf("Failed to notify, err=%s", err.Error())
					os.Exit(1)
				}

				logger.Infof("Message notification committed, message=%s, users=%v", opt.Notify.Message, opt.Notify.NotifyUser)
				os.Exit(0)
			})
		},
	}

	cmd.Flags().StringVar(&opt.Notify.Message, "message", "", lang.I("agent.cmd.notify.arg.message"))
	cmd.Flags().StringSliceVar(&opt.Notify.NotifyUser, "users", nil, lang.I("agent.cmd.notify.arg.users"))
	command.AddCommand(cmd)
}
