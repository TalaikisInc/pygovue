<template>
  <div class="post-content">
    <adsense
      :ad-client="adClient"
      :ad-slot="adSlot"
      ad-style="display: block"
      ad-format="auto">
    </adsense>
  </div>
</template>

<script>
import VueAdsense from './VueAdsense.vue'

export default {
  name: 'adsComp',
  data () {
    return {
      imgBaseUrl: process.env.imgBaseUrl,
      adClient: process.env.adClient,
      adSlot: process.env.adSlot
    }
  },
  components: {
    'adsense': VueAdsense
  }
}
</script>
