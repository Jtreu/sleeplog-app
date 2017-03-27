<template lang="html">
  <div class="signup-dialog">
    <div class="content">
      <!-- Select basic info -->
      <div v-if="currentForm === '1'" class="basic-prompt">
        <form class="form" v-on:submit="submitForm1($event)">
          <!-- Full Name -->
          <label for="fullname">{{ label.FULLNAME }}</label>
          <input v-model.trim="form1.fullname"
            type="text"
            id="fullname"
            class="fullname"
            name="name">
          <div v-if="form1.isSubmitted && $v.form1.fullname.$invalid" class="error">
            <span v-if="!$v.form1.fullname.required">Your Full Name is required.</span>
            <span v-else-if="!$v.form1.fullname.maxLength">Name must be shorter than 20 characters.</span>
          </div>
          <!-- Email -->
          <label for="email">{{ label.EMAIL }}</label>
          <input v-model.trim="form1.email"
            type="text"
            id="email"
            class="email"
            name="email">
            <div v-if="form1.isSubmitted && !$v.form1.email.$pending && $v.form1.email.$invalid" class="error">
              <span v-if="!$v.form1.email.required">Your Email is required.</span>
              <span v-else-if="!$v.form1.email.email">Email must be valid.</span>
              <span v-else-if="!$v.form1.email.isUnique">This email is unavailable.</span>
            </div>
          <!-- Password -->
          <label for="password">{{ label.PASSWORD }}</label>
          <input v-model.trim="form1.password"
            type="password"
            id="password"
            class="password">
          <div v-if="form1.isSubmitted && $v.form1.password.$invalid" class="error">
            <span v-if="!$v.form1.password.required">Password is required.</span>
            <span v-else-if="!$v.form1.password.passwordHasMinLength">Password must contain at least {{ passwordMinLength }} characters</span>
            <span v-else-if="!$v.form1.password.passwordContainsLetter">Password must contain at least 1 letter</span>
            <span v-else-if="!$v.form1.password.passwordContainsNumber">Password must contain at least 1 number</span>
            <span v-else-if="!$v.form1.password.passwordContainsUppercase">Password must contain at least 1 uppercase character</span>
            <span v-else-if="!$v.form1.password.passwordContainsSpecial">Password must contain at least 1 special character</span>
          </div>
          <input type="submit" class="signupBtn" :value="label.SIGNUP">
        </form>
        <div class="signin-prompt-container">
          <span>{{ label.OLD_USER }}</span>
          <button v-on:click="showDialog('signin')"
            class="signinBtn linkBtn">{{ label.SIGNIN }}</button>
        </div>
      </div>
      <!-- Select username -->
      <div v-if="currentForm === '2'" class="username-prompt">
        <button class="backBtn linkBtn" v-on:click="navigateToForm('1')">{{ label.BACK }}</button>
        <h2>{{ label.CHOOSE_USERNAME }}</h2>
        <p>{{ label.CHOOSE_USERNAME_DETAIL }}</p>
        <form class="form" v-on:submit="submitForm2($event)">
          <label for="username">sleeplog-app.com/{{ form2.username.toLowerCase() }}</label>
          <input v-model.trim="form2.username"
            type="text"
            id="username"
            :placeholder="label.USERNAME"
            class="username">
          <div v-if="form2.isSubmitted && !$v.form2.username.$pending && $v.form2.username.$invalid" class="error">
            <span v-if="!$v.form2.username.required">Username is required.</span>
            <span v-else-if="!$v.form2.username.maxLength">Username must be shorter than 15 characters.</span>
            <span v-else-if="!$v.form2.username.nonSpecial">No special characters.</span>
            <span v-else-if="!$v.form2.username.isUnique">This Username taken.</span>
          </div>
          <div v-else-if="$v.form2.username.$pending || !$v.form2.username.$invalid"  class="success">
            <span>This Username is available.</span>
          </div>
          <input type="submit" class="readyBtn" :value="label.READY">
        </form>
        <div v-if="registerAttemptStatus" class="signUpError error">
          <span>Error: {{ registerAttemptStatus }}</span>
        </div>
        <div v-show="showSpinner" class="spinner">
          <sl-spinner class="spinner-icon"></sl-spinner>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { required, email, maxLength } from 'vuelidate/lib/validators'
