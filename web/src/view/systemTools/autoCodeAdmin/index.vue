<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="goAutoCode(null)">新增</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="结构体名" min-width="150" prop="structName" />
        <el-table-column align="left" label="结构体描述" min-width="150" prop="structCNName" />
        <el-table-column align="left" label="表名称" min-width="150" prop="tableName" />
        <el-table-column align="left" label="回滚标记" min-width="150" prop="flag">
          <template #default="scope">
            <el-tag
              v-if="scope.row.flag"
              type="danger"
              size="mini"
              effect="dark"
            >
              已回滚
            </el-tag>
            <el-tag
              v-else
              size="mini"
              type="success"
              effect="dark"
            >
              未回滚
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="180">
          <template #default="scope">
            <div>
              <el-button size="mini" type="text" :disabled="scope.row.flag === 1" @click="rollback(scope.row)">回滚</el-button>
              <el-button size="mini" type="text" @click="goAutoCode(scope.row)">复用</el-button>
              <el-button size="mini" type="text" @click="deleteRow(scope.row)">删除</el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索
import { getSysHistory, rollback, delSysHistory } from '@/api/autoCode.js'
import infoList from '@/mixins/infoList'

export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getSysHistory
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    async deleteRow(row) {
      this.$confirm('此操作将删除本历史, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async() => {
        const res = await delSysHistory({ id: Number(row.ID) })
        if (res.code === 0) {
          this.$message.success('删除成功')
          this.getTableData()
        }
      })
    },
    async rollback(row) {
      this.$confirm('此操作将删除自动创建的文件和api, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async() => {
        const res = await rollback({ id: Number(row.ID) })
        if (res.code === 0) {
          this.$message.success('回滚成功')
          this.getTableData()
        }
      })
    },
    goAutoCode(row) {
      if (row) {
        this.$router.push({ name: 'autoCodeEdit', params: {
          id: row.ID
        }})
      } else {
        this.$router.push({ name: 'autoCode' })
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
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>
