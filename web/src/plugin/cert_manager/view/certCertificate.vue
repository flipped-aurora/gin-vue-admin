<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="flex flex-wrap items-center gap-2"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="项目名称" prop="category">
          <categorySelector
            v-model="searchInfo.category"
            placeholder="请选择项目"
            :auto-select-first="true"
            @change="onSubmit"
          />
        </el-form-item>

        <el-form-item label="域名" prop="domain">
          <el-input
            v-model="searchInfo.domain"
            placeholder="搜索域名"
            clearable
          />
        </el-form-item>

        <el-form-item label="域名状态" prop="domainStatus">
          <el-select
            v-model="searchInfo.domainStatus"
            placeholder="域名状态"
            clearable
            class="w-32"
          >
            <el-option label="正常" :value="1" />
            <el-option label="异常" :value="3" />
          </el-select>
        </el-form-item>

        <el-form-item label="证书状态" prop="certStatus">
          <el-select
            v-model="searchInfo.certStatus"
            placeholder="证书状态"
            clearable
            class="w-32"
          >
            <el-option label="正常" :value="1" />
            <el-option label="异常" :value="3" />
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">
          新增
        </el-button>
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
        >
          删除
        </el-button>
        <el-button
          type="success"
          icon="refresh"
          style="margin-left: 10px"
          :loading="updateAllLoading"
          @click="updateAllCertificatesFunc"
        >
          批量更新
        </el-button>
        <el-button
          type="primary"
          icon="edit"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onBatchUpdateCategory"
        >
          批量修改分类
        </el-button>
        <el-button
          type="warning"
          icon="refresh"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onBatchReprobe"
        >
          批量重新探测
        </el-button>
        <el-button
          type="info"
          icon="close"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onBatchIgnore"
        >
          批量忽略
        </el-button>
      </div>

      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        :row-class-name="getRowClassName"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="域名" prop="domain" width="200" />

        <el-table-column align="left" label="项目名称" prop="category" width="120" />

        <el-table-column align="left" label="证书品牌" prop="brand" width="150" />

        <el-table-column align="left" label="证书有效期" min-width="200">
          <template #default="scope">
            <div v-if="scope.row.expireAt && scope.row.expireAt !== '0001-01-01T00:00:00Z'" class="flex flex-col gap-1">
              <div class="text-xs text-gray-600">
                {{ formatYmd(scope.row.startAt) }} ~ {{ formatYmd(scope.row.expireAt) }}
              </div>
            </div>
            <div v-else class="text-xs text-gray-400 italic">
              未监测到 HTTPS 证书
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" label="证书剩余天数" width="120">
          <template #default="scope">
            <el-tag
              v-if="scope.row.expireAt && scope.row.expireAt !== '0001-01-01T00:00:00Z'"
              :type="getStatusTagType(scope.row.certStatus)"
              effect="dark"
            >
              {{ calcDays(scope.row.expireAt) }} 天
            </el-tag>
            <span v-else class="text-xs text-gray-400">-</span>
          </template>
        </el-table-column>

        <el-table-column align="left" label="域名有效期" min-width="200">
          <template #default="scope">
            <div v-if="scope.row.domainExpireAt && scope.row.domainExpireAt !== '0001-01-01T00:00:00Z'" class="flex flex-col gap-1">
              <div class="text-xs text-gray-600">
                {{ formatYmd(scope.row.domainStartAt) }} ~ {{ formatYmd(scope.row.domainExpireAt) }}
              </div>
            </div>
            <div v-else class="text-xs text-gray-400">
              暂无数据
            </div>
          </template>
        </el-table-column>

        <el-table-column align="left" label="TLS版本" prop="tlsVersion" width="120" />

        <el-table-column align="left" label="域名状态" prop="domainStatus" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.domainStatus)">
              {{ getStatusText(scope.row.domainStatus) }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column align="left" label="证书状态" prop="certStatus" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.certStatus)">
              {{ getStatusText(scope.row.certStatus) }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="refresh"
              class="table-button"
              :loading="scope.row.probeLoading"
              @click="probeCertificateFunc(scope.row)"
            >
              探测
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateCertFunc(scope.row)"
            >
              变更
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >
              删除
            </el-button>
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
      size="800"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg font-bold">{{
            type === "create" ? "新增证书监控" : "编辑证书监控"
          }}</span>
          <div>
            <el-button @click="closeDialog">取 消</el-button>
            <el-button
              type="primary"
              @click="enterDialog"
            >确 定</el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="top"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="项目名称:" prop="category">
          <categorySelector
            v-model="formData.category"
            placeholder="请选择项目"
          />
        </el-form-item>
        <el-form-item label="域名:" prop="domain">
          <el-input
            v-model="formData.domain"
            :clearable="true"
            placeholder="请输入域名（如：example.com）"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-dialog v-model="batchCategoryVisible" title="批量修改项目" width="400px">
      <el-form :model="batchCategoryData" label-width="100px">
        <el-form-item label="项目名称">
          <categorySelector v-model="batchCategoryData.category" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="batchCategoryVisible = false">取 消</el-button>
          <el-button type="primary" @click="enterBatchUpdateCategory">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCertCertificate,
  deleteCertCertificate,
  deleteCertCertificateByIds,
  updateCertCertificate,
  findCertCertificate,
  getCertCertificateList,
  probeCertificate,
  updateAllCertificates
} from '@/plugin/cert_manager/api/certCertificate'
import { batchUpdateDomainCategory } from '@/plugin/cert_manager/api/certCategory'
import {
  batchReprobe,
  batchIgnore
} from '@/plugin/cert_manager/api/certAdvanced'
import categorySelector from '../components/categorySelector.vue'
import { formatTimeToStr } from '@/utils/date'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'CertCertificate'
})

