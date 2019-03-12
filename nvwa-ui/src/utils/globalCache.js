
const KEY_LATEST_PROJECT = 'latest_project'

export default {
  getLatestProject() {
    const data = localStorage.getItem(KEY_LATEST_PROJECT)
    const res = JSON.parse(data)
    if (!res) {
      return false
    }

    return res
  },

  setLatestProject(project) {
    localStorage.setItem(KEY_LATEST_PROJECT, JSON.stringify(project))
  },

  removeLatestProject() {
    localStorage.removeItem(KEY_LATEST_PROJECT)
  }
}
