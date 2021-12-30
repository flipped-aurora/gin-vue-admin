<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-collapse v-model="activeNames">
        <el-collapse-item title="系统配置" name="1">
          <el-form-item label="环境值">
            <el-input v-model="config.system.env" />
          </el-form-item>
          <el-form-item label="端口值">
            <el-input v-model.number="config.system.addr" />
          </el-form-item>
          <el-form-item label="数据库类型">
            <el-select v-model="config.system.dbType" style="width:100%">
              <el-option value="mysql" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item label="Oss类型">
            <el-select v-model="config.system.ossType" style="width:100%">
              <el-option value="local" />
              <el-option value="qiniu" />
              <el-option value="tencent-cos" />
              <el-option value="aliyun-oss" />
              <el-option value="huawei-obs" />
            </el-select>
          </el-form-item>
          <el-form-item label="多点登录拦截">
            <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
          </el-form-item>
          <el-form-item label="限流次数">
            <el-input-number v-model.number="config.system.iplimitCount" />
          </el-form-item>
          <el-form-item label="限流时间">
            <el-input-number v-model.number="config.system.iplimitTime" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="jwt签名" name="2">
          <el-form-item label="jwt签名">
            <el-input v-model="config.jwt.signingKey" />
          </el-form-item>
          <el-form-item label="有效期（秒）">
            <el-input v-model="config.jwt.expiresTime" />
          </el-form-item>
          <el-form-item label="缓冲期（秒）">
            <el-input v-model="config.jwt.bufferTime" />
          </el-form-item>
          <el-form-item label="签发者">
            <el-input v-model="config.jwt.issuer" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Zap日志配置" name="3">
          <el-form-item label="级别">
            <el-input v-model.number="config.zap.level" />
          </el-form-item>
          <el-form-item label="输出">
            <el-input v-model="config.zap.format" />
          </el-form-item>
          <el-form-item label="日志前缀">
            <el-input v-model="config.zap.prefix" />
          </el-form-item>
          <el-form-item label="日志文件夹">
            <el-input v-model="config.zap.director" />
          </el-form-item>
          <el-form-item label="编码级">
            <el-input v-model="config.zap.encodeLevel" />
          </el-form-item>
          <el-form-item label="栈名">
            <el-input v-model="config.zap.stacktraceKey" />
          </el-form-item>
          <el-form-item label="显示行">
            <el-checkbox v-model="config.zap.showLine" />
          </el-form-item>
          <el-form-item label="输出控制台">
            <el-checkbox v-model="config.zap.logInConsole" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Redis admin数据库配置" name="4">
          <el-form-item label="db">
            <el-input v-model="config.redis.db" />
          </el-form-item>
          <el-form-item label="addr">
            <el-input v-model="config.redis.addr" />
          </el-form-item>
          <el-form-item label="password">
            <el-input v-model="config.redis.password" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="邮箱配置" name="5">
          <el-form-item label="接收者邮箱">
            <el-input v-model="config.email.to" placeholder="可多个，以逗号分隔" />
          </el-form-item>
          <el-form-item label="端口">
            <el-input v-model.number="config.email.port" />
          </el-form-item>
          <el-form-item label="发送者邮箱">
            <el-input v-model="config.email.from" />
          </el-form-item>
          <el-form-item label="host">
            <el-input v-model="config.email.host" />
          </el-form-item>
          <el-form-item label="是否为ssl">
            <el-checkbox v-model="config.email.isSSL" />
          </el-form-item>
          <el-form-item label="secret">
            <el-input v-model="config.email.secret" />
          </el-form-item>
          <el-form-item label="测试邮件">
            <el-button @click="email">测试邮件</el-button>
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="casbin配置" name="6">
          <el-form-item label="模型地址">
            <el-input v-model="config.casbin.modelPath" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="验证码配置" name="7">
          <el-form-item label="keyLong">
            <el-input v-model.number="config.captcha.keyLong" />
          </el-form-item>
          <el-form-item label="imgWidth">
            <el-input v-model.number="config.captcha.imgWidth" />
          </el-form-item>
          <el-form-item label="imgHeight">
            <el-input v-model.number="config.captcha.imgHeight" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="数据库配置" name="9">
          <template v-if="config.system.dbType === 'mysql'">
            <el-form-item label="username">
              <el-input v-model="config.mysql.username" />
            </el-form-item>
            <el-form-item label="password">
              <el-input v-model="config.mysql.password" />
            </el-form-item>
            <el-form-item label="path">
              <el-input v-model="config.mysql.path" />
            </el-form-item>
            <el-form-item label="dbname">
              <el-input v-model="config.mysql.dbname" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.mysql.maxIdleConns" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.mysql.maxOpenConns" />
            </el-form-item>
            <el-form-item label="logMode">
              <el-checkbox v-model="config.mysql.logMode" />
            </el-form-item>
          </template>
          <template v-if="config.system.dbType === 'pgsql'">
            <el-form-item label="username">
              <el-input v-model="config.pgsql.username" />
            </el-form-item>
            <el-form-item label="password">
              <el-input v-model="config.pgsql.password" />
            </el-form-item>
            <el-form-item label="path">
              <el-input v-model="config.pgsql.path" />
            </el-form-item>
            <el-form-item label="dbname">
              <el-input v-model="config.pgsql.dbname" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.pgsql.maxIdleConns" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.pgsql.maxOpenConns" />
            </el-form-item>
            <el-form-item label="logMode">
              <el-checkbox v-model="config.pgsql.logMode" />
            </el-form-item>
          </template>
        </el-collapse-item>

        <el-collapse-item title="oss配置" name="10">
          <template v-if="config.system.ossType === 'local'">
            <h2>本地文件配置</h2>
            <el-form-item label="本地文件路径">
              <el-input v-model="config.local.path" />
            </el-form-item>
          </template>
          <template v-if="config.system.ossType === 'qiniu'">
            <h2>qiniu上传配置</h2>
            <el-form-item label="存储区域">
              <el-input v-model="config.qiniu.zone" />
            </el-form-item>
            <el-form-item label="空间名称">
              <el-input v-model="config.qiniu.bucket" />
            </el-form-item>
            <el-form-item label="CDN加速域名">
              <el-input v-model="config.qiniu.imgPath" />
            </el-form-item>
            <el-form-item label="是否使用https">
              <el-checkbox v-model="config.qiniu.useHttps">开启</el-checkbox>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config.qiniu.accessKey" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.qiniu.secretKey" />
            </el-form-item>
            <el-form-item label="上传是否使用CDN上传加速">
              <el-checkbox v-model="config.qiniu.useCdnDomains">开启</el-checkbox>
            </el-form-item>
          </template>
          <template v-if="config.system.ossType === 'tencent-cos'">
            <h2>腾讯云COS上传配置</h2>
            <el-form-item label="bucket">
              <el-input v-model="config.tencentCOS.bucket" />
            </el-form-item>
            <el-form-item label="region">
              <el-input v-model="config.tencentCOS.region" />
            </el-form-item>
            <el-form-item label="secretID">
              <el-input v-model="config.tencentCOS.secretID" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.tencentCOS.secretKey" />
            </el-form-item>
            <el-form-item label="pathPrefix">
              <el-input v-model="config.tencentCOS.pathPrefix" />
            </el-form-item>
            <el-form-item label="baseURL">
              <el-input v-model="config.tencentCOS.baseURL" />
            </el-form-item>
          </template>
          <template v-if="config.system.ossType === 'aliyun-oss'">
            <h2>阿里云OSS上传配置</h2>
            <el-form-item label="endpoint">
              <el-input v-model="config.aliyunOSS.endpoint" />
            </el-form-item>
            <el-form-item label="accessKeyId">
              <el-input v-model="config.aliyunOSS.accessKeyId" />
            </el-form-item>
            <el-form-item label="accessKeySecret">
              <el-input v-model="config.aliyunOSS.accessKeySecret" />
            </el-form-item>
            <el-form-item label="bucketName">
              <el-input v-model="config.aliyunOSS.bucketName" />
            </el-form-item>
            <el-form-item label="bucketUrl">
              <el-input v-model="config.aliyunOSS.bucketUrl" />
            </el-form-item>
          </template>
          <template v-if="config.system.ossType === 'huawei-obs'">
            <h2>华为云Obs上传配置</h2>
            <el-form-item label="path">
              <el-input v-model="config.huaWeiObs.path" />
            </el-form-item>
            <el-form-item label="bucket">
              <el-input v-model="config.huaWeiObs.bucket" />
            </el-form-item>
            <el-form-item label="endpoint">
              <el-input v-model="config.huaWeiObs.endpoint" />
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config.huaWeiObs.AccessKey" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.huaWeiObs.secretKey" />
            </el-form-item>
          </template>

        </el-collapse-item>

        <el-collapse-item title="Excel上传配置" name="11">
          <el-form-item label="合成目标地址">
            <el-input v-model="config.excel.dir" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="自动化代码配置" name="12">
          <el-form-item label="是否自动重启(linux)">
            <el-checkbox v-model="config.autoCode.transferRestart" />
          </el-form-item>
          <el-form-item label="root(项目根路径)">
            <el-input v-model="config.autoCode.root" disabled />
          </el-form-item>
          <el-form-item label="Server(后端代码地址)">
            <el-input v-model="config.autoCode.transferRestart" />
          </el-form-item>
          <el-form-item label="SApi(后端api文件夹地址)">
            <el-input v-model="config.autoCode.serverApi" />
          </el-form-item>
          <el-form-item label="SInitialize(后端Initialize文件夹)">
            <el-input v-model="config.autoCode.serverInitialize" />
          </el-form-item>
          <el-form-item label="SModel(后端Model文件地址)">
            <el-input v-model="config.autoCode.serverModel" />
          </el-form-item>
          <el-form-item label="SRequest(后端Request文件夹地址)">
            <el-input v-model="config.autoCode.serverRequest" />
          </el-form-item>
          <el-form-item label="SRouter(后端Router文件夹地址)">
            <el-input v-model="config.autoCode.serverRouter" />
          </el-form-item>
          <el-form-item label="SService(后端Service文件夹地址)">
            <el-input v-model="config.autoCode.serverService" />
          </el-form-item>
          <el-form-item label="Web(前端文件夹地址)">
            <el-input v-model="config.autoCode.web" />
          </el-form-item>
          <el-form-item label="WApi(后端WApi文件夹地址)">
            <el-input v-model="config.autoCode.webApi" />
          </el-form-item>
          <el-form-item label="WForm(后端WForm文件夹地址)">
            <el-input v-model="config.autoCode.webForm" />
          </el-form-item>
          <el-form-item label="WTable(后端WTable文件夹地址)">
            <el-input v-model="config.autoCode.webTable" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="Timer(定时任务)" name="13">
          <el-form-item label="Start（是否启用）">
            <el-select v-model="config.timer.wTable" />
          </el-form-item>
          <el-form-item label="Spec(CRON表达式)">
            <el-input v-model="config.timer.spec" />
          </el-form-item>
          <template v-for="(item,k) in config.timer.detail">
            <div v-for="(key,k2) in item" :key="k2">
              <el-form-item :key="k+k2" :label="k2">
                <el-input v-model="item[k2]" />
              </el-form-item>
            </div>
          </template>
        </el-collapse-item>
      </el-collapse>
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" size="mini" @click="update">立即更新</el-button>
      <el-button type="primary" size="mini" @click="reload">重启服务（开发中）</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Config'
}
</script>
<script setup>
import { getSystemConfig, setSystemConfig } from '@/api/system'
import { emailTest } from '@/api/email'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

const activeNames = reactive([])
const config = ref({
  system: {
    iplimitCount: 0,
    iplimitTime: 0
  },
  jwt: {},
  casbin: {},
  mysql: {},
  pgsql: {},
  excel: {},
  autoCode: {},
  redis: {},
  qiniu: {},
  tencentCOS: {},
  aliyunOSS: {},
  huaWeiObs: {},
  captcha: {},
  zap: {},
  local: {},
  email: {},
  timer: {
    detail: {}
  }
})

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
}
initForm()
const reload = () => {}
const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置文件设置成功'
    })
    await initForm()
  }
}
const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '邮件发送成功'
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: '邮件发送失败'
    })
  }
}

</script>

<style lang="scss">
.system {
  background: #fff;
  padding:36px;
  border-radius: 2px;
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 0px 0px 0px #e7e8e8;
  }
  ::v-deep(.el-input-number__increase){
    top:5px !important;
  }
  .gva-btn-list{
    margin-top:16px;
  }
}
</style>
