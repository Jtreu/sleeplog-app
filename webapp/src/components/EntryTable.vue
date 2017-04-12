<!-- Hi -->
<template>
  <div class="entry-table">
    <sl-entry-dialog class="dialog"
      :open="dialogIsOpen"
      :entry="activeEntry"
      :default-activites="defaultActivites"
      v-on:add="addActivity()"
      v-on:remove="removeActivity($event)"
      v-on:update="updateProfileEntries()"
      v-on:close="closeDialog()"></sl-entry-dialog>
    <div class="outer-table-container">
      <div class="inner-table-container">
        <table>
          <tr>
            <th class="fixed-column">Date</th>
            <th class="fixed-column">Day of the Week</th>
            <th class="fixed-column">Type of Day
                <p id="day_type">Work, School, Off, Vacation (Vac)</p>
            </th class="fixed-column">
            <th v-for="time in times">{{time}}</th>
          </tr>
          <tr v-for="date in dates">
            <td class="fixed-column">{{ (new Date(date)).toLocaleDateString() }}</td>
            <td class="fixed-column">{{ (new Date(date)).toLocaleDateString('eng', {weekday: 'long'}) }}</td>
            <td class="fixed-column">
              <ui-select v-on:input="setDayType(date, $event)"
                 :placeholder="'Select Day'"
                 :options="dayOptions"
                 :value="entries[date].dayType || 'Off'"></ui-select>
            </td>
            <td v-for="time in times"
              :class="{ 'shaded': doesEntryContainSleep(date, time) }"
              v-on:click="openDialog(date, time)">{{displayInBox(date, time).toString()}}</td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import EntryDialog from './EntryDialog'
import UiSelect from './KeenUI/UiSelect'

export default {
  name: 'sl-entry-table',
  components: {
    'sl-entry-dialog': EntryDialog,
    'ui-select': UiSelect // <ui-button>
  },
  props: {
    userEntries: {
      type: Object
    },
    disableEdit: {
      type: Boolean,
      default () {
        return false
      }
    }
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
      dayType: ''
    }
  },
  mounted () {
    // Initialize times array
    this.generateTimes()
    // Initialize dates array
    this.generateDates()
    // Initialize entries object
    this.initEntries()

    if(isLoggedIn) {
      this.updateProfileEntries();
    }

    // If not your entries
    if (this.userEntries) {
      this.entries = this.userEntries
      return
    }

    // If your entries
    this.entries = Object.assign(this.entries, this.profileEntries)
  },
  watch: {
    userEntries () {
      this.entries = this.userEntries
    }
  },
  computed: {
    ...mapState({
      profileEntries: state => state.profile.entries
    })
  },
  methods: {
    doesEntryContainSleep (date, time) {
      if (!this.entries || !this.entries[date].times[time] || this.entries[date].times[time].activities.length < 1) {
        return false
      }
      let activities = this.entries[date].times[time].activities
      activities = Object.keys(activities).map(key => activities[key])
      for (let i = 0; i < activities.length; i++) {
        if (activities[i].name === 'Sleep' && activities[i].isDone) {
          return true
        }
      }
      return false
    },
    displayInBox (date, time) {
      let activities = this.entries[date].times[time].activities
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
        entries[`${this.dates[i]}`].dayType = ''
        entries[`${this.dates[i]}`].times = {}
        // this.selectedDayOptions[this.dates[i]] = ''
        for (var j = 0; j < this.times.length; j++) {
          entries[`${this.dates[i]}`].times[`${this.times[j]}`] = {
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
      for (var i = 1; i < 5; i++) {
        // get the next day after the previous
        var thisDate = new Date(dates[i - 1].getTime() - (24 * 60 * 60 * 1000))
        // thisDate = thisDate.toDateString()
        dates.push(thisDate)
      }
      dates = dates.map(date => (new Date(date)).toDateString())
      this.dates = dates
    },
    openDialog (date, time) {
      if (this.disableEdit) {
        return
      }
      this.activeEntry = this.entries[date].times[time]
      this.dialogIsOpen = true
    },
    closeDialog () {
      if (this.disableEdit) {
        return
      }
      this.dialogIsOpen = false
    },
    addActivity () {
      if (this.disableEdit) {
        return
      }

      if (!this.activeEntry.date || !this.activeEntry.time) {
        return
      }
      // Don't add more activities than available
      if (this.entries[this.activeEntry.date].times[this.activeEntry.time].activities.length >= this.defaultActivites.length) {
        return
      }
      let availableName = this.getAvailableActivityName(this.entries[this.activeEntry.date].times[this.activeEntry.time].activities)
      this.entries[this.activeEntry.date].times[this.activeEntry.time].activities.push({
        name: availableName,
        isDone: false
      })

      this.updateProfileEntries()
    },
    removeActivity (index) {
      if (this.disableEdit) {
        return
      }
      if (index > -1) {
        this.entries[this.activeEntry.date].times[this.activeEntry.time].activities.splice(index, 1)
      }

      this.updateProfileEntries()
    },
    getAvailableActivityName (selected) {
      let options = this.defaultActivites
      selected = JSON.parse(JSON.stringify(selected))
      selected = Object.keys(selected).map(key => selected[key].name)
      options = options.filter(i => selected.indexOf(i) < 0)
      return options[0]
    },
    setDayType (date, selectedName) {
      if (this.disableEdit) {
        // Make the text static here instead of having a return.
        return
      }
      this.entries[date].dayType = selectedName
      this.updateProfileEntries()
    },
    updateProfileEntries () {
      if (this.disableEdit) {
        return
      }
      this.$store.dispatch('updateEntries', { entries: this.entries })
        .then(() => {
          // console.log('Successfully updated profile')
        })
        .catch(() => {
          // console.log('Failed to update profile entries: ', err)
        })
    }
  }
}
</script>

<style lang="scss" scoped>
.outer-table-container {
  position: relative;
}

.inner-table-container {
  overflow-x: scroll;
  overflow-y: visible;
  width: 100%;
}

.fixed-column {
  position: relative;
  width: 100px;
}

.shaded {
  background-color: grey;
}

table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
  background-color: rgb(207, 231, 209);
  border: 2px solid #ccc;
}

td, th {
  vertical-align: top;
  border-top: 2px solid #ccc;
  border-right: 2px solid #ccc;
  padding: 10px;
  width: 100px;
}
#day_type {
	font-size: 0.95em;
	font-weight: 500;
}
</style>
