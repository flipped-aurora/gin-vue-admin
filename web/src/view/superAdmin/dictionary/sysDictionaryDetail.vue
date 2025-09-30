<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list justify-between flex items-center">
        <span class="text font-bold">字典详细内容</span>
        <div class="flex items-center gap-2">
          <!-- 层级视图切换按钮 -->
          <el-button-group>
            <el-button
              :type="viewMode === 'table' ? 'primary' : 'default'"
              @click="viewMode = 'table'"
              size="small"
            >
              表格视图
            </el-button>
            <el-button
              :type="viewMode === 'tree' ? 'primary' : 'default'"
              @click="viewMode = 'tree'"
              size="small"
            >
              层级视图
            </el-button>
          </el-button-group>

          <el-input
            placeholder="搜索展示值"
            v-model="searchName"
            clearable
            class="!w-64"
            @clear="clearSearchInput"
            :prefix-icon="Search"
            v-click-outside="handleCloseSearchInput"
            @keydown="handleInputKeyDown"
          >
            <template #append>
              <el-button
                :type="searchName ? 'primary' : 'info'"
                @click="getTableData"
                >搜索</el-button
              >
            </template>
          </el-input>
          <el-button type="primary" icon="plus" @click="openDrawer">
            新增字典项
          </el-button>
        </div>
      </div>
      <!-- 表格视图 -->
      <el-table
        v-if="viewMode === 'table'"
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="展示值" prop="label" />

        <el-table-column align="left" label="字典值" prop="value" />

        <el-table-column align="left" label="扩展值" prop="extend" />

        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.CreatedAt) }}
          </template>
        </el-table-column>

        <el-table-column align="left" label="层级" prop="level" width="80" />

        <el-table-column
          align="left"
          label="父级字典项"
          prop="parentID"
          width="150"
        >
          <template #default="scope">
            <span v-if="scope.row.parentID">
              {{ getParentLabel(scope.row.parentID) }}
            </span>
            <span v-else class="text-gray-400">根级</span>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="启用状态"
          prop="status"
          width="120"
        >
          <template #default="scope">
            {{ formatBoolean(scope.row.status) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="排序标记"
          prop="sort"
          width="120"
        />

        <el-table-column
          align="left"
          label="操作"
          :min-width="appStore.operateMinWith"
        >
          <template #default="scope">
            <el-button
              type="success"
              link
              icon="plus"
              @click="addChildNode(scope.row)"
            >
              添加子项
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updateSysDictionaryDetailFunc(scope.row)"
            >
              变更
            </el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteSysDictionaryDetailFunc(scope.row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 层级树形视图 -->
      <div v-if="viewMode === 'tree'" class="tree-view-container">
        <el-tree
          :data="treeData"
          :props="treeProps"
          node-key="ID"
          :expand-on-click-node="false"
          :default-expand-all="false"
          class="dictionary-tree"
        >
          <template #default="{ data }">
            <div class="tree-node-content">
              <div class="node-info">
                <span class="">{{ data.label }}</span>
                <span class="node-value">（{{ data.value }}）</span>
                <span v-if="data.extend" class="node-extend"
                  >- {{ data.extend }}</span
                >
                <el-tag
                  v-if="data.level !== undefined"
                  size="small"
                  class="ml-2"
                >
                  L{{ data.level }}
                </el-tag>
                <el-tag
                  :type="data.status ? 'success' : 'danger'"
                  size="small"
                  class="ml-1"
                >
                  {{ data.status ? '启用' : '禁用' }}
                </el-tag>
              </div>
              <div class="node-actions">
                <el-button
                  type="primary"
                  link
                  size="small"
                  @click="addChildNode(data)"
                >
                  添加子项
                </el-button>
                <el-button
                  type="primary"
                  link
                  size="small"
                  @click="updateSysDictionaryDetailFunc(data)"
                >
                  编辑
                </el-button>
                <el-button
                  type="danger"
                  link
                  size="small"
                  @click="deleteSysDictionaryDetailFunc(data)"
                >
                  删除
                </el-button>
              </div>
            </div>
          </template>
        </el-tree>
      </div>

      <!-- 分页 -->
      <div v-if="viewMode === 'table'" class="gva-pagination">
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
      v-model="drawerFormVisible"
      :size="appStore.drawerSize"
      :show-close="false"
      :before-close="closeDrawer"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            type === 'create' ? '添加字典项' : '修改字典项'
          }}</span>
          <div>
            <el-button @click="closeDrawer"> 取 消 </el-button>
            <el-button type="primary" @click="enterDrawer"> 确 定 </el-button>
          </div>
        </div>
      </template>
      <el-form
        ref="drawerForm"
        :model="formData"
        :rules="rules"
        label-width="110px"
      >
        <el-form-item label="父级字典项" prop="parentID">
          <el-cascader
            v-model="formData.parentID"
            :options="parentOptions"
            :props="cascaderProps"
            placeholder="请选择父级字典项（可选）"
            clearable
            filterable
            :style="{ width: '100%' }"
            @change="handleParentChange"
          />
        </el-form-item>
        <el-form-item label="展示值" prop="label">
          <el-input
            v-model="formData.label"
            placeholder="请输入展示值"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item label="字典值" prop="value">
          <el-input
            v-model="formData.value"
            placeholder="请输入字典值"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item label="扩展值" prop="extend">
          <el-input
            v-model="formData.extend"
            placeholder="请输入扩展值"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item label="启用状态" prop="status" required>
          <el-switch
            v-model="formData.status"
            active-text="开启"
            inactive-text="停用"
          />
        </el-form-item>
        <el-form-item label="排序标记" prop="sort">
          <el-input-number
            v-model.number="formData.sort"
            placeholder="排序标记"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createSysDictionaryDetail,
    deleteSysDictionaryDetail,
    updateSysDictionaryDetail,
    findSysDictionaryDetail,
    getSysDictionaryDetailList,
    getDictionaryTreeList
  } from '@/api/sysDictionaryDetail' // 此处请自行替换地址
  import { ref, watch } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { formatBoolean, formatDate } from '@/utils/format'
  import { useAppStore } from '@/pinia'
  import { Search } from '@element-plus/icons-vue'

  defineOptions({
    name: 'SysDictionaryDetail'
  })

  const appStore = useAppStore()
  const searchName = ref('')
  const viewMode = ref('table') // 'table' 或 'tree'

  const props = defineProps({
    sysDictionaryID: {
      type: Number,
      default: 0
    }
  })

  const formData = ref({
    label: null,
    value: null,
    status: true,
    sort: null,
    parentID: null
  })

  const rules = ref({
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
  })

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const treeData = ref([])
  const parentOptions = ref([])

  // 树形组件配置
  const treeProps = {
    children: 'children',
    label: 'label'
  }

  // 级联选择器配置
  const cascaderProps = {
    value: 'value', // 修改为使用value字段
    label: 'label',
    children: 'children',
    checkStrictly: true, // 允许选择任意级别
    emitPath: false // 只返回选中节点的值
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

  // 查询表格数据
  const getTableData = async () => {
    if (!props.sysDictionaryID) return

    // 首先获取树形数据以便查找父级标签
    try {
      const treeRes = await getDictionaryTreeList({
        sysDictionaryID: props.sysDictionaryID
      })
      if (treeRes.code === 0) {
        treeData.value = treeRes.data || []
        updateParentOptions(treeRes.data.list || [])
      }
    } catch (error) {
      console.error('获取树形数据失败:', error)
    }

    // 然后获取表格数据
    const table = await getSysDictionaryDetailList({
      page: page.value,
      pageSize: pageSize.value,
      sysDictionaryID: props.sysDictionaryID,
      label: searchName.value.trim()
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    }
  }

  // 获取树形数据
  const getTreeData = async () => {
    if (!props.sysDictionaryID) return
    try {
      const res = await getDictionaryTreeList({
        sysDictionaryID: props.sysDictionaryID
      })
      if (res.code === 0) {
        treeData.value = res.data.list || []
        // 同时更新父级选项
        updateParentOptions(res.data.list || [])
      }
    } catch (error) {
      console.error('获取树形数据失败:', error)
      ElMessage.error('获取层级数据失败')
    }
  }

  // 更新父级选项
  const updateParentOptions = (data) => {
    const convertToOptions = (items) => {
      return items.map((item) => ({
        ID: item.ID,
        label: item.label,
        value: item.ID, // 使用ID作为value
        children:
          item.children && item.children.length > 0
            ? convertToOptions(item.children)
            : undefined
      }))
    }

    // 添加根级选项（无父级）
    const rootOption = {
      ID: null,
      label: '无父级（根级）',
      value: null
    }

    parentOptions.value = [rootOption, ...convertToOptions(data)]
  }

  // 根据父级ID获取父级标签
  const getParentLabel = (parentID) => {
    if (!parentID) return '根级'

    // 从表格数据中查找父级项
    const parentItem = tableData.value.find((item) => item.ID === parentID)
    if (parentItem) {
      return parentItem.label
    }

    // 从树形数据中递归查找父级项
    const findInTree = (items) => {
      for (const item of items) {
        if (item.ID === parentID) {
          return item.label
        }
        if (item.children && item.children.length > 0) {
          const found = findInTree(item.children)
          if (found) return found
        }
      }
      return null
    }

    const treeResult = findInTree(treeData.value)
    return treeResult || `ID: ${parentID}`
  }

  // 根据视图模式获取数据
  const loadData = () => {
    if (viewMode.value === 'table') {
      getTableData()
    } else {
      getTreeData()
    }
  }

  // 监听视图模式变化
  watch(viewMode, () => {
    loadData()
  })

  // 初始加载
  loadData()

  const type = ref('')
  const drawerFormVisible = ref(false)

  const updateSysDictionaryDetailFunc = async (row) => {
    drawerForm.value && drawerForm.value.clearValidate()
    const res = await findSysDictionaryDetail({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data.reSysDictionaryDetail
      drawerFormVisible.value = true
    }
  }

  // 添加子节点
  const addChildNode = (parentNode) => {
    console.log(parentNode)
    type.value = 'create'
    formData.value = {
      label: null,
      value: null,
      status: true,
      sort: null,
      parentID: parentNode.ID,
      sysDictionaryID: props.sysDictionaryID
    }
    drawerForm.value && drawerForm.value.clearValidate()
    drawerFormVisible.value = true
  }

  // 处理父级选择变化
  const handleParentChange = (value) => {
    formData.value.parentID = value
  }

  const closeDrawer = () => {
    drawerFormVisible.value = false
    formData.value = {
      label: null,
      value: null,
      status: true,
      sort: null,
      parentID: null,
      sysDictionaryID: props.sysDictionaryID
    }
  }

  const deleteSysDictionaryDetailFunc = async (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const res = await deleteSysDictionaryDetail({ ID: row.ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        loadData() // 重新加载数据
      }
    })
  }

  const drawerForm = ref(null)
  const enterDrawer = async () => {
    drawerForm.value.validate(async (valid) => {
      formData.value.sysDictionaryID = props.sysDictionaryID
      if (!valid) return
      let res
      switch (type.value) {
        case 'create':
          res = await createSysDictionaryDetail(formData.value)
          break
        case 'update':
          res = await updateSysDictionaryDetail(formData.value)
          break
        default:
          res = await createSysDictionaryDetail(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDrawer()
        loadData() // 重新加载数据
      }
    })
  }

  const openDrawer = () => {
    type.value = 'create'
    formData.value.parentID = null
    drawerForm.value && drawerForm.value.clearValidate()
    drawerFormVisible.value = true
  }

  const clearSearchInput = () => {
    searchName.value = ''
    loadData()
  }

  const handleCloseSearchInput = () => {
    // 处理搜索输入框关闭
  }

  const handleInputKeyDown = (e) => {
    if (e.key === 'Enter' && searchName.value.trim() !== '') {
      loadData()
    }
  }

  watch(
    () => props.sysDictionaryID,
    () => {
      loadData()
    }
  )
</script>

<style scoped>
  .tree-view-container {
    min-height: 400px;
    border: 1px solid #e4e7ed;
    border-radius: 4px;
    padding: 16px;
  }

  .dictionary-tree {
    width: 100%;
  }

  .tree-node-content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    padding: 4px 8px;
  }

  .node-info {
    display: flex;
    align-items: center;
    flex: 1;
  }

  .node-value {
    color: #909399;
    font-size: 12px;
    margin-left: 8px;
  }

  .node-extend {
    color: #606266;
    font-size: 12px;
    margin-left: 4px;
  }

  .node-actions {
    display: flex;
    gap: 8px;
    opacity: 0;
    transition: opacity 0.2s;
  }

  .tree-node-content:hover .node-actions {
    opacity: 1;
  }

  .ml-1 {
    margin-left: 4px;
  }

  .ml-2 {
    margin-left: 8px;
  }
</style>
