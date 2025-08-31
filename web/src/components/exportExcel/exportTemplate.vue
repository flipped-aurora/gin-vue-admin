<template>
  <el-button type="primary" icon="download" @click="exportTemplateFunc"
    >{{ t('components.exportExcel.exportTemplate.exportTemplate') }}</el-button
  >
</template>

<script setup>
  import { ElMessage } from 'element-plus'
  import {exportTemplate} from "@/api/exportTemplate";
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

  const { t } = useI18n() // added by mohamed hassan to support multilingual
  
  const props = defineProps({
    templateId: {
      type: String,
      required: true
    }
  })

  const exportTemplateFunc = async () => {
    if (props.templateId === '') {
      ElMessage.error(t('components.exportExcel.exportTemplate.templateIdErr'))
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
      ElMessage.success(t('components.exportExcel.createTaskSuccess'))
      const url = `${baseUrl}${res.data}`
      window.open(url, '_blank')
    }

  }
</script>
