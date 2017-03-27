<template lang="html">
  <div class="search-view">
    <sl-user-list
      :user-list="searchResults"></sl-user-list>
  </div>
</template>

<script>
import { mapState } from 'vuex'

import UserList from '../components/UserList'

export default {
  name: 'search-view',
  components: {
    'sl-user-list': UserList
  },
  data () {
    return { }
  },
  computed: {
    ...mapState({
      searchResults: state => {
        let list = state.users.searchResults
        list = list.map(user => {
          return {
            imgPath: user.media.icon.source,
            name: user.alias,
            username: user.username,
            usertype: user.role || 'Patient'
          }
        })
        return list
      }
    })
  }
}
</script>

<style lang="scss" scoped>
.search-view {

}
</style>
