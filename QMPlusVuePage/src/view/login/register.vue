<template>
  <el-container class="login-register-box">
    <vue-particle-line></vue-particle-line>
    <el-main class="login-box">
      <h1 class="title-1">
        <img class="logo" :src="require('@/assets/logo.png')">
      </h1>
      <el-form :model="registerForm" :rules="rules" ref="registerForm">
        <el-form-item prop="username">
          <el-input placeholder="请输入用户名" v-model="registerForm.username"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            :type="lock==='lock'?'password':'text'"
            placeholder="请输入密码"
            v-model="registerForm.password"
          >
            <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
          </el-input>
        </el-form-item>
        <el-form-item prop="rePassword">
          <el-input
            :type="lock==='lock'?'password':'text'"
            placeholder="请再次输入密码"
            v-model="registerForm.rePassword"
          >
            <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="submitForm" style="width:100%">注 册</el-button>
        </el-form-item>
      </el-form>
    </el-main>
  </el-container>
</template>

<script>
import { mapActions } from 'vuex'
import { register } from '@/api/user'
export default {
  name: 'Register',
  data() {
    const ratioPassword = (rule, value, callback) => {
      if (value != this.registerForm.password) {
        return callback(new Error('两次密码不同'))
      } else {
        callback()
      }
    }
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
      registerForm: {
        username: '',
        password: '',
        rePassword: ''
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }],
        rePassword: [{ validator: ratioPassword, trigger: 'blur' }]
      }
    }
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    async submitForm() {
      this.$refs.registerForm.validate(async v => {
        if (v) {
          const res = await register(this.registerForm)
          if (res.success) {
            this.$message({
              type: 'success',
              message: '注册成功',
              showClose: true
            })
            this.$router.push({name:"login"})
          }
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写注册信息',
            showClose: true
          })
          return false
        }
      })
    },
    changeLock() {
      this.lock === 'lock' ? (this.lock = 'unlock') : (this.lock = 'lock')
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
}
</style>