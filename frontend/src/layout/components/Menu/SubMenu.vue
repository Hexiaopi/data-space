<template>
    <template v-for="menu in menus" :key="menu.path">
        <!-- 目录 -->
        <el-sub-menu v-if="menu.children?.length" :index="menu.path">
            <template #title>
                <el-icon v-if="menu.meta.icon">
                    <component :is="menu.meta.icon"></component>
                </el-icon>
                <span class="sle">{{ menu.meta.title }}</span>
            </template>
            <!-- 递归组件 -->
            <SubMenu :menus="menu.children"></SubMenu>
        </el-sub-menu>
        <!-- 菜单 -->
        <el-menu-item v-else :index="menu.path" @click="handleClickMenu(menu)">
            <el-icon v-if="menu.meta.icon">
                <component :is="menu.meta.icon"></component>
            </el-icon>
            <template #title>
                <span class="sle">{{ menu.meta.title }}</span>
            </template>
        </el-menu-item>
    </template>
</template>

<script setup lang='ts'>
import { useRouter } from 'vue-router'

let props = defineProps(['menus'])
console.log(props.menus)

const router = useRouter()
const handleClickMenu = (menu: any) => {
    if (menu.meta.isLink) {
        return window.open(menu.meta.isLink, '_blank')
    } else {
        router.push(menu.path)
    }
}

</script>
<style lang="scss">
.el-sub-menu .el-sub-menu__title:hover {
    color: var(--el-menu-hover-text-color) !important;
    background-color: transparent !important;
}

.el-menu--collapse {
    .is-active {
        .el-sub-menu__title {
            color: #ffffff !important;
            background-color: var(--el-color-primary) !important;
        }
    }
}

.el-menu-item {
    &:hover {
        color: var(--el-menu-hover-text-color);
    }

    &.is-active {
        color: var(--el-menu-active-color) !important;
        background-color: var(--el-menu-active-bg-color) !important;

        &::before {
            position: absolute;
            top: 0;
            bottom: 0;
            width: 4px;
            content: "";
            background-color: var(--el-color-primary);
        }
    }
}

.vertical,
.classic,
.transverse {
    .el-menu-item {
        &.is-active {
            &::before {
                left: 0;
            }
        }
    }
}

.columns {
    .el-menu-item {
        &.is-active {
            &::before {
                right: 0;
            }
        }
    }
}
</style>