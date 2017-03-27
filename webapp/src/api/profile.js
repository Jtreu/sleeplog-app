import config from './config'
import axios from 'axios'

const updateAlias = ({ username, alias }) => {
  return axios.post(`${config.apiBase}/users/${username}/alias`, {
    alias
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.user
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const updateUsername = ({ oldUsername, newUsername }) => {
  return axios.post(`${config.apiBase}/users/${oldUsername}/username`, {
    username: newUsername
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.user
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const updateEmail = ({ username, email }) => {
  return axios.post(`${config.apiBase}/users/${username}/email`, {
    email
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.user
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const updateDescription = ({ username, description }) => {
  return axios.post(`${config.apiBase}/users/${username}/description`, {
    description
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.user
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const updatePassword = ({ username, oldPassword, newPassword }) => {
  return axios.post(`${config.apiBase}/users/${username}/password`, {
    oldPassword,
    newPassword
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.message
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const updateEntries = ({ username, entries }) => {
  return axios.post(`${config.apiBase}/users/${username}/entries`, {
    entries
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.message
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const updateIcon = ({ username, icon }) => {
  return axios.post(`${config.staticApiBase}/users/${username}/icon`, icon)
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.user
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const updateBanner = ({ username, banner }) => {
  return axios.post(`${config.staticApiBase}/users/${username}/banner`, banner)
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.user
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const getProfile = () => {
  return axios.get(`${config.apiBase}/sessions`)
    .then(response => {
      let data = response.data

      if (!data.success) {
        return Promise.reject(data.message)
      }

      return data.user
    })
}

const register = (email, username, fullname, password) => {
  return axios.post(`${config.apiBase}/users`, {
    email,
    username,
    fullname,
    password
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.user
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const login = (usernameOrEmail, password) => {
  return axios.post(`${config.apiBase}/sessions`, {
    usernameOrEmail,
    password
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return {
      sessionToken: data.sessionToken,
      profile: data.user
    }
  })
  .catch(err => {
    return Promise.reject(err.response.data.message)
  })
}

const logout = () => {
  return axios.delete(`${config.apiBase}/sessions`)
    .then(response => {
      let data = response.data

      if (!data.success) {
        return Promise.reject(data.message)
      }

      return data.message
    })
    .catch((err) => {
      return Promise.reject(err.response.data.message)
    })
}

const confirmEmail = (code) => {
  return axios.post(`${config.apiBase}/users/email_confirmation`, {
    confirmationCode: code
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.message
  })
  .catch(() => {
    return Promise.reject('Invalid confirmation code')
  })
}

const resetPassword = (email) => {
  return axios.post(`${config.apiBase}/users/password_reset`, {
    email
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.message
  })
  .catch((err) => {
    return Promise.reject(err.response.data.message)
  })
}

const updatePasswordWithCode = (code, password) => {
  return axios.post(`${config.apiBase}/users/password_update`, {
    resetCode: code,
    password
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.message
  })
  .catch((err) => {
    return Promise.reject(err.response.data.message)
  })
}

const sendFeedback = ({ name, email, message }) => {
  return axios.post(`${config.apiBase}/users/feedback`, {
    name,
    email,
    message
  })
  .then(response => {
    let data = response.data

    if (!data.success) {
      return Promise.reject(data.message)
    }

    return data.message
  })
  .catch((err) => {
    return Promise.reject(err.response.data.message)
  })
}

export default {
  updateProfileAlias: updateAlias,
  updateProfileUsername: updateUsername,
  updateProfileEmail: updateEmail,
  updateProfileDescription: updateDescription,
  updateProfilePassword: updatePassword,
  updateProfileEntries: updateEntries,
  updateProfileIcon: updateIcon,
  updateProfileBanner: updateBanner,

  getProfile,

  register,
  login,
  logout,
  confirmEmail,
  resetPassword,
  updatePasswordWithCode,

  sendProfileFeedback: sendFeedback
}
