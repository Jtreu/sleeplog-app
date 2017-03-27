<template lang="html">
  <div class="home-view">
    <div v-if="!isLoggedIn" class="register-login-view">
      <ep-signin-dialog v-if="selectedDialog === dialogTypes.signin"
        v-on:dialog-selected="selectDialog($event)"
        class="signin-dialog"></ep-signin-dialog>
      <ep-password-reset-dialog v-if="selectedDialog === dialogTypes.passwordReset"
        v-on:back="selectDialog('signin')"
        class="password-reset-dialog"></ep-password-reset-dialog>
      <ep-signup-dialog v-if="selectedDialog === dialogTypes.signup"
        v-on:dialog-selected="selectDialog($event)"
        class="signup-dialog"></ep-signup-dialog>
    </div>
    <div v-else-if="isLoggedIn"
      class="profile-view">
      <sl-profile-header
        :username="profile.username"
        :alias="profile.alias"
        :icon="profile.media.icon.source"
        :banner="profile.media.banner.source"></sl-profile-header>
      <sl-profile-nav
        :selected="navOption"
        v-on:select="selectNavOption($event)"></sl-profile-nav>
      <div class="profile-content">
        <sl-profile-home v-if="isSelectedNavOption('home')"
          :contentString="profile.description"></sl-profile-home>
      </div>
      <sl-entry-table></sl-entry-table>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'
import EntryTable from '../components/EntryTable'

import SignInDialog from '../components/SignInDialog'
import SignUpDialog from '../components/SignUpDialog'
import PasswordResetDialog from '../components/PasswordResetDialog'

import ProfileHeader from '../components/ProfileHeader'
import ProfileNav from '../components/ProfileNav'
import ProfileHome from '../components/ProfileHome'

export default {
  name: 'home-view',
  components: {
    'sl-entry-table': EntryTable,
    'ep-signin-dialog': SignInDialog,
    'ep-signup-dialog': SignUpDialog,
    'ep-password-reset-dialog': PasswordResetDialog,
    'sl-profile-header': ProfileHeader,
    'sl-profile-nav': ProfileNav,
    'sl-profile-home': ProfileHome
  },
  data () {
    return {
      navOption: 'home',
      selectedDialog: 'signin',
      dialogTypes: {
        signin: 'signin',
        signup: 'signup',
        passwordReset: 'passwordReset'
      }
    }
  },
  computed: {
    ...mapState({
      isLoggedIn: state => state.profile.isLoggedIn
    }),
    ...mapGetters([
      'profile'
    ])
  },
  methods: {
    selectNavOption (option) {
      this.navOption = option
    },
    isSelectedNavOption (option) {
      return this.navOption === option
    },
    selectDialog (dialog, data) {
      this.selectedDialog = dialog

      // Media dialog
      if ((dialog === this.dialogTypes.media) && data) {
        this.activeMedia.type = data.type
        this.activeMedia.source = data.source
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.home-view {
  text-align: left;
  width: 100%;
}

.signin-dialog, .password-reset-dialog, .signup-dialog {
  margin: 0 auto;
}
</style>
