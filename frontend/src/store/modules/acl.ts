import { defineStore } from 'pinia'
import { login } from '@/api/acl'
import type { LoginParams } from '@/api/acl/type'

let useAclStore = defineStore('acl', {
    state: () => {
        return {
            token: localStorage.getItem('token'),
        }
    },
    actions: {
        async userLogin(params: LoginParams) {
            let result = await login(params)
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

export default useAclStore