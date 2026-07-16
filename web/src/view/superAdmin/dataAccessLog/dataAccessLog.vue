<template>
  <div>
    <warning-bar
      title="注：no_identity=无身份访问受控表（旁路盲区，需补 ctx 或 WithSystem）；blocked_write=数据范围过滤后写操作影响 0 行（疑似越权尝试，启发式信号）"
    />
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="事件类型">
          <el-select
            v-model="searchInfo.eventType"
            clearable
            placeholder="全部"
            class="!w-44"
          >
            <el-option label="无身份访问" value="no_identity" />
            <el-option label="疑似越权写" value="blocked_write" />
          </el-select>
        </el-form-item>
        <el-form-item label="业务表">
          <el-input v-model="searchInfo.targetTable" placeholder="表名(模糊)" />
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
          type="danger"
          icon="delete"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >批量删除</el-button
        >
      </div>
      <el-table
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="时间" min-width="160">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="事件类型" min-width="120">
          <template #default="scope">
            <el-tag :type="scope.row.eventType === 'no_identity' ? 'warning' : 'danger'">
              {{ scope.row.eventType === 'no_identity' ? '无身份访问' : '疑似越权写' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="业务表" min-width="150" prop="targetTable" />
        <el-table-column align="left" label="操作" min-width="90" prop="operation" />
        <el-table-column align="left" label="用户ID" min-width="80">
          <template #default="scope">
            {{ scope.row.userId || '-' }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="角色ID" min-width="80">
          <template #default="scope">
            {{ scope.row.authorityId || '-' }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="来源请求" min-width="220" show-overflow-tooltip>
          <template #default="scope">
            {{ scope.row.method ? `${scope.row.method} ${scope.row.path}` : '-' }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="详情" min-width="240" prop="detail" show-overflow-tooltip />
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
    getDataAccessLogList,
    deleteDataAccessLogByIds
  } from '@/api/dataAccessLog'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { formatDate } from '@/utils/format'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  defineOptions({
    name: 'DataAccessLog'
  })

  const tableData = ref([])
  const page = ref(1)
  const pageSize = ref(10)
  const total = ref(0)
  const searchInfo = ref({ eventType: '', targetTable: '' })
  const multipleSelection = ref([])

  const getTableData = async () => {
    const res = await getDataAccessLogList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (res.code === 0) {
      tableData.value = res.data.list || []
      total.value = res.data.total
    }
  }
  getTableData()

  const onSubmit = () => {
    page.value = 1
    getTableData()
  }
  const onReset = () => {
    searchInfo.value = { eventType: '', targetTable: '' }
    page.value = 1
    getTableData()
  }
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }
  const handleSizeChange = (val) => {
    pageSize.value = val
    page.value = 1
    getTableData()
  }
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  const onDelete = () => {
    ElMessageBox.confirm('确定要删除选中的审计记录吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const ids = multipleSelection.value.map((item) => item.ID)
      const res = await deleteDataAccessLogByIds({ ids })
      if (res.code === 0) {
        ElMessage.success('删除成功')
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }
</script>
