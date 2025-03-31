import useMenuStore from "@/store/modules/menu"
import router from "@/router/index"
import { RouteRecordRaw } from "vue-router"

export const initDynamicRouter = async () => {
    const menuStore = useMenuStore()

    try {
        await menuStore.getMenuList()
        menuStore.flatMenuList.forEach(item => {
            router.addRoute('layout', item as unknown as RouteRecordRaw)
        })
    } catch (error) {
        return Promise.reject(error)
    }
}