<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addMenu('0')" type="primary">新增根菜单</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="ID" min-width="40" prop="ID"></el-table-column>
      <el-table-column label="路径" min-width="100" prop="path"></el-table-column>
      <el-table-column label="名称" min-width="100" prop="name"></el-table-column>
      <el-table-column label="是否隐藏" min-width="80" prop="hidden">
        <template slot-scope="scope">
          <span>{{scope.row.hidden?"隐藏":"显示"}}</span>
        </template>
      </el-table-column>
      <el-table-column label="父节点Id" min-width="70" prop="parentId"></el-table-column>
      <el-table-column label="文件路径" min-width="250" prop="component"></el-table-column>
      <el-table-column label="展示名称" min-width="80" prop="authorityName">
        <template slot-scope="scope">
          <span>{{scope.row.meta.title}}</span>
        </template>
      </el-table-column>
      <el-table-column label="图标" min-width="180" prop="authorityName">
        <template slot-scope="scope">
          <span>{{scope.row.meta.icon}}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="deleteMenu(scope.row.ID)" size="small" type="text">删除菜单</el-button>
          <el-button @click="editMenu(scope.row.ID)" size="small" type="text">编辑菜单</el-button>
          <el-button @click="addMenu(scope.row.ID)" size="small" type="text">添加子菜单</el-button>
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

    <el-dialog :visible.sync="dialogFormVisible" title="新增角色">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="路径">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="名称">
          <el-input autocomplete="off" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="是否隐藏">
          <el-select placeholder="是否在列表隐藏" v-model="form.hidden">
            <el-option :value="false" label="否"></el-option>
            <el-option :value="true" label="是"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="父节点Id">
          <el-input autocomplete="off" disabled v-model="form.parentId"></el-input>
        </el-form-item>
        <el-form-item label="文件路径">
          <el-input autocomplete="off" v-model="form.component"></el-input>
        </el-form-item>
        <el-form-item label="展示名称">
          <el-input autocomplete="off" v-model="form.meta.title"></el-input>
        </el-form-item>
        <el-form-item label="图标">
          <el-input autocomplete="off" v-model="form.meta.icon"></el-input>
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
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成

import {
  updataBaseMenu,
  getMenuList,
  addBaseMenu,
  deleteBaseMenu,
  getBaseMenuById
} from '@/api/menu'
import infoList from '@/components/mixins/infoList'
export default {
  name: 'Menus',
  mixins: [infoList],
  data() {
    return {
      listApi: getMenuList,
      listKey: 'list',
      dialogFormVisible: false,
      form: {
        ID: 0,
        path: '',
        name: '',
        hidden: '',
        parentId: '',
        component: '',
        meta: {
          title: '',
          icon: ''
        }
      },
      isEdit: false
    }
  },
  methods: {
    // 删除菜单
    deleteMenu(ID) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteBaseMenu({ ID })
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
    // 初始化弹窗内表格方法
    initForm() {
      this.form = {
        path: '',
        name: '',
        hidden: '',
        parentId: '',
        component: '',
        meta: {
          title: '',
          icon: ''
        }
      }
    },
    // 关闭弹窗
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    // 添加menu
    async enterDialog() {
      let res
      if (this.isEdit) {
        res = await updataBaseMenu(this.form)
      } else {
        res = await addBaseMenu(this.form)
      }
      if (res.success) {
        this.$message({
          type: 'success',
          message: '添加成功!'
        })
        this.getTableData()
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
    // 添加菜单方法，id为 0则为添加根菜单
    addMenu(id) {
      this.form.parentId = String(id)
      this.isEdit = false
      this.dialogFormVisible = true
    },
    // 修改菜单方法
    async editMenu(id) {
      const res = await getBaseMenuById({ id })
      this.form = res.data.menu
      this.dialogFormVisible = true
      this.isEdit = true
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