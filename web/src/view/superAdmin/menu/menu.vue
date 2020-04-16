<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addMenu('0')" type="primary">新增根菜单</el-button>
    </div>

    <!-- 由于此处菜单跟左侧列表一一对应所以不需要分页 pageSize默认999 -->
    <el-table :data="tableData" border row-key="ID" stripe>
      <el-table-column label="ID" min-width="100" prop="ID"></el-table-column>
      <el-table-column label="路由Name" min-width="160" prop="name"></el-table-column>
      <el-table-column label="是否隐藏" min-width="100" prop="hidden">
        <template slot-scope="scope">
          <span>{{scope.row.hidden?"隐藏":"显示"}}</span>
        </template>
      </el-table-column>
      <el-table-column label="父节点" min-width="90" prop="parentId"></el-table-column>
      <el-table-column label="排序" min-width="70" prop="sort"></el-table-column>
      <el-table-column label="文件路径" min-width="360" prop="component"></el-table-column>
      <el-table-column label="展示名称" min-width="120" prop="authorityName">
        <template slot-scope="scope">
          <span>{{scope.row.meta.title}}</span>
        </template>
      </el-table-column>
      <el-table-column label="图标" min-width="140" prop="authorityName">
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

    <el-dialog :before-close="handleClose" :visible.sync="dialogFormVisible" :title="dialogTitle">
      <el-form :inline="true" :model="form" :rules="rules" label-width="85px" ref="menuForm">
        <el-form-item label="路由name" prop="path">
          <el-input autocomplete="off" placeholder="唯一英文字符串" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="是否隐藏">
          <el-select placeholder="是否在列表隐藏" v-model="form.hidden">
            <el-option :value="false" label="否"></el-option>
            <el-option :value="true" label="是"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="父节点Id">
           <el-select
            placeholder="请选择"
            v-model="form.parentId"
            :disabled="!this.isEdit"
            filterable 
          >
            <el-option
              :disabled="canSelect(item)"
              :key="item.ID"
              :label="item.title"
              :value="item.ID"
              v-for="item in menuOption"
            >
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="文件路径" prop="component">
          <el-input autocomplete="off" v-model="form.component"></el-input>
        </el-form-item>
        <el-form-item label="展示名称" prop="meta.title">
          <el-input autocomplete="off" v-model="form.meta.title"></el-input>
        </el-form-item>
        <el-form-item label="图标" prop="meta.icon">
          <el-input autocomplete="off" v-model="form.meta.icon"></el-input>
        </el-form-item>
        <el-form-item label="排序标记" prop="sort">
          <el-input autocomplete="off" v-model.number="form.sort"></el-input>
        </el-form-item>
        <el-form-item label="keepAlive" prop="meta.keepAlive">
          <el-select placeholder="是否keepAlive缓存页面" v-model="form.meta.keepAlive">
            <el-option :value="false" label="否"></el-option>
            <el-option :value="true" label="是"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div class="warning">新增菜单需要在角色管理内配置权限才可使用</div>
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
  updateBaseMenu,
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
      dialogFormVisible: false,
      dialogTitle:"新增菜单",
      menuOption:[
        {
          ID:"0",
          title:"根菜单"
        }
      ],
      form: {
        ID: 0,
        path: '',
        name: '',
        hidden: '',
        parentId: '',
        component: '',
        meta: {
          title: '',
          icon: '',
          defaultMenu:false,
          keepAlive:false
        }
      },
      rules: {
        path: [{ required: true, message: '请输入菜单name', trigger: 'blur' }],
        component: [
          { required: true, message: '请输入文件路径', trigger: 'blur' }
        ],
        'meta.title': [
          { required: true, message: '请输入菜单展示名称', trigger: 'blur' }
        ]
      },
      isEdit: false
    }
  },
  methods: {
    handleClose(done) {
      this.initForm()
      done()
    },
    // 懒加载子菜单
    load(tree, treeNode, resolve) {
      resolve([
        {
          id: 31,
          date: '2016-05-01',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1519 弄'
        },
        {
          id: 32,
          date: '2016-05-01',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1519 弄'
        }
      ])
    },
    // 删除菜单
    deleteMenu(ID) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteBaseMenu({ ID })
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
    // 初始化弹窗内表格方法
    initForm() {
      this.$refs.menuForm.resetFields()
      this.form = {
        ID: 0,
        path: '',
        name: '',
        hidden: '',
        parentId: '',
        component: '',
        meta: {
          title: '',
          icon: '',
          defaultMenu:false,
          keepAlive:""
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
      this.$refs.menuForm.validate(async valid => {
        if (valid) {
          let res
          this.form.name = this.form.path
          if (this.isEdit) {
            res = await updateBaseMenu(this.form)
          } else {
            res = await addBaseMenu(this.form)
          }
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: '添加成功!'
            })
            this.getTableData()
          } else {
            this.$message({
              type: 'error',
              message: '添加失败!'
            })
          }
          this.initForm()
          this.dialogFormVisible = false
        }
      })
    },
    // 添加菜单方法，id为 0则为添加根菜单
    addMenu(id) {
      this.dialogTitle = "新增菜单"
      this.form.parentId = String(id)
      this.isEdit = false
      this.dialogFormVisible = true
    },
    // 修改菜单方法
    async editMenu(id) {
      this.dialogTitle = "编辑菜单"
      const res = await getBaseMenuById({ id })
      this.form = res.data.menu
      this.dialogFormVisible = true
      this.isEdit = true
    },
    getMenuList(MenuData){
      MenuData.map(item=>{
        this.menuOption.push({
          ID:String(item.ID),
          title:item.meta.title
        })
        if(item.children){
          this.getMenuList(item.children)
        }
      })
    },
    findAuthoritySelf(mune,muneData,outData){
      muneData&&muneData.some(item=>{
        if(item.ID == mune.ID){
          outData.push(item)
          return true
        }
        this.findAuthoritySelf(mune,item.children,outData)
      })
    },
    findAllChild(menu,array){
      menu&&menu.map(item=>{
        array.push(String(item.ID))
        this.findAllChild(item.children,array)
      })
    },
    canSelect(authority){
      const array = []
      const arrayIds = []
      this.findAuthoritySelf({ID:this.form.ID},this.tableData,array)
      this.findAllChild(array,arrayIds)
      return arrayIds.indexOf(authority.ID)>-1
    },
  },
   async created() {
    this.pageSize = 999
    await this.getTableData()
    await this.getMenuList(this.tableData)
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
.warning {
  color: #dc143c;
}
</style>