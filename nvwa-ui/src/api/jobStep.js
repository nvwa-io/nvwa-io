import request from '@/utils/request'

export default {

  getByJobId(jobId) {
    const url = '/v1/job-steps/job/' + jobId
    return request({
      url: url,
      method: 'get'
    })
  }
}
