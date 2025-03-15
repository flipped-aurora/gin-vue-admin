<template>
  <el-upload
    :action="url"
    :show-file-list="false"
    :on-success="handleSuccess"
    :multiple="false"
    :headers="{'x-token': token}"
  >
    <el-button type="primary" icon="upload" class="ml-3"> 导入 </el-button>
  </el-upload>
</template>

<script setup>
  import { ElMessage } from 'element-plus'
  import { useUserStore } from "@/pinia";

  let baseUrl = import.meta.env.VITE_BASE_API
  if (baseUrl === "/"){
    baseUrl = ""
  }

  const props = defineProps({
    templateId: {
      type: String,
      required: true
    }
  })

  const userStore = useUserStore()

  const token = userStore.token

  const emit = defineEmits(['on-success'])

  const url = `${baseUrl}/sysExportTemplate/importExcel?templateID=${props.templateId}`

  const handleSuccess = (res) => {
    if (res.code === 0) {
      ElMessage.success('导入成功')
      emit('on-success')
    } else {
      ElMessage.error(res.msg)
    }
  }
</script>
