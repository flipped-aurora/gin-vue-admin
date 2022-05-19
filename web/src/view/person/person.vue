<template>
  <div>
    <el-row>
      <el-col :span="6">
        <div class="fl-left avatar-box">
          <div class="user-card">
            <div
              class="user-headpic-update"
              :style="{
                'background-image': `url(${
                  userStore.userInfo.headerImg &&
                  userStore.userInfo.headerImg.slice(0, 4) !== 'http'
                    ? path + userStore.userInfo.headerImg
                    : userStore.userInfo.headerImg
                })`,
                'background-repeat': 'no-repeat',
                'background-size': 'cover',
              }"
            >
              <span class="update" @click="openChooseImg">
                <el-icon>
                  <edit />
                </el-icon>
                重新上传</span>
            </div>
            <div class="user-personality">
              <p v-if="!editFlag" class="nickName">
                {{ userStore.userInfo.nickName }}
                <el-icon class="pointer" color="#66b1ff" @click="openEidt">
                  <edit />
                </el-icon>
              </p>
              <p v-if="editFlag" class="nickName">
                <el-input v-model="nickName" />
                <el-icon class="pointer" color="#67c23a" @click="enterEdit">
                  <check />
                </el-icon>
                <el-icon class="pointer" color="#f23c3c" @click="closeEdit">
                  <close />
                </el-icon>
              </p>
              <p class="person-info">这个家伙很懒，什么都没有留下</p>
            </div>
            <div class="user-information">
              <ul>
                <li>
                  <el-icon>
                    <user />
                  </el-icon>
                  {{ userStore.userInfo.nickName }}
                </li>
                <el-tooltip
                  class="item"
                  effect="light"
                  content="北京反转极光科技有限公司-技术部-前端事业群"
                  placement="top"
                >
                  <li>
                    <el-icon>
                      <data-analysis />
                    </el-icon>
                    北京反转极光科技有限公司-技术部-前端事业群
                  </li>
                </el-tooltip>
                <li>
                  <el-icon>
                    <video-camera />
                  </el-icon>
                  中国·北京市·朝阳区
                </li>
                <el-tooltip
                  class="item"
                  effect="light"
                  content="GoLang/JavaScript/Vue/Gorm"
                  placement="top"
                >
                  <li>
                    <el-icon>
                      <medal />
                    </el-icon>
                    GoLang/JavaScript/Vue/Gorm
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
                    已绑定手机:{{ userStore.userInfo.phone }}
                    <a href="javascript:void(0)" @click="changePhoneFlag = true">立即修改</a>
                  </p>
                </li>
                <li>
                  <p class="title">密保邮箱</p>
                  <p class="desc">
                    已绑定邮箱：{{ userStore.userInfo.email }}
                    <a href="javascript:void(0)" @click="changeEmailFlag = true">立即修改</a>
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
                    <a
                      href="javascript:void(0)"
                      @click="showPassword = true"
                    >修改密码</a>
                  </p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-col>
    </el-row>

    <ChooseImg ref="chooseImgRef" @enter-img="enterImg" />

    <el-dialog
      v-model="showPassword"
      title="修改密码"
      width="360px"
      @close="clearPassword"
    >
      <el-form
        ref="modifyPwdForm"
        :model="pwdModify"
        :rules="rules"
        label-width="80px"
      >
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
          <el-button
            size="small"
            @click="showPassword = false"
          >取 消</el-button>
          <el-button
            size="small"
            type="primary"
            @click="savePassword"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="changePhoneFlag" title="绑定手机" width="600px">
      <el-form :model="phoneForm">
        <el-form-item label="手机号" label-width="120px">
          <el-input v-model="phoneForm.phone" placeholder="请输入手机号" autocomplete="off" />
        </el-form-item>
        <el-form-item label="验证码" label-width="120px">
          <div class="code-box">
            <el-input v-model="phoneForm.code" autocomplete="off" placeholder="请自行设计短信服务，此处为模拟随便写" style="width:300px" />
            <el-button size="small" type="primary" :disabled="time>0" @click="getCode">{{ time>0?`(${time}s)后重新获取`:'获取验证码' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button
            size="small"
            @click="closeChangePhone"
          >取消</el-button>
          <el-button
            type="primary"
            size="small"
            @click="changePhone"
          >更改</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog v-model="changeEmailFlag" title="绑定邮箱" width="600px">
      <el-form :model="emailForm">
        <el-form-item label="邮箱" label-width="120px">
          <el-input v-model="emailForm.email" placeholder="请输入邮箱" autocomplete="off" />
        </el-form-item>
        <el-form-item label="验证码" label-width="120px">
          <div class="code-box">
            <el-input v-model="emailForm.code" placeholder="请自行设计邮件服务，此处为模拟随便写" autocomplete="off" style="width:300px" />
            <el-button size="small" type="primary" :disabled="emailTime>0" @click="getEmailCode">{{ emailTime>0?`(${emailTime}s)后重新获取`:'获取验证码' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button
            size="small"
            @click="closeChangeEmail"
          >取消</el-button>
          <el-button
            type="primary"
            size="small"
            @click="changeEmail"
          >更改</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Person',
}
</script>

<script setup>
import ChooseImg from '@/components/chooseImg/index.vue'
import { setSelfInfo, changePassword } from '@/api/user.js'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

const path = ref(import.meta.env.VITE_BASE_API + '/')
const activeName = ref('second')
const rules = reactive({
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' },
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: '请输入确认密码', trigger: 'blur' },
    { min: 6, message: '最少6个字符', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== pwdModify.value.newPassword) {
          callback(new Error('两次密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()
const modifyPwdForm = ref(null)
const showPassword = ref(false)
const pwdModify = ref({})
const nickName = ref('')
const editFlag = ref(false)
const savePassword = async() => {
  modifyPwdForm.value.validate((valid) => {
    if (valid) {
      changePassword({
        username: userStore.userInfo.userName,
        password: pwdModify.value.password,
        newPassword: pwdModify.value.newPassword,
      }).then((res) => {
        if (res.code === 0) {
          ElMessage.success('修改密码成功！')
        }
        showPassword.value = false
      })
    } else {
      return false
    }
  })
}

const clearPassword = () => {
  pwdModify.value = {
    password: '',
    newPassword: '',
    confirmPassword: '',
  }
  modifyPwdForm.value.clearValidate()
}

const chooseImgRef = ref(null)
const openChooseImg = () => {
  chooseImgRef.value.open()
}

const enterImg = async(url) => {
  const res = await setSelfInfo({ headerImg: url })
  if (res.code === 0) {
    userStore.ResetUserInfo({ headerImg: url })
    ElMessage({
      type: 'success',
      message: '设置成功',
    })
  }
}

const openEidt = () => {
  nickName.value = userStore.userInfo.nickName
  editFlag.value = true
}

const closeEdit = () => {
  nickName.value = ''
  editFlag.value = false
}

const enterEdit = async() => {
  const res = await setSelfInfo({
    nickName: nickName.value
  })
  if (res.code === 0) {
    userStore.ResetUserInfo({ nickName: nickName.value })
    ElMessage({
      type: 'success',
      message: '设置成功',
    })
  }
  nickName.value = ''
  editFlag.value = false
}

const handleClick = (tab, event) => {
  console.log(tab, event)
}

const changePhoneFlag = ref(false)
const time = ref(0)
const phoneForm = reactive({
  phone: '',
  code: ''
})

const getCode = async() => {
  time.value = 60
  let timer = setInterval(() => {
    time.value--
    if (time.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangePhone = () => {
  changePhoneFlag.value = false
  phoneForm.phone = ''
  phoneForm.code = ''
}

const changePhone = async() => {
  const res = await setSelfInfo({ phone: phoneForm.phone })
  if (res.code === 0) {
    ElMessage.success('修改成功')
    userStore.ResetUserInfo({ phone: phoneForm.phone })
    closeChangePhone()
  }
}

const changeEmailFlag = ref(false)
const emailTime = ref(0)
const emailForm = reactive({
  email: '',
  code: ''
})

const getEmailCode = async() => {
  emailTime.value = 60
  let timer = setInterval(() => {
    emailTime.value--
    if (emailTime.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangeEmail = () => {
  changeEmailFlag.value = false
  emailForm.email = ''
  emailForm.code = ''
}

const changeEmail = async() => {
  const res = await setSelfInfo({ email: emailForm.email })
  if (res.code === 0) {
    ElMessage.success('修改成功')
    userStore.ResetUserInfo({ email: emailForm.email })
    closeChangeEmail()
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
      .nickName {
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 26px;
      }
      .person-info {
        margin-top: 6px;
        font-size: 14px;
        color: #999;
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
          white-space: nowrap;
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
.user-headpic-update {
  width: 120px;
  height: 120px;
  line-height: 120px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  border-radius: 20px;
  &:hover {
    color: #fff;
    background: linear-gradient(
        to bottom,
        rgba(255, 255, 255, 0.15) 0%,
        rgba(0, 0, 0, 0.15) 100%
      ),
      radial-gradient(
          at top center,
          rgba(255, 255, 255, 0.4) 0%,
          rgba(0, 0, 0, 0.4) 120%
        )
        #989898;
    background-blend-mode: multiply, multiply;
    .update {
      color: #fff;
    }
  }
  .update {
    height: 120px;
    width: 120px;
    text-align: center;
    color: transparent;
  }
}
.pointer {
  cursor: pointer;
}
.code-box{
  display: flex;
  justify-content: space-between;
}
</style>
