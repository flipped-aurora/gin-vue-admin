<template>
  <div class="system">
    <el-form ref="form" :model="config" label-width="100px">
      <!--  System start  -->
      <h2>系统配置</h2>
      <el-form-item label="环境值">
        <el-input v-model="config.system.env" />
      </el-form-item>
      <el-form-item label="端口值">
        <el-input v-model.number="config.system.addr" />
      </el-form-item>
      <el-form-item label="数据库类型">
        <el-select v-model="config.system.dbType" style="width:100%">
          <el-option value="mysql" />
          <el-option value="sqlite" />
          <el-option value="sqlserver" />
          <el-option value="postgresql" />
        </el-select>
      </el-form-item>
      <el-form-item label="Oss类型">
        <el-select v-model="config.system.ossType" style="width:100%">
          <el-option value="local" />
          <el-option value="qiniu" />
          <el-option value="tencent-cos" />
          <el-option value="aliyun-oss" />
        </el-select>
      </el-form-item>
      <el-form-item label="配置文件环境变量名">
        <el-input v-model.number="config.system.configEnv" />
      </el-form-item>
      <el-form-item label="数据初始化">
        <el-checkbox v-model="config.system.needInitData">开启</el-checkbox>
      </el-form-item>
      <el-form-item label="多点登录拦截">
        <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
      </el-form-item>
      <!--  System end  -->

      <!--  JWT start  -->
      <h2>jwt签名</h2>
      <el-form-item label="jwt签名">
        <el-input v-model="config.jwt.signingKey" />
      </el-form-item>
      <!--  JWT end  -->

      <!--  Zap start  -->
      <h2>Zap日志配置</h2>
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
      <el-form-item label="软链接名称">
        <el-input v-model="config.zap.linkName" />
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
      <!--  Zap end  -->

      <!--  Redis start  -->
      <h2>Redis admin数据库配置</h2>
      <el-form-item label="db">
        <el-input v-model="config.redis.db" />
      </el-form-item>
      <el-form-item label="addr">
        <el-input v-model="config.redis.addr" />
      </el-form-item>
      <el-form-item label="password">
        <el-input v-model="config.redis.password" />
      </el-form-item>
      <!--  Redis end  -->

      <!--  Email start  -->
      <h2>邮箱配置</h2>
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
      <!--  Email end  -->

      <!--  Casbin start  -->
      <h2>casbin配置</h2>
      <el-form-item label="模型地址">
        <el-input v-model="config.casbin.modelPath" />
      </el-form-item>
      <!--  Casbin end  -->

      <!--  Captcha start  -->
      <h2>验证码配置</h2>
      <el-form-item label="keyLong">
        <el-input v-model.number="config.captcha.keyLong" />
      </el-form-item>
      <el-form-item label="imgWidth">
        <el-input v-model.number="config.captcha.imgWidth" />
      </el-form-item>
      <el-form-item label="imgHeight">
        <el-input v-model.number="config.captcha.imgHeight" />
      </el-form-item>
      <!--  Captcha end  -->

      <!--  dbType start  -->
      <template v-if="config.system.dbType === 'mysql'">
        <h2>mysql admin数据库配置</h2>
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
      <template v-if="config.system.dbType === 'sqlite'">
        <h2>sqlite admin数据库配置</h2>
        <el-form-item label="path">
          <el-input v-model="config.mysql.path" />
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.mysql.maxIdleConns" />
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.mysql.maxOpenConns" />
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.mysql.logger" />
        </el-form-item>
      </template>
      <template v-if="config.system.dbType === 'sqlserver'">
        <h2>sqlserver admin数据库配置</h2>
        <el-form-item label="username">
          <el-input v-model="config.sqlserver.username" />
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.sqlserver.password" />
        </el-form-item>
        <el-form-item label="path">
          <el-input v-model="config.sqlserver.path" />
        </el-form-item>
        <el-form-item label="dbname">
          <el-input v-model="config.sqlserver.dbname" />
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.sqlserver.maxIdleConns" />
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.sqlserver.maxOpenConns" />
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.sqlserver.logger" />
        </el-form-item>
      </template>
      <template v-if="config.system.dbType === 'postgresql'">
        <h2>postgresql admin数据库配置</h2>
        <el-form-item label="username">
          <el-input v-model="config.mysql.username" />
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="config.mysql.password" />
        </el-form-item>
        <el-form-item label="dbName">
          <el-input v-model="config.mysql.dbName" />
        </el-form-item>
        <el-form-item label="port">
          <el-input v-model="config.mysql.port" />
        </el-form-item>
        <el-form-item label="config">
          <el-input v-model="config.mysql.config" />
        </el-form-item>
        <el-form-item label="maxIdleConns">
          <el-input v-model.number="config.mysql.maxIdleConns" />
        </el-form-item>
        <el-form-item label="maxOpenConns">
          <el-input v-model.number="config.mysql.maxOpenConns" />
        </el-form-item>
        <el-form-item label="logger">
          <el-checkbox v-model="config.mysql.logger" />
        </el-form-item>
        <el-form-item label="prefer-simple-protocol">
          <el-checkbox v-model="config.mysql.preferSimpleProtocol" />
        </el-form-item>
      </template>
      <!--  dbType end  -->

      <!--  ossType start  -->
      <template v-if="config.system.ossType === 'local'">
        <h2>本地上传配置</h2>
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
      <!--  ossType end  -->
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" size="mini" @click="update">立即更新</el-button>
      <el-button type="primary" size="mini" @click="reload">重启服务（开发中）</el-button>
    </div>
  </div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from '@/api/system'
import { emailTest } from '@/api/email'
export default {
  name: 'Config',
  data() {
    return {
      config: {
        system: {},
        jwt: {},
        casbin: {},
        mysql: {},
        sqlite: {},
        redis: {},
        qiniu: {},
        tencentCOS: {},
        aliyunOSS: {},
        captcha: {},
        zap: {},
        local: {},
        email: {}
      }
    }
  },
  async created() {
    await this.initForm()
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig()
      if (res.code === 0) {
        this.config = res.data.config
      }
    },
    reload() {},
    async update() {
      const res = await setSystemConfig({ config: this.config })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '配置文件设置成功'
        })
        await this.initForm()
      }
    },
    async email() {
      const res = await emailTest()
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '邮件发送成功'
        })
        await this.initForm()
      } else {
        this.$message({
          type: 'error',
          message: '邮件发送失败'
        })
      }
    }
  }
}
</script>

<style lang="scss">
.system {
  background: #fff;
  padding:12px;
  border-radius: 2px;
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 0px 0px 0px #e7e8e8;
  }
}
</style>
