<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="50" prop="ID"></el-table-column>        
      <el-table-column label="useId" min-width="250" prop="userId"></el-table-column>
      <el-table-column label="省" min-width="150" prop="province"></el-table-column>
      <el-table-column label="市" min-width="150" prop="city"></el-table-column>
      <el-table-column label="区/镇/村" min-width="150" prop="town"></el-table-column>
      <el-table-column label="具体地址" min-width="150" prop="specAddress"></el-table-column>
      <el-table-column label="收货人" min-width="150" prop="consignee"></el-table-column>        
      <el-table-column label="电话" min-width="150" prop="phone"></el-table-column>           
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="editAddress(scope.row)" size="small" type="text">编辑地址</el-button>
          <el-button @click="deleteAddress(scope.row)" size="small" type="text">删除地址</el-button>
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
  <el-dialog title="编辑地址" :visible.sync="editAddressDialog">
      <el-form :model="addressInfo" size="mini">
        <span>请选择地址</span>
        <el-form-item>
          <el-cascader
          :options="addressInfo.options"
          v-model="addressInfo.selectedOptions"
          @change="changeAddress">
          </el-cascader>            
        </el-form-item>
        <el-form-item label="具体地址" label-width="80px">
          <el-input v-model="addressInfo.specAddress" size="mini"></el-input>
        </el-form-item>
        <el-form-item label="手机号" label-width="80px">
          <el-input v-model="addressInfo.phone" size="mini"></el-input>
        </el-form-item>
        <el-form-item label="收货人" label-width="80px">
          <el-input v-model="addressInfo.consignee" size="mini"></el-input>
        </el-form-item>
      </el-form>
      <el-button @click="closeAddressDialog">取消</el-button>
      <el-button @click="AddAddressDialog" type="primary">确定</el-button>
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false">取 消</el-button>
    <el-button type="primary" @click="dialogFormVisible = false">确 定</el-button>
  </div>
</el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;

import {
  getCustomerAddressList,
  updateCustomerAddress,
  getAddressById,
  delCustomerAddress
} from "@/api/customer";
import infoList from "@/components/mixins/infoList";
import { regionData,TextToCode } from 'element-china-area-data'
export default {
  name: "addressList",
  mixins: [infoList],
  data() {
    return {
      listApi: getCustomerAddressList,
      listKey: "addressList",
      path: path,
      addressInfo: {
        options: regionData,
        selectedOptions: [],
        specAddress: '',
        phone: '',
        consignee: '',
        isDefault: false
      },
      editAddressDialog: false,
      province: "",
      city: "",
      region: ""
    }
  },
  methods: {
    async enterAddAddressDialog() {
      // eslint-disable-next-line no-console
      //console.log(this.coffeeInfo)
      const res = await updateCustomerAddress(this.addressInfo)
      if (res.success) {
        this.$message({ type: "success", message: "修改地址成功" });
      }
      await this.getTableData()
      this.closeAddAddressDialog()
    },
    closeAddAddressDialog() {
      this.addressInfo = {
        options: regionData,
        selectedOptions: [],
        province: '',
        city: '',
        town: '',
        specAddress: '',
        phone: '',
        consignee: '',
        isDefault: false
      },
      this.editAddressDialog = false
    },
    async editAddress(row) {
      const res = await getAddressById({ id: row.ID })
      this.addressInfo = res.data.address
      console.log(this.addressInfo)
      this.addressInfo.selectedOptions = [TextToCode[this.addressInfo.province].code,TextToCode[this.addressInfo.province][this.addressInfo.city].code,TextToCode[this.addressInfo.province][this.addressInfo.city][this.addressInfo.town].code]
      this.addressInfo.options = regionData
      this.editAddressDialog = true

    },
    async deleteAddress(row) {
      this.$confirm('此操作将永久删除该客户下的地址, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async ()=> {
        const res = await delCustomerAddress({ id: row.id })
        if (res.success) {
          this.$message({ type: "success", message: "删除地址成功" })
        }
        await this.getTableData()
      }).catch(() => {
           this.$message({
            type: 'info',
            message: '已取消删除'
          })       
      })
    },
    changeAddress(value) {
      this.province =CodeToText[value[0]]
      this.city = CodeToText[value[1]]
      this.region = CodeToText[value[2]]
    },
    closeAddressDialog() {
      this.addressInfo = {
        options: regionData,
        selectedOptions: [],
        specAddress: ''       
      }
      this.editAddressDialog = false
    },
    async AddAddressDialog() {
      this.editAddressDialog = false
    }
  },
  async created() {
    this.page = 1
    this.pageSize = 999
  }
};
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