<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addAuthority" type="primary">新增角色</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="角色id" min-width="180" prop="authorityId"></el-table-column>
      <el-table-column label="角色名称" min-width="180" prop="authorityName"></el-table-column>
      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="deleteAuth(scope.row)" size="small" type="text">删除角色</el-button>
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
      hide-on-single-page
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :visible.sync="dialogFormVisible" title="新增角色">
      <el-form :model="form">
        <el-form-item label="角色ID">
          <el-input autocomplete="off" v-model="form.authorityId"></el-input>
        </el-form-item>
        <el-form-item label="角色姓名">
          <el-input autocomplete="off" v-model="form.authorityName"></el-input>
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
import { getAuthorityList, deleteAuthority, createAuthority } from '@/api/authority'
export default {
  name: 'Authority',
  data() {
    return {
      page: 1,
      total: 10,
      pageSize: 10,
      tableData: [],
      dialogFormVisible: false,
      form:{
          authorityId:"",
          authorityName:""
      }
    }
  },
  methods: {
    handleSizeChange(val) {
      this.pageSize = val
      this.getAuthList()
    },
    handleCurrentChange(val) {
      this.page = val
      this.getAuthList()
    },
    deleteAuth(row) {
      this.$confirm('此操作将永久删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          try {
            const res = await deleteAuthority({ authorityId: row.authorityId })
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            this.getAuthList()
          } catch (err) {
            this.$message({
              type: 'error',
              message: '删除失败!' + err
            })
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    initForm(){
        for(const key in this.form){
            this.form[key] = ''
        }
    },
    closeDialog(){
        this.initForm()
        this.dialogFormVisible = false
    },
    async enterDialog(){
        const res = await createAuthority(this.form)
        if(res.success){
            this.$message({
              type: 'success',
              message: '添加成功!'
            })
            this.getAuthList()
            this.closeDialog()
        }else{
            this.$message({
              type: 'error',
              message: '添加失败!'
            })
            this.closeDialog()
        }
        this.initForm()
        this.dialogFormVisible = false
    },
    addAuthority() {
        this.dialogFormVisible = true
    },
    async getAuthList(page = this.page, pageSize = this.pageSize) {
      try {
        const table = await getAuthorityList({ page, pageSize })
        this.tableData = table.data.authList
      } catch (err) {
        console.log(err)
      }
    }
  },
  created() {
    this.getAuthList()
  }
}
</script>
<style lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
</style>