<template>
<div>
  <header-component></header-component>
    <div class="col-sm-9">
      <div class="row">
        <div class="col-sm-8">
          <div class="tr-content">
            <ad-component></ad-component>

            <div class="tr-section bg-transparent">

            <div class="row">
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
                          </ul>
                        </div>
                        <h2 class="entry-title">
                        <a :href="baseUrl+post.slug+'/'">{{ post.title }}</a>&nbsp;
                        <div v-bind:class="[(post.sentiment >= 0) ? 'sentiment-pos' : 'sentiment-neg']" v-if="post.sentiment">[{{ post.sentiment }}]</div></h2>
                        <div class="post-inner-image">
                          <p><span v-for="tag in tags" :key="tag.id"><a :href="baseUrl + 'tag/' + tag.slug + '/'" class="card-link">{{ tag.title }}</a> | </span></p>
                        </div>
                        <div v-if="post.status == 0">
                          <p>{{ post.summary }}</p>
                        </div>
                        <div v-else>
                          <span v-html="post.content"></span>
                        </div>
                        <ad-component></ad-component>
                        <a :href="post.url"><button class="btn btn-primary pull-right">Read more...</button></a>
                        <div class="post-inner-image">
                          <social-sharing :url="baseUrl + post.slug + '/'" :title="post.title"></social-sharing>
                        </div>
                        <div class="post-inner-image" v-if="post.wordcloud">
                          <img :src="imgBaseUrl+post.wordcloud" class="img-responsive" alt="Article analysis">
                        </div>
                        <div class="comment-list">
                          <disqus :shortname="disqusID" :identifier="post.title" :url="baseUrl + post.slug + '/'"></disqus>
                        </div>
                  </div>
                </div>
            </div>
            <!-- paginator -->
          </div>
        </div>
      </div>

      <div class="col-sm-4 tr-sidebar">
        <div>
          <ad-component></ad-component>
          <div class="tr-section tr-widget tr-ad">
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
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import Popular from '../components/PopularSidebar.vue'
import SocialSharing from '../components/SocialSharing.vue'
import Ads from '../components/Ads.vue'
import VueDisqus from 'vue-disqus/VueDisqus.vue'

export default {
  name: 'post',
  data () {
    return {
      post: null,
      tags: [],
      baseUrl: process.env.baseUrl,
      imgBaseUrl: process.env.imgBaseUrl,
      disqusID: process.env.disqusID,
      title: process.env.siteName
    }
  },
  asyncData ({ req, params, error }) {
    return axios.all([axios.get('/post/' + params.postSlug + '/'),
      axios.get('/post_tags/' + params.postSlug + '/')])
      .then(axios.spread(function (posts, tags) {
        return { post: posts.data[0], tags: tags.data }
      }))
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  components: {
    'header-component': Header,
    'footer-component': Footer,
    'popular-posts': Popular,
    'social-sharing': SocialSharing,
    'ad-component': Ads,
    'disqus': VueDisqus
  },
  head () {
    return {
      title: this.post.title + ' | ' + this.title
    }
  }
}
</script>
