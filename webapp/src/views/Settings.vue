<template lang="html">
  <div class="settings-view">
    <sl-settings-profile v-if="profile.isLoggedIn"></sl-settings-profile>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

import SettingsProfile from '../components/SettingsProfile'

export default {
  name: 'sl-settings-view',
  components: {
    'sl-settings-profile': SettingsProfile
  },
  data () {
    return { }
  },
  computed: {
    ...mapGetters([
      'profile'
    ])
  }
}
</script>

<style lang="scss" scoped>
.settings-view {

}
</style>
