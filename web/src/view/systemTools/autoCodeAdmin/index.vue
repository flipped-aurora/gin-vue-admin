<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="goAutoCode(null)">{{ $t('general.add') }}</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" :label="$t('general.date')" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="$t('autoCode.structName')" min-width="150" prop="structName" />
        <el-table-column align="left" :label="$t('autoCode.structChineseName')" min-width="150" prop="structCNName" />
        <el-table-column align="left" :label="$t('autoCode.tableName')" min-width="150" prop="tableName" />
        <el-table-column align="left" :label="$t('autoCodeAdmin.rollBackMark')" min-width="150" prop="flag">
          <template #default="scope">
            <el-tag
              v-if="scope.row.flag"
              type="danger"
              size="mini"
              effect="dark"
            >
              {{ $t('autoCodeAdmin.rolledBack') }}
            </el-tag>
            <el-tag
              v-else
              size="mini"
              type="success"
              effect="dark"
            >
              {{ $t('autoCodeAdmin.notRolledBack') }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" :lable="$t('general.operations')" min-width="180">
          <template #default="scope">
            <div>
              <el-button size="mini" type="text" :disabled="scope.row.flag === 1" @click="rollback(scope.row)">{{ $t('autoCodeAdmin.rollBack') }}</el-button>
              <el-button size="mini" type="text" @click="goAutoCode(scope.row)">{{ $t('autoCodeAdmin.reuse') }}</el-button>
              <el-button size="mini" type="text" @click="deleteRow(scope.row)">{{ $t('general.delete') }}</el-button>
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
      this.$confirm(this.$t('autoCodeAdmin.deleteHistoryConfirm'), this.$t('general.hint'), {
        confirmButtonText: this.$t('general.confirm'),
        cancelButtonText: this.$t('general.cancel'),
        type: 'warning'
      }).then(async() => {
        const res = await delSysHistory({ id: Number(row.ID) })
        if (res.code === 0) {
          this.$message.success(this.$t('general.deleteSuccess'))
          this.getTableData()
        }
      })
    },
    async rollback(row) {
      this.$confirm(this.$t('autoCodeAdmin.rollbackConfirm'), this.$t('general.hint'), {
        confirmButtonText: this.$t('general.confirm'),
        cancelButtonText: this.$t('general.cancel'),
        type: 'warning'
      }).then(async() => {
        const res = await rollback({ id: Number(row.ID) })
        if (res.code === 0) {
          this.$message.success(this.$t('autoCodeAdmin.rollbackSuccess'))
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
