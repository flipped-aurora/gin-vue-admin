<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="240px">
      <!--  System start  -->
      <el-collapse v-model="activeNames">
        <el-collapse-item :title="t('view.systemTools.system.systemConfig')" name="1">
          <el-form-item :label="t('view.systemTools.system.envValue')">
            <!-- <el-input v-model="config.system.env" />-->
            <el-select v-model="config.system.env" style="width:100%">
              <el-option value="public" />
              <el-option value="develop" />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.portValue')">
            <el-input v-model.number="config.system.addr" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.dbType')">
            <el-select v-model="config.system['db-type']" style="width:100%">
              <el-option value="mysql" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ossType')">
            <el-select v-model="config.system['oss-type']" style="width:100%">
              <el-option value="local" />
              <el-option value="qiniu" />
              <el-option value="tencent-cos" />
              <el-option value="aliyun-oss" />
              <el-option value="huawei-obs" />
            </el-select>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.blockMultiSignOn')">
            <el-checkbox v-model="config.system['use-multipoint']">{{ t('general.enable') }}</el-checkbox>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.enableRedis')">
            <el-checkbox v-model="config.system['use-redis']">{{ t('general.enable') }}</el-checkbox>
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ipLimitCount')">
            <el-input-number v-model.number="config.system['iplimit-count']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.ipLimitTime')">
            <el-input-number v-model.number="config.system['iplimit-time']" />
          </el-form-item>
          <el-tooltip
            content="请修改完成后，注意一并修改前端env环境下的VITE_BASE_PATH"
            placement="top-start"
          >
            <el-form-item label="全局路由前缀">
              <el-input v-model="config.system['router-prefix']" />
            </el-form-item>
          </el-tooltip>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.jwtSignature')" name="2">
          <el-form-item :label="t('view.systemTools.system.jwtSignature')">
            <el-input v-model="config.jwt['signing-key']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.expirartionSec')">
            <el-input v-model="config.jwt['expires-time']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.bufferPeriodSec')">
            <el-input v-model="config.jwt['buffer-time']" />
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
            <el-input v-model="config.zap['encode-level']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.stackName')">
            <el-input v-model="config.zap['stacktrace-key']" />
          </el-form-item>
          <el-form-item label="日志留存时间(默认以天为单位)">
            <el-input v-model.number="config.zap['max-age']" />
          </el-form-item>
          <el-form-item label="显示行">
            <el-checkbox v-model="config.zap['show-line']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.outputConsole')">
            <el-checkbox v-model="config.zap['log-in-console']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.redisAdminDBConfig')" name="4">
          <el-form-item label="库">
            <el-input v-model.number="config.redis.db" />
          </el-form-item>
          <el-form-item label="地址">
            <el-input v-model="config.redis.addr" />
          </el-form-item>
          <el-form-item label="密码">
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
            <el-checkbox v-model="config.email['is-ssl']" />
          </el-form-item>
          <el-form-item label="Secret">
            <el-input v-model="config.email.secret" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.testEmail')">
            <el-button @click="email">{{ t('view.systemTools.system.testEmail') }}</el-button>
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.verCodeConfig')" name="7">
          <el-form-item label="字符长度">
            <el-input v-model.number="config.captcha['key-long']" />
          </el-form-item>
          <el-form-item label="平台宽度">
            <el-input v-model.number="config.captcha['img-width']" />
          </el-form-item>
          <el-form-item label="图片高度">
            <el-input v-model.number="config.captcha['img-height']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item :title="t('view.systemTools.system.dbConfig')" name="9">
          <template v-if="config.system['db-type'] === 'mysql'">
            <el-form-item label="用户名">
              <el-input v-model="config.mysql.username" />
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="config.mysql.password" />
            </el-form-item>
            <el-form-item label="地址">
              <el-input v-model="config.mysql.path" />
            </el-form-item>
            <el-form-item label="数据库">
              <el-input v-model="config.mysql['db-name']" />
            </el-form-item>
            <el-form-item label="前缀">
              <el-input v-model="config.mysql['refix']" />
            </el-form-item>
            <el-form-item label="复数表">
              <el-switch v-model="config.mysql['singular']" />
            </el-form-item>
            <el-form-item label="引擎">
              <el-input v-model="config.mysql['engine']" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.mysql['max-idle-conns']" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.mysql['max-open-conns']" />
            </el-form-item>
            <el-form-item label="写入日志">
              <el-checkbox v-model="config.mysql['log-zap']" />
            </el-form-item>
            <el-form-item label="日志模式">
              <el-input v-model="config.mysql['log-mode']" />
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'pgsql'">
            <el-form-item label="用户名">
              <el-input v-model="config.pgsql.username" />
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="config.pgsql.password" />
            </el-form-item>
            <el-form-item label="地址">
              <el-input v-model="config.pgsql.path" />
            </el-form-item>
            <el-form-item label="数据库">
              <el-input v-model="config.pgsql.dbname" />
            </el-form-item>
            <el-form-item label="前缀">
              <el-input v-model="config.pgsql['refix']" />
            </el-form-item>
            <el-form-item label="复数表">
              <el-switch v-model="config.pgsql['singular']" />
            </el-form-item>
            <el-form-item label="引擎">
              <el-input v-model="config.pgsql['engine']" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.pgsql['max-idle-conns']" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.pgsql['max-open-conns']" />
            </el-form-item>
            <el-form-item label="写入日志">
              <el-checkbox v-model="config.pgsql['log-zap']" />
            </el-form-item>
            <el-form-item label="日志模式">
              <el-input v-model="config.pgsql['log-mode']" />
            </el-form-item>
          </template>
        </el-collapse-item>

        <el-collapse-item :title="t('view.systemTools.system.ossConfig')" name="10">
          <template v-if="config.system['oss-type'] === 'local'">
            <h2>{{ t('view.systemTools.system.localFileConfig') }}</h2>
            <el-form-item :label="t('view.systemTools.system.localFilePath')">
              <el-input v-model="config.local.path" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'qiniu'">
            <h2>qiniu上传配置</h2>
            <el-form-item label="存储区域">
              <el-input v-model="config.qiniu.zone" />
            </el-form-item>
            <el-form-item label="空间名称">
              <el-input v-model="config.qiniu.bucket" />
            </el-form-item>
            <el-form-item label="CDN加速域名">
              <el-input v-model="config.qiniu['img-path']" />
            </el-form-item>
            <el-form-item label="是否使用https">
              <el-checkbox v-model="config.qiniu['use-https']">{{ t('general.enable') }}</el-checkbox>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config.qiniu['access-key']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.qiniu['secret-key']" />
            </el-form-item>
            <el-form-item label="上传是否使用CDN上传加速">
              <el-checkbox v-model="config.qiniu['use-cdn-domains']">{{ t('general.enable') }}</el-checkbox>
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'tencent-cos'">
            <h2>腾讯云COS上传配置</h2>
            <el-form-item label="存储桶名称">
              <el-input v-model="config['tencent-cos']['bucket']" />
            </el-form-item>
            <el-form-item label="所属地域">
              <el-input v-model="config['tencent-cos'].region" />
            </el-form-item>
            <el-form-item label="secretID">
              <el-input v-model="config['tencent-cos']['secret-id']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config['tencent-cos']['secret-key']" />
            </el-form-item>
            <el-form-item label="路径前缀">
              <el-input v-model="config['tencent-cos']['path-prefix']" />
            </el-form-item>
            <el-form-item label="访问域名">
              <el-input v-model="config['tencent-cos']['base-url']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'aliyun-oss'">
            <h2>阿里云OSS上传配置</h2>
            <el-form-item label="区域">
              <el-input v-model="config['aliyun-oss'].endpoint" />
            </el-form-item>
            <el-form-item label="accessKeyId">
              <el-input v-model="config['aliyun-oss']['access-key-id']" />
            </el-form-item>
            <el-form-item label="accessKeySecret">
              <el-input v-model="config['aliyun-oss']['access-key-secret']" />
            </el-form-item>
            <el-form-item label="存储桶名称">
              <el-input v-model="config['aliyun-oss']['bucket-name']" />
            </el-form-item>
            <el-form-item label="访问域名">
              <el-input v-model="config['aliyun-oss']['bucket-url']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'huawei-obs'">
            <h2>华为云Obs上传配置</h2>
            <el-form-item label="路径">
              <el-input v-model="config['hua-wei-obs'].path" />
            </el-form-item>
            <el-form-item label="存储桶名称">
              <el-input v-model="config['hua-wei-obs'].bucket" />
            </el-form-item>
            <el-form-item label="区域">
              <el-input v-model="config['hua-wei-obs'].endpoint" />
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config['hua-wei-obs']['access-key']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config['hua-wei-obs']['secret-key']" />
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
            <el-checkbox v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.projectRootPath')">
            <el-input v-model="config.autocode.root" disabled />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendCodePath')">
            <el-input v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendApiPath')">
            <el-input v-model="config.autocode['server-api']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendInitPath')">
            <el-input v-model="config.autocode['server-initialize']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendModelPath')">
            <el-input v-model="config.autocode['server-model']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRequestPath')">
            <el-input v-model="config.autocode['server-request']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendRouterPath')">
            <el-input v-model="config.autocode['server-router']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.backendServicePath')">
            <el-input v-model="config.autocode['server-service']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendCodePath')">
            <el-input v-model="config.autocode.web" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendApiPath')">
            <el-input v-model="config.autocode['web-api']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendFormPath')">
            <el-input v-model="config.autocode['web-form']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.frontendTablePath')">
            <el-input v-model="config.autocode['web-table']" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item :title="t('view.systemTools.system.timedTask')" name="13">
          <el-form-item :label="t('view.systemTools.system.startEnableDisable')">
            <el-checkbox v-model="config.timer['start']" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.system.cronExp')">
            <el-input v-model="config.timer.spec" />
          </el-form-item>
          <template v-for="(item,k) in config.timer.detail">
            <div v-for="(_,k2) in item" :key="k2">
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
      <el-button type="primary" @click="update">{{ t('view.systemTools.system.updateNow') }}</el-button>
      <el-button type="primary" @click="reload">{{ t('view.systemTools.system.restartService') }}</el-button>
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
    'iplimit-count': 0,
    'iplimit-time': 0
  },
  jwt: {},
  mysql: {},
  pgsql: {},
  excel: {},
  autocode: {},
  redis: {},
  qiniu: {},
  'tencent-cos': {},
  'aliyun-oss': {},
  'hua-wei-obs': {},
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
