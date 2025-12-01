<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAtRange">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>

          <el-date-picker v-model="searchInfo.createdAtRange" class="w-[380px]" type="datetimerange" range-separator="至"
            start-placeholder="开始时间" end-placeholder="结束时间" />
        </el-form-item>

        <el-form-item label="版本名称" prop="versionName">
          <el-input v-model="searchInfo.versionName" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="版本号" prop="versionCode">
          <el-input v-model="searchInfo.versionCode" placeholder="搜索条件" />
        </el-form-item>



        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery = true"
            v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery = false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="success" icon="download" @click="openExportDialog">创建发版</el-button>
        <el-button type="warning" icon="upload" @click="openImportDialog">导入版本</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="onDelete">删除</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="版本名称" prop="versionName" width="120" />

        <el-table-column align="left" label="版本号" prop="versionCode" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="320">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看</el-button>
            <el-button type="success" link icon="download" class="table-button"
              @click="downloadJson(scope.row)">下载发版包</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true"
      :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="版本名称">
          {{ detailForm.versionName }}
        </el-descriptions-item>
        <el-descriptions-item label="版本号">
          {{ detailForm.versionCode }}
        </el-descriptions-item>
        <el-descriptions-item label="版本描述">
          {{ detailForm.description }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <!-- 导出版本抽屉 -->
    <el-drawer v-model="exportDialogVisible" title="创建发版" direction="rtl" size="80%" :before-close="closeExportDialog"
      :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">创建发版</span>
          <div>
            <el-button @click="closeExportDialog">取消</el-button>
            <el-button type="primary" @click="handleExport" :loading="exportLoading">创建发版</el-button>
          </div>
        </div>
      </template>
      <el-form :model="exportForm" label-width="100px">
        <el-form-item label="版本名称" required>
          <el-input v-model="exportForm.versionName" placeholder="请输入版本名称" />
        </el-form-item>
        <el-form-item label="版本号" required>
          <el-input v-model="exportForm.versionCode" placeholder="请输入版本号" />
        </el-form-item>
        <el-form-item label="版本描述">
          <el-input v-model="exportForm.description" type="textarea" placeholder="请输入版本描述" />
        </el-form-item>
        <el-form-item label="发版信息">
          <div class="flex gap-3 w-full">
            <!-- 菜单选择 -->
            <div class="card-col card-vertical">
              <div class="card-header">
                <span class="card-title">选择菜单</span>
              </div>
              <div class="card-filter">
                <el-input v-model="menuFilterText" placeholder="输入关键字进行过滤" clearable size="small" />
              </div>
              <div class="card-body">
                <el-tree ref="menuTreeRef" :data="menuTreeData" :default-checked-keys="selectedMenuIds"
                  :props="menuTreeProps" default-expand-all highlight-current node-key="ID" show-checkbox
                  :filter-node-method="filterMenuNode" @check="onMenuCheck" class="menu-tree">
                  <template #default="{ node }">
                    <span class="flex-1 flex items-center justify-between text-sm pr-2">
                      <span>{{ node.label }}</span>
                    </span>
                  </template>
                </el-tree>
              </div>
            </div>

            <!-- API选择 -->
            <div class="card-col card-vertical">
              <div class="card-header">
                <span class="card-title">选择API</span>
              </div>
              <div class="card-filter">
                <el-input v-model="apiFilterTextName" placeholder="按名称过滤" clearable size="small"
                  style="margin-bottom: 8px" />
                <el-input v-model="apiFilterTextPath" placeholder="按路径过滤" clearable size="small" />
              </div>
              <div class="card-body">
                <el-tree ref="apiTreeRef" :data="apiTreeData" :default-checked-keys="selectedApiIds"
                  :props="apiTreeProps" default-expand-all highlight-current node-key="onlyId" show-checkbox
                  :filter-node-method="filterApiNode" @check="onApiCheck" class="api-tree">
                  <template #default="{ data }">
                    <div class="flex items-center justify-between w-full pr-1">
                      <span>{{ data.description }}</span>
                      <el-tooltip :content="data.path">
                        <span
                          class="max-w-[240px] break-all overflow-ellipsis overflow-hidden text-gray-500 dark:text-gray-400">
                          {{ data.path }}
                        </span>
                      </el-tooltip>
                    </div>
                  </template>
                </el-tree>
              </div>
            </div>

            <!-- 字典选择 -->
            <div class="card-col card-vertical">
              <div class="card-header">
                <span class="card-title">选择字典</span>
              </div>
              <div class="card-filter">
                <el-input v-model="dictFilterText" placeholder="输入关键字进行过滤" clearable size="small" />
              </div>
              <div class="card-body">
                <el-tree ref="dictTreeRef" :data="dictTreeData" :default-checked-keys="selectedDictIds"
                  :props="dictTreeProps" default-expand-all highlight-current node-key="ID" show-checkbox
                  :filter-node-method="filterDictNode" @check="onDictCheck" class="dict-tree">
                  <template #default="{ data }">
                    <div class="flex items-center justify-between w-full pr-1">
                      <span>{{ data.name || data.label }}</span>
                      <el-tooltip :content="data.desc || (data.value ? `值: ${data.value}` : '')">
                        <span class="text-gray-500 dark:text-gray-400 text-xs ml-2">
                          {{ data.type || (data.value ? `值: ${data.value}` : '') }}
                        </span>
                      </el-tooltip>
                    </div>
                  </template>
                </el-tree>
              </div>
            </div>
          </div>
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 导入版本抽屉 -->
    <el-drawer v-model="importDialogVisible" title="导入版本" direction="rtl" size="80%" :before-close="closeImportDialog"
      :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">导入版本</span>
          <div>
            <el-button @click="closeImportDialog">取消</el-button>
            <el-button type="primary" @click="handleImport" :loading="importLoading"
              :disabled="!importJsonContent.trim()">导入</el-button>
          </div>
        </div>
      </template>
      <el-form label-width="100px">
        <el-form-item label="上传文件">
          <el-upload ref="uploadRef" :auto-upload="false" :show-file-list="true" :limit="1" accept=".json"
            :on-change="handleFileChange" :on-remove="handleFileRemove" drag>
            <el-icon class="el-icon--upload"><upload-filled /></el-icon>
            <div class="el-upload__text">
              将JSON文件拖到此处，或<em>点击上传</em>
            </div>
            <template #tip>
              <div class="el-upload__tip">
                只能上传JSON文件
              </div>
            </template>
          </el-upload>
        </el-form-item>
        <el-form-item label="版本JSON">
          <el-input v-model="importJsonContent" type="textarea" :rows="20" placeholder="请粘贴版本JSON"
            @input="handleJsonContentChange" />
        </el-form-item>
        <el-form-item label="预览内容" v-if="importPreviewData">
          <div class="preview-wrap">
            <div class="flex gap-3 w-full">
              <div class="card-col">
                <div class="card-vertical">
                  <div class="card-header">
                    <h3 class="card-title">菜单 ({{ getTotalMenuCount() }}项)</h3>
                  </div>
                  <div class="card-body">
                    <el-tree :data="previewMenuTreeData" :props="menuTreeProps" node-key="name"
                      :expand-on-click-node="false" :check-on-click-node="false" :show-checkbox="false"
                      default-expand-all>
                      <template #default="{ data }">
                        <div class="flex-1 flex items-center justify-between text-sm pr-2">
                          <span>{{ data.meta?.title || data.title }}</span>
                          <span class="text-gray-500 dark:text-gray-400 text-xs ml-2">{{ data.path }}</span>
                        </div>
                      </template>
                    </el-tree>
                  </div>
                </div>
              </div>
              <div class="card-col">
                <div class="card-vertical">
                  <div class="card-header">
                    <h3 class="card-title">API ({{ importPreviewData.apis?.length || 0 }}项)</h3>
                  </div>
                  <div class="card-body">
                    <el-tree :data="previewApiTreeData" :props="apiTreeProps" node-key="ID"
                      :expand-on-click-node="false" :check-on-click-node="false" :show-checkbox="false"
                      default-expand-all>
                      <template #default="{ data }">
                        <div class="flex-1 flex items-center justify-between text-sm pr-2">
                          <span>{{ data.description }}</span>
                          <span class="text-gray-500 dark:text-gray-400 text-xs ml-2">{{ data.path }} [{{ data.method
                            }}]</span>
                        </div>
                      </template>
                    </el-tree>
                  </div>
                </div>
              </div>
              <div class="card-col">
                <div class="card-vertical">
                  <div class="card-header">
                    <h3 class="card-title">字典 ({{ importPreviewData.dictionaries?.length || 0 }}项)</h3>
                  </div>
                  <div class="card-body">
                    <el-tree :data="previewDictTreeData" :props="dictTreeProps" node-key="ID"
                      :expand-on-click-node="false" :check-on-click-node="false" :show-checkbox="false"
                      default-expand-all>
                      <template #default="{ data }">
                        <div class="flex-1 flex items-center justify-between text-sm pr-2">
                          <span>{{ data.name || data.label }}</span>
                          <span class="text-gray-500 dark:text-gray-400 text-xs ml-2">
                            {{ data.type || (data.value ? `值: ${data.value}` : '') }}
                          </span>
                        </div>
                      </template>
                    </el-tree>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </el-form-item>
      </el-form>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  deleteSysVersion,
  deleteSysVersionByIds,
  findSysVersion,
  getSysVersionList,
  exportVersion,
  importVersion,
  downloadVersionJson
} from '@/api/version'

