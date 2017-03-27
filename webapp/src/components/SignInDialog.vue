<template lang="html">
  <div class="signin-dialog">
    <div class="content">
      <!-- Sign in form -->
      <form class="form" v-on:submit="submitForm($event)">
        <!-- Username or Email -->
        <label for="usernameOrEmail">{{ label.USERNAME_OR_EMAIL }}</label>
        <input v-model.trim="form.usernameOrEmail"
          type="text"
          id="usernameOrEmail"
          class="usernameOrEmail">
        <div v-if="form.isSubmitted && $v.form.usernameOrEmail.$invalid" class="error">
          <span v-if="!$v.form.usernameOrEmail.required">Username or Email is required.</span>
        </div>
        <!-- Password -->
        <label for="password">{{ label.PASSWORD }}</label>
        <input v-model.trim="form.password"
          type="password"
          id="password"
          class="password">
        <div v-if="form.isSubmitted && $v.form.password.$invalid" class="error">
          <span v-if="!$v.form.password.required">Password is required.</span>
        </div>
        <input type="submit" class="loginBtn" :value="label.LOGIN">
      </form>
      <div v-if="loginAttemptStatus" class="signInError error">
        <span>Error: {{ loginAttemptStatus }}</span>
      </div>
      <!-- Forgot password & Sign up -->
      <button v-on:click="showDialog('passwordReset')"
        class="forgotBtn linkBtn">{{ label.FORGOT }}</button>
      <div class="signup-prompt-container">
        <span>{{ label.NEW }}</span>
        <button v-on:click="showDialog('signup')"
          class="signupBtn linkBtn">{{ label.SIGNUP }}</button>
      </div>
    </div>
  </div>
</template>

<script>
import { required } from 'vuelidate/lib/validators'

export default {
  name: 'sl-signin-dialog',
  data () {
    return {
      loginAttemptStatus: '',
      form: {
        usernameOrEmail: '',
        password: '',
        isSubmitted: false
      },
      label: {
        LOGIN: 'Log in',
        SIGNUP: 'Sign up',
        USERNAME_OR_EMAIL: 'Username or email',
        PASSWORD: 'Password',
        FORGOT: 'Forgot Password?',
        NEW: 'New to Sleeplog-app?'
      }
    }
  },
  validations: {
    form: {
      usernameOrEmail: {
        required
      },
      password: {
        required
      }
    }
  },
  methods: {
    submitForm (event) {
      event.preventDefault()

      this.form.isSubmitted = true
      this.loginAttemptStatus = ''

      // Validate Input
      if (this.$v.form.$invalid) {
        return
      }

      this.$store.dispatch('login', {
        usernameOrEmail: this.form.usernameOrEmail,
        password: this.form.password
      }).then(() => {
        // Login success!
        this.showDialog('welcome')
      }).catch(err => {
        this.loginAttemptStatus = err
      })
    },
    showDialog (option) {
      this.$emit('dialog-selected', option)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';

$primary-color: #22C19A;

.signin-dialog {
  position: relative;
  display: block;
  background: #FFFFFF;
  width: 280px;
}

.content {
  position: relative;
  padding: 40px 20px;
  margin: auto;
  text-align: center;
}

.form {
  text-align: left;
}

label[for] {
  display: block;
}

input[type=text], input[type=password] {
  width: 100%;
  padding: 12px 20px;
  font-size: inherit;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

input[type=submit] {
    width: 100%;
    background-color: $form-primary-color;
    color: #FFFFFF;
    padding: 14px 20px;
    font-size: 16px;
    margin: 8px 0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.forgotBtn, .signupBtn {
  margin-top: 16px;
}

.linkBtn {
  background:none!important;
  border:none;
  padding:0!important;
  font: inherit;
  color: $form-primary-color;
  cursor: pointer;
}

.signup-prompt-container {
  * {
    display: inline-block;
  }
}

.error {
  margin-bottom: 20px;
  color: #cc0000;

  &.signInError {
    margin-top: 12px;
    margin-bottom: 0;
  }
}
</style>
