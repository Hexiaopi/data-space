<template>
    <div>
        <el-scrollbar>
            <el-menu :collapse="appStore.collapsed" class="el-menu-vertical-demo">
                <template v-for="menu in menus" :key="menu.name">
                    <template v-if="!menu.meta.hidden">
                        <el-sub-menu v-if="menu.children" :index="menu.path">
                            <template #title>
                                <el-icon>
                                    <component :is="menu.meta.icon"></component>
                                </el-icon>
                                <span>{{ menu.meta.title }}</span>
                            </template>
                            <SideMenu :menus="menu.children"></SideMenu>
                        </el-sub-menu>
                        <el-menu-item v-else :index="menu.path" @click="goRoute(menu.path)">
                            <el-icon>
                                <component :is="menu.meta.icon"></component>
                            </el-icon>
                            <template #title>
                                <span>{{ menu.meta.title }}</span>
                            </template>
                        </el-menu-item>
                    </template>
                </template>
            </el-menu>
        </el-scrollbar>
    </div>
</template>

<script setup lang='ts'>

import { defineProps } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/store/modules/app'
const router = useRouter()
defineProps(['menus'])
const goRoute = (path: string) => {
    router.push(path)
}
const appStore = useAppStore()
</script>
<script lang="ts">
export default {
    name: 'SideMenu'
}
</script>
<style scoped>
.scrollbar {
    width: 100%;
    height: 100%;

    .el-menu {
        border-right: none;
    }
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
}
</style>