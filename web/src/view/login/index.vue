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
        <div style="padding-left: 92%; padding-bottom: 20px;">
          <el-dropdown trigger="click" @command="handleSetLanguage">
            <span class="el-dropdown-link">
              <img src="@/assets/language.svg" style="width: 30px; height: 30px;">
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item :disabled="$i18n.locale==='en'" command="en"><img src="@/assets/flags/en.svg" class="img">English</el-dropdown-item>
                <el-dropdown-item :disabled="$i18n.locale==='zh'" command="zh"><img src="@/assets/flags/zh.svg" class="img">中文</el-dropdown-item>
                <el-dropdown-item :disabled="$i18n.locale==='ar'" command="ar"><img src="@/assets/flags/ar.svg" class="img">العربية</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
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
import Cookies from 'js-cookie'
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
    }
  },
  computed: {
    // language() {
    //   var lang = Cookies.get('language')
    //   console.log('computed: Loaded language from cookies: ' + lang)
    //   return (lang || 'en')
    // }
  },
  created() {
    this.getLanguage()

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
    setLocale(local) {
      console.log('Selected langauge is: ' + local.value)
      this.$i18n.locale = local.value
    },
    getLanguage() {
      var lang = Cookies.get('language')
      return (lang || 'en')
    },
    handleSetLanguage(lang) {
      this.$i18n.locale = lang

      Cookies.set('language', lang)
      // this.$store.dispatch('setLanguage', lang)

      // if (lang === 'ar') {
      //   console.log('Arabic language selected changing to RTL')
      //   document.querySelector('html').classList.add('is-rtl')
      // } else {
      //   console.log('Non Arabic language selected changing to LTR')
      //   document.querySelector('html').classList.add('is-ltr')
      // }

      // const htmlEl = document.querySelector('html')

      // if (this.$i18n.locale === 'ar') {
      //   console.log('change language to arabic and ltr to rtl')
      //   htmlEl.setAttribute('dir', 'rtl')
      // } else {
      //   console.log('change language to english and rtl to ltr')
      //   htmlEl.setAttribute('dir', 'ltr')
      // }

      // htmlEl.setAttribute('lang', lang)

      this.$message({
        message: this.$t('general.langSwitch'),
        type: 'success'
      })

      // this.$emit('handerevent')
    }
  }
}

</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";

img {
  padding-right: 20px;
  width: 20px;
  height: 20px;
}

prefix {
  margin-top: 10px;
  width: 100px;
  height: 100px;
}

.international-icon {
  font-size: 20px;
  cursor: pointer;
  vertical-align: -5px!important;
}

html.is-rtl * {
    direction: rtl;
}

html.is-ltr * {
    direction: ltr;
}
</style>
