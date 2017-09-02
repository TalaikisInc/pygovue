<template>
  <ins class="adsbygoogle"
    :data-ad-client="adClient"
    :data-ad-slot="adSlot"
    :data-ad-format="adFormat"
    :style="adStyle"></ins>
</template>

<script>
// Component from https://github.com/darrynten/vue-adsense
export default {
  props: {
    adClient: {
      type: String,
      required: true
    },
    adSlot: {
      type: String,
      required: true
    },
    adFormat: {
      type: String,
      required: false,
      default: 'auto'
    },
    adStyle: {
      type: String,
      required: false,
      default: 'display: block'
    }
  },
  mounted () {
    (window.adsbygoogle = window.adsbygoogle || []).push({})
  }
}
</script>