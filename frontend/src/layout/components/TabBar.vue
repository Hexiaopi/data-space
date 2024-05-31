<template>
    <div class="tab-bar">
        <div @click="appStore.switchCollapsed" style="text-align: left;">
            <el-icon>
                <Fold v-if="appStore.collapsed"></Fold>
                <Expand v-else></Expand>
            </el-icon>
        </div>
        <el-breadcrumb :separator-icon="ArrowRight" style="padding:0px 20px;">
            <el-breadcrumb-item v-for="(item, index) in route.matched" :key="index" :to="{ path: item.path }">
                <el-icon>
                    <component :is="item.meta.icon"></component>
                </el-icon>
                {{ item.meta.title }}
            </el-breadcrumb-item>
        </el-breadcrumb>
        <div style="text-align: right">
            <el-dropdown>
                <el-icon style="margin-right: 8px; margin-top: 1px">
                    <setting />
                </el-icon>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item>View</el-dropdown-item>
                        <el-dropdown-item>Add</el-dropdown-item>
                        <el-dropdown-item>Delete</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
            <span>Tom</span>
        </div>
    </div>
</template>

<script setup lang='ts'>
import { useAppStore } from '@/store/modules/app'
import { ArrowRight } from '@element-plus/icons-vue'
import { useRoute } from 'vue-router'
let route = useRoute()
console.log(route.matched)
const appStore = useAppStore()
</script>
<style scoped>
.tab-bar {
    height: 40px;
    display: inline-flex;
    align-items: center;
}
</style>