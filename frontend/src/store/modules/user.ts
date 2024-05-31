import { defineStore } from 'pinia'
import { login } from '@/api/user'
import type { LoginParams } from '@/api/user/type'

let useUserStore = defineStore('user', {
    state: () => {
        return {
            token: localStorage.getItem('token'),
        }
    },
    actions: {
        async userLogin(params: LoginParams) {
            let result = await login(params)
            console.log(result)
            if (result.code === '000000') {
                this.token = result.data.access_token
                localStorage.setItem('token', result.data.access_token)
                return 'ok'
            } else {
                return Promise.reject(new Error(result.desc))
            }
        }
    },
    getters: {

    }
})

export default useUserStore