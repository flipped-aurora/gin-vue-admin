<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="flex flex-wrap items-center gap-2">
        <el-form-item label="项目名称">
          <el-input v-model="searchInfo.name" placeholder="搜索项目名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增项目</el-button>
      </div>
      <el-table :data="tableData" style="width: 100%" row-key="ID">
        <el-table-column align="left" label="项目名称" prop="name" />
        <el-table-column align="left" label="项目描述" prop="desc" />
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="updateCategoryFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteCategoryFunc(scope.row)">删除</el-button>
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

    <el-dialog v-model="dialogFormVisible" :title="type==='create'?'新增项目':'编辑项目'" destroy-on-close>
      <el-form :model="formData" label-width="100px">
        <el-form-item label="项目名称">
          <el-input v-model="formData.name" autocomplete="off" placeholder="请输入项目名称" />
        </el-form-item>
        <el-form-item label="项目描述">
          <el-input v-model="formData.desc" type="textarea" autocomplete="off" placeholder="请输入项目描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  createCertCategory,
  deleteCertCategory,
  updateCertCategory,
  getCertCategoryList
} from '../api/certCategory'
import { ElMessage, ElMessageBox } from 'element-plus'

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const getTableData = async() => {
  const table = await getCertCategoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const formData = ref({
  name: '',
  desc: ''
})

const type = ref('')
const dialogFormVisible = ref(false)

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    desc: ''
  }
}

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const updateCategoryFunc = (row) => {
  type.value = 'update'
  formData.value = { ...row }
  dialogFormVisible.value = true
}

const deleteCategoryFunc = async(row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteCertCategory(row)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createCertCategory(formData.value)
      break
    case 'update':
      res = await updateCertCategory(formData.value)
      break
    default:
      res = await createCertCategory(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage.success('操作成功')
    closeDialog()
    getTableData()
  }
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}
</script>
