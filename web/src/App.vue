<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import { checkDB } from "@/api/initdb"
export default {
  name: 'app',
  async created(){
    const res = await checkDB()
    if(res.code == 0 && res.data.needInit){
      this.$message({
        type:"info",
        message:"您是第一次使用，请初始化"
      })
        this.$store.commit("user/NeedInit")
        this.$router.push({name:"init"})
    }
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
