import axios from 'axios'
import { ElMessage } from 'element-plus'
import useAclStore from '@/store/modules/acl'

const request = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_API,
    timeout: 5000
})

request.interceptors.request.use(
    (config) => {
        const { token } = useAclStore()
        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config;
    },
    (error) => {
        return Promise.reject(error)
    }
)

request.interceptors.response.use(
    (response) => {
        let resp = response.data
        if (resp.code !== '000000') {
            ElMessage({ message: resp.desc, type: 'error' })
            if (resp.code === '000010') {
                // 跳转到登录页面
                const aclStore = useAclStore()
                aclStore.removeToken()
                window.location.href = '/login'
            }
            return Promise.reject(resp)
        }
        return response.data;
    },
    (error) => {
        let message = ''
        const status = error.response.status
        switch (status) {
            case 400:
                message = '请求错误'
                break;
            case 401:
                message = '未授权，请登录'
                break;
            case 403:
                message = '拒绝访问'
                break;
            case 404:
                message = `请求地址出错: ${error.response.config.url}`
                break;
            case 500:
                message = '服务器内部错误'
                break;
            default:
                message = `连接错误${error.message}`
                break;
        }
        ElMessage({ message: message, type: 'error' })
        return Promise.reject(error)
    }
)

export default request