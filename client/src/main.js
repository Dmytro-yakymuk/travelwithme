import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from "./store";
// import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import "./constants/axios";

Vue.config.productionTip = false

import DefaultLayout from './views/layouts/DefaultLayout.vue'
import AdminLayout from './views/layouts/AdminLayout'

Vue.component('default-layout', DefaultLayout)
Vue.component('admin-layout', AdminLayout)

// Specify vue plugins
// Vue.use(IconsPlugin);
// Vue.use(BootstrapVue);


// scroll
var VueScrollTo = require('vue-scrollto');
Vue.use(VueScrollTo)


router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!store.getters.loggedIn) {
      next({
        name: 'login',
      })
    } else {
      next()
    }
  } 
  if (to.matched.some(record => record.meta.requiresVisitor)) {
    if (store.getters.loggedIn) {
      next({
        // якщо в нас прописаний мета true на цей роут то ми перенаправляєм по нижче уназаному роуту
        name: 'admin',
      })
    } else {
      next()
    }
  } 
  if (to.matched.some(record => record.meta.requiresOwnerOrAdmin)) {
    if (localStorage.getItem('role') != "owner" && localStorage.getItem('role') != "admin") {
      next({
        name: 'admin',
      })
    } else {
      next()
    } 
  } 
  if (to.matched.some(record => record.meta.requiresAdmin)) {
    if (localStorage.getItem('role') != "admin") {
      next({
        name: 'admin',
      })
    } else {
      next()
    } 
  } else {
    next()
  }
})


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
