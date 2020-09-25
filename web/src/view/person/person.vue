<template>
  <div>
    <div class="fl-left left-mg-xs">
      <el-avatar :size="120" :src="userInfo.headerImg" shape="square" @click.native="openChooseImg"></el-avatar>
    </div>
    <div class="fl-left left-mg-lg">
      <div>用户ID：{{userInfo.uuid}}</div>
      <div>用户昵称：{{userInfo.nickName}}</div>
      <div>用户组：{{userInfo.authority&&userInfo.authority.authorityName}}</div>
    </div>
    <ChooseImg ref="chooseImg" @enter-img="enterImg"/>
  </div>
</template>
<script>
import ChooseImg from "@/components/chooseImg";
import {setUserInfo} from "@/api/user"
import { mapGetters, mapMutations } from 'vuex'
const path = process.env.VUE_APP_BASE_API
export default {
  name: 'Person',
  data(){
    return {
      path:path
    }
  },
  components: {
		ChooseImg
	},
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods:{
    ...mapMutations('user',['ResetUserInfo']),
      openChooseImg(){
        this.$refs.chooseImg.open()
      },
      async enterImg(url){
        const res = await setUserInfo({headerImg:url,ID:this.userInfo.ID})
        if(res.code == 0){
          this.ResetUserInfo({headerImg:url})
          this.$message({
            type:"success",
            message:"设置成功"
          }
          )
        }
      },
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