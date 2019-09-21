<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addApi" type="primary">新增api</el-button>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="180" prop="ID"></el-table-column>
      <el-table-column label="api路径" min-width="180" prop="path"></el-table-column>
      <el-table-column label="api简介" min-width="180" prop="description"></el-table-column>
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

import { getApiList } from '@/api/api'
import infoList from '@/view/superAdmin/mixins/infoList'

export default {
  name: 'Api',
  mixins:[infoList],
  data() {
    return {
      listApi: getApiList,
      listKey:'list'
    }
  },
  methods: {

    addApi() {},
    editApi() {},
    deleteApi() {}
  },
  created() {
    this.getTableData()
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