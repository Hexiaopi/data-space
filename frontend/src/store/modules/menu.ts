import { defineStore } from 'pinia'

import { globalRouters } from '@/router/routers'


let useMenuStore = defineStore('menu', {
    state: () => {
        return {
            menus: globalRouters,
        }
    },
    actions: {

    },
    getters: {

    }
})

export default useMenuStore