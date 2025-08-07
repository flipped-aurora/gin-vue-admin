<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item :label="t('view.api.path')">
          <el-input
            v-model="searchInfo.path"
            :placeholder="t('view.api.path')"
          />
        </el-form-item>
        <el-form-item :label="t('general.description')">
          <el-input
            v-model="searchInfo.description"
            :placeholder="t('general.description')"
          />
        </el-form-item>
        <el-form-item :label="t('view.api.apiGroup')">
          <el-select
            v-model="searchInfo.apiGroup"
            clearable
            :placeholder="t('view.api.apiGroup')"
          >
            <el-option
              v-for="item in apiGroupOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('general.request')">
          <el-select
            v-model="searchInfo.method"
            clearable
            :placeholder="t('general.pleaseSelect')"
          >
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            {{ t('general.search') }}
          </el-button>
          <el-button icon="refresh" @click="onReset">
            {{ t('general.reset') }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog('addApi')">
          {{ t('general.add') }}
        </el-button>
        <el-button icon="delete" :disabled="!apis.length" @click="onDelete">
          {{ t('general.delete') }}
        </el-button>
        <el-button icon="Refresh" @click="onFresh">
          {{ t('view.api.refreshCache') }}
        </el-button>
        <el-button icon="Compass" @click="onSync">
          {{ t('view.api.synchronousAPI') }}
        </el-button>
        <ExportTemplate template-id="api" />
        <ExportExcel template-id="api" :limit="9999" />
        <ImportExcel template-id="api" @on-success="getTableData" />
      </div>
      <el-table
        :data="tableData"
        @sort-change="sortChange"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column
          align="left"
          label="ID"
          min-width="60"
          prop="ID"
          sortable="custom"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiPath')"
          min-width="150"
          prop="path"
          sortable="custom"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiGrouping')"
          min-width="150"
          prop="apiGroup"
          sortable="custom"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiDescription')"
          min-width="150"
          prop="description"
          sortable="custom"
        />
        <el-table-column
          align="left"
          :label="t('general.request')"
          min-width="150"
          prop="method"
          sortable="custom"
        >
          <template #default="scope">
            <div>
              {{ scope.row.method }} / {{ methodFilter(scope.row.method) }}
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" fixed="right" :label="t('general.operations')" :min-width="appStore.operateMinWith" width="180px">
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editApiFunc(scope.row)"
            >
              {{ t('general.edit') }}
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >
              {{ t('general.delete') }}
            </el-button>
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

    <el-drawer
      v-model="syncApiFlag"
      :size="appStore.drawerSize"
      :before-close="closeSyncDialog"
      :show-close="false"
    >
      <warning-bar :title="t('view.api.synchronousAPINote')" />
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg"> {{ t('view.api.synchronousRouting') }} </span>
          <div>
            <el-button :loading="apiCompletionLoading" @click="closeSyncDialog">
              {{ t('general.close') }}
            </el-button>
            <el-button
              type="primary"
              :loading="syncing || apiCompletionLoading"
              @click="enterSyncDialog"
            >
              {{ t('general.confirm') }}
            </el-button>
          </div>
        </div>
      </template>

      <h4>
        {{ t('view.api.newAddedRouteNote1') }}
        <span class="text-xs text-gray-500 mx-2 font-normal">{{
          t('view.api.newAddedRouteNote2')
        }}</span>
        <el-button type="primary" size="small" @click="apiCompletion">
          <el-icon size="18">
            <ai-gva />
          </el-icon>
          自动填充
        </el-button>
      </h4>
      <el-table
        v-loading="syncing || apiCompletionLoading"
        element-loading-text="小淼正在思考..."
        :data="syncApiData.newApis"
      >
        <el-table-column
          align="left"
          :label="t('view.api.apiPath')"
          min-width="150"
          prop="path"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiGrouping')"
          min-width="150"
          prop="apiGroup"
        >
          <template #default="{ row }">
            <el-select
              v-model="row.apiGroup"
              :placeholder="t('view.api.selectOrAdd')"
              allow-create
              filterable
              default-first-option
            >
              <el-option
                v-for="item in apiGroupOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.api.apiDescription')"
          min-width="150"
          prop="description"
        >
          <template #default="{ row }">
            <el-input v-model="row.description" autocomplete="off" />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('general.request')"
          min-width="150"
          prop="method"
        >
          <template #default="scope">
            <div>
              {{ scope.row.method }} / {{ methodFilter(scope.row.method) }}
            </div>
          </template>
        </el-table-column>
        <el-table-column
          :label="t('general.operations')"
          min-width="150"
          fixed="right"
        >
          <template #default="{ row }">
            <el-button icon="plus" type="primary" link @click="addApiFunc(row)">
              单条新增
            </el-button>
            <el-button
              icon="sunrise"
              type="primary"
              link
              @click="ignoreApiFunc(row, true)"
            >
              {{ t('view.api.ignore') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <h4>
        {{ t('view.api.deletedRouteNote1') }}
        <span class="text-xs text-gray-500 ml-2 font-normal">{{
          t('view.api.deletedRouteNote2')
        }}</span>
      </h4>
      <el-table :data="syncApiData.deleteApis">
        <el-table-column
          align="left"
          :label="t('view.api.apiPath')"
          min-width="150"
          prop="path"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiGrouping')"
          min-width="150"
          prop="apiGroup"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiDescription')"
          min-width="150"
          prop="description"
        />
        <el-table-column
          align="left"
          :label="t('general.request')"
          min-width="150"
          prop="method"
        >
          <template #default="scope">
            <div>
              {{ scope.row.method }} / {{ methodFilter(scope.row.method) }}
            </div>
          </template>
        </el-table-column>
      </el-table>

      <h4>
        {{ t('view.api.ignoreRouteNote1') }}
        <span class="text-xs text-gray-500 ml-2 font-normal">{{
          t('view.api.ignoreRouteNote2')
        }}</span>
      </h4>
      <el-table :data="syncApiData.ignoreApis">
        <el-table-column
          align="left"
          :label="t('view.api.apiPath')"
          min-width="150"
          prop="path"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiGrouping')"
          min-width="150"
          prop="apiGroup"
        />
        <el-table-column
          align="left"
          :label="t('view.api.apiDescription')"
          min-width="150"
          prop="description"
        />
        <el-table-column
          align="left"
          :label="t('general.request')"
          min-width="150"
          prop="method"
        >
          <template #default="scope">
            <div>
              {{ scope.row.method }} / {{ methodFilter(scope.row.method) }}
            </div>
          </template>
        </el-table-column>
        <el-table-column
          :label="t('general.operation')"
          min-width="150"
          fixed="right"
        >
          <template #default="{ row }">
            <el-button
              icon="sunny"
              type="primary"
              link
              @click="ignoreApiFunc(row, false)"
            >
              {{ t('view.api.cancelIgnore') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-drawer>

    <el-drawer
      v-model="dialogFormVisible"
      :size="appStore.drawerSize"
      :before-close="closeDialog"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ dialogTitle }}</span>
          <div>
            <el-button @click="closeDialog">
              {{ t('general.close') }}
            </el-button>
            <el-button type="primary" @click="enterDialog">
              {{ t('general.confirm') }}
            </el-button>
          </div>
        </div>
      </template>

      <warning-bar :title="t('view.api.newApiNote')" />
      <el-form ref="apiForm" :model="form" :rules="rules" label-width="120px">
        <el-form-item :label="t('view.api.path')" prop="path">
          <el-input v-model="form.path" autocomplete="off" />
        </el-form-item>
        <el-form-item :label="t('general.request')" prop="method">
          <el-select
            v-model="form.method"
            :placeholder="t('general.pleaseSelect')"
            style="width: 100%"
          >
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('view.api.apiGrouping')" prop="apiGroup">
          <el-select
            v-model="form.apiGroup"
            :placeholder="t('view.api.selectOrAdd')"
            allow-create
            filterable
            default-first-option
          >
            <el-option
              v-for="item in apiGroupOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('view.api.apiDescription')" prop="description">
          <el-input v-model="form.description" autocomplete="off" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    getApiById,
    getApiList,
    createApi,
    updateApi,
    deleteApi,
    deleteApisByIds,
    freshCasbin,
    syncApi,
    getApiGroups,
    ignoreApi,
    enterSyncApi
  } from '@/api/api'
  import { toSQLLine } from '@/utils/stringFun'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import ExportExcel from '@/components/exportExcel/exportExcel.vue'
  import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'
  import ImportExcel from '@/components/exportExcel/importExcel.vue'
  import { butler } from '@/api/autoCode'
  import { useAppStore } from "@/pinia";
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

  const { t } = useI18n() // added by mohamed hassan to support multilingual

  defineOptions({
    name: 'Api'
  })

  const appStore = useAppStore()

  const methodFilter = (value) => {
    const target = methodOptions.value.filter((item) => item.value === value)[0]
    return target && `${target.label}`
  }

  const apis = ref([])
  const form = ref({
    path: '',
    apiGroup: '',
    method: '',
    description: ''
  })
  const methodOptions = ref([
    {
      value: 'POST',
      label: t('view.api.create'),
      type: 'success'
    },
    {
      value: 'GET',
      label: t('view.api.view'),
      type: ''
    },
    {
      value: 'PUT',
      label: t('view.api.update'),
      type: 'warning'
    },
    {
      value: 'DELETE',
      label: t('general.delete'),
      type: 'danger'
    }
  ])

  const type = ref('')
  const rules = ref({
    path: [
      { required: true, message: t('view.api.enterApiPath'), trigger: 'blur' }
    ],
    apiGroup: [
      { required: true, message: t('view.api.enterGroupName'), trigger: 'blur' }
    ],
    method: [
      {
        required: true,
        message: t('view.api.selectRequestMethod'),
        trigger: 'blur'
      }
    ],
    description: [
      {
        required: true,
        message: t('view.api.enterApiDescription'),
        trigger: 'blur'
      }
    ]
  })

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  const apiGroupOptions = ref([])
  const apiGroupMap = ref({})

  const getGroup = async () => {
    const res = await getApiGroups()
    if (res.code === 0) {
      const groups = res.data.groups
      apiGroupOptions.value = groups.map((item) => ({
        label: item,
        value: item
      }))
      apiGroupMap.value = res.data.apiGroupMap
    }
  }

  const ignoreApiFunc = async (row, flag) => {
    const res = await ignoreApi({ path: row.path, method: row.method, flag })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg
      })
      if (flag) {
        syncApiData.value.newApis = syncApiData.value.newApis.filter(
          (item) => !(item.path === row.path && item.method === row.method)
        )
        syncApiData.value.ignoreApis.push(row)
        return
      }
      syncApiData.value.ignoreApis = syncApiData.value.ignoreApis.filter(
        (item) => !(item.path === row.path && item.method === row.method)
      )
      syncApiData.value.newApis.push(row)
    }
  }

  const addApiFunc = async (row) => {
    if (!row.apiGroup) {
      ElMessage({
        type: 'error',
        message: '请先选择API分组'
      })
      return
    }
    if (!row.description) {
      ElMessage({
        type: 'error',
        message: '请先填写API描述'
      })
      return
    }
    const res = await createApi(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '添加成功，请到角色管理页面分配权限',
        showClose: true
      })
      syncApiData.value.newApis = syncApiData.value.newApis.filter(
        (item) => !(item.path === row.path && item.method === row.method)
      )
    }
    getTableData()
    getGroup()
  }

  const closeSyncDialog = () => {
    syncApiFlag.value = false
  }

  const syncing = ref(false)

  const enterSyncDialog = async () => {
    if (
      syncApiData.value.newApis.some(
        (item) => !item.apiGroup || !item.description
      )
    ) {
      ElMessage({
        type: 'error',
        message: '存在API未分组或未填写描述'
      })
      return
    }

    syncing.value = true
    const res = await enterSyncApi(syncApiData.value)
    syncing.value = false
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg
      })
      syncApiFlag.value = false
      getTableData()
    }
  }

  const onReset = () => {
    searchInfo.value = {}
    getTableData()
  }
  // 搜索

  const onSubmit = () => {
    page.value = 1
    getTableData()
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 排序
  const sortChange = ({ prop, order }) => {
    if (prop) {
      if (prop === 'ID') {
        prop = 'id'
      }
      searchInfo.value.orderKey = toSQLLine(prop)
      searchInfo.value.desc = order === 'descending'
    }
    getTableData()
  }

  // 查询
  const getTableData = async () => {
    const table = await getApiList({
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
  getGroup()
  // 批量操作
  const handleSelectionChange = (val) => {
    apis.value = val
  }

  const onDelete = async () => {
    ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    }).then(async () => {
      const ids = apis.value.map((item) => item.ID)
      const res = await deleteApisByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: res.msg
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
  }
  const onFresh = async () => {
    ElMessageBox.confirm(t('view.api.cacheConfirmNote'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    }).then(async () => {
      const res = await freshCasbin()
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: res.msg
        })
      }
    })
  }

  const syncApiData = ref({
    newApis: [],
    deleteApis: [],
    ignoreApis: []
  })

  const syncApiFlag = ref(false)

  const onSync = async () => {
    const res = await syncApi()
    if (res.code === 0) {
      res.data.newApis.forEach((item) => {
        item.apiGroup = apiGroupMap.value[item.path.split('/')[1]]
      })

      syncApiData.value = res.data
      syncApiFlag.value = true
    }
  }

  // 弹窗相关
  const apiForm = ref(null)
  const initForm = () => {
    apiForm.value.resetFields()
    form.value = {
      path: '',
      apiGroup: '',
      method: '',
      description: ''
    }
  }

  const dialogTitle = ref(t('view.api.newApi'))
  const dialogFormVisible = ref(false)
  const openDialog = (key) => {
    switch (key) {
      case 'addApi':
        dialogTitle.value = t('view.api.newApi')
        break
      case 'edit':
        dialogTitle.value = t('view.api.editApi')
        break
      default:
        break
    }
    type.value = key
    dialogFormVisible.value = true
  }
  const closeDialog = () => {
    initForm()
    dialogFormVisible.value = false
  }

  const editApiFunc = async (row) => {
    const res = await getApiById({ id: row.ID })
    form.value = res.data.api
    openDialog('edit')
  }

  const enterDialog = async () => {
    apiForm.value.validate(async (valid) => {
      if (valid) {
        switch (type.value) {
          case 'addApi':
            {
              const res = await createApi(form.value)
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: t('general.addSuccess'),
                  showClose: true
                })
              }
              getTableData()
              getGroup()
              closeDialog()
            }

            break
          case 'edit':
            {
              const res = await updateApi(form.value)
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: t('general.editSuccess'),
                  showClose: true
                })
              }
              getTableData()
              closeDialog()
            }
            break
          default:
            {
              ElMessage({
                type: 'error',
                message: t('view.api.unknownOperation'),
                showClose: true
              })
            }
            break
        }
      }
    })
  }

  const deleteApiFunc = async (row) => {
    ElMessageBox.confirm(t('view.api.deleteApiConfirm'), t('general.hint'), {
      confirmButtonText: t('general.confirm'),
      cancelButtonText: t('general.cancel'),
      type: 'warning'
    }).then(async () => {
      const res = await deleteApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
        getGroup()
      }
    })
  }
  const apiCompletionLoading = ref(false)
  const apiCompletion = async () => {
    apiCompletionLoading.value = true
    const routerPaths = syncApiData.value.newApis
      .filter((item) => !item.apiGroup || !item.description)
      .map((item) => item.path)
    const res = await butler({ data: routerPaths, command: 'apiCompletion' })
    apiCompletionLoading.value = false
    if (res.code === 0) {
      try {
        const data = JSON.parse(res.data)
        syncApiData.value.newApis.forEach((item) => {
          const target = data.find((d) => d.path === item.path)
          if (target) {
            if (!item.apiGroup) {
              item.apiGroup = target.apiGroup
            }
            if (!item.description) {
              item.description = target.description
            }
          }
        })
      } catch (_) {
        ElMessage({
          type: 'error',
          message: 'AI自动填充失败,请重新生成'
        })
      }
    }
  }
</script>

<style scoped lang="scss">
  .warning {
    color: #dc143c;
  }
</style>
