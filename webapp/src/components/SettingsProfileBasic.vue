<template lang="html">
  <div class="settings-profile-basic">
    <div v-if="showFormSpinner"
      class="spinner">
      <sl-spinner class="spinner-icon"></sl-spinner>
    </div>
    <div v-else-if="formUpdateError"
      class="error saveMessage">
      {{ formUpdateError }}
    </div>
    <div v-else-if="formUpdateSuccess"
      class="success saveMessage">
      {{ formUpdateSuccess }}
    </div>
    <form v-on:submit="submitForm($event)">
      <div>
        <label for="fullname">
          {{ label.FULL_NAME }}
          <span v-if="$v.form.fullname.$dirty && form.fullname !== fullname"
            class="change">{{ label.REQUIRES_SAVE }}</span>
        </label>
        <input id="fullname"
          v-model.trim="form.fullname"
          @input="$v.form.fullname.$touch()"
          type="text">
        <div v-if="form.fullname !== fullname && $v.form.fullname.$invalid" class="error">
          <span v-if="!$v.form.fullname.required">Your Full Name is required.</span>
          <span v-else-if="!$v.form.fullname.maxLength">Name must be shorter than 20 characters.</span>
        </div>
      </div>
      <div>
        <label for="username">
          {{ label.USERNAME }}
          <span v-if="$v.form.username.$dirty && form.username !== username"
            class="change">{{ label.REQUIRES_SAVE }}</span>
        </label>
        <input id="username"
          v-model.trim="form.username"
          @input="$v.form.username.$touch()"
          type="text">
        <div v-if="form.username !== username && !$v.form.username.$pending && $v.form.username.$invalid" class="error">
          <span v-if="!$v.form.username.required">Username is required.</span>
          <span v-else-if="!$v.form.username.maxLength">Username must be shorter than 15 characters.</span>
          <span v-else-if="!$v.form.username.nonSpecial">No special characters.</span>
          <span v-else-if="!$v.form.username.isUnique">This Username taken.</span>
        </div>
        <div v-else-if="$v.form.username.$pending || !$v.form.username.$invalid"  class="success">
          <span>This Username is available.</span>
        </div>
      </div>
      <div>
        <label for="email">
          {{ label.EMAIL }}
          <span v-if="$v.form.email.$dirty && form.email !== email"
            class="change">{{ label.REQUIRES_SAVE }}</span>
        </label>
        <input id="email"
          v-model.trim="form.email"
          @input="$v.form.email.$touch()"
          type="text">
        <div v-if="form.email !== email && !$v.form.email.$pending && $v.form.email.$invalid" class="error">
          <span v-if="!$v.form.email.required">Your Email is required.</span>
          <span v-else-if="!$v.form.email.email">Email must be valid.</span>
          <span v-else-if="!$v.form.email.isUnique">This email is unavailable.</span>
        </div>
      </div>
      <div class="icon-container">
        <label for="icon" class="input-label">
          {{ label.ICON }}
          <span v-if="fileInputDirty('icon')"
            class="change">{{ label.REQUIRES_SAVE }}</span>
        </label>
        <div class="user-img-container">
          <img class="user-img" :src="form.icon">
        </div>
        <input id="icon"
          ref="icon"
          type="file"
          v-on:change="handleImageSelect($event, 'icon')">
        <div v-if="!validIconFile"
          class="error">
          Input must be an image.
        </div>
      </div>
      <div class="banner-container">
        <label for="banner" class="input-label">
          {{ label.BANNER }}
          <span v-if="fileInputDirty('banner')"
            class="change">{{ label.REQUIRES_SAVE }}</span>
        </label>
        <div class="banner-container">
          <div class="banner-img-container">
            <div class="banner-img"
              :style="{ 'background-image': bannerURL }"></div>
          </div>
        </div>
        <input id="banner"
          ref="banner"
          type="file"
          v-on:change="handleImageSelect($event, 'banner')">
        <div v-if="!validBannerFile"
          class="error">
          Input must be an image.
        </div>
      </div>
      <input type="submit" :value="label.SAVE">
    </form>
  </div>
