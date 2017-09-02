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
