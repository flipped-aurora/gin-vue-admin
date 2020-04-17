<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addCustomer" type="primary">新增客户</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="图片" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.image" height="50" width="50" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="uuid" min-width="250" prop="uuid"></el-table-column>
      <el-table-column label="客户名称" min-width="150" prop="username">
      </el-table-column>
      <el-table-column label="密码" min-width="150" prop="password"></el-table-column>
      <el-table-column label="别名" min-width="150" prop="nickname"></el-table-column>
      <el-table-column label="邮箱" min-width="150" prop="email"></el-table-column>
      <el-table-column label="手机号" min-width="150" prop="phone"></el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="deleteCustomer(scope.row)" size="small" type="text">删除客户</el-button>
          <el-button @click="editCustomer(scope.row)" size="small" type="text">编辑客户</el-button>
          <el-button @click="addAddress(scope.row)" size="small" type="text">添加地址</el-button>
          <el-button @click="editAddress(scope.row)" size="mini" type="text">选择地址作为默认地址</el-button>
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

    <el-dialog :visible.sync="addCustomerDialog" custom-class="user-dialog" :title="titleMap[dialogTitle]">
      <el-form :model="customerInfo" :rules="rules" ref="customerInfo">
        <el-form-item label="客户名称" label-width="80px" required prop="username">
          <el-input v-model="customerInfo.username" ></el-input>
        </el-form-item>
        <el-form-item label="密码" label-width="80px" required prop="password">
          <el-input v-model="customerInfo.password" show-password ></el-input>
        </el-form-item>
        <el-form-item label="别名" label-width="80px" required>
          <el-input v-model="customerInfo.nickname" ></el-input>
        </el-form-item>
        <el-form-item label="图片" label-width="80px" required prop="image">
          <el-upload
            :on-success="handleAvatarSuccess"
            :show-file-list="false"
            :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
            class="avatar-uploader"
            name="file"
          >
            <img :src="customerInfo.image" class="avatar" v-if="customerInfo.image" />
            <i class="el-icon-plus avatar-uploader-icon" v-else></i>
          </el-upload>
        </el-form-item>
        <el-form-item label="邮箱" label-width="80px" required prop="email">
            <el-input v-model="customerInfo.email"></el-input>
        </el-form-item>
        <el-form-item label="手机号" label-width="80px" required prop="phone">
            <el-input v-model="customerInfo.phone" ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="closeAddCustomerDialog">取 消</el-button>
          <el-button @click="submitForm('customerInfo')" type="primary">确 定</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
    <el-dialog :visible.sync="addAddressDialog" custom-class="user-dialog" :title="addressTitleMap[addressDialogTitle]">
      <el-form :model="addressInfo" size="mini">
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
        <el-form-item label="是否默认?" label-width="80px">
          <el-checkbox v-model="addressInfo.isDefault"></el-checkbox>
        </el-form-item>
      </el-form>
      <el-button @click="closeAddressDialog">取消</el-button>
      <el-button @click="AddAddressDialog" type="primary">确定</el-button>
    </el-dialog>
    <el-dialog
    title="设置默认地址"
    :visible.sync="defaultAdressDialog"
    width="30%"
    center
    >
<el-dropdown trigger="click" ref="ddlAddress">
  <span class="el-dropdown-link">
    选择地址<span v-html="selectedAddress"></span><i class="el-icon-arrow-down el-icon--right"></i>
  </span>
  <el-dropdown-menu slot="dropdown">
    <el-table
      size="mini"
      stripe
      :data="userAddress"
      style="width: 100%"
      @row-click="handleRowClicked">
      <el-table-column prop="ID" label="id" width="50"></el-table-column>
      <el-table-column prop="consignee" label="收货人" width="80"></el-table-column>
      <el-table-column prop="userId" label="客户Id" width="280"></el-table-column> 
      <el-table-column prop="specAddress" label="详细地址" width="280"></el-table-column>    
    </el-table>
  </el-dropdown-menu>
