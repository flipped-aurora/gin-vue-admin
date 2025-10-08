<template>
  <el-button type="primary" icon="download" @click="exportExcelFunc"
    >导出</el-button
  >
</template>

<script setup>

import { exportExcel } from '@/api/exportTemplate'

  const props = defineProps({
    filterDeleted: {
      type: Boolean,
      default: true
    },
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

  const exportExcelFunc = async () => {
    if (props.templateId === '') {
      ElMessage.error('组件未设置模板ID')
      return
    }
    let baseUrl = import.meta.env.VITE_BASE_API
    if (baseUrl === "/"){
      baseUrl = ""
    }
    const paramsCopy = JSON.parse(JSON.stringify(props.condition))

    if (props.filterDeleted) {
      paramsCopy.filterDeleted = 'true'
    }

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
      .map(
        ([key, value]) =>
          `${encodeURIComponent(key)}=${encodeURIComponent(value)}`
      )
      .join('&')

    const res = await exportExcel({
      templateID: props.templateId,
      params
    })

    if(res.code === 0){
      ElMessage.success('创建导出任务成功，开始下载')
      const url = `${baseUrl}${res.data}`
      window.open(url, '_blank')
    }


  }
</script>
