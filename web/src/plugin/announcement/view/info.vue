<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item :label="t('general.createDate')" prop="createdAt">
          <template #label>
            <span>
              {{ t('general.createDate') }}
              <el-tooltip :content="t('general.searchDesc')">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" :placeholder="t('general.startData')" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" :placeholder="t('general.endData')" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false" />
        </el-form-item>


        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            {{ t('general.search') }}
          </el-button>
          <el-button icon="refresh" @click="onReset">
            {{ t('general.reset') }}
          </el-button>
          <el-button v-if="!showAllQuery" link type="primary" icon="arrow-down" @click="showAllQuery=true">
            {{ t('general.expand') }}
          </el-button>
          <el-button v-else link type="primary" icon="arrow-up" @click="showAllQuery=false">
            {{ t('general.collapse') }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">
          {{ t('general.add') }}
        </el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">
          {{ t('general.delete') }}
        </el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" :label="t('general.createdAt')" prop="createdAt" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column align="left" :label="t('plugins.announcement.info.title')" prop="title" width="120" />
        <el-table-column :label="t('plugins.announcement.info.content')" prop="content" width="200">
          <template #default="scope">
           {{ t('plugins.announcement.info.richTextContent') }}
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('plugins.announcement.info.author')" prop="userID" width="120">
          <template #default="scope">
            <span>{{ filterDataSource(dataSource.userID,scope.row.userID) }}</span>
          </template>
        </el-table-column>
        <el-table-column :label="t('plugins.announcement.info.attachments')" prop="attachments" width="200">
          <template #default="scope">
            <div class="file-list">
              <el-tag v-for="file in scope.row.attachments" :key="file.uid" @click="downloadFile(file.url)">
                {{ file.name }}
              </el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.operations')" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateInfoFunc(scope.row)">
              {{ t('general.edit') }}
            </el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">
              {{ t('general.delete') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer v-model="dialogFormVisible" destroy-on-close size="800" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type==='create'? t('general.addTo') :t('general.modify') }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">
              {{ t('general.confirm') }}
            </el-button>
            <el-button @click="closeDialog">
              {{ t('general.cancel') }}
            </el-button>
          </div>
        </div>
      </template>

      <el-form ref="elFormRef" :model="formData" label-position="top" :rules="rule" label-width="80px">
        <el-form-item :label="t('plugins.announcement.info.title') + ':'" prop="title">
          <el-input v-model="formData.title" :clearable="true" :placeholder="t('plugins.announcement.info.enterTitleNote')" />
        </el-form-item>
        <el-form-item :label="t('plugins.announcement.info.content') + ':'" prop="content">
          <RichEdit v-model="formData.content" />
        </el-form-item>
        <el-form-item :label="t('plugins.announcement.info.author') + ':'" prop="userID">
          <el-select v-model="formData.userID" :placeholder="t('plugins.announcement.info.selectAuthorNote')" style="width:100%" :clearable="true">
            <el-option v-for="(item,key) in dataSource.userID" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('plugins.announcement.info.attachments') + ':'" prop="attachments">
          <SelectFile v-model="formData.attachments" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
    getInfoDataSource,
  createInfo,
  deleteInfo,
  deleteInfoByIds,
  updateInfo,
  findInfo,
  getInfoList
} from '@/plugin/announcement/api/info'
import { getUrl } from '@/utils/image'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'

// 全量引入格式化工具 请按需保留
import {  formatDate, filterDataSource } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

const { t } = useI18n() // added by mohamed hassan to support multilingual

defineOptions({
    name: 'Info'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        title: '',
        content: '',
        userID: undefined,
        attachments: [],
        })
  const dataSource = ref([])
  const getDataSourceFunc = async()=>{
    const res = await getInfoDataSource()
    if (res.code === 0) {
      dataSource.value = res.data
    }
  }
  getDataSourceFunc()



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error(t('general.placeInputEndData')))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error(t('general.placeInputStartData')))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error(t('general.startDataMustBeforeEndData')))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
        confirmButtonText: t('general.confirm'),
        cancelButtonText: t('general.cancel'),
        type: 'warning'
    }).then(() => {
            deleteInfoFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: t('general.selectDataToDelete')
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteInfoByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateInfoFunc = async(row) => {
    const res = await findInfo({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteInfoFunc = async (row) => {
    const res = await deleteInfo({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: t('general.deleteSuccess')
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        title: '',
        content: '',
        userID: undefined,
        attachments: [],
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createInfo(formData.value)
                  break
                case 'update':
                  res = await updateInfo(formData.value)
                  break
                default:
                  res = await createInfo(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: t('general.createUpdateSuccess')
                })
                closeDialog()
                getTableData()
              }
      })
}

const downloadFile = (url) => {
    window.open(getUrl(url), '_blank')
}

</script>

<style>

.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}

</style>
