const vars = {
  deploy: {
    home: '首页',
    id: 'ID',
    addProject: '新建项目',
    editProject: '修改项目',
    myProject: '我的项目',
    projectName: '项目名称',
    deleteConfirm: ' 确定删除',
    tipDeleteProject: '确定删除该项目？',

    role: '角色',
    memberManagement: '成员管理',
    nvwaArch: '女娲架构原理',
    username: '用户名',
    remove: '移除',
    editRole: '修改角色',
    selectRole: '选择角色 ',
    loadingRole: '加载角色中...',
    loadingUsers: '加载用户中...',
    selectUser: '选择用户',
    titleRemoveMember: '移除确认',
    tipRemoveMember: '确定移除成员？',
    rolePerm: '角色权限',
    permScope: '权限域',
    permItem: '权限项',
    addMember: '添加成员',
    projectManagement: '项目管理',
    auditUser: '审核人',
    titlePassAudit: '通过确认',
    tipPassAudit: '确认通过该部署单？',
    titleRejectAudit: '驳回确认',
    tipRejectAudit: '确认驳回该部署单？',
    titleCancelAudit: '取消确认',
    tipCancelAudit: '确认取消该部署单？',
    auditStatusLabels: {
      // AUDIT_STATUS_WAITING = 10
      // AUDIT_STATUS_PASS    = 40
      // AUDIT_STATUS_REJECT  = 50
      // AUDIT_STATUS_CANCELED  = 60
      10: '待审核',
      40: '已通过',
      50: '驳回',
      60: '已取消'
    },

    // deployment audit
    order: '工单',
    waitAudit: '待审核',
    myDeploymentOrder: '我的工单',
    audited: '已审核',
    pass: '通过',
    reject: '驳回',
    noAuditPerm: '无审核权限',

    project: '项目',
    member: '成员',
    app: '应用',
    env: '部署环境',
    cluster: '服务器分组',
    deployment: '部署',

    // project role perm is : xxx.xxx
    // for convenience, we replace . to ___ to match the field
    project___update: '修改',
    project___create: '新建',
    member___add: '添加',
    member___remove: '移除',
    member___change___role: '修改角色',
    app___create: '新建',
    app___update: '修改',
    app___delete: '删除',
    env___create: '新建',
    env___update: '修改',
    env___delete: '删除',
    env___audit: '审核',
    cluster___create: '新建',
    cluster___update: '修改',
    cluster___delete: '删除',
    deployment___create: '新建',

    index: '概览',
    appList: '应用列表',
    buildList: '构建打包',
    envDeploy: '环境部署',
    clusterDeploy: '分组部署',
    deploymentList: '部署单',
    buildCmd: '构建命令',
    addApp: '新建应用',
    description: '简介',
    appType: '应用类型',
    deployCmd: '部署命令',
    groupHost: '主机分组',
    appStatus: '应用状态',
    applyPermission: '申请权限',
    buildDeploy: '构建/部署',
    creator: '创建人',

    appName: '应用名称',
    appDescription: '应用简介',
    gitRepo: 'Git 仓库',
    gitUsername: '用户名',
    gitPassword: '密码',
    deployUser: '部署用户',
    deployPath: '部署路径',
    gitSshTip: '需要女娲所在主机能 SSH 免密拉去/更新代码。',
    gitHttpTip: '需要配置仓库账号密码或使用系统默认账号。',

    addEnv: '新建环境',
    clusterName: '分组名称',
    viewHostGroup: '查看分组',
    clusterHostNum: '服务器数量',
    audit: '审核',
    detection: '检测',
    detectionCond: '条件检测',
    // codeBranch: '代码分支',
    addCluster: '添加分组',
    configCluster: '配置分组',
    enableAudit: '开启审核',
    envName: '环境名称',
    clusterHosts: '服务器列表',
    operationSuccess: '操作成功',

    selectBranches: '选择分支',
    selectApp: '选择应用',
    loadingApps: '加载应用中...',
    loadingBranches: '加载分支中...',
    pleaseSelectApp: '请选择应用',
    pleaseSelectBranch: '请选择分支',
    branch: '分支',
    commit: 'commit',
    ctime: '创建时间',
    utime: '更新时间',
    status: '状态',
    statusLabels: {
      // BUILD_STATUS_CREATED         = 10
      // BUILD_STATUS_BUIDING         = 20
      // BUILD_STATUS_BUILD_SUCC      = 30
      // BUILD_STATUS_BUILD_FAILED    = 40
      // BUILD_STATUS_PACK_SUCC       = 50
      // BUILD_STATUS_PACK_FAILED     = 60
      // BUILD_STATUS_PKG_PUSH_SUCC   = 70
      // BUILD_STATUS_PKG_PUSH_FAILED = 80
      10: '就绪',
      20: '构建中',
      30: '构建通过',
      40: '构建失败',
      50: '已打包',
      60: '打包失败',
      70: '完成',
      80: '失败'
    },
    process: '进度',
    log: '日志',
    packageName: '版本包',
    launchBuild: '发起构建',
    launchDeploy: '发起部署',
    appTypeLabels: {
      1: '自定义',
      2: 'Spring Boot',
      3: 'NodeJs',
      4: 'PM2 部署'
    },
    permitBranches: '可选分支',
    deployClusters: '分批部署',
    deploymentStatusLabels: {
      // DEPLOYMENT_STATUS_CREATED      = 10
      // DEPLOYMENT_STATUS_NO_AUDIT     = 20
      // DEPLOYMENT_STATUS_WAIT_AUDIT   = 30
      // DEPLOYMENT_STATUS_AUDIT_PASS   = 40
      // DEPLOYMENT_STATUS_AUDIT_REJECT = 50
      // DEPLOYMENT_STATUS_CANCELED     = 60
      // DEPLOYMENT_STATUS_DEALING      = 70
      // DEPLOYMENT_STATUS_SUCC         = 80
      // DEPLOYMENT_STATUS_FAILED       = 90
      10: '待部署',
      20: '审核通过',
      30: '待审核',
      40: '审核通过',
      50: '审核拒绝',
      60: '已取消',
      70: '部署中',
      80: '成功',
      90: '失败'
    },
    jobStatusLabels: {
      // JOB_STATUS_CREATED = 10
      // JOB_STATUS_READY   = 20
      // JOB_STATUS_DEALING = 30
      // JOB_STATUS_SUCC    = 40
      // JOB_STATUS_FAILED  = 50
      10: '待部署',
      20: '就绪',
      30: '部署中',
      40: '成功',
      50: '失败'
    },
    jobStepLabels: {
      // JOB_STEP_INIT_WORKSPACE         = 10
      // JOB_STEP_SYNC_VERISON_PACKAGE   = 20
      // JOB_STEP_UNPACK_VERISON_PACKAGE = 30
      // JOB_STEP_CMD_BEFORE_DEPLOY      = 40
      // JOB_STEP_DO_DEPLOY              = 50
      // JOB_STEP_CMD_AFTER_DEPLOY       = 60
      // JOB_STEP_CMD_HEALTH_CHECK       = 70
      // JOB_STEP_CMD_ONLINE             = 80
      // JOB_STEP_END_CLEAN              = 90
      10: '初始化',
      20: '同步版本包',
      30: '解压版本包',
      40: '部署前命令',
      50: '执行部署',
      60: '部署后命令',
      70: '应用检测',
      80: '上线命令',
      90: '清理工作'
    },
    startDeploy: '开始部署',
    tipStartDeploy: '确定开始部署？',
    confirmDeploy: '确认部署',
    confirm: '确认',
    cancel: '取消'
  },
  dycompute: {
    applyServer: '申请服务器',
    hostname: '主机名',
    innerIp: '内网IP',
    outerIp: '外网IP',
    diskType: '磁盘类型',
    status: '状态',
    monitor: '监控',
    applyAccount: '申请账号',
    installSoftware: '安装软件',
    tagManagement: '标签管理',
    applyDomain: '申请域名',
    domain: '域名',
    parseInfo: '解析信息',
    editParse: '修改解析',
    applyLvs: '申请 LVS',
    vip: 'VIP',
    port: '端口',
    zone: '分区',
    rsWithStatus: 'RS(正常/异常)',
    healthCheckType: '健康检查方式',
    note: '备注',
    viewBandwidth: '查看带宽',
    manageRs: '管理 RS',
    addPort: '增加端口',
    tempOffline: '临时下线',
    addRs: '新增 RS',
    batchDeleteRs: '批量删除 RS',
    ip: 'IP'
  },
  db: {
    label: '数据库',
    mysql: 'MySQL',
    redis: 'Redis',
    mongodb: 'MongoDB',
    addIns: '新建实例',
    port: '端口',
    zone: '分区',
    insName: '实例名称',
    usedCapacity: '已用容量',
    optimizeAdvice: '优化建议',
    linkInfo: '连接信息',
    basicInfo: '基本信息',
    manageAccount: '账号管理',
    manageDb: '数据库管理',
    editConfig: '参数设置',
    editPackage: '变更配置',
    addTable: '建表',
    editTable: '改表',
    deleteIns: '删除实例',
    version: '版本',
    monitor: '监控',
    address: '地址',
    link: '连接',
    apply: '申请',
    innerOut: '内网/外网'
  },
  admin: {
    deploy: {
      dashboard: '概览',
      manageUser: '用户管理',
      manageProject: '项目管理',
      projectList: '项目列表',
      manageProjectRole: '项目角色',
      manageApp: '应用管理',
      manageOrder: '工单管理',
      systemConfig: '系统设置',
      addProjectRole: '添加项目角色',
      editProjectRole: '修改项目角色',
      projectRoleName: '角色名称',

      account: '账号',
      displayName: '姓名',
      email: '邮箱',
      systemRole: '系统角色'
    }
  }
}

