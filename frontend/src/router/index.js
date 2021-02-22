import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../pages/auth/Login.vue'
import {Token} from '../utils/token'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/users',
    name: "Users",
    component: () => import("../pages/users/Users.vue"),
    beforeEnter:(to, from, next) => {
      if (!Token.getToken()) {
        return next("/")
      }

      next()

    }
  }
  // {
  //   path: '/about',
  //   name: 'About',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  // }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
