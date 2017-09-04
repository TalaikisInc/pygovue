<template>
<div>
  <header-component></header-component>
    <div class="col-sm-9">
      <div class="row">
        <div class="col-sm-8">
          <ad-component :type="0"></ad-component>
          <div class="col-md-12">
            <h1>{{ post.title }}</h1>
            <div v-if="post.image">
              <a :href="baseUrl + post.slug+'/'"><img class="img-responsive" :src="imgBaseUrl + post.image" :alt="post.title"></a>
            </div>
            <div>
              >By <a :href="baseUrl+'/source/'+post.category_id.Slug + '/'">{{ post.category_id.Title }}</a>
                | {{ post.date | formatDate }}
            </div>
            <h2><a :href="baseUrl + post.slug + '/'">{{ post.title }}</a></h2>
            <p v-if="post.content">{{ post.content }}</p>
            <div>
              <social-sharing :url="baseUrl + post.slug + '/'" :title="post.title"></social-sharing>
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
import axios from 'axios'
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import SocialSharing from '../components/SocialSharing.vue'
import Ads from '../components/Ads.vue'

export default {
  name: 'post',
  data () {
    return {
      post: null,
      baseUrl: process.env.BASE_URL,
      imgBaseUrl: process.env.IMG_URL,
      title: process.env.SITE_NAME
    }
  },
  asyncData ({ req, params, error }) {
    return axios.get('/post/' + params.postSlug + '/')
      .then((response) => {
        return { post: response.data[0] }
      })
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  components: {
    'header-component': Header,
    'footer-component': Footer,
    'social-sharing': SocialSharing,
    'ad-component': Ads
  },
  head () {
    return {
      title: this.post.title + ' | ' + this.title
    }
  }
}
</script>
