<template>
  <div class="entry-table">
    <sl-entry-dialog class="dialog"
      :open="dialogIsOpen"
      :entry="activeEntry"
      :default-activites="defaultActivites"
      v-on:add="addActivity()"
      v-on:remove="removeActivity($event)"
      v-on:close="closeDialog()"></sl-entry-dialog>
    <div class="table-container">
      <table class="description-table">
        <tr>
          <th>DATE:</th>
          <th>Day of the Week</th>
          <th>Type of Day
              <p id="day_type">Work, School, Off, Vacation (Vac)</p>
          </th>
        </tr>
        <tr v-for="date in dates">
          <td>{{ date.toLocaleDateString() }}</td>
          <td>{{ date.toLocaleDateString('eng', {weekday: 'long'}) }}</td>
          <td>
            <ui-select v-on:input="setDayType(date, $event)"
               :placeholder="'Select Day'"
               :options="dayOptions"
               :value="selectedDayOptions[date] || 'Off'"></ui-select>
          </td>
        </tr>
      </table>
      <table class="times-table">
        <tr>
          <th class="verticalTableHeader" v-for="time in times">{{time}}</th>
        </tr>
        <tr v-for="date in dates">
          <td v-for="time in times"
            :class="{ 'shaded': doesEntryContainSleep(date, time) }"
            v-on:click="openDialog(date, time)">{{displayInBox(date, time).toString()}}</td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import EntryDialog from './EntryDialog'
import UiSelect from './KeenUI/UiSelect'

export default {
  name: 'app',
  components: {
    'sl-entry-dialog': EntryDialog,
    'ui-select': UiSelect // <ui-button>
  },
  data () {
    return {
      entries: {},
      times: [],
      dates: [],
      dialogIsOpen: false,
      activeEntry: { date: '', time: '', activities: [] },
      defaultActivites: ['Sleep', 'Caffeine', 'Medicine', 'Alcohol', 'Exercise'],
      dayOptions: ['Work', 'School', 'Off', 'Vacation'],
      selectedDayOptions: {}
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
        // this.selectedDayOptions[this.dates[i]] = ''
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
        dates[j] = dates[j]
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
      selected = Object.keys(selected).map(key => selected[key].name)
      options = options.filter(i => selected.indexOf(i) < 0)
      return options[0]
    },
    setDayType (date, selectedName) {
      this.$set(this.selectedDayOptions, date, selectedName)
    }
  }
}
</script>

<style>
.table-container {
  width: 50px;
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
    text-align: center;
    white-space: nowrap;
    transform-origin: 50% 50%;
    -webkit-transform: rotate(90deg);
    -moz-transform: rotate(90deg);
    -ms-transform: rotate(90deg);
    -o-transform: rotate(90deg);
    transform: rotate(90deg);
}
.verticalTableHeader:before {
    content: '';
    padding-top:80%;
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
