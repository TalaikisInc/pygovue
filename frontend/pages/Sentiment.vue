<template>
<div>
  <header-component></header-component>
    <div class="col-sm-9">
      <div class="row">
        <div class="col-sm-12">
          <div class="tr-content">
            <ad-component></ad-component>

          <div class="tr-section bg-transparent">
            <div class="row">

            <line-chart :data="getS"></line-chart>

            </div>
          </div>
          <ad-component></ad-component>
        </div>
      </div>
    </div>
  </div>
<footer-component></footer-component>
</div>
</template>

<script>
import SentimentChart from '../plugins/chart.js'

import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import Ads from '../components/Ads.vue'

import axios from 'axios'

export default {
  data () {
    return {
      sentiment: [],
      data: null,
      baseUrl: process.env.baseUrl,
      imgBaseUrl: process.env.imgBaseUrl,
      title: process.env.siteName
    }
  },
  asyncData ({ req, params }) {
    return axios.get('/sentiment/')
      .then((response) => {
        return { data: response.data }
      })
      .catch((e) => {
        error({ statusCode: 500, message: e })
      })
  },
  components: {
    'header-component': Header,
    'footer-component': Footer,
    'ad-component': Ads,
    'line-chart': SentimentChart
  },
  head () {
    return {
      title: 'Sentiment | ' + this.title
    }
  },
  computed: {
    getS () {
      var s = []
      var d = []
      for (var i = 0; i < this.data.length; i++) {
        s.push(this.data[i].sentiment)
        d.push(this.data[i].date)
      }

      var avg = s.reduce((a, b) => a + b, 0) / s.length

      return {
        labels: d,
        datasets: [
          {
            label: 'Sentiment',
            backgroundColor: '#f87979',
            data: s.map(x => x - avg)
          }
        ]
      }
    }
  }
}
</script>

<style>
</style>
