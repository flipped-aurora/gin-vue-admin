<template>
  <div>
    <el-upload
        multiple
        :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
        :headers="{ 'x-token': userStore.token }"
        :on-error="uploadError"
        :on-success="uploadSuccess"
        :show-file-list="true"
        :file-list="fileList"
        class="upload-btn"
    >
      <el-button type="primary">上传文件</el-button>
    </el-upload>
  </div>
</template>

<script setup>

import {ref, watch} from 'vue'
import {ElMessage} from 'element-plus'
import {useUserStore} from '@/pinia/modules/user'

defineOptions({
  name: 'UploadCommon',
})

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  }
})

const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

const fileList = ref(props.modelValue)

const emits = defineEmits(['update:modelValue'])

watch(fileList.value, (val) => {
  console.log(val)
  emits('update:modelValue', val)
})

watch(
    () => props.modelValue,
    value => {
      fileList.value = value
    },
    {immediate: true}
)
const uploadSuccess = (res) => {
  const {data} = res
  if (data.file) {
    fileList.value.push({
      name: data.file.name,
      url: data.file.url
    })
    fullscreenLoading.value = false
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}

</script>

