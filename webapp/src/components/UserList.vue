<template lang="html">
<div id="user_list-wrapper">
  <ul style="list-style: none;">
    <li v-for="user in userList">
      <router-link :to="{ name: 'profile', params: { username: user.username }}">
        <sl-user-display
          :user="user"></sl-user-display>
      </router-link>
    </li>
    <hr>
  </ul>
</div>
</template>

<script>
import UserDisplay from 'components/UserDisplay'

export default {
  name: 'sl-user-list',

  mounted () {
    this.generateExamples()
    console.log(this.userList[1])
  },
  components: {
    'sl-user-display': UserDisplay
  },

  props: {
    userList: {
      type: Array,
      default () {
        return []
      }
    }
  },

  data () {
    return {
      user: { imgPath: '', name: '', username: '', usertype: '' }
    }
  },

  methods: {
    generateExamples () {
      var poop = {
        imgPath: 'http://1.media.collegehumor.cvcdn.com/14/45/7d51a082762115b56229fc6b741c1438.jpg',
        name: 'Poop',
        userType: 'user'
      }

      for (var i = 0; i < 5; i++) {
        this.userList.push(poop)
      }
    }
  }
}
</script>

<style lang="css">
#user_list-wrapper {

}
hr {
  margin: 5px;
}
</style>
