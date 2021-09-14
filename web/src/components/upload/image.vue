
<template>
  <div>
    <el-upload
      :action="`${path}/fileUploadAndDownload/upload`"
      :headers="{ 'x-token': token }"
      :show-file-list="false"
      :on-success="handleImageSuccess"
      :before-upload="beforeImageUpload"
      :multiple="false"
    >
      <el-button size="mini" type="primary">压缩上传</el-button>
    </el-upload>
  </div>
</template>

<script>
const path = import.meta.env.VITE_BASE_API
import { mapGetters } from 'vuex'
import ImageCompress from '@/utils/image'
export default {
  name: 'UploadImage',
  model: {
    prop: 'imageUrl',
    event: 'change'
  },
  props: {
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
    }
  },
  data() {
    return {
      path: path
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token']),
    showImageUrl() {
      return (this.imageUrl && this.imageUrl.slice(0, 4) !== 'http') ? path + this.imageUrl : this.imageUrl
    }
  },
  methods: {
    beforeImageUpload(file) {
      const isJPG = file.type === 'image/jpeg'
      const isPng = file.type === 'image/png'
      if (!isJPG && !isPng) {
        this.$message.error('上传头像图片只能是 jpg或png 格式!')
        return false
      }

      const isRightSize = file.size / 1024 < this.fileSize
      if (!isRightSize) {
        // 压缩
        const compress = new ImageCompress(file, this.fileSize, this.maxWH)
        return compress.compress()
      }
      return isRightSize
    },
    handleImageSuccess(res) {
      // this.imageUrl = URL.createObjectURL(file.raw);
      const { data } = res
      if (data.file) {
        this.$emit('change', data.file.url)
        this.$emit('on-success')
      }
    }
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
