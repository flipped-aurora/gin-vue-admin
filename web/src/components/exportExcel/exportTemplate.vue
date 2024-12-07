<template>
  <el-button type="primary" icon="download" @click="exportTemplate"
    >下载模板</el-button
  >
</template>

<script setup>
  const props = defineProps({
    templateId: {
      type: String,
      required: true
    }
  })

  import { ElMessage } from 'element-plus'

  const exportTemplate = async () => {
    if (props.templateId === '') {
      ElMessage.error('组件未设置模板ID')
      return
    }
    let baseUrl = import.meta.env.VITE_BASE_API
    if (baseUrl === "/"){
      baseUrl = ""
    }
    const url = `${baseUrl}/sysExportTemplate/exportTemplate?templateID=${props.templateId}`
    window.open(url, '_blank')
  }
</script>
