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
        <th>Day of the Week</th>
        <th>Type of Day
            <p id="day_type">Work, School, Off, Vacation (Vac)</p>
        </th>
        <th class="verticalTableHeader" v-for="time in times">{{time}}</th>
      </tr>
      <tr v-for="date in dates">
        <td>{{date}}</td>
        <!-- Day of Wk/Type of Day Entries -->
        <td></td>
        <td></td>
        <!-- Timeslot Entries -->
        <td v-for="time in times"
          :class="{ 'shaded': doesEntryContainSleep(date, time) }"
          v-on:click="openDialog(date, time)">{{displayInBox(date, time).toString()}}</td>
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
      defaultActivites: ['Sleep', 'Caffeine', 'Medicine', 'Alcohol', 'Exercise']
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
    doesEntryContainSleep (date, time) {
      if (!this.entries || !this.entries[date][time] || this.entries[date][time].activities.length < 1) {
        return false
      }

      let activities = this.entries[date][time].activities
      activities = Object.keys(activities).map(key => activities[key])

      for (let i = 0; i < activities.length; i++) {
        if (activities[i].name === 'Sleep' && activities[i].isDone) {
          return true
        }
      }

      return false
    },

    displayInBox (date, time) {
      let activities = this.entries[date][time].activities
      activities = Object.keys(activities).map(key => activities[key])
      let displayThis = []

      for (let i = 0; i < activities.length; i++) {
        if (activities[i].name === 'Caffeine' && activities[i].isDone) {
          displayThis.push('C')
        }
        if (activities[i].name === 'Medicine' && activities[i].isDone) {
          displayThis.push('M')
        }

        if (activities[i].name === 'Alcohol' && activities[i].isDone) {
          displayThis.push('A')
        }
        if (activities[i].name === 'Exercise' && activities[i].isDone) {
          displayThis.push('E')
        }
        if (activities[i].name === 'Sleep' && activities[i].isDone) {
          displayThis = []
          break
        }
      }
      return displayThis
    },
    initEntries () {
      var entries = {}

      for (var i = 0; i < this.dates.length; i++) {
        entries[`${this.dates[i]}`] = {}
        for (var j = 0; j < this.times.length; j++) {
          entries[`${this.dates[i]}`][`${this.times[j]}`] = {
            date: this.dates[i],
            time: this.times[j],
            activities: []
          }
        }
      }

      this.entries = entries
    },
    generateTimes () {
      var am = []
      var pm = []
      pm.push('NOON')

      for (var i = 1; i < 12; i++) {
        am.push(i + 'a')
        pm.push(i + 'p')
      }

      pm.push('MIDNIGHT')

      this.times = pm.concat(am)
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
    generateDayOfWeek () {
      var d = new Date();
      var weekday = new Array(7);
      weekday[0] =  "Sunday";
      weekday[1] = "Monday";
      weekday[2] = "Tuesday";
      weekday[3] = "Wednesday";
      weekday[4] = "Thursday";
      weekday[5] = "Friday";
      weekday[6] = "Saturday";

      var day = weekday[dates.getDay()];

      return day;
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

.verticalTableHeader {
    text-align:center;
    white-space:nowrap;
    transform-origin:50% 50%;
    -webkit-transform: rotate(90deg);
    -moz-transform: rotate(90deg);
    -ms-transform: rotate(90deg);
    -o-transform: rotate(90deg);
    transform: rotate(90deg);
}
.verticalTableHeader:before {
    content:'';
    padding-top:80%;/* takes width as reference, + 10% for faking some extra padding */
    display:inline-block;
    vertical-align:middle;
}
#day_type {
	font-size: 0.95em;
	font-weight: 500;
}
.shaded {
  background-color: grey;
}
</style>
