<template>
  <div>
    <a :href="baseUrl"><img src="~/assets/logo/logo.png" :alt="logoAlt" :title="logoAlt"></a>
  </div>
</template>

<script>
export default {
  name: 'headerComponent',
  data () {
    return {
      baseUrl: process.env.BASE_URL,
      logoAlt: process.env.SITE_NAME
    }
  }
}
</script>