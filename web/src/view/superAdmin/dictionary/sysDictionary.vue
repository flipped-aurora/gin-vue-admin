<template>
  <div>
    <warning-bar
      :title="t('dic.warningBarTitle')"
    />
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item :label="t('dic.nameMiddle')">
          <el-input v-model="searchInfo.name" :placeholder="t('general.searchCondition')" />
        </el-form-item>
        <el-form-item :label="t('dic.nameEnglish')">
          <el-input v-model="searchInfo.type" :placeholder="t('general.searchCondition')" />
        </el-form-item>
        <el-form-item :label="t('general.status')" prop="status">
          <el-select v-model="searchInfo.status" clear :placeholder="t('general.choose')">
            <el-option key="true" label="Yes" value="true" />
            <el-option key="false" label="No" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('general.des')">
          <el-input v-model="searchInfo.desc" :placeholder="t('general.searchCondition')" />
        </el-form-item>
        <el-form-item>
          <el-button
            size="small"
            type="primary"
            icon="search"
            @click="onSubmit"
          >{{t('general.date')}}</el-button>
          <el-button
            size="small"
            icon="refresh"
            @click="onReset"
          >{{t('general.reset')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          size="small"
          type="primary"
          icon="plus"
          @click="openDialog"
        >{{t('general.add')}}</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" :label="t('general.date')" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          :label="t('dic.nameMiddle')"
          prop="name"
          width="160"
        />

        <el-table-column
          align="left"
          :label="t('dic.nameEnglish')"
          prop="type"
          width="120"
        />

        <el-table-column align="left" :label="t('general.status')" prop="status" width="120">
          <template #default="scope">{{
            formatBoolean(scope.row.status)
          }}</template>
        </el-table-column>

        <el-table-column align="left" :label="t('general.des')" prop="desc" width="280" />

        <el-table-column align="left" :label="t('general.action')">
          <template #default="scope">
            <el-button
              size="small"
              icon="document"
              type="text"
              @click="toDetile(scope.row)"
            >{{t('general.detail')}}</el-button>
            <el-button
              size="small"
              icon="edit"
              type="text"
              @click="updateSysDictionaryFunc(scope.row)"
            >{{t('general.change')}}</el-button>
            <el-popover
              v-model:visible="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>{{t('general.confirmDelete')}}</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button
                  size="small"
                  type="text"
                  @click="scope.row.visible = false"
                >Cancel</el-button>
                <el-button
                  type="primary"
                  size="small"
                  @click="deleteSysDictionaryFunc(scope.row)"
                >{{t('general.yes')}}</el-button>
              </div>
              <template #reference>
                <el-button
                  type="text"
                  icon="delete"
                  size="small"
                  style="margin-left: 10px"
                  @click="scope.row.visible = true"
                >{{t('general.delete')}}</el-button>
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="t('general.operations')"
    >
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        size="medium"
        label-width="110px"
      >
        <el-form-item :label="t('dic.nameMiddle')" prop="name">
          <el-input
            v-model="formData.name"
            :placeholder="t('dic.pleaseEnterADictionaryNameIn')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item :label="t('dic.nameEnglish')" prop="type">
          <el-input
            v-model="formData.type"
            :placeholder="t('dic.pleaseEnterADictionaryNameEn')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item :label="t('general.status')" prop="status" required>
          <el-switch
            v-model="formData.status"
            :active-text="t('general.turOn')"
            :inactive-text="t('general.disable')"
          />
        </el-form-item>
        <el-form-item :label="t('general.des')" prop="desc">
          <el-input
            v-model="formData.desc"
            :placeholder="t('dic.pleaseEnterADescription')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">{{t('general.cancel')}}</el-button>
          <el-button
            size="small"
            type="primary"
            @click="enterDialog"
          >{{t('general.yes')}}</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'SysDictionary',
}
</script>

<script setup>
import {
  createSysDictionary,
  deleteSysDictionary,
  updateSysDictionary,
  findSysDictionary,
  getSysDictionaryList,
} from '@/api/sysDictionary' //  此处请自行替换地址
import warningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { formatBoolean, formatDate } from '@/utils/format'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const router = useRouter()

const formData = ref({
  name: null,
  type: null,
  status: true,
  desc: null,
})
const rules = ref({
  name: [
    {
      required: true,
      message: t('dic.pleaseEnterADictionaryNameIn'),
      trigger: 'blur',
    },
  ],
  type: [
    {
      required: true,
      message: t('dic.pleaseEnterADictionaryNameEn'),
      trigger: 'blur',
    },
  ],
  desc: [
    {
      required: true,
      message: t('dic.pleaseEnterADescription'),
      trigger: 'blur',
    },
  ],
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}

// Conditional search front-end to see this method
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
  getTableData()
}

// pagination
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// Search
const getTableData = async() => {
  const table = await getSysDictionaryList({
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

const toDetile = (row) => {
  router.push({
    name: 'dictionaryDetail',
    params: {
      id: row.ID,
    },
  })
}

const dialogFormVisible = ref(false)
const type = ref('')
const updateSysDictionaryFunc = async(row) => {
  const res = await findSysDictionary({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysDictionary
    dialogFormVisible.value = true
  }
}
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: null,
    type: null,
    status: true,
    desc: null,
  }
}
const deleteSysDictionaryFunc = async(row) => {
  row.visible = false
  const res = await deleteSysDictionary({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Successfully Deleted',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const dialogForm = ref(null)
const enterDialog = async() => {
  dialogForm.value.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysDictionary(formData.value)
        break
      case 'update':
        res = await updateSysDictionary(formData.value)
        break
      default:
        res = await createSysDictionary(formData.value)
        break
    }
    if (res.code === 0) {
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

<style></style>
