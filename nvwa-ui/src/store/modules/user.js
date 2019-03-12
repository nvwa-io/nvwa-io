import { loginByUsername, logout, getUserInfo } from '@/api/login'
import { getToken, setToken, removeToken } from '@/utils/auth'
import apiSite from '@/api/site'
import apiUser from '@/api/user'
import apiAudit from '@/api/audit'
import { Message } from 'element-ui'

const user = {
  state: {
    curUser: {},
    token: getToken(),
    roles: [],
    auditNum: 0
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_CUR_USER: (state, user) => {
      state.curUser = user
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_AUDIT_NUM: (state, auditNum) => {
      state.auditNum = auditNum
    }
  },

  actions: {
    // login and set token to local storage
    Login({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        apiSite.login(userInfo.loginType, userInfo.username, userInfo.password).then(response => {
          console.log(response)

          const data = response.data
          commit('SET_TOKEN', data.token)
          setToken(response.data.token, response.data.expire)

          resolve(response)
        }).catch(error => {
          console.log(error)
          reject(error)
        })
      })
    },

    // 用户名登录
    LoginByUsername({ commit }, userInfo) {
      return new Promise((resolve, reject) => {
        loginByUsername(userInfo.username.trim(), userInfo.password).then(response => {
          const data = response.data
          commit('SET_TOKEN', data.token)
          // @TODO token expire
          setToken(response.data.token, 10000000000000000)
          console.log(getToken())
          resolve()
        }).catch(error => {
          console.log(error)
          reject(error)
        })
      })
    },
    GetWaitAuditNum({ commit }) {
      return new Promise((resolve, reject) => {
        apiAudit.getWaitAuditNum().then(response => {
          commit('SET_AUDIT_NUM', response.data.num)
        }).catch(error => {
          console.log(error)
        })
      })
    },

    // 获取用户信息
    GetUserInfo({ commit, state }) {
      return new Promise((resolve, reject) => {
        apiUser.detail().then(response => {
          // @TODO user role
          commit('SET_ROLES', [response.data.user.role])
          commit('SET_CUR_USER', response.data.user)

          resolve(response)
        }).catch(error => {
          console.log(error)
          Message.error(error)
          reject(error)
        })

        // getUserInfo(state.token).then(response => {
        //   // if (!response.data) { // 由于mockjs 不支持自定义状态码只能这样hack
        //   //   reject('error')
        //   // }
        //   const data = response.data
        //
        //
        //   commit('SET_NAME', data.name)
        //   commit('SET_AVATAR', data.avatar)
        //   commit('SET_INTRODUCTION', data.introduction)
        //   resolve(response)
        // }).catch(error => {
        //   reject(error)
        // })
      })
    },

    // 登出
    LogOut({ commit, state }) {
      return new Promise((resolve, reject) => {
        logout(state.token).then(() => {
          commit('SET_TOKEN', '')
          commit('SET_ROLES', [])
          removeToken()
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },

    // 前端 登出
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        removeToken()
        resolve()
      })
    },

    // @TODO 动态修改权限
    ChangeRoles({ commit, dispatch }, role) {
      return new Promise(resolve => {
        commit('SET_TOKEN', role)
        setToken(role)
        getUserInfo(role).then(response => {
          const data = response.data
          commit('SET_ROLES', data.roles)
          // commit('SET_NAME', data.name)
          // commit('SET_AVATAR', data.avatar)
          // commit('SET_INTRODUCTION', data.introduction)
          dispatch('GenerateRoutes', data) // 动态修改权限后 重绘侧边菜单
          resolve()
        })
      })
    }
  }
}

export default user
