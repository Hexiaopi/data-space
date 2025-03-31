<template>
    <el-container class="layout-container">
        <div>
            <Logo :collapse="appStore.collapsed" :title="title"></Logo>
            <el-scrollbar>
                <!-- <SideMenu :menus="menuStore.menus"></SideMenu> -->
                <el-menu :collapse="appStore.collapsed">
                    <SubMenu :menus="getShowMenus(menuStore.menus)"></SubMenu>
                </el-menu>
            </el-scrollbar>
        </div>
        <el-container>
            <el-header>
                <Header></Header>
            </el-header>
            <el-main>
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<script lang="ts" setup>
import useMenuStore from '@/store/modules/menu'
import useAppStore from '@/store/modules/app'
import SubMenu from "@/layout/components/Menu/SubMenu.vue"
import Logo from "@/components/Logo/index.vue"
import Header from "@/layout/components/Header/index.vue"
const title = import.meta.env.VITE_APP_TITLE

const menuStore = useMenuStore()
const appStore = useAppStore()

function getShowMenus(menus: any) {
    let newMenuList = JSON.parse(JSON.stringify(menus));
    return newMenuList.filter(item => {
        item.children?.length && (item.children = getShowMenus(item.children));
        return !item.meta?.hidden;
    });
}

</script>

<style scoped>
.layout-container {
    height: 100%;
}
</style>