<template lang="html">
  <div class="settings-profile-description">
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
        <label for="description">
          {{ label.DESCRIPTION }}
        </label>
        <textarea id="description"
          v-model.trim="form.description"
          rows="10"></textarea>
        <div v-if="(form.description || formSubmitted) && $v.form.description.$invalid" class="error"></div>
      </div>
      <input type="submit" :value="label.SAVE">
    </form>
  </div>
</template>

<script>
import { mapState } from 'vuex'

import Spinner from '../components/Spinner'

export default {
  name: 'sl-settings-profile-description',
  components: {
    'sl-spinner': Spinner
  },
  data () {
    return {
      formSubmitted: false,
      form: {
        description: ''
      },
      label: {
        DESCRIPTION: 'Description',
        SAVE: 'Save Description',
        REQUIRES_SAVE: '- requires save'
      },
      showFormSpinner: false,
      formUpdateError: '',
      formUpdateSuccess: ''
    }
  },
  validations: {
    form: {
      description: { }
    }
  },
  computed: {
    ...mapState({
      description: state => state.profile.description
    })
  },
  mounted () {
    this.initForm()
  },
  methods: {
    saveSettings () {
      this.$store.dispatch('updateProfileDescription', this.form.description)
        .then(() => {
          this.showFormSpinner = false
          this.showUpdateSuccess('Description updated!')
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
    showUpdateSuccess (message) {
      this.formUpdateSuccess = message
      window.setTimeout(() => {
        this.formUpdateSuccess = ''
      }, 5000)
    },
    initForm () {
      this.form.description = this.description
    }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/input';

$icon-size: 70px;

.settings-profile-description {
  display: block;
  position: relative;
}

#description {

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
