<template>
  <div id="app" class="bg-gray-100 h-screen overflow-hidden font-sans">
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
// element 2.3.8之前使用下面的语句
// import zhCn from 'element-plus/lib/locale/lang/zh-cn'
</script>
<style lang="scss">
@tailwind base;
@tailwind components;
@tailwind utilities;
// 引入初始化样式

</style>
