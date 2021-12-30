<template>
  <div class="init_page">
    <div class="init_page_panle">
      <div v-if="hello < 2" id="hello" :class="[hello < 1 ? 'slide-in-fwd-top' : 'slide-out-right']" class="hello  ">
        <div>
          <div class="hello_title">{{ $GIN_VUE_ADMIN.appName }}</div>
          <p class="in-two a-fadeinT">{{ $t('init.note') }}</p>
          <p class="init_p">{{ $t('init.note1') }}</p>
          <p class="init_p">{{ $t('init.note2') }}</p>
          <p class="init_p">{{ $t('init.note3') }}</p>
          <p class="init_p">{{ $t('init.note4') }}</p>
          <p class="init_btn">
            <el-button type="primary" @click="goDoc">
              {{ $t('init.readDocs') }}
            </el-button>
            <el-button type="primary" @click="showNext">
              {{ $t('init.confirm') }}
            </el-button>
          </p>
        </div>
      </div>
      <div v-if="hello > 0 " :class="[(hello > 0 && !out)? 'slide-in-left' : '' , out ? 'slide-out-right' : '']" class=" form">
        <el-form ref="form" :model="form" label-width="120px">
          <el-form-item :label="$t('init.dbType')">
            <el-select v-model="form.dbType" placeholder="$('general.pleaseSelect')" @change="changeDB">
              <el-option key="mysql" label="MySQL" value="mysql" />
              <el-option key="pgsql" :label="'PostgreSQL (' + $t('init.beta') + ')'" value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('init.dbHost')">
            <el-input v-model="form.host" :placeholder="$t('init.enterDBHost')" />
          </el-form-item>
          <el-form-item :label="$t('init.dbPort')">
            <el-input v-model="form.port" :placeholder="$t('init.enterDBPort')" />
          </el-form-item>
          <el-form-item :label="$t('init.dbUsername')">
            <el-input v-model="form.userName" :placeholder="$t('init.enterDBUsername')" />
          </el-form-item>
          <el-form-item :label="$t('init.dbPassword')">
            <el-input
              v-model="form.password"
              :placeholder="$t('init.enterDBPassword')"
            />
          </el-form-item>
          <el-form-item :label="$t('init.dbName')">
            <el-input v-model="form.dbName" :placeholder="$t('init.enterDBName')" />
          </el-form-item>
          <el-form-item>
            <div style="text-align: right">
              <el-button type="primary" @click="onSubmit">{{ $t('init.initNow') }}</el-button>
            </div>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>
<script>
import { initDB } from '@/api/initdb'
export default {
  name: 'Init',
  data() {
    return {
      hello: 0,
      out: false,
      form: {
        dbType: 'mysql',
        host: '127.0.0.1',
        port: '3306',
        userName: 'root',
        password: '',
        dbName: 'gva'
      }
    }
  },
  methods: {
    changeDB(val) {
      switch (val) {
        case 'mysql':
          this.form = {
            dbType: 'mysql',
            host: '127.0.0.1',
            port: '3306',
            userName: 'root',
            password: '',
            dbName: 'gva'
          }
          break
        case 'pgsql':
          this.form = {
            dbType: 'pgsql',
            host: '127.0.0.1',
            port: '5432',
            userName: 'postgres',
            password: '',
            dbName: 'gva'
          }
          break
        default:
          this.form = {
            dbType: 'mysql',
            host: '127.0.0.1',
            port: '3306',
            userName: 'root',
            password: '',
            dbName: 'gva'
          }
      }
    },
    showNext() {
      this.hello = this.hello + 1
    },
    goDoc() {
      window.open('https://www.gin-vue-admin.com/docs/first_master#3-init')
    },
    async onSubmit() {
      const loading = this.$loading({
        lock: true,
        text: this.$t('init.pleaseWait'),
        spinner: 'loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      try {
        const res = await initDB(this.form)
        if (res.code === 0) {
          this.out = true
          this.$message({
            type: 'success',
            message: res.msg
          })
          this.$router.push({ name: 'Login' })
        }
        loading.close()
      } catch (err) {
        loading.close()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.init_page{
  margin: 0;
  padding: 0;
  background-image: url("@/assets/login_background.jpg");
  background-size: cover;
  width: 100%;
  height: 100%;
  position: relative;
  .init_page_panle{
    position: absolute;
    top: 3vh;
    left: 2vw;
    width: 96vw;
    height: 94vh;
    background-color: rgba(255,255,255,.8);
    backdrop-filter: blur(5px);
    border-radius: 10px;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    .hello{
      position: absolute;
      z-index: 2;
      text-align: center;
      width: 100%;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
      cursor: pointer;
      .hello_title{
        font-size: 32px;
        line-height: 98px;
      }
      .in-two{
        font-size: 22px;
      }
      .init_p{
        margin-top: 20px;
        color: #777777;
      }
      .init_btn{
        margin-top: 20px;
      }
    }
    .form{
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
  -webkit-animation: slide-in-fwd-top 0.4s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
  animation: slide-in-fwd-top 0.4s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
}
.slide-out-right {
  -webkit-animation: slide-out-right 0.5s cubic-bezier(0.550, 0.085, 0.680, 0.530) both;
  animation: slide-out-right 0.5s cubic-bezier(0.550, 0.085, 0.680, 0.530) both;
}
.slide-in-left {
  -webkit-animation: slide-in-left 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
  animation: slide-in-left 0.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both;
}
@-webkit-keyframes slide-in-fwd-top {
  0% {
    -webkit-transform: translateZ(-1400px) translateY(-800px);
    transform: translateZ(-1400px) translateY(-800px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateZ(0) translateY(0);
    transform: translateZ(0) translateY(0);
    opacity: 1;
  }
}
@keyframes slide-in-fwd-top {
  0% {
    -webkit-transform: translateZ(-1400px) translateY(-800px);
    transform: translateZ(-1400px) translateY(-800px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateZ(0) translateY(0);
    transform: translateZ(0) translateY(0);
    opacity: 1;
  }
}
@-webkit-keyframes slide-out-right {
  0% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
  100% {
    -webkit-transform: translateX(1000px);
    transform: translateX(1000px);
    opacity: 0;
  }
}
@keyframes slide-out-right {
  0% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
  100% {
    -webkit-transform: translateX(1000px);
    transform: translateX(1000px);
    opacity: 0;
  }
}
@-webkit-keyframes slide-in-left {
  0% {
    -webkit-transform: translateX(-1000px);
    transform: translateX(-1000px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
}
@keyframes slide-in-left {
  0% {
    -webkit-transform: translateX(-1000px);
    transform: translateX(-1000px);
    opacity: 0;
  }
  100% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    opacity: 1;
  }
}
@media (max-width: 750px) {
  .form{
    width: 94vw !important;
    padding: 0;
  }
}

</style>
