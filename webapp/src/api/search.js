import config from './config'
import axios from 'axios'

const searchUsers = (query) => {
  return axios.get(`${config.apiBase}/users?q=${query}`)
    .then(response => {
      let data = response.data

      if (!data.success) {
        return Promise.reject(data.message)
      }

      return data.users
    })
    .catch(err => {
      return Promise.reject(err.response.data.message)
    })
}

export default {
  searchUsers
}
