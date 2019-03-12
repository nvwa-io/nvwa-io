import request from '@/utils/request'

export default {

  listByProjectId(projectId, ...appId) {
    let url = '/v1/builds/project/' + projectId
    if (appId.length > 0) {
      url += '?app_id=' + appId[0]
    }
    return request({
      url: url,
      method: 'get'
    })
  },

  create(appId, branch) {
    const url = '/v1/builds'
    return request({
      url: url,
      method: 'post',
      data: JSON.stringify({
        'build': {
          app_id: appId,
          branch: branch
        }
      })
    })
  }
}
