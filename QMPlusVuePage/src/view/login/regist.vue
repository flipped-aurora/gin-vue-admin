<template>
  <el-container class="login-regist-box">
       <vue-particle-line>
  </vue-particle-line>
    <el-main class="login-box">
      <el-form :model="registForm" :rules="rules" label-width="100px" ref="registForm" status-icon>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="registForm.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input :type="lock==='lock'?'password':'text'" v-model="registForm.password">
            <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
          </el-input>
        </el-form-item>
        <el-form-item label="重复密码" prop="rePassword">
          <el-input :type="lock==='lock'?'password':'text'" v-model="registForm.rePassword">
            <i :class="'el-input__icon el-icon-' + lock" @click="changeLock" slot="suffix"></i>
          </el-input>
        </el-form-item>
      </el-form>
      <el-button style="float:right;width:calc(100% - 100px)" @click="submitForm">注册</el-button>
    </el-main>
  </el-container>
</template>

<script>
import { mapActions } from 'vuex'
import { regist } from '@/api/user'
export default {
  name: 'Regist',
  data() {
    const ratioPassword = (rule, value, callback) => {
      if (value != this.registForm.password) {
        return callback(new Error('两次密码不同'))
      } else {
        callback()
      }
    }
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
      registForm: {
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
      this.$refs.registForm.validate(async v => {
        if (v) {
          const res = await regist(this.registForm)
          if (res.success) {
            this.$message({
              type: 'success',
              message: '注册成功',
              showClose: true
            })
            await this.LoginIn(this.registForm)
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

<style lang="scss">
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