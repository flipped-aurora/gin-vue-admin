<template>
  <div class="init_page">
    <div class="init_page_panel">
      <div v-if="hello < 2" id="hello" :class="[hello < 1 ? 'slide-in-fwd-top' : 'slide-out-right']" class="hello">
        <div>
          <div class="hello_title">GIN-VUE-ADMIN</div>
          <p class="in-two a-fadeinT">初始化须知</p>
          <p class="init_p">1.您需有用一定的VUE和GOLANG基础</p>
          <p class="init_p">2.请您确认是否已经阅读过官方文档</p>
          <p class="init_p">3.请您确认是否了解后续的配置流程</p>
          <p class="init_p">4.如果您使用mysql数据库，请确认数据库引擎为<span style="color: red">innoDB</span></p>
          <p class="init_p">注：开发组不为文档中书写过的内容提供无偿服务</p>
          <p class="init_btn">
            <el-button type="primary" @click="goDoc">
              阅读文档
            </el-button>
            <el-button type="primary" @click="showNext">
              我已确认
            </el-button>
          </p>
        </div>
      </div>
      <div
        v-if="hello > 0 "
        :class="[(hello > 0 && !out)? 'slide-in-left' : '' , out ? 'slide-out-right' : '']"
        class="form"
      >
        <el-form ref="formRef" :model="form" label-width="100px">
          <el-form-item label="数据库类型">
            <el-select v-model="form.dbType" placeholder="请选择" @change="changeDB">
              <el-option key="mysql" label="mysql" value="mysql" />
              <el-option key="pgsql" label="pgsql" value="pgsql" />
              <el-option key="oracle" label="oracle" value="oracle" />
              <el-option key="mssql" label="mssql" value="mssql" />
            </el-select>
          </el-form-item>
          <el-form-item label="host">
            <el-input v-model="form.host" placeholder="请输入数据库链接" />
          </el-form-item>
          <el-form-item label="port">
            <el-input v-model="form.port" placeholder="请输入数据库端口" />
          </el-form-item>
          <el-form-item label="userName">
            <el-input v-model="form.userName" placeholder="请输入数据库用户名" />
          </el-form-item>
          <el-form-item label="password">
            <el-input v-model="form.password" placeholder="请输入数据库密码（没有则为空）" />
          </el-form-item>
          <el-form-item label="dbName">
            <el-input v-model="form.dbName" placeholder="请输入数据库名称" />
          </el-form-item>
          <el-form-item>
            <div style="text-align: right">
              <el-button type="primary" @click="onSubmit">立即初始化</el-button>
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
      })
  }
}
const onSubmit = async() => {
  const loading = ElLoading.service({
    lock: true,
    text: '正在初始化数据库，请稍候',
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
</script>

<style lang="scss" scoped>
.init_page {
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
