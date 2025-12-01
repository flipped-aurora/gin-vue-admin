<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list justify-between flex items-center">
        <span class="text font-bold">字典详细内容</span>
        <div class="flex items-center gap-2">
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
                @click="getTreeData"
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
        :data="treeData"
        style="width: 100%"
        tooltip-effect="dark"
        :tree-props="{ children: 'children'}"
        row-key="ID"
        default-expand-all
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="展示值" prop="label" min-width="100"/>

        <el-table-column align="left" label="字典值" prop="value" />

        <el-table-column align="left" label="扩展值" prop="extend" />

        <el-table-column align="left" label="层级" prop="level" width="80" />

        <el-table-column
          align="left"
          label="启用状态"
          prop="status"
          width="100"
        >
          <template #default="scope">
            {{ formatBoolean(scope.row.status) }}
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          label="排序标记"
          prop="sort"
          width="100"
        />

        <el-table-column
          align="left"
          label="操作"
          :min-width="appStore.operateMinWith"
          fixed="right"
        >
          <template #default="scope">
            <el-button
              type="primary"
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
            :options="[rootOption,...treeData]"
            :props="cascadeProps"
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

  const treeData = ref([])

  // 级联选择器配置
  const cascadeProps = {
    value: 'ID',
    label: 'label',
    children: 'children',
    checkStrictly: true, // 允许选择任意级别
    emitPath: false // 只返回选中节点的值
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
      }
    } catch (error) {
      console.error('获取树形数据失败:', error)
      ElMessage.error('获取层级数据失败')
    }
  }

  const rootOption = {
    ID: null,
    label: '无父级（根级）'
  }


  // 初始加载
  getTreeData()

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
        await getTreeData() // 重新加载数据
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
        await getTreeData() // 重新加载数据
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
    getTreeData()
  }

  const handleCloseSearchInput = () => {
    // 处理搜索输入框关闭
  }

  const handleInputKeyDown = (e) => {
    if (e.key === 'Enter' && searchName.value.trim() !== '') {
      getTreeData()
    }
  }

  watch(
    () => props.sysDictionaryID,
    () => {
      getTreeData()
    }
  )
</script>

<style scoped>

</style>
