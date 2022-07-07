<template>
  <div>
    <warning-bar title="需要提前配置email配置文件，为防止不必要的垃圾邮件，在线体验功能不开放此功能体验。" />
    <div class="gva-form-box">
      <el-form ref="emailForm" label-position="right" label-width="80px" :model="form">
        <el-form-item label="目标邮箱">
          <el-input v-model="form.to" />
        </el-form-item>
        <el-form-item label="邮件">
          <el-input v-model="form.subject" />
        </el-form-item>
        <el-form-item label="邮件内容">
          <el-input v-model="form.body" type="textarea" />
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTestEmail">发送测试邮件</el-button>
          <el-button @click="sendEmail">发送邮件</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

</template>

<script>
export default {
  name: 'Email',
}
</script>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { emailTest } from '@/plugin/email/api/email.js'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
const emailForm = ref(null)
const form = reactive({
  to: '',
  subject: '',
  body: '',
})
const sendTestEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('发送成功')
  }
}

const sendEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('发送成功,请查收')
  }
}
</script>