</template>

<script>
import { mapState } from 'vuex'

import { required, email, maxLength } from 'vuelidate/lib/validators'

import Spinner from '../components/Spinner'

export default {
  name: 'sl-settings-profile-basic',
  components: {
    'sl-spinner': Spinner
  },
  data () {
    return {
      validIconFile: true,
      validBannerFile: true,
      form: {
        fullname: '',
        username: '',
        email: '',
        icon: 'https://yt3.ggpht.com/-NcHEcMW_jJc/AAAAAAAAAAI/AAAAAAAAAAA/0kW1e4vlxpY/s100-c-k-no-mo-rj-c0xffffff/photo.jpg',
        banner: 'https://yt3.ggpht.com/oDJaL_HFqg0HKL9yKjk0bwVS7FLG4JSOX6d-PK_vHyYoL6NcK7e-m7dAC3bJpnNiTORftYg1jw=w2120-fcrop64=1,00005a57ffffa5a8-nd-c0xffffffff-rj-k-no'
      },
      label: {
        PROFILE: 'Profile',
        FULL_NAME: 'Full Name',
        USERNAME: 'Username',
        EMAIL: 'Email',
        PASSWORD: 'Password',
        ICON: 'Photo',
        BANNER: 'Banner Image',
        SAVE: 'Save Changes',
        REQUIRES_SAVE: '- requires save',
        OLD_PASSWORD: 'Old Password',
        NEW_PASSWORD: 'New Password',
        NEW_PASSWORD_REPEAT: 'Repeat New Password'
      },
      showFormSpinner: false,
      formUpdateError: '',
      formUpdateSuccess: ''
    }
  },
  validations: {
    form: {
      fullname: {
        required,
        maxLength: maxLength(20)
      },
      username: {
        required,
        maxLength: maxLength(15),
        nonSpecial (value) {
          return /^[a-zA-Z0-9_]*$/.test(value)
        },
        isUnique (value) {
          if (value === '') return true
          return this.$store.dispatch('checkUsernameAvailability', value) || true
        }
      },
      email: {
        required,
        email,
        isUnique (value) {
          if (value === '') return true
          return this.$store.dispatch('checkEmailAvailability', value) || true
        }
      }
    }
  },
  mounted () {
    this.initForm()
  },
  computed: {
    bannerURL () {
      return `url('${this.form.banner}')`
    },
    ...mapState({
      fullname: state => state.profile.alias,
      username: state => state.profile.username,
      email: state => state.profile.email,
      icon: state => state.profile.media.icon.source,
      banner: state => state.profile.media.banner.source
    })
  },
  methods: {
    fileInputDirty (name) {
      if (!this.$refs[name]) {
        return false
      }

      return this.$refs[name].files.length > 0
    },
    handleImageSelect (event, name) {
      event.preventDefault()

      if (name === 'icon') {
        this.validIconFile = this.validImageFile(name)
      } else if (name === 'banner') {
        this.validBannerFile = this.validImageFile(name)
      }

      let files = event.target.files

      for (let i = 0; i < files.length; i++) {
        let f = files[i]

        if (!f.type.match('image.*')) {
          continue
        }

        let reader = new window.FileReader()

        reader.onload = ((file) => {
          return (e) => {
            let image = e.target.result
            if (name === 'icon') {
              this.form.icon = image
            } else if (name === 'banner') {
              this.form.banner = image
            }
          }
        })(f)

        reader.readAsDataURL(f)
      }
    },
    saveSettings () {
      let updatedFullname = this.form.fullname !== this.fullname
      let updatedUsername = this.form.username !== this.username
      let updatedEmail = this.form.email !== this.email
      let updatedIcon = this.$refs.icon.files.length > 0
      let updatedBanner = this.$refs.banner.files.length > 0

      let promises = []

      if (updatedFullname) {
        promises.push(
          this.$store.dispatch('updateProfileAlias', this.form.fullname)
        )
      }

      if (updatedUsername) {
        promises.push(
          this.$store.dispatch('updateProfileUsername', this.form.username)
        )
      }

      if (updatedEmail) {
        promises.push(
          this.$store.dispatch('updateProfileEmail', this.form.email)
        )
      }

      if (updatedIcon) {
        let formData = new window.FormData()
        let imageData = this.$refs.icon.files[0]
        formData.append('icon', imageData)

        promises.push(
          this.$store.dispatch('updateProfileIcon', formData)
            .then(() => {
              this.clearFile('icon')
            })
        )
      }

      if (updatedBanner) {
        let formData = new window.FormData()
        let imageData = this.$refs.banner.files[0]
        formData.append('banner', imageData)

        promises.push(
          this.$store.dispatch('updateProfileBanner', formData)
            .then(() => {
              this.clearFile('banner')
            })
        )
      }

      if (promises.length < 1) {
        this.showFormSpinner = false
        return
      }

      return Promise.all(promises)
        .then(() => {
          this.showUpdateSuccess('Profile updated!')
          this.showFormSpinner = false
        })
        .catch(err => {
          console.error('File error: ', err)
          this.formUpdateError = err
          this.showFormSpinner = false
        })
    },
    submitForm (event) {
      event.preventDefault()

      this.formUpdateSuccess = ''
      this.formUpdateError = ''
      this.showFormSpinner = true

      // Validate Input
      if (!this.validForm()) {
        this.showFormSpinner = false
        return
      }

      this.saveSettings()
    },
    validForm () {
      let validFullname = (this.form.fullname === this.fullname) || (!this.$v.form.fullname.$invalid)
      let validUsername = (this.form.username === this.username) || (!this.$v.form.username.$invalid)
      let validEmail = (this.form.email === this.email) || (!this.$v.form.email.$invalid)
      let validIcon = this.validIconFile
      let validBanner = this.validBannerFile

      return validFullname && validUsername && validEmail && validIcon && validBanner
    },
    validImageFile (name) {
      let image = this.$refs[name]

      if (!image) {
        return true
      }

      let addedNewFile = image.files.length > 0

      if (addedNewFile) {
        return image.files[0].type.match('image.*')
      }

      return true
    },
    clearFile (name) {
      let element = this.$refs[name]

      if (!element) {
        return
      }

      element.value = null
    },
    initForm () {
      this.form.fullname = this.fullname
      this.form.username = this.username
      this.form.email = this.email
      this.form.icon = this.icon
      this.form.banner = this.banner
    },
    showUpdateSuccess (message) {
      this.formUpdateSuccess = message
      window.setTimeout(() => {
        this.formUpdateSuccess = ''
      }, 5000)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/input';

$icon-size: 70px;

.settings-profile-basic {
  display: block;
  position: relative;
}

.icon-container {
  .user-img-container {
    width: $icon-size;
    height: $icon-size;
    overflow: hidden;
    text-align: center;
    white-space: nowrap;
    float: left;
    background-color: #68747F;
  }

  .user-img {
    border: 0;
    height: $icon-size;
  }
}

.banner-container {
  position: relative;
  width: 100%;

  .banner-img-container {
    position: relative;
    width: 100%;
    padding-bottom: 16%;
  }

  .banner-img {
    position: absolute;
    background-color: #68747F;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    bottom: 0;
    top: 0;
    right: 0;
    left: 0;
    overflow: hidden;
  }
}

.input-label {
  margin-bottom: 8px;
}

.change {
  color: $form-primary-color;
}

.success {
  margin-bottom: 20px;
  color: $form-primary-color;
}

.error {
  margin-bottom: 20px;
  color: #cc0000;
}

.saveMessage {
  text-align: center;
  margin-top: 12px;
  margin-bottom: 0;
}

.spinner {
  position: relative;
}
</style>
