import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '../views/Home.vue'
import Catalog from '../views/Catalog.vue'
import TourDetail from '../views/TourDetail.vue'

import Register from '../views/auth/Register.vue'
import Login from '../views/auth/Login.vue'
import Logout from '../views/auth/Logout.vue'

import AdminIndex from '../views/admin/Index.vue'
import AdminTours from '../views/admin/tour/Tours.vue'
import AdminUpdateTour from '../views/admin/tour/UpdateTour.vue'
import AdminCreateTour from '../views/admin/tour/CreateTour.vue'
import AdminDeleteTour from '../views/admin/tour/DeleteTour.vue'


Vue.use(VueRouter)

const routes = [

  // Site
  {
    path: '/',
    name: 'home',
    component: Home,
    meta: {
      layout: 'default-layout'
    }
  },
  {
    path: '/tours',
    name: 'catalog',
    component: Catalog,
    meta: {
      layout: 'default-layout'
    }
  },
  {
    path: '/tour/:slug',
    name: 'tourDetail',
    component: TourDetail,
    props: true,
    meta: {
      layout: 'default-layout'
    }
  },


  // Auth
  {
    path: '/register',
    name: 'register',
    component: Register,
    meta: {
      requiresVisitor: true,
    }
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      requiresVisitor: true,
      layout: 'default-layout'
    }
  },
  {
    path: '/logout',
    name: 'logout',
    component: Logout
  },



  // Admin
  {
    path: '/admin',
    name: 'admin',
    component: AdminIndex,
    meta: {
      requiresAuth: true,
      layout: 'admin-layout'
    }
  },

  // Admin Tour
  {
    path: '/admin/tours',
    name: 'adminGetTours',
    component: AdminTours,
    meta: {
      requiresAuth: true,
      requiresOwnerOrAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/tour',
    name: 'adminCreateTour',
    component: AdminCreateTour,
    meta: {
      requiresAuth: true,
      requiresOwnerOrAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/tour/:slug',
    name: 'adminUpdateTour',
    component: AdminUpdateTour,
    props: true,
    meta: {
      requiresAuth: true,
      requiresOwnerOrAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/tour/delete/:slug',
    name: 'adminDeleteTour',
    component: AdminDeleteTour,
    props: true,
    meta: {
      requiresAuth: true,
      requiresOwnerOrAdmin: true,
      layout: 'admin-layout'
    }
  },


]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
