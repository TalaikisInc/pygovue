<template>
  <div class="col-md-4">
    <ul class="nav flex-column">
      <li class="nav-item" v-for="cat in categories">
        <a class="nav-link" :href="baseUrl + 'category/' + cat.slug + '/'">{{ cat.title }}</a>
      </li>
      <li class="nav-item">
        <strong><a :href="baseUrl + 'categories/1/'" class="nav-link">All categories</a></strong>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'sidebarComponent',
  data () {
    return {
      categories: this.categories,
      baseUrl: process.env.BASE_URL
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