import request from '@/utils/request'

export default {
  all() {
    return request({
      url: '/v1/project-role-perms/admin/',
      method: 'get'
    })
  },

  batchUpdate(projectRoleId, projectRoleName, perms) {
    return request({
      url: '/v1/project-role-perms/admin/batch-update',
      method: 'post',
      data: JSON.stringify({
        project_role_id: projectRoleId,
        project_role_name: projectRoleName,
        perms: perms
      })
    })
  },

  batchCreate(projectRoleName, perms) {
    return request({
      url: '/v1/project-role-perms/admin/batch-create',
      method: 'post',
      data: JSON.stringify({
        project_role_name: projectRoleName,
        perms: perms
      })
    })
  }
}
