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
      <el-table-column label="客户名称" min-width="150" prop="username"></el-table-column>
      <el-table-column label="密码" min-width="150" prop="password"></el-table-column>
      <el-table-column label="别名" min-width="150" prop="nickname"></el-table-column>
      <el-table-column label="邮箱" min-width="150" prop="email"></el-table-column>
      <el-table-column label="手机号" min-width="150" prop="phone"></el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="deleteCustomer(scope.row)" size="small" type="text">删除客户</el-button>
          <el-button @click="editCustomer(scope.row)" size="small" type="text">编辑客户</el-button>
          <el-button @click="editAddress(scope.row)" size="small" type="text">添加地址</el-button>
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
    <el-dialog :visible.sync="editAddressDialog" custom-class="user-dialog" title="添加地址">
      <el-form :model="addressInfo">
        <el-form-item>
          <el-cascader
          size="larget"
          :options="addressInfo.options"
          v-model="addressInfo.selectedOptions"
          @change="changeAddress">
          </el-cascader>            
        </el-form-item>
        <el-form-item label="具体地址" label-width="80px">
          <el-input v-model="addressInfo.specAddress"></el-input>
        </el-form-item>
      </el-form>
      <el-button @click="closeAddressDialog">取消</el-button>
      <el-button @click="AddAddressDialog" type="primary">确定</el-button>
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
  addCustomerAddress
} from "@/api/customer"
import infoList from "@/components/mixins/infoList"
import { regionData  } from 'element-china-area-data'
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
      editAddressDialog: false,
      addressInfo: {
        options: regionData,
        selectedOptions: [],
        specAddress: ''
      },
      province: '',
      city: '',
      region: ''
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
      const res = await delCustomer({ uuid: row.uuid });
      if (res.success) {
        this.$message({ type: "success", message: "删除客户成功" });
      }
      await this.getTableData();
    },
    editAddress(row) {
      //alert(row.uuid)
      this.uuid = row.uuid
      this.editAddressDialog = true
    },
    changeAddress(value) {
      this.province =value[0]
      this.city = value[1]
      this.region = value[2]
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
      if(this.province && this.city && this.region) {
        //alert(this.uuid)
        const res = await addCustomerAddress({userId: this.uuid,province: this.province, city: this.city,town: this.region, specAddress: this.addressInfo.specAddress})
        if(res.success) {
          this.$message({ type: "success", message: "添加地址成功" })
        }
      }else {
        this.$message.Error("请选择省市区")
      }

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