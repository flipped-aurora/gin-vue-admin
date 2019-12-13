<template>
  <div>
    <div class="button-box clearflex"></div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="头像" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.headerImg" height="50" width="50" />
          </div>
        </template>
      </el-table-column>
      <el-table-column label="uuid" min-width="250" prop="uuid"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="userName"></el-table-column>
      <el-table-column label="昵称" min-width="150" prop="nickName"></el-table-column>
      <el-table-column label="用户角色" min-width="150">
        <template slot-scope="scope">
          <el-select
            @change="changeAuthority(scope.row)"
            placeholder="请选择"
            v-model="scope.row.authority.authorityId"
          >
            <el-option
              :key="item.authorityId"
              :label="item.authorityName"
              :value="item.authorityId"
              v-for="item in authOptions"
            ></el-option>
          </el-select>
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
import { getUserList, setUserAuthority } from '@/api/user'
import { getAuthorityList } from '@/api/authority'
import infoList from '@/components/mixins/infoList'

export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getUserList,
      listKey: 'userList',
      authOptions: []
    }
  },
  methods: {
    async changeAuthority(row) {
      const res = await setUserAuthority({
        uuid: row.uuid,
        authorityId: row.authority.authorityId
      })
      if (res.success) {
        this.$message({ type: 'success', message: '角色设置成功' })
      }
    }
  },
  async created() {
    const res = await getAuthorityList({ page: 1, pageSize: 999 })
    this.authOptions = res.data.list
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