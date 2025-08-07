<template>
  <div>
    <warning-bar :title="t('view.superAdmin.params.paramNote')" />
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item :label="t('general.createDate')" prop="createdAt">
          <template #label>
            <span>
              {{ t('general.createDate') }}
              <el-tooltip
                :content="t('general.searchDesc')"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            :placeholder="t('general.startData')"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            :placeholder="t('general.endDate')"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>

        <el-form-item :label="t('view.superAdmin.params.paramName')" prop="name">
          <el-input v-model="searchInfo.name" :placeholder="t('general.searchCriteria')" />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.params.paramKey')" prop="key">
          <el-input v-model="searchInfo.key" :placeholder="t('general.searchCriteria')" />
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >{{ t('general.search') }}</el-button
          >
          <el-button icon="refresh" @click="onReset">{{ t('general.reset') }}</el-button>
          <el-button
            link
            type="primary"
            icon="arrow-down"
            @click="showAllQuery = true"
            v-if="!showAllQuery"
            >{{ t('general.expand') }}</el-button
          >
          <el-button
            link
            type="primary"
            icon="arrow-up"
            @click="showAllQuery = false"
            v-else
            >{{ t('general.collapse') }}</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >{{ t('general.add') }}</el-button
        >
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >{{ t('general.delete') }}</el-button
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

        <el-table-column align="left" :label="t('general.createdAt')" prop="createdAt" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          :label="t('view.superAdmin.params.paramName')"
          prop="name"
          width="200"
        />
        <el-table-column align="left" :label="t('view.superAdmin.params.paramKey')" prop="key" width="200" />
        <el-table-column align="left" :label="t('view.superAdmin.params.paramValue')" prop="value" width="200" />
        <el-table-column
          align="left"
          :label="t('view.superAdmin.params.paramDesc')"
          prop="desc"
          width="200"
        />
        <el-table-column
          align="left"
          :label="t('general.operations')"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
              ><el-icon style="margin-right: 5px"><InfoFilled /></el-icon
              >{{ t('general.desc') }}</el-button
            >
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateSysParamsFunc(scope.row)"
              >{{ t('general.change') }}</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >{{ t('general.delete') }}</el-button
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
      size="800"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? t('general.addTo') : t('general.modify') }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">{{ t('general.confirm') }}</el-button>
            <el-button @click="closeDialog">{{ t('general.cancel') }}</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item :label="t('view.superAdmin.params.paramName') + ':'" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            :placeholder="t('view.superAdmin.params.enterParamName')"
          />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.params.paramKey') + ':'" prop="key">
          <el-input
            v-model="formData.key"
            :clearable="true"
            :placeholder="t('view.superAdmin.params.enterParamKey')"
          />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.params.paramValue') + ':'" prop="value">
          <el-input
            type="textarea"
            :rows="5"
            v-model="formData.value"
            :clearable="true"
            :placeholder="t('view.superAdmin.params.enterParamValue')"
          />
        </el-form-item>
        <el-form-item :label="t('view.superAdmin.params.paramDesc') + ':'" prop="desc">
          <el-input
            v-model="formData.desc"
            :clearable="true"
            :placeholder="t('view.superAdmin.params.enterParamDesc')"
          />
        </el-form-item>
      </el-form>

      <div
        class="usage-instructions bg-gray-100 border border-gray-300 rounded-lg p-4 mt-5"
      >
        <h3 class="mb-3 text-lg text-gray-800">{{ t('view.superAdmin.params.instruction') }}</h3>
        <p class="mb-2 text-sm text-gray-600">
          {{ t('view.superAdmin.params.instructionNote1') }}
          <code class="bg-blue-100 px-1 py-0.5 rounded"
            >import { getParams } from '@/utils/params'</code
          >
          {{ t('view.superAdmin.params.instructionNote2') }}
          <code class="bg-blue-100 px-1 py-0.5 rounded"
            >await getParams("{{ formData.key }}")</code
          >
          {{ t('view.superAdmin.params.instructionNote3') }}
        </p>
        <p class="text-sm text-gray-600">
          {{ t('view.superAdmin.params.instructionNote4') }}
          <code class="bg-blue-100 px-1 py-0.5 rounded"
            >import
            "github.com/flipped-aurora/gin-vue-admin/server/service/system"</code
          >
        </p>
        <p class="mb-2 text-sm text-gray-600">
          {{ t('view.superAdmin.params.instructionNote5') }}
          <code class="bg-blue-100 px-1 py-0.5 rounded"
            >new(system.SysParamsService).GetSysParam("{{
              formData.key
            }}")</code
          >
          {{ t('view.superAdmin.params.instructionNote6') }}
        </p>
      </div>
    </el-drawer>

    <el-drawer
      destroy-on-close
      size="800"
      v-model="detailShow"
      :show-close="true"
      :before-close="closeDetailShow"
    >
      <el-descriptions :column="1" border>
        <el-descriptions-item :label="t('view.superAdmin.params.paramName')">
          {{ detailFrom.name }}
        </el-descriptions-item>
        <el-descriptions-item :label="t('view.superAdmin.params.paramKey')">
          {{ detailFrom.key }}
        </el-descriptions-item>
        <el-descriptions-item :label="t('view.superAdmin.params.paramValue')">
          {{ detailFrom.value }}
        </el-descriptions-item>
        <el-descriptions-item :label="t('view.superAdmin.params.paramDesc')">
          {{ detailFrom.desc }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createSysParams,
    deleteSysParams,
    deleteSysParamsByIds,
    updateSysParams,
    findSysParams,
    getSysParamsList
  } from '@/api/sysParams'

  // 全量引入格式化工具 请按需保留
  import { formatDate } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'
  import WarningBar from "@/components/warningBar/warningBar.vue";
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

  const { t } = useI18n() // added by mohamed hassan to support multilingual

  defineOptions({
    name: 'SysParams'
  })

  // 控制更多查询条件显示/隐藏状态
  const showAllQuery = ref(false)

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    name: '',
    key: '',
    value: '',
    desc: ''
  })

  // 验证规则
  const rule = reactive({
    name: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      },
      {
        whitespace: true,
        message: t('general.noOnlySpace'),
        trigger: ['input', 'blur']
      }
    ],
    key: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      },
      {
        whitespace: true,
        message: t('general.noOnlySpace'),
        trigger: ['input', 'blur']
      }
    ],
    value: [
      {
        required: true,
        message: '',
        trigger: ['input', 'blur']
      },
      {
        whitespace: true,
        message: t('general.noOnlySpace'),
        trigger: ['input', 'blur']
      }
    ]
  })

  const searchRule = reactive({
    createdAt: [
      {
        validator: (rule, value, callback) => {
          if (
            searchInfo.value.startCreatedAt &&
            !searchInfo.value.endCreatedAt
          ) {
            callback(new Error(t('general.placeInputEndData')))
          } else if (
            !searchInfo.value.startCreatedAt &&
            searchInfo.value.endCreatedAt
          ) {
            callback(new Error(t('general.placeInputStartData')))
          } else if (
            searchInfo.value.startCreatedAt &&
            searchInfo.value.endCreatedAt &&
            (searchInfo.value.startCreatedAt.getTime() ===
              searchInfo.value.endCreatedAt.getTime() ||
              searchInfo.value.startCreatedAt.getTime() >
                searchInfo.value.endCreatedAt.getTime())
          ) {
            callback(new Error(t('general.startDataMustBeforeEndData')))
          } else {
            callback()
          }
        },
        trigger: 'change'
      }
    ]
  })

  const elFormRef = ref()
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
    const table = await getSysParamsList({
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
      ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    }).then(() => {
      deleteSysParamsFunc(row)
    })
  }

  // 多选删除
  const onDelete = async () => {
      ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    }).then(async () => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: t('general.selectDataToDelete')
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map((item) => {
          IDs.push(item.ID)
        })
      const res = await deleteSysParamsByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 更新行
  const updateSysParamsFunc = async (row) => {
    const res = await findSysParams({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteSysParamsFunc = async (row) => {
    const res = await deleteSysParams({ ID: row.ID })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: t('general.deleteSuccess')
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  }

  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
      name: '',
      key: '',
      value: '',
      desc: ''
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return
      let res
      switch (type.value) {
        case 'create':
          res = await createSysParams(formData.value)
          break
        case 'update':
          res = await updateSysParams(formData.value)
          break
        default:
          res = await createSysParams(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.createUpdateSuccess')
        })
        closeDialog()
        getTableData()
      }
    })
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
    const res = await findSysParams({ ID: row.ID })
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
</script>

<style></style>
