package lang

var zhCN = map[string]string{
    // !-- errors -- begin
    "success":            "成功",
    "errs.param":         "参数错误",
    "errs.unknown":       "未知错误",
    "errs.sign":          "签名认证失败",
    "errs.system":        "系统异常",
    "errs.operate":       "操作失败",
    "errs.invalid.token": "未登录或 Token 无效",
    "errs.login":         "登录失败，密码错误或用户不存在",
    "errs.no.record":     "记录不存在",
    // !-- errors -- end

    // project info
    "project.name.not.empty": "项目名称不能为空",
    "project.id.invalid":     "项目 ID 无效",
    "project.exist":          "项目已存在",

    // project role
    "project_role.exist":               "项目角色已存在",
    "project_role.not.exist":           "项目角色不存在",
    "project_role.name.not.empty":      "项目角色名称不能为空",
    "project_role.id.invalid":          "项目角色 ID 无效",
    "project_role_perm.perm.not.empty": "角色权限不能为空",
    "project_role_perm.perm.invalid":   "无效权限字段",

    // user info
    "user.exist":                 "用户已存在",
    "user.username.invalid":      "用户名格式错误: [0-9a-zA-Z_-]",
    "user.password.not.empty":    "密码不能为空",
    "user.password.length.error": "密码长度至少 6 位",
    "user.email.format.error":    "邮箱格式错误",
    "user.id.invalid":            "用户 ID 无效",

    // member
    "member.exist":      "项目成员已存在",
    "member.id.invalid": "项目成员 ID 无效",

    // app
    "app.name.not.empty":        "应用名不能为空: [0-9a-zA-Z_-]",
    "app.name.invalid":          "应用名格式错误: [0-9a-zA-Z_-]",
    "app.exist":                 "应用名已被占用",
    "app.repo_url.not.empty":    "仓库地址不能为空",
    "app.deploy_user.not.empty": "部署用户不能为空",
    "app.repo_type.not.empty":   "仓库类型不能为空",
    "app.deploy_type.invalid":   "应用部署类型无效",
    "app.app_type.invalid":      "应用类型无效",
    "app.cmd_timeout.not.empty": "命令超时时间不能为空",
    "app.id.invalid":            "应用 ID 无效",

    // env
    "env.exist":                     "环境已存在",
    "env.name.not.empty":            "环境名称不能为空",
    "env.id.invalid":                "环境 ID 无效",
    "env.permit_branches.not.empty": "允许分支不能为空",
    "env.is_auto_deploy.not.empty":  "参数'是否自动部署'不能为空",
    "env.is_need_audit.not.empty":   "参数'是否需要审核'不能为空",
    "env.not.found":                 "环境不存在 ",

    // init 4 environments while create app
    "env.dev":     "开发环境",
    "env.test":    "测试环境",
    "env.staging": "预发布环境",
    "env.prod":    "生成环境",

    // cluster
    "cluster.exist":           "分组已存在",
    "cluster.name.not.empty":  "分组名称不能为空",
    "cluster.hosts.not.empty": "分组主机不能为空",
    "cluster.default":         "默认分组",

    // build
    "build.branch.not.empty": "分支不能为空",

    // deployment
    "deployment.cluster_ids.not.empty": "部署分组不能为空",

    // job
    "job.id.invalid":                "任务 ID 无效",

    // pkg
    "pkg.id.invalid": "版本包 ID 无效",
    "pkg.not.found":  "版本包不存在",

    // !-- permissions -- begin
    // project permission
    "project.create": "新建项目",
    "project.update": "修改项目",

    // project member permissions
    "member.add":    "添加成员",
    "member.remove": "移除成员",
    // tips: change other member's role
    "member.change.role": "修改角色",

    // app permissions
    "app.create": "新建应用",
    "app.update": "修改应用",
    "app.delete": "删除应用",

    // app's env permission
    "env.create": "新建环境",
    "env.update": "修改环境",
    "env.delete": "删除环境",
    // permission to configure environment's audit
    // and permission pass or reject the environment's deployments
    "env.audit": "审核权限",

    // app's cluster permissions
    "cluster.create": "新建分组",
    "cluster.update": "修改分组",
    "cluster.delete": "删除分组",

    // deployment's permissions
    "deployment.create": "新建部署单",
    // !-- permissions -- end
}
