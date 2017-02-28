<template>
  <!-- App -->
  <div id="app">

    <!--
    Use the sl-entry-dialog component.

    :open means we are passing data to the property (aka prop) called 'open' that's defined in this component.

    Since dialogIsOpen is just some data.. we can pass it. Vuejs handles all the propogation stuff when dialogIsOpen
    changes. So for instance, when dialogIsOpen changes to true or false, it will be updated in the sl-entry-dialog component.

    'entry' is yet another prop. If you look in the EntryDialog.vue component you'll see that it is expected to be of type Object.
    Again. Objects in javascript are just key:value pairs.

    For example:

    {
      'isCool': true,
      'isOpen': true,
      'initWorld': function () { console.log('Hello World') },
      'anotherObject': {
        'anotherKey': 'anotherValue',
        'anArray': [
          'stuff',
          'stuff2',
          'stuff3',
          5,
          6
          ]
      }
    }

    keys must be strings. However, values can be whatever.

    v-on:add allows us to listen to events that the component will propogate up to the parent (this is the parent component).
    In this case, we are listening to the 'add' event. In the component, when this.$emit('add') occurs, then we can know about it by using
    v-on:add="doSomething()".

    v-on:remove is a way for us to handle when the child component emits an even called 'remove'. Keep in mind these names are arbitrary. You
    could call it 'poop' or whatever. So if you wanted to listen to the event 'poop' you would do

    v-on:poop="" // in the string you call a function you want to do something in response to the poop. For example

    v-on:poop="console.log('mmmm mmm Poopy')"

    You also notice the $event that we pass. If the child component passes an argument with their event emission:

    this.$emit('poop', someArgument, someArgument2)

    Then we can use $event to get those arguments. In this case we are getting the index that is passed up and calling
    the removeActivity() function that we've defined in our methods object.

    v-on:close allows us to handle the emission of an event called 'close'
    -->
    <sl-entry-dialog class="dialog"
      :open="dialogIsOpen"
      :entry="activeEntry"
      v-on:add="addActivity()"
      v-on:remove="removeActivity($event)"
      v-on:close="closeDialog()"></sl-entry-dialog>

    <!--
    A Table where we display our stuff
    -->
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
/*
Delete these comments later after you've understood how things are connecting to one another
*/
/*
Import other components you want to use
*/
import EntryDialog from './components/EntryDialog'

export default {
  /*
  Name the component (required)
  */
  name: 'app',
  /*
  If you use components in other files, reference them here
  */
  components: {
    /*
    In our markup, we'll reference EntryDialog as 'sl-entry-dialog', but you can call it whatever.. just stay consistent.

    Staying consistent means if you call this 'sl-whatever-component' then in your markup you will need to write

    <sl-whatever-component></sl-whatever-component>
    */
    'sl-entry-dialog': EntryDialog
  },
  data () {
    return {
      /*
      entries is an object that contains the entries for each day and hour. The object is structured as follows:

      entries: {
        'todaysDate': {
          '1am': {
            date: 'todaysDate',
            time: '1am',
            activities: [{
              name: 'Sleep',
              duration: 60 // measured in minutes
            }, {
              name: 'Caffeine',
              duration: 60 // You said you don't want this so then you will have to think about how to restructure this object to fit your needs
            }]
          },
          '2am': { ... },
          '3am': { ... }, ...
        },
        'tomorrowsDate': { ... },
        'dateAfterThat': { ... }, ...
      }
      */
      entries: {},

      /*
      times is an array that holds the times that we'll display with v-for in our markup
      */
      times: [],

      /*
      dates is an array that holds the dates that we'll display with v-for in our markup
      */
      dates: [],

      /*
      dialogIsOpen is a boolean that we toggle to display the dialog or not
      */
      dialogIsOpen: false,

      /*
      activeEntry is the current entry that's being modified in the dialog / modal .. this will be set
      when we open the dialog.. (see the openDialog method below)
      */
      activeEntry: { date: '', time: '', activities: [] }
    }
  },
  /*
  This is a lifecycle hook for vuejs components. There are other lifecycle hooks.

  Within the mounted() lifecycle hook we can do some setup. In this case, we're setting up
  the array of times and dates that we will use in the the markup with v-for.
  */
  mounted () {
    // Initialize times array
    this.generateTimes()

    // Initialize dates array
    this.generateDates()

    // Initialize entries object
    this.initEntries()
  },
  /*
  methods are functions that our component can use to do things
  */
  methods: {
    /*
    initEntries initializes the entries object that we use to store the data of each entry in the table.
    */
    initEntries () {
      /*
      Initialize entries object
      */
      var entries = {}

      /*
      Have a default activity for each entry so each entry's dialog will display at least 1 activity to start with.

      Activities are displayed in EntryDialog.vue component.
      */
      var defaultActivity = {
        name: 'Sleep',
        duration: 30
      }

      /*
      For each date
        For each time of the day
          create an entry for that day and time and set date, time and activities for that day/time entry
      */
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

      /*
      Set out component's entries data to the entries that we've created here
      */
      this.entries = entries
    },
    /*
    generateTimes intitializes the times array that will be used to display the times columns
    */
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
    /*
    generateDates initializes the dates array that will be used to display the dates rows
    */
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
    /*
    openDialog opens the dialog component by setting dialogIsOpen to true and setting the activeEntry so
    the dialog component knows what data to display.

    This allows us to use a single dialog component for all entries rather than creating a dialog component
    for each entry.
    */
    openDialog (date, time) {
      this.activeEntry = this.entries[date][time]
      this.dialogIsOpen = true
    },
    /*
    closeDialog closes the dialog by setting dialogIsOpen to false. Doing this will make Vue propogate that change
    to they dialog component as a property (aka prop). That prop is used to set the styling on the component
    to display: none; to hide it or display: block; to show it normally.
    */
    closeDialog () {
      this.dialogIsOpen = false
    },
    /*
    addActivity adds an activity to the activeEntry. So if you look in the EntryDialog component you'll see there is a
    button for adding entries. That button emits a click event that then calls a function which itself calls this.$emit('add')

    We can call 'add' whatever we'd like, but we have to be consistent and handle that event on the component by using
    :whatever-we-call-it="addActivity()"

    addActivity makes sure that the activeEntry has a date and time before continuing. If there is a date and time, it
    will add an activity by the name of 'Sleep' and has a duration of 20.
    */
    addActivity () {
      if (!this.activeEntry.date || !this.activeEntry.time) {
        return
      }
      this.entries[this.activeEntry.date][this.activeEntry.time].activities.push({
        name: 'Sleep',
        duration: 20
      })
    },
    /*
    removeActivity removes an activity from the activeEntry for a given index. Since activities is an array we need an
    index to remove something from it.
    */
    removeActivity (index) {
      if (index > -1) {
        this.entries[this.activeEntry.date][this.activeEntry.time].activities.splice(index, 1)
      }
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