</el-dropdown>
<div slot="footer" class="dialog-footer">
    <el-button @click="cloesDefaultAddressDialog">取 消</el-button>
    <el-button type="primary" @click="setDefaultAddress">确 定</el-button>
  </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getCustomerList,
  addCustomer,
  updateCustomer,
  delCustomer,
  getCustomerById,
  addCustomerAddress,
  getCustomerAddressByUserId,
  setCustomerDefaultAdress
} from "@/api/customer"
import infoList from "@/components/mixins/infoList"
import { regionData,CodeToText  } from 'element-china-area-data'
export default {
  name: "customer",
  mixins: [infoList],
  data() {
    const checkUsername = (rule, value, callback) => {
      if(!value) {
        return callback(new Error('用户名不能为空'))
      }
      if (value.length < 5 || value.length > 12) {
        return callback(new Error('请输入正确的用户名'))
      } else {
        callback()
      }
    }
    const checkPassword = (rule, value, callback) => {
      if(!value) {
        return callback(new Error('密码不能为空'))
      }
      if (value.length < 6 || value.length > 12) {
        return callback(new Error('请输入正确的密码'))
      } else {
        callback()
      }
    }
    const checkEmail = (rule, value, callback) => {
      const mailReg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
      if (!value) {
        return callback(new Error('邮箱不能为空'))
      }
      setTimeout(() => {
       if (mailReg.test(value)) {
         callback()
       } else {
          callback(new Error('请输入正确的邮箱格式'))
        }
      }, 100)
    }
    const checkPhone = (rule, value, callback) => {
      const phoneReg = /^1[3|4|5|7|8][0-9]{9}$/
      if (!value) {
        return callback(new Error('电话号码不能为空'))
      }
      setTimeout(() => {
      // Number.isInteger是es6验证数字是否为整数的方法,但是我实际用的时候输入的数字总是识别成字符串
      // 所以我就在前面加了一个+实现隐式转换
        if (!Number.isInteger(+value)) {
          callback(new Error('请输入数字值'))
        } else {
          if (phoneReg.test(value)) {
            callback()
          } else {
            callback(new Error('电话号码格式不正确'))
          }
      }
    }, 100)
    }
    return {
      listApi: getCustomerList,
      listKey: "cusList",
      path: path,
      addCustomerDialog: false,
      dialogTitle: '',
      titleMap : {
        addData : "添加客户",
        updateData : "修改客户"
      },
      addressDialogTitle: "",
      addressTitleMap: {
        addData: "添加地址",
        updateData: "修改地址"
      },
      isEdit: false,
      customerInfo: {
        username: "",
        password: "",
        nickname: "",
        image: "",
        email: "",
        phone: ""
      },
      uuid:'',
      rules: {
        username: [{required: true, validator: checkUsername, trigger: 'blur' }],
        password: [{required: true, validator: checkPassword, trigger: 'blur' }],
        email:  [{required: true, validator: checkEmail, trigger: 'blur' }],
        phone: [{required: true, validator: checkPhone, trigger: 'blur'}]
      },
      addAddressDialog: false,
      addressInfo: {
        options: regionData,
        selectedOptions: [],
        specAddress: '',
        phone: '',
        consignee: '',
        isDefault: false
      },
      province: '',
      city: '',
      region: '',
      defaultAdressDialog: false,
      userAddress: [],
      selectedAddress: "",
      addressId: 0
    };
  },
  methods: {
    async submitForm(formName) {
      // eslint-disable-next-line no-console
      //console.log(this.coffeeInfo)
      this.$refs[formName].validate(async v => {
        //alert(v)
        if (v) {
          let res
          if (this.isEdit) {
            res = await updateCustomer(this.customerInfo);
          } else {
            res = await addCustomer(this.customerInfo);
          }
          if (res.success) {
            this.$message({ type: "success", message: "创建成功" });
          }
          await this.getTableData();
          this.closeAddCustomerDialog();
        } else {
          this.$message({
            type: 'error',
            message: '请正确填写信息',
            showClose: true
          })
          return false
        }
      })
    },
    closeAddCustomerDialog() {
      this.customerInfo =  {
        username: "",
        password: "",
        nickname: "",
        image: "",
        email: "",
        phone: ""
      }
      this.addCustomerDialog = false;
    },
    handleAvatarSuccess(res) {
      this.customerInfo.image = res.data.file.url;
    },
    addCustomer() {
      this.customerInfo =  {
        username: "",
        password: "",
        nickname: "",
        image: "",
        email: "",
        phone: ""
      }
      this.dialogTitle = "addData";
      this.isEdit = false;
      this.addCustomerDialog = true;
    },
    async editCustomer(row) {
      const res = await getCustomerById({ uuid: row.uuid });
      this.customerInfo = res.data.customer;
      this.isEdit = true;
      this.dialogTitle = "updateData";
      this.addCustomerDialog = true;

    },
    async deleteCustomer(row) {
      this.$confirm('此操作将永久删除该客户, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        const res = await delCustomer({ uuid: row.uuid });
        if (res.success) {
          this.$message({ type: "success", message: "删除客户成功" });
        }
        await this.getTableData();
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
    },
    addAddress(row) {
      //alert(row.uuid)
      this.uuid = row.uuid
      this.addAddressDialog = true
      this.addressDialogTitle = "addData"
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
      this.addAddressDialog = false
    },
    async AddAddressDialog() {
      if(this.province && this.city && this.region) {
        //alert(this.uuid)
        var isDefault
        if(this.addressInfo.isDefault == true) {
          isDefault = 1
        } else {
          isDefault = 0
        }
        const res = await addCustomerAddress({userId: this.uuid,province: this.province, city: this.city,town: this.region, specAddress: this.province + this.city + this.region + this.addressInfo.specAddress,phone: this.addressInfo.phone,consignee: this.addressInfo.consignee,isDefault: isDefault})
        if(res.success) {
          this.$message({ type: "success", message: "添加地址成功" })
        }
      }else {
        this.$message.Error("请选择省市区")
      }

      this.addAddressDialog = false
    },
    handleRowClicked(row, column, event){
      this.selectedAddress = "("+row.specAddress +")";
      this.addressId = row.ID
      this.uuid = row.userId
      this.$refs.ddlAddress.hide();
    },
    async editAddress(row) {
      this.uuid = row.uuid
      const res = await getCustomerAddressByUserId({pageInfo: {page:1,pageSize: 999},user_id: this.uuid})
      if(res.success) {
        this.userAddress = res.data.addressList
        //console.log(res.data.addressList)
        for (let i = 0; i < res.data.addressList.length; i++) {
            if(res.data.addressList[i].isDefault == 1) {
              this.selectedAddress = "("+res.data.addressList[i].specAddress +")";
            }
        }
      }
      this.defaultAdressDialog = true
    },
    cloesDefaultAddressDialog() {
      this.selectedAddress = ""
      this.defaultAdressDialog = false
    },
    async setDefaultAddress() {
      const res = await setCustomerDefaultAdress({userId: this.uuid,id: this.addressId})
      console.log(res)
      if(res.success) {
        this.$message({ type: "success", message: "设置成功" })
      } else {
        this.$message.Error("设置失败")
      }
      this.defaultAdressDialog = false
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
  .el-icon-arrow-down {
    font-size: 12px;
  }
}
</style>