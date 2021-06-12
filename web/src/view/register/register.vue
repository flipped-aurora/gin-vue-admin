<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/logo_login.png" alt="">
        </div>
        <div class="header">
          <a href="/"><span class="title">Gin-Vue-Admin</span></a>
        </div>
      </div>
      <div class="main">
        <el-form
          ref="registerForm"
          :model="registerForm"
          :rules="rules"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="username">
            <el-input v-model="registerForm.username" placeholder="请输入用户名">
              <i slot="suffix" class="el-input__icon el-icon-user" />
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="registerForm.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <i
                slot="suffix"
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
              />
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="registerForm.passwordConfirm"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请再次输入密码"
            >
              <i
                slot="suffix"
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
              />
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              style="width: 100%"
              @click="submitForm"
            >注 册</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>
<script>
import { mapActions } from 'vuex'
import { captcha } from '@/api/user'
export default {
  name: 'Login',
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5) {
        return callback(new Error('用户名不能少于5个字符'))
      } else {
        callback()
      }
    }
    const checkPwdError = (rule, value, callback) => {
      if (value.length < 6) {
        return callback(new Error('长度不能小于6位'))
      } else if (/(?=.*?[A-Z])(?=.*?[a-z])(?=.*?\d)(?=.*?[#/\\?!@=,<>;|"$%^&*\-_~+:.'`{}()[\]])/.test(value)) {
        callback()
      } else {
        return callback(new Error('密码必须包含数字字母以及特殊字符'))
      }
    }
    return {
      curYear: 0,
      lock: 'lock',
      registerForm: {
        username: '',
        password: '',
        passwordConfirm: ''
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPwdError, trigger: 'blur' }],
        passwordConfirm: [{ validator: checkPwdError, trigger: 'blur' }]
      },
      logVerify: '',
      picPath: ''
    }
  },
  created() {
    this.curYear = new Date().getFullYear()
  },
  methods: {
    ...mapActions('user', ['Register']),
    async register() {
      const { username, password } = this.registerForm
      return await this.Register({
        username,
        password
      })
    },

    checkPassword() {
      return this.registerForm.password === this.registerForm.passwordConfirm
    },
    async submitForm() {
      this.$refs.registerForm.validate(async(v) => {
        if (v && this.checkPassword()) {
          const flag = await this.register()
        } else if (this.checkPassword()) {
          this.$message({
            type: 'error',
            message: '请正确填写注册信息',
            showClose: true
          })
          return false
        } else {
          this.$message({
            type: 'error',
            message: '两次输入密码不一致',
            showClose: true
          })
        }
      })
    },
    changeLock() {
      this.lock = this.lock === 'lock' ? 'unlock' : 'lock'
    }
  }
}
</script>

<style scoped lang="scss">
@import "@/style/login.scss";
</style>
