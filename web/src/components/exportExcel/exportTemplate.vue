<template>
  <el-button type="primary" icon="download" @click="exportTemplateFunc"
    >下载模板</el-button
  >
</template>

<script setup>
  import { ElMessage } from 'element-plus'
  import {exportTemplate} from "@/api/exportTemplate";

  const props = defineProps({
    templateId: {
      type: String,
      required: true
    }
  })


  const exportTemplateFunc = async () => {
    if (props.templateId === '') {
      ElMessage.error('组件未设置模板ID')
      return
    }
    let baseUrl = import.meta.env.VITE_BASE_API
    if (baseUrl === "/"){
      baseUrl = ""
    }

    const res = await exportTemplate({
      templateID: props.templateId
    })

    if(res.code === 0){
      ElMessage.success('创建导出任务成功，开始下载')
      const url = `${baseUrl}${res.data}`
      window.open(url, '_blank')
    }

  }
</script>
