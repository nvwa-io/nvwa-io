import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import HomeLayout from '@/views/layout/HomeLayout'
import Layout from '@/views/layout/Layout'
import SubLayout from '@/views/layout/SubLayout'

import AdminLayout from '@/views/admin/layout/Layout'

Vue.use(Router)

/* Router Modules */

/** note: Submenu only appear when children.length>=1
 *  detail see  https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 **/

/**
 * hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
 *                                if not set alwaysShow, only more than one route under the children
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noredirect           if `redirect:noredirect` will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']     will control the page roles (you can set multiple roles)
    title: 'title'               the name show in submenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar,
    noCache: true                if true ,the page will no be cached(default is false)
  }
 **/
export const constantRouterMap = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/home',
    component: HomeLayout,
    redirect: 'home/index',
    hidden: true,
    children: [
      {
        path: 'index',
        component: () => import('@/views/home/index'),
        name: 'HomeIndex',
        meta: { title: 'deploy.index', icon: 'dashboard' }
      },
      {
        path: 'project',
        component: () => import('@/views/home/project'),
        name: 'HomeProject'
      },
      {
        path: 'audit',
        component: () => import('@/views/home/audit'),
        name: 'HomeAudit'
      }
    ]
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/authredirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/errorPage/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/errorPage/401'),
    hidden: true
  }
]

export default new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})

export const asyncRouterMap = [
  {
    path: '',
    component: Layout,
    redirect: 'deploy/index',
    children: [
      {
        path: 'deploy/index',
        component: () => import('@/views/deploy/index/index'),
        name: 'DeployIndex',
        meta: { title: 'deploy.index', icon: 'dashboard' }
      }
    ]
  },
  {
    path: '/app-list',
    component: Layout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/deploy/appList/index'),
        name: 'app-list/index',
        meta: { title: 'deploy.appList', icon: 'component' }
      },
      {
        path: '',
        component: SubLayout,
        name: 'app-list-sub-page',
        hidden: true,
        meta: { title: 'deploy.appList' },
        redirect: 'index',
        children: [
          {
            path: 'env',
            name: 'app-list/env',
            component: () => import('@/views/deploy/appList/env'),
            meta: { title: 'deploy.env' }
          }
        ]
      }
    ]
  },
  {
    path: '/build',
    component: Layout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/deploy/buildList/index'),
        name: 'build-list/index',
        meta: { title: 'deploy.buildList', icon: 'example' }
      }
    ]
  },
  {
    path: '/env-deploy',
    component: Layout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/deploy/envDeploy/index'),
        name: 'env-deploy/index',
        meta: { title: 'deploy.launchDeploy', icon: 'guide' }
      }
    ]
  },
  {
    path: '/deployment',
    component: Layout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/deploy/deploymentList/index'),
        name: 'deployment-list/index',
        meta: { title: 'deploy.deploymentList', icon: 'nested' }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

// admin end routers
export const adminRouterMap = [
  {
    path: '/admin',
    component: AdminLayout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/admin/index/index'),
        name: 'admin/index',
        meta: { title: 'admin.deploy.dashboard', icon: 'dashboard' }
      }
    ]
  },
  {
    path: '/admin/manage-project',
    component: AdminLayout,
    redirect: 'index',
    meta: { title: 'admin.deploy.manageProject', icon: 'list' },
    children: [
      {
        path: 'index',
        component: () => import('@/views/admin/project/index'),
        name: 'admin/manage-user',
        meta: { title: 'admin.deploy.projectList' }
      },
      {
        path: 'role',
        component: () => import('@/views/admin/project/role'),
        name: 'admin/manage-project-role',
        meta: { title: 'admin.deploy.manageProjectRole' }
      }
    ]
  },
  {
    path: '/admin/manage-app',
    component: AdminLayout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/admin/app/index'),
        name: 'admin/manage-user',
        meta: { title: 'admin.deploy.manageApp', icon: 'component' }
      }
    ]
  },
  {
    path: '/admin/manage-user',
    component: AdminLayout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/admin/user/index'),
        name: 'admin/manage-user',
        meta: { title: 'admin.deploy.manageUser', icon: 'peoples' }
      }
    ]
  },
  {
    path: '/admin/manage-audit',
    component: AdminLayout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/admin/audit/index'),
        name: 'admin/manage-user',
        meta: { title: 'admin.deploy.manageOrder', icon: 'guide' }
      }
    ]
  },
  {
    path: '/admin/system-config',
    component: AdminLayout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/admin/system/index'),
        name: 'admin/manage-user',
        meta: { title: 'admin.deploy.systemConfig', icon: 'lock' }
      }
    ]
  }
]
