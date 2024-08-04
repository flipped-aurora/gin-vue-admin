<template>
  <div
    id="userLayout"
    class="w-full h-full relative"
  >
    <div
      class="rounded-lg flex items-center justify-evenly w-full h-full md:w-screen md:h-screen md:bg-[#194bfb]"
    >
      <div class="md:w-3/5 w-10/12 h-full flex items-center justify-evenly">
        <div class="oblique h-[130%] w-3/5 bg-white dark:bg-slate-900 transform -rotate-12 absolute -ml-52" />
        <!-- 分割斜块 -->
        <div class="z-[999] pt-12 pb-10 md:w-96 w-full  rounded-lg flex flex-col justify-between box-border">
          <div>
            <div class="flex items-center justify-center">

              <img
                class="w-24"
                :src="$GIN_VUE_ADMIN.appLogo"
                alt
              >
            </div>
            <div class="mb-9">
              <p class="text-center text-4xl font-bold">{{ $GIN_VUE_ADMIN.appName }}</p>
              <p class="text-center text-sm font-normal text-gray-500 mt-2.5">A management platform using Golang and Vue
              </p>
            </div>
            <div style="padding-left: 92%; padding-bottom: 20px;">
              <el-dropdown
                trigger="click"
                @command="handleSetLanguage"
              >
                <span class="el-dropdown-link">
                  <img
                    src="@/assets/language.svg" alt = "Languages"
                    style="width: 30px; height: 30px;"
                  >
                </span>
                <template #dropdown style="width: 100px">
                  <el-dropdown-menu>
                    <el-dropdown-item
                      :disabled="$i18n.locale==='en'"
                      command="en"
                    ><img
                      src="@/assets/flags/en.svg" alt = "English"
                      class="img"
                    >English</el-dropdown-item>
                    <el-dropdown-item
                      :disabled="$i18n.locale==='zh'"
                      command="zh"
                    ><img
                      src="@/assets/flags/zh.svg" alt = "中文"
                      class="img"
                    >中文</el-dropdown-item>
                    <el-dropdown-item
                      :disabled="$i18n.locale==='ar'"
                      command="ar"
                    ><img
                      src="@/assets/flags/ar.svg" alt = "العربية"
                      class="img"
                    >العربية</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
            <el-form
              ref="loginForm"
              :model="loginFormData"
              :rules="rules"
              :validate-on-rule-change="false"
              @keyup.enter="submitForm"
            >
              <el-form-item
                prop="username"
                class="mb-6"
              >
                <el-input
                  v-model="loginFormData.username"
                  size="large"
                  :placeholder="t('login.entUserName')"
                  suffix-icon="user"
                />
              </el-form-item>
              <el-form-item
                prop="password"
                class="mb-6"
              >
                <el-input
                  v-model="loginFormData.password"
                  show-password
                  size="large"
                  type="password"
                  :placeholder="t('login.entPassword')"
                />
              </el-form-item>
              <el-form-item
                v-if="loginFormData.openCaptcha"
                prop="captcha"
                class="mb-6"
              >
                <div class="flex w-full justify-between">
                  <el-input
                    v-model="loginFormData.captcha"
                    :placeholder="t('login.entVerificationCode')"
                    size="large"
                    class="flex-1 mr-5"
                  />
                  <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                    <img
                      v-if="picPath"
                      class="w-full h-full"
                      :src="picPath"
                      :alt="t('login.entVerificationCode')"
                      @click="loginVerify()"
                    >
                  </div>
                </div>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-active h-11 w-full"
                  type="primary"
                  size="large"
                  @click="submitForm"
                >{{ t('login.login') }}</el-button>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-active h-11 w-full"
                  type="primary"
                  size="large"
                  @click="checkInit"
                >{{ t('login.init') }}</el-button>

              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
      <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]"><img
        class="h-full"
        src="@/assets/login_right_banner.jpg"
        alt="banner"
      ></div>
    </div>

    <BottomInfo class="left-0 right-0 absolute bottom-3 mx-auto  w-full z-20">
      <div class="links items-center justify-center gap-2 hidden md:flex">
        <a
          href="https://www.gin-vue-admin.com/"
          target="_blank"
        >
          <img
            src="@/assets/docs.png"
            class="w-8 h-8"
            alt="文档"
          >
        </a>
        <a
          href="https://support.qq.com/product/371961"
          target="_blank"
        >
          <img
            src="@/assets/kefu.png"
            class="w-8 h-8"
            alt="客服"
          >
        </a>
        <a
          href="https://github.com/flipped-aurora/gin-vue-admin"
          target="_blank"
        >
          <img
            src="@/assets/github.png"
            class="w-8 h-8"
            alt="github"
          >
        </a>
        <a
          href="https://space.bilibili.com/322210472"
          target="_blank"
        >
          <img
            src="@/assets/video.png"
            class="w-8 h-8"
            alt="视频站"
          >
        </a>
      </div>
    </BottomInfo>
  </div>
</template>

<script setup>
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'
import Cookies from 'js-cookie' // added by mohamed hassan to support multilanguage
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const i18n = useI18n() // added by mohamed hassan to support multilanguage
const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: "Login",
})

const router = useRouter()
// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error(t('login.errUserName')))
  } else {
    callback()
  }
}

const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error(t('login.errPassword')))
  } else {
    callback()
  }
}

// 获取验证码
const loginVerify = async() => {
  const ele = await captcha()
  rules.captcha.push({
    max: ele.data.captchaLength,
    min: ele.data.captchaLength,
    message: `t('login.pleaseEnter') ${ele.data.captchaLength} t('login.verificationCode')`,
    trigger: 'blur',
  })
  picPath.value = ele.data.picPath
  loginFormData.captchaId = ele.data.captchaId
  loginFormData.openCaptcha = ele.data.openCaptcha
}

const getLanguage = () => {
  var lang = Cookies.get('language')
  return (lang || 'en')
}

getLanguage()

loginVerify()

// 登录相关操作
const loginForm = ref(null)
const picPath = ref('')

const loginFormData = reactive({
  username: 'admin',
  password: '',
  captcha: '',
  captchaId: '',
  openCaptcha: false,
})

const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    {
      message: t('login.errVerificationCode'),
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()

const login = async() => {
  return await userStore.LoginIn(loginFormData)
}

const submitForm = () => {
  loginForm.value.validate(async(v) => {
    if (!v) {
      // 未通过前端静态验证
      ElMessage({
        type: 'error',
        message: t('login.errLogin'),
        showClose: true,
      })
      loginVerify()
      return false
    }

    // 通过验证，请求登陆
    const flag = await login()

    // 登陆失败，刷新验证码
    if (!flag) {
      loginVerify()
      return false
    }

    // 登陆成功
    return true
  })
}

// 跳转初始化
const checkInit = async() => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: t('login.errInit'),
      })
    }
  }
}

const handleSetLanguage = (lang) => {
  // console.log('handleSetLanguage() called with value: ' + lang)
  i18n.locale.value = lang

  userStore.setLanguage(lang)

  // console.log('userStore handleSetLanguage() called with value: ' + userStore.getLanguage())

  Cookies.set('language', lang)

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

  ElMessage({
    message: t('general.langSwitch'),
    type: 'success'
  })

  // this.$emit('handerevent')
}
</script>
