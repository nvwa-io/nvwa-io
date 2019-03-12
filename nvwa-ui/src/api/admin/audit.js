import request from '@/utils/request'

export default {
  listByStatus(status) {
    return request({
      url: '/v1/audits/admin/status/' + status,
      method: 'get'
    })
  }
}
