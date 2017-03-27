<template lang="html">
  <div class="email-confirmation-view">
    <sl-spinner v-if="!message" class="spinner-icon"></sl-spinner>
    <h2>{{ message }}</h2>
  </div>
</template>

<script>
import Spinner from '../components/Spinner'

export default {
  name: 'email-confirmation',
  components: {
    'sl-spinner': Spinner
  },
  data () {
    return {
      message: '',
      label: {
        SUCCESS_MESSAGE: 'Fantastic! Thank you for confirming your email.',
        ERROR_MESSAGE: 'Error: The confirmation code is invalid or your email has already been confirmed.'
      }
    }
  },
  mounted () {
    let code = this.$route.params.code

    this.$store.dispatch('confirmEmail', code)
      .then(() => {
        this.message = this.label.SUCCESS_MESSAGE
      })
      .catch(() => {
        this.message = this.label.ERROR_MESSAGE
      })
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';

.email-confirmation-view {
  margin: 0 16px;
  color: $text-color;
}

.spinner-icon {
  left: 0;
  right: 0;
  top: 0;
  margin: auto;
}
</style>
