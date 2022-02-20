<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-upload
          :action="`${path}/fileUploadAndDownload/upload`"
          :before-upload="checkFile"
          :headers="{ 'x-token': userStore.token }"
          :on-error="uploadError"
          :on-success="uploadSuccess"
          :show-file-list="false"
          class="upload-btn"
        >
          <el-button size="small" type="primary">Ordinary upload</el-button>
        </el-upload>
        <upload-image
          v-model:imageUrl="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          class="upload-btn"
          @on-success="getTableData"
        />
      </div>

      <el-table :data="tableData">
        <el-table-column align="left" label="Preview" width="100">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.url" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="Date" prop="UpdatedAt" width="180">
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="File Name" prop="name" width="180" />
        <el-table-column align="left" label="Link" prop="url" min-width="300" />
        <el-table-column align="left" label="Tag" prop="tag" width="100">
          <template #default="scope">
            <el-tag
              :type="scope.row.tag === 'jpg' ? 'primary' : 'success'"
              disable-transitions
            >{{ scope.row.tag }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Action" width="160">
          <template #default="scope">
            <el-button size="small" icon="download" type="text" @click="downloadFile(scope.row)">Download</el-button>
            <el-button size="small" icon="delete" type="text" @click="deleteFileFunc(scope.row)">Delete</el-button>
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

<script setup>
import { getFileList, deleteFile } from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import { useUserStore } from '@/pinia/modules/user'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import { formatDate } from '@/utils/format'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const path = ref(import.meta.env.VITE_BASE_API)
const userStore = useUserStore()

const imageUrl = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getFileList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const deleteFileFunc = async(row) => {
  ElMessageBox.confirm('Does this action will permanently file, do you continue?', 'hint', {
    confirmButtonText: 'Sure',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteFile(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Successfull deleted!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Delete'
      })
    })
}

const fullscreenLoading = ref(false)
const checkFile = (file) => {
  fullscreenLoading.value = true
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isLt2M = file.size / 1024 / 1024 < 0.5
  if (!isJPG && !isPng) {
    ElMessage.error('Upload images can only be JPG or PNG format!')
    fullscreenLoading.value = false
  }
  if (!isLt2M) {
    ElMessage.error('Uncompressed, no upper image size cannot exceed 500KB, please use compression upload')
    fullscreenLoading.value = false
  }
  return (isPng || isJPG) && isLt2M
}
const uploadSuccess = (res) => {
  fullscreenLoading.value = false
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '上传成功'
    })
    if (res.code === 0) {
      getTableData()
    }
  } else {
    ElMessage({
      type: 'warning',
      message: res.msg
    })
  }
}
const uploadError = () => {
  ElMessage({
    type: 'error',
    message: 'upload failed'
  })
  fullscreenLoading.value = false
}
const downloadFile = (row) => {
  if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    downloadImage(row.url, row.name)
  } else {
    downloadImage(path.value + row.url, row.name)
  }
}

</script>

<script>

export default {
  name: 'Upload'
}
</script>
<style scoped>
.upload-btn+.upload-btn {
            margin-left: 12px;
        }
</style>
