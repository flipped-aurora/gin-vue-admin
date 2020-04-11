<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addCoffeeType" type="primary">新增咖啡类型</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="咖啡类型名称" min-width="150" prop="name"></el-table-column>
      <el-table-column label="咖啡类型代码" min-width="150" prop="code"></el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="deleteCoffeeType(scope.row)" size="small" type="text">删除咖啡类型</el-button>
          <el-button @click="editCoffeeType(scope.row)" size="small" type="text">编辑咖啡类型</el-button>
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

    <el-dialog :visible.sync="addCoffeeTypeDialog" custom-class="user-dialog" title="新增咖啡类型">
      <el-form :model="coffeeTypeInfo">
        <el-form-item label="类型名称" label-width="80px">
          <el-input v-model="coffeeTypeInfo.name"></el-input>
        </el-form-item>
        <el-form-item label="类型代码" label-width="80px">
          <el-input v-model="coffeeTypeInfo.code" :disabled = dis></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddCoffeeTypeDialog">取 消</el-button>
        <el-button @click="enterAddCoffeeTypeDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API
import {
  getCoffeeTypeList,
  updateCoffeeType,
  addCoffeeType,
  delCoffeeType,
  getCoffeeTypeById
  
} from "@/api/coffee"
import infoList from "@/components/mixins/infoList"
export default {
  name: "coffeetype",
  mixins: [infoList],
  data() {
    return {
      listApi: getCoffeeTypeList,
      listKey: "coffeetype",
      path: path,
      addCoffeeTypeDialog: false,
      isEdit: false,
      coffeeTypeInfo: {
        code: "",
        name: ""
      }
    }
  },
  methods: {
    async enterAddCoffeeTypeDialog() {
      // eslint-disable-next-line no-console
      //console.log(this.coffeeInfo)
      let res
      if(this.isEdit) {
        res = await updateCoffeeType(this.coffeeTypeInfo)
      } else {
        res = await addCoffeeType(this.coffeeTypeInfo)
      }

      if (res.success) {
        this.$message({ type: "success", message: "创建成功" })
      }
      await this.getTableData()
      this.closeAddCoffeeDialog()
    },
    closeAddCoffeeTypeDialog() {
      this.coffeeTypeInfo = {
        name: "",
        code: ""
      }
      this.addCoffeeTypeDialog = false
    },
    addCoffeeType() {
      this.coffeeTypeInfo = {
        name: "",
        code: ""
      }
      this.isEdit = false
      this.dis = false
      this.addCoffeeTypeDialog = true
    },
    async editCoffeeType(row) {
        const res = await getCoffeeTypeById({code: row.code})
        if(res.success){
            this.coffeeTypeInfo = res.data.coffeetype
        }
        this.isEdit = true
        this.dis = true
        this.addCoffeeTypeDialog = true
    },
    async deleteCoffeeType(row) {
      const res = await delCoffeeType({ code: row.code })
      if(res.success) {
        this.$message({type: "success", message: "删除咖啡类型成功"})
      }
      await this.getTableData()
    }
  },
  created() {
      this.pageSize = 999
  },
}
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}
</style>