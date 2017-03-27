<template lang="html">
  <div class="profile-view">
    <div v-if="userProfileExists" class="profile-view-content">
      <sl-profile-header
        :username="activeUser.username"
        :alias="activeUser.alias"
        :icon="activeUser.media.icon.source"
        :banner="activeUser.media.banner.source"></sl-profile-header>
      <sl-profile-nav
        :selected="navOption"
        v-on:select="selectNavOption($event)"></sl-profile-nav>
      <div class="profile-content">
        <sl-profile-home v-if="isSelectedNavOption('home')"
          :contentString="activeUser.description"></sl-profile-home>
        <sl-entry-table :user-entries="activeUser.entries"
          :disable-edit="true"></sl-entry-table>
      </div>
    </div>
    <div v-if="!userProfileExists" class="not-found-content">
      <h2>404 - Looking for something?</h2>
      <h3>Email me at: jamesontreu [at] gmail [dot] com</h3>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

import ProfileHeader from '../components/ProfileHeader'
import ProfileNav from '../components/ProfileNav'
import ProfileHome from '../components/ProfileHome'

import EntryTable from '../components/EntryTable'

export default {
  name: 'profile',
  components: {
    'sl-profile-header': ProfileHeader,
    'sl-profile-nav': ProfileNav,
    'sl-profile-home': ProfileHome,
    'sl-entry-table': EntryTable
  },
  data () {
    return {
      userProfileExists: true,
      navOption: 'home'
    }
  },
  mounted () {
    let username = this.$route.path.substr(1)
    this.$store.dispatch('getUserByUsername', username)
      .catch(() => {
        this.userProfileExists = false
      })
  },
  computed: {
    ...mapGetters([
      'activeUser'
    ])
  },
  methods: {
    selectNavOption (option) {
      this.navOption = option
    },
    isSelectedNavOption (option) {
      return this.navOption === option
    }
  }
}
</script>

<style lang="scss" scoped>

.profile-view {
  text-align: left;
  color: #3D4E5E;
}

.not-found-content {
  position: relative;
  margin: 0 16px;
}
</style>
