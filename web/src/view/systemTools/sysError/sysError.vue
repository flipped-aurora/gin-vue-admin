<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAtRange">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>

          <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="!w-380px"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>

        <el-form-item label="错误来源" prop="form">
          <el-input v-model="searchInfo.form" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="错误内容" prop="info">
          <el-input v-model="searchInfo.info" placeholder="搜索条件" />
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button
            link
            type="primary"
            icon="arrow-down"
            @click="showAllQuery = true"
            v-if="!showAllQuery"
            >展开</el-button
          >
          <el-button
            link
            type="primary"
            icon="arrow-up"
            @click="showAllQuery = false"
            v-else
            >收起</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >删除</el-button
        >
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column
          sortable
          align="left"
          label="日期"
          prop="CreatedAt"
          width="180"
        >
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="错误来源"
          prop="form"
          width="120"
        />

        <el-table-column
          align="left"
          label="错误等级"
          prop="level"
          width="120"
        >
          <template #default="scope">
            <el-tag
              effect="dark"
              :type="levelTagMap[scope.row.level] || 'info'"
            >
              {{ levelLabelMap[scope.row.level] || defaultLevelLabel }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="处理状态"
          prop="status"
          width="140"
        >
          <template #default="scope">
            <el-tag
              effect="light"
              :type="statusTagMap[scope.row.status] || 'info'"
            >
              {{ statusLabelMap[scope.row.status] || defaultStatusLabel }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="错误内容"
          prop="info"
          show-overflow-tooltip
          width="240"
        />

        <el-table-column
          align="left"
          label="解决方案"
          show-overflow-tooltip
          prop="solution"
          width="120"
        />

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          :min-width="appStore.operateMinWith"
        >
          <template #default="scope">
            <el-button
              v-if="scope.row.status !== '处理中'"
              type="primary"
              link
              class="table-button"
              @click="getSolution(scope.row.ID)"
            >
              <el-icon><ai-gva /></el-icon>方案
            </el-button>
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
              ><el-icon style="margin-right: 5px"><InfoFilled /></el-icon
              >查看</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer
      destroy-on-close
      :size="appStore.drawerSize"
      v-model="detailShow"
      :show-close="true"
      :before-close="closeDetailShow"
      title="查看"
    >
      <el-descriptions :column="2" border direction="vertical">
        <el-descriptions-item label="错误来源">
          {{ detailForm.form }}
        </el-descriptions-item>
        <el-descriptions-item label="错误等级">
          <el-tag
            effect="dark"
            :type="levelTagMap[detailForm.level] || 'info'"
          >
            {{ levelLabelMap[detailForm.level] || defaultLevelLabel }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="处理状态">
          <el-tag
            effect="light"
            :type="statusTagMap[detailForm.status] || 'info'"
          >
            {{ statusLabelMap[detailForm.status] || defaultStatusLabel }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="错误内容" :span="2">
          <pre class="whitespace-pre-wrap break-words">{{ detailForm.info }}</pre>
        </el-descriptions-item>
        <el-descriptions-item label="解决方案" :span="2">
          <pre class="whitespace-pre-wrap break-words">{{ detailForm.solution }}</pre>
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    deleteSysError,
    deleteSysErrorByIds,
    findSysError,
    getSysErrorList,
    getSysErrorSolution
  } from '@/api/system/sysError'

  import { formatDate } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref } from 'vue'
  import { useAppStore } from '@/pinia'

  defineOptions({
    name: 'SysError'
  })

  const appStore = useAppStore()

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  const elSearchFormRef = ref()

  // =========== 表格控制部分 ===========
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  // 重置
  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }

  const getSolution = async (id) => {
    const confirmed = await ElMessageBox.confirm(
      '日志将通过 AI-PATH 传输至 GVA AI 用于错误分析，并在 GVA 官方平台短暂存储作为 AI 上下文。是否确认进行 AI 处理？（此功能仅向授权用户开放）',
      '提示(Beta)',
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning'
      }
    ).catch(() => false)
    if (!confirmed) return
    const res = await getSysErrorSolution({ id })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: res.msg || '处理已提交，1分钟后完成' })
      getTableData()
    }
  }
  // 搜索
  const onSubmit = () => {
    elSearchFormRef.value?.validate(async (valid) => {
      if (!valid) return
      page.value = 1
      getTableData()
    })
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  // 修改页面容量
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getSysErrorList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  getTableData()

  // ============== 表格控制部分结束 ===============

  // 获取需要的字典 可能为空 按需保留
  const setOptions = async () => {}

  // 获取需要的字典 可能为空 按需保留
  setOptions()

  // 多选数据
  const multipleSelection = ref([])
  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 删除行
  const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      deleteSysErrorFunc(row)
    })
  }

  // 多选删除
  const onDelete = async () => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map((item) => {
          IDs.push(item.ID)
        })
      const res = await deleteSysErrorByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }

  // 删除行
  const deleteSysErrorFunc = async (row) => {
    const res = await deleteSysError({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }

  const detailForm = ref({})

  // 查看详情控制标记
  const detailShow = ref(false)

  // 打开详情弹窗
  const openDetailShow = () => {
    detailShow.value = true
  }

  // 打开详情
  const getDetails = async (row) => {
    // 打开弹窗
    const res = await findSysError({ ID: row.ID })
    if (res.code === 0) {
      detailForm.value = res.data
      openDetailShow()
    }
  }

  // 关闭详情弹窗
  const closeDetailShow = () => {
    detailShow.value = false
    detailForm.value = {}
  }

  const statusLabelMap = {
    未处理: '未处理',
    处理中: '处理中',
    处理完成: '处理完成',
    处理失败: '处理失败'
  }
  const statusTagMap = {
    未处理: 'info',
    处理中: 'warning',
    处理完成: 'success',
    处理失败: 'danger'
  }
  const defaultStatusLabel = '未处理'

  const levelLabelMap = {
    fatal: '致命错误',
    error: '一般错误'
  }
  const levelTagMap = {
    fatal: 'danger',
    error: 'warning'
  }
  const defaultLevelLabel = '一般错误'
</script>
