require('dotenv').config({ path: '../.env' })

const cat = '/' + process.env.KEYWORD + '/:catSlug/'
const catPaged = '/' + process.env.KEYWORD + '/:catSlug/page/:page/'
const cats = '/' + process.env.CATEGORIES_KEY + '/:page/'
const search = '/' + process.env.SEARCH_KEYWORD + '/:catSlug/'
const searchPaged = '/' + process.env.SEARCH_KEYWORD + '/:catSlug/page/:page/'

module.exports = {
  head: {
    title: 'PyGoVue',
    manifest: {
      name: process.env.SITE_NAME,
      short_name: 'PyGoVue',
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' }
    ],
    script: [
      { src: 'https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js' },
      { src: '/js/vue-social-sharing.min.js' },
      { src: 'https://code.jquery.com/jquery-3.1.1.slim.min.js' },
      { src: 'https://cdnjs.cloudflare.com/ajax/libs/tether/1.4.0/js/tether.min.js' },
      { src: 'https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: 'stylesheet', href: 'https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css' },
      { rel: 'stylesheet', href: 'https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css' }
    ]
  },
  css: [
    '@/assets/css/main.css'
  ],
  plugins: [
    '~plugins/filters.js',
    '~plugins/axios.js'
  ],
  modules: [
    ['@nuxtjs/component-cache', { maxAge: 1000 * 60 * 60 * 24 }],
    ['@nuxtjs/google-analytics', { ua: process.env.GOOGLE_ANALYTICS }],
    ['@nuxtjs/google-tag-manager', { id: process.env.GTM }],
    '@nuxtjs/icon',
    '@nuxtjs/meta',
    '@nuxtjs/workbox',
    '@nuxtjs/manifest'
  ],
  env: {
    API_URL: process.env.API_URL,
    BASE_URL: process.env.BASE_URL,
    IMG_URL: process.env.IMG_URL,
    SITE_NAME: process.env.SITE_NAME,
    ADSENSE_AD_CLIENT: process.env.ADSENSE_AD_CLIENT,
    ADSENSE_AD_SLOT: process.env.ADSENSE_AD_SLOT,
    TWITTER_HANDLE: process.env.TWITTER_HANDLE,
    FACEBOOK_HANDLE: process.env.FACEBOOK_HANDLE,
    KEYWORD: process.env.KEYWORD,
    CATEGORIES_KEY: process.env.CATEGORIES_KEY
  },
  router: {
    extendRoutes (routes, resolve) {
      routes.push(
        { path: '', component: resolve(__dirname, 'pages', 'Posts.vue') },
        { path: '/page/:page/', component: resolve(__dirname, 'pages', 'Posts.vue') },
        { path: '/:postSlug/', component: resolve(__dirname, 'pages', 'Post.vue') },
        { path: cat, component: resolve(__dirname, 'pages', 'PostsByCat.vue') },
        { path: catPaged, component: resolve(__dirname, 'pages', 'PostsByCat.vue') },
        { path: cats, component: resolve(__dirname, 'pages', 'Categories.vue') },
        { path: search, component: resolve(__dirname, 'pages', 'Searchable.vue') },
        { path: searchPaged, component: resolve(__dirname, 'pages', 'Searchable.vue') },
        { path: '*', component: resolve(__dirname, 'pages', 'Error404.vue') }
      )
    }
  },
  loading: '~/components/Loading.vue',
  build: {
    analyze: true,
    extractCSS: true,
    vendor: ['axios', 'moment'],
    extend (config, ctx) {
      if (ctx.dev && ctx.isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    }
  }
}
