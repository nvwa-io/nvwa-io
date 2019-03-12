import request from '@/utils/request'

export default {
  // loginType: ldap|default
  login(loginType, username, password) {
    return request({
      url: '/v1/site/login',
      method: 'post',
      data: JSON.stringify({
        username: username,
        password: password,
        'login-type': loginType
      })
    })
  },

  register(username, email, password) {
    return request({
      url: '/v1/site/register',
      method: 'post',
      data: JSON.stringify({
        username: username,
        email: email,
        password: password
      })
    })
  }
}
