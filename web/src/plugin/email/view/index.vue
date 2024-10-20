<template>
  <div>
    <warning-bar :title="t('plugins.email.emailConfigNote')" />
    <div class="gva-form-box">
      <el-form
        ref="emailForm"
        label-position="right"
        label-width="140px"
        :model="form"
      >
        <el-form-item :label="t('plugins.email.targetEmail')">
          <el-input v-model="form.to" />
        </el-form-item>
        <el-form-item :label="t('plugins.email.email')">
          <el-input v-model="form.subject" />
        </el-form-item>
        <el-form-item :label="t('plugins.email.emailContent')">
          <el-input
            v-model="form.body"
            type="textarea"
          />
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTestEmail">{{ t('plugins.email.sendTestEmail') }}</el-button>
          <el-button @click="sendEmail">{{ t('plugins.email.sendEmail') }}</el-button>
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
    ElMessage.success(t('plugins.email.sendSuccess'))
  }
}

const sendEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success(t('plugins.email.sentSuccessPleaseCheck'))
  }
}
</script>

