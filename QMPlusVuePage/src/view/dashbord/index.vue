<template>
  <div>
    <h2 style="text-align:center">gin-vue-admin by 奇淼 And krank666</h2>
    <h2 style="text-align:center">点击动画获取联系方式</h2>

    <el-drawer
      :before-close="handleClose"
      :direction="direction"
      :visible.sync="drawer"
      title="联系方式"
    >
      <div style="text-align:center;margin-top:30px;">
        <img src="http://qmplusimg.henrongyi.top/yx.jpg" width="180" />
        <H3>krank666微信</H3>
      </div>
      <div style="text-align:center;margin-top:30px;">
        <img src="http://qmplusimg.henrongyi.top/jjz.jpg" width="180" />
        <H3>Mr.奇淼微信</H3>
      </div>
    </el-drawer>
    <div @click="drawer = true">
      <Animition />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Animition from '@/view/dashbord/component/animition.vue'
export default {
  name: 'Dashbord',
  data() {
    return {
      drawer: false
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  components: {
    Animition
  }
}
</script>

<style lang="scss" scoped>
</style>