import request from '@/utils/request'

export default {

  get() {
    const url = '/v1/systems/'
    return request({
      url: url,
      method: 'get'
    })
  }
}
