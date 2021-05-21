import Vue from 'vue'
import VueRouter from 'vue-router'

import Home from '../views/Home.vue'
import Catalog from '../views/Catalog.vue'
import TourDetail from '../views/TourDetail.vue'

import Overview from '../components/tour/Overview.vue'
import Map from '../components/tour/Map.vue'
import Comments from '../components/tour/Comments.vue'
import dateAndPrice from '../components/tour/DateAndPrice.vue'

import Cart from '../views/Cart.vue'
import Checkout from '../views/Checkout.vue'

import Register from '../views/auth/Register.vue'
import Login from '../views/auth/Login.vue'
import Logout from '../views/auth/Logout.vue'

import AdminIndex from '../views/admin/Index.vue'

import AdminCategories from '../views/admin/category/Categories.vue'
import AdminUpdateCategory from '../views/admin/category/UpdateCategory.vue'
import AdminCreateCategory from '../views/admin/category/CreateCategory.vue'

import AdminRegions from '../views/admin/region/Regions.vue'
import AdminUpdateRegion from '../views/admin/region/UpdateRegion.vue'
import AdminCreateRegion from '../views/admin/region/CreateRegion.vue'

import AdminTours from '../views/admin/tour/Tours.vue'
import AdminUpdateTour from '../views/admin/tour/UpdateTour.vue'
import AdminCreateTour from '../views/admin/tour/CreateTour.vue'
import AdminDeleteTour from '../views/admin/tour/DeleteTour.vue'

import AdminEvents from '../views/admin/event/Events.vue'
import AdminUpdateEvent from '../views/admin/event/UpdateEvent.vue'
import AdminCreateEvent from '../views/admin/event/CreateEvent.vue'

import AdminTrips from '../views/admin/trip/Trips.vue'
import AdminUpdateTrip from '../views/admin/trip/UpdateTrip.vue'
import AdminCreateTrip from '../views/admin/trip/CreateTrip.vue'

import AdminOrders from '../views/admin/order/Orders.vue'
import AdminViewOrder from '../views/admin/order/ViewOrder.vue'

import AdminComments from '../views/admin/comment/Comments.vue'
import AdminViewComment from '../views/admin/comment/ViewComment.vue'
import AdminCreateComment from '../views/admin/comment/CreateComment.vue'
import AdminUpdateComment from '../views/admin/comment/UpdateComment.vue'

import AdminUpdateProfile from '../views/admin/profile/UpdateProfile.vue'




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
    children: [
      {
        path: '',
        name: 'overview',
        component: Overview
      },
      {
        path: 'map',
        name: 'map',
        component: Map
      },
      {
        path: 'comments',
        name: 'comments',
        component: Comments
      },
      {
        path: 'dateAndPrice',
        name: 'dateAndPrice',
        component: dateAndPrice
      },
      
    ],
    props: true,
    meta: {
      layout: 'default-layout'
    }
  },
  {
    path: '/cart',
    name: 'cart',
    component: Cart,
    meta: {
      requiresAuth: true,
      requiresClient: true,
      layout: 'default-layout'
    }
  },
  {
    path: '/checkout',
    name: 'checkout',
    component: Checkout,
    meta: {
      requiresAuth: true,
      requiresClient: true,
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

  // Admin Category
  {
    path: '/admin/categories',
    name: 'adminGetCategories',
    component: AdminCategories,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/category',
    name: 'adminCreateCategory',
    component: AdminCreateCategory,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/category/:slug',
    name: 'adminUpdateCategory',
    component: AdminUpdateCategory,
    props: true,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },


  // Admin Region
  {
    path: '/admin/regions',
    name: 'adminGetRegions',
    component: AdminRegions,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/region',
    name: 'adminCreateRegion',
    component: AdminCreateRegion,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/region/:slug',
    name: 'adminUpdateRegion',
    component: AdminUpdateRegion,
    props: true,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
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


  // Admin Event
  {
    path: '/admin/events',
    name: 'adminGetEvents',
    component: AdminEvents,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/event',
    name: 'adminCreateEvent',
    component: AdminCreateEvent,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/event/:id',
    name: 'adminUpdateEvent',
    component: AdminUpdateEvent,
    props: true,
    meta: {
      requiresAuth: true,
      requiresAdmin: true,
      layout: 'admin-layout'
    }
  },


  // Admin Trip
  {
    path: '/admin/trips',
    name: 'adminGetTrips',
    component: AdminTrips,
    meta: {
      requiresAuth: true,
      requiresOwnerOrAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/trip',
    name: 'adminCreateTrip',
    component: AdminCreateTrip,
    meta: {
      requiresAuth: true,
      requiresOwnerOrAdmin: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/trip/:id',
    name: 'adminUpdateTrip',
    component: AdminUpdateTrip,
    props: true,
    meta: {
      requiresAuth: true,
      requiresOwnerOrAdmin: true,
      layout: 'admin-layout'
    }
  },

  // Admin Order
  {
    path: '/admin/orders',
    name: 'adminGetOrders',
    component: AdminOrders,
    meta: {
      requiresAuth: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/order/:id',
    name: 'adminViewOrder',
    component: AdminViewOrder,
    props: true,
    meta: {
      requiresAuth: true,
      layout: 'admin-layout'
    }
  },

  // Admin Comment
  {
    path: '/admin/comments',
    name: 'adminGetComments',
    component: AdminComments,
    meta: {
      requiresAuth: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/comment/view/:id',
    name: 'adminViewComment',
    component: AdminViewComment,
    props: true,
    meta: {
      requiresAuth: true,
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/comment',
    name: 'adminCreateComment',
    component: AdminCreateComment,
    meta: {
      requiresAuth: true,
      requiresClient: true, 
      layout: 'admin-layout'
    }
  },
  {
    path: '/admin/comment/edit/:id',
    name: 'adminUpdateComment',
    component: AdminUpdateComment,
    props: true,
    meta: {
      requiresAuth: true,
      requiresClient: true,
      layout: 'admin-layout'
    }
  },

  // Edit Profile
  {
    path: '/admin/editProfile',
    name: 'adminEditProfile',
    component: AdminUpdateProfile,
    meta: {
      requiresAuth: true,
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
