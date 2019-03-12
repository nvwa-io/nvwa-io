import request from '@/utils/request'

export default {

  create(cluster) {
    return request({
      url: '/v1/clusters',
      method: 'post',
      data: JSON.stringify({
        cluster: cluster
      })
    })
  },
  update(clusterId, cluster) {
    return request({
      url: '/v1/clusters/' + clusterId,
      method: 'put',
      data: JSON.stringify({
        cluster: cluster
      })
    })
  }
}
