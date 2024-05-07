<template>
  <div id="app" class="bg-gray-50 text-slate-700 dark:text-slate-500  dark:bg-slate-800">
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
// element 2.3.8之前使用下面的语句
// import zhCn from 'element-plus/lib/locale/lang/zh-cn'

defineOptions({
  name: 'App'
})

</script>
<style lang="scss">

// 引入初始化样式
#app {
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}
.el-button{
  font-weight: 400 !important;
}

</style>
