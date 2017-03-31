<template lang="html">
  <div class="header">
    <link href="https://fonts.googleapis.com/css?family=Oxygen" rel="stylesheet">
    <router-link :to="{ name: 'home' }">
      <div class="logo-container" >
        <img src="/static/logo.png" class="logo"><span>{{ label.APP_NAME }}</span>
      </div>
    </router-link>
    <form class="searchForm" v-on:submit="submitForm($event)">
      <input id="searchBar"
        type="text"
        v-model="form.query"
        :placeholder="label.FIND_USERS">
      <input id="searchBtn" type="submit" :value="label.SEARCH">
    </form>
    <nav class="header-navs">
      <a v-show="isLoggedIn" v-on:click="logOut()" href="/">LOGOUT</a>
    </nav>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex'

export default {
  name: 'sl-header',
  data () {
    return {
      form: {
        query: ''
      },
      label: {
        APP_NAME: 'Sleeplog App',
        FIND_USERS: 'Find Users',
        SEARCH: 'Search'
      }
    }
  },
  mounted () {
    let query = this.$route.query.q
    this.form.query = query

    if (query) {
      this.search(query)
    }
  },
  computed: {
    ...mapState({
      isLoggedIn: state => state.profile.isLoggedIn
    }),
    ...mapGetters([
      'profile'
    ])
  },
  methods: {
    search (query) {
      return this.$store.dispatch('searchUsers', query)
        .then(() => {
          this.$router.push({ name: 'results', query: { q: query } })
        })
    },
    submitForm (event) {
      event.preventDefault()

      // Search
      let query = this.form.query
      if (query) {
        this.search(query)
      }
    },
    clearQuery () {
      this.form.query = ''
    },
    logOut () {
      /* Add functionality to log user out */
      this.$store.dispatch('logout')
      this.$store.isLoggedIn = false
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';
@import '../assets/styles/input';

$input-height: 40px;

.header {
  width: 100%;
  background-color: $app-theme-color;
}

.searchForm {
  display: inline-block;
  font-size: 0;
  width: 70%;
}

.logo-container {
  margin-left: 20px;
  float:left;
  font-weight: bold;
}
.logo {
  height: $input-height;
}
.header-navs {
  display: inline-block;
  float:right;
  margin-top: 18px;
  font-size: 16px;
  font-family: 'Oxygen', sans-serif;
}
.header-navs a {
  width: 100%;
  background-color: #998DA0;
  color: #fff;
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
#searchBar {
  width: 50%;
  border: 0;
  padding: 10px;
  border-radius: 2px;
  font-size: 18px;
  height: $input-height;
}

#searchBtn {
  height: $input-height;
  width: initial;
}
</style>
