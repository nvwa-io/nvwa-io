import request from '@/utils/request'

export default {

  create(env) {
    return request({
      url: '/v1/envs',
      method: 'post',
      data: JSON.stringify({
        env: env
      })
    })
  },
  update(envId, env) {
    return request({
      url: '/v1/envs/' + envId,
      method: 'put',
      data: JSON.stringify({
        env: env
      })
    })
  }
}
