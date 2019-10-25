<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addAuthority" type="primary">新增角色</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="180" prop="ID"></el-table-column>
      <el-table-column label="角色id" min-width="180" prop="authorityId"></el-table-column>
      <el-table-column label="角色名称" min-width="180" prop="authorityName"></el-table-column>
      <el-table-column fixed="right" label="操作" width="500">
        <template slot-scope="scope">
          <el-button @click="addAuthMenu(scope.row)" size="small" type="text">变更角色菜单</el-button>
          <el-button @click="addAuthApi(scope.row)" size="small" type="text">变更角色Api</el-button>
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
        :data="menuTreeData"
        :default-checked-keys="menuTreeIds"
        :props="menuDefaultProps"
        default-expand-all
        highlight-current
        node-key="ID"
        ref="menuTree"
        show-checkbox
        v-if="menuDialogFlag"
      ></el-tree>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="relation" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <!-- 关联api弹窗 -->
    <el-dialog :visible.sync="apiDialogFlag" title="关联api">
      <el-tree
        :data="apiTreeData"
        :default-checked-keys="apiTreeIds"
        :props="apiDefaultProps"
        default-expand-all
        highlight-current
        node-key="ID"
        ref="apiTree"
        show-checkbox
        v-if="apiDialogFlag"
      ></el-tree>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="authApiEnter" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成

import {
  getAuthorityList,
  deleteAuthority,
  createAuthority
} from '@/api/authority'
import { getBaseMenuTree, addMenuAuthority, getMenuAuthority } from '@/api/menu'
import { getAllApis, getAuthAndApi, setAuthAndApi } from '@/api/api'
import infoList from '@/components/mixins/infoList'
export default {
  name: 'Authority',
  mixins: [infoList],
  data() {
    return {
      listApi: getAuthorityList,
      listKey: 'list',
      activeUserId: 0,
      menuTreeData: [],
      menuTreeIds: [],
      menuDefaultProps: {
        children: 'children',
        label: 'nickName'
      },

      apiTreeData: [],
      apiTreeIds: [],
      apiDefaultProps: {
        children: 'children',
        label: 'description'
      },
      dialogFormVisible: false,
      menuDialogFlag: false,
      apiDialogFlag: false,
      form: {
        authorityId: '',
        authorityName: ''
      }
    }
  },
  methods: {
    // 删除角色
    deleteAuth(row) {
      this.$confirm('此操作将永久删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteAuthority({ authorityId: row.authorityId })
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
      this.apiDialogFlag = false
    },
    // 确定弹窗

    async enterDialog() {
      const res = await createAuthority(this.form)
      if (res.success) {
        this.$message({
          type: 'success',
          message: '添加成功!'
        })
        this.getTableData()
        this.closeDialog()
      }
      this.initForm()
      this.dialogFormVisible = false
    },
    // 增加角色
    addAuthority() {
      this.dialogFormVisible = true
    },

    // 关联用户列表关系
    async addAuthMenu(row) {
      const res1 = await getMenuAuthority({ authorityId: row.authorityId })
      const menus = res1.data.menus
      const arr = []
      menus.map(item => {
        // 防止直接选中父级造成全选
        if (!menus.some(same => same.parentId === item.menuId)) {
          arr.push(Number(item.menuId))
        }
      })
      this.menuTreeIds = arr
      this.activeUserId = row.authorityId
      this.menuDialogFlag = true
    },
    // 关联树 确认方法
    async relation() {
      const checkArr = this.$refs.menuTree.getCheckedNodes(false, true)
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
    },
    // 创建api树方法
    buildApiTree(apis) {
      const apiObj = new Object()
      apis &&
        apis.map(item => {
          if (apiObj.hasOwnProperty(item.group)) {
            apiObj[item.group].push(item)
          } else {
            Object.assign(apiObj, { [item.group]: [item] })
          }
        })
      const apiTree = []
      for (const key in apiObj) {
        const treeNode = {
          ID: key,
          description: key + '组',
          children: apiObj[key]
        }
        apiTree.push(treeNode)
      }
      return apiTree
    },
    // 关联用户api关系
    async addAuthApi(row) {
      const res = await getAuthAndApi({ authorityId: row.authorityId })
      this.activeUserId = row.authorityId
      this.apiTreeIds = res.data.apis || []
      this.apiDialogFlag = true
    },
    async authApiEnter() {
      const checkArr = this.$refs.apiTree.getCheckedKeys(true)
      const res = await setAuthAndApi({
        authorityId: this.activeUserId,
        apiIds: checkArr
      })
      if (res.success) {
        this.$message({
          type: 'success',
          message: '添加成功!'
        })
      }
      this.closeDialog()
    }
  },
  async created() {
    // 获取所有菜单树
    const res = await getBaseMenuTree()
    this.menuTreeData = res.data.menus
    // 获取api并整理成树结构
    const res2 = await getAllApis()
    const apis = res2.data.apis
    this.apiTreeData = this.buildApiTree(apis)
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