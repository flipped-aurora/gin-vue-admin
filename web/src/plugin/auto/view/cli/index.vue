<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="CLI名称">
          <el-input v-model="searchInfo.name" placeholder="搜索 CLI 名称" />
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
        <el-button type="primary" icon="plus" @click="openCreate">新增CLI</el-button>
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column label="名称" prop="name" min-width="140" />
        <el-table-column label="主命令" prop="command" min-width="140" />
        <el-table-column label="显示名称" prop="displayName" min-width="140" />
        <el-table-column label="版本" prop="version" width="100" />
        <el-table-column label="状态" prop="status" width="100" />
        <el-table-column label="API数" prop="apiCount" width="80" />
        <el-table-column label="操作" min-width="320">
          <template #default="scope">
            <el-button type="primary" link @click="openEdit(scope.row)">编辑</el-button>
            <el-button type="primary" link @click="openBindings(scope.row)">管理API</el-button>
            <el-button type="primary" link @click="openManifest(scope.row)">预览命令</el-button>
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

    <cli-form v-model="formVisible" :form-data="currentCli" @submit="submitForm" />
    <cli-api-binding v-model="bindingVisible" :cli="currentCli" @refresh="getTableData" />
    <cli-manifest-preview v-model="manifestVisible" :cli="currentCli" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { createCli, deleteCli, getCliList, updateCli } from '@/plugin/auto/api/cli'
import CliForm from './components/cliForm.vue'
import CliApiBinding from './components/cliApiBinding.vue'
import CliManifestPreview from './components/cliManifestPreview.vue'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const searchInfo = ref({})
const formVisible = ref(false)
const bindingVisible = ref(false)
const manifestVisible = ref(false)
const currentCli = ref({})

const getTableData = async () => {
  const res = await getCliList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}

const openCreate = () => {
  currentCli.value = {}
  formVisible.value = true
}

const openEdit = (row) => {
  currentCli.value = { ...row }
  formVisible.value = true
}

const submitForm = async (formData) => {
  const api = formData.id ? updateCli : createCli
  const res = await api(formData)
  if (res.code === 0) {
    ElMessage.success(formData.id ? '更新成功' : '创建成功')
    formVisible.value = false
    getTableData()
  }
}

const openBindings = (row) => {
  currentCli.value = { ...row }
  bindingVisible.value = true
}

const openManifest = (row) => {
  currentCli.value = { ...row }
  manifestVisible.value = true
}

const handleDelete = async (row) => {
  await ElMessageBox.confirm(`确定删除 CLI ${row.name} 吗？`, '提示', { type: 'warning' })
  const res = await deleteCli({ id: row.ID || row.id })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getTableData()
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
