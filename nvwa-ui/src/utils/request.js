import axios from 'axios'
import store from '@/store'
import { getToken } from '@/utils/auth'
import { Message, MessageBox } from 'element-ui'

// create an axios instance
const service = axios.create({
  baseURL: process.env.BASE_API, // api 的 base_url
  timeout: 30000, // request timeout
  headers: {},
  params: {},
  data: {},

  // response type
  // e.g: 'arraybuffer', 'blob', 'document', 'json', 'text', 'stream'
  // responseType: 'json', // default
  maxRedirects: 3,

  // core creadentials, default false
  withCredentials: false
})

// request interceptor
service.interceptors.request.use(
  config => {
    if (getToken()) {
      config.headers['Nvwa-Token'] = getToken()
    }
    return config
  },
  error => {
    // Do something with request error
    console.log(error) // for debug
    Promise.reject(error)
  }
)

const errs = {
  SUCCESS: 200,
  INVALID_TOKEN: 401,
  LOGIN_FAILED: 402
}

// response interceptor
service.interceptors.response.use(
  // do nothing with response
  // response => response,
  // check by response code
  response => {
    const res = response.data
    if (res.code !== errs.SUCCESS) {
      // show error message
      // Message({
      //   message: res.msg,
      //   type: 'error',
      //   duration: 5 * 1000
      // })
      // Message.error(res.msg)

      // invalid token
      if (res.code === errs.INVALID_TOKEN) {
        MessageBox.confirm('你已被登出，可以取消继续留在该页面，或者重新登录', '确定登出', {
          confirmButtonText: '重新登录',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          store.dispatch('FedLogOut').then(() => {
            location.reload() // 为了重新实例化vue-router对象 避免bug
          })
        })
      }

      return Promise.reject(res.msg)
    } else {
      return response.data
    }
  },
  error => { // e.g: server is not running, will get error: Error: Network Error
    console.log('error: ' + error)
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
