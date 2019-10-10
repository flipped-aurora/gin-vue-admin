<template>
  <el-card class="box-card" shadow="hover">
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
  </el-card>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'Dashbord',
  computed: {
    ...mapGetters('user', ['userInfo'])
  }
}
</script>

<style scoped>
.box-card{
    width: 600px;
}
</style>