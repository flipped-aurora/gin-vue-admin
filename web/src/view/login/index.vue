<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <div class="login_panle_form_title">
          <img
            class="login_panle_form_title_logo"
            :src="$GIN_VUE_ADMIN.appLogo"
            alt
          >
          <p class="login_panle_form_title_p">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <el-form
          ref="loginForm"
          :model="loginForm"
          :rules="rules"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input v-model="loginForm.username" placeholder="请输入用户名">
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <user />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginForm.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon><component :is="lock" @click="changeLock" /></el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative" prop="captcha">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                alt="请输入验证码"
                @click="loginVerify()"
              >
            </div>
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              style="width: 46%"
              @click="checkInit"
            >前往初始化</el-button>
            <el-button
              type="primary"
              style="width: 46%; margin-left: 8%"
              @click="submitForm"
            >登 录</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="login_panle_right" />
      <div class="login_panle_foot">
        <div class="links">
          <a href="http://doc.henrongyi.top/" target="_blank">
            <img src="@/assets/docs.png" class="link-icon">
          </a>
          <a href="https://support.qq.com/product/371961" target="_blank">
            <img src="@/assets/kefu.png" class="link-icon">
          </a>
          <a href="https://github.com/flipped-aurora/gin-vue-admin" target="_blank">
            <img src="@/assets/github.png" class="link-icon">
          </a>
          <a href="https://space.bilibili.com/322210472" target="_blank">
            <img src="@/assets/video.png" class="link-icon">
          </a>
        </div>
        <div class="copyright">
          <bootomInfo />
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { mapActions } from 'vuex'
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import bootomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
export default {
  name: 'Login',
  components: {
    bootomInfo
  },
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5) {
        return callback(new Error('请输入正确的用户名'))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6) {
        return callback(new Error('请输入正确的密码'))
      } else {
        callback()
      }
    }
    return {
      lock: 'lock',
      loginForm: {
        username: 'admin',
        password: '123456',
        captcha: '',
        captchaId: ''
      },
      rules: {
        username: [{ validator: checkUsername, trigger: 'blur' }],
        password: [{ validator: checkPassword, trigger: 'blur' }],
        captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' },
          {
            message: '验证码格式不正确',
            trigger: 'blur',
          }]
      },
      logVerify: '',
      picPath: ''
    }
  },
  created() {
    this.loginVerify()
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    async checkInit() {
      const res = await checkDB()
      if (res.code === 0) {
        if (res.data?.needInit) {
          this.$store.commit('user/NeedInit')
          this.$router.push({ name: 'Init' })
        } else {
          this.$message({
            type: 'info',
            message: '已配置数据库信息，无法初始化'
          })
        }
      }
    },
    async login() {
      return await this.LoginIn(this.loginForm)
    },
    async submitForm() {
      this.$refs.loginForm.validate(async(v) => {
        if (v) {
          const flag = await this.login()
          if (!flag) {
            this.loginVerify()
          }
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写登录信息',
            showClose: true
          })
          this.loginVerify()
          return false
        }
      })
    },
    changeLock() {
      this.lock = this.lock === 'lock' ? 'unlock' : 'lock'
    },
    loginVerify() {
      captcha({}).then((ele) => {
        this.rules.captcha[1].max = ele.data.captchaLength
        this.rules.captcha[1].min = ele.data.captchaLength
        this.picPath = ele.data.picPath
        this.loginForm.captchaId = ele.data.captchaId
      })
    }
  }
}

</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";
</style>
