<template>
<div class="row align-items-center">
  <div class="col-sm-3"></div>
  <div class="col-sm-6">
    <ul class="pagination" v-if="paginatorType === 0">
      <li v-for="n in paginationRange" :class="activePage(n)">
        <a :href="baseUrl + 'page/' + n + '/'" @click="pageChanged(n)" class="page-link">{{ n }}</a>
      </li>
    </ul>
    <ul class="pagination" v-if="paginatorType === 1">
      <li v-for="n in paginationRange" :class="activePage(n)">
        <a :href="baseUrl + keyword + '/' + value + '/page/' + n + '/'" @click="pageChanged(n)" class="page-link">{{ n }}</a>
      </li>
    </ul>
    <ul class="pagination" v-if="paginatorType === 2">
      <li v-for="n in paginationRange" :class="activePage(n)">
        <a :href="baseUrl + catKey + '/' + n + '/'" @click="pageChanged(n)" class="page-link">{{ n }}</a>
      </li>
    </ul>
    <ul class="pagination" v-if="paginatorType === 3">
      <li v-for="n in paginationRange" :class="activePage(n)">
        <a :href="baseUrl + searchKey + '/' + value + '/page/' + n + '/'" @click="pageChanged(n)" class="page-link">{{ n }}</a>
      </li>
    </ul>
  </div>
  <div class="col-sm-3"></div>
</div>
</template>

<script>
import Util from '../plugins/util.js'

export default {
  name: 'paginatorComponent',
  data () {
    return {
      baseUrl: process.env.BASE_URL,
      keyword: process.env.KEYWORD,
      catKey: process.env.CATEGORIES_KEY,
      searchKey: process.env.SEARCH_KEYWORD
    }
  },
  methods: {
    activePage (pageNum) {
      return this.currentPage === pageNum ? 'page-item active' : 'page-item'
    },
    pageChanged (pageNum) {
      this.$emit('page-changed', pageNum)
    }
  },
  computed: {
    lastPage () {
      if (this.totalPages) {
        return this.totalPages
      } else {
        return this.totalItems % this.itemsPerPage === 0
          ? this.totalItems / this.itemsPerPage
          : Math.floor(this.totalItems / this.itemsPerPage)
      }
    },
    paginationRange () {
      let start =
        this.currentPage - this.visiblePages / 2 <= 0
          ? 1 : this.currentPage + this.visiblePages / 2 > this.lastPage
            ? Util.lowerBound(this.lastPage - this.visiblePages + 1, 1)
            : Math.ceil(this.currentPage - this.visiblePages / 2)

      let range = []

      for (let i = 0; i < this.visiblePages && i < this.lastPage; i++) {
        range.push(start + i)
      }

      return range
    }
  },
  props: {
    currentPage: {
      type: Number,
      required: true
    },
    value: {
      type: String,
      required: false
    },
    paginatorType: {
      type: Number,
      required: true
    },
    totalPages: Number,
    itemsPerPage: Number,
    totalItems: Number,
    visiblePages: {
      type: Number,
      default: 20,
      coerce: (val) => parseInt(val)
    }
  }
}
</script>

<style></style>