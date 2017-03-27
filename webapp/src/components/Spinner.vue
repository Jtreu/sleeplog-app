<template lang="html">
  <div class="spinner-container">
    <div class="spinner"></div>
  </div>
</template>

<script>
export default {
  name: 'sl-spinner',
  data () {
    return { }
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/styles/theme';

$width: 50px;
$thickness: 5px;

.spinner-container {
  padding: calc(#{$thickness} + 10px) 0;
}

.spinner {
  width: $width;
  height: $width;
  border-radius: 100%;
  position: relative;
  margin: 0 auto;

  &:before{
    content: "";
    animation: spin 1s infinite;
  }

  &:after{
    border: $thickness solid #ccc;
  }

  &:before, &:after {
    position: absolute;
    top: -$thickness;
    left: -$thickness;
    width: 100%;
    height: 100%;
    border-radius: 100%;
    border: $thickness solid transparent;
    border-top-color: $app-theme-color;
  }
}

@keyframes spin{
  0% {
    -webkit-transform: rotate(0deg);
    -ms-transform: rotate(0deg);
    -o-transform: rotate(0deg);
    transform: rotate(0deg);
  }

  100% {
    -webkit-transform: rotate(360deg);
    -ms-transform: rotate(360deg);
    -o-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
</style>
