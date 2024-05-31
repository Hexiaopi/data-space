<template>
    <div class="logo">
        <img src="@/assets/images/logo.png" />
        <span v-show="!appStore.collapsed">
            {{ title }}
        </span>
    </div>
</template>

<script setup lang='ts'>
import { useAppStore } from '@/store/modules/app'
const title = import.meta.env.VITE_APP_TITLE
const appStore = useAppStore()
</script>
<style scoped>
.logo {
    width: 100%;
    height: 64px;
    line-height: 64px;
    padding-left: 20px;
    align-items: center;

    img {
        width: 32px;
        height: 32px;
        vertical-align: middle;
    }

    span {
        margin-left: 10px;

    }
}
</style>