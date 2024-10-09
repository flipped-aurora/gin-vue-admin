<template>
  <div>
    <el-upload
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
      :before-upload="checkFile"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      multiple
      class="upload-btn"
    >
      <el-button type="primary">{{ t('components.upload.common.normalUpload') }}</el-button>
    </el-upload>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { isVideoMime, isImageMime } from '@/utils/image'
import { getBaseUrl } from '@/utils/format'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: 'UploadCommon',
})

const emit = defineEmits(['on-success'])

const fullscreenLoading = ref(false)

const checkFile = (file) => {
  fullscreenLoading.value = true
  const isLt500K = file.size / 1024 / 1024 < 0.5 // 500K, @todo 应支持在项目中设置
  const isLt5M = file.size / 1024 / 1024 < 5 // 5MB, @todo 应支持项目中设置
  const isVideo = isVideoMime(file.type)
  const isImage = isImageMime(file.type)
  let pass = true
  if (!isVideo && !isImage) {
    ElMessage.error(t('components.upload.common.fileFormatNote'))
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt5M && isVideo) {
    ElMessage.error(t('components.upload.common.videoSizeError'))
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt500K && isImage) {
    ElMessage.error(t('components.upload.common.imageSizeError'))
    fullscreenLoading.value = false
    pass = false
  }

  console.log('upload file check result: ', pass)

  return pass
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: t('selectFile.uploadFailed')
  })
  fullscreenLoading.value = false
}

</script>

