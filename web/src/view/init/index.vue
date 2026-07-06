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

    <!-- 左侧内容区 -->
    <div class="relative z-10 h-full w-full md:w-1/2">
      <div class="flex h-full items-center justify-center">
        <div class="w-4/5 md:w-96">
          <!-- 顶部 logo + 大写字标（无副标题） -->
          <EntryBrand class="mb-9" />
          <div v-if="!page.showForm" class="anim-fade-up">
            <h2 class="mb-6 text-2xl font-bold dark:text-white">初始化须知</h2>
            <ul class="flex flex-col gap-4">
              <li
                class="flex items-start gap-3 text-gray-600 dark:text-gray-300"
              >
                <span
                  class="mt-2 h-1.5 w-1.5 shrink-0 rounded-full bg-[#194bfb]"
                />
                <span>您需有用一定的 VUE 和 GOLANG 基础</span>
              </li>
              <li
                class="flex items-start gap-3 text-gray-600 dark:text-gray-300"
              >
                <span
                  class="mt-2 h-1.5 w-1.5 shrink-0 rounded-full bg-[#194bfb]"
                />
                <span>
                  请您确认是否已经阅读过<a
                    class="mx-1 font-bold text-[#194bfb]"
                    href="https://www.gin-vue-admin.com"
                    target="_blank"
                    >官方文档</a
                  ><a
                    class="font-bold text-[#194bfb]"
                    href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=2"
                    target="_blank"
                    >初始化视频</a
                  >
                </span>
              </li>
              <li
                class="flex items-start gap-3 text-gray-600 dark:text-gray-300"
              >
                <span
                  class="mt-2 h-1.5 w-1.5 shrink-0 rounded-full bg-[#194bfb]"
                />
                <span>请您确认是否了解后续的配置流程</span>
              </li>
              <li
                class="flex items-start gap-3 text-gray-600 dark:text-gray-300"
              >
                <span
                  class="mt-2 h-1.5 w-1.5 shrink-0 rounded-full bg-[#194bfb]"
                />
                <span
                  >如果您使用 mysql 数据库，请确认数据库引擎为
                  <span class="font-bold text-[#194bfb]">innoDB</span></span
                >
              </li>
            </ul>
            <p class="mt-4 text-sm text-gray-400 dark:text-gray-500">
              注：开发组不为文档中书写过的内容提供无偿服务
            </p>
            <div class="mt-8 flex gap-4">
              <el-button
                type="primary"
                size="large"
                class="flex-1"
                @click="goDoc"
              >
                阅读文档
              </el-button>
              <el-button
                type="primary"
                plain
                size="large"
                class="btn-hollow flex-1"
                @click="showNext"
              >
                我已确认
              </el-button>
            </div>
          </div>
          <div v-if="page.showForm" class="anim-fade-up">
            <el-form
              ref="formRef"
              :model="form"
              label-position="top"
              size="large"
            >
              <el-form-item label="管理员密码">
                <el-input
                  v-model="form.adminPassword"
                  show-password
                  placeholder="admin账号的默认密码"
                />
              </el-form-item>
              <el-form-item label="数据库类型">
                <el-select
                  v-model="form.dbType"
                  placeholder="请选择"
                  class="w-full"
                  @change="changeDB"
                >
                  <el-option key="mysql" label="mysql" value="mysql" />
                  <el-option key="pgsql" label="pgsql" value="pgsql" />
                  <el-option key="oracle" label="oracle" value="oracle" />
                  <el-option key="mssql" label="mssql" value="mssql" />
                  <el-option key="sqlite" label="sqlite" value="sqlite" />
                </el-select>
              </el-form-item>
              <div class="flex gap-4">
                <el-form-item
                  v-if="form.dbType !== 'sqlite'"
                  label="host"
                  class="flex-1"
                >
                  <el-input
                    v-model="form.host"
                    placeholder="请输入数据库链接"
                  />
                </el-form-item>
                <el-form-item
                  v-if="form.dbType !== 'sqlite'"
                  label="port"
                  class="flex-1"
                >
                  <el-input
                    v-model="form.port"
                    placeholder="请输入数据库端口"
                  />
                </el-form-item>
              </div>
              <div class="flex gap-4">
                <el-form-item
                  v-if="form.dbType !== 'sqlite'"
                  label="userName"
                  class="flex-1"
                >
                  <el-input
                    v-model="form.userName"
                    placeholder="请输入数据库用户名"
                  />
                </el-form-item>
                <el-form-item label="dbName" class="flex-1">
                  <el-input
                    v-model="form.dbName"
                    placeholder="请输入数据库名称"
                  />
                </el-form-item>
              </div>
              <el-form-item v-if="form.dbType !== 'sqlite'" label="password">
                <el-input
                  v-model="form.password"
                  placeholder="请输入数据库密码（没有则为空）"
                />
              </el-form-item>
              <el-form-item v-if="form.dbType === 'sqlite'" label="dbPath">
                <el-input
                  v-model="form.dbPath"
                  placeholder="请输入sqlite数据库文件存放路径"
                />
              </el-form-item>
              <el-form-item v-if="form.dbType === 'pgsql'" label="template">
                <el-input
                  v-model="form.template"
                  placeholder="请输入postgresql指定template"
                />
              </el-form-item>
              <el-form-item>
                <el-button
                  type="primary"
                  size="large"
                  class="w-full"
                  @click="onSubmit"
                >
                  立即初始化
                </el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  // @ts-ignore
  import { initDB } from '@/api/initdb'
  import { reactive, ref } from 'vue'
  import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
  import { useRouter } from 'vue-router'
  import EntryBrand from '@/components/entryBrand/index.vue'

  defineOptions({
    name: 'Init'
  })

  const router = useRouter()

  const page = reactive({
    showReadme: false,
    showForm: false
  })

  const showNext = () => {
    page.showReadme = false
    setTimeout(() => {
      page.showForm = true
    }, 20)
  }

  const goDoc = () => {
    window.open('https://www.gin-vue-admin.com/guide/start-quickly/env.html')
  }

  const out = ref(false)

  const form = reactive({
    adminPassword: '123456',
    dbType: 'mysql',
    host: '127.0.0.1',
    port: '3306',
    userName: 'root',
    password: '',
    dbName: 'gva',
    dbPath: ''
  })

  const changeDB = (val) => {
    switch (val) {
      case 'mysql':
        Object.assign(form, {
          adminPassword: '123456',
          reAdminPassword: '',
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
        break
      case 'pgsql':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'pgsql',
          host: '127.0.0.1',
          port: '5432',
          userName: 'postgres',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: 'template0'
        })
        break
      case 'oracle':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'oracle',
          host: '127.0.0.1',
          port: '1521',
          userName: 'oracle',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
        break
      case 'mssql':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'mssql',
          host: '127.0.0.1',
          port: '1433',
          userName: 'mssql',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
        break
      case 'sqlite':
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'sqlite',
          host: '',
          port: '',
          userName: '',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
        break
      default:
        Object.assign(form, {
          adminPassword: '123456',
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'gva',
          dbPath: ''
        })
    }
  }
  const onSubmit = async () => {
    if (form.adminPassword.length < 6) {
      ElMessage({
        type: 'error',
        message: '密码长度不能小于6位'
      })
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text: '正在初始化数据库，请稍候',
      spinner: 'loading',
      background: 'rgba(0, 0, 0, 0.7)'
    })
    try {
      const res = await initDB(form)
      if (res.code === 0) {
        out.value = true
        ElMessage({
          type: 'success',
          message: res.msg
        })
        
        // 显示AI助手配置提示弹窗
        ElMessageBox.confirm(
          '已经完成基础数据库初始化！建议先进行编辑器AI助手配置，以获得更好的开发体验。',
          '配置完成',
          {
            confirmButtonText: '查看AI配置文档',
            cancelButtonText: '稍后配置',
            type: 'success',
            center: true
          }
        ).then(() => {
          // 点击确认按钮，打开AI配置文档
          window.open('https://www.gin-vue-admin.com/guide/server/mcp.html', '_blank')
          router.push({ name: 'Login' })
        }).catch(() => {
          // 点击取消按钮或关闭弹窗，直接跳转到登录页
          router.push({ name: 'Login' })
        })
      }
      loading.close()
    } catch (_) {
      loading.close()
    }
  }
</script>

<style lang="scss" scoped>
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
  /* 淡入上浮：替换原来的飞入动画 */
  .anim-fade-up {
    animation: anim-fade-up 0.45s cubic-bezier(0.22, 0.61, 0.36, 1) both;
  }
  @keyframes anim-fade-up {
    0% {
      opacity: 0;
      transform: translateY(18px);
    }
    100% {
      opacity: 1;
      transform: translateY(0);
    }
  }
  @media (max-width: 750px) {
    .form {
      width: 94vw !important;
      padding: 0;
    }
  }
</style>
