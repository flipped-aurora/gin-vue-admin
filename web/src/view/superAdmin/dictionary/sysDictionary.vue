<template>
  <div>
    <warning-bar
      :title="t('view.dictionary.sysDictionary.dictNote')"
    />
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item :label="t('view.dictionary.sysDictionary.dictName')">
          <el-input v-model="searchInfo.name" :placeholder="t('general.searchCriteria')" />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionary.dictNameEn')">
          <el-input v-model="searchInfo.type" :placeholder="t('general.searchCriteria')" />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionary.status')" prop="status">
          <el-select v-model="searchInfo.status" clear :placeholder="t('general.pleaseSelect')">
            <el-option key="true" :label="t('general.yes')" value="true" />
            <el-option key="false" :label="t('general.no')" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('general.description')">
          <el-input v-model="searchInfo.desc" :placeholder="t('general.searchCriteria')" />
        </el-form-item>
        <el-form-item>
          <el-button

            type="primary"
            icon="search"
            @click="onSubmit"
          >{{ t('general.search') }}</el-button>
          <el-button

            icon="refresh"
            @click="onReset"
          >{{ t('general.reset') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button

          type="primary"
          icon="plus"
          @click="openDialog"
        >{{ t('general.add') }}</el-button>
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
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          :label="t('view.dictionary.sysDictionary.dictName')"
          prop="name"
          width="200"
        />

        <el-table-column
          align="left"
          :label="t('view.dictionary.sysDictionary.dictNameEn')"
          prop="type"
          width="200"
        />

        <el-table-column align="left" :label="t('view.dictionary.sysDictionary.status')" prop="status" width="120">
          <template #default="scope">{{
            formatBoolean(scope.row.status)
          }}</template>
        </el-table-column>

        <el-table-column align="left" :label="t('general.description')" prop="desc" width="280" />

        <el-table-column align="left" :label="t('general.operations')">
          <template #default="scope">
            <el-button

              icon="document"
              type="primary"
              link
              @click="toDetail(scope.row)"
            >{{ t('view.dictionary.sysDictionary.details') }}</el-button>
            <el-button

              icon="edit"
              type="primary"
              link
              @click="updateSysDictionaryFunc(scope.row)"
            >{{ t('general.change') }}</el-button>
            <el-popover
              v-model="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>{{ t('general.deleteConfirm') }}</p>
              <div style="text-align: right; margin-top: 8px">
                <el-button

                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >{{ t('general.cancel') }}</el-button>
                <el-button
                  type="primary"

                  @click="deleteSysDictionaryFunc(scope.row)"
                >{{ t('general.confirm') }}</el-button>
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"

                  style="margin-left: 10px"
                  @click="scope.row.visible = true"
                >{{ t('general.delete') }}</el-button>
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
      :title="type==='create'?t('view.dictionary.sysDictionary.addDictionary'):t('view.dictionary.sysDictionary.editDictionary')"
    >
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="110px"
      >
        <el-form-item :label="t('view.dictionary.sysDictionary.dictName')" prop="name">
          <el-input
            v-model="formData.name"
            :placeholder="t('view.dictionary.sysDictionary.enterDictName')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionary.dictNameEn')" prop="type">
          <el-input
            v-model="formData.type"
            :placeholder="t('view.dictionary.sysDictionary.enterDictNameEn')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item :label="t('view.dictionary.sysDictionary.status')" prop="status" required>
          <el-switch
            v-model="formData.status"
            :active-text="t('general.enable')"
            :inactive-text="t('general.disable')"
          />
        </el-form-item>
        <el-form-item :label="t('general.description')" prop="desc">
          <el-input
            v-model="formData.desc"
            :placeholder="t('view.dictionary.sysDictionary.enterDescription')"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">{{ t('general.close') }}</el-button>
          <el-button

            type="primary"
            @click="enterDialog"
          >{{ t('general.confirm') }}</el-button>
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
} from '@/api/sysDictionary' // 此处请自行替换地址
import WarningBar from '@/components/warningBar/warningBar.vue'
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
      message: t('view.dictionary.sysDictionary.enterDictName'),
      trigger: 'blur',
    },
  ],
  type: [
    {
      required: true,
      message: t('view.dictionary.sysDictionary.enterDictNameEn'),
      trigger: 'blur',
    },
  ],
  desc: [
    {
      required: true,
      message: t('view.dictionary.sysDictionary.enterDescription'),
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

const toDetail = (row) => {
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
  const res = await findSysDictionary({ ID: row.ID, status: row.status })
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
      message: t('general.deleteSuccess'),
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
      ElMessage.success('操作成功')
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
