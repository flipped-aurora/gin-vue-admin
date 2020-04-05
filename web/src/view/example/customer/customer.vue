<template>
  <div>
      <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增客户</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>
      <el-table-column label="接入日期" width="180">
        <template slot-scope="scope">{{ scope.row.CreatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="姓名" prop="customerName" width="120"></el-table-column>
      <el-table-column label="电话" prop="customerPhoneData" width="120"></el-table-column>
      <el-table-column label="接入人ID" prop="sysUserId" width="120"></el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateCustomer(scope.row)" size="small" type="text">变更</el-button>
          <el-popover
          placement="top"
          width="160"
          v-model="visible">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin: 0">
            <el-button size="mini" type="text" @click="visible = false">取消</el-button>
            <el-button type="primary" size="mini" @click="deleteCustomer(scope.row)">确定</el-button>
          </div>
          <el-button type="text" size="mini" slot="reference">删除</el-button>
        </el-popover>
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增Api">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="客户名">
          <el-input autocomplete="off" v-model="form.customerName"></el-input>
        </el-form-item>
        <el-form-item label="客户电话">
          <el-input autocomplete="off" v-model="form.customerPhoneData"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
    在资源权限中将此角色的资源权限清空 或者不包含创建者的角色 即可屏蔽此客户资源的显示
  </div>
</template>

<script>

import {
  createExaCustomer,
  updateExaCustomer,
  deleteExaCustomer,
  getExaCustomer,
  getExaCustomerList
} from '@/api/customer'
import { formatTimeToStr } from '@/utils/data'
import infoList from '@/components/mixins/infoList'
import { mapGetters } from 'vuex'

export default {
  name: 'Customer',
   mixins: [infoList],
  data(){
    return{
      listApi: getExaCustomerList,
      listKey: 'customer',
      dialogFormVisible:false,
      visible:false,
      type:"",
      form:{
        customerName:"",
        customerPhoneData:""
      }
    }
  },
  computed:{
    ...mapGetters('user', ['token'])
  },
   filters: {
    formatDate: function(time) {
      if (time != null && time != '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    }
  },
  methods:{
    async updateCustomer(row){
      const res = await getExaCustomer(row)
      this.type = "update"
      if(res.code == 0){
        this.form = res.data.customer        
        this.dialogFormVisible = true
      }
    },
    closeDialog(){
      this.dialogFormVisible = false
    },
    async deleteCustomer(row){
      this.visible = false
      const res = await deleteExaCustomer(row)
      if (res.code == 0){
         this.getTableData()
      }
    },
    async enterDialog(){
      let res 
      switch (this.type) {
        case "create":
          res =await createExaCustomer(this.form)
             break;
        case "update":
           res =await updateExaCustomer(this.form)
              break;
        default:
          res =await createExaCustomer(this.form)
             break;

      }
     
      if(res.code == 0){
        this.dialogFormVisible = false
        this.getTableData()
      }
    },
     openDialog() {
      this.type = "create"
      this.dialogFormVisible = true
    }
  }
}
</script>

<style>
</style>