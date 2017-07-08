// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import store from './store'
import router from './router'
import { sync } from 'vuex-router-sync'

import Vuelidate from 'vuelidate'
import BootstrapVue from 'bootstrap-vue'

// sync the router with the vuex store.
// this registers `store.state.route`
sync(store, router)

Vue.use(Vuelidate)
Vue.use(BootstrapVue)

/* eslint-disable no-new */
const app = new Vue({
  router,
  store,
  ...App
})

// actually mount to DOM
app.$mount('#app')

// export { app, router, store }
