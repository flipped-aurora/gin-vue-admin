<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-upload
          :action="`${path}/fileUploadAndDownload/upload`"
          :before-upload="checkFile"
          :headers="{ 'x-token': token }"
          :on-error="uploadError"
          :on-success="uploadSuccess"
          :show-file-list="false"
          class="upload-btn"
        >
          <el-button size="mini" type="primary">普通上传</el-button>
        </el-upload>
        <upload-image
          v-model="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          class="upload-btn"
          @on-success="getTableData"
        />
      </div>

      <el-table :data="tableData">
        <el-table-column align="left" label="预览" width="100">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.url" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" prop="UpdatedAt" width="180">
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="文件名" prop="name" width="180" />
        <el-table-column align="left" label="链接" prop="url" min-width="300" />
        <el-table-column align="left" label="标签" prop="tag" width="100">
          <template #default="scope">
            <el-tag
              :type="scope.row.tag === 'jpg' ? 'primary' : 'success'"
              disable-transitions
            >{{ scope.row.tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button size="small" icon="el-icon-download" type="text" @click="downloadFile(scope.row)">下载</el-button>
            <el-button size="small" icon="el-icon-delete" type="text" @click="deleteFile(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :style="{ float: 'right', padding: '20px' }"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script>
const path = import.meta.env.VITE_BASE_API
import { mapGetters } from 'vuex'
import infoList from '@/mixins/infoList'
import { getFileList, deleteFile } from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
export default {
  name: 'Upload',
  components: {
    CustomPic,
    UploadImage
  },
  mixins: [infoList],
  data() {
    return {
      fullscreenLoading: false,
      listApi: getFileList,
      path: path,
      tableData: [],
      imageUrl: ''
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  created() {
    this.getTableData()
  },
  methods: {
    async deleteFile(row) {
      this.$confirm('此操作将永久文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          const res = await deleteFile(row)
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            if (this.tableData.length === 1 && this.page > 1) {
              this.page--
            }
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    checkFile(file) {
      this.fullscreenLoading = true
      const isJPG = file.type === 'image/jpeg'
      const isPng = file.type === 'image/png'
      const isLt2M = file.size / 1024 / 1024 < 0.5
      if (!isJPG && !isPng) {
        this.$message.error('上传图片只能是 jpg或png 格式!')
        this.fullscreenLoading = false
      }
      if (!isLt2M) {
        this.$message.error('未压缩未见上传图片大小不能超过 500KB，请使用压缩上传')
        this.fullscreenLoading = false
      }
      return (isPng || isJPG) && isLt2M
    },
    uploadSuccess(res) {
      this.fullscreenLoading = false
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '上传成功'
        })
        if (res.code === 0) {
          this.getTableData()
        }
      } else {
        this.$message({
          type: 'warning',
          message: res.msg
        })
      }
    },
    uploadError() {
      this.$message({
        type: 'error',
        message: '上传失败'
      })
      this.fullscreenLoading = false
    },
    downloadFile(row) {
      if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
        downloadImage(row.url, row.name)
      } else {
        downloadImage(this.path + row.url, row.name)
      }
    }
  }
}
</script>
<style scoped>
.upload-btn+.upload-btn {
            margin-left: 12px;
        }
</style>
