<template>
  <div class="upload">
    <div class="gva-table-box">

      <div class="gva-btn-list">
        <el-upload
          class="excel-btn"
          :action="`${path}/excel/importExcel`"
          :headers="{'x-token':token}"
          :on-success="loadExcel"
          :show-file-list="false"
        >
          <el-button size="mini" type="primary" icon="el-icon-upload2">导入</el-button>
        </el-upload>
        <el-button class="excel-btn" size="mini" type="primary" icon="el-icon-download" @click="handleExcelExport('ExcelExport.xlsx')">导出</el-button>
        <el-button class="excel-btn" size="mini" type="success" icon="el-icon-download" @click="downloadExcelTemplate()">下载模板</el-button>
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
const path = import.meta.env.VITE_BASE_API
import { mapGetters } from 'vuex'
import infoList from '@/mixins/infoList'
import { exportExcel, loadExcelData, downloadTemplate } from '@/api/excel'
import { getMenuList } from '@/api/menu'
export default {
  name: 'Excel',
  mixins: [infoList],
  data() {
    return {
      listApi: getMenuList,
      path: path
    }
  },
  computed: {
    ...mapGetters('user', ['userInfo', 'token'])
  },
  created() {
    this.pageSize = 999
    this.getTableData()
  },
  methods: {
    handleExcelExport(fileName) {
      if (!fileName || typeof fileName !== 'string') {
        fileName = 'ExcelExport.xlsx'
      }
      exportExcel(this.tableData, fileName)
    },
    loadExcel() {
      this.listApi = loadExcelData
      this.getTableData()
    },
    downloadExcelTemplate() {
      downloadTemplate('ExcelTemplate.xlsx')
    }
  }
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
