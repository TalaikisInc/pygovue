<template>
  <social-sharing :url="url" :title="title" inline-template>
  <div>
    <network network="facebook">
      <i class="fa fa-fw fa-facebook"></i>&nbsp;
    </network>
    <network network="googleplus">
      <i class="fa fa-fw fa-google-plus"></i>&nbsp;
    </network>
    <network network="linkedin">
      <i class="fa fa-fw fa-linkedin"></i>&nbsp;
    </network>
    <network network="reddit">
      <i class="fa fa-fw fa-reddit"></i>&nbsp;
    </network>
    <network network="twitter">
      <i class="fa fa-fw fa-twitter"></i>&nbsp;
    </network>
  </div>
</social-sharing>
</template>

<script>
const SocialSharing = require('vue-social-sharing')

export default {
  name: 'socialSharingComp',
  props: {
    url: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    }
  },
  components: {
    'social-sharing': SocialSharing
  }
}
</script>