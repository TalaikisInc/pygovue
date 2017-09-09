<template>
    <footer>
        <div class="row align-items-center">
            <div class="col-sm-3"></div>
            <div class="col-sm-6 text-center">
                <div>
                    <a :href="baseUrl"><img class="img-responsive" src="~/assets/logo/logo.png" :alt="logoAlt"></a>
                </div>
                <div>
                    <p><a :href="'https://twitter.com/' + twHandle"><i class="fa fa-twitter"></i></a>&nbsp;
                    <a :href="'https://www.facebook.com/' + fbHandle"><i class="fa fa-facebook"></i></a></p>
                    <address>
                        <p>&copy; 2017 <a :href="baseUrl">{{ siteName }}</a> | 
                        Developed by <a href="https://talaikis.com">Talaikis Inc.</a></p>
                    </address>					
                </div>
            </div>
            <div class="col-sm-3"></div>
        </div>
    </footer>
</template>

<script>
export default {
  name: 'footerComponent',
  data () {
    return {
      baseUrl: process.env.BASE_URL,
      logoAlt: process.env.SITE_NAME,
      twHandle: process.env.TWITTE_HANDLE,
      fbHandle: process.env.FACEBOOK_HANDLE,
      siteName: process.env.SITE_NAME
    }
  }
}
</script>

<style></style>