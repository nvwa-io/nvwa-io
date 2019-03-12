import request from '@/utils/request'

export default {
// {
//   "project_id": 8,
//   "name": "demo-02",
//   "description": "测试 APP",
//   "repo_url": "git@github.com/demo-01.git",
//   "repo_username": "hikoqiu",
//   "repo_password": "@hikoqiu",
//   "deploy_user": "nvwa",
//   "deploy_path": "/data/nvwa/deploy/demo-01",
// }
  create(app) {
    return request({
      url: '/v1/apps',
      method: 'post',
      data: JSON.stringify({
        app: app
      })
    })
  },

  // @TODO update app
  update(id, name, description) {
    // return request({
    //   url: '/v1/projects/' + id,
    //   method: 'put',
    //   data: JSON.stringify({
    //     project: {
    //       name: name,
    //       description: description
    //     }
    //   })
    // })
  },

  updateCmds(id, params) {
    return request({
      url: '/v1/apps/' + id + '/commands',
      method: 'put',
      data: JSON.stringify(params)
    })
  },

  getById(id) {
    return request({
      url: '/v1/apps/' + id,
      method: 'get'
    })
  },

  // get branches by app id
  getBranches(id) {
    return request({
      url: '/v1/apps/' + id + '/branches',
      method: 'get'
    })
  },

  listByProjectId(projectId) {
    return request({
      url: '/v1/apps/project/' + projectId,
      method: 'get'
    })
  },

  listAppEnvsByProjectId(projectId) {
    return request({
      url: '/v1/apps/app-and-envs/project/' + projectId,
      method: 'get'
    })
  },

  listAppEnvsByAppId(appId) {
    return request({
      url: '/v1/apps/app-and-envs/app/' + appId,
      method: 'get'
    })
  },

  deleteById(id) {
    return request({
      url: '/v1/apps/' + id,
      method: 'delete'
    })
  }
}
