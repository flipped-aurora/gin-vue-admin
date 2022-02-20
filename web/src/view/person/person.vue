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
                re-upload</span>
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
              <p class="person-info">This guy is too lazy, nothing left</p>
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
                  content="Beijing Inverse Aurora Science and Technology Co., Ltd. - Technical Department - Front End Business Group"
                  placement="top"
                >
                  <li>
                    <el-icon>
                      <data-analysis />
                    </el-icon>
                    Beijing Inverse Aurora Science and Technology Co., Ltd. - Technical Department - Front End Business Group
                  </li>
                </el-tooltip>
                <li>
                  <el-icon>
                    <video-camera />
                  </el-icon>
                  China · Beijing Chaoyang District
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
            <el-tab-pane label="Account binding" name="second">
              <ul>
                <li>
                  <p class="title">Phone</p>
                  <p class="desc">
                    Binding mobile phone:{{ userStore.userInfo.phone }}
                    <a href="javascript:void(0)" @click="changePhoneFlag = true">Modification</a>
                  </p>
                </li>
                <li>
                  <p class="title">Secret mailbox</p>
                  <p class="desc">
                    Binding mailbox:{{ userStore.userInfo.email }}
                    <a href="javascript:void(0)" @click="changeEmailFlag = true">Modification</a>
                  </p>
                </li>
                <li>
                  <p class="title">Security Question</p>
                  <p class="desc">
                    No security issues are set
                    <a href="javascript:void(0)">Set</a>
                  </p>
                </li>
                <li>
                  <p class="title">Change Password</p>
                  <p class="desc">
                    Modify personal password
                    <a
                      href="javascript:void(0)"
                      @click="showPassword = true"
                    >Change Password</a>
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
      title="Change Password"
      width="360px"
      @close="clearPassword"
    >
      <el-form
        ref="modifyPwdForm"
        :model="pwdModify"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item :minlength="6" label="Password" prop="password">
          <el-input v-model="pwdModify.password" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="New Password" prop="newPassword">
          <el-input v-model="pwdModify.newPassword" show-password />
        </el-form-item>
        <el-form-item :minlength="6" label="Confirm Password" prop="confirmPassword">
          <el-input v-model="pwdModify.confirmPassword" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button
            size="small"
            @click="showPassword = false"
          >Cancel</el-button>
          <el-button
            size="small"
            type="primary"
            @click="savePassword"
          >Save</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="changePhoneFlag" title="Bind mobile phone" width="600px">
      <el-form :model="phoneForm">
        <el-form-item label="Phone" label-width="120px">
          <el-input v-model="phoneForm.phone" placeholder="Please enter the phone" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Verification code" label-width="120px">
          <div class="code-box">
            <el-input v-model="phoneForm.code" autocomplete="off" placeholder="Please design SMS service yourself, here to write casually" style="width:300px" />
            <el-button size="small" type="primary" :disabled="time>0" @click="getCode">{{ time>0?`(${time}s)Re-acquire`:'get verification code' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button
            size="small"
            @click="closeChangePhone"
          >Cancel</el-button>
          <el-button
            type="primary"
            size="small"
            @click="changePhone"
          >Change</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog v-model="changeEmailFlag" title="Binding mailbox" width="600px">
      <el-form :model="emailForm">
        <el-form-item label="Mail" label-width="120px">
          <el-input v-model="emailForm.email" placeholder="Please input your email" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Code" label-width="120px">
          <div class="code-box">
            <el-input v-model="emailForm.code" placeholder="Please design the mail service yourself, here to write casually" autocomplete="off" style="width:300px" />
            <el-button size="small" type="primary" :disabled="emailTime>0" @click="getEmailCode">{{ emailTime>0?`(${emailTime}s)Re-acquire`:'get verification code' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button
            size="small"
            @click="closeChangeEmail"
          >Cancel</el-button>
          <el-button
            type="primary"
            size="small"
            @click="changeEmail"
          >Change</el-button>
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

const path = ref(import.meta.env.VITE_BASE_API)
const activeName = ref('second')
const rules = reactive({
  password: [
    { required: true, message: 'Please enter your password', trigger: 'blur' },
    { min: 6, message: 'At least 6 characters', trigger: 'blur' },
  ],
  newPassword: [
    { required: true, message: 'Please enter a new password', trigger: 'blur' },
    { min: 6, message: 'At least 6 characters', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: 'Please enter a confirmation password', trigger: 'blur' },
    { min: 6, message: 'At least 6 characters', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== pwdModify.value.newPassword) {
          callback(new Error('Two passwords are inconsistent'))
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
          ElMessage.success('successfully change password!')
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
      message: 'Set success',
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
      message: 'Set success',
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
    ElMessage.success('Successfully modified')
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
    ElMessage.success('Successfully modified')
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
