<template>
<div>
  <ul class="pagination" v-if="source === 0">
    <li v-for="p in pages" :class="{ active: active === p }"><a :href="'/page/' + p + '/'">{{ p }}</a></li>
  </ul>
  <ul class="pagination" v-if="source === 1">
    <li v-for="p in pages" :class="{ active: active === p }"><a :href="'/category/' + value + '/page/' + p + '/'">{{ p }}</a></li>
  </ul>
/div>
</template>

<script>
export default {
  name: 'paginatorComp',
  props: {
    pages: {
      type: Number,
      required: true
    },
    source: {
      type: Number
    },
    value: {
      type: String
    },
    active: {
      type: Number,
      required: true
    }
  }
}
</script>

<style></style>