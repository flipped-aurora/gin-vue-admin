<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="gva-table-box">
      <warning-bar
        :title="t('view.example.upload.editFileNote')"
      />
      <div class="gva-btn-list gap-3">
        <upload-common
          :image-common="imageCommon"
          @on-success="getTableData"
        />
        <upload-image
          :image-url="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          @on-success="getTableData"
        />
        <el-button
            type="primary"
            icon="upload"
            @click="importUrlFunc"
        >
          {{ t('view.example.upload.importURL') }}
        </el-button>
        <el-input
          v-model="search.keyword"
          class="w-72"
          :placeholder="t('view.example.upload.enterFileName')"
        />
        <el-button
          type="primary"
          icon="search"
          @click="getTableData"
        >{{ t('general.search') }}</el-button>
      </div>

      <el-table :data="tableData">
        <el-table-column
          align="left"
          :label="t('view.example.upload.review')"
          width="100"
        >
          <template #default="scope">
            <CustomPic
              pic-type="file"
              :pic-src="scope.row.url"
              preview
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('general.createdAt')"
          prop="UpdatedAt"
          width="180"
        >
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.example.upload.fileNameComments')"
          prop="name"
          width="180"
        >
          <template #default="scope">
            <div
              class="name"
              @click="editFileNameFunc(scope.row)"
            >{{ scope.row.name }}</div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.example.upload.link')"
          prop="url"
          min-width="300"
        />
        <el-table-column
          align="left"
          :label="t('view.example.upload.label')"
          prop="tag"
          width="100"
        >
          <template #default="scope">
            <el-tag
              :type="scope.row.tag === 'jpg' ? 'info' : 'success'"
              disable-transitions
            >{{ scope.row.tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('general.operations')"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="download"
              type="primary"
              link
              @click="downloadFile(scope.row)"
            >{{ t('view.example.upload.download') }}</el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteFileFunc(scope.row)"
            >{{ t('general.delete') }}</el-button>
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
import {getFileList, deleteFile, editFileName, importURL} from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import {CreateUUID, formatDate} from '@/utils/format'
import WarningBar from '@/components/warningBar/warningBar.vue'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: 'Upload',
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
const getTableData = async() => {
  const table = await getFileList({ page: page.value, pageSize: pageSize.value, ...search.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const deleteFileFunc = async(row) => {
  ElMessageBox.confirm(t('view.example.upload.deleteAllFilesNote'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning',
  })
    .then(async() => {
      const res = await deleteFile(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess'),
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
        message: t('general.undeleted'),
      })
    })
}

const downloadFile = (row) => {
  if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    downloadImage(row.url, row.name)
  } else {
    debugger
    downloadImage(path.value + '/' + row.url, row.name)
  }
}

/**
 * 编辑文件名或者备注
 * @param row
 * @returns {Promise<void>}
 */
const editFileNameFunc = async(row) => {
  ElMessageBox.prompt(t('view.example.upload.enterFileNameOrComment'), t('general.edit'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.close'),
    inputPattern: /\S/,
    inputErrorMessage: t('general.cannotBeEmpty'),
    inputValue: row.name
  }).then(async({ value }) => {
    row.name = value
    // console.log(row)
    const res = await editFileName(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: t('general.editSuccess'),
      })
      await getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: t('general.cancelModification')
    })
  })
}

/**
 * 导入URL
 */
const importUrlFunc = () => {
  ElMessageBox.prompt(t('view.example.upload.formatNote'), t('general.import'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    inputType: 'textarea',
    inputPlaceholder: t('view.example.upload.inputPlaceholder'),
    inputPattern: /\S/,
    inputErrorMessage: t('general.cannotBeEmpty'),
  }).then(async({ value }) => {
    let data = value.split('\n')
    let importData = []
    data.forEach(item => {
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
          tag: url.substring(url.lastIndexOf(".") + 1),
          key: CreateUUID()
        })
      }
    })

    const res = await importURL(importData)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: t('view.example.upload.importSuccess'),
      })
      await getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: t('view.example.upload.cancelImport'),
    })
  })
}
</script>

<style scoped>
.name {
  cursor: pointer;
}
</style>
