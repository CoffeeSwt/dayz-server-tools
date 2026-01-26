import type { RouteRecordRaw } from 'vue-router'

export const routes: Array<RouteRecordRaw> = [
    { path: '/', redirect: '/layout' },

    {
        path: '/layout', name: 'Layout', component: () => import('@/layout/Index.vue'),
        children: [
            { path: '', name: 'Server', component: () => import('@/views/Server.vue') },
            { path: 'map', name: 'Map', component: () => import('@/views/Map.vue') },
        ]
    },
]