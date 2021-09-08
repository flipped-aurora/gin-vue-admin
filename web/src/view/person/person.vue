<template>
  <div>
    <el-row>
      <el-col :span="6">
        <div class="fl-left avatar-box">
          <div class="user-card">
            <div class="user-headpic-update" :style="{ 'background-image': `url(${(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg})`,'background-repeat':'no-repeat','background-size':'cover' }">
              <span class="update" @click="openChooseImg">
                <i class="el-icon-edit" />
                重新上传</span>
            </div>
            <div class="user-personality">
              <p class="nickname">{{ userInfo.nickName }}</p>
              <p class="person-info">这个家伙很懒，什么都没有留下</p>
            </div>
            <div class="user-information">
              <ul>
                <li>
                  <i class="el-icon-user" />{{ userInfo.nickName }}
                </li>
                <el-tooltip class="item" effect="light" content="北京反转极光科技有限公司-技术部-前端事业群" placement="top">
                  <li>
                    <i class="el-icon-data-analysis" />北京反转极光科技有限公司-技术部-前端事业群
                  </li>
                </el-tooltip>
                <li>
                  <i class="el-icon-video-camera-solid" />中国·北京市·朝阳区
                </li>
                <el-tooltip class="item" effect="light" content="goLang/JavaScript/Vue/Gorm" placement="top">
                  <li>
                    <i class="el-icon-medal-1" />goLang/JavaScript/Vue/Gorm
                  </li>
                </el-tooltip>

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
                  <p class="desc">
                    已绑定手机:1245678910
                    <a href="javascript:void(0)">立即修改</a>
                  </p>
                </li>
                <li>
                  <p class="title">密保邮箱</p>
                  <p class="desc">
                    已绑定邮箱：gin-vue-admin@google.com.cn
                    <a href="javascript:void(0)">立即修改</a>
                  </p>
                </li>
                <li>
                  <p class="title">密保问题</p>
                  <p class="desc">
                    未设置密保问题
                    <a href="javascript:void(0)">去设置</a>
                  </p>
                </li>
                <li>
                  <p class="title">修改密码</p>
                  <p class="desc">
                    修改个人密码
                    <a href="javascript:void(0)" @click="showPassword=true">修改密码</a>
                  </p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-col>
    </el-row>

    <ChooseImg ref="chooseImg" @enter-img="enterImg" />

    <el-dialog v-model="showPassword" title="修改密码" width="360px" @close="clearPassword">
      <el-form ref="modifyPwdForm" :model="pwdModify" :rules="rules" label-width="80px">
        <el-form-item :minlength="6" label="原密码" prop="password">
          <el-input v-model="pwdModify.password" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="新密码" prop="newPassword">
          <el-input v-model="pwdModify.newPassword" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="确认密码" prop="confirmPassword">
          <el-input v-model="pwdModify.confirmPassword" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showPassword=false">取 消</el-button>
          <el-button type="primary" @click="savePassword">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import ChooseImg from '@/components/chooseImg/index.vue'
import { setUserInfo, changePassword } from '@/api/user.js'

import { mapGetters, mapMutations } from 'vuex'
const path = import.meta.env.VITE_BASE_API
export default {
  name: 'Person',
  components: {
    ChooseImg
  },
  data() {
    return {
      path: path,
      activeName: 'second',
      showPassword: false,
      pwdModify: {},
      rules: {
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, message: '最少6个字符', trigger: 'blur' }
        ],
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          { min: 6, message: '最少6个字符', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请输入确认密码', trigger: 'blur' },
          { min: 6, message: '最少6个字符', trigger: 'blur' },
          {
            validator: (rule, value, callback) => {
              if (value !== this.pwdModify.newPassword) {
                callback(new Error('两次密码不一致'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      }
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods: {
    ...mapMutations('user', ['ResetUserInfo']),
    savePassword() {
      this.$refs.modifyPwdForm.validate(valid => {
        if (valid) {
          changePassword({
            username: this.userInfo.userName,
            password: this.pwdModify.password,
            newPassword: this.pwdModify.newPassword
          }).then((res) => {
            if (res.code === 0) {
              this.$message.success('修改密码成功！')
            }
            this.showPassword = false
          })
        } else {
          return false
        }
      })
    },
    clearPassword() {
      this.pwdModify = {
        password: '',
        newPassword: '',
        confirmPassword: ''
      }
      this.$refs.modifyPwdForm.clearValidate()
    },
    openChooseImg() {
      this.$refs.chooseImg.open()
    },
    async enterImg(url) {
      const res = await setUserInfo({ headerImg: url, ID: this.userInfo.ID })
      if (res.code === 0) {
        this.ResetUserInfo({ headerImg: url })
        this.$message({
          type: 'success',
          message: '设置成功'
        })
      }
    },
    handleClick(tab, event) {
      console.log(tab, event)
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
.avatar-box {
  box-shadow: -2px 0 20px -16px;
  width: 80%;
  height: 100%;
  .user-card {
    min-height: calc(90vh - 200px);
    padding: 30px 20px;
    text-align: center;
    .el-avatar {
      border-radius: 50%;
    }
    .user-personality {
      padding: 24px 0;
      text-align: center;
      p {
        font-size: 16px;
      }
      .nickname {
        font-size: 26px;
      }
      .person-info{
        margin-top: 6px;
        font-size: 14px;
        color:#999
      }
    }
    .user-information {
      width: 100%;
      height: 100%;
      text-align: left;
      ul {
        display: inline-block;
        height: 100%;
        width: 100%;
        li {
          width: 100%;
          white-space:nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
          i {
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
.user-addcount {
  ul {
    li {
      .title {
        padding: 10px;
        font-size: 18px;
        color: #696969;
      }
      .desc {
        font-size: 16px;
        padding: 0 10px 20px 10px;
        color: #a9a9a9;
        a {
          color: rgb(64, 158, 255);
          float: right;
        }
      }
      border-bottom: 2px solid #f0f2f5;
    }
  }
}
.user-headpic-update{
    width: 120px;
    height: 120px;
    line-height: 120px;
    margin: 0 auto;
    display: flex;
    justify-content: center;
    border-radius: 20px;
     &:hover{
      color: #fff;
      background: linear-gradient(to bottom, rgba(255,255,255,0.15) 0%, rgba(0,0,0,0.15) 100%), radial-gradient(at top center, rgba(255,255,255,0.40) 0%, rgba(0,0,0,0.40) 120%) #989898;
      background-blend-mode: multiply,multiply;
      .update{
        color:#fff ;
      }
    }
    .update{
      height: 120px;
      width: 120px;
      text-align: center;
      color:transparent;
    }
  }
</style>
