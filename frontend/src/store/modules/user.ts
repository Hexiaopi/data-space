import { defineStore } from "pinia"
import { info } from "@/api/user"

let useUserStore = defineStore("user", {
    state: () => {
        return {
            userInfo: JSON.parse(localStorage.getItem("userInfo") || "{}")
        }
    },
    getters: {
        id: state => state.userInfo?.id,
        name: state => state.userInfo?.name,
        avatar: state => state.userInfo?.avatar,
        roles: state => state.userInfo?.roles,
        role: state => state.userInfo?.current_role,
    },
    actions: {
        async getUserInfo() {
            let result = await info()
            this.userInfo = result.data
            localStorage.setItem("userInfo", JSON.stringify(result.data))
        }
    }
})

export default useUserStore