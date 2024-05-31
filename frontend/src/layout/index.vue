<template>
    <el-container class="layout-container">
        <div>
            <SideLogo></SideLogo>
            <SideMenu :menus="menuStore.menus"></SideMenu>
        </div>

        <el-container>
            <el-header>
                <TabBar></TabBar>
            </el-header>

            <el-main>
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<script lang="ts" setup>
import { SideLogo, SideMenu, TabBar } from '@/layout/components'
import useMenuStore from '@/store/modules/menu'

const menuStore = useMenuStore()

</script>

<style scoped>
.layout-container {
    height: 100%;
}
</style>