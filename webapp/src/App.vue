<template>
  <div id="app">
    <sl-entry-table></sl-entry-table>
  </div>
</template>

<script>
import EntryTable from './components/EntryTable'

export default {
  name: 'app',
  components: {
    'sl-entry-table': EntryTable
  },
  data () {
    return { }
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
  margin-top: 60px;
}
</style>
