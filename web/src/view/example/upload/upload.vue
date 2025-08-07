<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="flex gap-4 p-2">
      <div class="flex-none w-64 bg-white text-slate-700 dark:text-slate-400  dark:bg-slate-900 rounded p-4">
        <el-scrollbar style="height: calc(100vh - 300px)">
          <el-tree
              :data="categories"
              node-key="id"
              :props="defaultProps"
              @node-click="handleNodeClick"
              default-expand-all
          >
            <template #default="{ node, data }">
              <div class="w-36" :class="search.classId === data.ID ? 'text-blue-500 font-bold' : ''">{{ data.name }}
              </div>
              <el-dropdown>
                <el-icon class="ml-3 text-right" v-if="data.ID > 0"><MoreFilled /></el-icon>
                <el-icon class="ml-3 text-right mt-1" v-else><Plus /></el-icon>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item @click="addCategoryFun(data)">{{ t('components.selectImage.selectImage.addCategory') }}</el-dropdown-item>
                    <el-dropdown-item @click="editCategory(data)" v-if="data.ID > 0">{{ t('components.selectImage.selectImage.editCategory') }}</el-dropdown-item>
                    <el-dropdown-item @click="deleteCategoryFun(data.ID)" v-if="data.ID > 0">{{ t('components.selectImage.selectImage.deleteCategory') }}</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </template>
          </el-tree>
        </el-scrollbar>
      </div>
      <div class="flex-1 bg-white text-slate-700 dark:text-slate-400  dark:bg-slate-900">
        <div class="gva-table-box mt-0 mb-0">
          <warning-bar :title="t('components.selectImage.selectImage.editCategoryNote')" />
          <div class="gva-btn-list gap-3">
            <upload-common :image-common="imageCommon" :classId="search.classId" @on-success="onSuccess" />
            <cropper-image :classId="search.classId" @on-success="onSuccess" />
            <QRCodeUpload :classId="search.classId" @on-success="onSuccess" />
            <upload-image
                :image-url="imageUrl"
                :file-size="512"
                :max-w-h="1080"
                :classId="search.classId"
                @on-success="onSuccess"
            />
            <el-button type="primary" icon="upload" @click="importUrlFunc">
              {{ t('view.example.upload.importURL') }}
            </el-button>
            <el-input
                v-model="search.keyword"
                class="w-72"
                :placeholder="t('view.example.upload.enterFileName')"
            />
            <el-button type="primary" icon="search" @click="onSubmit"
            >{{ t('general.search') }}
            </el-button
            >
          </div>

          <el-table :data="tableData">
            <el-table-column align="left" :label="t('view.example.upload.review')" width="100">
              <template #default="scope">
                <CustomPic pic-type="file" :pic-src="scope.row.url" preview/>
              </template>
            </el-table-column>
            <el-table-column align="left" :label="t('general.createdAt')" prop="UpdatedAt" width="180">
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
                <div class="cursor-pointer" @click="editFileNameFunc(scope.row)">
                  {{ scope.row.name }}
                </div>
              </template>
            </el-table-column>
            <el-table-column align="left" :label="t('view.example.upload.link')" prop="url" min-width="300"/>
            <el-table-column align="left" :label="t('view.example.upload.label')" prop="tag" width="80">
              <template #default="scope">
                <el-tag
                    :type="scope.row.tag?.toLowerCase() === 'jpg' ? 'info' : 'success'"
                    disable-transitions
                >{{ scope.row.tag }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column align="left" :label="t('general.operations')" width="200">
              <template #default="scope">
                <el-button
                    icon="download"
                    type="primary"
                    link
                    @click="downloadFile(scope.row)"
                >{{ t('view.example.upload.download') }}
                </el-button
                >
                <el-button
                    icon="delete"
                    type="primary"
                    link
                    @click="deleteFileFunc(scope.row)"
                >{{ t('general.delete') }}
                </el-button
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
    </div>

    <!-- 添加分类弹窗 -->
    <el-dialog v-model="categoryDialogVisible" @close="closeAddCategoryDialog" width="520"
               :title="(categoryFormData.ID === 0 ? t('general.addTo') : t('general.edit')) + t('components.selectImage.selectImage.category')"
               draggable
    >
      <el-form ref="categoryForm" :rules="rules" :model="categoryFormData" label-width="120px">
        <el-form-item :label="t('components.selectImage.selectImage.parentCategory')">
          <el-tree-select
              v-model="categoryFormData.pid"
              :data="categories"
              check-strictly
              :props="defaultProps"
              :render-after-expand="false"
              style="width: 240px"
          />
        </el-form-item>
        <el-form-item :label="t('components.selectImage.selectImage.categoryName')" prop="name">
          <el-input v-model.trim="categoryFormData.name" :placeholder="t('components.selectImage.selectImage.categoryName')"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="closeAddCategoryDialog">{{ t('general.close') }}</el-button>
        <el-button type="primary" @click="confirmAddCategory">{{ t('general.confirm') }}</el-button>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  getFileList,
  deleteFile,
  editFileName,
  importURL
} from '@/api/fileUploadAndDownload'
import {downloadImage} from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import {CreateUUID, formatDate} from '@/utils/format'
import WarningBar from '@/components/warningBar/warningBar.vue'

