<template>
  <div>
    <el-upload
      :action="`${path}/fileUploadAndDownload/upload`"
      :before-remove="beforeRemove"
      :file-list="fileList"
      :headers="{'x-token':token}"
      :limit="10"
      :on-exceed="handleExceed"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      class="upload-demo"
      multiple
    >
      <el-button size="small" type="primary">点击上传</el-button>
      <div class="el-upload__tip" slot="tip">未对文件格式及大小做校验</div>
    </el-upload>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
const path = process.env.VUE_APP_BASE_API
export default {
  name: 'Excel',
  data() {
    return {
      path: path,

      fileList: [
        {
          name: 'food.jpeg',
          url:
            'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        },
        {
          name: 'food2.jpeg',
          url:
            'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        }
      ]
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  methods: {
    handleRemove(file, fileList) {
      this.$message.warning(
        `共有 ${fileList.length} 个文件，移除了${file.name}`
      )
    },
    handlePreview(file) {
      this.$message.warning(`${file.name}选择完成`)
    },
    handleExceed(files, fileList) {
      this.$message.warning(
        `当前限制选择 3 个文件，本次选择了 ${
          files.length
        } 个文件，共选择了 ${files.length + fileList.length} 个文件`
      )
    },
    beforeRemove(file, fileList) {
      return this.$confirm(
        `共有 ${fileList.length} 个文件，确定移除 ${file.name}？`
      )
    }
  }
}
</script>