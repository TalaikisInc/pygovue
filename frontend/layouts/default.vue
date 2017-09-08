<template>
  <div class="container">
    <header class="app-header navbar">
      <header-component></header-component>
    </header>
    <div class="app-body">
      <div class="sidebar">
        <sidebar-component></sidebar-component>
      </div>
      <main class="main">
        <nuxt/>
      </main>
      <!-- <aside class="aside-menu">
      </aside> -->
    </div>
    <footer class="app-footer">
      <footer-component></footer-component>
    </footer>
  </div>
</template>

<script>
import Header from '../components/Header.vue'
import Sidebar from '../components/Sidebar.vue'
import Footer from '../components/Footer.vue'

export default {
  data () {
    return { }
  },
  components: {
    'header-component': Header,
    'sidebar-component': Sidebar,
    'footer-component': Footer
  }
}
</script>

<style>
</style>
