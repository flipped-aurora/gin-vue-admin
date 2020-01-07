<template>
  <div>
    <div class="button-box clearflex">
    </div>
        <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
  <el-form-item label="路径">
    <el-input v-model="searchInfo.path" placeholder="路径"></el-input>
  </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="onSubmit">查询</el-button>
  </el-form-item>
   <el-form-item >
      <el-button @click="openDialog('addApi')" type="primary">新增api</el-button>
  </el-form-item>
</el-form>
  </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="ID"></el-table-column>
      <el-table-column label="api路径" min-width="150" prop="path"></el-table-column>
      <el-table-column label="api分组" min-width="150" prop="group"></el-table-column>
      <el-table-column label="api简介" min-width="150" prop="description"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteApi(scope.row)" size="small" type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :visible.sync="dialogFormVisible" title="新增Api">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="路径">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="api分组">
          <el-input autocomplete="off" v-model="form.group"></el-input>
        </el-form-item>
        <el-form-item label="api简介">
          <el-input autocomplete="off" v-model="form.description"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getApiById,
  getApiList,
  createApi,
  updataApi,
  deleteApi
} from '@/api/api'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getApiList,
      listKey: 'list',
      dialogFormVisible: false,
      form: {
        path: '',
        group: '',
        description: ''
      },
      type: ''
    }
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit(){
      this.page = 1 
      this.pageSize = 10
      this.getTableData()
    },
    initForm() {
      this.form = {
        path: '',
        group: '',
        description: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      this.type = type
      this.dialogFormVisible = true
    },
    async editApi(row) {
      const res = await getApiById({ id: row.ID })
      this.form = res.data.api
      this.openDialog('edit')
    },
    async deleteApi(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteApi(row)
          if (res.success) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    async enterDialog() {
      switch (this.type) {
        case 'addApi':
          {
            const res = await createApi(this.form)
            if (res.success) {
              this.$message({
                type: 'success',
                message: '添加成功',
                showClose: true
              })
            }
            this.getTableData()
            this.closeDialog()
          }

          break
        case 'edit':
          {
            const res = await updataApi(this.form)
            if (res.success) {
              this.$message({
                type: 'success',
                message: '添加成功',
                showClose: true
              })
            }
            this.getTableData()
            this.closeDialog()
          }
          break
        default:
          {
            this.$message({
              type: 'error',
              message: '未知操作',
              showClose: true
            })
          }
          break
      }
    }
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
</style>