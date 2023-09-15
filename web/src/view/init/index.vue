<template>
  <div class="rounded-lg  flex items-center justify-evenly w-full h-full relative bg-white md:w-screen md:h-screen md:bg-[#194bfb] overflow-hidden">
    <div class="rounded-md w-full h-full flex items-center justify-center overflow-hidden">
      <div class="oblique h-[130%] w-3/5 bg-white transform -rotate-12 absolute -ml-80" />
      <div
        v-if="!page.showForm"
        :class="[page.showReadme ?'slide-out-right' :'slide-in-fwd-top' ]"
      >
        <div class=" text-lg">
          <div class="font-sans text-4xl font-bold text-center mb-4">{{ $GIN_VUE_ADMIN.appName }}</div>
          <p class="text-gray-600 mb-2">{{ t('init.note') }}</p>
          <p class="text-gray-600 mb-2">{{ t('init.note1') }}</p>
          <p class="text-gray-600 mb-2">{{ t('init.note2-1') }}<a
            class="text-blue-600 font-bold"
            href="https://www.gin-vue-admin.com"
            target="_blank"
          >{{ t('init.note2-2') }}</a> <a
            class="text-blue-600 font-bold"
            href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=2"
            target="_blank"
          >{{ t('init.note2-3') }}</a></p>
          <p class="text-gray-600 mb-2">{{ t('init.note3') }}</p>
          <p class="text-gray-600 mb-2">{{ t('init.note4') }}<span class="text-red-600 font-bold text-3xl ml-2 ">innoDB</span></p>
          <p class="text-gray-600 mb-2">{{ t('init.note5') }}</p>
          <p class="flex items-center justify-between mt-8">
            <el-button
              type="primary"
              size="large"
              @click="goDoc"
            >
            {{ t('init.readDocs') }}
            </el-button>
            <el-button
              type="primary"
              size="large"
              @click="showNext"
            >
            {{ t('init.confirm') }}
            </el-button>
          </p>
        </div>
      </div>
      <div
        v-if="page.showForm "
        :class="[ page.showForm ? 'slide-in-left' : 'slide-out-right' ]"
        class="w-96"
      >
        <el-form
          ref="formRef"
          :model="form"
          label-width="100px"
          size="large"
        >
          <!-- added by mohamed hassan to support multilangauge -->
          <el-form-item :label="t('init.language')">
            <el-select v-model="form.language"
:placeholder="t('general.pleaseSelect')" @change="changeLanguage"
>
              <el-option key="en"
label="English" value="en"
><img src="@/assets/flags/en.svg"
class="img"
>English</el-option>
              <el-option key="zh"
label="中文" value="zh"
><img src="@/assets/flags/zh.svg"
class="img"
>中文</el-option>
              <el-option key="ar"
label="العربية" value="ar"
><img src="@/assets/flags/ar.svg"
class="img"
>العربية</el-option>
            </el-select>
          </el-form-item>
          <!-- end of adding -->
          <el-form-item :label="t('init.dbType')">
            <el-select
              v-model="form.dbType"
              :placeholder="t('general.pleaseSelect')"
              class="w-full"
              @change="changeDB"
            >
              <el-option
                key="mysql"
                label="mysql"
                value="mysql"
              />
              <el-option
                key="pgsql"
                label="pgsql"
                value="pgsql"
              />
              <el-option
                key="oracle"
                label="oracle"
                value="oracle"
              />
              <el-option
                key="mssql"
                label="mssql"
                value="mssql"
              />
              <el-option
                key="sqlite"
                label="sqlite"
                value="sqlite"
              />
            </el-select>
          </el-form-item>
          <el-form-item
            v-if="form.dbType !== 'sqlite'"
            :label="t('init.dbHost')"
          >
            <el-input
              v-model="form.host"
              :placeholder="t('init.enterDBHost')"
            />
          </el-form-item>
          <el-form-item
            v-if="form.dbType !== 'sqlite'"
            :label="t('init.dbPort')"
          >
            <el-input
              v-model="form.port"
              :placeholder="t('init.enterDBPort')"
            />
          </el-form-item>
          <el-form-item
            v-if="form.dbType !== 'sqlite'"
            :label="t('init.dbUsername')"
          >
            <el-input
              v-model="form.userName"
              :placeholder="t('init.enterDBUsername')"
            />
          </el-form-item>
          <el-form-item
            v-if="form.dbType !== 'sqlite'"
            label="password"
          >
            <el-input
              v-model="form.password"
              :placeholder="t('init.enterDBPassword')"
            />
          </el-form-item>
          <el-form-item :label="t('init.dbName')">
            <el-input
              v-model="form.dbName"
              :placeholder="t('init.enterDBName')"
            />
          </el-form-item>
          <el-form-item
            v-if="form.dbType === 'sqlite'"
            label="dbPath"
          >
            <el-input
              v-model="form.dbPath"
              placeholder="请输入sqlite数据库文件存放路径"
            />
          </el-form-item>
          <el-form-item>
            <div style="text-align: right">
              <el-button
                type="primary"
                @click="onSubmit"
              >{{ t('init.initNow') }}</el-button>
            </div>
          </el-form-item>
        </el-form>
      </div>
    </div>

    <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]"><img
      class="h-full"
      src="@/assets/login_right_banner.jpg"
      alt="banner"
    ></div>
  </div>
</template>

<script setup>
// @ts-ignore
import { initDB } from '@/api/initdb'
import { reactive, ref } from 'vue'
import { ElLoading, ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const i18n = useI18n() // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: 'Init',
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
  dbType: 'mysql',
  host: '127.0.0.1',
  port: '3306',
  userName: 'root',
  password: '',
  dbName: 'gva',
  dbPath: '',
  language: 'en' // added by mohamed hassan to support multilanguage
})

const changeDB = (val) => {
  switch (val) {
    case 'mysql':
      Object.assign(form, {
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
        dbType: 'pgsql',
        host: '127.0.0.1',
        port: '5432',
        userName: 'postgres',
        password: '',
        dbName: 'gva',
        dbPath: ''
      })
      break
    case 'oracle':
      Object.assign(form, {
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
const onSubmit = async() => {
  const loading = ElLoading.service({
    lock: true,
    text: t('init.pleaseWait'),
    spinner: 'loading',
    background: 'rgba(0, 0, 0, 0.7)',
  })
  try {
    const res = await initDB(form)
    if (res.code === 0) {
      out.value = true
      ElMessage({
        type: 'success',
        message: res.msg,
      })
      router.push({ name: 'Login' })
    }
    loading.close()
  } catch (err) {
    loading.close()
  }
}
// added by mohamed hassan to support multilanguage
const changeLanguage = (val) => {
  i18n.locale.value = val
}
</script>

<style lang="scss" scoped>

img {
  padding-right: 20px;
  width: 20px;
  height: 20px;
}

.slide-in-fwd-top {
  -webkit-animation: slide-in-fwd-top 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94)
    both;
  animation: slide-in-fwd-top 0.4s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
}
.slide-out-right {
  -webkit-animation: slide-out-right 0.5s cubic-bezier(0.55, 0.085, 0.68, 0.53)
    both;
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
