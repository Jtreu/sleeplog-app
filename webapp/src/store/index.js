import Vue from 'vue'
import Vuex from 'vuex'
import actions from './actions'
import getters from './getters'
import profile from './modules/profile'
import users from './modules/users'

Vue.use(Vuex)

const store = new Vuex.Store({
  actions,
  getters,
  modules: {
    profile,
    users
  }
})

export default store
