import { RouteRecordRaw } from "vue-router"

export const staticRouter: RouteRecordRaw[] = [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/login/index.vue'),
    },
    {
        path: '/',
        name: 'layout',
        component: () => import('@/layout/index.vue'),
        children: []
    }
]

export const errorRouter = [
    {
        path: '/404',
        name: '404',
        component: () => import('@/views/error-page/404.vue'),
        meta: {
            title: '404',
            hidden: true,
            icon: '',
        }
    },
    {
        path: '/403',
        name: '403',
        component: () => import('@/views/error-page/403.vue'),
        meta: {
            title: '403',
            hidden: true,
            icon: '',
        }
    },
    {
        path: '/:pathMatch(.*)*',
        redirect: '/404',
    },
]