import request from '@/utils/request'

export default {

  start(jobId) {
    const url = '/v1/jobs/' + jobId + '/start'
    console.log(url)
    return request({
      url: url,
      method: 'put'
    })
  }
}
