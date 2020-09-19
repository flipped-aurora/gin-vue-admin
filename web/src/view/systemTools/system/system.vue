<template>
	<div class="system">
		<el-form :model="config" label-width="100px" ref="form" class="system">
			<h2>系统配置</h2>
			<el-form-item label="多点登录拦截">
				<el-checkbox v-model="config.system.useMultipoint">开启</el-checkbox>
			</el-form-item>
			<el-form-item label="环境值">
				<el-input v-model="config.system.env"></el-input>
			</el-form-item>
			<el-form-item label="端口值">
				<el-input v-model.number="config.system.addr"></el-input>
			</el-form-item>
			<el-form-item label="数据库类型">
				<el-select v-model="config.system.dbType">
					<el-option value="sqlite"></el-option>
					<el-option value="mysql"></el-option>
				</el-select>
			</el-form-item>
			<h2>jwt签名</h2>
			<el-form-item label="jwt签名">
				<el-input v-model="config.jwt.signingKey"></el-input>
			</el-form-item>
			<h2>casbin配置</h2>
			<el-form-item label="模型地址">
				<el-input v-model="config.casbin.modelPath"></el-input>
			</el-form-item>
			<template v-show="config.system.dbType == 'mysql'">
				<h2>mysql admin数据库配置</h2>
				<el-form-item label="username">
					<el-input v-model="config.mysql.username"></el-input>
				</el-form-item>
				<el-form-item label="password">
					<el-input v-model="config.mysql.password"></el-input>
				</el-form-item>
				<el-form-item label="path">
					<el-input v-model="config.mysql.path"></el-input>
				</el-form-item>
				<el-form-item label="dbname">
					<el-input v-model="config.mysql.dbname"></el-input>
				</el-form-item>
				<el-form-item label="maxIdleConns">
					<el-input v-model.number="config.mysql.maxIdleConns"></el-input>
				</el-form-item>
				<el-form-item label="maxOpenConns">
					<el-input v-model.number="config.mysql.maxOpenConns"></el-input>
				</el-form-item>
				<el-form-item label="logMode">
					<el-checkbox v-model="config.mysql.logMode"></el-checkbox>
				</el-form-item>
			</template>
			<template v-show="config.system.dbType == 'sqlite'">
				<h2>sqlite admin数据库配置</h2>
				<el-form-item label="path">
					<el-input v-model="config.sqlite.path"></el-input>
				</el-form-item>
				<el-form-item label="config">
					<el-input v-model="config.sqlite.config"></el-input>
				</el-form-item>
				<el-form-item label="logMode">
					<el-checkbox v-model="config.sqlite.logMode"></el-checkbox>
				</el-form-item>
			</template>
			<h2>Redis admin数据库配置</h2>
			<el-form-item label="addr">
				<el-input v-model="config.redis.addr"></el-input>
			</el-form-item>
			<el-form-item label="password">
				<el-input v-model="config.redis.password"></el-input>
			</el-form-item>
			<el-form-item label="db">
				<el-input v-model="config.redis.db"></el-input>
			</el-form-item>
			<h2>上传配置</h2>
				<el-form-item label="本地或七牛云">
					<el-checkbox v-model="config.localUpload.local">本地</el-checkbox>
				</el-form-item>
				<el-form-item label="本地文件路径">
					<el-input v-model="config.localUpload.filePath"></el-input>
				</el-form-item>
			<h2>七牛密钥配置</h2>
			<el-form-item label="accessKey">
				<el-input v-model="config.qiniu.accessKey"></el-input>
			</el-form-item>
			<el-form-item label="secretKey">
				<el-input v-model="config.qiniu.secretKey"></el-input>
			</el-form-item>
			<h2>验证码配置</h2>
			<el-form-item label="keyLong">
				<el-input v-model.number="config.captcha.keyLong"></el-input>
			</el-form-item>
			<el-form-item label="imgWidth">
				<el-input v-model.number="config.captcha.imgWidth"></el-input>
			</el-form-item>
			<el-form-item label="imgHeight">
				<el-input v-model.number="config.captcha.imgHeight"></el-input>
			</el-form-item>
			<h2>日志配置</h2>
			<el-form-item label="prefix">
				<el-input v-model.number="config.log.prefix"></el-input>
			</el-form-item>
			<el-form-item label="logFile">
				<el-checkbox v-model="config.log.logFile"></el-checkbox>
			</el-form-item>
			<h2>邮箱配置</h2>
			<el-form-item label="emailFrom">
				<el-input v-model="config.email.emailFrom"></el-input>
			</el-form-item>
			<el-form-item label="emailNickName">
				<el-input v-model="config.email.emailNickName"></el-input>
			</el-form-item>
			<el-form-item label="emailSecret">
				<el-input v-model="config.email.emailSecret"></el-input>
			</el-form-item>
			<el-form-item label="emailTo">
				<el-input v-model="config.email.emailTo" placeholder="可多个，以逗号分隔"></el-input>
			</el-form-item>
			<el-form-item label="emailHost">
				<el-input v-model="config.email.emailHost"></el-input>
			</el-form-item>
			<el-form-item label="emailPort">
				<el-input v-model.number="config.email.emailPort"></el-input>
			</el-form-item>
			<el-form-item label="emailIsSSL">
				<el-checkbox v-model="config.email.emailIsSSL"></el-checkbox>
			</el-form-item>
			<el-form-item label="测试邮件">
				<el-button @click="email">测试邮件</el-button>
			</el-form-item>
			<el-form-item>
				<el-button @click="update" type="primary">立即更新</el-button>
				<el-button @click="reload" type="primary">重启服务（开发中）</el-button>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
import { getSystemConfig, setSystemConfig } from "@/api/system";
import { emailTest } from "@/api/email";
export default {
  name: "Config",
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
        captcha: {},
		log: {},
		localUpload: {},
		email: {},
      }
    };
  },
  async created() {
    await this.initForm();
  },
  methods: {
    async initForm() {
      const res = await getSystemConfig();
      if (res.code == 0) {
        this.config = res.data.config;
      }
    },
    reload() {},
    async update() {
      const res = await setSystemConfig({ config: this.config });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "配置文件设置成功"
        });
        await this.initForm();
      }
    },
    async email() {
      const res = await emailTest();
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "邮件发送成功"
        });
        await this.initForm();
      } else {
		this.$message({
          type: "error",
          message: "邮件发送失败"
        });
	  }
    }
  }
};
</script>
<style lang="scss">
.system {
  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 1px 3px 0px #e7e8e8;
  }
}
</style>
