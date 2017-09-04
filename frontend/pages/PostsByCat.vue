<template>
<div>
  <header-component></header-component>
    <div class="col-sm-9">
      <div class="row">
        <div class="col-sm-8">
          <ad-component :type="0"></ad-component>
          <h1>{{ title }}<span v-if="page > 0">, page {{ page }}</span></h1>
          <div class="row" v-for="chunk in chunkPosts">
            <div class="col-md-6" v-for="post in chunk">
              <div v-if="post.image">
                <a :href="baseUrl + post.slug + '/'">
                <img class="img-responsive" :src="imgBaseUrl + post.image" :alt="post.title">
                </a>
              </div>
              <div>
                By <a :href="baseUrl+'/source/'+post.category_id.Slug + '/'">{{ post.category_id.Title }}</a>
                 | {{ post.date | formatDate }}
              </div>
              <h2><a :href="baseUrl + post.slug + '/'">{{ post.title }}</a></h2>
              <p v-if="post.content">{{ post.content }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  <paginator-component v-once :pages="calcPages" :source="type" value="" :active="page"></paginator-component>
  <footer-component></footer-component>
</div>
</template>

<script>
import chunk from '../plugins/chunk'

import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import Paginator from '../components/Paginator.vue'
import Ads from '../components/Ads.vue'

import axios from 'axios'

export default {
  data () {
    return {
      posts: [],
      baseUrl: process.env.BASE_URL,
      imgBaseUrl: process.env.IMG_URL,
      title: process.env.SITE_NAME,
      category: null,
      page: null,
      type: 1
    }
  },
  asyncData ({ req, params, error }) {
    return axios.get('/cat/' + params.catSlug + '/' + (Number(params.page) || '0') + '/')
      .then((response) => {
        return { posts: response.data, category: params.catSlug, page: Number(params.page) || 0 }
      })
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  components: {
    'header-component': Header,
    'footer-component': Footer,
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
      title: Number(this.$route.params.page) ? 'Page ' + Number(this.$route.params.page) + ' ' + this.$route.params.catSlug + ' | ' + this.title : this.title
    }
  }
}
</script>

<style>
</style>
