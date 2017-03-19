<template>
  <div id="app">
    <sl-entry-dialog class="dialog"
      :open="dialogIsOpen"
      :entry="activeEntry"
      :default-activites="defaultActivites"
      v-on:add="addActivity()"
      v-on:remove="removeActivity($event)"
      v-on:close="closeDialog()"></sl-entry-dialog>
    <table class="table">
      <tr>
        <th>DATE:</th>
        <th v-for="time in times">{{time}}</th>
      </tr>
      <tr v-for="date in dates">
        <td>{{date}}</td>
        <td v-for="time in times" v-on:click="openDialog(date, time)"></td>
      </tr>
    </table>
  </div>
</template>

<script>
import EntryDialog from './components/EntryDialog'

export default {
  name: 'app',
  components: {
    'sl-entry-dialog': EntryDialog
  },
  data () {
    return {
      entries: {},
      times: [],
      dates: [],
      dialogIsOpen: false,
      activeEntry: { date: '', time: '', activities: [] },
      defaultActivites: ['Sleep', 'Caffeine', 'Medicine']
    }
  },
  mounted () {
    // Initialize times array
    this.generateTimes()

    // Initialize dates array
    this.generateDates()

    // Initialize entries object
    this.initEntries()
  },
  methods: {
    initEntries () {
      var entries = {}
      var defaultActivity = {
        name: 'Sleep',
        isDone: false
      }

      for (var i = 0; i < this.dates.length; i++) {
        entries[`${this.dates[i]}`] = {}
        for (var j = 0; j < this.times.length; j++) {
          entries[`${this.dates[i]}`][`${this.times[j]}`] = {
            date: this.dates[i],
            time: this.times[j],
            activities: [defaultActivity]
          }
        }
      }

      this.entries = entries
    },
    generateTimes () {
      var am = []
      var pm = []

      for (var i = 1; i < 12; i++) {
        am.push(i + 'a')
        pm.push(i + 'p')
      }

      am.push('NOON')
      pm.push('MIDNIGHT')

      this.times = am.concat(pm)
    },
    generateDates () {
      var dates = [new Date()]

      for (var i = 1; i < 10; i++) {
        // get the next day after the previous
        var thisDate = new Date(dates[i - 1].getTime() - (24 * 60 * 60 * 1000))
        dates.push(thisDate)
      }

      for (var j = 0; j < dates.length; j++) {
        dates[j] = dates[j].toLocaleDateString()
      }
      this.dates = dates
    },
    openDialog (date, time) {
      this.activeEntry = this.entries[date][time]
      this.dialogIsOpen = true
    },
    closeDialog () {
      this.dialogIsOpen = false
    },
    addActivity () {
      if (!this.activeEntry.date || !this.activeEntry.time) {
        return
      }

      // Don't add more activities than available
      if (this.entries[this.activeEntry.date][this.activeEntry.time].activities.length >= this.defaultActivites.length) {
        return
      }

      let availableName = this.getAvailableActivityName(this.entries[this.activeEntry.date][this.activeEntry.time].activities)

      this.entries[this.activeEntry.date][this.activeEntry.time].activities.push({
        name: availableName,
        isDone: false
      })
    },
    removeActivity (index) {
      if (index > -1) {
        this.entries[this.activeEntry.date][this.activeEntry.time].activities.splice(index, 1)
      }
    },
    getAvailableActivityName (selected) {
      let options = this.defaultActivites
      selected = JSON.parse(JSON.stringify(selected))
      selected = Object.keys(selected).map(key => selected[key].name) // e.g. ['Medicine', 'Sleep']

      options = options.filter(i => selected.indexOf(i) < 0)
      return options[0]
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
  margin-top: 60px;
}

table {
    font-family: arial, sans-serif;
    border-collapse: collapse;
    width: 100%;
}

td, th {
    border: 1px solid #dddddd;
    text-align: left;
    padding: 8px;
}
</style>
