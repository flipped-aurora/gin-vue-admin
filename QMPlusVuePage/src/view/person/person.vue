<template>
  <div>
    <div class="fl-left left-mg-xs">
      <el-upload
        :headers="{'x-token':token}"
        :on-success="handleAvatarSuccess"
        :show-file-list="false"
        :action="`${path}/user/uploadHeaderImg`"
        class="avatar-uploader"
        name="headerImg"
      >
        <img :src="userInfo.headerImg" class="avatar" v-if="userInfo.headerImg" />
        <i class="el-icon-plus avatar-uploader-icon" v-else></i>
      </el-upload>

      <!-- <el-avatar :size="120" :src="userInfo.headerImg" shape="square"></el-avatar> -->
    </div>
    <div class="fl-left left-mg-lg">
      <div>用户ID：{{userInfo.uuid}}</div>
      <div>用户昵称：{{userInfo.nickName}}</div>
      <div>用户组：{{userInfo.authority&&userInfo.authority.authorityName}}</div>
    </div>
  </div>
</template>
<script>
import { mapGetters, mapMutations } from 'vuex'
const path = process.env.VUE_APP_BASE_API
export default {
  name: 'Person',
  data(){
    return {
      path:path
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods:{
    ...mapMutations('user',['ResetUserInfo']),
      handleAvatarSuccess(res){
        this.ResetUserInfo({headerImg:res.data.user.headerImg})
      }
  }
}
</script>
<style lang="scss">
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>