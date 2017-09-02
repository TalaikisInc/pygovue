<template>
<div>
  <header-component></header-component>
    <div class="col-sm-9">
      <div class="row">
        <div class="col-sm-8">
          <div class="tr-content">
            <ad-component :type="0"></ad-component>

          <div class="tr-section bg-transparent">
            <h1>Popular<span v-if="page > 0">, page {{ page }}</span></h1>
            <div class="row" v-for="chunk in chunkPosts">
              <div class="col-md-6 medium-post" v-for="post in chunk">
                <div class="tr-post">
                  <div class="entry-header">
                    <div class="entry-thumbnail" v-if="post.image">
                      <a :href="baseUrl+post.slug+'/'">
                      <img class="img-responsive" :src="imgBaseUrl+post.image" :alt="post.title"></a>
                    </div>
                  </div>
                  <div class="post-content">
                    <div v-bind:class="[(post.image) ? 'crop' : 'crop-no-img']">
                          <a :href="baseUrl+'source/'+post.category_id.Slug+'/'" v-if="post.category_id.Thumbnail">
                          <img class="img-responsive circle-img" :src="imgBaseUrl+post.category_id.Thumbnail" :alt="post.category_id.Title"></a>
                        </div>
                        <div class="entry-meta">
                          <ul>
                            <li>By <a :href="baseUrl+'source/'+post.category_id.Slug+'/'">{{ post.category_id.Title }}</a></li>
                            <li>{{ post.date | formatDate }}</li>
                            <li>Hits: {{ post.hits }}</li>
                          </ul>
                        </div>
                        <h2 class="entry-title">
                        <a :href="baseUrl+post.slug+'/'">{{ post.title }}</a>&nbsp;
                        <div v-bind:class="[(post.sentiment >= 0) ? 'sentiment-pos' : 'sentiment-neg']" v-if="post.sentiment">[{{ post.sentiment }}]</div></h2>
                        <p c-if="post.summary">{{ post.summary }}</p>
                  </div>
                </div>
              </div>
            </div>
            <paginator-component v-once :pages="calcPages" :soure="type" value="" :active="page"></paginator-component>            
          </div>
        </div>
      </div>

      <div class="col-sm-4 tr-sidebar">
        <div>
          <ad-component :type="1"></ad-component>
          <div class="tr-section tr-widget tr-ad ad-before">
            <recent-posts></recent-posts>
          </div>
        </div>
      </div>
    </div>
  </div>
<footer-component></footer-component>
</div>
</template>

<script>
import chunk from '../plugins/chunk'

import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import Paginator from '../components/Paginator.vue'
import Recent from '../components/RecentSidebar.vue'
import Ads from '../components/Ads.vue'

import axios from 'axios'

export default {
  data () {
    return {
      posts: [],
      baseUrl: process.env.baseUrl,
      imgBaseUrl: process.env.imgBaseUrl,
      title: process.env.siteName,
      page: null,
      type: 3
    }
  },
  asyncData ({ req, params }) {
    return axios.get('/most_popular/' + (Number(params.page) || '0') + '/')
      .then((response) => {
        return { posts: response.data, page: Number(params.page) || 0 }
      })
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  components: {
    'header-component': Header,
    'footer-component': Footer,
    'recent-posts': Recent,
    'paginator-component': Paginator,
    'ad-component': Ads
  },
  computed: {
    chunkPosts () {
      return chunk(this.posts, 2)
    },
    calcPages () {
      const pages = Math.floor(this.posts[0].total_posts / 20)
      return pages <= 250 ? pages : 250
    }
  },
  head () {
    return {
      title: Number(this.$route.params.page) ? 'Page ' + Number(this.$route.params.page) + ' Popular articles | ' + this.title : this.title
    }
  }
}
</script>

<style>
</style>
