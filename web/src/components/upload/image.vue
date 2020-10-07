<template>
  <div>
    <el-upload class="image-uploader"
      :action="`${path}/fileUploadAndDownload/upload`"
      :show-file-list="false"
      :on-success="handleImageSuccess"
      :before-upload="beforeImageUpload"
      :multiple="false"
    >
      <img v-if="imageUrl" :src="imageUrl" class="image" />
      <i v-else class="el-icon-plus image-uploader-icon"></i>
    </el-upload>
  </div>
</template>
<script>
const path = process.env.VUE_APP_BASE_API;
import ImageCompress from '@/utils/image.js'
export default {
  name: "upload-image",
  model: {
    prop: "imageUrl",
    event: "change",
  },
  props: {
    imageUrl: {
      type: String,
      default: "",
    },
    fileSize: {
      type: Number,
      default: 2048 // 2M
    },
  },
  data() {
    return {
      path: path,
    };
  },
  methods: {
    beforeImageUpload(file) {
      let isRightSize = file.size / 1024 < this.fileSize;
      if (!isRightSize) {
        // 压缩
        let compress = new ImageCompress(file, this.fileSize)
        return compress.compress()
        // compress.then(res => {
        // }).catch((err) => {
        //   console.error(err)
        //   this.$message.error("文件大小超过 2MB");
        // })
      }
      return isRightSize;
    },
    handleImageSuccess(res, file) {
      // this.imageUrl = URL.createObjectURL(file.raw);
      const { code, msg, data } = res;
      if (data.file) {
        this.$emit("change", data.file.url);
      }
    },
  },
};
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