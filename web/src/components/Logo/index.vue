<template>
    <div class="logo" :style="{ width: collapse ? '65px' : '210px' }">
        <img class="logo-img" src="@/assets/images/logo.png" alt="logo" />
        <span v-show="!collapse" class="logo-text">
            {{ title }}
        </span>
    </div>
</template>

<script setup lang='ts'>
defineProps({
    title: {
        type: String,
        default: ''
    },
    collapse: {
        type: Boolean,
        default: false
    }
})

</script>
<style scoped>
.logo {
    box-sizing: border-box;
    height: 55px;
    padding-left: 20px;
    align-items: center;

    .logo-img {
        width: 32px;
        height: 32px;
        margin-right: 6px;
        vertical-align: middle;
    }

    .logo-text {
        font-size: 21.5px;
        color: var(--el-aside-logo-text-color);
        white-space: nowrap;
    }
}
</style>