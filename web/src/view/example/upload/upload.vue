<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="gva-table-box">
      <warning-bar title="点击“文件名/备注”可以编辑文件名或者备注内容。" />
      <div class="gva-btn-list gap-3">
        <upload-common :image-common="imageCommon" @on-success="getTableData" />
        <upload-image
          :image-url="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          @on-success="getTableData"
        />
        <el-button type="primary" icon="upload" @click="importUrlFunc">
          导入URL
        </el-button>
        <el-input
          v-model="search.keyword"
          class="w-72"
          placeholder="请输入文件名或备注"
        />
        <el-button type="primary" icon="search" @click="getTableData"
          >查询</el-button
        >
      </div>

      <el-table :data="tableData">
        <el-table-column align="left" label="预览" width="100">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.url" preview />
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" prop="UpdatedAt" width="180">
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="文件名/备注"
          prop="name"
          width="180"
        >
          <template #default="scope">
            <div class="name" @click="editFileNameFunc(scope.row)">
              {{ scope.row.name }}
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="链接" prop="url" min-width="300" />
        <el-table-column align="left" label="标签" prop="tag" width="100">
          <template #default="scope">
            <el-tag
              :type="scope.row.tag?.toLowerCase() === 'jpg' ? 'info' : 'success'"
              disable-transitions
              >{{ scope.row.tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button
              icon="download"
              type="primary"
              link
              @click="downloadFile(scope.row)"
              >下载</el-button
            >
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteFileFunc(scope.row)"
              >删除</el-button
            >
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
  import {
    getFileList,
    deleteFile,
    editFileName,
    importURL
  } from '@/api/fileUploadAndDownload'
  import { downloadImage } from '@/utils/downloadImg'
  import CustomPic from '@/components/customPic/index.vue'
  import UploadImage from '@/components/upload/image.vue'
  import UploadCommon from '@/components/upload/common.vue'
  import { CreateUUID, formatDate } from '@/utils/format'
  import WarningBar from '@/components/warningBar/warningBar.vue'

  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  defineOptions({
    name: 'Upload'
  })

  const path = ref(import.meta.env.VITE_BASE_API)

  const imageUrl = ref('')
  const imageCommon = ref('')

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const search = ref({})
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
  const getTableData = async () => {
    const table = await getFileList({
      page: page.value,
      pageSize: pageSize.value,
      ...search.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }
  getTableData()

  const deleteFileFunc = async (row) => {
    ElMessageBox.confirm('此操作将永久删除文件, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async () => {
        const res = await deleteFile(row)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '删除成功!'
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
          message: '已取消删除'
        })
      })
  }

  const downloadFile = (row) => {
    if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
      downloadImage(row.url, row.name)
    } else {
      downloadImage(path.value + '/' + row.url, row.name)
    }
  }

  /**
   * 编辑文件名或者备注
   * @param row
   * @returns {Promise<void>}
   */
  const editFileNameFunc = async (row) => {
    ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /\S/,
      inputErrorMessage: '不能为空',
      inputValue: row.name
    })
      .then(async ({ value }) => {
        row.name = value
        // console.log(row)
        const res = await editFileName(row)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '编辑成功!'
          })
          await getTableData()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消修改'
        })
      })
  }

  /**
   * 导入URL
   */
  const importUrlFunc = () => {
    ElMessageBox.prompt('格式：文件名|链接或者仅链接。', '导入', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputType: 'textarea',
      inputPlaceholder:
        '我的图片|https://my-oss.com/my.png\nhttps://my-oss.com/my_1.png',
      inputPattern: /\S/,
      inputErrorMessage: '不能为空'
    })
      .then(async ({ value }) => {
        let data = value.split('\n')
        let importData = []
        data.forEach((item) => {
          let oneData = item.trim().split('|')
          let url, name
          if (oneData.length > 1) {
            name = oneData[0].trim()
            url = oneData[1]
          } else {
            url = oneData[0].trim()
            let str = url.substring(url.lastIndexOf('/') + 1)
            name = str.substring(0, str.lastIndexOf('.'))
          }
          if (url) {
            importData.push({
              name: name,
              url: url,
              tag: url.substring(url.lastIndexOf('.') + 1),
              key: CreateUUID()
            })
          }
        })

        const res = await importURL(importData)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '导入成功!'
          })
          await getTableData()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消导入'
        })
      })
  }
</script>

<style scoped>
  .name {
    cursor: pointer;
  }
</style>
