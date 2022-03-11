<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-collapse v-model="activeNames">
        <el-collapse-item :title="t('view.systemTools.system.systemConfig')" name="1">
          <el-form-item :label="t('view.systemTools.system.envValue')">
            <el-input v-model="config.system.env" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.portValue')">
            <el-input v-model.number="config.system.addr" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.dbType')">
            <el-select v-model="config.system.dbType" style="width:100%">
              <el-option value="mysql" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ossType')">
            <el-select v-model="config.system.ossType" style="width:100%">
              <el-option value="local" />
              <el-option value="qiniu" />
              <el-option value="tencent-cos" />
              <el-option value="aliyun-oss" />
              <el-option value="huawei-obs" />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.blockMultiSignOn')">
            <el-checkbox v-model="config.system.useMultipoint">{{ t('general.enable') }}</el-checkbox>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.enableRedis')">
            <el-checkbox v-model="config.system.useRedis">{{ t('general.enable') }}</el-checkbox>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ipLimitCount')">
            <el-input-number v-model.number="config.system.iplimitCount" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ipLimitTime')">
            <el-input-number v-model.number="config.system.iplimitTime" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.jwtSignature')" name="2">
          <el-form-item :label="t('view.systemTools.system.jwtSignature')">
            <el-input v-model="config.jwt.signingKey" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.expirartionSec')">
            <el-input v-model="config.jwt.expiresTime" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.bufferPeriodSec')">
            <el-input v-model="config.jwt.bufferTime" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.issuer')">
            <el-input v-model="config.jwt.issuer" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.zapLogConfig')" name="3">
          <el-form-item :label="t('view.systemTools.system.level')">
            <el-input v-model.number="config.zap.level" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.output')">
            <el-input v-model="config.zap.format" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.logPrefix')">
            <el-input v-model="config.zap.prefix" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.logFolder')">
            <el-input v-model="config.zap.director" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.encodeLevel')">
            <el-input v-model="config.zap.encodeLevel" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.stackName')">
            <el-input v-model="config.zap.stacktraceKey" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.showLine')">
            <el-checkbox v-model="config.zap.showLine" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.outputConsole')">
            <el-checkbox v-model="config.zap.logInConsole" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.redisAdminDBConfig')" name="4">
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

        <el-collapse-item :title="t('view.systemTools.system.emailConfig')" name="5">
          <el-form-item :label="t('view.systemTools.system.recipientEmail')">
            <el-input v-model="config.email.to" :placeholder="t('view.systemTools.system.emailNote')" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.port')">
            <el-input v-model.number="config.email.port" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.emailSender')">
            <el-input v-model="config.email.from" />
          </el-form-item>
          <el-form-item label="Host">
            <el-input v-model="config.email.host" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.enableSSL')">
            <el-checkbox v-model="config.email.isSSL" />
          </el-form-item>
          <el-form-item label="Secret">
            <el-input v-model="config.email.secret" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.testEmail')">
            <el-button @click="email">{{ t('view.systemTools.system.testEmail') }}</el-button>
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.casbinConfig')" name="6">
          <el-form-item :label="t('view.systemTools.system.modelAddress')">
            <el-input v-model="config.casbin.modelPath" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item :title="t('view.systemTools.system.verCodeConfig')" name="7">
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
        <el-collapse-item :title="t('view.systemTools.system.dbConfig')" name="9">
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

        <el-collapse-item :title="t('view.systemTools.system.ossConfig')" name="10">
          <template v-if="config.system.ossType === 'local'">
            <h2>{{ t('view.systemTools.system.localFileConfig') }}</h2>
            <el-form-item :label="t('view.systemTools.system.localFilePath')">
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
              <el-checkbox v-model="config.qiniu.useHttps">{{ t('general.enable') }}</el-checkbox>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config.qiniu.accessKey" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.qiniu.secretKey" />
            </el-form-item>
            <el-form-item label="上传是否使用CDN上传加速">
              <el-checkbox v-model="config.qiniu.useCdnDomains">{{ t('general.enable') }}</el-checkbox>
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

        <el-collapse-item :title="t('view.systemTools.system.excelUploadConfig')" name="11">
          <el-form-item :label="t('view.systemTools.system.excelDir')">
            <el-input v-model="config.excel.dir" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item :title="t('view.systemTools.system.autoCodeConfig')" name="12">
          <el-form-item :label="t('view.systemTools.system.autoRestart')">
            <el-checkbox v-model="config.autoCode.transferRestart" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.projectRootPath')">
            <el-input v-model="config.autoCode.root" disabled />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendCodePath')">
            <el-input v-model="config.autoCode.transferRestart" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendApiPath')">
            <el-input v-model="config.autoCode.serverApi" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendInitPath')">
            <el-input v-model="config.autoCode.serverInitialize" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendModelPath')">
            <el-input v-model="config.autoCode.serverModel" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRequestPath')">
            <el-input v-model="config.autoCode.serverRequest" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRouterPath')">
            <el-input v-model="config.autoCode.serverRouter" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendServicePath')">
            <el-input v-model="config.autoCode.serverService" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendCodePath')">
            <el-input v-model="config.autoCode.web" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendApiPath')">
            <el-input v-model="config.autoCode.webApi" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendFormPath')">
            <el-input v-model="config.autoCode.webForm" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendTablePath')">
            <el-input v-model="config.autoCode.webTable" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item :title="t('view.systemTools.system.timedTask')" name="13">
          <el-form-item :label="t('view.systemTools.system.startEnableDisable')">
            <el-select v-model="config.timer.wTable" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.cronExp')">
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

        <el-collapse-item :title="t('view.systemTools.system.i18n')" name="14">
          <el-form-item :label="t('view.systemTools.system.langFilesPath')">
            <el-input v-model="config.language.dir" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.language')">
            <el-input v-model="config.language.language" />
          </el-form-item>
        </el-collapse-item>
      </el-collapse>
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" size="small" @click="update">{{ t('view.systemTools.system.updateNow') }}</el-button>
      <el-button type="primary" size="small" @click="reload">{{ t('view.systemTools.system.restartService') }}</el-button>
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
  },
  language: {}
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
      message: t('view.systemTools.system.configSetupSuccess')
    })
    await initForm()
  }
}
const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: t('view.systemTools.system.emailSentSuccess')
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: t('view.systemTools.system.emailSentError')
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
