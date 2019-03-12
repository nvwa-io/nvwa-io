/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'
import SubLayout from '@/views/layout/SubLayout'

const router = {
  path: '/deploy-system',
  component: Layout,
  // redirect: '/service/complex-table',
  name: 'deploySystem',
  redirect: 'noredirect',
  meta: {
    title: 'deploySystem',
    icon: 'component'
  },
  children: [
    {
      path: 'index',
      component: () => import('@/views/deploy/index/index'),
      name: 'index',
      meta: { title: 'deploy.index' }
    },
    {
      path: 'app-list/index',
      component: () => import('@/views/deploy/appList/index'),
      name: 'app-list/index',
      meta: { title: 'deploy.appList' }
    },
    {
      path: 'app-list',
      component: SubLayout,
      name: 'app-list',
      hidden: true,
      meta: { title: 'deploy.appList' },
      redirect: '/deploy-system/app-list/index',
      children: [
        {
          path: 'env',
          name: 'app-list/env',
          component: () => import('@/views/deploy/appList/env'),
          meta: { title: 'deploy.env' }
        }
      ]
    },
    {
      path: 'build-list/index',
      component: () => import('@/views/deploy/buildList/index'),
      name: 'build-list/index',
      meta: { title: 'deploy.buildList' }
    },
    {
      path: 'env-deploy/index',
      component: () => import('@/views/deploy/envDeploy/index'),
      name: 'env-deploy/index',
      meta: { title: 'deploy.launchDeploy' }
    },
    {
      path: 'deployment-list/index',
      component: () => import('@/views/deploy/deploymentList/index'),
      name: 'deployment-list/index',
      meta: { title: 'deploy.deploymentList' }
    }
  ]
}

export default router
