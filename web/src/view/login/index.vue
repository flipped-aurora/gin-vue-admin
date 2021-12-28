<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <!--
        <label for="locale">Select Language:</label>
        <select v-model="$i18n.locale" width="50px">
          <option>en</option>
          <option>cn</option>
          <option>ar</option>
        </select>
        -->
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
            <el-input v-model="loginForm.username" :placeholder="$t('login.entUserName')">
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
              :placeholder="$t('login.entPassword')"
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
              :placeholder="$t('login.entVerificationCode')"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                :alt="$t('login.entVerificationCode')"
                @click="loginVerify()"
              >
            </div>
          </el-form-item>
          <!-- added by mohamed hassan to support multi language -->
          <el-form-item>
            <label for="locale" style="padding-right: 88px">Select Language:</label>
            <el-select v-model="$i18n.locale" value-key="value" placeholder="Select Language">
              <el-option v-for="item in langs" :key="item.value" :label="item.label" :value="item.value">
                <img :src="item.photo"> {{ item.label }}
              </el-option>
            </el-select>
          </el-form-item>
          <!-- end of adding -->
          <el-form-item>
            <el-button
              type="primary"
              style="width: 46%"
              @click="checkInit"
            >{{ $t('login.init') }}</el-button>
            <el-button
              type="primary"
              style="width: 46%; margin-left: 8%"
              @click="submitForm"
            >{{ $t('login.login') }}</el-button>
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
        return callback(new Error(this.$t('login.errUserName')))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6) {
        return callback(new Error(this.$t('login.errPassword')))
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
        captcha: [{ required: true, message: this.$t('login.entVerificationCode'), trigger: 'blur' },
          {
            message: this.$t('login.errVerificationCode'),
            trigger: 'blur',
          }]
      },
      logVerify: '',
      picPath: '',
      langs: [{
        value: 'en',
        label: 'English',
        photo: '@/assets/flags/us.svg'
      }, {
        value: 'cn',
        label: '中文',
        photo: '@/assets/flags/cn.svg'
      }, {
        value: 'ar',
        label: 'العربية',
        photo: '@/assets/flags/eg.svg'
      }],
    }
  },
  created() {
    this.$i18n.locale = this.langs[0].value

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
            message: this.$t('login.errInit')
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
            message: this.$t('login.errLogin'),
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
    },
    setLocale(locale) {
      this.$$i18n.locale = locale
    }
  }
}

</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";

img {
  width: 20px;
  height: 20px;
}

.prefix {
  margin-top: 10px;
}
</style>
