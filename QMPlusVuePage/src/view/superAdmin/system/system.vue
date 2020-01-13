<template>
  <div>
    <el-form :model="config" label-width="100px" ref="form">
      <h2>系统配置</h2>
      <el-form-item label="多点登录拦截">
        <el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
      </el-form-item>
      <el-form-item label="环境值">
        <el-input v-model="config.system.env"></el-input>
      </el-form-item>
      <h2>jwt签名</h2>
      <el-form-item label="jwt签名">
        <el-input v-model="config.jwt.signingKey"></el-input>
      </el-form-item>
      <h2>casbin配置</h2>
      <el-form-item label="模型地址">
        <el-input v-model="config.casbinConfig.modelPath"></el-input>
      </el-form-item>
      <h2>mysql admin数据库配置</h2>
      <el-form-item label="username">
        <el-input v-model="config.mysqlAdmin.username"></el-input>
      </el-form-item>
      <el-form-item label="password">
        <el-input v-model="config.mysqlAdmin.password"></el-input>
      </el-form-item>
      <el-form-item label="path">
        <el-input v-model="config.mysqlAdmin.path"></el-input>
      </el-form-item>
      <el-form-item label="dbname">
        <el-input v-model="config.mysqlAdmin.dbname"></el-input>
      </el-form-item>
      <el-form-item label="config">
        <el-input v-model="config.mysqlAdmin.config"></el-input>
      </el-form-item>
      <h2>Redis admin数据库配置</h2>
      <el-form-item label="addr">
        <el-input v-model="config.redisAdmin.addr"></el-input>
      </el-form-item>
      <el-form-item label="password">
        <el-input v-model="config.redisAdmin.password"></el-input>
      </el-form-item>
      <el-form-item label="db">
        <el-input v-model="config.redisAdmin.db"></el-input>
      </el-form-item>
      <h2>七牛密钥配置</h2>
      <el-form-item label="accessKey">
        <el-input v-model="config.qiniu.accessKey"></el-input>
      </el-form-item>
      <el-form-item label="secretKey">
        <el-input v-model="config.qiniu.secretKey"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="update" type="primary">立即更新</el-button>
        <el-button @click="reload" type="primary">重启服务（开发中）</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from '@/api/system'
export default {
  name: 'Config',
  data() {
    return {
      config: {
        system: {},
        jwt: {},
        casbinConfig: {},
        mysqlAdmin: {},
        redisAdmin: {},
        qiniu: {}
      }
    }
  },
  async created() {
      await this.initForm()
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig()
      if (res.success) {
        this.config = res.data.config
      }
    },
    reload() {},
    async update() {
      const res = await setSystemConfig({ config: this.config })
      if (res.success) {
        this.$message({
          type:"success",
          message:"配置文件设置成功"
        })
          await this.initForm()
      }
    }
  }
}
</script>
<style lang="scss">
h2 {
  padding: 10px;
  border-bottom: 1px dashed #ccc;
  margin: 10px 0;
  font-size: 16px;
}
</style>