const project = {
  state: {
    curProject: {
      id: 0
    }
  },
  mutations: {
    SET_CUR_PROJECT: (state, project) => {
      state.curProject = project
    }
  },
  actions: {
    SetCurProject({ commit }, project) {
      return new Promise((resolve) => {
        commit('SET_CUR_PROJECT', project)
        resolve(project)
      }).catch(error => {
        console.log(error)
      })
    }
  }
}

export default project
