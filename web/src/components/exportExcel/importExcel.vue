<template>
  <el-upload
    :action="url"
    :show-file-list="false"
    :on-success="handleSuccess"
    :multiple="false"
  >
    <el-button
      type="primary"
      icon="upload"
      class="ml-3"
    >
      {{ t('components.exportExcel.importExcel.import') }}
    </el-button>
  </el-upload>
</template>

<script setup>
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const baseUrl = import.meta.env.VITE_BASE_API

const props = defineProps({
  templateId: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['on-success'])

const url = `${baseUrl}/sysExportTemplate/importExcel?templateID=${props.templateId}`

const handleSuccess = (res) => {
  if (res.code === 0) {
    ElMessage.success(t('components.exportExcel.importExcel.importSuccess'))
    emit('on-success')
  } else {
    ElMessage.error(res.msg)
  }
}
</script>
