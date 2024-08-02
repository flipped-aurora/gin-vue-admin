<template>
  <el-button
    type="primary"
    icon="download"
    @click="exportTemplate"
  >{{ t('components.exportExcel.exportTemplate.exportTemplate') }}</el-button>
</template>

<script setup>
const props = defineProps({
  templateId: {
    type: String,
    required: true
  }
})

import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const exportTemplate = async() => {
  if (props.templateId === '') {
    ElMessage.error(t('components.exportExcel.exportTemplate.templateIdErr'))
    return
  }
  const baseUrl = import.meta.env.VITE_BASE_API
  const url = `${baseUrl}/sysExportTemplate/exportTemplate?templateID=${props.templateId}`
  window.open(url, '_blank')
}
</script>
