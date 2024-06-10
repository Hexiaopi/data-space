import { defineStore } from 'pinia'
import { tree } from '@/api/acl'

const modules = import.meta.glob("@/views/**/*.vue")

let useMenuStore = defineStore('menu', {
    state: () => {
        return {
            //menus: globalRouters.filter(item => item.name == 'layout')[0].children,
            menus: [],
        }
    },
    getters: {

    },
    actions: {
        async getMenuList() {
            const { data } = await tree()
            for (let i = 0; i < data.length; i++) {
                let menu = this.generateRoute(data[i])
                this.menus.push(menu)
            }
            console.log(this.menus)
        },
        generateRoute(item) {
            return {
                name: item.name,
                path: item.path,
                component: modules["/src/views" + item.component + ".vue"],
                meta: {
                    icon: item.icon,
                    title: item.desc,
                    layout: item.layout,
                    btns: item.children
                        ?.filter((item) => item.type === 'BUTTON')
                        .map((item) => ({ name: item.name })),
                },
                children: item.children?.map((item) => this.generateRoute(item)),
            }
        },
    }
})

export default useMenuStore