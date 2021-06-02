<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import { checkDB } from '@/api/initdb'
export default {
  name: 'App',
  async created() {
    await checkDB()
  }
}
</script>

<style lang="scss">
// 引入初始化样式
@import '@/style/main.scss';
@import '@/style/base.scss';
@import '@/style/mobile.scss';
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
}
</style>
