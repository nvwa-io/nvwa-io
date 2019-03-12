import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css'// progress bar style
import { getToken } from '@/utils/auth' // getToken from cookie

NProgress.configure({ showSpinner: false })// NProgress Configuration

// permission judge function
function hasPermission(roles, permissionRoles) {
  if (roles.indexOf('admin') >= 0) return true // admin permission passed directly
  if (!permissionRoles) return true
  return roles.some(role => permissionRoles.indexOf(role) >= 0)
}

const whiteList = ['/home/index', '/login', '/auth-redirect']// no redirect whitelist

router.beforeEach((to, from, next) => {
  NProgress.start() // start progress bar
  if (getToken()) { // determine if there has token
    /* has token*/
    if (to.path === '/login') {
      next({ path: '/' })
      // if current page is dashboard will not trigger	afterEach hook, so manually handle it
      NProgress.done()
    } else {
      // check whether there is user info
      if (store.getters.roles.length === 0) {
        store.dispatch('GetUserInfo').then(res => {
          // note: roles must be a array! such as: ['admin']
          // generate routers according to user's role
          const roles = [res.data.user.role]
          store.dispatch('GenerateRoutes', { roles }).then(() => {
            // @TODO 判断角色，添加管理员路由
            console.log(store.getters.adminRouters)
            router.addRoutes(store.getters.adminRouters)

            // add route to router table dynamically
            router.addRoutes(store.getters.addRouters)

            // hack, ensure addRoutes finish
            // set the replace: true so the navigation will not leave a history record
            next({ ...to, replace: true })
          })
        }).catch((err) => {
          store.dispatch('FedLogOut').then(() => {
            Message.error(err || 'Verification failed, please login again')
            next({ path: '/' })
          })
        })
      } else {
        // 没有动态改变权限的需求可直接next() 删除下方权限判断 ↓
        if (hasPermission(store.getters.roles, to.meta.roles)) {
          next()
        } else {
          next({ path: '/401', replace: true, query: { noGoBack: true }})
        }
        // 可删 ↑
      }
    }
  } else {
    /* has no token*/
    if (whiteList.indexOf(to.path) !== -1) { // 在免登录白名单，直接进入
      next()
    } else {
      // next(`/login?redirect=${to.path}`) // 否则全部重定向到登录页
      next(`/home?redirect=${to.path}`)
      NProgress.done() // if current page is login will not trigger afterEach hook, so manually handle it
    }
  }
})

router.afterEach(() => {
  NProgress.done() // finish progress bar
})
