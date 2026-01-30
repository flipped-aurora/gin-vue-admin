<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.username" placeholder="搜索用户名" />
        </el-form-item>
        <el-form-item label="状态">
             <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                 <el-option label="成功" :value="true" />
                 <el-option label="失败" :value="false" />
             </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          icon="delete"
          style="margin-left: 10px;"
          :disabled="!multipleSelection.length"
          @click="onDelete"
        >删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="用户名" prop="username" width="150" />
        <el-table-column align="left" label="登录IP" prop="ip" width="150" />
        <el-table-column align="left" label="状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status ? 'success' : 'danger'">
              {{ scope.row.status ? '成功' : '失败' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="详情" show-overflow-tooltip>
             <template #default="scope">
                 {{ scope.row.status ? '登录成功' : scope.row.errorMessage }}
             </template>
        </el-table-column>
        <el-table-column align="left" label="浏览器/设备" prop="agent" show-overflow-tooltip />
        <el-table-column align="left" label="登录时间" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="120">
          <template #default="scope">
            <el-popover v-model:visible="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button size="small" type="primary" link @click="scope.row.visible = false">取消</el-button>
                <el-button size="small" type="primary" @click="deleteRow(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button icon="delete" type="primary" link @click="scope.row.visible = true">删除</el-button>
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
  getLoginLogList,
  deleteLoginLog,
  deleteLoginLogByIds
} from '@/api/sysLoginLog'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const multipleSelection = ref([])

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const getTableData = async () => {
  const table = await getLoginLogList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const deleteRow = async (row) => {
  row.visible = false
  const res = await deleteLoginLog(row)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const onDelete = async() => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(async() => {
        const ids = multipleSelection.value.map(item => item.ID)
        const res = await deleteLoginLogByIds({ ids })
        if (res.code === 0) {
            ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === ids.length && page.value > 1) {
                page.value--
            }
            getTableData()
        }
    })
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 首次加载
getTableData()
</script>

<style scoped>
</style>
