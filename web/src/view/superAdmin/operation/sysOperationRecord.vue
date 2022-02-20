<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item :label="t('operationRecord.requestMethod')">
          <el-input v-model="searchInfo.method" :placeholder="t('general.searchCondition')" />
        </el-form-item>
        <el-form-item :label="t('operationRecord.requestPath')">
          <el-input v-model="searchInfo.path" :placeholder="t('general.searchCondition')" />
        </el-form-item>
        <el-form-item :label="t('operationRecord.statusCode')">
          <el-input v-model="searchInfo.status" :placeholder="t('general.searchCondition')" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">{{t('general.search')}}</el-button>
          <el-button size="small" icon="refresh" @click="onReset">{{t('general.reset')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">

        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>{{t('general.confirmDelete')}}</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="text" @click="deleteVisible = false">{{t('general.cancel')}}</el-button>
            <el-button size="small" type="primary" @click="onDelete">{{t('general.yes')}}</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length">{{t('general.delete')}}</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column align="left" type="selection" width="55" />
        <el-table-column align="left" :label="t('operationRecord.operator')" width="140">
          <template #default="scope">
            <div>{{ scope.row.user.userName }}({{ scope.row.user.nickName }})</div>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.date')" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('operationRecord.statusCode')" prop="status" width="110">
          <template #default="scope">
            <div>
              <el-tag type="success">{{ scope.row.status }}</el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('operationRecord.requestIP')" prop="ip" width="120" />
        <el-table-column align="left" :label="t('operationRecord.requestMethod')" prop="method" width="130" />
        <el-table-column align="left" :label="t('operationRecord.requestPath')" prop="path" width="240" />
        <el-table-column align="left" :label="t('operationRecord.ask')" prop="path" width="80">
          <template #default="scope">
            <div>
              <el-popover v-if="scope.row.body" placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.body) }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>

              <span v-else>without</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.response')" prop="path" width="90">
          <template #default="scope">
            <div>
              <el-popover v-if="scope.row.resp" placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.resp) }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>
              <span v-else>without</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.action')">
          <template #default="scope">
            <el-popover v-model:visible="scope.row.visible" placement="top" width="160">
              <p>{{t('general.confirmDelete')}}</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="small" type="text" @click="scope.row.visible = false">{{t('general.cancel')}}</el-button>
                <el-button size="small" type="primary" @click="deleteSysOperationRecordFunc(scope.row)">{{t('general.yes')}}</el-button>
              </div>
              <template #reference>
                <el-button icon="delete" size="small" type="text" @click="scope.row.visible = true">{{t('general.delete')}}</el-button>
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
  </div>
</template>

<script setup>
import {
  deleteSysOperationRecord,
  getSysOperationRecordList,
  deleteSysOperationRecordByIds
} from '@/api/sysOperationRecord' // 此处请自行替换地址
import { formatDate } from '@/utils/format'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

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
  const table = await getSysOperationRecordList({
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

const deleteVisible = ref(false)
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}
const onDelete = async() => {
  const ids = []
  multipleSelection.value &&
        multipleSelection.value.forEach(item => {
          ids.push(item.ID)
        })
  const res = await deleteSysOperationRecordByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Successfully deleted'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}
const deleteSysOperationRecordFunc = async(row) => {
  row.visible = false
  const res = await deleteSysOperationRecord({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Successfully deleted'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}
const fmtBody = (value) => {
  try {
    return JSON.parse(value)
  } catch (err) {
    return value
  }
}

</script>

<script>

export default {
  name: 'SysOperationRecord'
}
</script>

<style lang="scss">
.table-expand {
  padding-left: 60px;
  font-size: 0;
  label {
    width: 90px;
    color: #99a9bf;
    .el-form-item {
      margin-right: 0;
      margin-bottom: 0;
      width: 50%;
    }
  }
}
.popover-box {
  background: #112435;
  color: #f08047;
  height: 600px;
  width: 420px;
  overflow: auto;
}
.popover-box::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>

// without
