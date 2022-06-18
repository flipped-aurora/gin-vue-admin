<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="goAutoCode(null)">新增</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="结构体名" min-width="150" prop="structName" />
        <el-table-column align="left" label="结构体描述" min-width="150" prop="structCNName" />
        <el-table-column align="left" label="表名称" min-width="150" prop="tableName" />
        <el-table-column align="left" label="回滚标记" min-width="150" prop="flag">
          <template #default="scope">
            <el-tag
              v-if="scope.row.flag"
              type="danger"
              size="small"
              effect="dark"
            >
              已回滚
            </el-tag>
            <el-tag
              v-else
              size="small"
              type="success"
              effect="dark"
            >
              未回滚
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="240">
          <template #default="scope">
            <div>
              <el-button size="small" type="primary" link :disabled="scope.row.flag === 1" @click="rollbackFunc(scope.row,true)">回滚(删表)</el-button>
              <el-button size="small" type="primary" link :disabled="scope.row.flag === 1" @click="rollbackFunc(scope.row,false)">回滚(不删表)</el-button>
              <el-button size="small" type="primary" link @click="goAutoCode(scope.row)">复用</el-button>
              <el-button size="small" type="primary" link @click="deleteRow(scope.row)">删除</el-button>
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
  ElMessageBox.confirm('此操作将删除本历史, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await delSysHistory({ id: Number(row.ID) })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}
const rollbackFunc = async(row, flag) => {
  if (flag) {
    ElMessageBox.confirm(`此操作将删除自动创建的文件和api（会删除表！！！）, 是否继续?`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async() => {
      ElMessageBox.confirm(`此操作将删除自动创建的文件和api（会删除表！！！）, 请继续确认！！！`, '会删除表', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async() => {
        ElMessageBox.confirm(`此操作将删除自动创建的文件和api（会删除表！！！）, 请继续确认！！！`, '会删除表', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async() => {
          const res = await rollback({ id: Number(row.ID), deleteTable: flag })
          if (res.code === 0) {
            ElMessage.success('回滚成功')
            getTableData()
          }
        })
      })
    })
  } else {
    ElMessageBox.confirm(`此操作将删除自动创建的文件和api, 是否继续?`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async() => {
      const res = await rollback({ id: Number(row.ID), deleteTable: flag })
      if (res.code === 0) {
        ElMessage.success('回滚成功')
        getTableData()
      }
    })
  }
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
