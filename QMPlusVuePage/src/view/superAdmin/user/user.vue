<template>
  <div>
    <div class="button-box clearflex"></div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="头像" min-width="50">
        <template slot-scope="scope">
          <img :src="scope.row.headerImg" width="50" height="50" />
        </template>
      </el-table-column>
      <el-table-column label="id" min-width="60" prop="ID"></el-table-column>
      <el-table-column label="用户名" min-width="150" prop="userName"></el-table-column>
      <el-table-column label="昵称" min-width="150" prop="nickName"></el-table-column>
      <el-table-column label="用户级别" min-width="150">
        <template slot-scope="scope">
          <div>{{scope.row.authority.authorityName}}</div>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="text">编辑</el-button>
          <el-button @click="deleteApi(scope.row)" size="small" type="text">删除</el-button>
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
      hide-on-single-page
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
import { getUserList } from '@/api/user'
import infoList from '@/view/superAdmin/mixins/infoList'

export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getUserList,
      listKey: 'userList',
      dialogFormVisible: false,
      type: ''
    }
  },
  methods: {
    initForm() {
      this.form = {
        path: '',
        group: '',
        description: ''
      }
    }
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