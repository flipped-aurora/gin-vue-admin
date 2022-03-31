<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="goAutoCode(null)">{{ t('general.add') }}</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" :label="t('general.createdAt')" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('autoCode.structName')" min-width="150" prop="structName" />
        <el-table-column align="left" :label="t('autoCode.structChineseName')" min-width="150" prop="structCNName" />
        <el-table-column align="left" :label="t('autoCode.tableName')" min-width="150" prop="tableName" />
        <el-table-column align="left" :label="t('autoCodeAdmin.rollBackMark')" min-width="150" prop="flag">
          <template #default="scope">
            <el-tag
              v-if="scope.row.flag"
              type="danger"
              size="mini"
              effect="dark"
            >
              {{ t('autoCodeAdmin.rolledBack') }}
            </el-tag>
            <el-tag
              v-else
              size="mini"
              type="success"
              effect="dark"
            >
              {{ t('autoCodeAdmin.notRolledBack') }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" :lable="t('general.operations')" min-width="240">
          <template #default="scope">
            <div>
              <el-button size="small" type="text" :disabled="scope.row.flag === 1" @click="rollbackFunc(scope.row,true)">{{ t('autoCodeAdmin.rollBackDeleteTable') }}</el-button>
              <el-button size="small" type="text" :disabled="scope.row.flag === 1" @click="rollbackFunc(scope.row,false)">{{ t('autoCodeAdmin.rollBackWithoutDeleteTable') }}</el-button>
              <el-button size="small" type="text" @click="goAutoCode(scope.row)">{{ t('autoCodeAdmin.reuse') }}</el-button>
              <el-button size="small" type="text" @click="deleteRow(scope.row)">{{ t('general.delete') }}</el-button>
            </div>
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

<script>
export default {
  name: 'AutoCodeAdmin',
}
</script>

<script setup>
import { getSysHistory, rollback, delSysHistory } from '@/api/autoCode.js'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { formatDate } from '@/utils/format'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

const router = useRouter()

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
  const table = await getSysHistory({
    page: page.value,
    pageSize: pageSize.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const deleteRow = async(row) => {
  ElMessageBox.confirm(t('autoCodeAdmin.deleteHistoryConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  }).then(async() => {
    const res = await delSysHistory({ id: Number(row.ID) })
    if (res.code === 0) {
      ElMessage.success(t('general.deleteSuccess'))
      getTableData()
    }
  })
}
const rollbackFunc = async(row, flag) => {
  ElMessageBox.confirm(t('autoCodeAdmin.rollbackConfirm') + `${flag ? t('autoCodeAdmin.includeDBTables') : ' ,'}` + t('autoCodeAdmin.rollBackContinue'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  }).then(async() => {
    const res = await rollback({ id: Number(row.ID), deleteTable: flag })
    if (res.code === 0) {
      ElMessage.success(t('autoCodeAdmin.rollbackSuccess'))
      getTableData()
    }
  })
}
const goAutoCode = (row) => {
  if (row) {
    router.push({ name: 'autoCodeEdit', params: {
      id: row.ID
    }})
  } else {
    router.push({ name: 'autoCode' })
  }
}

</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>