export default {
  route: {
    dashboard: '总览',
    introduction: '简述',
    documentation: '文档',
    guide: '引导页',
    permission: '权限测试页',
    pagePermission: '页面权限',
    directivePermission: '指令权限',
    icons: '图标',
    deploySystem: '应用部署',

    deploy: vars.deploy,
    admin: vars.admin,

    codeManagement: {
      label: '代码仓库'
    },
    dynamicCompute: '弹性计算',
    compute: {
      ecs: '云服务器',
      domain: '域名解析',
      lvs: '负载均衡'
    },
    db: vars.db,
    monitor: {
      label: '服务监控',
      index: '概览',
      app: '应用监控',
      ecs: '服务器监控',
      lvs: '负载均衡监控',
      mysql: 'MySQL 监控',
      redis: 'Redis 监控',
      mongodb: 'MongoDB 监控'
    },
    communication: {
      label: '云通信',
      index: '概览',
      sms: '短信服务',
      push: '移动推送',
      email: '邮件推送'
    },
    secret: {
      label: '密钥管理',
      index: '密钥列表'
    },
    components: '组件',
    componentIndex: '介绍',
    tinymce: '富文本编辑器',
    markdown: 'Markdown',
    jsonEditor: 'JSON编辑器',
    dndList: '列表拖拽',
    splitPane: 'Splitpane',
    avatarUpload: '头像上传',
    dropzone: 'Dropzone',
    sticky: 'Sticky',
    countTo: 'CountTo',
    componentMixin: '小组件',
    backToTop: '返回顶部',
    dragDialog: '拖拽 Dialog',
    dragSelect: '拖拽 Select',
    dragKanban: '可拖拽看板',
    charts: '图表',
    keyboardChart: '键盘图表',
    lineChart: '折线图',
    mixChart: '混合图表',
    example: '综合实例',
    nested: '路由嵌套',
    menu1: '菜单1',
    'menu1-1': '菜单1-1',
    'menu1-2': '菜单1-2',
    'menu1-2-1': '菜单1-2-1',
    'menu1-2-2': '菜单1-2-2',
    'menu1-3': '菜单1-3',
    menu2: '菜单2',
    Table: 'Table',
    dynamicTable: '动态Table',
    dragTable: '拖拽Table',
    inlineEditTable: 'Table内编辑',
    complexTable: '综合Table',
    treeTable: '树形表格',
    customTreeTable: '自定义树表',
    tab: 'Tab',
    form: '表单',
    createArticle: '创建文章',
    editArticle: '编辑文章',
    articleList: '文章列表',
    errorPages: '错误页面',
    page401: '401',
    page404: '404',
    errorLog: '错误日志',
    excel: 'Excel',
    exportExcel: 'Export Excel',
    selectExcel: 'Export Selected',
    uploadExcel: 'Upload Excel',
    zip: 'Zip',
    exportZip: 'Export Zip',
    theme: '换肤',
    clipboardDemo: 'Clipboard',
    i18n: '国际化',
    externalLink: '外链'
  },
  // 页面变量
  page: {
    viewHelp: '查看帮助',
    save: ' 保存',
    delete: '删除',
    ctime: '创建时间',
    utime: '更新时间',
    batchOp: '批量操作',
    action: '操作',
    more: '更多',
    status: '状态',

    opTable: {
      utime: '操作时间',
      edit: '修改',
      delete: '删除'
    },
    codeManagement: {
      add: '新建仓库',
      id: 'ID',
      repo: '仓库',
      repoUrl: '仓库地址',
      permission: '权限',
      action: '操作',
      utime: '操作时间',
      memberManage: '成员管理',
      applyPermission: '申请权限',
      delete: '删除',
      sshKeyManagement: 'SSH key管理'
    },
    deploy: vars.deploy,
    admin: vars.admin
  },
  navbar: {
    logOut: '退出登录',
    dashboard: '首页',
    github: '项目地址',
    screenfull: '全屏',
    theme: '换肤',
    size: '布局大小'
  },
  login: {
    title: '系统登录',
    logIn: '登录',
    username: '账号',
    password: '密码',
    any: '随便填',
    thirdparty: '第三方登录',
    thirdpartyTips: '本地不能模拟，请结合自己业务进行模拟！！！'
  },
  documentation: {
    documentation: '文档',
    github: 'Github 地址'
  },
  permission: {
    roles: '你的权限',
    switchRoles: '切换权限'
  },
  guide: {
    description: '引导页对于一些第一次进入项目的人很有用，你可以简单介绍下项目的功能。本 Demo 是基于',
    button: '打开引导'
  },
  components: {
    documentation: '文档',
    tinymceTips: '富文本是管理后台一个核心的功能，但同时又是一个有很多坑的地方。在选择富文本的过程中我也走了不少的弯路，市面上常见的富文本都基本用过了，最终权衡了一下选择了Tinymce。更详细的富文本比较和介绍见',
    dropzoneTips: '由于我司业务有特殊需求，而且要传七牛 所以没用第三方，选择了自己封装。代码非常的简单，具体代码你可以在这里看到 @/components/Dropzone',
    stickyTips: '当页面滚动到预设的位置会吸附在顶部',
    backToTopTips1: '页面滚动到指定位置会在右下角出现返回顶部按钮',
    backToTopTips2: '可自定义按钮的样式、show/hide、出现的高度、返回的位置 如需文字提示，可在外部使用Element的el-tooltip元素',
    imageUploadTips: '由于我在使用时它只有vue@1版本，而且和mockjs不兼容，所以自己改造了一下，如果大家要使用的话，优先还是使用官方版本。'
  },
  table: {
    dynamicTips1: '固定表头, 按照表头顺序排序',
    dynamicTips2: '不固定表头, 按照点击顺序排序',
    dragTips1: '默认顺序',
    dragTips2: '拖拽后顺序',
    title: '标题',
    importance: '重要性',
    type: '类型',
    remark: '点评',
    search: '搜索',
    add: '添加',
    export: '导出',
    reviewer: '审核人',
    id: '序号',
    date: '时间',
    author: '作者',
    readings: '阅读数',
    status: '状态',
    actions: '操作',
    edit: '编辑',
    publish: '发布',
    draft: '草稿',
    delete: '删除',
    cancel: '取 消',
    confirm: '确 定'
  },
  errorLog: {
    tips: '请点击右上角bug小图标',
    description: '现在的管理后台基本都是spa的形式了，它增强了用户体验，但同时也会增加页面出问题的可能性，可能一个小小的疏忽就导致整个页面的死锁。好在 Vue 官网提供了一个方法来捕获处理异常，你可以在其中进行错误处理或者异常上报。',
    documentation: '文档介绍'
  },
  excel: {
    export: '导出',
    selectedExport: '导出已选择项',
    placeholder: '请输入文件名(默认excel-list)'
  },
  zip: {
    export: '导出',
    placeholder: '请输入文件名(默认file)'
  },
  theme: {
    change: '换肤',
    documentation: '换肤文档',
    tips: 'Tips: 它区别于 navbar 上的 theme-pick, 是两种不同的换肤方法，各自有不同的应用场景，具体请参考文档。'
  },
  tagsView: {
    refresh: '刷新',
    close: '关闭',
    closeOthers: '关闭其它',
    closeAll: '关闭所有'
  }
}
