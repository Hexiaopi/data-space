<template>
  <el-config-provider :locale="zhCn">
    <router-view />
  </el-config-provider>
</template>

<script setup lang='ts'>
import zhCn from "element-plus/es/locale/lang/zh-cn"
import layout from '@/layout/index.vue'
</script>
<style scoped></style>