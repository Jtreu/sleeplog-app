<template lang="html">
  <div class="password-reset-dialog">
    <div class="content">
      <button class="backBtn linkBtn" v-on:click="navigateBack()">{{ label.BACK }}</button>
      <form v-if="!formSubmitSuccess"
        class="form" v-on:submit="submitForm($event)">
        <label for="email">{{ label.EMAIL }}</label>
        <input v-model.trim="form.email"
          type="text"
          id="username"
          :placeholder="label.USERNAME"
          class="username">
        <div v-if="form.isSubmitted && $v.form.email.$invalid" class="error">
          <span v-if="!$v.form.email.required">Your Email is required.</span>
          <span v-else-if="!$v.form.email.email">Email must be valid.</span>
        </div>
        <input type="submit" class="sendBtn" :value="label.SEND_RESET_CODE">
      </form>
      <div v-if="formSubmitSuccess"
        class="">{{ label.RESET_CODE_SENT_MESSAGE }}</div>
    </div>
  </div>
</template>

<script>
import { required, email } from 'vuelidate/lib/validators'

export default {
  name: 'sl-password-reset-dialog',
  data () {
    return {
      formSubmitSuccess: false,
      form: {
        email: '',
        isSubmitted: false
      },
      label: {
        EMAIL: 'Email',
        BACK: 'Back',
        SEND_RESET_CODE: 'Send Reset Code',
        RESET_CODE_SENT_MESSAGE: 'A password reset code has been sent to {{ email }}. Follow the instructions to reset your password.'
      }
    }
  },
  validations: {
    form: {
      email: {
        required,
        email
      }
    }
  },
  methods: {
    submitForm (event) {
      event.preventDefault()

      this.form.isSubmitted = true

      if (this.$v.form.$invalid) {
        return
      }

      let email = this.form.email

      this.$store.dispatch('resetPassword', email)
        .then(() => {
          // Password Reset code sent!
          this.label.RESET_CODE_SENT_MESSAGE = this.label.RESET_CODE_SENT_MESSAGE.replace('{{ email }}', email)
          this.formSubmitSuccess = true
        })
    },
    navigateBack () {
      this.$emit('back')
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';
@import '../assets/styles/input';

.password-reset-dialog {
  position: relative;
  display: block;
  background: #FFFFFF;
  width: 280px;
}

.content {
  position: relative;
  padding: 40px 20px;
  margin: auto;
  text-align: left;
}

.form {
  text-align: left;
}

.linkBtn {
  background: none !important;
  border: none;
  padding: 0 !important;
  font: inherit;
  color: $form-primary-color;
  cursor: pointer;
}

.backBtn {
  margin-bottom: 20px;
}

.error {
  margin-bottom: 20px;
  color: #cc0000;
}
</style>
