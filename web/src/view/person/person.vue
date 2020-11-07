<template>
  <div>
    <el-row>
      <el-col :span="6">
        <div class="fl-left  avatar-box">
          <div class="user-card">
            <el-avatar  :size="160" :src="userInfo.headerImg" shape="square" @click.native="openChooseImg"></el-avatar>
            <div class="user-personality">
              <p class="nickname">{{userInfo.nickName}}</p>
              <p>我是个性签名</p>
            </div>
            <div class="user-information">
              <ul>
                <li>
                  <i class="el-icon-user"></i>资深前端工程师
                </li>
                <li>
                  <i class="el-icon-data-analysis"></i>北京反转极光科技有限公司-技术部-前端事业群
                </li>
                <li>
                  <i class="el-icon-video-camera-solid"></i>中国·北京市·朝阳区
                </li>
                <li>
                  <i class="el-icon-medal-1"></i>goLang/JavaScript/Vue/Gorm
                </li>
              </ul>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="18">
        <div class="user-addcount">
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="账号绑定" name="second">
              <ul>
                <li>
                  <p class="title">密保手机</p>
                  <p class="desc">已绑定手机:1245678910 <a href="#">立即修改</a></p>
                </li>
                <li>
                  <p class="title">密保邮箱</p>
                  <p class="desc">已绑定邮箱：gin-vue-admin@google.com.cn<a href="#">立即修改</a></p>
                </li>
                <li>
                  <p class="title">密保问题</p>
                  <p class="desc">未设置密保问题 <a href="#">去设置</a></p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-col>
    </el-row>

    <ChooseImg ref="chooseImg" @enter-img="enterImg"/>
    <!--      <div>用户ID：{{userInfo.uuid}}</div>-->
    <!--      <div>用户昵称：{{userInfo.nickName}}</div>-->
    <!--      <div>用户组：{{userInfo.authority&&userInfo.authority.authorityName}}</div>-->
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
      path:path,
      activeName: 'second',

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
    handleClick(tab, event) {
      console.log(tab, event);
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
  .avatar-box{
    box-shadow: -2px 0 20px -16px;
    width: 80%;
    height: 100%;
    .user-card{
      min-height: calc(90vh - 200px);
      padding: 30px 20px;
      text-align: center;
      .el-avatar{
        border-radius: 50%;
      }
      .user-personality{
        padding: 24px 0;
        text-align: center;
        p{
          font-size: 16px;
        }
        .nickname{
          font-size: 26px;
        }
      }
      .user-information{
        width: 100%;
        height: 100%;
        text-align: left;
        ul{
          display: inline-block;
          height: 100%;
          li{
            i{
              margin-right: 8px;
            }
            padding: 20px 0;
            font-size: 16px;
            font-weight: 400;
            color: #606266;
          }
        }
      }
    }
  }
  .user-addcount{
    ul{
      li{
        .title{
          padding: 10px;
          font-size: 18px;
          color: 	#696969;
        }
        .desc{
            font-size: 16px;
          padding: 0 10px 20px 10px ;
          color: 	#A9A9A9;
          a{
            color: rgb(64, 158, 255);
            float: right;
          }
        }
        border-bottom: 2px solid #f0f2f5;
      }
    }
  }

</style>