<template>
  <div class="init_page">
    <div class="init_page_panel">
      <div v-if="hello < 2" id="hello" :class="[hello < 1 ? 'slide-in-fwd-top' : 'slide-out-right']" class="hello">
        <div>
          <div class="hello_title">{{ $GIN_VUE_ADMIN.appName }}</div>
          <p class="in-two a-fadeinT">{{ t('init.note') }}</p>
          <p class="init_p">{{ t('init.note1') }}</p>
          <p class="init_p">{{ t('init.note2') }}</p>
          <p class="init_p">{{ t('init.note3') }}</p>
          <p class="init_p">{{ t('init.note4') }}<span style="color: red">innoDB</span></p>
          <p class="init_p">{{ t('init.note5') }}</p>
          <p class="init_btn">
            <el-button type="primary" @click="goDoc">
              {{ t('init.readDocs') }}
            </el-button>
            <el-button type="primary" @click="showNext">
              {{ t('init.confirm') }}
            </el-button>
          </p>
        </div>
      </div>
      <div
        v-if="hello > 0 "
        :class="[(hello > 0 && !out)? 'slide-in-left' : '' , out ? 'slide-out-right' : '']"
        class="form"
      >
        <el-form ref="formRef" :model="form" label-width="130px">
          <!-- added by mohamed hassan to support multilangauge -->
          <el-form-item :label="t('init.language')">
            <el-select v-model="form.language" :placeholder="t('general.pleaseSelect')" @change="changeLanguage">
              <el-option key="en" label="English" value="en"><img src="@/assets/flags/en.svg" class="img">English</el-option>
              <el-option key="zh" label="中文" value="zh"><img src="@/assets/flags/zh.svg" class="img">中文</el-option>
              <el-option key="ar" label="العربية" value="ar"><img src="@/assets/flags/ar.svg" class="img">العربية</el-option>
            </el-select>
          </el-form-item>
          <!-- end of adding -->
          <el-form-item :label="t('init.dbType')">
            <el-select v-model="form.dbType" :placeholder="t('general.pleaseSelect')" @change="changeDB">
              <el-option key="mysql" label="mysql" value="mysql" />
              <el-option key="pgsql" label="pgsql" value="pgsql" />
              <el-option key="oracle" label="oracle" value="oracle" />
              <el-option key="mssql" label="mssql" value="mssql" />
              <el-option key="sqlite" label="sqlite" value="sqlite"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="t('init.dbHost')" v-if="form.dbType !== 'sqlite'">
            <el-input v-model="form.host" :placeholder="t('init.enterDBHost')" />
          </el-form-item>
          <el-form-item :label="t('init.dbPort')" v-if="form.dbType !== 'sqlite'">
            <el-input v-model="form.port" :placeholder="t('init.enterDBPort')" />
          </el-form-item>
          <el-form-item :label="t('init.dbUsername')" v-if="form.dbType !== 'sqlite'">
            <el-input v-model="form.userName" :placeholder="t('init.enterDBUsername')" />
          </el-form-item>
          <el-form-item label="password" v-if="form.dbType !== 'sqlite'">
            <el-input v-model="form.password" :placeholder="t('init.enterDBPassword')" />
          </el-form-item>
          <el-form-item :label="t('init.dbName')">
            <el-input v-model="form.dbName" :placeholder="t('init.enterDBName')" />
          </el-form-item>
          <el-form-item label="dbPath" v-if="form.dbType === 'sqlite'">
            <el-input v-model="form.dbPath" placeholder="请输入sqlite数据库文件存放路径"></el-input>
          </el-form-item>
          <el-form-item>
            <div style="text-align: right">
              <el-button type="primary" @click="onSubmit">{{ t('init.initNow') }}</el-button>
            </div>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Init',
}
</script>

<script setup>
// @ts-ignore
import { initDB } from '@/api/initdb'
import { reactive, ref } from 'vue'
import { ElLoading, ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const i18n = useI18n() // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const router = useRouter()

const hello = ref(0)
const showNext = () => {
  hello.value = hello.value + 1
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

.init_page{
  margin: 0;
  padding: 0;
  background-image: url('@/assets/login_background.jpg');
  background-size: cover;
  width: 100%;
  height: 100%;
  position: relative;
  .init_page_panel {
    position: absolute;
    top: 3vh;
    left: 2vw;
    width: 96vw;
    height: 94vh;
    background-color: rgba(255, 255, 255, 0.8);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    .hello {
      position: absolute;
      z-index: 2;
      text-align: center;
      width: 100%;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      .hello_title {
        font-size: 32px;
        line-height: 98px;
      }
      .in-two {
        font-size: 22px;
      }
      .init_p {
        margin-top: 20px;
        color: #777777;
      }
      .init_btn {
        margin-top: 20px;
      }
    }
    .form {
      position: absolute;
      z-index: 3;
      margin-top: 60px;
      width: 600px;
      height: auto;
      padding: 20px;
      border-radius: 6px;
    }
  }
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
