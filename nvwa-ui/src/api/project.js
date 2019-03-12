import request from '@/utils/request'

export default {
  create(name, description) {
    return request({
      url: '/v1/projects',
      method: 'post',
      data: JSON.stringify({
        project: {
          name: name,
          description: description
        }
      })
    })
  },

  update(id, name, description) {
    return request({
      url: '/v1/projects/' + id,
      method: 'put',
      data: JSON.stringify({
        project: {
          name: name,
          description: description
        }
      })
    })
  },

  getById(id) {
    return request({
      url: '/v1/projects/' + id,
      method: 'get'
    })
  },

  deleteById(id) {
    return request({
      url: '/v1/projects/' + id,
      method: 'delete'
    })
  },

  // get projects of current login user
  list() {
    return request({
      url: '/v1/projects/',
      method: 'get'
    })
  }
}
