<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="订单编号" min-width="250" prop="orderId"></el-table-column>        
      <el-table-column label="用户编号" min-width="250" prop="userId"></el-table-column>
      <el-table-column label="收货人" min-width="150" prop="consignee"></el-table-column>
      <el-table-column label="总价格" min-width="150" prop="value"></el-table-column>
      <el-table-column label="电话" min-width="150" prop="phone"></el-table-column>
      <el-table-column label="订单类型" min-width="150" prop="orderType"></el-table-column>      
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button @click="deleteOrder(scope.row)" size="small" type="text">删除订单</el-button>
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

  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API;
import {
  getCustomerOrderList,
  getCustomerOrderByOrderId,
  deleteCustomerOrder
} from "@/api/customer";
import infoList from "@/components/mixins/infoList";
export default {
  name: "orderList",
  mixins: [infoList],
  data() {
    return {
      listApi: getCustomerOrderList,
      listKey: "orderList",
      path: path,
      orderInfo: {
        userId: "",
        orderId: "",
        consignee: "",
        value: "",
        specAddress: "",
        phone: "",
        orderType: 0,
        orderDetail: {
          orderId: "",
          coffee: {
            uuid: "",
            name: "",
            value: 0.0,
            des: "",
            img: "",
            type: {
              code: "",
              name: "",
              image: ""
            },
            code: "",
            spec: []
          },
          coffee_id: "",
          spec: "",
          count: 0,
          value: 0.0
        }
      }
    }
  },
  methods: {
    async deleteOrder(row) {
      const res = await deleteCustomerOrder({ orderId: row.orderId })
      if (res.success) {
        this.$message({ type: "success", message: "删除订单成功" })
      }
      await this.getTableData()
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