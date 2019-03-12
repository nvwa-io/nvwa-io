package errs

import (
    "fmt"
    "github.com/nvwa-io/nvwa-io/nvwa-server/lang"
    "strconv"
    "strings"
)

type ErrStr string

func e(code int, msg string) ErrStr {
    return ErrStr(fmt.Sprintf("%d%s%s", code, ERR_SEP, msg))
}

var (
    ERR_NONE          = 200
    ERR_SEP           = "=>"
    ERR_PARAM         = e(400, lang.I("errs.param"))
    ERR_INVALID_TOKEN = e(401, lang.I("errs.invalid.token"))
    ERR_LOGIN         = e(402, lang.I("errs.login"))
    ERR_SIGN          = e(403, lang.I("errs.sign"))
    ERR_NO_RECORD     = e(404, lang.I("errs.no.record"))

    ERR_OPERATE = e(500, lang.I("errs.operate"))
    ERR_UNKNOWN = e(501, lang.I("errs.unknown"))
    ERR_SYSTEM  = e(503, lang.I("errs.system"))

    // users
    ERR_USER_EXIST = e(1001, lang.I("user.exist"))

    // projects
    ERR_PROJECT_EXIST = e(2001, lang.I("project.exist"))

    // project roles
    ERR_PROJECT_ROLE_EXIST = e(2101, lang.I("project_role.exist"))

    // member
    ERR_MEMBER_EXIST = e(2201, lang.I("member.exist"))

    // app
    ERR_APP_EXIST = e(2301, lang.I("app.exist"))

    // env
    ERR_ENV_EXIST = e(2401, lang.I("env.exist"))

    // cluster
    ERR_CLUSTER_EXIST = e(2501, lang.I("cluster.exist"))
)

func NoFound(err error, info ...string) ErrStr {
    msg := ""
    if len(info) > 0 {
        msg = fmt.Sprintf("%s:%s:%s", lang.I("errs.no.record"), info[0], err.Error())
    } else {
        msg = fmt.Sprintf("%s:%s", lang.I("errs.no.record"), err.Error())
    }
    return e(404, msg)
}

func (ErrStr ErrStr) GetErrno() int {
    s := strings.Split(string(ErrStr), ERR_SEP)
    errno, _ := strconv.Atoi(s[0])
    return errno
}

func (ErrStr ErrStr) GetErrmsg() string {
    return strings.Split(string(ErrStr), ERR_SEP)[1]
}

func (ErrStr ErrStr) Error() string {
    return string(ErrStr)
}

func IsErrStr(err error) bool {
    switch err.(type) {
    case *ErrStr:
        return true
    default:
        return false
    }
}
