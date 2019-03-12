import request from '@/utils/request'

export default {

  listByEnvId(envId, ...limit) {
    let url = '/v1/pkgs/env/' + envId
    if (limit.length > 0) {
      url += '?limit=' + limit[0]
    }
    return request({
      url: url,
      method: 'get'
    })
  }
}
