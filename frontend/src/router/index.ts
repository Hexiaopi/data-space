import { createRouter, createWebHistory } from 'vue-router'
import { globalRouters } from './routers'

let router = createRouter({
    history: createWebHistory(),
    routes: globalRouters,
    scrollBehavior: () => ({ left: 0, top: 0 }),
})

export default router