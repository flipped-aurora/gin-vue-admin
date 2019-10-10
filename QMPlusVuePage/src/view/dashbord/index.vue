<template>
  <div>
    <h2 style="text-align:center">
    gin-vue-admin by 奇淼 And Spike
    </h2>
    <h4 style="text-align:center">
      微信：shouzi_1994
    </h4>
    <!-- <el-card class="box-card" shadow="hover">
      <div class="clearfix" slot="header">
        <span>用户信息</span>
      </div>
      <el-row>
        <div class="fl-left left-mg-xs">
          <el-avatar :size="120" :src="userInfo.headerImg" shape="square"></el-avatar>
        </div>
        <div class="fl-left left-mg-lg">
          <div>用户ID：{{userInfo.uuid}}</div>
          <div>用户昵称：{{userInfo.nickName}}</div>
          <div>用户组：{{userInfo.authority&&userInfo.authority.authorityName}}</div>
        </div>
      </el-row>
    </el-card> -->
    <Animition />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Animition from '@/view/dashbord/component/animition.vue'
export default {
  name: 'Dashbord',
  computed: {
    ...mapGetters('user', ['userInfo'])
  },
  components: {
    Animition
  }
}
</script>

<style scoped>
.box-card {
  width: 600px;
}
</style>