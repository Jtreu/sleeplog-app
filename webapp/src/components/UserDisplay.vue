<template lang="html">
  <div id="user_wrapper">
    <link href="https://fonts.googleapis.com/css?family=Oxygen" rel="stylesheet">
    <div class="img_container">
      <img :src="this.user.imgPath" alt="User Picture" style="width:304px;height:228px;">
    </div>

    <div class="info_container">
      <ul style="list-style: none;">
        <li><strong>Name: </strong>{{this.user.name}}</li>
        <li><strong>Occupation: </strong>{{this.user.name}}</li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'sl-user-display',

  props: {
    user: {
      type: Object,
      default () {
        return {
          imgPath: '',
          name: '',
          usertype: ''
        }
      }
    }
  },
  data () {
    return { }
  },
  methods: {

  }
}
</script>

<style lang="scss" scoped>
#app {
  width: 100%;
}
#user_wrapper {
  width: 100%;
  background-color: #4be5cb;
  box-shadow: 10px 10px 5px #888888;
}
div {
  margin-right: 5px;
}
.img_container {
  display:inline-block;
  // Box shadow won't work
  box-shadow: 10px 20px 20px #888888;
  // border: 1px solid black;
}
.info_container {
  display:inline-block;
  font-family: 'Oxygen', sans-serif;
  vertical-align: top;
  width: 400px;
  height:100%;
  font-size: 1.10em;
}
ul li {
  margin: 20px;
}
</style>
