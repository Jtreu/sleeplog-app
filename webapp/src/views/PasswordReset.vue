<template lang="html">
  <div class="password-reset-view">
    <form class="form" v-on:submit="submitForm($event)">
      <label for="password">{{ label.PASSWORD }}</label>
      <input v-model.trim="form.password"
        type="password"
        id="password"
        class="password">
      <div v-if="form.isSubmitted && $v.form.password.$invalid" class="error">
        <span v-if="!$v.form.password.required">Password is required.</span>
        <span v-else-if="!$v.form.password.passwordHasMinLength">Password must contain at least {{ passwordMinLength }} characters</span>
        <span v-else-if="!$v.form.password.passwordContainsLetter">Password must contain at least 1 letter</span>
        <span v-else-if="!$v.form.password.passwordContainsNumber">Password must contain at least 1 number</span>
        <span v-else-if="!$v.form.password.passwordContainsUppercase">Password must contain at least 1 uppercase character</span>
        <span v-else-if="!$v.form.password.passwordContainsSpecial">Password must contain at least 1 special character</span>
      </div>
      <label for="passwordRepeat">{{ label.REPEAT_PASSWORD }}</label>
      <input v-model.trim="form.passwordRepeat"
        type="password"
        id="passwordRepeat"
        class="password">
      <div v-if="form.isSubmitted && $v.form.passwordRepeat.$invalid" class="error">
        <span v-if="!$v.form.passwordRepeat.required">Repeat Password is required</span>
        <span v-else-if="!$v.form.passwordRepeat.sameAsPassword">Passwords must match</span>
      </div>
      <input type="submit" :value="label.RESET_PASSWORD">
    </form>
    <div v-if="submissionMessage" v-bind:class="{ 'error': !submissionStatusSuccess, 'success': submissionStatusSuccess }" class="submission-message">
      <span>{{ submissionMessage }}</span>
    </div>
  </div>
</template>

<script>
import { required, sameAs } from 'vuelidate/lib/validators'
import { passwordMinLength, passwordHasMinLength, passwordContainsLetter, passwordContainsNumber, passwordContainsUppercase, passwordContainsSpecial } from '../utils/formValidator'

import Spinner from '../components/Spinner'

export default {
  name: 'email-confirmation',
  components: {
    'sl-spinner': Spinner
  },
  data () {
    return {
      submissionMessage: '',
      submissionStatusSuccess: false,
      form: {
        password: '',
        passwordRepeat: '',
        isSubmitted: false
      },
      label: {
        PASSWORD: 'Password',
        REPEAT_PASSWORD: 'Repeat Password',
        RESET_PASSWORD: 'Reset Password',
        SUCCESS_MESSAGE: '',
        ERROR_MESSAGE: ''
      },
      passwordMinLength: passwordMinLength
    }
  },
  validations: {
    form: {
      password: {
        required,
        passwordHasMinLength,
        passwordContainsLetter,
        passwordContainsNumber,
        passwordContainsUppercase,
        passwordContainsSpecial
      },
      passwordRepeat: {
        required,
        sameAsPassword: sameAs('password')
      }
    }
  },
  methods: {
    submitForm (event) {
      event.preventDefault()

      this.form.isSubmitted = true
      this.submissionMessage = ''

      if (this.$v.form.$invalid) {
        return
      }

      let code = this.$route.params.code
      let password = this.form.password

      this.$store.dispatch('updatePasswordWithCode', { code, password })
        .then(message => {
          this.submissionMessage = message
          this.submissionStatusSuccess = true
        })
        .catch(errMessage => {
          this.submissionMessage = errMessage
          this.submissionStatusSuccess = false
        })
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';
@import '../assets/styles/input';

.password-reset-view {
  margin: 0 16px;
  color: $text-color;
}

.form {
  text-align: left;
}

.submission-message {
  margin-top: 20px;
  text-align: center;

  &.success {
    color: $form-primary-color;
  }
}

.error {
  margin-bottom: 20px;
  color: #cc0000;
}
</style>
