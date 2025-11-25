import { createRouter, createWebHashHistory } from 'vue-router'
import { routes } from './routes.ts'

const router = createRouter({
  history: createWebHashHistory(),
  routes, // `routes: routes` 的缩写
})

export default router
