<template lang="html">
  <div class="profile-header">
    <div class="banner-container">
      <div class="banner-img-container">
        <img class="banner-img" :src="banner" :alt="label.BANNER">
      </div>
    </div>
    <div class="user-container">
      <div class="user">
        <div class="user-img-container">
          <img class="user-img" :src="icon" :alt="label.PROFILE">
        </div>
        <div class="user-name-container">
          <div class="user-alias">{{alias}}</div>
          <div class="user-username">@{{username}}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'sl-profile-header',
  props: {
    username: {
      type: String,
      default () {
        return 'johnwick'
      }
    },
    alias: {
      type: String,
      default () {
        return 'John Wick'
      }
    },
    icon: {
      type: String,
      default () {
        return ''
      }
    },
    banner: {
      type: String,
      default () {
        return ''
      }
    }
  },
  data () {
    return {
      label: {
        PROFILE: 'PROFILE',
        BANNER: 'BANNER'
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';

$profile-image-size: 55px;

.profile-header {

}

.banner-container {
  position: relative;
  display: block;
  border: none;
  overflow: hidden;
  height: auto;

  .banner-img-container {
    position: relative;
    left: 0;
    right: 0;
    text-align: center;
  }

  .banner-img {
    height: auto;
    width: 100%;
    display: block;
    text-align: center;
  }
}

.user-container {
  background-color: #f6f6f6;
  position: relative;
  border-bottom: none;
  padding: 16px 8px 0 16px;
  overflow: visible;

  .user {
    overflow: hidden;
    white-space: nowrap;
    word-wrap: normal;
    text-overflow: ellipsis;
  }

  .user-img-container {
    position: absolute;
    top: -27px;
    width: $profile-image-size;
    height: $profile-image-size;
    overflow: hidden;
    text-align: center;
    white-space: nowrap;
    float: left;
    border-radius: $profile-image-size;
  }

  .user-img {
    border: 0;
    height: $profile-image-size;
  }

  .user-name-container {
    padding-top: 12px;
    line-height: 23px;
    margin-bottom: 2px;
  }

  .user-alias {
    font-size: 18px;
    color: $text-color;
  }

  .user-username {
    color: #444;
  }
}
</style>
