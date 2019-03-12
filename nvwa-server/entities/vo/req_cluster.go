package vo

import (
    "github.com/astaxie/beego/validation"
    "github.com/nvwa-io/nvwa-io/nvwa-server/entities"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
)

type ReqCluster struct {
    Cluster entities.ClusterEntity `json:"cluster"`
}

func (t *ReqCluster) Valid() error {
    valid := validation.Validation{}
    valid.Min(t.Cluster.AppId, 1, "app_id").Message(lang.I("app.id.invalid"))
    valid.Required(t.Cluster.Name, "name").Message(lang.I("cluster.name.not.empty"))
    valid.Required(t.Cluster.Hosts, "hosts").Message(lang.I("cluster.hosts.not.empty"))
    if valid.HasErrors() {
        return valid.Errors[0]
    }

    return nil
}
