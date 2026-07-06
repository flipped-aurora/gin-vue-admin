<template>
  <div
    id="userLayout"
    class="relative h-full w-full bg-white dark:bg-slate-900 md:h-screen"
  >
    <!-- 右侧蓝色斜切 banner（仅桌面端显示） -->
    <div
      class="banner-oblique absolute inset-y-0 right-0 hidden w-[56%] overflow-hidden bg-[#194bfb] md:block"
    >
      <img
        class="absolute right-0 top-0 h-full w-auto max-w-none"
        src="@/assets/login_right_banner.jpg"
        alt="banner"
      />
    </div>

    <!-- 左侧登录区 -->
    <div
      class="relative z-10 flex h-full w-full items-center justify-center md:w-1/2"
    >
      <div class="w-4/5 md:w-96">
        <EntryBrand class="mb-9" />
        <el-form
          ref="loginForm"
          :model="loginFormData"
          :rules="rules"
          :validate-on-rule-change="false"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username" class="mb-6">
            <el-input
              v-model="loginFormData.username"
              size="large"
              placeholder="请输入用户名"
              suffix-icon="user"
            />
          </el-form-item>
          <el-form-item prop="password" class="mb-6">
            <el-input
              v-model="loginFormData.password"
              show-password
              size="large"
              type="password"
              placeholder="请输入密码"
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
                placeholder="请输入验证码"
                size="large"
                class="flex-1 mr-5"
              />
              <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                <img
                  v-if="picPath"
                  class="w-full h-full"
                  :src="picPath"
                  alt="请输入验证码"
                  @click="loginVerify()"
                />
              </div>
            </div>
          </el-form-item>
          <el-form-item class="mb-6">
            <el-button
              class="shadow shadow-active h-11 w-full"
              type="primary"
              size="large"
              @click="submitForm"
              >登 录</el-button
            >
          </el-form-item>
          <el-form-item v-if="isDev" class="mb-6">
            <el-button
              class="btn-hollow h-11 w-full"
              type="primary"
              plain
              size="large"
              @click="checkInit"
              >前往初始化</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>

    <BottomInfo class="left-0 right-0 absolute bottom-3 mx-auto w-full z-20">
      <div class="links items-center justify-center gap-2 hidden md:flex">
        <a href="https://www.gin-vue-admin.com/" target="_blank">
          <img src="@/assets/docs.png" class="w-8 h-8" alt="文档" />
        </a>
        <a href="https://support.qq.com/product/371961" target="_blank">
          <img src="@/assets/kefu.png" class="w-8 h-8" alt="客服" />
        </a>
        <a
          href="https://github.com/flipped-aurora/gin-vue-admin"
          target="_blank"
        >
          <img src="@/assets/github.png" class="w-8 h-8" alt="github" />
        </a>
        <a href="https://space.bilibili.com/322210472" target="_blank">
          <img src="@/assets/video.png" class="w-8 h-8" alt="视频站" />
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
  import EntryBrand from '@/components/entryBrand/index.vue'
  import { isDev } from '@/utils/env.js'

  defineOptions({
    name: 'Login'
  })

  const router = useRouter()
  const captchaRequiredLength = ref(6)
  // 验证函数
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
  const checkCaptcha = (rule, value, callback) => {
    if (!loginFormData.openCaptcha) {
      return callback()
    }
    const sanitizedValue = (value || '').replace(/\s+/g, '')
    if (!sanitizedValue) {
      return callback(new Error('请输入验证码'))
    }
    if (!/^\d+$/.test(sanitizedValue)) {
      return callback(new Error('验证码须为数字'))
    }
    if (sanitizedValue.length < captchaRequiredLength.value) {
      return callback(
        new Error(`请输入至少${captchaRequiredLength.value}位数字验证码`)
      )
    }
    if (sanitizedValue !== value) {
      loginFormData.captcha = sanitizedValue
    }
    callback()
  }

  // 获取验证码
  const loginVerify = async () => {
    const ele = await captcha()
    captchaRequiredLength.value = Number(ele.data?.captchaLength) || 0
    picPath.value = ele.data?.picPath
    loginFormData.captchaId = ele.data?.captchaId
    loginFormData.openCaptcha = ele.data?.openCaptcha
  }
  loginVerify()

  // 登录相关操作
  const loginForm = ref(null)
  const picPath = ref('')
  const loginFormData = reactive({
    username: 'admin',
    password: '',
    captcha: '',
    captchaId: '',
    openCaptcha: false
  })
  const rules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur' }],
    password: [{ validator: checkPassword, trigger: 'blur' }],
    captcha: [{ validator: checkCaptcha, trigger: 'blur' }]
  })

  const userStore = useUserStore()
  const login = async () => {
    return await userStore.LoginIn(loginFormData)
  }
  const submitForm = () => {
    loginForm.value.validate(async (v) => {
      if (!v) {
        // 未通过前端静态验证
        ElMessage({
          type: 'error',
          message: '请正确填写登录信息',
          showClose: true
        })
        return false
      }

      // 通过验证，请求登陆
      const flag = await login()

      // 登陆失败，刷新验证码
      if (!flag) {
        await loginVerify()
        return false
      }

      // 登陆成功
      return true
    })
  }

  // 跳转初始化
  const checkInit = async () => {
    const res = await checkDB()
    if (res.code === 0) {
      if (res.data?.needInit) {
        userStore.NeedInit()
        await router.push({ name: 'Init' })
      } else {
        ElMessage({
          type: 'info',
          message: '已配置数据库信息，无法初始化'
        })
      }
    }
  }
</script>

<style scoped>
  /* 右侧蓝色面板的左边缘斜切：顶部靠右、底部靠左，蓝色区上窄下宽 */
  .banner-oblique {
    clip-path: polygon(22% 0, 100% 0, 100% 100%, 4% 100%);
  }

  /* 镂空按钮：常态透明底；hover 用浅色底 + 主色文字表现（不再变白字实心） */
  .btn-hollow {
    --el-button-bg-color: transparent;
    --el-button-hover-bg-color: var(--el-color-primary-light-9);
    --el-button-hover-text-color: var(--el-color-primary);
    --el-button-hover-border-color: var(--el-color-primary);
    --el-button-active-bg-color: var(--el-color-primary-light-8);
    --el-button-active-text-color: var(--el-color-primary);
    --el-button-active-border-color: var(--el-color-primary);
  }
</style>