// 导入菜单和API相关接口
import { getMenuList } from '@/api/menu'
import { getApiList } from '@/api/api'
import { getSysDictionaryList } from '@/api/sysDictionary'

// 全量引入格式化工具 请按需保留
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { UploadFilled } from '@element-plus/icons-vue'
import { ref, watch } from 'vue'
import { useAppStore } from "@/pinia"

defineOptions({
  name: 'SysVersion'
})

const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 导出相关数据
const exportDialogVisible = ref(false)
const exportLoading = ref(false)
const exportForm = ref({
  versionName: '',
  versionCode: '',
  description: '',
  menuIds: [],
  apiIds: [],
  dictIds: []
})

// 树形结构相关数据
const menuTreeData = ref([])
const apiTreeData = ref([])
const dictTreeData = ref([])
const selectedMenuIds = ref([])
const selectedApiIds = ref([])
const selectedDictIds = ref([])
const menuFilterText = ref('')
const apiFilterTextName = ref('')
const apiFilterTextPath = ref('')
const dictFilterText = ref('')

// 树形组件引用
const menuTreeRef = ref(null)
const apiTreeRef = ref(null)
const dictTreeRef = ref(null)

// 树形属性配置
const menuTreeProps = ref({
  children: 'children',
  label: function (data) {
    return data.meta?.title || data.title
  }
})

