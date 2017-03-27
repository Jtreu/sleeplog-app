import Vue from 'vue'
import api from '../../api'

import types from '../mutation-types'
import dataTypes from '../data-types'
import { updateIconSource, updateBannerSource } from '../../utils/imageSourceManipulator'

const state = {
  all: {/* [uid: string]: User */},
  activeUserUID: '',

  foundActiveUser: false,

  searchQuery: '',
  searchResults: []
}

const getters = {
  activeUser (state) {
    return state.all[state.activeUserUID] || dataTypes.defaultUser
  },
  foundActiveUser (state) {
    return state.foundActiveUser
  }
}

const actions = {
  searchUsers ({ state, commit }, query) {
    commit(types.SET_SEARCH_QUERY, query)

    return api.searchUsers(query)
      .then(results => {
        commit(types.SET_SEARCH_RESULTS, results)
        return 'Results achieved'
      })
      .catch(err => {
        console.log(err)
      })
  },
  getUserByUsername ({ state, commit }, username) {
    let uid = null

    for (let u in state.all) {
      if (state.all[u].username.toLowerCase() === username.toLowerCase()) {
        uid = state.all[u].uid
        break
      }
    }

    if (uid && state.all[uid]) {
      commit(types.SET_ACTIVE_USER, uid)
      commit(types.SET_FOUND_ACTIVE_USER, true)
      return Promise.resolve(state.all[uid])
    }

    return api.getUserByUsername(username)
      .then(user => {
        commit(types.SET_USER, user)
        commit(types.SET_ACTIVE_USER, user.uid)
        commit(types.SET_FOUND_ACTIVE_USER, true)
        return user
      })
      .catch(err => {
        commit(types.SET_FOUND_ACTIVE_USER, false)
        return Promise.reject(err)
      })
  },
  getUsersByUIDs ({ commit }, uids) {
    uids = uids.filter(uid => {
      if (!state.all[uid]) {
        return true
      }
    })

    if (uids.length < 1) {
      return Promise.resolve([])
    }

    return api.getUsersByUIDs(uids)
      .then(users => {
        commit(types.SET_USERS, users)
        return users
      })
      .catch(err => console.error(err))
  }
}

const mutations = {
  [types.SET_SEARCH_RESULTS] (state, results) {
    state.searchResults = results
  },
  [types.SET_SEARCH_QUERY] (state, query) {
    state.searchQuery = query
  },
  [types.SET_FOUND_ACTIVE_USER] (state, isFound) {
    state.foundActiveUser = isFound
  },
  [types.SET_ACTIVE_USER] (state, uid) {
    state.activeUserUID = uid
  },
  [types.SET_USERS] (state, users) {
    users.forEach(user => {
      if (user.media && user.media.icon) {
        user.media.icon.source = updateIconSource(user.media.icon.source)
      }

      if (user.media && user.media.banner) {
        user.media.banner.source = updateBannerSource(user.media.banner.source)
      }

      if (user) {
        Vue.set(state.all, user.uid, user)
      }
    })
  },
  [types.SET_USER] (state, user) {
    if (user.media && user.media.icon) {
      user.media.icon.source = updateIconSource(user.media.icon.source)
    }

    if (user.media && user.media.banner) {
      user.media.banner.source = updateBannerSource(user.media.banner.source)
    }

    if (user) {
      Vue.set(state.all, user.uid, user)
    }
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
