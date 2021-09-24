<template>
  <div class="authority">
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="addAuthority('0')">新增角色</el-button>
      </div>
      <el-table
        :data="tableData"
        :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
        row-key="authorityId"
        style="width: 100%"
      >
        <el-table-column label="角色ID" min-width="180" prop="authorityId" />
        <el-table-column align="left" label="角色名称" min-width="180" prop="authorityName" />
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button
              icon="el-icon-setting"
              size="mini"
              type="text"
              @click="opdendrawer(scope.row)"
            >设置权限</el-button>
            <el-button
              icon="el-icon-plus"
              size="mini"
              type="text"
              @click="addAuthority(scope.row.authorityId)"
            >新增子角色</el-button>
            <el-button
              icon="el-icon-copy-document"
              size="mini"
              type="text"
              @click="copyAuthority(scope.row)"
            >拷贝</el-button>
            <el-button
              icon="el-icon-edit"
              size="mini"
              type="text"
              @click="editAuthority(scope.row)"
            >编辑</el-button>
            <el-button
              icon="el-icon-delete"
              size="mini"
              type="text"
              @click="deleteAuth(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增角色弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
      <el-form ref="authorityForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="父级角色" prop="parentId">
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="dialogType=='add'"
            :options="AuthorityOption"
            :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="角色ID" prop="authorityId">
          <el-input v-model="form.authorityId" :disabled="dialogType=='edit'" autocomplete="off" />
        </el-form-item>
        <el-form-item label="角色姓名" prop="authorityName">
          <el-input v-model="form.authorityName" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-drawer v-if="drawer" v-model="drawer" :with-header="false" size="40%" title="角色配置">
      <el-tabs :before-leave="autoEnter" class="role-box" type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <Apis ref="apis" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="资源权限">
          <Datas ref="datas" :authority="tableData" :row="activeRow" @changeRow="changeRow" />
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
  createAuthority,
  updateAuthority,
  copyAuthority
} from '@/api/authority'

import Menus from '@/view/superAdmin/authority/components/menus.vue'
import Apis from '@/view/superAdmin/authority/components/apis.vue'
import Datas from '@/view/superAdmin/authority/components/datas.vue'
import warningBar from '@/components/warningBar/warningBar.vue'

import infoList from '@/mixins/infoList'
export default {
  name: 'Authority',
  components: {
    Menus,
    Apis,
    Datas,
    warningBar
  },
  mixins: [infoList],
  data() {
    var mustUint = (rule, value, callback) => {
      if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
        return callback(new Error('请输入正整数'))
      }
      return callback()
    }

    return {
      AuthorityOption: [
        {
          authorityId: '0',
          authorityName: '根角色'
        }
      ],
      listApi: getAuthorityList,
      drawer: false,
      dialogType: 'add',
      activeRow: {},
      activeUserId: 0,
      dialogTitle: '新增角色',
      dialogFormVisible: false,
      apiDialogFlag: false,
      copyForm: {},
      form: {
        authorityId: '',
        authorityName: '',
        parentId: '0'
      },
      rules: {
        authorityId: [
          { required: true, message: '请输入角色ID', trigger: 'blur' },
          { validator: mustUint, trigger: 'blur' }
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
  async created() {
    this.pageSize = 999
    await this.getTableData()
  },
  methods: {
    changeRow(key, value) {
      this.activeRow[key] = value
    },
    autoEnter(activeName, oldActiveName) {
      const paneArr = ['menus', 'apis', 'datas']
      if (oldActiveName) {
        if (this.$refs[paneArr[oldActiveName]].needConfirm) {
          this.$refs[paneArr[oldActiveName]].enterAndNext()
          this.$refs[paneArr[oldActiveName]].needConfirm = false
        }
      }
    },
    // 拷贝角色
    copyAuthority(row) {
      this.setOptions()
      this.dialogTitle = '拷贝角色'
      this.dialogType = 'copy'
      for (const k in this.form) {
        this.form[k] = row[k]
      }
      this.copyForm = row
      this.dialogFormVisible = true
    },
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
        .then(async() => {
          const res = await deleteAuthority({ authorityId: row.authorityId })
          if (res.code === 0) {
            this.$message({
              type: 'success',
              message: '删除成功!'
            })
            if (this.tableData.length === 1 && this.page > 1) {
              this.page--
            }
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
      if (this.$refs.authorityForm) {
        this.$refs.authorityForm.resetFields()
      }
      this.form = {
        authorityId: '',
        authorityName: '',
        parentId: '0'
      }
    },
    // 关闭窗口
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
      this.apiDialogFlag = false
    },
    // 确定弹窗

    async enterDialog() {
      if (this.form.authorityId === '0') {
        this.$message({
          type: 'error',
          message: '角色id不能为0'
        })
        return false
      }
      this.$refs.authorityForm.validate(async valid => {
        if (valid) {
          switch (this.dialogType) {
            case 'add':
              {
                const res = await createAuthority(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功!'
                  })
                  this.getTableData()
                  this.closeDialog()
                }
              }
              break
            case 'edit':
              {
                const res = await updateAuthority(this.form)
                if (res.code === 0) {
                  this.$message({
                    type: 'success',
                    message: '添加成功!'
                  })
                  this.getTableData()
                  this.closeDialog()
                }
              }
              break
            case 'copy': {
              const data = {
                authority: {
                  authorityId: 'string',
                  authorityName: 'string',
                  datauthorityId: [],
                  parentId: 'string'
                },
                oldAuthorityId: 0
              }
              data.authority.authorityId = this.form.authorityId
              data.authority.authorityName = this.form.authorityName
              data.authority.parentId = this.form.parentId
              data.authority.dataAuthorityId = this.copyForm.dataAuthorityId
              data.oldAuthorityId = this.copyForm.authorityId
              const res = await copyAuthority(data)
              if (res.code === 0) {
                this.$message({
                  type: 'success',
                  message: '复制成功！'
                })
                this.getTableData()
              }
            }
          }

          this.initForm()
          this.dialogFormVisible = false
        }
      })
    },
    setOptions() {
      this.AuthorityOption = [
        {
          authorityId: '0',
          authorityName: '根角色'
        }
      ]
      this.setAuthorityOptions(this.tableData, this.AuthorityOption, false)
    },
    setAuthorityOptions(AuthorityData, optionsData, disabled) {
      this.form.authorityId = String(this.form.authorityId)
      AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId === this.form.authorityId,
              children: []
            }
            this.setAuthorityOptions(
              item.children,
              option.children,
              disabled || item.authorityId === this.form.authorityId
            )
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId === this.form.authorityId
            }
            optionsData.push(option)
          }
        })
    },
    // 增加角色
    addAuthority(parentId) {
      this.initForm()
      this.dialogTitle = '新增角色'
      this.dialogType = 'add'
      this.form.parentId = parentId
      this.setOptions()
      this.dialogFormVisible = true
    },
    // 编辑角色
    editAuthority(row) {
      this.setOptions()
      this.dialogTitle = '编辑角色'
      this.dialogType = 'edit'
      for (const key in this.form) {
        this.form[key] = row[key]
      }
      this.setOptions()
      this.dialogFormVisible = true
    }
  }
}
</script>

<style lang="scss">
.authority {
  .el-input-number {
    margin-left: 15px;
    span {
      display: none;
    }
  }
}
.role-box {
  .el-tabs__content {
    height: calc(100vh - 72px);
    overflow: auto;
  }
}
</style>
