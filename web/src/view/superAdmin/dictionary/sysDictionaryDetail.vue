<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item :label="t('view.dictionary.sysDictionaryDetail.displayValue')">
          <el-input v-model="searchInfo.label" :placeholder="t('general.searchCriteria')" />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionaryDetail.dictValue')">
          <el-input v-model="searchInfo.value" :placeholder="t('general.searchCriteria')" min="-2147483648" max="2147483647" />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionaryDetail.enabledStatus')" prop="status">
          <el-select v-model="searchInfo.status" :placeholder="t('general.pleaseSelect')">
            <el-option key="true" :label="t('general.yes')" value="true" />
            <el-option key="false" :label="t('general.no')" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">{{ t('general.search') }}</el-button>
          <el-button icon="refresh" @click="onReset">{{ t('general.reset') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">{{ t('view.dictionary.sysDictionaryDetail.addDictEntry') }}</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" :label="t('general.createdAt')" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" :label="t('view.dictionary.sysDictionaryDetail.displayValue')" prop="label" width="120" />

        <el-table-column align="left" :label="t('view.dictionary.sysDictionaryDetail.dictValue')" prop="value" width="140" />

        <el-table-column align="left" :label="t('view.dictionary.sysDictionaryDetail.enabledStatus')" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>

        <el-table-column align="left" :label="t('general.order')" prop="sort" width="120" />

        <el-table-column align="left" :label="t('general.operations')">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="updateSysDictionaryDetailFunc(scope.row)">{{ t('general.change') }}</el-button>
            <el-popover v-model="scope.row.visible" placement="top" width="160">
              <p>{{ t('general.deleteConfirm') }}</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="scope.row.visible = false">{{ t('general.cancel') }}</el-button>
                <el-button type="primary" @click="deleteSysDictionaryDetailFunc(scope.row)">{{ t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button type="primary" link icon="delete" @click="scope.row.visible = true">{{ t('general.delete') }}</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?t('view.dictionary.sysDictionaryDetail.addDictionaryItem'):t('view.dictionary.sysDictionaryDetail.editDictionaryItem')">
      <el-form ref="dialogForm" :model="formData" :rules="rules" label-width="110px">
        <el-form-item :label="t('view.dictionary.sysDictionaryDetail.displayValue')" prop="label">
          <el-input
            v-model="formData.label"
            :placeholder="t('view.dictionary.sysDictionaryDetail.enterDisplayValue')"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionaryDetail.dictValue')" prop="value">
          <el-input-number
            v-model.number="formData.value"
            step-strictly
            :step="1"
            :placeholder="t('view.dictionary.sysDictionaryDetail.enterDictValue')"
            clearable
            :style="{width: '100%'}"
            min="-2147483648"
            max="2147483647"
          />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionaryDetail.enabledStatus')" prop="status" required>
          <el-switch v-model="formData.status" :active-text="t('general.enable')" :inactive-text="t('general.disable')" />
        </el-form-item>
        <el-form-item :label="t('general.order')" prop="sort">
          <el-input-number v-model.number="formData.sort" :placeholder="t('view.dictionary.sysDictionaryDetail.enabledStatus')" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">{{ t('general.close') }}</el-button>
          <el-button type="primary" @click="enterDialog">{{ t('general.confirm') }}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'SysDictionaryDetail'
}
</script>

<script setup>
import {
  createSysDictionaryDetail,
  deleteSysDictionaryDetail,
  updateSysDictionaryDetail,
  findSysDictionaryDetail,
  getSysDictionaryDetailList
} from '@/api/sysDictionaryDetail' // 此处请自行替换地址
import { ref } from 'vue'
import { useRoute, onBeforeRouteUpdate } from 'vue-router'
import { ElMessage } from 'element-plus'
import { formatBoolean, formatDate } from '@/utils/format'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const route = useRoute()

onBeforeRouteUpdate((to) => {
  if (to.name === 'dictionaryDetail') {
    searchInfo.value.sysDictionaryID = to.params.id
    getTableData()
  }
})

const formData = ref({
  label: null,
  value: null,
  status: true,
  sort: null
})
const rules = ref({
  label: [
    {
      required: true,
      message: t('view.dictionary.sysDictionaryDetail.enterDisplayValue'),
      trigger: 'blur'
    }
  ],
  value: [
    {
      required: true,
      message: t('view.dictionary.sysDictionaryDetail.enterDictValue'),
      trigger: 'blur'
    }
  ],
  sort: [
    {
      required: true,
      message: t('general.order'),
      trigger: 'blur'
    }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({ sysDictionaryID: Number(route.params.id) })
const onReset = () => {
  searchInfo.value = { sysDictionaryID: Number(route.params.id) }
  getTableData()
}

// 条件搜索前端看此方法
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
  getTableData()
}

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
  const table = await getSysDictionaryDetailList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const type = ref('')
const dialogFormVisible = ref(false)
const updateSysDictionaryDetailFunc = async(row) => {
  const res = await findSysDictionaryDetail({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reSysDictionaryDetail
    dialogFormVisible.value = true
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    label: null,
    value: null,
    status: true,
    sort: null,
    sysDictionaryID: ''
  }
}
const deleteSysDictionaryDetailFunc = async(row) => {
  row.visible = false
  const res = await deleteSysDictionaryDetail({ ID: row.ID })
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

const dialogForm = ref(null)
const enterDialog = async() => {
  formData.value.sysDictionaryID = Number(route.params.id)
  dialogForm.value.validate(async valid => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysDictionaryDetail(formData.value)
        break
      case 'update':
        res = await updateSysDictionaryDetail(formData.value)
        break
      default:
        res = await createSysDictionaryDetail(formData.value)
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
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

</script>

<style>
</style>
