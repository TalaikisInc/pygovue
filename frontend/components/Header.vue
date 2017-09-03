<template>
<div class="col-sm-3">
    <nav>
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar-collapse">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a :href="baseUrl"><img class="img-responsive" src="~/assets/logo/qprob.png" :alt="logoAlt"></a>
        </div>
        <div class="collapse navbar-collapse" id="navbar-collapse">
            <ul class="navbar-nav">
                <li v-for="cat in categories">
                    <a :href="baseUrl+'category/' + cat.slug + '/'" :title="cat.title">
                    <i class="fa fa-list" aria-hidden="true">&nbsp;</i>
                    {{ cat.title }} [{{ cat.post_count }}]</a>
                </li>
                <li>All categories</li>
            </ul>
        </div>
    </nav>
</div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'headerComp',
  data () {
    return {
      categories: this.categories,
      baseUrl: process.env.BASE_URL,
      logoAlt: process.env.SITE_NAME,
      imgBaseUrl: process.env.IMG_URL
    }
  },
  methods: {
    fetchData () {
      axios.get('/cats/0/').then(response => {
        this.categories = response.data
      }).catch(e => {
        console.log(e)
      })
    }
  },
  mounted () {
    this.fetchData()
  }
}
</script>