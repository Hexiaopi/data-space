<template>
    <div class="tab-bar">
        <div class="tab-left">
            <div @click="appStore.switchCollapsed" style="text-align: left;">
                <el-icon>
                    <Fold v-if="appStore.collapsed"></Fold>
                    <Expand v-else></Expand>
                </el-icon>
            </div>
            <el-breadcrumb :separator-icon="ArrowRight" style="padding:0px 20px;">
                <el-breadcrumb-item v-for="(item, index) in breadcrumbList" :key="index" :to="{ path: item.path }">
                    <el-icon>
                        <component :is="item.meta.icon"></component>
                    </el-icon>
                    {{ item.meta.title }}
                </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="tab-right">
            <Avatar />
        </div>
    </div>
</template>

<script setup lang='ts'>
import { useAppStore } from '@/store/modules/app'
import { ArrowRight } from '@element-plus/icons-vue'
import { useRoute } from 'vue-router'
import Avatar from './components/avatar.vue'
import { computed } from 'vue';

let route = useRoute()
const appStore = useAppStore()
const breadcrumbList = computed(() => {
    // 过滤掉meta为空的路由，并且meta.title不为空
    return route.matched.filter(item => item.meta && item.meta.title)
})
</script>

<style scoped>
.tab-bar {
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .tab-left {
        display: flex;
        align-items: center;
        justify-content: flex-start;
        align-items: center;
    }

    .tab-right {
        display: flex;
        align-items: center;
        justify-content: flex-end;
        align-items: center;
    }
}
</style>