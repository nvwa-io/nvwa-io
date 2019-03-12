const getters = {
  sidebar: state => state.app.sidebar,
  language: state => state.app.language,
  size: state => state.app.size,
  device: state => state.app.device,
  visitedViews: state => state.tagsView.visitedViews,
  cachedViews: state => state.tagsView.cachedViews,

  token: state => state.user.token,
  avatar: state => state.user.curUser.avatar,
  // name: state => state.user.name,
  // introduction: state => state.user.introduction,
  // status: state => state.user.status,
  auditNum: state => state.user.auditNum,
  roles: state => state.user.roles,
  // setting: state => state.user.setting,
  // user end routers
  permission_routers: state => state.permission.routers,
  addRouters: state => state.permission.addRouters,
  // admin end routers
  adminRouters: state => state.permission.adminRouters,
  errorLogs: state => state.errorLog.logs
}
export default getters
