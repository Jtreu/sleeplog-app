import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Home from '../views/Home'
import Intro from '../views/Intro'

import Search from '../views/Search'
import Profile from '../views/Profile'
import Settings from '../views/Settings'

import EmailConfirmation from '../views/EmailConfirmation'
import PasswordReset from '../views/PasswordReset'

const router = new Router({
  mode: 'history',
  scrollBehavior: () => ({ y: 0 }),
  routes: [
    { path: '/', name: 'home', component: Home },

    { path: '/results', name: 'results', component: Search },

    { path: '/settings', name: 'settings', component: Settings },

    { path: '/account/email-confirmation/:code', name: 'email-confirmation', component: EmailConfirmation },
    { path: '/account/password-reset/:code', name: 'password-reset', component: PasswordReset },

    { path: '/:username', name: 'profile', component: Profile }
  ]
})

export default router
