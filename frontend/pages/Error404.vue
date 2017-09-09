<template>
<div clas="row align-items-center">
  <div class="row-sm-3"></div>
  <div class="row-sm-6 text-center">
    <div>
      <h1>404</h1>
      <p>Page not found.</p>
      <a :href="baseUrl">Back to Home</a>
    </div>
  </div>
  <div class="row-sm-3"></div>
</div>
</template>

<script>
export default {
  name: 'errorPage',
  data () {
    return {
      baseUrl: process.env.BASE_URL,
      logoAlt: process.env.SITE_NAME
    }
  }
}
</script>

<style>
</style>