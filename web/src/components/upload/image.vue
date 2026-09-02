<template>
  <div>
    <el-upload
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
      :show-file-list="false"
      :on-success="handleImageSuccess"
      :before-upload="beforeImageUpload"
      :multiple="false"
      :data="{'classId': props.classId}"
      :headers="{'x-token': token}"
    >
      <el-button type="primary" :icon="Upload">压缩上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
  import ImageCompress from '@/utils/image'
  import { ElMessage } from 'element-plus'
  import { getBaseUrl } from '@/utils/format'
  import { Upload } from "@element-plus/icons-vue";
  import { useUserStore } from "@/pinia";

  defineOptions({
    name: 'UploadImage'
  })

  const emit = defineEmits(['on-success'])
  const props = defineProps({
    imageUrl: {
      type: String,
      default: ''
    },
    fileSize: {
      type: Number,
      default: 2048 // 2M 超出后执行压缩
    },
    maxWH: {
      type: Number,
      default: 1920 // 图片长宽上限
    },
    classId: {
      type: Number,
      default: 0
    }
  })

  const userStore = useUserStore()

  const token = userStore.token

  const beforeImageUpload = (file) => {
    const isJPG = file.type?.toLowerCase() === 'image/jpeg'
    const isPng = file.type?.toLowerCase() === 'image/png'
    if (!isJPG && !isPng) {
      ElMessage.error('上传头像图片只能是 jpg或png 格式!')
      return false
    }

    const isRightSize = file.size / 1024 < props.fileSize
    if (!isRightSize) {
      // 压缩
      const compress = new ImageCompress(file, props.fileSize, props.maxWH)
      return compress.compress()
    }
    return isRightSize
  }

  const handleImageSuccess = (res) => {
    const { data } = res
    if (data.file) {
      emit('on-success', data.file.url)
    }
  }
</script>

<style lang="scss" scoped>
  .image-uploader {
    border: 1px dashed #d9d9d9;
    width: 180px;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .image-uploader {
    border-color: #409eff;
  }
  .image-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .image {
    width: 178px;
    height: 178px;
    display: block;
  }
</style>
