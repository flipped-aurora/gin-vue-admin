<template>
  <div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="250" prop="ID"></el-table-column>        
      <el-table-column label="useId" min-width="250" prop="userId"></el-table-column>
      <el-table-column label="省" min-width="150" prop="province"></el-table-column>
      <el-table-column label="市" min-width="150" prop="city"></el-table-column>
      <el-table-column label="区/镇/村" min-width="150" prop="town"></el-table-column>
      <el-table-column label="具体地址" min-width="150" prop="specAddress"></el-table-column>      
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
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
export default {
  name: "addressList",
  mixins: [infoList],
  data() {
    return {
      listApi: getCustomerAddressList,
      listKey: "addressList",
      path: path,
      addAddressDialog: false,
      addressInfo: {
        userId: "",
        province: 0.0,
        city: "",
        town: "",
        specAddress: ""
      }
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
        userId: "",
        province: 0.0,
        city: "",
        town: "",
        specAddress: ""
      }
      this.addAddressDialog = false
    },
    async editAddress(row) {
      const res = await getAddressById({ uuid: row.uuid })
      this.addressInfo = res.data.address
      this.addAddressDialog = true

    },
    async deleteAddress(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
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