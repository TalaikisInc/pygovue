import Vue from 'vue'
import moment from 'moment'

Vue.filter('formatDate', function (value) {
  if (value) {
    return moment.utc(value).format('YYYY-MM-DD  hh:mm')
  }
})

Vue.filter('capFirst', function (value) {
  if (value) {
    return value.charAt(0).toUpperCase() + value.slice(1)
  }
})

Vue.filter('truncate', function (value) {
  let length = 300
  let ending = '...'

  if (value.length > length) {
    return value.substring(0, length - ending.length) + ending
  } else {
    return value
  }
})
