import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Catalog from '../views/Catalog.vue'
import About from '../views/About.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/catalog', component: Catalog },
  { path: '/about', component: About }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router