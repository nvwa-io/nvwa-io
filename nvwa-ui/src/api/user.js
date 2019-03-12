import request from '@/utils/request'

export default {
  detail() {
    return request({
      url: '/v1/users/detail',
      method: 'get'
    })
  },
  all() {
    return request({
      url: '/v1/users/all',
      method: 'get'
    })
  }
}
