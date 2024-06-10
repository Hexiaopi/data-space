export const globalRouters = [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/login/index.vue'),
    },
    {
        path: '/',
        component: () => import('@/layout/index.vue'),
        name: 'layout',
        children: [
            {
                path: '/home',
                name: 'home',
                component: () => import('@/views/home/index.vue'),
                meta: {
                    title: '首页',
                    hidden: false,
                    icon: 'HomeFilled',
                },
            },
            {
                path: '/system',
                name: 'system',
                meta: {
                    title: '系统管理',
                    hidden: false,
                    icon: 'Setting',
                },
                children: [
                    {
                        path: '/system/department',
                        name: 'department',
                        component: () => import('@/views/system/department/index.vue'),
                        meta: {
                            title: '部门管理',
                            hidden: false,
                            icon: 'School',
                        },
                    },
                    {
                        path: '/system/user',
                        name: 'user',
                        component: () => import('@/views/system/user/index.vue'),
                        meta: {
                            title: '用户管理',
                            hidden: false,
                            icon: 'User',
                        }
                    },
                    {
                        path: '/system/role',
                        name: 'role',
                        component: () => import('@/views/system/role/index.vue'),
                        meta: {
                            title: '角色管理',
                            hidden: false,
                            icon: 'User',
                        }
                    },
                    {
                        path: '/system/menu',
                        name: 'menu',
                        component: () => import('@/views/system/menu/index.vue'),
                        meta: {
                            title: '菜单管理',
                            hidden: false,
                            icon: 'Menu',
                        }
                    },
                ]
            },
        ]
    },
    {
        path: '/404',
        name: '404',
        component: () => import('@/views/error-page/404.vue'),
    },
    {
        path: '/403',
        name: '403',
        component: () => import('@/views/error-page/403.vue'),
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'not-found',
        redirect: '/404',
    },

]