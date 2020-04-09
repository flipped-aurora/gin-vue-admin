<template>
  <div class="authority">
    <div class="button-box clearflex">
      <el-button @click="addAuthority('0')" type="primary">新增角色</el-button>
    </div>
    <el-table
      :data="tableData"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
      border
      row-key="authorityId"
      stripe
      style="width: 100%"
    >
      <el-table-column label="角色id" min-width="180" prop="authorityId"></el-table-column>
      <el-table-column label="角色名称" min-width="180" prop="authorityName"></el-table-column>
      <el-table-column fixed="right" label="操作" min-width="300">
        <template slot-scope="scope">
          <el-button @click="opdendrawer(scope.row)" size="small" type="text">设置权限</el-button>
          <el-button @click="deleteAuth(scope.row)" size="small" type="text">删除角色</el-button>
          <el-button @click="addAuthority(scope.row.authorityId)" size="small" type="text">新增子角色</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 新增角色弹窗 -->
    <el-dialog :visible.sync="dialogFormVisible" :title="dialogTitle">
      <el-form :model="form" :rules="rules" ref="authorityForm">
        <el-form-item label="父级角色ID" prop="parentId">
          <el-input autocomplete="off" disabled v-model="form.parentId"></el-input>
        </el-form-item>
        <el-form-item label="角色ID" prop="authorityId">
          <el-input autocomplete="off" v-model="form.authorityId"></el-input>
        </el-form-item>
        <el-form-item label="角色姓名" prop="authorityName">
          <el-input autocomplete="off" v-model="form.authorityName"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-drawer :visible.sync="drawer" :with-header="false" size="40%" title="角色配置" v-if="drawer">
      <el-tabs class="role-box" type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus :row="activeRow" />
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <apis :row="activeRow" />
        </el-tab-pane>
        <el-tab-pane label="资源权限">
          <Datas :authority="tableData" :row="activeRow" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成

import {
  getAuthorityList,
  deleteAuthority,
  createAuthority
} from '@/api/authority'

import Menus from '@/view/superAdmin/authority/components/menus'
import Apis from '@/view/superAdmin/authority/components/apis'
import Datas from '@/view/superAdmin/authority/components/datas'

import infoList from '@/components/mixins/infoList'
export default {
  name: 'Authority',
  mixins: [infoList],
  data() {
    return {
      listApi: getAuthorityList,
      drawer: false,
      activeRow: {},
      activeUserId: 0,
      dialogTitle:"新增角色",
      dialogFormVisible: false,
      apiDialogFlag: false,
      form: {
        authorityId: '',
        authorityName: '',
        parentId: '0'
      },
      rules: {
        authorityId: [
          { required: true, message: '请输入角色ID', trigger: 'blur' }
        ],
        authorityName: [
          { required: true, message: '请输入角色名', trigger: 'blur' }
        ],
        parentId: [
          { required: true, message: '请选择请求方式', trigger: 'blur' }
        ]
      }
    }
  },
  components: {
    Menus,
    Apis,
    Datas
  },
  methods: {
    opdendrawer(row) {
      this.drawer = true
      this.activeRow = row
    },
    // 删除角色
    deleteAuth(row) {
      this.$confirm('此操作将永久删除该角色, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteAuthority({ authorityId: row.authorityId })
          if (res.code == 0) {
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
      this.$refs.authorityForm.resetFields()
    },
    // 关闭窗口
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
      this.apiDialogFlag = false
    },
    // 确定弹窗

    async enterDialog() {
      if (this.form.authorityId == '0') {
        this.$message({
          type: 'error',
          message: '角色id不能为0'
        })
        return false
      }
      this.$refs.authorityForm.validate(async valid => {
        if (valid) {
          const res = await createAuthority(this.form)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: '添加成功!'
            })
            this.getTableData()
            this.closeDialog()
          }
          this.initForm()
          this.dialogFormVisible = false
        }
      })
    },
    // 增加角色
    addAuthority(parentId) {
      this.form.parentId = parentId
      this.dialogFormVisible = true
    }
  },
  created() {
    this.pageSize = 999
  }
}
</script>
<style lang="scss">
.authority {
  .button-box {
    padding: 10px 20px;
    .el-button {
      float: right;
    }
  }
}
.role-box {
  .el-tabs__content {
    height: calc(100vh - 150px);
    overflow: auto;
  }
}
</style>