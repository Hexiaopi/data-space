import { createRouter, createWebHistory } from 'vue-router'
import NProgress from '@/config/nprogress'
import useMenuStore from '@/store/modules/menu'
import useAclStore from '@/store/modules/acl'
import { staticRouter, errorRouter } from "@/router/modules/staticRouter"
import { initDynamicRouter } from './modules/dynamicRouter'



let router = createRouter({
    history: createWebHistory(),
    routes: [...staticRouter, ...errorRouter],
    strict: false,
    scrollBehavior: () => ({ left: 0, top: 0 }),
})

const WHITE_LIST = ['/login', '/404', '/403']

// 路由守卫
router.beforeEach((to, from, next) => {
    const aclStore = useAclStore()
    const menuStore = useMenuStore()

    NProgress.start()

    // 设置动态标题
    const title = import.meta.env.VITE_APP_TITLE
    document.title = to.meta.title ? `${to.meta.title || ''} - ${title}` : title

    // 登录状态下访问登录页
    if (to.path.toLocaleLowerCase() === '/login') {
        if (aclStore.token) {
            return next(from.fullPath)
        }
        resetRouter()
        return next()
    }

    // 白名单路由直接放行
    if (WHITE_LIST.includes(to.path)) {
        return next()
    }

    // 未登录状态下访问其他页面
    if (!aclStore.token) {
        return next({ path: '/login', replace: true })
    }

    // 动态路由
    if (menuStore.menus.length === 0) {
        initDynamicRouter()
    }

    next()
})

// 重置路由
export const resetRouter = () => {
    const menuStore = useMenuStore()
    menuStore.menus?.forEach(route => {
        const { name } = route
        if (name && router.hasRoute(name)) {
            router.removeRoute(name)
        }
    })
}

// 路由错误处理
router.onError(error => {
    NProgress.done()
    console.warn('路由错误', error.meesage)
})

// 路由跳转结束
router.afterEach(() => {
    NProgress.done()
})

export default router