<template>
  <div
    ref="rootRef"
    class="p-2 flex flex-col overflow-hidden"
    :style="{ height: rootHeight }"
  >
    <el-splitter class="flex-1 min-h-0">
      <el-splitter-panel size="380px" min="260px" max="700px" collapsible>
        <div
          class="gva-table-box !my-0 mr-2 h-full flex flex-col overflow-hidden"
        >
          <div class="flex justify-between items-center flex-none gap-2">
            <span class="font-bold flex-none">岗位列表</span>
            <el-input
              v-model="searchName"
              placeholder="搜索岗位"
              clearable
              class="!flex-1"
              @keyup.enter="getTableData"
              @clear="getTableData"
            />
            <el-button
              type="primary"
              icon="plus"
              class="flex-none"
              @click="addPosition"
              >新增</el-button
            >
          </div>
          <el-scrollbar class="mt-3 flex-1">
            <div
              v-for="pos in tableData"
              :key="pos.ID"
              class="rounded flex justify-between items-center px-3 py-3 cursor-pointer mt-2 gap-2 group bg-gray-50 dark:bg-gray-800 hover:bg-blue-50 dark:hover:bg-blue-900"
              :class="
                selectedPos.ID === pos.ID
                  ? 'text-active'
                  : 'text-slate-700 dark:text-slate-50'
              "
              @click="selectedPos = pos"
            >
              <div class="truncate">
                {{ pos.name }}
                <span v-if="pos.code" class="text-sm text-gray-400"
                  >（{{ pos.code }}）</span
                >
                <el-tag v-if="!pos.status" size="small" type="info" class="ml-1"
                  >停用</el-tag
                >
              </div>
              <div class="hidden group-hover:inline-flex items-center gap-2">
                <el-icon
                  class="!text-blue-500"
                  title="编辑"
                  @click.stop="editPosition(pos)"
                >
                  <Edit />
                </el-icon>
                <el-icon
                  class="!text-red-500"
                  title="删除"
                  @click.stop="deleteRow(pos)"
                >
                  <Delete />
                </el-icon>
              </div>
            </div>
            <el-empty
              v-if="!tableData.length"
              description="暂无岗位"
              :image-size="80"
            />
          </el-scrollbar>
        </div>
      </el-splitter-panel>
      <el-splitter-panel :min="300">
        <div class="gva-table-box !my-0 ml-2 h-full overflow-hidden">
          <org-member-assign
            kind="position"
            :target-id="selectedPos.ID || 0"
            :target-name="selectedPos.name || ''"
          />
        </div>
      </el-splitter-panel>
    </el-splitter>

    <el-drawer
      v-model="formVisible"
      :size="appStore.drawerSize"
      :show-close="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">{{ titleForm }}</span>
          <div>
            <el-button @click="closeForm">取 消</el-button>
            <el-button type="primary" @click="submitForm">确 定</el-button>
          </div>
        </div>
      </template>
      <el-form ref="posForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="岗位名称" prop="name">
          <el-input
            v-model="form.name"
            autocomplete="off"
            placeholder="请输入岗位名称"
          />
        </el-form-item>
        <el-form-item label="岗位编码" prop="code">
          <el-input
            v-model="form.code"
            autocomplete="off"
            placeholder="请输入岗位编码"
          />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-switch v-model="form.status" active-text="启用" inactive-text="停用" />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input
            v-model="form.remark"
            type="textarea"
            autocomplete="off"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    getPositionList,
    createPosition,
    updatePosition,
    deletePosition
  } from '@/api/position'
  import OrgMemberAssign from '@/view/superAdmin/orgMemberAssign.vue'
  import { Edit, Delete } from '@element-plus/icons-vue'
  import {
    ref,
    onMounted,
    onActivated,
    onBeforeUnmount,
    nextTick
  } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { useAppStore } from '@/pinia'

  defineOptions({
    name: 'Position'
  })

  const appStore = useAppStore()
  const tableData = ref([])
  const selectedPos = ref({})
  const searchName = ref('')

  // 整页填满视口高度（照字典管理）
  const rootRef = ref(null)
  const rootHeight = ref('auto')
  const updateRootHeight = () => {
    const el = rootRef.value
    if (!el) return
    const top = el.getBoundingClientRect().top
    rootHeight.value = `${Math.max(320, window.innerHeight - top - 32)}px`
  }
  onMounted(() => {
    nextTick(updateRootHeight)
    window.addEventListener('resize', updateRootHeight)
  })
  onActivated(updateRootHeight)
  onBeforeUnmount(() => window.removeEventListener('resize', updateRootHeight))

  const getTableData = async () => {
    const res = await getPositionList({
      page: 1,
      pageSize: 1000,
      name: searchName.value.trim()
    })
    if (res.code === 0) {
      tableData.value = res.data.list || []
      if (!selectedPos.value.ID && tableData.value.length) {
        selectedPos.value = tableData.value[0]
      }
    }
  }
  getTableData()

  const formVisible = ref(false)
  const titleForm = ref('新增岗位')
  const dialogType = ref('add')
  const posForm = ref(null)
  const defaultForm = () => ({
    ID: 0,
    name: '',
    code: '',
    sort: 0,
    status: true,
    remark: ''
  })
  const form = ref(defaultForm())

  const rules = ref({
    name: [{ required: true, message: '请输入岗位名称', trigger: 'blur' }]
  })

  const initForm = () => {
    posForm.value && posForm.value.resetFields()
    form.value = defaultForm()
  }
  const closeForm = () => {
    initForm()
    formVisible.value = false
  }

  const addPosition = () => {
    initForm()
    titleForm.value = '新增岗位'
    dialogType.value = 'add'
    formVisible.value = true
  }

  const editPosition = (row) => {
    titleForm.value = '编辑岗位'
    dialogType.value = 'edit'
    form.value = {
      ID: row.ID,
      name: row.name,
      code: row.code,
      sort: row.sort,
      status: !!row.status,
      remark: row.remark
    }
    posForm.value && posForm.value.clearValidate()
    formVisible.value = true
  }

  const submitForm = () => {
    posForm.value.validate(async (valid) => {
      if (!valid) return
      const api = dialogType.value === 'add' ? createPosition : updatePosition
      const res = await api(form.value)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: dialogType.value === 'add' ? '添加成功!' : '编辑成功!'
        })
        getTableData()
        closeForm()
      }
    })
  }

  const deleteRow = (row) => {
    ElMessageBox.confirm('此操作将永久删除该岗位, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async () => {
        const res = await deletePosition({ id: row.ID })
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '删除成功!' })
          if (selectedPos.value.ID === row.ID) selectedPos.value = {}
          getTableData()
        }
      })
      .catch(() => {
        ElMessage({ type: 'info', message: '已取消删除' })
      })
  }
</script>
