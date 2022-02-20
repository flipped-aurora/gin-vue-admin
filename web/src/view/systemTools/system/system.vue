<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-collapse v-model="activeNames">
        <el-collapse-item title="System Configuration" name="1">
          <el-form-item label="Environmental value">
            <el-input v-model="config.system.env" />
          </el-form-item>
          <el-form-item label="Port value">
            <el-input v-model.number="config.system.addr" />
          </el-form-item>
          <el-form-item label="Database type">
            <el-select v-model="config.system.dbType" style="width:100%">
              <el-option value="mysql" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item label="Oss type">
            <el-select v-model="config.system.ossType" style="width:100%">
              <el-option value="local" />
              <el-option value="qiniu" />
              <el-option value="tencent-cos" />
              <el-option value="aliyun-oss" />
              <el-option value="huawei-obs" />
            </el-select>
          </el-form-item>
          <el-form-item label="Use Multipoint">
            <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
          </el-form-item>
          <el-form-item label="IP limit count">
            <el-input-number v-model.number="config.system.iplimitCount" />
          </el-form-item>
          <el-form-item label="IP limit time">
            <el-input-number v-model.number="config.system.iplimitTime" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="JWT signature" name="2">
          <el-form-item label="JWT signature">
            <el-input v-model="config.jwt.signingKey" />
          </el-form-item>
          <el-form-item label="Experies time">
            <el-input v-model="config.jwt.expiresTime" />
          </el-form-item>
          <el-form-item label="Buffer time">
            <el-input v-model="config.jwt.bufferTime" />
          </el-form-item>
          <el-form-item label="issuer">
            <el-input v-model="config.jwt.issuer" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="ZAP log configuration" name="3">
          <el-form-item label="Level">
            <el-input v-model.number="config.zap.level" />
          </el-form-item>
          <el-form-item label="Format">
            <el-input v-model="config.zap.format" />
          </el-form-item>
          <el-form-item label="Perfix">
            <el-input v-model="config.zap.prefix" />
          </el-form-item>
          <el-form-item label="Dicrector">
            <el-input v-model="config.zap.director" />
          </el-form-item>
          <el-form-item label="Encode level">
            <el-input v-model="config.zap.encodeLevel" />
          </el-form-item>
          <el-form-item label="Stack trace key">
            <el-input v-model="config.zap.stacktraceKey" />
          </el-form-item>
          <el-form-item label="Show line">
            <el-checkbox v-model="config.zap.showLine" />
          </el-form-item>
          <el-form-item label="Log in console">
            <el-checkbox v-model="config.zap.logInConsole" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Redis Admin Database Configuration" name="4">
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

        <el-collapse-item title="Mailbox configuration" name="5">
          <el-form-item label="Receiver mailbox">
            <el-input v-model="config.email.to" placeholder="可多个，以逗号分隔" />
          </el-form-item>
          <el-form-item label="Port">
            <el-input v-model.number="config.email.port" />
          </el-form-item>
          <el-form-item label="Sender mailbox">
            <el-input v-model="config.email.from" />
          </el-form-item>
          <el-form-item label="host">
            <el-input v-model="config.email.host" />
          </el-form-item>
          <el-form-item label="Is SSL">
            <el-checkbox v-model="config.email.isSSL" />
          </el-form-item>
          <el-form-item label="secret">
            <el-input v-model="config.email.secret" />
          </el-form-item>
          <el-form-item label="Test email">
            <el-button @click="email">Test email</el-button>
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item title="Casbin configuration" name="6">
          <el-form-item label="Model Address">
            <el-input v-model="config.casbin.modelPath" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="Verification code configuration" name="7">
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
        <el-collapse-item title="Database configuration" name="9">
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

        <el-collapse-item title="OSS configuration" name="10">
          <template v-if="config.system.ossType === 'local'">
            <h2>Local file configuration</h2>
            <el-form-item label="Local file path">
              <el-input v-model="config.local.path" />
            </el-form-item>
          </template>
          <template v-if="config.system.ossType === 'qiniu'">
            <h2>QiniU upload configuration</h2>
            <el-form-item label="Storage area">
              <el-input v-model="config.qiniu.zone" />
            </el-form-item>
            <el-form-item label="Spatial name">
              <el-input v-model="config.qiniu.bucket" />
            </el-form-item>
            <el-form-item label="CDN accelerated domain name">
              <el-input v-model="config.qiniu.imgPath" />
            </el-form-item>
            <el-form-item label="Use HTTPs">
              <el-checkbox v-model="config.qiniu.useHttps">开启</el-checkbox>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config.qiniu.accessKey" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.qiniu.secretKey" />
            </el-form-item>
            <el-form-item label="Use CDN domains">
              <el-checkbox v-model="config.qiniu.useCdnDomains">Open</el-checkbox>
            </el-form-item>
          </template>
          <template v-if="config.system.ossType === 'tencent-cos'">
            <h2>Tencent Cloud COS Upload Configuration</h2>
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
            <h2>Ali Cloud OSS upload configuration</h2>
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
            <h2>Huawei cloud OBS upload configuration</h2>
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

        <el-collapse-item title="Excel upload configuration" name="11">
          <el-form-item label="Synthetic target address">
            <el-input v-model="config.excel.dir" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="Automation code configuration" name="12">
          <el-form-item label="Whether it is automatically restarted (linux)">
            <el-checkbox v-model="config.autoCode.transferRestart" />
          </el-form-item>
          <el-form-item label="root (Project root path)">
            <el-input v-model="config.autoCode.root" disabled />
          </el-form-item>
          <el-form-item label="Server (Back end code address)">
            <el-input v-model="config.autoCode.transferRestart" />
          </el-form-item>
          <el-form-item label="SApi (Back end API folder address)">
            <el-input v-model="config.autoCode.serverApi" />
          </el-form-item>
          <el-form-item label="SInitialize (Back-end Initialize folder)">
            <el-input v-model="config.autoCode.serverInitialize" />
          </el-form-item>
          <el-form-item label="SModel (Back end MODEL file address)">
            <el-input v-model="config.autoCode.serverModel" />
          </el-form-item>
          <el-form-item label="SRequest (Back-end Request folder address)">
            <el-input v-model="config.autoCode.serverRequest" />
          </el-form-item>
          <el-form-item label="SRouter (Back end ROUTER folder address)">
            <el-input v-model="config.autoCode.serverRouter" />
          </el-form-item>
          <el-form-item label="SService (Back end service folder address)">
            <el-input v-model="config.autoCode.serverService" />
          </el-form-item>
          <el-form-item label="Web (Front end folder address)">
            <el-input v-model="config.autoCode.web" />
          </el-form-item>
          <el-form-item label="WApi (Back end WAPI folder address)">
            <el-input v-model="config.autoCode.webApi" />
          </el-form-item>
          <el-form-item label="WForm (Back end WFORM folder address)">
            <el-input v-model="config.autoCode.webForm" />
          </el-form-item>
          <el-form-item label="WTable (Back end WTABLE folder address)">
            <el-input v-model="config.autoCode.webTable" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item title="Timer (Timing task)" name="13">
          <el-form-item label="START (Is it enabled)">
            <el-select v-model="config.timer.wTable" />
          </el-form-item>
          <el-form-item label="Spec(Cron expression)">
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
      <el-button type="primary" size="small" @click="update">Update</el-button>
      <el-button type="primary" size="small" @click="reload">Restart service (in development)</el-button>
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
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

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
      message: 'The configuration file is set successfully'
    })
    await initForm()
  }
}
const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Mail sent successfully'
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: 'Mail delivery failed'
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
