<template>
  <div class="col-md-12">
    <ad-component></ad-component>
    <h1>Categories<span v-if="page > 0">, page {{ page }}</span></h1>
    <div class="row" v-for="chunk in chunkCats">
      <div class="col-md-6 card bg-light mb3" style="max-width: 20rem;" v-for="cat in chunk">
        <div  class="card-body">
          <h2 class="card-title"><a :href="baseUrl + 'category/' + cat.slug + '/'">{{ cat.title }}</a></h2>
        </div>
      </div>
    </div>
    <paginator-component v-once :totalPages="calcPages" :paginatorType="paginatorType" value="" :currentPage="page" :itemsPerPage="itemsPerPage" :totalItems="categories[0].total_cats">
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
      categories: [],
      baseUrl: process.env.BASE_URL,
      title: process.env.SITE_NAME,
      page: null,
      paginatorType: 2,
      itemsPerPage: 40
    }
  },
  asyncData ({ req, params, error }) {
    return axios.get('/cats/' + (Number(params.page) || '0') + '/')
      .then((response) => {
        return { categories: response.data, page: Number(params.page) || 0 }
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
    chunkCats () {
      return chunk(this.categories, 4)
    },
    calcPages () {
      const pages = Math.floor(this.categories[0].total_cats / 40)
      return pages <= 250 ? pages : 250
    }
  },
  head () {
    return {
      title: Number(this.$route.params.page) ? 'Page ' + Number(this.$route.params.page) + ' Categories | ' + this.title : this.title
    }
  }
}
</script>

<style>
</style>
