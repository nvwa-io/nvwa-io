import request from '@/utils/request'

export default {

  listWait() {
    return request({
      url: '/v1/audits/wait',
      method: 'get'
    })
  },
  getWaitAuditNum() {
    return request({
      url: '/v1/audits/wait-num',
      method: 'get'
    })
  },
  listMine() {
    return request({
      url: '/v1/audits/mine',
      method: 'get'
    })
  },
  listAudited() {
    return request({
      url: '/v1/audits/audited',
      method: 'get'
    })
  },
  pass(auditId) {
    return request({
      url: '/v1/audits/' + auditId + '/pass',
      method: 'put'
    })
  },
  reject(auditId) {
    return request({
      url: '/v1/audits/' + auditId + '/reject',
      method: 'put'
    })
  },
  cancel(auditId) {
    return request({
      url: '/v1/audits/' + auditId + '/cancel',
      method: 'put'
    })
  }
}