import { passwordMinLength,
  passwordHasMinLength,
  passwordContainsLetter,
  passwordContainsNumber,
  passwordContainsUppercase,
  passwordContainsSpecial } from '../utils/formValidator'

import Spinner from '../components/Spinner'

export default {
  name: 'sl-signup-dialog',
  components: {
    'sl-spinner': Spinner
  },
  data () {
    return {
      currentForm: '1',
      registerAttemptStatus: '',
      showSpinner: false,
      form1: {
        fullname: '',
        email: '',
        password: '',
        isSubmitted: false
      },
      form2: {
        username: '',
        isSubmitted: false
      },
      label: {
        FULLNAME: 'Full name',
        EMAIL: 'Email',
        USERNAME: 'Username',
        PASSWORD: 'Password',
        SIGNUP: 'Sign up',
        SIGNIN: 'Sign in',
        READY: 'Ready!',
        BACK: 'Back',
        CHOOSE_USERNAME: 'Choose a Username',
        CHOOSE_USERNAME_DETAIL: 'Don\'t worry, you can always change it later.',
        OLD_USER: 'Have an account?'
      },
      passwordMinLength: passwordMinLength
    }
  },
  validations: {
    form1: {
      fullname: {
        required,
        maxLength: maxLength(20)
      },
      email: {
        required,
        email,
        isUnique (value) {
          if (value === '') return true
          return this.$store.dispatch('checkEmailAvailability', value) || true
        }
      },
      password: {
        required,
        passwordHasMinLength,
        passwordContainsLetter,
        passwordContainsNumber,
        passwordContainsUppercase,
        passwordContainsSpecial
      }
    },
    form2: {
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
      }
    },
    validationGroup: ['form1.fullname', 'form1.email', 'form1.password', 'form2.username']
  },
  methods: {
    submitForm1 (event) {
      event.preventDefault()

      this.form1.isSubmitted = true

      // Validate Input
      if (this.$v.form1.$invalid) {
        window.setTimeout(this.submitForm1AfterAsync, 500)
        return
      }

      // Next form
      this.navigateToForm('2')
    },
    submitForm2 (event) {
      event.preventDefault()

      this.showSpinner = true
      this.form2.isSubmitted = true

      // Validate Input
      if (this.$v.form2.$invalid) {
        window.setTimeout(this.submitForm2AfterAsync, 500)
        return
      }

      this.register()
    },
    submitForm1AfterAsync () {
      if (this.$v.form1.$invalid) {
        return
      }

      this.navigateToForm('2')
    },
    submitForm2AfterAsync () {
      if (this.$v.form2.$invalid) {
        this.showSpinner = false
        return
      }

      this.register()
    },
    register () {
      if (this.$v.validationGroup.$invalid) {
        return
      }

      this.$store.dispatch('register', {
        email: this.form1.email,
        username: this.form2.username,
        alias: this.form1.fullname,
        password: this.form1.password
      }).then(() => {
        // Registration complete!

        this.showSpinner = false
        this.$emit('show-welcome')
      })
      .catch(err => {
        this.showSpinner = false
        this.registerAttemptStatus = err
      })
    },
    navigateToForm (num) {
      this.currentForm = num
    },
    showDialog (option) {
      this.$emit('dialog-selected', option)
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';
@import '../assets/styles/input';

.signup-dialog {
  position: relative;
  display: block;
  background: #FFFFFF;
  width: 280px;
  overflow: hidden;
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

.linkBtn {
  background:none!important;
  border:none;
  padding:0!important;
  font: inherit;
  color: $form-primary-color;
  cursor: pointer;
}

.signinBtn {
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

.backBtn {
  margin-bottom: 20px;
}

.username-prompt {
  text-align: left;
}

.success {
  margin-bottom: 20px;
  color: $form-primary-color;
}

.error {
  margin-bottom: 20px;
  color: #cc0000;

  &.signUpError {
    margin-top: 12px;
    margin-bottom: 0;
  }
}

</style>
