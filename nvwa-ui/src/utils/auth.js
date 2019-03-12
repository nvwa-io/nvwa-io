// import Cookies from 'js-cookie'

// const TokenKey = 'Admin-Token'

const KEY_TOKEN = 'Nvwa-Token'

export function getToken() {
  const data = localStorage.getItem(KEY_TOKEN)
  const res = JSON.parse(data)
  const timestamp = new Date().getTime() / 1000 // second
  if (!res || res.expire < timestamp) { // no token or token has expired
    return false
  }

  return res.token
}

export function setToken(token, expire) {
  const data = { token: token, expire: expire }
  localStorage.setItem(KEY_TOKEN, JSON.stringify(data))
}

export function removeToken() {
  localStorage.removeItem(KEY_TOKEN)
}
