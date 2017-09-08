<template>
<div>
    <header-component></header-component>
        <div>
            <a class="navbar-brand" :href="baseUrl"><img class="img-responsive" src="~/assets/logo/logo.png" :alt="logoAlt"></a>
        </div>
        <div>
            <h1>404</h1>
            <p>Page not found.</p>
            <a :href="baseUrl">Back to Home</a>
        </div>
    <footer-component></footer-component>
</div>
</template>

<script>
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'

export default {
  name: 'errorPage',
  data () {
    return {
      baseUrl: process.env.BASE_URL,
      logoAlt: process.env.SITE_NAME
    }
  },
  components: {
    'header-component': Header,
    'footer-component': Footer
  }
}
</script>

<style>
</style>