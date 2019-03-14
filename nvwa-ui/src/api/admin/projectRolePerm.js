import request from '@/utils/request'

export default {
  all() {
    return request({
      url: '/v1/project-role-perms/admin/',
      method: 'get'
    })
  }
}
