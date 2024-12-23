<template>
  <div>
    <WarningBar
      :title="t('view.systemTools.exportTemplate.syncTableExportFeature')"
      href="https://flipped-aurora.feishu.cn/docx/KwjxdnvatozgwIxGV0rcpkZSn4d"
    />
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
              <el-tooltip :content="t('general.searchDesc')">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            :placeholder="t('general.endData')"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          />
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            :placeholder="t('general.startData')"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          />
        </el-form-item>
        <el-form-item
          :label="t('view.systemTools.exportTemplate.templateName')"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            :placeholder="t('general.searchCriteria')"
          />
        </el-form-item>
        <el-form-item
          :label="t('view.systemTools.exportTemplate.tableName')"
          prop="tableName"
        >
          <el-input
            v-model="searchInfo.tableName"
            :placeholder="t('general.searchCriteria')"
          />
        </el-form-item>
        <el-form-item
          :label="t('view.systemTools.exportTemplate.templateIdentifier')"
          prop="templateID"
        >
          <el-input
            v-model="searchInfo.templateID"
            :placeholder="t('general.searchCriteria')"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">{{
            t('general.search')
          }}</el-button>
          <el-button icon="refresh" @click="onReset">{{
            t('general.reset')
          }}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">{{
          t('general.add')
        }}</el-button>

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
        <el-table-column
          align="left"
          :label="t('general.createdAt')"
          width="180"
        >
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoCode.dbName')"
          width="240"
        >
          <template #default="scope">
            <span>{{
              scope.row.dbName ||
              t('view.systemTools.exportTemplate.gvaDatabase')
            }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.systemTools.exportTemplate.templateIdentifier')"
          prop="templateID"
          width="180"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.exportTemplate.templateName')"
          prop="name"
          width="260"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.exportTemplate.tableName')"
          prop="tableName"
          width="120"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.exportTemplate.templateInfo')"
          prop="templateInfo"
          min-width="120"
          show-overflow-tooltip
        />
        <el-table-column align="left" :label="t('components.commandMenu.operate')" min-width="280">
          <template #default="scope">
            <el-button
                type="primary"
                link
                icon="documentCopy"
                class="table-button"
                @click="copyFunc(scope.row)"
            >{{ t('view.systemTools.autoCode.copy') }}</el-button
            >
            <el-button
              type="primary"
              link
              icon="edit-pen"
              class="table-button"
              @click="showCode(scope.row)"
              >{{ t('view.systemTools.exportTemplate.code') }}</el-button
            >
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateSysExportTemplateFunc(scope.row)"
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
      v-model="dialogFormVisible"
      size="40%"
      :before-close="closeDialog"
      :title="
        type === 'create'
          ? t('view.systemTools.exportTemplate.add')
          : t('view.systemTools.exportTemplate.edit')
      "
      :show-close="false"
      destroy-on-close
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            type === 'create'
              ? t('view.systemTools.exportTemplate.add')
              : t('view.systemTools.exportTemplate.edit')
          }}</span>
          <div>
            <el-button @click="closeDialog">{{ t('general.close') }}</el-button>
            <el-button type="primary" @click="enterDialog">{{
              t('general.confirm')
            }}</el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="160px"
        v-loading="aiLoading"
        :element-loading-text="t('view.systemTools.exportTemplate.xiaoMiaoIsThinking')"
      >
        <el-form-item
          label-width="160px"
          :label="t('view.systemTools.autoCode.businessLibrary')"
          prop="dbName"
        >
          <template #label>
            <el-tooltip
              :content="t('view.systemTools.exportTemplate.dbListNote')"
              placement="bottom"
              effect="light"
            >
              <div>
                {{ t('view.systemTools.autoCode.businessLibrary') }}
                <el-icon><QuestionFilled /></el-icon>
              </div>
            </el-tooltip>
          </template>
          <el-select
            v-model="formData.dbName"
            clearable
            @change="dbNameChange"
            :placeholder="t('view.systemTools.autoCode.selectBusinessLibrary')"
          >
            <el-option
              v-for="item in dbList"
              :key="item.aliasName"
              :value="item.aliasName"
              :label="item.aliasName"
              :disabled="item.disable"
            >
              <div>
                <span>{{ item.aliasName }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{
                  item.dbName
                }}</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="t('view.systemTools.exportTemplate.tableToBeUsed')" prop="tables">
          <el-select
            multiple
            v-model="tables"
            clearable
            :placeholder="t('view.systemTools.exportTemplate.selectWhenUSingAi')"
          >
            <el-option
              v-for="item in tableOptions"
              :key="item.tableName"
              :label="item.tableName"
              :value="item.tableName"
            />
          </el-select>
        </el-form-item>

        <el-form-item :label="t('view.systemTools.exportTemplate.aiHelpWriting')" prop="ai">
          <div class="relative w-full">
            <el-input
              type="textarea"
              v-model="prompt"
              :clearable="true"
              :rows="5"
              :placeholder="t('view.systemTools.exportTemplate.aiNote')"
            />
            <el-button
              class="absolute bottom-2 right-2"
              type="primary"
              @click="autoExport"
              ><el-icon><ai-gva /></el-icon>{{ t('view.systemTools.exportTemplate.helpWrite') }}</el-button
            >
          </div>
        </el-form-item>

        <el-form-item
          :label="t('view.systemTools.exportTemplate.tableName') + ':'"
          clearable
          prop="tableName"
        >
          <div class="w-full flex gap-4">
            <el-select
              v-model="formData.tableName"
              class="flex-1"
              filterable
              :placeholder="t('view.systemTools.autoCode.selectTable')"
            >
              <el-option
                v-for="item in tableOptions"
                :key="item.tableName"
                :label="item.tableName"
                :value="item.tableName"
              />
            </el-select>
            <el-button
              :disabled="!formData.tableName"
              type="primary"
              @click="getColumnFunc(true)"
              ><el-icon><ai-gva /></el-icon>{{ t('view.systemTools.exportTemplate.autoComplete') }}</el-button
            >
            <el-button
              :disabled="!formData.tableName"
              type="primary"
              @click="getColumnFunc(false)"
              >{{ t('view.systemTools.exportTemplate.autoGenerateTemplates') }}</el-button
            >
          </div>
        </el-form-item>

        <el-form-item :label="t('view.systemTools.exportTemplate.templateName2')" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            :placeholder="t('view.systemTools.exportTemplate.templateNameNote')"
          />
        </el-form-item>

        <el-form-item
          label-width="160px"
          :label="t('view.systemTools.exportTemplate.templateIdentifier')"
          prop="templateID"
        >
          <el-input
            v-model="formData.templateID"
            :clearable="true"
            :placeholder="
              t('view.systemTools.exportTemplate.templateIdentifierInfo')
            "
          />
        </el-form-item>

        <el-form-item
          label-width="160px"
          :label="t('view.systemTools.exportTemplate.associationCondition')"
        >
          <div
            v-for="(join, key) in formData.joinTemplate"
            :key="key"
            class="flex gap-4 w-full mb-2"
          >
            <el-select
              v-model="join.joins"
              :placeholder="
                t('view.systemTools.exportTemplate.selectAssociationMethod')
              "
            >
              <el-option label="LEFT JOIN" value="LEFT JOIN" />
              <el-option label="INNER JOIN" value="INNER JOIN" />
              <el-option label="RIGHT JOIN" value="RIGHT JOIN" />
            </el-select>
            <el-input
              v-model="join.table"
              :placeholder="
                t('view.systemTools.exportTemplate.enterAssociationTable')
              "
            />
            <el-input
              v-model="join.on"
              :placeholder="
                t('view.systemTools.exportTemplate.associationCondition')
              "
            />
            <el-button
              type="danger"
              icon="delete"
              @click="() => formData.joinTemplate.splice(key, 1)"
              >{{ t('general.delete') }}</el-button
            >
          </div>
          <div class="flex justify-end w-full">
            <el-button type="primary" icon="plus" @click="addJoin">{{
              t('view.systemTools.exportTemplate.addCondition')
            }}</el-button>
          </div>
        </el-form-item>

        <el-form-item
          label-width="160px"
          :label="t('view.systemTools.exportTemplate.templateInfo')"
          prop="templateInfo"
        >
          <el-input
            v-model="formData.templateInfo"
            type="textarea"
            :rows="12"
            :clearable="true"
            :placeholder="templatePlaceholder"
          />
        </el-form-item>
        <el-form-item
          label-width="240px"
          :label="t('view.systemTools.exportTemplate.defaultExportCount')"
        >
          <el-input-number
            v-model="formData.limit"
            :step="1"
            :step-strictly="true"
            :precision="0"
          />
        </el-form-item>
        <el-form-item
          label-width="160px"
          :label="t('view.systemTools.exportTemplate.defaultSortCondition')"
        >
          <el-input
            v-model="formData.order"
            :placeholder="
              t('view.systemTools.exportTemplate.sortingCriteriaNote')
            "
          />
        </el-form-item>
        <el-form-item
          label-width="160px"
          :label="t('view.systemTools.exportTemplate.exportCondition')"
        >
          <div
            v-for="(condition, key) in formData.conditions"
            :key="key"
            class="flex gap-4 w-full mb-2"
          >
            <el-input
              v-model="condition.from"
              :placeholder="
                t('view.systemTools.exportTemplate.jsonKeyFromQuery')
              "
            />
            <el-input
              v-model="condition.column"
              :placeholder="t('view.systemTools.exportTemplate.tableColumn')"
            />
            <el-select
              v-model="condition.operator"
              :placeholder="
                t('view.systemTools.exportTemplate.selectQueryCondition')
              "
            >
              <el-option
                v-for="item in typeSearchOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
            <el-button
              type="danger"
              icon="delete"
              @click="() => formData.conditions.splice(key, 1)"
              >{{ t('general.delete') }}</el-button
            >
          </div>
          <div class="flex justify-end w-full">
            <el-button type="primary" icon="plus" @click="addCondition">{{
              t('view.systemTools.exportTemplate.addCondition')
            }}</el-button>
          </div>
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      v-model="codeVisible"
      size="60%"
      :before-close="closeDialog"
      :title="type === 'create' ? t('general.addTo') : t('general.modify')"
      :show-close="false"
      destroy-on-close
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.systemTools.autoPkg.template') }}</span>
          <div>
            <el-button type="primary" @click="closeDialog">{{ t('general.confirm') }}</el-button>
          </div>
        </div>
      </template>
      <v-ace-editor
        v-model:value="webCode"
        lang="vue"
        theme="github_dark"
        class="h-full"
      />
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createSysExportTemplate,
    deleteSysExportTemplate,
    deleteSysExportTemplateByIds,
    updateSysExportTemplate,
    findSysExportTemplate,
    getSysExportTemplateList
  } from '@/api/exportTemplate.js'

  // 全量引入格式化工具 请按需保留
  import { formatDate } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { getDB, getTable, getColumn, butler } from '@/api/autoCode'
  import { useI18n } from 'vue-i18n'
  import { getCode } from './code'
  import { VAceEditor } from 'vue3-ace-editor'

  import 'ace-builds/src-noconflict/mode-vue'
  import 'ace-builds/src-noconflict/theme-github_dark'

  const { t } = useI18n()
  defineOptions({
    name: 'ExportTemplate'
  })

  const templatePlaceholder =
    t('view.systemTools.exportTemplate.templatePlaceholder1') +
    `{
  "table_column1":"` +
    t('view.systemTools.exportTemplate.templatePlaceholder2') +
    `",
  "table_column3":"` +
    t('view.systemTools.exportTemplate.templatePlaceholder3') +
    `",
  "table_column4":"` +
    t('view.systemTools.exportTemplate.templatePlaceholder4') +
    `",
  "\`rows\`":"` +
    t('view.systemTools.exportTemplate.templatePlaceholder5') +
    `",
}` +
    t('view.systemTools.exportTemplate.templatePlaceholder6')

  // 自动化生成的字典（可能为空）以及字段
  const formData = ref({
    name: '',
    tableName: '',
    dbName: '',
    templateID: '',
    templateInfo: '',
    limit: 0,
    order: '',
    conditions: [],
    joinTemplate: []
  })

  const prompt = ref('')
  const tables = ref([])

  const typeSearchOptions = ref([
    {
      label: '=',
      value: '='
    },
    {
      label: '<>',
      value: '<>'
    },
    {
      label: '>',
      value: '>'
    },
    {
      label: '<',
      value: '<'
    },
    {
      label: 'LIKE',
      value: 'LIKE'
    },
    {
      label: 'BETWEEN',
      value: 'BETWEEN'
    },
    {
      label: 'NOT BETWEEN',
      value: 'NOT BETWEEN'
    }
  ])

  const addCondition = () => {
    formData.value.conditions.push({
      from: '',
      column: '',
      operator: ''
    })
  }

  const addJoin = () => {
    formData.value.joinTemplate.push({
      joins: 'LEFT JOIN',
      table: '',
      on: ''
    })
  }

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
    tableName: [
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
    templateID: [
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
    templateInfo: [
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

  const dbList = ref([])
  const tableOptions = ref([])
  const aiLoading = ref(false)

  const getTablesCloumn = async () => {
    const tablesMap = {}
    const promises = tables.value.map(async (item) => {
      const res = await getColumn({
        businessDB: formData.value.dbName,
        tableName: item
      })
      if (res.code === 0) {
        tablesMap[item] = res.data.columns
      }
    })
    await Promise.all(promises)
    return tablesMap
  }

  const autoExport = async () => {
    if (tables.value.length === 0) {
      ElMessage({
        type: 'error',
        message: '请先选择需要参与导出的表'
      })
      return
    }
    aiLoading.value = true
    const tableMap = await getTablesCloumn()
    const aiRes = await butler({
      prompt: prompt.value,
      businessDB: formData.value.dbName || '',
      tableMap: tableMap,
      command: 'autoExportTemplate'
    })
    aiLoading.value = false
    if (aiRes.code === 0) {
      const aiData = JSON.parse(aiRes.data)
      formData.value.name = aiData.name
      formData.value.tableName = aiData.tableName
      formData.value.templateID = aiData.templateID
      formData.value.templateInfo = JSON.stringify(aiData.templateInfo, null, 2)
      formData.value.joinTemplate = aiData.joinTemplate
    }
  }

  const getDbFunc = async () => {
    const res = await getDB()
    if (res.code === 0) {
      dbList.value = res.data.dbList
    }
  }

  getDbFunc()

  const dbNameChange = () => {
    formData.value.tableName = ''
    formData.value.templateInfo = ''
    tables.value = []
    getTableFunc()
  }

  const getTableFunc = async () => {
    const res = await getTable({ businessDB: formData.value.dbName })
    if (res.code === 0) {
      tableOptions.value = res.data.tables
    }
    formData.value.tableName = ''
  }
  getTableFunc()
  const getColumnFunc = async (aiFLag) => {
    if (!formData.value.tableName) {
      ElMessage({
        type: 'error',
        message: t('view.systemTools.selectBusinessDbAndTable')
      })
      return
    }
    formData.value.templateInfo = ''
    aiLoading.value = true
    const res = await getColumn({
      businessDB: formData.value.dbName,
      tableName: formData.value.tableName
    })
    if (res.code === 0) {
      if (aiFLag) {
        const aiRes = await butler({
          data: res.data.columns,
          command: 'exportCompletion'
        })
        if (aiRes.code === 0) {
          const aiData = JSON.parse(aiRes.data)
          aiLoading.value = false
          formData.value.templateInfo = JSON.stringify(
            aiData.templateInfo,
            null,
            2
          )
          formData.value.name = aiData.name
          formData.value.templateID = aiData.templateID
          return
        }
        ElMessage.warning(t('view.systemTools.exportTemplate.aiAutoCompleteFail'))
      }

      // 把返回值的data.columns做尊换，制作一组JSON数据，columnName做key，columnComment做value
      const templateInfo = {}
      res.data.columns.forEach((item) => {
        templateInfo[item.columnName] = item.columnComment || item.columnName
      })
      formData.value.templateInfo = JSON.stringify(templateInfo, null, 2)
    }
    aiLoading.value = false
  }

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
    const table = await getSysExportTemplateList({
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
      deleteSysExportTemplateFunc(row)
    })
  }

  // 多选删除
  const onDelete = async () => {
    ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    }).then(async () => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: t('general.selectDataToDelete')
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map((item) => {
          ids.push(item.ID)
        })
      const res = await deleteSysExportTemplateByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }

  // 行为控制标记（弹窗内部需要增还是改）
  const type = ref('')

  // 复制
  const copyFunc = async (row) => {
    let copyData
    const res = await findSysExportTemplate({ ID: row.ID })
    if (res.code === 0) {
      copyData = JSON.parse(JSON.stringify(res.data.resysExportTemplate))
      if (!copyData.conditions) {
        copyData.conditions = []
      }
      if (!copyData.joinTemplate) {
        copyData.joinTemplate = []
      }
      delete copyData.ID
      delete copyData.CreatedAt
      delete copyData.UpdatedAt
      copyData.templateID = copyData.templateID + '_copy'
      copyData.name = copyData.name + '_copy'
      formData.value = copyData
      dialogFormVisible.value = true
    }
  }

  // 更新行
  const updateSysExportTemplateFunc = async (row) => {
    const res = await findSysExportTemplate({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data.resysExportTemplate
      if (!formData.value.conditions) {
        formData.value.conditions = []
      }
      if (!formData.value.joinTemplate) {
        formData.value.joinTemplate = []
      }
      dialogFormVisible.value = true
    }
  }

  // 删除行
  const deleteSysExportTemplateFunc = async (row) => {
    const res = await deleteSysExportTemplate({ ID: row.ID })
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
  const codeVisible = ref(false)
  // 弹窗控制标记
  const dialogFormVisible = ref(false)

  const webCode = ref('')

  const showCode = (row) => {
    webCode.value = getCode(row.templateID)
    codeVisible.value = true
  }

  // 打开弹窗
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 关闭弹窗
  const closeDialog = () => {
    codeVisible.value = false
    dialogFormVisible.value = false
    formData.value = {
      name: '',
      tableName: '',
      templateID: '',
      templateInfo: '',
      limit: 0,
      order: '',
      conditions: [],
      joinTemplate: []
    }
  }
  // 弹窗确定
  const enterDialog = async () => {
    // 判断 formData.templateInfo 是否为标准json格式 如果不是标准json 则辅助调整
    try {
      JSON.parse(formData.value.templateInfo)
    } catch (_) {
      ElMessage({
        type: 'error',
        message: t('view.systemTools.templateFormatIncorrect')
      })
      return
    }

    const reqData = JSON.parse(JSON.stringify(formData.value))
    for (let i = 0; i < reqData.conditions.length; i++) {
      if (
        !reqData.conditions[i].from ||
        !reqData.conditions[i].column ||
        !reqData.conditions[i].operator
      ) {
        ElMessage({
          type: 'error',
          message: t('view.systemTools.completeExportConditions')
        })
        return
      }
      reqData.conditions[i].templateID = reqData.templateID
    }

    for (let i = 0; i < reqData.joinTemplate.length; i++) {
      if (!reqData.joinTemplate[i].joins || !reqData.joinTemplate[i].on) {
        ElMessage({
          type: 'error',
          message: t('view.systemTools.completeAssociation')
        })
        return
      }
      reqData.joinTemplate[i].templateID = reqData.templateID
    }

    elFormRef.value?.validate(async (valid) => {
      if (!valid) return
      let res
      switch (type.value) {
        case 'create':
          res = await createSysExportTemplate(reqData)
          break
        case 'update':
          res = await updateSysExportTemplate(reqData)
          break
        default:
          res = await createSysExportTemplate(reqData)
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
</script>

<style></style>
