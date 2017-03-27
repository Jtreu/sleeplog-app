<template lang="html">
  <div id="modal"
    v-on:click="close($event)"
    class="entry-dialog modal"
    :style="{display: open ? 'block' : 'none'}">
    <div class="modal-content-container">
      <div class="modal-content">
        <div class="activityList-container">
          <div class="activityItem" v-for="(activity, index) in entry.activities">
            <div class="name-container">
              <ui-select
                v-model="activity.name"
                v-on:select="setActivity()"
                :placeholder="'Select an activity'"
                :options="activityOptions[index]"></ui-select>
            </div>
            <div class="input-container">
              <ui-checkbox v-on:input="updateEntries()"
                v-model="activity.isDone"></ui-checkbox>
            </div>
            <div class="remove-btn-container">
              <ui-button v-on:click="removeActivity(index)">Remove</ui-button>
            </div>
          </div>
        </div>
        <div class="add-btn-container">
          <ui-button v-on:click="addActivity()">Add Activity</ui-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import UiSelect from './KeenUI/UiSelect'
import UiCheckbox from './KeenUI/UiCheckbox'
import UiButton from './KeenUI/UiButton'

export default {
  name: 'sl-entry-dialog', // The html tag in the markup. <sl-entry-dialog>
  components: {
    'ui-select': UiSelect, // <ui-select>
    'ui-checkbox': UiCheckbox, // <ui-checkbox>
    'ui-button': UiButton // <ui-button>
  },
  props: { // props is data that is stored locally in this component but comes from the parent component..
    open: {
      type: Boolean,
      default () {
        return false
      }
    },
    entry: {
      type: Object,
      default () {
        return {
          date: '',
          time: '',
          activities: [{
            name: 'Sleep',
            isDone: false
          }],
          typeday: []
        }
      }
    },
    defaultActivites: {
      type: Array,
      default () {
        return []
      }
    }
  },
  data () {
    return { }
  },
  computed: {
    activityOptions () {
      let activities = JSON.parse(JSON.stringify(this.entry.activities))
      activities = Object.keys(activities).map(key => activities[key].name)
      return this.calculateAvailableOptions(activities)
    }
  },
  methods: {
    setActivity () {
      let activities = JSON.parse(JSON.stringify(this.entry.activities))
      activities = Object.keys(activities).map(key => activities[key].name)
      let duplicate = this.checkForDuplicateOptions(activities)
      if (duplicate.exists) {
        this.removeActivity(duplicate.index)
      }
    },
    checkForDuplicateOptions (activities) {
      let options = {}
      let duplicate = { exists: false, index: 0 }
      for (let i = 0; i < activities.length; i++) {
        if (!options[activities[i]]) {
          options[activities[i]] = true
        } else {
          duplicate.exists = true
          duplicate.index = i
        }
      }
      return duplicate
    },
    calculateAvailableOptions (activities) {
      let options = []
      let selected = []
      if (!selected) {
        return
      }
      for (let i = 0; i < activities.length; i++) {
        options[i] = JSON.parse(JSON.stringify(this.defaultActivites)) || []
        selected[i] = activities.slice(0, i) || []
        options[i] = options[i].filter(arrIndex => selected[i].indexOf(arrIndex) < 0)
      }
      return options
    },
    close (event) {
      if (event.target.id === 'modal') {
        this.$emit('close')
      }
    },
    updateEntries () {
      this.$emit('update')
    },
    addActivity () {
      this.$emit('add')
    },
    removeActivity (index) {
      this.$emit('remove', index)
    }
  }
}
</script>
<style lang="scss" scoped>
.entry-dialog {
  text-align: left;
}

.activityItem {
  display: block;
  position: relative;
  .name-container, .input-container, .remove-btn-container {
    vertical-align: top;
    display: inline-block;
  }
  .name-container, .input-container {
    width: 200px;
  }
  .name-container {
    margin-right: 30px;
  }
}
.modal {
    display: none; /* Hidden by default */
    position: fixed; /* Stay in place */
    z-index: 1; /* Sit on top */
    left: 0;
    top: 0;
    width: 100%; /* Full width */
    height: 100%; /* Full height */
    overflow: auto; /* Enable scroll if needed */
    background-color: rgb(0,0,0); /* Fallback color */
    background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
}
.modal-content-container {
    background-color: #fefefe;
    margin: 15% auto; /* 15% from the top and centered */
    padding: 20px;
    border: 1px solid #888;
    width: 80%; /* Could be more or less, depending on screen size */
}
</style>
