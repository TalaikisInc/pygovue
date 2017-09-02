import Vue from 'vue'
import VueLocalStorage from 'vue-ls'

const options = {
  namespace: 'vuejs__'
}
  
Vue.use(VueLocalStorage, options)