const formData = ref({
  domain: '',
  category: ''
})

const rule = reactive({
  domain: [
    { required: true, message: '请输入域名', trigger: 'blur' },
    {
      pattern: /^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](?:\.[a-zA-Z0-9]{2,})+$/,
      message: '请输入有效的域名格式',
      trigger: 'blur'
    }
  ]
})

const elFormRef = ref()
const elSearchFormRef = ref()

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const updateAllLoading = ref(false)

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async () => {
  const table = await getCertCertificateList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list.map(item => ({
      ...item,
      probeLoading: false
    }))
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const multipleSelection = ref([])

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const batchCategoryVisible = ref(false)
const batchCategoryData = ref({
  category: ''
})

const onBatchUpdateCategory = () => {
  batchCategoryData.value.category = ''
  batchCategoryVisible.value = true
}

const onBatchReprobe = async () => {
  ElMessageBox.confirm(`确定要重新探测选中的 ${multipleSelection.value.length} 个域名吗?`, '二次确认', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = multipleSelection.value.map(item => item.ID)
    const res = await batchReprobe({ ids })
    if (res.code === 0) {
      ElMessage.success('批量探测完成')
      getTableData()
    }
  })
}

const onBatchIgnore = async () => {
  ElMessageBox.confirm(`确定要忽略选中的 ${multipleSelection.value.length} 个域名吗?`, '二次确认', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = multipleSelection.value.map(item => item.ID)
    const res = await batchIgnore({ ids })
    if (res.code === 0) {
      ElMessage.success('批量忽略成功')
      getTableData()
    }
  })
}

const enterBatchUpdateCategory = async () => {
  if (!batchCategoryData.value.category) {
    ElMessage.warning('请选择项目')
    return
  }
  ElMessageBox.confirm(`确定要将选中的 ${multipleSelection.value.length} 个域名修改为 [${batchCategoryData.value.category}] 吗?`, '二次确认', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = multipleSelection.value.map(item => item.ID)
    const res = await batchUpdateDomainCategory({
      ids,
      category: batchCategoryData.value.category
    })
    if (res.code === 0) {
      ElMessage.success('批量修改成功')
      batchCategoryVisible.value = false
      getTableData()
    }
  })
}

const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteCertFunc(row)
  })
}

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
    const res = await deleteCertCertificateByIds({ IDs })
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

const type = ref('')

const updateCertFunc = async (row) => {
  const res = await findCertCertificate({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

const deleteCertFunc = async (row) => {
  const res = await deleteCertCertificate({ ID: row.ID })
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

const dialogFormVisible = ref(false)

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    domain: '',
    category: ''
  }
}

const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCertCertificate(formData.value)
        break
      case 'update':
        res = await updateCertCertificate(formData.value)
        break
      default:
        res = await createCertCertificate(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

const probeCertificateFunc = async (row) => {
  row.probeLoading = true
  try {
    const res = await probeCertificate({ domain: row.domain })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '探测成功'
      })
      getTableData()
    }
  } finally {
    row.probeLoading = false
  }
}

const updateAllCertificatesFunc = async () => {
  updateAllLoading.value = true
  try {
    const res = await updateAllCertificates()
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '批量更新成功'
      })
      getTableData()
    }
  } finally {
    updateAllLoading.value = false
  }
}

const getStatusTagType = (status) => {
  switch (status) {
    case 1: return 'success'
    case 2: return 'warning'
    case 3: return 'danger'
    default: return 'info'
  }
}

const getStatusText = (status) => {
  switch (status) {
    case 1: return '正常'
    case 2: return '告警'
    case 3: return '异常'
    default: return '未知'
  }
}

const calcDays = (expireAt) => {
  if (!expireAt || expireAt === '0001-01-01T00:00:00Z') return '-'
  const end = new Date(expireAt)
  const now = new Date()
  const diff = end.getTime() - now.getTime()
  return Math.ceil(diff / (1000 * 60 * 60 * 24))
}

const getRowClassName = ({ row }) => {
  if (row.domainStatus === 3 || row.certStatus === 3) {
    return '!bg-red-50'
  }
  if (row.domainStatus === 2 || row.certStatus === 2) {
    return '!bg-amber-50'
  }
  return ''
}

const formatYmd = (val) => {
  if (!val || val === '0001-01-01T00:00:00Z') return ''
  return formatTimeToStr(val, 'yyyy-MM-dd')
}
</script>
