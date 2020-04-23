<template>
  <el-container class="login-register-box">
    <vue-particle-line></vue-particle-line>
    <el-main class="login-box">
      <h1 class="title-1">
        <img class="logo" :src="require('@/assets/logo.png')" alt="" srcset="">
      </h1>
      <el-form :model="loginForm" :rules="rules" ref="loginForm">
        <el-form-item prop="username">
          <el-input placeholder="请输入用户名" v-model="loginForm.username"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            :type="lock==='lock'?'password':'text'"
            placeholder="请输入密码"
            v-model="loginForm.password"
          >
            <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
          </el-input>
        </el-form-item>
        <el-form-item style="position:relative">
          <el-input
            v-model="loginForm.captcha"
            name="logVerify"
            placeholder="请输入验证码"
            maxlength="10"
          />
          <img v-if="picPath" :src="path + picPath" alt="请输入验证码" @click="loginVefify()" class="vPic">
        </el-form-item>
        <el-form-item>
          <el-button @click="submitForm" style="width:100%">登 录</el-button>
        </el-form-item>
      </el-form>
      <h3 class="title-3 fl-right">测试用户:admin 密码:123456</h3>
    </el-main>
  </el-container>
</template>

<script>
import { mapActions } from 'vuex'
import { captcha } from '@/api/user'
const path = process.env.VUE_APP_BASE_API
export default {
  name: 'Login',
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error('请输入正确的用户名'))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error('请输入正确的密码'))
      } else {
        callback()
      }
    }

    return {
      lock: 'lock',
      loginForm: {
        username: '',
        password: '',
        captcha:'',
        captchaId: '',
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }]
      },
      path:path,
      logVerify:'',
      picPath:''
    }
  },
  created() {
    this.loginVefify()
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    async login() {
      await this.LoginIn(this.loginForm)
    },
    async submitForm() {
      this.$refs.loginForm.validate(async v => {
        if (v) {
          this.login()
          this.loginVefify()
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写登录信息',
            showClose: true
          })
          this.loginVefify()
          return false
        }
      })
    },
    changeLock() {
      this.lock === 'lock' ? (this.lock = 'unlock') : (this.lock = 'lock')
    },
    loginVefify() {
      captcha({}).then(ele=>{
        this.picPath = ele.data.picPath
        this.loginForm.captchaId = ele.data.captchaId
      })
    }
  }
}
</script>

<style scoped lang="scss">
.login-register-box {
  background: #fff;
  height: 100vh;
  .login-box {
    width: 40vw;
    position: absolute;
    left: 50%;
    margin-left: -22vw;
    top:5vh;
    .logo{
      height: 35vh;
      width: 35vh;
    }
  }
  .vPic{
    position: absolute;
    right: 10px;
    bottom: 0px;   // 适配ie
  }
}
</style>
