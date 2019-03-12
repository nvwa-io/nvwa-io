package vo

import (
    "github.com/astaxie/beego/validation"
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqApp struct {
    App entities.AppEntity `json:"app"`
}

// validate request app params
func (t *ReqApp) Valid() error {
    valid := validation.Validation{}
    valid.Min(t.App.ProjectId, 1, "project_id").Message(lang.I("project.id.invalid"))
    valid.Required(t.App.Name, "name").Message(lang.I("app.name.not.empty"))
    valid.Required(t.App.RepoUrl, "repo_url").Message(lang.I("app.repo_url.not.empty"))
    valid.Required(t.App.DeployUser, "deploy_user").Message(lang.I("app.deploy_user.not.empty"))
    //valid.Min(t.App.DeployType, 1, "deploy_type").Message(lang.I("app.deploy_type.invalid"))
    //valid.Min(t.App.AppType, 1, "app_type").Message(lang.I("app.app_type.invalid"))
    //valid.Required(t.App.RepoType, "repo_type").Message(lang.I("app.repo_type.not.empty"))
    //valid.Required(t.App.CmdTimeout, "cmd_timeout").Message(lang.I("app.cmd_timeout.not.empty"))
    if valid.HasErrors() {
        return valid.Errors[0]
    }

    return nil
}
