<template lang="html">
  <div class="profile-home">
    <div v-html="contentString"></div>
  </div>
</template>

<script>
export default {
  name: 'sl-profile-home',
  props: {
    contentString: {
      type: String,
      default () {
        return ''
      }
    }
  },
  data () {
    return { }
  }
}
</script>

<style lang="scss" scoped>
.profile-home {
  display: block;
  position: relative;
  margin: 0;
  padding: 16px;
  border-right: none;
  border-left: none;
}
</style>
