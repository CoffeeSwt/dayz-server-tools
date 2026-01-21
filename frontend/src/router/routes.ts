import type { RouteRecordRaw } from 'vue-router'

export const routes: Array<RouteRecordRaw> = [
    { path: '/', redirect: '/layout' },

    {
        path: '/layout', name: 'Layout', component: () => import('@/layout/Index.vue'),
        children: [
            { path: '', name: 'Home', component: () => import('@/views/Home.vue') },
        ]
    },
]