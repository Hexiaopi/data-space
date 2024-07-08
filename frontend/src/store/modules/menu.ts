import { defineStore } from 'pinia'
import { tree } from '@/api/acl'
import { getFlatMenuList } from '@/utils/index'

const modules = import.meta.glob("@/views/**/*.vue")

let useMenuStore = defineStore('menu', {
    state: () => {
        return {
            //menus: globalRouters.filter(item => item.name == 'layout')[0].children,
            menus: [],
        }
    },
    getters: {
        flatMenuList: state => getFlatMenuList(state.menus),
    },
    actions: {
        async getMenuList() {
            const { data } = await tree()
            for (let i = 0; i < data.length; i++) {
                let menu = this.generateRoute(data[i])
                this.menus.push(menu)
            }
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
                        ?.filter((item) => item.type === 'Button')
                        .map((item) => ({ name: item.name })),
                },
                children: item.children?.filter((menu) => menu.type === 'Menu').map((menu) => this.generateRoute(menu)),
            }
        },
        flatMenuTree(tree) {
            return tree.reduce((prev, item) => {
                prev.push(item)
                if (item.children) {
                    prev = prev.concat(this.flatMenuTree(item.children))
                }
                return prev
            }, [])
        },
    }
})

export default useMenuStore