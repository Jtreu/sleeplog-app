<template lang="html">
  <div class="settings-profile">
    <button class="collapsible-title-button"
      v-on:click="toggleExpanded()"
      v-bind:class="{ expanded: expanded, condensed: !expanded }">
      <span>{{ label.PROFILE }}</span>
    </button>
    <div v-show="expanded" class="content">
      <sl-settings-profile-description></sl-settings-profile-description>
      <hr class="section-division">
      <sl-settings-profile-basic></sl-settings-profile-basic>
      <hr class="section-division">
      <sl-settings-profile-password></sl-settings-profile-password>
      <hr class="section-division">
      <input type="submit"
        id="logoutBtn"
        :value="label.LOGOUT"
        v-on:click="logout()">
    </div>
  </div>
</template>

<script>
import SettingsProfileBasic from './SettingsProfileBasic'
import SettingsProfileDescription from './SettingsProfileDescription'
import SettingsProfilePassword from './SettingsProfilePassword'

export default {
  name: 'sl-settings-profile',
  components: {
    'sl-settings-profile-basic': SettingsProfileBasic,
    'sl-settings-profile-description': SettingsProfileDescription,
    'sl-settings-profile-password': SettingsProfilePassword
  },
  data () {
    return {
      expanded: true,
      label: {
        PROFILE: 'Profile',
        LOGOUT: 'Logout'
      }
    }
  },
  methods: {
    logout () {
      this.$store.dispatch('logout')
        .then(() => {
          this.goHome()
        })
        .catch(() => {
          this.goHome()
        })
    },
    goHome () {
      this.$router.push({ name: 'home' })
    },
    toggleExpanded () {
      this.expanded = !this.expanded
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';
@import '../assets/styles/input';

.settings-profile {
  display: block;
  position: relative;
  margin: 0;
  padding: 16px;
  border-right: none;
  border-left: none;
  border-bottom: 1px solid #dcdcdc;
}

#logoutBtn {
  background-color: $orange;
}

.content {
  margin-top: 20px;
}

.section-division {
  margin: 24px 0;
  border: none;
  border-bottom: 1px solid $dark-gray;
}
</style>
