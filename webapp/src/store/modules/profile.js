import localforage from 'localforage'
import axios from 'axios'

import api from '../../api'
import types from '../mutation-types'
import { updateIconSource, updateBannerSource } from '../../utils/imageSourceManipulator'

const SESSION_TOKEN_KEY = 'session_token'

const state = {
  isLoggedIn: false,
  sessionToken: '',

  registerAttemptStatus: null,
  loginAttemptStatus: null,

  uid: '',
  email: '',
  username: '',
  alias: '',
  description: '',

  entries: { },
  media: { },

  status: '',
  settings: { }, // nightmode, etc..
  deviceSettings: { }
}

const getters = {
  profile (state) {
    return {
      isLoggedIn: state.isLoggedIn,

      uid: state.uid,
      email: state.email,
      username: state.username,
      alias: state.alias,
      status: state.status,
      media: state.media,

      settings: state.settings
    }
  },
  deviceIsMobile (state) {
    return state.deviceSettings.isMobile || false
  },
  deviceIsTablet (state) {
    return state.deviceSettings.isTablet || false
  },
  deviceSettingsInitialized (state) {
    return state.deviceSettings.initialized || false
  }
}
const actions = {
  updateProfileAlias ({ state, commit }, alias) {
    return api.updateProfileAlias({
      username: state.username,
      alias
    })
    .then(profile => {
      commit(types.SET_PROFILE_INFO, profile)
      return 'Updated Full Name'
    })
  },
  updateProfileUsername ({ state, commit }, username) {
    return api.updateProfileUsername({
      oldUsername: state.username,
      newUsername: username
    })
    .then(profile => {
      commit(types.SET_PROFILE_INFO, profile)
      return 'Updated Username'
    })
  },
  updateProfileEmail ({ state, commit }, email) {
    return api.updateProfileEmail({
      username: state.username,
      email
    })
    .then(profile => {
      commit(types.SET_PROFILE_INFO, profile)
      return 'Updated Email'
    })
  },
  updateProfileDescription ({ state, commit }, description) {
    return api.updateProfileDescription({
      username: state.username,
      description
    })
    .then(profile => {
      commit(types.SET_PROFILE_INFO, profile)
      return 'Updated Description'
    })
  },
  updateProfilePassword ({ state, commit }, { oldPassword, newPassword }) {
    return api.updateProfilePassword({
      username: state.username,
      oldPassword,
      newPassword
    })
  },
  updateEntries ({ state, commit }, { entries }) {
    return api.updateProfileEntries({
      username: state.username,
      entries
    })
  },
  updateProfileIcon ({ state, commit }, icon) {
    return api.updateProfileIcon({
      username: state.username,
      icon
    })
    .then(profile => {
      commit(types.SET_PROFILE_INFO, profile)
      return 'Updated Icon'
    })
  },
  updateProfileBanner ({ state, commit }, banner) {
    return api.updateProfileBanner({
      username: state.username,
      banner
    })
    .then(profile => {
      commit(types.SET_PROFILE_INFO, profile)
      return 'Updated Banner'
    })
  },
  checkEmailAvailability (context, email) {
    return api.checkEmailAvailability(email)
  },
  checkUsernameAvailability (context, username) {
    return api.checkUsernameAvailability(username)
  },
  getProfileInfo ({ commit, dispatch }) {
    return dispatch('retrieveSessionToken')
      .then(token => {
        if (!token) {
          return Promise.reject()
        }
        return api.getProfile()
          .then(profile => {
            commit(types.LOGIN_SUCCESS)
            commit(types.SET_PROFILE_INFO, profile)
          })
          .catch(() => {
            commit(types.LOGIN_FAILURE)
            return dispatch('resetProfile')
          })
      })
      .catch(() => {
        // User is not logged in
      })
  },
  register ({ commit, dispatch }, { email, username, alias, password }) {
    return dispatch('resetProfile')
      .then(() => {
        return api.register(email, username, alias, password)
          .then(profile => {
            commit(types.REGISTER_SUCCESS)
            commit(types.LOGIN_SUCCESS)
            commit(types.SET_PROFILE_INFO, profile)

            // Login after registration
            return dispatch('login', { usernameOrEmail: email, password })
          })
          .catch(err => {
            commit(types.REGISTER_FAILURE, err)
            commit(types.LOGIN_FAILURE)

            return Promise.reject(err)
          })
      })
  },
  login ({ commit, dispatch }, { usernameOrEmail, password }) {
    return dispatch('resetProfile')
      .then(() => {
        return api.login(usernameOrEmail, password)
          .then(({ sessionToken, profile }) => {
            commit(types.LOGIN_SUCCESS)
            commit(types.SET_PROFILE_INFO, profile)

            return dispatch('cacheSessionToken', sessionToken)
          })
          .catch(err => {
            commit(types.LOGIN_FAILURE)
            return Promise.reject(err)
          })
      })
  },
  logout ({ commit, dispatch }) {
    return api.logout()
      .then(() => {
        return dispatch('resetProfile')
      })
      .then(() => {
        window.location.reload()
      })
  },
  refreshSessionToken ({ dispatch }) {
    return api.resetSessionToken()
      .then(sessionToken => {
        return dispatch('cacheSessionToken', sessionToken)
      })
      .catch(err => {
        console.error('Error resetting session token: ', err)
      })
  },
  cacheSessionToken ({ commit }, token) {
    return localforage.setItem(SESSION_TOKEN_KEY, token)
      .then(() => {
        return localforage.getItem(SESSION_TOKEN_KEY)
      })
      .then(val => {
        commit(types.SET_SESSION_TOKEN, val)
      })
      .catch(err => {
        console.error('Error setting session token:', err)
      })
  },
  retrieveSessionToken ({ commit }) {
    return localforage.getItem(SESSION_TOKEN_KEY)
      .then(val => {
        commit(types.SET_SESSION_TOKEN, val)
        return val
      })
  },
  resetProfile ({ commit }) {
    commit(types.RESET_PROFILE)
    return localforage.removeItem(SESSION_TOKEN_KEY)
      .then(() => {
        commit(types.REMOVE_SESSION_TOKEN)
      })
      .catch(err => {
        console.error('Error removing session token:', err)
      })
  }
}
const mutations = {
  [types.SET_PROFILE_INFO] (state, profile) {
    state.uid = profile.uid
    state.email = profile.email || state.email
    state.username = profile.username || state.username
    state.alias = profile.alias || state.alias
    state.description = profile.description || state.description

    state.entries = profile.entries || state.entries

    if (profile.media && profile.media.icon) {
      profile.media.icon.source = updateIconSource(profile.media.icon.source)
    }

    if (profile.media && profile.media.banner) {
      profile.media.banner.source = updateBannerSource(profile.media.banner.source)
    }

    state.media = profile.media || state.media || {}
  },
  [types.RESET_PROFILE] (state) {
    state.isLoggedIn = false
    state.uid = ''
    state.email = ''
    state.username = ''
    state.alias = ''
    state.description = ''

    state.entries = {}
    state.media = {}

    state.registerAttemptStatus = null
    state.loginAttemptStatus = null
  },
  [types.REGISTER_SUCCESS] (state) {
    state.registerAttemptStatus = null
  },
  [types.REGISTER_FAILURE] (state, status) {
    state.registerAttemptStatus = status
  },
  [types.LOGIN_SUCCESS] (state) {
    state.isLoggedIn = true
    state.loginAttemptStatus = null
  },
  [types.LOGIN_FAILURE] (state) {
    state.isLoggedIn = false
  },
  [types.SET_SESSION_TOKEN] (state, token) {
    state.sessionToken = token
    axios.defaults.headers.common['Bearer'] = token
  },
  [types.REMOVE_SESSION_TOKEN] (state) {
    state.sessionToken = ''
    axios.defaults.headers.common['Bearer'] = ''
  },
  [types.SET_DEVICE_SETTINGS] (state, settings) {
    // window.alert(window.JSON.stringify(settings))
    settings.initialized = true
    state.deviceSettings = settings
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