const apiTreeProps = ref({
  children: 'children',
  label: 'description'
})

const dictTreeProps = ref({
  children: 'sysDictionaryDetails',
  label: function (data) {
    // 如果是字典主项，显示字典名称
    if (data.name) {
      return data.name
    }
    // 如果是字典详情项，显示标签
    if (data.label) {
      return data.label
    }
    return '未知项'
  }
})

// 导入相关数据
const importDialogVisible = ref(false)
const importLoading = ref(false)
const importJsonContent = ref('')
const importPreviewData = ref(null)
const uploadRef = ref(null)
const previewMenuTreeData = ref([])
const previewApiTreeData = ref([])
const previewDictTreeData = ref([])



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
  const table = await getSysVersionList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

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
    deleteSysVersionFunc(row)
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
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deleteSysVersionByIds({ IDs })
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
const deleteSysVersionFunc = async (row) => {
  const res = await deleteSysVersion({ ID: row.ID })
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
  const res = await findSysVersion({ ID: row.ID })
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



// 获取菜单和API列表
const getMenuAndApiList = async () => {
  try {
    // 获取菜单列表
    const menuRes = await getMenuList()
    if (menuRes.code === 0) {
      menuTreeData.value = menuRes.data || []
    }

    // 获取API列表
    const apiRes = await getApiList({ page: 1, pageSize: 9999 })
    if (apiRes.code === 0) {
      console.log('原始API数据:', apiRes.data)
      const apis = apiRes.data.list || []
      apiTreeData.value = buildApiTree(apis)
    }
  } catch (error) {
    console.error('获取数据失败:', error)
    ElMessage.error('获取菜单或API数据失败')
  }
}

// 获取字典列表
const getDictList = async () => {
  try {
    const dictRes = await getSysDictionaryList({ page: 1, pageSize: 9999 })
    if (dictRes.code === 0) {
      dictTreeData.value = dictRes.data || []
    }
  } catch (error) {
    console.error('获取字典数据失败:', error)
    ElMessage.error('获取字典数据失败')
  }
}

// 构建API树形结构
const buildApiTree = (apis) => {
  const apiObj = {}
  apis.forEach((item) => {
    item.onlyId = 'p:' + item.path + 'm:' + item.method
    if (Object.prototype.hasOwnProperty.call(apiObj, item.apiGroup)) {
      apiObj[item.apiGroup].push(item)
    } else {
      Object.assign(apiObj, { [item.apiGroup]: [item] })
    }
  })
  const apiTree = []
  for (const key in apiObj) {
    const treeNode = {
      ID: key,
      description: key + '组',
      children: apiObj[key]
    }
    apiTree.push(treeNode)
  }
  return apiTree
}

// 树形组件事件处理方法
const filterMenuNode = (value, data) => {
  if (!value) return true
  const title = data.meta?.title || data.title || ''
  return title.indexOf(value) !== -1
}

const filterApiNode = (value, data) => {
  if (!apiFilterTextName.value && !apiFilterTextPath.value) return true
  let matchesName, matchesPath
  if (!apiFilterTextName.value) {
    matchesName = true
  } else {
    matchesName = data.description && data.description.includes(apiFilterTextName.value)
  }
  if (!apiFilterTextPath.value) {
    matchesPath = true
  } else {
    matchesPath = data.path && data.path.includes(apiFilterTextPath.value)
  }
  return matchesName && matchesPath
}

const filterDictNode = (value, data) => {
  if (!value) return true
  const name = data.name || ''
  const type = data.type || ''
  const desc = data.desc || ''
  const label = data.label || ''
  const dataValue = data.value || ''
  return name.indexOf(value) !== -1 ||
    type.indexOf(value) !== -1 ||
    desc.indexOf(value) !== -1 ||
    label.indexOf(value) !== -1 ||
    dataValue.indexOf(value) !== -1
}

const onMenuCheck = (data, checked) => {
  if (checked.checkedKeys) {
    selectedMenuIds.value = checked.checkedKeys
  }
}

const onApiCheck = (data, checked) => {
  if (checked.checkedKeys) {
    selectedApiIds.value = checked.checkedKeys
  }
}

const onDictCheck = (data, checked) => {
  if (checked.checkedKeys) {
    selectedDictIds.value = checked.checkedKeys
  }
}

// 监听过滤文本变化
watch(menuFilterText, (val) => {
  if (menuTreeRef.value) {
    menuTreeRef.value.filter(val)
  }
})

watch([apiFilterTextName, apiFilterTextPath], () => {
  if (apiTreeRef.value) {
    apiTreeRef.value.filter('')
  }
})

watch(dictFilterText, (val) => {
  if (dictTreeRef.value) {
    dictTreeRef.value.filter(val)
  }
})

// 导出相关方法
const openExportDialog = async () => {
  exportDialogVisible.value = true
  await getMenuAndApiList()
  await getDictList()
}

const closeExportDialog = () => {
  exportDialogVisible.value = false
  exportForm.value = {
    versionName: '',
    versionCode: '',
    description: '',
    menuIds: [],
    apiIds: [],
    dictIds: []
  }
  selectedMenuIds.value = []
  selectedApiIds.value = []
  selectedDictIds.value = []
  menuFilterText.value = ''
  apiFilterTextName.value = ''
  apiFilterTextPath.value = ''
  dictFilterText.value = ''
}

const handleExport = async () => {
  if (!exportForm.value.versionName || !exportForm.value.versionCode) {
    ElMessage.warning('请填写版本名称和版本号')
    return
  }

  exportLoading.value = true
  try {
    // 获取选中的菜单、API和字典
    const checkedMenus = menuTreeRef.value ? menuTreeRef.value.getCheckedNodes(false, true) : []
    const checkedApis = apiTreeRef.value ? apiTreeRef.value.getCheckedNodes(true) : []
    const checkedDicts = dictTreeRef.value ? dictTreeRef.value.getCheckedNodes(true) : []

    const menuIds = checkedMenus.map(menu => menu.ID)
    const apiIds = checkedApis.map(api => api.ID)
    const dictIds = checkedDicts.map(dict => dict.ID)

    exportForm.value.menuIds = menuIds
    exportForm.value.apiIds = apiIds
    exportForm.value.dictIds = dictIds

    const res = await exportVersion(exportForm.value)
    if (res.code !== 0) {
      ElMessage.error(res.msg || '创建发版失败')
      return
    }

    ElMessage.success('创建发版成功')
    closeExportDialog()
    getTableData() // 刷新表格数据
  } catch (error) {
    console.error('创建发版失败:', error)
    ElMessage.error('创建发版失败')
  } finally {
    exportLoading.value = false
  }
}

// 导入相关方法
const openImportDialog = () => {
  importDialogVisible.value = true
}

const closeImportDialog = () => {
  importDialogVisible.value = false
  importJsonContent.value = ''
  importPreviewData.value = null
  previewMenuTreeData.value = []
  previewApiTreeData.value = []
  // 清理上传文件
  if (uploadRef.value) {
    uploadRef.value.clearFiles()
  }
}

// 文件上传处理函数
const handleFileChange = (file) => {
  if (!file.raw) return

  // 验证文件类型
  if (!file.name.toLowerCase().endsWith('.json')) {
    ElMessage.error('只能上传JSON文件')
    uploadRef.value.clearFiles()
    return
  }

  // 读取文件内容
  const reader = new FileReader()
  reader.onload = (e) => {
    try {
      const content = e.target.result
      // 验证JSON格式
      JSON.parse(content)
      importJsonContent.value = content
      handleJsonContentChange()
      ElMessage.success('文件上传成功')
    } catch (error) {
      ElMessage.error('JSON文件格式错误')
      uploadRef.value.clearFiles()
    }
  }
  reader.readAsText(file.raw)
}

const handleFileRemove = () => {
  importJsonContent.value = ''
  importPreviewData.value = null
  previewMenuTreeData.value = []
  previewApiTreeData.value = []
}

// 计算菜单总数（递归计算所有菜单项）
const getTotalMenuCount = () => {
  if (!importPreviewData.value?.menus) return 0

  const countMenus = (menus) => {
    let count = 0
    menus.forEach(menu => {
      count += 1 // 当前菜单
      if (menu.children && menu.children.length > 0) {
        count += countMenus(menu.children) // 递归计算子菜单
      }
    })
    return count
  }

  return countMenus(importPreviewData.value.menus)
}



const handleJsonContentChange = () => {
  if (!importJsonContent.value.trim()) {
    importPreviewData.value = null
    previewMenuTreeData.value = []
    previewApiTreeData.value = []
    previewDictTreeData.value = []
    return
  }

  try {
    const data = JSON.parse(importJsonContent.value)

    // 构建预览数据
    importPreviewData.value = {
      menus: data.menus || [],
      apis: data.apis || [],
      dictionaries: data.dictionaries || []
    }

    // 直接使用菜单数据，因为它已经是树形结构（包含children字段）
    if (data.menus && data.menus.length > 0) {
      previewMenuTreeData.value = data.menus
    } else {
      previewMenuTreeData.value = []
    }

    // 构建API树形数据（按分组组织）
    if (data.apis && data.apis.length > 0) {
      const apiGroups = {}
      data.apis.forEach(api => {
        const group = api.apiGroup || '未分组'
        if (!apiGroups[group]) {
          apiGroups[group] = {
            ID: `group_${group}`,
            description: group,
            path: '',
            method: '',
            children: []
          }
        }
        apiGroups[group].children.push(api)
      })
      previewApiTreeData.value = Object.values(apiGroups)
    } else {
      previewApiTreeData.value = []
    }

    // 处理字典数据
    if (data.dictionaries && data.dictionaries.length > 0) {
      previewDictTreeData.value = data.dictionaries
    } else {
      previewDictTreeData.value = []
    }
  } catch (error) {
    console.error('JSON解析失败:', error)
    importPreviewData.value = null
    previewMenuTreeData.value = []
    previewApiTreeData.value = []
    previewDictTreeData.value = []
  }
}

const handleImport = async () => {
  if (!importJsonContent.value.trim()) {
    ElMessage.warning('请输入版本JSON')
    return
  }

  try {
    JSON.parse(importJsonContent.value)
  } catch (error) {
    ElMessage.error('JSON格式错误，请检查输入内容')
    return
  }

  importLoading.value = true
  try {
    const data = JSON.parse(importJsonContent.value)
    const res = await importVersion(data)
    if (res.code === 0) {
      ElMessage.success('导入成功')
      closeImportDialog()
      getTableData() // 刷新表格数据
    } else {
      ElMessage.error(res.msg || '导入失败')
    }
  } catch (error) {
    console.error('导入失败:', error)
    ElMessage.error('导入失败')
  } finally {
    importLoading.value = false
  }
}

// 下载版本JSON
const downloadJson = async (row) => {
  try {
    const res = await downloadVersionJson({ ID: row.ID })
    // 处理axios响应，获取实际的blob数据
    // 当responseType为blob时，axios拦截器会返回完整的response对象
    let blob
    if (res instanceof Blob) {
      blob = res
    } else if (res.data instanceof Blob) {
      blob = res.data
    } else {
      // 如果不是blob，可能是错误响应，尝试从response中获取
      blob = res
    }

    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = `${row.versionName}_${row.versionCode}.json`
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)

    ElMessage.success('下载成功')
  } catch (error) {
    console.error('下载失败:', error)
    ElMessage.error('下载失败')
  }
}

