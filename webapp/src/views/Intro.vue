<template lang="html">
  <!DOCTYPE html>
  <html>
    <head>
      <link href="https://fonts.googleapis.com/css?family=Slabo+27px" rel="stylesheet">
      <meta charset="utf-8">
      <title>Intro</title>
    </head>
    <body>
      <h1>Sleep Log App</h1>
      <h2>Welcome!</h2>
      <h3>Instructions</h3>
        <div class="content">
          <div class="row">
            <div class="col-xs-8">
            1. If you don't have an account, create one! It's free!<br />
            2. Log in.<br />
            3. When viewing your profile, you have will a table of dates and times, each of these cells contains the activities you did during that hour.<br />
            4. If you want to add an activity, find the correct date and time, and click inside the box.<br />
            5. This will bring up a dialog to add activities, you can add as many as you need, but you can only add each activity once.<br />
            6. When you have added all activities, make sure you click the checkmark next to the each activity.<br />
            7. When you click out of the dialog window, you will have succesffully added activites to your log!<br />
            8. Sleep Tight!<br />
            </div>
          </div>
        </div>
    </body>
  </html>

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

h1 {
    font-size: 4em;
}

h2 {
  font-size: 2em;
}

h3 {
  font-size: 1.25em;
  text-decoration: underline;
}

.content {
  min-width: 500px;
  margin-top: 2%;
  margin-left: 27%;
  text-align: left;
  font-family: 'Slabo 27px', serif;
  font-size: 1.25em;
  font-weight: bold;
  background-color: rgb(229, 240, 227);
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
