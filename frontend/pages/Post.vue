<template>
<div>
  <ad-component :type="0"></ad-component>
  <div class="row">
    <div class="col-sm-3"></div>
    <div class="col-sm-6">
      <h1>{{ post.Title }}</h1>
    </div>
    <div class="col-sm-3"></div>
  </div>
  <div class="row">
    <div class="col-sm-3"></div>
    <div class="col-sm-6">
      <div v-if="post.Image">
        <a :href="baseUrl + post.Slug+'/'">
          <img class="img-fluid" :src="imgBaseUrl + post.Image" :alt="post.Title">
        </a>
      </div>
      <div>
        <a :href="baseUrl + keyword + '/' + post.CategoryID.Slug + '/'">
          {{ post.CategoryID.Title }}
        </a>
         | {{ post.Date | formatDate }}
      </div>
      <p v-if="post.Content">
        {{ post.Content }}
      </p>
      <div>
        <ad-component :type="0"></ad-component>
        <social-sharing :url="baseUrl + post.Slug + '/'" :title="post.Title">
        </social-sharing>
      </div>
      <div class="col">
        <a :href="post.URL" class="btn btn-danger btn-lg pull-right" role="button" aria-disabled="true">
          Read more...
        </a>
      </div>
    </div>
    <div class="col-sm-3"></div>
  </div>
</div>
</template>

<script>
import axios from 'axios'
import SocialSharing from '../components/SocialSharing.vue'
import Ads from '../components/Ads.vue'

export default {
  name: 'post',
  data () {
    return {
      post: null,
      baseUrl: process.env.BASE_URL,
      imgBaseUrl: process.env.IMG_URL,
      title: process.env.SITE_NAME,
      keyword: process.env.KEYWORD
    }
  },
  asyncData ({ req, params, error }) {
    return axios.get('/post/' + params.postSlug + '/')
      .then((response) => {
        return { post: response.data }
      })
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  components: {
    'social-sharing': SocialSharing,
    'ad-component': Ads
  },
  head () {
    return {
      title: this.post.Title + ' | ' + this.title
    }
  }
}
</script>
