package vo

import (
    "github.com/astaxie/beego/validation"
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqDeployment struct {
    Deployment entities.DeploymentEntity `json:"deployment"`
}

func (t *ReqDeployment) Valid() error {
    valid := validation.Validation{}
    valid.Min(t.Deployment.ProjectId, 1, "project_id").Message(lang.I("project.id.invalid"))
    valid.Min(t.Deployment.AppId, 1, "app_id").Message(lang.I("app.id.invalid"))
    valid.Min(t.Deployment.EnvId, 1, "env_id").Message(lang.I("env.id.invalid"))
    valid.Required(t.Deployment.ClusterIds, "cluster_ids").Message(lang.I("deployment.cluster_ids.not.empty"))
    valid.Min(t.Deployment.PkgId, 1, "pkg_id").Message(lang.I("pkg.id.invalid"))
    if valid.HasErrors() {
        return valid.Errors[0]
    }

    return nil
}
