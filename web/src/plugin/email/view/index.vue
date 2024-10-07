<template>
  <div>
    <warning-bar :title="t('view.plugins.emailConfigNote')" />
    <div class="gva-form-box">
      <el-form
        ref="emailForm"
        label-position="right"
        label-width="140px"
        :model="form"
      >
        <el-form-item :label="t('view.plugins.targetEmail')">
          <el-input v-model="form.to" />
        </el-form-item>
        <el-form-item :label="t('view.plugins.email')">
          <el-input v-model="form.subject" />
        </el-form-item>
        <el-form-item :label="t('view.plugins.emailContent')">
          <el-input
            v-model="form.body"
            type="textarea"
          />
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTestEmail">{{ t('view.plugins.sendTestEmail') }}</el-button>
          <el-button @click="sendEmail">{{ t('view.plugins.sendEmail') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

</template>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { emailTest } from '@/plugin/email/api/email.js'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
defineOptions({
  name: 'Email',
})

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

