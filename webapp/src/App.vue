<template>
  <div id="app">
    <sl-header></sl-header>
    <router-view class="view"></router-view>
  </div>
</template>

<script>
import Header from './components/Header'

import device from './utils/device'

export default {
  name: 'app',
  components: {
    'sl-header': Header
  },
  data () {
    return { }
  },
  mounted () {
    this.initProfile()
  },
  methods: {
    initProfile () {
      this.$store.dispatch('getProfileInfo')
    },
    initDeviceSettings () {
      this.$store.commit('SET_DEVICE_SETTINGS', device.getSettings())
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body {
  margin: 0;
}
</style>
