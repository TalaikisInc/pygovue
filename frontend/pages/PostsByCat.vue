<template>
  <div class="col-md-12">
    <ad-component></ad-component>
    <h1>{{ catTitle }}<span v-if="page > 0">, page {{ page }}</span></h1>
    <div class="row" v-for="chunk in chunkPosts">
      <div class="col-md-6 card bg-light mb3" style="max-width: 20rem;" v-for="post in chunk">
        <div class="card-header">
          <a :href="baseUrl+'/source/'+post.category_id.Slug + '/'">{{ post.category_id.Title }}</a>
            | {{ post.date | formatDate }}
        </div>
        <div v-if="post.image" class="card-img-top">
          <a :href="baseUrl + post.slug + '/'">
            <img class="img-responsive" :src="imgBaseUrl + post.image" :alt="post.title">
          </a>
        </div>
        <div  class="card-body">
          <h2 class="card-title"><a :href="baseUrl + post.slug + '/'">{{ post.title }}</a></h2>
          <p class="card-text" v-if="post.content">{{ post.content }}</p>
        </div>
      </div>
    </div>
    <paginator-component v-once :totalPages="calcPages" :paginatorType="paginatorType" :value="catSlug" :currentPage="page" :itemsPerPage="itemsPerPage" :totalItems="posts[0].total_posts">
    </paginator-component>
  </div>
</template>

<script>
import chunk from '../plugins/chunk'
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
      page: null,
      paginatorType: 1,
      itemsPerPage: 20,
      catSlug: null,
      catTitle: null
    }
  },
  asyncData ({ req, params, error }) {
    return axios.get('/cat/' + params.catSlug + '/' + (Number(params.page) || '0') + '/')
      .then((response) => {
        return { posts: response.data, catSlug: params.catSlug, catTitle: response.data[0].category_id.Title, page: Number(params.page) || 0 }
      })
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  components: {
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
      title: Number(this.$route.params.page) ? 'Page ' + Number(this.$route.params.page) + ' ' + this.catTitle + ' | ' + this.title : this.title
    }
  }
}
</script>

<style>
</style>