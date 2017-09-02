<template>
<div>
  <header-component></header-component>
    <div class="col-sm-9">
      <div class="row">
        <div class="col-sm-8">
          <div class="tr-content">
            <ad-component></ad-component>
          <div class="tr-section bg-transparent">
            <h1>Top tags</h1>
            <div class="row" v-for="chunk in chunkPosts">
              <div class="col-md-3 medium-post" v-for="tag in chunk">
                <div class="tr-post">
                  <div class="post-content">
                    <a :href="baseUrl + 'tag/' + tag.slug + '/'">{{ tag.title }} [{{ tag.post_count }}]</a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="col-sm-4 tr-sidebar">
        <div>
          <ad-component></ad-component>
          <div class="tr-section tr-widget tr-ad ad-before">
            <popular-posts></popular-posts>
          </div>
        </div>
      </div>
    </div>
  </div>
<footer-component></footer-component>
</div>
</template>

<script>
import axios from 'axios'
import chunk from '../plugins/chunk'

import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import Popular from '../components/PopularSidebar.vue'
import Ads from '../components/Ads.vue'

export default {
  data () {
    return {
      tags: [],
      baseUrl: process.env.baseUrl
    }
  },
  asyncData ({ req, params }) {
    return axios.get('/top_tags/')
      .then((response) => {
        return { tags: response.data }
      })
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  computed: {
    chunkPosts () {
      return chunk(this.tags, 4)
    }
  },
  components: {
    'header-component': Header,
    'footer-component': Footer,
    'popular-posts': Popular,
    'ad-component': Ads
  },
  head () {
    return {
      title: 'Top subjects | ' + this.title
    }
  }
}
</script>

<style></style>