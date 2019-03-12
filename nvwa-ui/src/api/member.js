import request from '@/utils/request'

export default {

  create(member) {
    console.log('member:', member)
    return request({
      url: '/v1/members',
      method: 'post',
      data: JSON.stringify({
        member: member
      })
    })
  },

  getListByProjectId(projectId) {
    return request({
      url: '/v1/members/project/' + projectId,
      method: 'get'
    })
  },

  editRole(memberId, projectRoleId) {
    return request({
      url: '/v1/members/' + memberId + '/role',
      method: 'put',
      data: JSON.stringify({
        member: {
          project_role_id: projectRoleId
        }
      })
    })
  },
  removeMember(memberId) {
    return request({
      url: '/v1/members/' + memberId,
      method: 'delete'
    })
  }

}
