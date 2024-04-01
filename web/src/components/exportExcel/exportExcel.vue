<template>
  <el-button
    type="primary"
    icon="download"
    @click="exportExcelFunc"
  >导出</el-button>
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

const exportExcelFunc = async() => {
  if (props.templateId === '') {
    ElMessage.error('组件未设置模板ID')
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
