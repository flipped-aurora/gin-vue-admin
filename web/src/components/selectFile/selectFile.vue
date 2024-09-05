<template>
  <div>
    <el-upload
      v-model:file-list="fileList"
      multiple
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload?noSave=1`"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="true"
      :limit="limit"
      :accept="accept"
      class="upload-btn"
    >
      <el-button type="primary">
        上传文件
      </el-button>
    </el-upload>
  </div>
</template>

<script setup>

import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { getBaseUrl } from '@/utils/format'

defineOptions({
  name: 'UploadCommon',
})

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  },
  limit: {
    type: Number,
    default: 3
  },
  accept: {
    type: String,
    default: ''
  },
})

const fullscreenLoading = ref(false)

const fileList = ref(props.modelValue)

const emits = defineEmits(['update:modelValue', 'on-success', 'on-error'])

watch(fileList.value, (val) => {
  emits('update:modelValue', val)
})

watch(
  () => props.modelValue,
  value => {
    fileList.value = value
  },
  { immediate: true }
)
const uploadSuccess = (res) => {
  const { data,code } = res
  if(code !== 0){
    ElMessage({
      type: 'error',
      message: '上传失败'+res.msg
    })
    fileList.value.pop()
    return
  }
  if (data.file) {
    fileList.value.push({
      name: data.file.name,
      url: data.file.url
    })
    fullscreenLoading.value = false
  }
  emits('on-success', res)
}

const uploadError = (err) => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
  emits('on-error',err)
}

</script>

