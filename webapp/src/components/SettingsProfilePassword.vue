<template lang="html">
  <div class="settings-profile-password">
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
        <label for="oldPassword">
          {{ label.OLD_PASSWORD }}
        </label>
        <input id="oldPassword"
          v-model.trim="form.oldPassword"
          type="password">
        <div v-if="(form.oldPassword || formSubmitted) && $v.form.oldPassword.$invalid" class="error">
          <span v-if="!$v.form.oldPassword.required">Password is required.</span>
          <span v-else-if="!$v.form.oldPassword.passwordHasMinLength">Password must contain at least {{ passwordMinLength }} characters</span>
          <span v-else-if="!$v.form.oldPassword.passwordContainsLetter">Password must contain at least 1 letter</span>
          <span v-else-if="!$v.form.oldPassword.passwordContainsUppercase">Password must contain at least 1 uppercase character</span>
          <span v-else-if="!$v.form.oldPassword.passwordContainsSpecial">Password must contain at least 1 special character</span>
        </div>
      </div>
      <div>
        <label for="newPassword">
          {{ label.NEW_PASSWORD }}
        </label>
        <input id="newPassword"
          v-model.trim="form.newPassword"
          type="password">
        <div v-if="(form.newPassword || formSubmitted) && $v.form.newPassword.$invalid" class="error">
          <span v-if="!$v.form.newPassword.required">Password is required.</span>
          <span v-else-if="!$v.form.newPassword.passwordHasMinLength">Password must contain at least {{ passwordMinLength }} characters</span>
          <span v-else-if="!$v.form.newPassword.passwordContainsLetter">Password must contain at least 1 letter</span>
          <span v-else-if="!$v.form.newPassword.passwordContainsUppercase">Password must contain at least 1 uppercase character</span>
          <span v-else-if="!$v.form.newPassword.passwordContainsSpecial">Password must contain at least 1 special character</span>
        </div>
      </div>
      <div>
        <label for="newPasswordRepeat">
          {{ label.NEW_PASSWORD_REPEAT }}
        </label>
        <input id="newPasswordRepeat"
          v-model.trim="form.newPasswordRepeat"
          type="password">
        <div v-if="(form.newPasswordRepeat || formSubmitted) && $v.form.newPasswordRepeat.$invalid" class="error">
          <span v-if="!$v.form.newPasswordRepeat.required">Password is required.</span>
          <span v-else-if="!$v.form.newPasswordRepeat.sameAsPassword">Repeat Password must match New Password</span>
        </div>
      </div>
      <input type="submit" :value="label.SAVE">
    </form>
  </div>
</template>

<script>
import { required, sameAs } from 'vuelidate/lib/validators'
import { passwordMinLength,
  passwordHasMinLength,
  passwordContainsLetter,
  passwordContainsNumber,
  passwordContainsUppercase,
  passwordContainsSpecial } from '../utils/formValidator'

import Spinner from '../components/Spinner'

export default {
  name: 'sl-settings-profile-password',
  components: {
    'sl-spinner': Spinner
  },
  data () {
    return {
      formSubmitted: false,
      form: {
        oldPassword: '',
        newPassword: '',
        newPasswordRepeat: ''
      },
      label: {
        SAVE: 'Save Password',
        REQUIRES_SAVE: '- requires save',
        OLD_PASSWORD: 'Old Password',
        NEW_PASSWORD: 'New Password',
        NEW_PASSWORD_REPEAT: 'Repeat New Password'
      },
      passwordMinLength: passwordMinLength,
      showFormSpinner: false,
      formUpdateError: '',
      formUpdateSuccess: ''
    }
  },
  validations: {
    form: {
      oldPassword: {
        required,
        passwordHasMinLength,
        passwordContainsLetter,
        passwordContainsNumber,
        passwordContainsUppercase,
        passwordContainsSpecial
      },
      newPassword: {
        required,
        passwordHasMinLength,
        passwordContainsLetter,
        passwordContainsNumber,
        passwordContainsUppercase,
        passwordContainsSpecial
      },
      newPasswordRepeat: {
        required,
        sameAsPassword: sameAs('newPassword')
      }
    }
  },
  methods: {
    saveSettings () {
      this.$store.dispatch('updateProfilePassword', {
        oldPassword: this.form.oldPassword,
        newPassword: this.form.newPassword
      }).then(() => {
        this.showFormSpinner = false
        this.showUpdateSuccess('Password updated!')
        this.clearForm()
      }).catch(err => {
        this.showFormSpinner = false
        this.formUpdateError = err
      })
    },
    submitForm (event) {
      event.preventDefault()

      this.formUpdateSuccess = ''
      this.formUpdateError = ''
      this.showFormSpinner = true
      this.formSubmitted = true

      // Validate Input
      if (this.$v.form.$invalid) {
        this.showFormSpinner = false
        return
      }

      this.saveSettings()
    },
    clearForm () {
      this.form.oldPassword = ''
      this.form.newPassword = ''
      this.form.newPasswordRepeat = ''
      this.formSubmitted = false
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

.settings-profile-password {
  display: block;
  position: relative;
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
