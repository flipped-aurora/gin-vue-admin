<template>
  <el-button
    type="primary"
    icon="download"
    @click="exportExcelFunc"
  >{{ t('components.exportExcel.exportExcel.export') }}</el-button>
</template>

<script setup>
const props = defineProps({
  templateId: {
    type: String,
    required: true
  },
  condition: {
    type: Object,
    default: () => ({})
  },
  limit: {
    type: Number,
    default: 0
  },
  offset: {
    type: Number,
    default: 0
  },
  order: {
    type: String,
    default: ''
  }
})

import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const exportExcelFunc = async() => {
  if (props.templateId === '') {
    ElMessage.error(t('components.exportExcel.exportExcel.templateIdErr'))
    return
  }
  const baseUrl = import.meta.env.VITE_BASE_API
  const paramsCopy = JSON.parse(JSON.stringify(props.condition))
  if (props.limit) {
    paramsCopy.limit = props.limit
  }
  if (props.offset) {
    paramsCopy.offset = props.offset
  }
  if (props.order) {
    paramsCopy.order = props.order
  }
  const params = Object.entries(paramsCopy)
    .map(([key, value]) => `${encodeURIComponent(key)}=${encodeURIComponent(value)}`)
    .join('&')
  const url = `${baseUrl}/sysExportTemplate/exportExcel?templateID=${props.templateId}${params ? '&' + params : ''}`

  window.open(url, '_blank')
}
</script>
