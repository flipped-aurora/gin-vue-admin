<template>
  <div class="upload">
    <div class="gva-table-box">

      <div class="gva-btn-list">
        <el-upload
          class="excel-btn"
          :action="`${path}/excel/importExcel`"
          :headers="{'x-token':userStore.token}"
          :on-success="loadExcel"
          :show-file-list="false"
        >
          <el-button size="mini" type="primary" icon="upload">导入</el-button>
        </el-upload>
        <el-button class="excel-btn" size="mini" type="primary" icon="download" @click="handleExcelExport('ExcelExport.xlsx')">导出</el-button>
        <el-button class="excel-btn" size="mini" type="success" icon="download" @click="downloadExcelTemplate()">下载模板</el-button>
      </div>
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="100" prop="ID" />
        <el-table-column align="left" show-overflow-tooltip label="路由Name" min-width="160" prop="name" />
        <el-table-column align="left" show-overflow-tooltip label="路由Path" min-width="160" prop="path" />
        <el-table-column align="left" label="是否隐藏" min-width="100" prop="hidden">

          <template #default="scope">
            <span>{{ scope.row.hidden?"隐藏":"显示" }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="父节点" min-width="90" prop="parentId" />
        <el-table-column align="left" label="排序" min-width="70" prop="sort" />
        <el-table-column align="left" label="文件路径" min-width="360" prop="component" />
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Excel',
}
</script>

<script setup>
import { useUserStore } from '@/pinia/modules/user'
import { exportExcel, loadExcelData, downloadTemplate } from '@/api/excel'
import { getMenuList } from '@/api/menu'
import { ref } from 'vue'
const path = ref(import.meta.env.VITE_BASE_API)

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])

// 查询
const getTableData = async(f = () => {}) => {
  const table = await f({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData(getMenuList)

const userStore = useUserStore()

const handleExcelExport = (fileName) => {
  if (!fileName || typeof fileName !== 'string') {
    fileName = 'ExcelExport.xlsx'
  }
  exportExcel(tableData.value, fileName)
}
const loadExcel = () => {
  getTableData(loadExcelData)
}
const downloadExcelTemplate = () => {
  downloadTemplate('ExcelTemplate.xlsx')
}

</script>

<style lang="scss" scoped>
.btn-list{
  display: flex;
  margin-bottom: 12px;
  justify-content: flex-end;

}
.excel-btn+.excel-btn{
  margin-left: 10px;
}
</style>
