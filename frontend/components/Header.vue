<template>
    <div class="col-sm-3 tr-sidebar">
        <div>
            <div class="tr-menu sidebar-menu">
                <nav class="navbar navbar-default">
                    <div class="navbar-header">
                        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar-collapse">
                            <span class="sr-only">Toggle navigation</span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                        </button>
                        <a :href="baseUrl"><img class="img-responsive header-logo" src="~/assets/logo/qprob.png" :alt="logoAlt"></a>
                    </div>

                    <div class="collapse navbar-collapse" id="navbar-collapse">
                        <span class="discover"><!-- Discover --></span>
                        <ul class="navbar-nav">
                            <li><a :href="baseUrl + 'today/'"><i class="fa fa-calendar-o" aria-hidden="true"></i> Today</a></li>
                            <li v-for="cat in categories">
                                <a :href="baseUrl+'source/'+cat.slug+'/'" :title="cat.title">
                                <div v-if="cat.thumbnail">
                                    <img class="img-responsive img-circle menuCat" :src="imgBaseUrl+cat.thumbnail" />
                                </div>
                                <div v-else>
                                    <i class="fa fa-list" aria-hidden="true"></i>
                                </div>
                                <i>&nbsp;</i>
                                {{ cat.title }} [{{ cat.post_count }}]</a>
                            </li>
                            <li class="active dropdown"><a data-toggle="dropdown" href="#">
                            <i class="fa fa-building-o" aria-hidden="true"></i>Business</a>
                                <ul class="sub-menu">
                                    <li class="active"><a href="https://bsnssnws.com/" alt="Daily busines news">
                                    <i class="fa fa-building-o" aria-hidden="true"></i>Business news</a></li>
                                    <li><a href="https://entreprnrnws.com/" alt="Daily entrepreneurship news">
                                    <i class="fa fa-google-wallet" aria-hidden="true"></i>Entrepreneurs news</a></li>
                                </ul>
                            </li>
                            <li class="active dropdown"><a data-toggle="dropdown" href="#">
                            <i class="fa fa-signal" aria-hidden="true"></i>Markets</a>
                                <ul class="sub-menu">
                                    <li class="active"><a href="https://stckmrktnws.com/" alt="Daily stock market news">
                                    <i class="fa fa-shopping-cart" aria-hidden="true"></i>Stock market news</a></li>
                                    <li><a href="https://realestenews.com/" alt="Daily real estate investing news">
                                    <i class="fa fa-home" aria-hidden="true"></i>Real estate investing news</a></li>
                                    <li><a href="https://qprob.com/" alt="Daily quantitative trading news">
                                    <i class="fa fa-wrench" aria-hidden="true"></i>Trading news</a></li>
                                    <li><a href="https://quantrade.co.uk/" alt="Free trading signals "><i class="fa fa-signal" aria-hidden="true"></i>Trading signals</a></li>
                                </ul>
                            </li>
                            <li><a href="https://parameterless.com/" alt="Daily technology news"><i class="fa fa-cloud" aria-hidden="true"></i>Technology news</a></li>
                            <li><a href="https://webdnl.com/" alt="Daily insurance news"><i class="fa fa-shield" aria-hidden="true"></i>Insurance news</a></li>
                            <!--
                            <router-link to="/quotes/" tag="li" active-class="active" exact><i class="fa fa-shield" aria-hidden="true"></i>Quotes<router-link/></li>
                            <router-link to="/funnyvideos/" tag="li" active-class="active" exact><i class="fa fa-shield" aria-hidden="true"></i>Funny videos</router-link></li>
                            <router-link to="/celebrities/" tag="li" active-class="active" exact><i class="fa fa-shield" aria-hidden="true"></i>Celebrities</router-link></li>
                            -->
                        </ul>
                    </div>
                </nav>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'headerComp',
  data () {
    return {
      categories: this.categories,
      baseUrl: process.env.baseUrl,
      logoAlt: process.env.logoAlt,
      imgBaseUrl: process.env.imgBaseUrl,
      keyword: process.env.KEYWORD
    }
  },
  methods: {
    fetchData () {
      axios.get('/all_cats/').then(response => {
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