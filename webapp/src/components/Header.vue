<template lang="html">
  <div class="header">
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
  </div>
</template>

<script>
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
  font-size: 0;
}

.logo-container {
  margin-left: 20px;
  float: left;
  font-weight: bold;
}
.logo {
  height: $input-height;
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
