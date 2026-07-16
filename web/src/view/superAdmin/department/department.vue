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
          <div class="flex justify-between items-center flex-none">
            <span class="font-bold">部门列表</span>
            <el-button type="primary" icon="plus" @click="addDepartment(0)"
              >新增部门</el-button
            >
          </div>
          <el-scrollbar class="mt-3 flex-1">
            <el-tree
              :data="tableData"
              node-key="ID"
              :props="{ label: 'name', children: 'children' }"
              :expand-on-click-node="false"
              :current-node-key="selectedDept.ID"
              highlight-current
              default-expand-all
              @node-click="onNodeClick"
            >
              <template #default="{ data }">
                <div
                  class="flex-1 flex items-center justify-between pr-2 group"
                >
                  <span class="truncate">
                    {{ data.name }}
                    <el-tag
                      v-if="!data.status"
                      size="small"
                      type="info"
                      class="ml-1"
                      >停用</el-tag
                    >
                  </span>
                  <span class="hidden group-hover:inline-flex items-center gap-2">
                    <el-icon
                      class="!text-blue-500"
                      title="新增子部门"
                      @click.stop="addDepartment(data.ID)"
                    >
                      <Plus />
                    </el-icon>
                    <el-icon
                      class="!text-blue-500"
                      title="编辑"
                      @click.stop="editDepartment(data)"
                    >
                      <Edit />
                    </el-icon>
                    <el-icon
                      class="!text-red-500"
                      title="删除"
                      @click.stop="deleteRow(data)"
                    >
                      <Delete />
                    </el-icon>
                  </span>
                </div>
              </template>
            </el-tree>
          </el-scrollbar>
        </div>
      </el-splitter-panel>
      <el-splitter-panel :min="300">
        <div class="gva-table-box !my-0 ml-2 h-full overflow-hidden">
          <org-member-assign
            kind="department"
            :target-id="selectedDept.ID || 0"
            :target-name="selectedDept.name || ''"
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
      <el-form ref="deptForm" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="上级部门" prop="parentId">
          <el-cascader
            v-model="form.parentId"
            style="width: 100%"
            :options="deptOptions"
            :props="{
              checkStrictly: true,
              label: 'name',
              value: 'ID',
              disabled: 'disabled',
              emitPath: false
            }"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="部门名称" prop="name">
          <el-input
            v-model="form.name"
            autocomplete="off"
            placeholder="请输入部门名称"
          />
        </el-form-item>
        <el-form-item label="负责人" prop="leaderId">
          <el-select
            v-model="form.leaderId"
            style="width: 100%"
            filterable
            clearable
            placeholder="请选择负责人"
          >
            <el-option
              v-for="u in userOptions"
              :key="u.ID"
              :label="u.nickName || u.userName"
              :value="u.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="联系电话">
          <el-input
            :model-value="selectedLeader?.phone || ''"
            disabled
            placeholder="选择负责人后自动带出"
          />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input
            :model-value="selectedLeader?.email || ''"
            disabled
            placeholder="选择负责人后自动带出"
          />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" :min="0" />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-switch v-model="form.status" active-text="启用" inactive-text="停用" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    getDepartmentList,
    createDepartment,
    updateDepartment,
    deleteDepartment
  } from '@/api/department'
  import { getUserList } from '@/api/user'
  import OrgMemberAssign from '@/view/superAdmin/orgMemberAssign.vue'
  import { Plus, Edit, Delete } from '@element-plus/icons-vue'
  import {
    ref,
    computed,
    onMounted,
    onActivated,
    onBeforeUnmount,
    nextTick
  } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { useAppStore } from '@/pinia'

  defineOptions({
    name: 'Department'
  })

  const appStore = useAppStore()
  const tableData = ref([])
  const deptOptions = ref([])
  const selectedDept = ref({})

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

  // 负责人从用户列表选取；联系电话/邮箱从所选负责人用户带出
  const userOptions = ref([])
  const selectedLeader = computed(() =>
    userOptions.value.find((u) => u.ID === form.value.leaderId)
  )
  const getUsers = async () => {
    const res = await getUserList({ page: 1, pageSize: 1000 })
    if (res.code === 0) {
      userOptions.value = res.data.list || []
    }
  }
  getUsers()

  const getTableData = async () => {
    const res = await getDepartmentList()
    if (res.code === 0) {
      tableData.value = res.data || []
      // 默认选中第一个根部门，保持右侧成员面板有内容
      if (!selectedDept.value.ID && tableData.value.length) {
        selectedDept.value = tableData.value[0]
      }
    }
  }
  getTableData()

  const onNodeClick = (data) => {
    selectedDept.value = data
  }

  // 组装级联选择器选项（编辑时禁用自身及子孙，防止形成环）
  const buildOptions = (source, disabled) => {
    const result = []
    source &&
      source.forEach((item) => {
        const selfDisabled = disabled || item.ID === form.value.ID
        const option = {
          ID: item.ID,
          name: item.name,
          disabled: selfDisabled
        }
        if (item.children && item.children.length) {
          option.children = buildOptions(item.children, selfDisabled)
        }
        result.push(option)
      })
    return result
  }
  const setOptions = () => {
    deptOptions.value = [
      { ID: 0, name: '顶级部门', disabled: false },
      ...buildOptions(tableData.value, false)
    ]
  }

  const formVisible = ref(false)
  const titleForm = ref('新增部门')
  const dialogType = ref('add')
  const deptForm = ref(null)
  const defaultForm = () => ({
    ID: 0,
    name: '',
    parentId: 0,
    sort: 0,
    leaderId: undefined,
    status: true
  })
  const form = ref(defaultForm())

  const rules = ref({
    name: [{ required: true, message: '请输入部门名称', trigger: 'blur' }],
    parentId: [{ required: true, message: '请选择上级部门', trigger: 'change' }]
  })

  const initForm = () => {
    deptForm.value && deptForm.value.resetFields()
    form.value = defaultForm()
  }
  const closeForm = () => {
    initForm()
    formVisible.value = false
  }

  const addDepartment = (parentId) => {
    initForm()
    titleForm.value = '新增部门'
    dialogType.value = 'add'
    form.value.parentId = parentId
    setOptions()
    formVisible.value = true
  }

  const editDepartment = (row) => {
    titleForm.value = '编辑部门'
    dialogType.value = 'edit'
    form.value = {
      ID: row.ID,
      name: row.name,
      parentId: row.parentId,
      sort: row.sort,
      leaderId: row.leaderId || undefined,
      status: !!row.status
    }
    setOptions()
    deptForm.value && deptForm.value.clearValidate()
    formVisible.value = true
  }

  const submitForm = () => {
    deptForm.value.validate(async (valid) => {
      if (!valid) return
      const api = dialogType.value === 'add' ? createDepartment : updateDepartment
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
    ElMessageBox.confirm('此操作将永久删除该部门, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async () => {
        const res = await deleteDepartment({ id: row.ID })
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '删除成功!' })
          if (selectedDept.value.ID === row.ID) selectedDept.value = {}
          getTableData()
        }
      })
      .catch(() => {
        ElMessage({ type: 'info', message: '已取消删除' })
      })
  }
</script>
