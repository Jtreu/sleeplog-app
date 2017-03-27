<template lang="html">
  <div class="profile-nav">
    <div class="nav-container">
      <h1 class="nav-item-container"
          v-on:click="select('home')"
          v-bind:class="{ 'nav-item-selected': isSelected('home')}">
        <div class="nav-item">
          <div class="nav-item-text">{{label.HOME}}</div>
        </div>
      </h1>
    </div>
  </div>
</template>

<script>
export default {
  name: 'sl-profile-nav',
  props: {
    selected: {
      type: String,
      default () {
        return 'home'
      }
    }
  },
  data () {
    return {
      label: {
        GAMES: 'games',
        HOME: 'about'
      }
    }
  },
  methods: {
    select (option) {
      this.$emit('select', option)
    },
    isSelected (option) {
      if (option === this.selected) {
        return true
      }

      return false
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';

$nav-height: 44px;

.profile-nav {
  position: relative;
  display: table;
  table-layout: fixed;
  width: 100%;
  height: $nav-height;
  background-color: #f6f6f6;
  border-bottom: 1px solid #e9e9e9;
  border-image: none;
}

.nav-container {
  margin-right: 40px;
}

.nav-item-container {
  width: 33.333333333333336%;
  background-image: none;
  display: inline-table;
  font-size: 14px;
  line-height: 20px;
  text-align: center;
  padding: 7px 0;
  vertical-align: top;
  margin: 0;
  color: $text-color;
  font-weight: normal;
}

.nav-item {
  border-right: none;
  display: table-cell;
  vertical-align: middle;
  padding: 0;
  height: 34px;
  color: $dark-gray;
  box-sizing: border-box;

  .nav-item-text {
    text-transform: uppercase;
    word-break: break-word;
    overflow: hidden;
    -webkit-line-clamp: 2;
    text-overflow: ellipsis;
    line-height: normal;
  }
}

.nav-item-selected {
  border-bottom: 4px solid $app-theme-color;
}
</style>
