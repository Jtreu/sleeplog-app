import api from '../api'

/* Profile */

const confirmEmail = ({ state }, code) => {
  return api.confirmEmail(code)
}

const resetPassword = ({ state }, email) => {
  return api.resetPassword(email)
}

const updatePasswordWithCode = ({ state }, { code, password }) => {
  return api.updatePasswordWithCode(code, password)
}

/* Users */

const ensureUsers = ({ state, dispatch }, uids) => {
  return dispatch('getUsersByUIDs', uids)
}

export default {
  confirmEmail,
  resetPassword,
  updatePasswordWithCode,

  ensureUsers
}
