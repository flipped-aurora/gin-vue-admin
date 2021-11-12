<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="展示值">
          <el-input v-model="searchInfo.label" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="字典值">
          <el-input v-model="searchInfo.value" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="启用状态" prop="status">
          <el-select v-model="searchInfo.status" placeholder="请选择">
            <el-option key="true" label="是" value="true" />
            <el-option key="false" label="否" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="el-icon-refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增字典项</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="展示值" prop="label" width="120" />

        <el-table-column align="left" label="字典值" prop="value" width="120" />

        <el-table-column align="left" label="启用状态" prop="status" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>

        <el-table-column align="left" label="排序标记" prop="sort" width="120" />

        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button size="small" type="text" icon="el-icon-edit" @click="updateSysDictionaryDetail(scope.row)">变更</el-button>
            <el-popover :visible="scope.row.visible" placement="top" width="160">
              <p>确定要删除吗？</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                <el-button type="primary" size="mini" @click="deleteSysDictionaryDetail(scope.row)">确定</el-button>
              </div>
              <template #reference>
                <el-button type="text" icon="el-icon-delete" size="mini">删除</el-button>
              </template>
            </el-popover>
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

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="110px">
        <el-form-item label="展示值" prop="label">
          <el-input
            v-model="formData.label"
            placeholder="请输入展示值"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item label="字典值" prop="value">
          <el-input-number
            v-model.number="formData.value"
            step-strictly
            :step="1"
            placeholder="请输入字典值"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item label="启用状态" prop="status" required>
          <el-switch v-model="formData.status" active-text="开启" inactive-text="停用" />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort">
          <el-input-number v-model.number="formData.sort" placeholder="排序标记" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createSysDictionaryDetail,
  deleteSysDictionaryDetail,
  updateSysDictionaryDetail,
  findSysDictionaryDetail,
  getSysDictionaryDetailList
} from '@/api/sysDictionaryDetail' //  此处请自行替换地址
import infoList from '@/mixins/infoList'

export default {
  name: 'SysDictionaryDetail',
  mixins: [infoList],
  data() {
    return {
      listApi: getSysDictionaryDetailList,
      dialogFormVisible: false,
      type: '',
      formData: {
        label: null,
        value: null,
        status: true,
        sort: null
      },
      rules: {
        label: [
          {
            required: true,
            message: '请输入展示值',
            trigger: 'blur'
          }
        ],
        value: [
          {
            required: true,
            message: '请输入字典值',
            trigger: 'blur'
          }
        ],
        sort: [
          {
            required: true,
            message: '排序标记',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  created() {
    this.searchInfo.sysDictionaryID = Number(this.$route.params.id)
    this.getTableData()
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.status === '') {
        this.searchInfo.status = null
      }
      this.getTableData()
    },
    async updateSysDictionaryDetail(row) {
      const res = await findSysDictionaryDetail({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.resysDictionaryDetail
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        label: null,
        value: null,
        status: true,
        sort: null,
        sysDictionaryID: ''
      }
    },
    async deleteSysDictionaryDetail(row) {
      row.visible = false
      const res = await deleteSysDictionaryDetail({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      this.formData.sysDictionaryID = Number(this.$route.params.id)
      this.$refs['elForm'].validate(async valid => {
        if (!valid) return
        let res
        switch (this.type) {
          case 'create':
            res = await createSysDictionaryDetail(this.formData)
            break
          case 'update':
            res = await updateSysDictionaryDetail(this.formData)
            break
          default:
            res = await createSysDictionaryDetail(this.formData)
            break
        }
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '创建/更改成功'
          })
          this.closeDialog()
          this.getTableData()
        }
      })
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  }
}
</script>

<style>
</style>
