import config from './config'
import axios from 'axios'

const checkEmailAvailability = (email) => {
  return axios.get(`${config.apiBase}/users?field=email&q=${email}`)
    .then(response => {
      let data = response.data

      // If no user is found, email is available
      if (!data.success) {
        return Promise.resolve(true)
      }

      return data.users.length < 1
    })
    .catch(() => {
      return Promise.reject(true)
    })
}

const checkUsernameAvailability = (username) => {
  return axios.get(`${config.apiBase}/users/${username}`)
    .then(response => {
      let data = response.data

      // If no user is found, username is available
      if (!data.success) {
        return true
      }

      return false
    })
    .catch(() => {
      return true
    })
}

const getUserByUsername = (username) => {
  return axios.get(`${config.apiBase}/users/${username}`)
    .then(response => {
      let data = response.data

      if (!data.success) {
        return Promise.reject(data.message)
      }

      return data.user
    })
}

const getUsersByUIDs = (uids) => {
  let promises = []

  for (let i = 0; i < uids.length; i++) {
    promises.push(
      axios.get(`${config.apiBase}/users?field=uid&q=${uids[i]}`)
        .then(response => {
          let data = response.data

          if (!data.success) {
            return Promise.reject(data.message)
          }

          return data.users
        })
    )
  }

  return Promise.all(promises)
    .then(arrays => {
      let users = [].concat.apply([], arrays)
      return users
    })
}

export default {
  checkEmailAvailability,
  checkUsernameAvailability,
  getUserByUsername,
  getUsersByUIDs
}
