import type { RouteRecordRaw } from 'vue-router'

export const routes: Array<RouteRecordRaw> = [
    { path: '/', redirect: '/home' },

    { path: '/home', name: 'Home', component: () => import('@/views/Home.vue') },
]