<template>

  <el-container class="login-regist-box">
    <vue-particle-line>
  </vue-particle-line>
    <el-main class="login-box">
      <el-form :model="loginForm" :rules="rules" label-width="100px" ref="loginForm" status-icon>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input :type="lock==='lock'?'password':'text'" v-model="loginForm.password">
            <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
          </el-input>
        </el-form-item>
      </el-form>
      <el-button @click="submitForm" style="float:right;width:calc(100% - 100px)">登 录</el-button>
    </el-main>
  </el-container>

</template>

<script>
import { mapActions } from 'vuex'
export default {
  name: 'Login',
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
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
        password: ''
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }]
      }
    }
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    async login() {
      await this.LoginIn(this.loginForm)
    },
    async submitForm() {
      this.$refs.loginForm.validate(async v => {
        if (v) {
          await this.login()
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写登录信息',
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
.login-regist-box {
  background: #409eff;
  height: 100vh;
  .login-box{
    width: 40vw;
    position: absolute;
    left: 50%;
    margin-left: -22vw;
    top:25vh;
  }
}
</style>