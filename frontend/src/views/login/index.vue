<template>
    <div class="login-container">
        <el-form :rules="rules" :model="loginForm" ref="loginFormRef" class="login-form">
            <el-row>
                <el-col :span="14">
                    <img src="@/assets/images/login_banner.webp" height="300">
                </el-col>
                <el-col :span="10">
                    <h1>
                        <Logo :title="title"></Logo>
                    </h1>
                    <el-form-item prop="username">
                        <el-input placeholder=" 请输入用户名" :prefix-icon="User" v-model="loginForm.username"></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input placeholder="请输入密码" type="password" :prefix-icon="Lock" v-model="loginForm.password"
                            show-password></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button :loading="loading" type="primary" style="width: 100%;"
                            @click="handler">登录</el-button>
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
    </div>
</template>

<script setup lang='ts'>
import { User, Lock } from '@element-plus/icons-vue'
import { ElNotification, FormRules } from 'element-plus'
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import useAclStore from '@/store/modules/acl'
import useUserStore from '@/store/modules/user'
import { GetCNTime } from '@/utils/time'
import Logo from '@/components/Logo/index.vue'
const title = import.meta.env.VITE_APP_TITLE

let router = useRouter()
let aclStore = useAclStore()
let userStore = useUserStore()

let loading = ref(false)
let loginForm = reactive({
    username: 'admin',
    password: '123456'
})
let loginFormRef = ref()

const rules = reactive<FormRules>({
    username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 3, max: 10, message: '用户名长度在3到10个字符', trigger: 'change' }
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 15, message: '密码长度在6到15个字符', trigger: 'change' }
    ]
})
const handler = () => {
    loginFormRef.value.validate((valid: any) => {
        if (valid) {
            login()
        } else {
            ElNotification({
                title: '提示',
                message: '请输入正确的用户名和密码',
                type: 'warning',
                duration: 3000
            })
        }
    })
}
const login = () => {
    loading.value = true
    aclStore.userLogin(loginForm).then(() => {

        //获取用户信息
        userStore.getUserInfo()

        router.push('/')
        ElNotification({
            title: '登录成功',
            message: `${GetCNTime()}好，欢迎回来！`,
            type: 'success',
            duration: 3000
        })
        loading.value = false
    }).catch((err) => {
        ElNotification({
            title: '登录失败',
            message: (err as Error).message,
            type: 'error',
            duration: 3000
        })
        loading.value = false
    })
}
</script>
<style scoped>
.login-container {
    background: url('@/assets/images/login_bg.webp') no-repeat;
    background-size: cover;
    width: 100%;
    height: 100vh;

    .login-form {
        position: relative;
        top: 30vh;
        left: 50vh;
        padding: 20px;
        background: #fff;
        border-radius: 5px;
        box-shadow: 0 0 10px #ccc;
        max-width: 600px;
    }
}
</style>