<template lang="html">
  <!--
  v-on:click allows us to do something when the click event is emitted. In this case, when click is fired, we will
  call our function called close() that's defined locally. We pass the $event because our close function wants to make sure that
  we are clicking outside of the modal content.

  :style is a way for us to pass data down to the property (or prop) called style. In this case it's basically allowing us
  to define styles that is connected to our javascript. So in this code we are saying, if open is true then set display: block;
  if open in false set display: 'none'..

  display is a css property of course.
  -->
  <div id="modal"
    v-on:click="close($event)"
    class="entry-dialog modal"
    :style="{display: open ? 'block' : 'none'}">
    <div class="modal-content-container">
      <div class="modal-content">
        <div class="activityList-container">
          <!--
          Loop through the activities of this entry and display them here. Notice how we're also interested in the index of the item
          in the array. So we have "(activity, index) in entry.activities". Instead of "activity in entry.activities"
          -->
          <div class="activityItem" v-for="(activity, index) in entry.activities">
            <div class="name-container">
              <!--
              This is a custom component from the library called KeenUI.

              Documentation on how to use these components: https://josephuspaye.github.io/Keen-UI/#/ui-alert
              -->
              <ui-select
                v-model="activity.name"
                :placeholder="'Select and activity'"
                :options="activityOptions"></ui-select>
            </div>
            <div class="duration-container">
              <!--
              Another KeenUI component. It comes with a bunch of styling that we can use for free. You can try creating
              your own slider but this was faster to use someone else's
              -->
              <ui-slider
                v-model="activity.duration"
                :step="15"
                :snap-to-steps="true"
                :show-marker="true"
                :marker-value="Math.floor((activity.duration / 100) * 60)"></ui-slider>
              <span>{{Math.floor((activity.duration / 100) * 60)}} minutes</span>
            </div>
            <div class="remove-btn-container">
              <!--
              Another KeenUI component. You can use <button></button> which is just an html element but it is styled
              ugly. Using other people's code like this allows you to have styling out of the box with little effort.
              -->
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
// Import the 3rd party components we'll use in our markup
import UiSelect from './KeenUI/UiSelect'
import UiSlider from './KeenUI/UiSlider'
import UiButton from './KeenUI/UiButton'

// Export this component so other components can use it
export default {
  // Name the component (required)
  name: 'sl-entry-dialog',
  // Declare how you'd like to name the components in the markup
  components: {
    'ui-select': UiSelect,
    'ui-slider': UiSlider,
    'ui-button': UiButton
  },
  // Define the properties that any parent component can pass down this component (the child)
  props: {
    // open is a prop that is expected to be of type Boolean and have a default value of false (if the parent doesn't pass anything down)
    open: {
      type: Boolean,
      default () {
        return false
      }
    },
    // entry is a prop that is expected to be an Object and have a default (if the parent doesn't pass anything)
    entry: {
      type: Object,
      default () {
        return {
          date: '',
          time: '',
          activities: [{
            name: 'Sleep',
            duration: 30
          }]
        }
      }
    }
  },
  /*
  data() returns an object of the local state of this component. State is just data. it's just a representation of the
  information that this component cares about. For instance activityOptions is part of this component's state.

  You can add more data as you see fit. Normally it is data like strings, objects, arrays etc.. not functions.. those belong
  in methods object
  */
  data () {
    return {
      // activityOptions is used by out markup to generate the names of the select options
      activityOptions: ['Sleep', 'Caffeine', 'Medicine']
    }
  },
  /*
  methods allows us to write functions that we can use in our markup or even in lifecycle hooks like mounted() if we want to
  do some setup or stuff that is specific to this component.
  */
  methods: {
    /*
    close is a function that takes the event and determines if the id of the target is 'modal'. if the id is modal
    then we know the user has clicked outside of the modal-content div. They have clicked in the dark area. So
    we interpret that as them wanting to leave.

    We then emit an event that the parent can handle if they want. The event we emit is called 'close'. If the parent
    component wants to know when we emit this event, they can use

    v-on:close="doSomething()"
    */
    close (event) {
      if (event.target.id === 'modal') {
        this.$emit('close')
      }
    },
    /*
    addActivity is a function that we use to tell the parent that we want to add an activity to this current entry
    */
    addActivity () {
      this.$emit('add')
    },
    /*
    removeActivity is a function that we use to tell the parent that we want to remove an activity with a given
    index from the current entry.
    */
    removeActivity (index) {
      this.$emit('remove', index)
    }
  }
}
</script>

<!-- SASS/CSS code -->
<style lang="scss" scoped>

.entry-dialog {
  text-align: left;
}

.activityItem {
  display: block;
  position: relative;

  .name-container, .duration-container, .remove-btn-container {
    vertical-align: top;
    display: inline-block;
  }

  .name-container, .duration-container {
    width: 200px;
  }

  .name-container {
    margin-right: 30px;
  }
}

/* The Modal (background) */
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

/* Modal Content/Box */
.modal-content-container {
    background-color: #fefefe;
    margin: 15% auto; /* 15% from the top and centered */
    padding: 20px;
    border: 1px solid #888;
    width: 80%; /* Could be more or less, depending on screen size */
}
</style>
