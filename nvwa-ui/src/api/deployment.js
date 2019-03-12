import request from '@/utils/request'

export default {

  listByProjectId(projectId, ...appId) {
    let url = '/v1/deployments/project/' + projectId
    if (appId.length > 0) {
      url += '?app_id=' + appId[0]
    }
    return request({
      url: url,
      method: 'get'
    })
  }
}
