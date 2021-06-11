<template>
  <div class="init">
    <p class="in-one a-fadeinT">欢迎使用GIN-VUE-ADMIN</p>
    <p class="in-two a-fadeinT">您需要初始化您的数据库并且填充初始数据</p>
    <div class="form-card in-three a-fadeinB">
      <el-form ref="form" :model="form" label-width="100px">
        <el-form-item label="数据库类型">
          <el-select v-model="form.sqlType" disabled placeholder="请选择">
            <el-option key="mysql" label="mysql(目前只支持mysql)" value="mysql" />
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
          <el-input
            v-model="form.password"
            placeholder="请输入数据库密码（没有则为空）"
          />
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
</template>

<script>
import { initDB } from '@/api/initdb'
export default {
  name: 'Init',
  data() {
    return {
      form: {
        sqlType: 'mysql',
        host: '127.0.0.1',
        port: '3306',
        userName: 'root',
        password: '',
        dbName: 'gva'
      }
    }
  },
  methods: {
    async onSubmit() {
      const loading = this.$loading({
        lock: true,
        text: '正在初始化数据库，请稍候',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      try {
        const res = await initDB(this.form)
        if (res.code === 0) {
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

<style lang="scss">
.init {
  height: 100vh;
  flex-direction: column;
  display: flex;
  align-items: center;
  background: #fff;
}
.in-one {
  font-size: 26px;
  line-height: 98px;
}
.in-two {
  font-size: 22px;
}
.form-card {
  margin-top: 60px;
  box-shadow: 0px 0px 5px 0px rgba(5, 12, 66, 0.15);
  width: 60vw;
  height: auto;
  background: #fff;
  padding: 20px;
  border-radius: 6px;
}
</style>
