<template>
    <el-dropdown>
        <el-row :gutter="2" style="align-items: center;display: flex">
            <el-col :span="12">
                <el-avatar :src="userStore.avatar" />
            </el-col>
            <el-col :span="12">
                <el-row>
                    {{ userStore.name }}
                </el-row>
                <el-row style="white-space: nowrap;margin-top: 5px">
                    [{{ userStore.role?.name }}]
                </el-row>
            </el-col>
        </el-row>
        <template #dropdown>
            <el-dropdown-menu>
                <el-dropdown-item>个人信息</el-dropdown-item>
                <el-dropdown-item>退出登录</el-dropdown-item>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
</template>

<script setup lang='ts'>
import useUserStore from '@/store/modules/user'
const userStore = useUserStore()
</script>

<style scoped>
.avatar {
    width: 40px;
    height: 40px;
    overflow: hidden;
    cursor: pointer;
    border-radius: 50%;

    img {
        width: 100%;
        height: 100%;
    }
}
</style>