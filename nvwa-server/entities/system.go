package entities

import (
    "encoding/json"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs"
    "github.com/nvwa-io/nvwa-io/nvwa-server/libs/logger"
)

const (
    // if git url is http(s), need to check which auth type is configured.
    // if git url is SSH, use host's ssh public key and no need to check auth type.
    GIT_CI_AUTH_TYPE_BASIC = 1
    GIT_CI_AUTH_TYPE_TOKEN = 2

    PKG_STORAGE_TYPE_LOCAL  = "local"
    PKG_STORAGE_TYPE_OSS    = "oss"
    PKG_STORAGE_TYPE_COS    = "cos"
    PKG_STORAGE_TYPE_AWS_S3 = "aws-s3"
)

type (
    SystemEntity struct {
        BaseEntity

        Version              string `json:"version"`
        PkgRootPath          string `json:"pkg_root_path"`
        DeployRootPath       string `json:"deploy_root_path"`
        CustomDeployPath     bool   `json:"custom_deploy_path"`
        DeployUser           string `json:"deploy_user"`
        CustomDeployUser     bool   `json:"custom_deploy_user"`
        PkgLimit             int    `json:"pkg_limit"`
        DefaultProjectRoleId int    `json:"default_project_role_id"`
        UseJenkins           bool   `json:"use_jenkins"`
        JenkinsUrl           string `json:"jenkins_url"`
        JenkinsTemplate      string `json:"jenkins_template"`
        JenkinsCredentialId  string `json:"jenkins_credential_id"`
        JenkinsUser          string `json:"jenkins_user"`
        JenkinsPassword      string `json:"jenkins_password"`

        RepoRootPath string `json:"repo_root_path"`

        PkgStorageType   string `json:"pkg_storage_type"`
        PkgStorageConfig string `json:"pkg_storage_config"`

        GitCIAuthType int    `json:"git_ci_auth_type"`
        GitCIUser     string `json:"git_ci_user"`
        GitCIPassword string `json:"git_ci_password"`
        GitCIToken    string `json:"git_ci_token"`

        NotifyEnableTypes string `json:"notify_enable_types"`
        NotifyConfig      string `json:"notify_config"`
    }

    PkgStorageConfig struct {
        // local config
        Local *PkgStorageConfigLocal `json:",omitempty"`

        // oss config
        Oss *PkgStorageConfigOss `json:"oss,omitempty"`
    }

    PkgStorageConfigLocal struct{}

    PkgStorageConfigOss struct {
        Endpoint     string `json:"endpoint"`
        AccessKey    string `json:"access_key"`
        AccessSecret string `json:"access_secret"`
        Bucket       string `json:"bucket"`
    }

    NotifyConfig struct {
        Email struct {
            Host       string `json:"host"`
            Port       string `json:"port"`
            Username   string `json:"username"`
            Password   string `json:"password"`
            Encryption string `json:"encryption"` // tls / ssl
        } `json:"email"`

        // @TODO support wechat notification
        WechatWork struct{}
    }
)

func (t *SystemEntity) FormatJenkinsTemplate(app *AppEntity) (string, error) {
    content, err := libs.ParseTemplate(t.JenkinsTemplate, struct {
        App                 *AppEntity
        JenkinsCredentialId string
    }{
        App:                 app,
        JenkinsCredentialId: t.JenkinsCredentialId,
    })

    if err != nil {
        return "", err
    }

    return content, nil
}

// Decode PkgStorageConfig
func (t *SystemEntity) DecodePkgStorageConfig() (*PkgStorageConfig, error) {
    c := new(PkgStorageConfig)
    err := json.Unmarshal([]byte(t.PkgStorageConfig), c)
    if err != nil {
        logger.Errorf("Failed to Decode package storage config, err=%s", err)
        return nil, err
    }

    return c, nil
}