import {ref} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {addCategory, deleteCategory, getCategoryList} from "@/api/attachmentCategory";
import CropperImage from "@/components/upload/cropper.vue";
import QRCodeUpload from "@/components/upload/QR-code.vue";
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

const { t } = useI18n() // added by mohamed hassan to support multilingual

defineOptions({
  name: 'Upload'
})

const fullscreenLoading = ref(false)
const path = ref(import.meta.env.VITE_BASE_API)

const imageUrl = ref('')
const imageCommon = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const search = ref({
  keyword: null,
  classId: 0
})
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

const onSubmit = () => {
  search.value.classId = 0
  page.value = 1
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
  ElMessageBox.confirm(t('view.example.upload.deleteAllFilesNote'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  })
      .then(async () => {
        const res = await deleteFile(row)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: t('general.deleteSuccess')
          })
          if (tableData.value.length === 1 && page.value > 1) {
            page.value--
          }
          await getTableData()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: t('general.undeleted')
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
  ElMessageBox.prompt(t('view.example.upload.enterFileNameOrComment'), t('general.edit'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.close'),
    inputPattern: /\S/,
    inputErrorMessage: t('general.cannotBeEmpty'),
    inputValue: row.name
  })
      .then(async ({value}) => {
        row.name = value
        // console.log(row)
        const res = await editFileName(row)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: t('general.editSuccess')
          })
          await getTableData()
        }
      })
      .catch(() => {
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
    inputPlaceholder:
    t('view.example.upload.inputPlaceholder'),
    inputPattern: /\S/,
    inputErrorMessage: t('general.cannotBeEmpty')
  })
      .then(async ({value}) => {
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
              classId: search.value.classId,
              tag: url.substring(url.lastIndexOf('.') + 1),
              key: CreateUUID()
            })
          }
        })

        const res = await importURL(importData)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: t('view.example.upload.importSuccess')
          })
          await getTableData()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: t('view.example.upload.cancelImport')
        })
      })
}

const onSuccess = () => {
  search.value.keyword = null
  page.value = 1
  getTableData()
}

const defaultProps = {
  children: 'children',
  label: 'name',
  value: 'ID'
}

const categories = ref([])
const fetchCategories = async () => {
  const res = await getCategoryList()
  let data = {
    name: t('components.selectImage.selectImage.allCategories'),
    ID: 0,
    pid: 0,
    children:[]
  }
  if (res.code === 0) {
    categories.value = res.data || []
    categories.value.unshift(data)
  }
}

const handleNodeClick = (node) => {
  search.value.keyword = null
  search.value.classId = node.ID
  page.value = 1
  getTableData()
}

const categoryDialogVisible = ref(false)
const categoryFormData = ref({
  ID: 0,
  pid: 0,
  name: ''
})

const categoryForm = ref(null)
const rules = ref({
  name: [
    {required: true, message: t('components.selectImage.selectImage.enterCategoryName'), trigger: 'blur'},
    {max: 20, message: t('components.selectImage.selectImage.categoryNameNote'), trigger: 'blur'}
  ]
})

const addCategoryFun = (category) => {
  categoryDialogVisible.value = true
  categoryFormData.value.ID = 0
  categoryFormData.value.pid = category.ID
}

const editCategory = (category) => {
  categoryFormData.value = {
    ID: category.ID,
    pid: category.pid,
    name: category.name
  }
  categoryDialogVisible.value = true
}

const deleteCategoryFun = async (id) => {
  const res = await deleteCategory({id: id})
  if (res.code === 0) {
    ElMessage.success({type: 'success', message: t('general.deleteSuccess')})
    await fetchCategories()
  }
}

const confirmAddCategory = async () => {
  categoryForm.value.validate(async valid => {
    if (valid) {
      const res = await addCategory(categoryFormData.value)
      if (res.code === 0) {
        ElMessage({type: 'success', message: t('view.systemTools.system.operationSuccess')})
        await fetchCategories()
        closeAddCategoryDialog()
      }
    }
  })
}

const closeAddCategoryDialog = () => {
  categoryDialogVisible.value = false
  categoryFormData.value = {
    ID: 0,
    pid: 0,
    name: ''
  }
}

fetchCategories()
</script>
