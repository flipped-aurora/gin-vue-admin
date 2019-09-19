<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addAuthority" type="primary">新增角色</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="角色id" min-width="180" prop="authorityId"></el-table-column>
      <el-table-column label="角色名称" min-width="180" prop="authorityName"></el-table-column>
      <el-table-column fixed="right" label="操作" width="500">
        <template slot-scope="scope">
          <el-button @click="addAuthMenu(scope.row)" size="small" type="text">增加角色菜单</el-button>
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
    <!-- 新增角色弹窗 -->
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

    <!-- 关联menu弹窗 -->
    <el-dialog :visible.sync="menuDialogFlag" title="关联菜单">
      <el-tree
        :data="treeData"
        :props="defaultProps"
        default-expand-all
        highlight-current
        node-key="ID"
        :default-checked-keys="treeIds"
        ref="tree"
        show-checkbox
      ></el-tree>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="relation" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getAuthorityList,
  deleteAuthority,
  createAuthority
} from '@/api/authority'
import { getBaseMenuTree, addMenuAuthority, getMenuAuthority } from '@/api/menu'
export default {
  name: 'Authority',
  data() {
    return {
      activeUserId: 0,
      page: 1,
      total: 10,
      pageSize: 10,
      tableData: [],
      treeData: [],
      treeIds:[],
      defaultProps: {
        children: 'children',
        label: 'nickName'
      },
      dialogFormVisible: false,
      menuDialogFlag: false,
      form: {
        authorityId: '',
        authorityName: ''
      }
    }
  },
  methods: {
    // 条数
    handleSizeChange(val) {
      this.pageSize = val
      this.getAuthList()
    },
    // 页码
    handleCurrentChange(val) {
      this.page = val
      this.getAuthList()
    },
    // 删除角色
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
    // 初始化表单
    initForm() {
      for (const key in this.form) {
        this.form[key] = ''
      }
    },
    // 关闭窗口
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
      this.menuDialogFlag = false
    },
    // 确定弹窗
    async enterDialog() {
      const res = await createAuthority(this.form)
      if (res.success) {
        this.$message({
          type: 'success',
          message: '添加成功!'
        })
        this.getAuthList()
        this.closeDialog()
      } else {
        this.$message({
          type: 'error',
          message: '添加失败!'
        })
        this.closeDialog()
      }
      this.initForm()
      this.dialogFormVisible = false
    },
    // 增加角色
    addAuthority() {
      this.dialogFormVisible = true
    },
    // 获取用户列表
    async getAuthList(page = this.page, pageSize = this.pageSize) {
      try {
        const table = await getAuthorityList({ page, pageSize })
        this.tableData = table.data.authList
      } catch (err) {
        console.log(err)
      }
    },
    async addAuthMenu(row) {
      const res1 = await getMenuAuthority({authorityId:row.authorityId})
      const menus = res1.data.menus
      const arr = []
      menus.map(item=>{arr.push(Number(item.menuId))})
      this.treeIds = arr
      const res2 = await getBaseMenuTree()
      this.treeData = res2.data.menus
      console.log(this.treeData)
      this.activeUserId = row.authorityId
      this.menuDialogFlag = true
    },
    // 关联树
    async relation() {
      const checkArr = this.$refs.tree
        .getCheckedNodes()
        .concat(this.$refs.tree.getHalfCheckedNodes())
      const res = await addMenuAuthority({
        menus: checkArr,
        authorityId: this.activeUserId
      })
      if (res.success) {
        this.$message({
          type: 'success',
          message: '添加成功!'
        })
      }
      this.closeDialog()
    }
    // 获取基础menu树
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