<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addCoffeeSpec" type="primary">新增咖啡规格</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="80" prop="ID"></el-table-column>  
      <el-table-column label="咖啡规格编号" min-width="150" prop="spec_id"></el-table-column>
      <el-table-column label="咖啡规格名称" min-width="150" prop="name"></el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="deleteCoffeeSpec(scope.row)" size="small" type="text">删除咖啡规格</el-button>
          <el-button @click="editCoffeeSpec(scope.row)" size="small" type="text">编辑咖啡规格</el-button>
          <el-button @click="addCoffeeSpecDetail(scope.row)" size="small" type="text">添加咖啡规格明细</el-button>
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

    <el-dialog :visible.sync="addCoffeeSpecDialog" custom-class="user-dialog" :title="titleMap[dialogTitle]">
      <el-form :model="coffeeSpecInfo" :rules="rules" ref="coffeeSpecInfo">
        <el-form-item label="规格名称" label-width="80px" required>
          <el-input v-model="coffeeSpecInfo.name"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddCoffeeSpecDialog">取 消</el-button>
        <el-button @click="enterAddCoffeeSpecDialog('coffeeSpecInfo')" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="addCoffeeSpecDetailDialog" custom-class="user-dialog" title="添加咖啡规格明细">
      <el-form :model="coffeeSpecDetailInfo"  ref="coffeeSpecDetailInfo">
        <el-form-item label="咖啡" label-width="80px" required prop="coffee_id">
          <el-select placeholder="请选择" v-model="coffeeSpecDetailInfo.coffee_id">
            <el-option
              :key="item.uuid"
              :label="item.name"
              :value="item.uuid"
              v-for="item in coffeeOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="规格" label-width="80px" required>
          <el-input v-model="coffeeSpecDetailInfo.value"></el-input>
        </el-form-item>
        <el-form-item label="加价" label-width="80px" required>
          <el-input v-model="coffeeSpecDetailInfo.price_incre"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeAddCoffeeSpecDetailDialog">取 消</el-button>
        <el-button @click="enterCoffeeSpecDetailDialog('coffeeSpecDetailInfo')" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API
import {
  getCoffeeSpecListFunc,
  updateCoffeeSpecFunc,
  addCoffeeSpecFunc,
  delCoffeeSpecFunc,
  getCoffeeSpecByIdFunc,
  addCoffeeSpecDetailFunc,
  getCoffeeList,
  getCoffeeSpecDetail
} from "@/api/coffee"
import infoList from "@/components/mixins/infoList"
export default {
  name: "coffeespec",
  mixins: [infoList],
  data() {
    const checkname = (rule, value, callback) => {
      if(!value) {
        return callback(new Error('咖啡规格名不能为空'))
      } else {
        callback()
      }
    }     
    return {
      listApi: getCoffeeSpecListFunc,
      listKey: "coffeeSpec",
      path: path,
      addCoffeeSpecDialog: false,
      addCoffeeSpecDetailDialog: false,
      isEdit: false,
      coffeeSpecInfo: {
        spec_id: "",
        name: "",
        coffee_spec_detail: []
      },
      coffeeSpecDetailInfo: {
        spec_id: "",
        coffee_id: "",
        value: "",
        price_incre: 0.0
      },
      dialogTitle: '',
      titleMap : {
        addData : "添加咖啡规格",
        updateData : "修改咖啡规格"
      },
      rules: {
        name: [{required: true, validator: checkname, trigger: 'blur' }],
      },
      coffeeOptions: []
    }
  },
  methods: {
    async enterAddCoffeeSpecDialog(formName) {
      // eslint-disable-next-line no-console
      //console.log(this.coffeeInfo)
    this.$refs[formName].validate(async v => {
      let res
      if(this.isEdit) {
        res = await updateCoffeeSpec(this.coffeeSpecInfo)
      } else {
        res = await addCoffeeSpec(this.coffeeSpecInfo)
      }

      if (res.success) {
        this.$message({ type: "success", message: "创建成功" })
      }
      await this.getTableData()
      this.closeAddCoffeeSpecDialog()
    })
  },
    closeAddCoffeeSpecDialog() {
      this.coffeeSpecInfo = {
        spec_id: "",
        name: "",
        coffee_spec_detail: []
      }
      this.addCoffeeSpecDialog = false
    },
    addCoffeeSpec() {
      this.coffeeSpecInfo = {
        spec_id: "",
        name: "",
        coffee_spec_detail: []
      }
      this.isEdit = false
      this.dialogTitle = 'addData'
      this.addCoffeeSpecDialog = true
    },
    async editCoffeeSpec(row) {
        console.log(row.ID)
        const res = await getCoffeeSpecByIdFunc({id: row.ID})
        console.log(res)
        if(res.success){
            this.coffeeSpecInfo = res.data.coffeeSpec
        }
        this.isEdit = true
        this.dialogTitle = 'updateData'
        this.addCoffeeSpecDialog = true
    },
    async deleteCoffeeSpec(row) {
      const res = await delCoffeeSpecFunc({ id: row.ID })
      if(res.success) {
        this.$message({type: "success", message: "删除咖啡规格成功"})
      }
      await this.getTableData()
    },
    async addCoffeeSpecDetail(row) {
        this.coffeeSpecDetailInfo.spec_id = row.spec_id
        this.addCoffeeSpecDetailDialog = true
    },
    closeAddCoffeeSpecDetailDialog() {
        this.coffeeSpecDetailInfo = {
            spec_id: "",
            coffee_id: "",
            value: "",
            price_incre: 0.0           
        }
        this.addCoffeeSpecDetailDialog = false
    },
    async enterCoffeeSpecDetailDialog(formName) {
    this.$refs[formName].validate(async v => {
      this.coffeeSpecDetailInfo.price_incre = Number(this.coffeeSpecDetailInfo.price_incre)
      const res = await addCoffeeSpecDetailFunc({"spec_id": this.coffeeSpecDetailInfo.spec_id,"coffee_id": this.coffeeSpecDetailInfo.coffee_id,"value": this.coffeeSpecDetailInfo.value,"price_incre": this.coffeeSpecDetailInfo.price_incre})
      if (res.success) {
        this.$message({ type: "success", message: "创建成功" })
      }
      await this.getTableData()
      this.closeAddCoffeeSpecDetailDialog()
    })
    }
  },
  async created() {
      var res = await getCoffeeList({"page": 1, "pageSize": 999})
      this.coffeeOptions = res.data.coffeeList
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