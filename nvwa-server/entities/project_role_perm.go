package entities

import "github.com/nvwa-io/nvwa-io/nvwa-server/lang"

const (
    // project permission
    PERM_PROJECT_CREATE = "project.create"
    PERM_PROJECT_UPDATE = "project.update"

    // project member permissions
    PERM_MEMBER_ADD    = "member.add"
    PERM_MEMBER_REMOVE = "member.remove"
    // tips: change other member's role
    PERM_MEMBER_CHANGE_ROLE = "member.change.role"

    // app permissions
    PERM_APP_CREATE = "app.create"
    PERM_APP_UPDATE = "app.update"
    PERM_APP_DELETE = "app.delete"

    // app's env permission
    PERM_ENV_CREATE = "env.create"
    PERM_ENV_UPDATE = "env.update"
    PERM_ENV_DELETE = "env.delete"
    // permission to configure environment's audit
    // and permission pass or reject the environment's deployments
    PERM_ENV_AUDIT = "env.audit"

    // app's cluster permissions
    PERM_CLUSTER_CREATE = "cluster.create"
    PERM_CLUSTER_UPDATE = "cluster.update"
    PERM_CLUSTER_DELETE = "cluster.delete"

    // deployment's permissions
    PERM_DEPLOYMENT_CREAET = "deployment.create"
)

// permission labels
var PERM_LABELS = map[string]string{
    // project permission
    PERM_PROJECT_CREATE: lang.I("project.create"),
    PERM_PROJECT_UPDATE: lang.I("project.update"),

    // project member permissions
    PERM_MEMBER_ADD:    lang.I("member.add"),
    PERM_MEMBER_REMOVE: lang.I("member.remove"),
    // tips: change other member's role
    PERM_MEMBER_CHANGE_ROLE: lang.I("member.change.role"),

    // app permissions
    PERM_APP_CREATE: lang.I("app.create"),
    PERM_APP_UPDATE: lang.I("app.update"),
    PERM_APP_DELETE: lang.I("app.delete"),

    // app's env permission
    PERM_ENV_CREATE: lang.I("env.create"),
    PERM_ENV_UPDATE: lang.I("env.update"),
    PERM_ENV_DELETE: lang.I("env.delete"),
    // permission to configure environment's audit
    // and permission pass or reject the environment's deployments
    PERM_ENV_AUDIT: lang.I("env.audit"),

    // app's cluster permissions
    PERM_CLUSTER_CREATE: lang.I("cluster.create"),
    PERM_CLUSTER_UPDATE: lang.I("cluster.update"),
    PERM_CLUSTER_DELETE: lang.I("cluster.delete"),

    // deployment's permissions
    PERM_DEPLOYMENT_CREAET: lang.I("deployment.create"),
}

type ProjectRolePermEntity struct {
    BaseEntity

    ProjectRoleId int64  `json:"project_role_id"`
    Perm          string `json:"perm"`
}
