<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="搜索 MCP 名称" />
        </el-form-item>
        <el-form-item label="显示名称">
          <el-input v-model="searchInfo.displayName" placeholder="搜索显示名称" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" clearable placeholder="请选择">
            <el-option label="启用" value="enabled" />
            <el-option label="禁用" value="disabled" />
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
        <el-button type="primary" icon="plus" @click="openCreate">新增MCP</el-button>
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column label="名称" prop="name" min-width="140" />
        <el-table-column label="显示名称" prop="displayName" min-width="140" />
        <el-table-column label="版本" prop="version" width="100" />
        <el-table-column label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-switch
              :model-value="row.status === 'enabled'"
              active-text="启用"
              inactive-text="禁用"
              inline-prompt
              @change="(val) => toggleStatus(row, val)"
            />
          </template>
        </el-table-column>
        <el-table-column label="API数" prop="apiCount" width="80" />
        <el-table-column label="操作" min-width="300">
          <template #default="scope">
            <el-button type="primary" link @click="openEdit(scope.row)">编辑</el-button>
            <el-button type="primary" link @click="openBindings(scope.row)">管理API</el-button>
            <el-button type="primary" link @click="openScenarios(scope.row)">管理场景</el-button>
            <el-button type="danger" link @click="handleDelete(scope.row)">删除</el-button>
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

    <mcp-form v-model="formVisible" :form-data="currentMcp" @submit="submitForm" />
    <mcp-api-binding v-model="bindingVisible" :mcp="currentMcp" @refresh="getTableData" />
    <mcp-scenario-flow v-model="scenarioVisible" :mcp="currentMcp" @refresh="getTableData" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { createMcp, deleteMcp, getMcpList, updateMcp } from '@/plugin/ai/api/mcpApi'
import McpForm from './components/mcpForm.vue'
import McpApiBinding from './components/mcpApiBinding.vue'
import McpScenarioFlow from './components/mcpScenarioFlow.vue'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})
const formVisible = ref(false)
const bindingVisible = ref(false)
const scenarioVisible = ref(false)
const currentMcp = ref({})

const getTableData = async () => {
  const res = await getMcpList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}

const openCreate = () => {
  currentMcp.value = {}
  formVisible.value = true
}

const openEdit = (row) => {
  currentMcp.value = { ...row }
  formVisible.value = true
}

const submitForm = async (formData) => {
  const api = formData.id ? updateMcp : createMcp
  const res = await api(formData)
  if (res.code === 0) {
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    formVisible.value = false
    getTableData()
  }
}

const openBindings = (row) => {
  currentMcp.value = { ...row }
  bindingVisible.value = true
}

const openScenarios = (row) => {
  currentMcp.value = { ...row }
  scenarioVisible.value = true
}

const handleDelete = async (row) => {
  await ElMessageBox.confirm(`确定删除 MCP ${row.name} 吗？`, '提示', { type: 'warning' })
  const res = await deleteMcp({ id: row.ID || row.id })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
  }
}

const toggleStatus = async (row, val) => {
  const res = await updateMcp({
    id: row.ID || row.id,
    name: row.name,
    displayName: row.displayName,
    description: row.description,
    status: val ? 'enabled' : 'disabled',
    version: row.version,
    scenariosJson: row.scenariosJson
  })
  if (res.code === 0) {
    row.status = val ? 'enabled' : 'disabled'
    ElMessage.success(val ? '已启用' : '已禁用')
  } else {
    ElMessage.error(res.msg || '切换失败')
  }
}

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

getTableData()
</script>
