<template>
  <div class="container">
    <header-component></header-component>
    <nuxt/>
    <footer-component></footer-component>
  </div>
</template>

<script>
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'

export default {
  data () {
    return { }
  },
  components: {
    'header-component': Header,
    'footer-component': Footer
  }
}
</script>

<style>
</style>
