<template lang="html">
  <div class="profile-home">
    <button v-on:click="print()" class="button_print">Print</button>
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
  },
  methods: {
    print () {
      window.print()
    }
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
.button_print {
  width: 200px;
  background-color: #c34949;
  color: #fff;
  box-shadow: 10px 10px 5px #888888;
  overflow: auto;
  padding: 14px 20px;
  font-size: 13px;
  font-weight: bold;
  margin: 8px 0;
  border: none;
  border-radius: 2px;
  cursor: pointer;
  text-transform: uppercase;
  text-decoration: none;
}
</style>