</script>

<style scoped>
/* 复用卡片样式（支持暗色） */
.card-col {
  @apply border border-gray-300 dark:border-gray-600 rounded overflow-hidden flex-1 bg-white dark:bg-gray-900;
}

.card-vertical {
  @apply flex flex-col h-full;
}

.card-header {
  @apply flex justify-between items-center px-4 py-3 bg-gray-50 dark:bg-gray-800 border-b border-gray-300 dark:border-gray-600;
}

.card-title {
  @apply m-0 text-gray-800 dark:text-gray-200 text-base font-medium;
}

.card-filter {
  @apply px-4 py-3 border-b border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-800;
}

.card-body {
  @apply flex-1 p-2 min-h-[300px] max-h-[400px] overflow-y-auto;
}

.preview-wrap {
  @apply flex flex-col flex-1 gap-4 border border-gray-300 dark:border-gray-600 rounded p-4 bg-gray-50 dark:bg-gray-900;
}

/* Element Plus 树形组件样式优化 */
:deep(.el-tree) {
  background-color: transparent;
}

:deep(.el-tree-node__content) {
  height: 32px;
  line-height: 32px;
}

:deep(.el-tree-node__label) {
  font-size: 14px;
}

:deep(.el-scrollbar__view) {
  padding: 0;
}
</style>
