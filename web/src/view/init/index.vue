<template>
  <div
    class="rounded-lg flex items-center justify-evenly w-full h-full relative md:w-screen md:h-screen md:bg-[#194bfb] overflow-hidden"
  >
    <div
      class="rounded-md w-full h-full flex items-center justify-center overflow-hidden"
    >
      <div
        class="oblique h-[130%] w-3/5 bg-white dark:bg-slate-900 transform -rotate-12 absolute -ml-80"
      />
      <div
        v-if="!page.showForm"
        :class="[page.showReadme ? 'slide-out-right' : 'slide-in-fwd-top']"
      >
        <div class="text-lg">
          <div
            class="font-sans text-4xl font-bold text-center mb-4 dark:text-white"
          >
            GIN-VUE-ADMIN
          </div>
          <p class="text-gray-600 dark:text-gray-300 mb-2">初始化须知</p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            1.您需有用一定的VUE和GOLANG基础
          </p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            2.请您确认是否已经阅读过<a
              class="text-blue-600 font-bold"
              href="https://www.gin-vue-admin.com"
              target="_blank"
              >官方文档</a
            >
            <a
              class="text-blue-600 font-bold"
              href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=2"
              target="_blank"
              >初始化视频</a
            >
          </p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            3.请您确认是否了解后续的配置流程
          </p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            4.如果您使用mysql数据库，请确认数据库引擎为<span
              class="text-red-600 font-bold text-3xl ml-2"
              >innoDB</span
            >
          </p>
          <p class="text-gray-600 dark:text-gray-300 mb-2">
            注：开发组不为文档中书写过的内容提供无偿服务
          </p>
          <p class="flex items-center justify-between mt-8">
            <el-button type="primary" size="large" @click="goDoc">
              阅读文档
            </el-button>
            <el-button type="primary" size="large" @click="showNext">
              我已确认
            </el-button>
          </p>
        </div>
      </div>
      <div
        v-if="page.showForm"
        :class="[page.showForm ? 'slide-in-left' : 'slide-out-right']"
        class="w-96"
      >
        <el-form ref="formRef" :model="form" label-width="100px" size="large">
          <el-form-item label="管理员密码">
            <el-input
              v-model="form.adminPassword"
              placeholder="admin账号的默认密码"
            ></el-input>
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
          <el-form-item v-if="form.dbType !== 'sqlite'" label="host">
            <el-input v-model="form.host" placeholder="请输入数据库链接" />
          </el-form-item>
          <el-form-item v-if="form.dbType !== 'sqlite'" label="port">
            <el-input v-model="form.port" placeholder="请输入数据库端口" />
          </el-form-item>
          <el-form-item v-if="form.dbType !== 'sqlite'" label="userName">
            <el-input
              v-model="form.userName"
              placeholder="请输入数据库用户名"
            />
          </el-form-item>
          <el-form-item v-if="form.dbType !== 'sqlite'" label="password">
            <el-input
              v-model="form.password"
              placeholder="请输入数据库密码（没有则为空）"
            />
          </el-form-item>
          <el-form-item label="dbName">
            <el-input v-model="form.dbName" placeholder="请输入数据库名称" />
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
            <div style="text-align: right">
              <el-button type="primary" @click="onSubmit">立即初始化</el-button>
            </div>
          </el-form-item>
        </el-form>
      </div>
    </div>

    <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]">
      <img class="h-full" src="@/assets/login_right_banner.jpg" alt="banner" />
    </div>
  </div>
</template>

<script setup>
  // @ts-ignore
  import { initDB } from '@/api/initdb'
  import { reactive, ref } from 'vue'
  import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
  import { useRouter } from 'vue-router'

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
  .slide-in-fwd-top {
    -webkit-animation: slide-in-fwd-top 0.4s
      cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
    animation: slide-in-fwd-top 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }
  .slide-out-right {
    -webkit-animation: slide-out-right 0.5s
      cubic-bezier(0.55, 0.085, 0.68, 0.53) both;
    animation: slide-out-right 0.5s cubic-bezier(0.55, 0.085, 0.68, 0.53) both;
  }
  .slide-in-left {
    -webkit-animation: slide-in-left 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94)
      both;
    animation: slide-in-left 0.5s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
  }
  @-webkit-keyframes slide-in-fwd-top {
    0% {
      transform: translateZ(-1400px) translateY(-800px);
      opacity: 0;
    }
    100% {
      transform: translateZ(0) translateY(0);
      opacity: 1;
    }
  }
  @keyframes slide-in-fwd-top {
    0% {
      transform: translateZ(-1400px) translateY(-800px);
      opacity: 0;
    }
    100% {
      transform: translateZ(0) translateY(0);
      opacity: 1;
    }
  }
  @-webkit-keyframes slide-out-right {
    0% {
      transform: translateX(0);
      opacity: 1;
    }
    100% {
      transform: translateX(1000px);
      opacity: 0;
    }
  }
  @keyframes slide-out-right {
    0% {
      transform: translateX(0);
      opacity: 1;
    }
    100% {
      transform: translateX(1000px);
      opacity: 0;
    }
  }
  @-webkit-keyframes slide-in-left {
    0% {
      transform: translateX(-1000px);
      opacity: 0;
    }
    100% {
      transform: translateX(0);
      opacity: 1;
    }
  }
  @keyframes slide-in-left {
    0% {
      transform: translateX(-1000px);
      opacity: 0;
    }
    100% {
      transform: translateX(0);
      opacity: 1;
    }
  }
  @media (max-width: 750px) {
    .form {
      width: 94vw !important;
      padding: 0;
    }
  }
</style>
