<template>
  <div class="authority">
    <warning-bar title="注：右上角头像下拉可切换角色" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addAuthority(0)"
          >新增角色</el-button
        >
      </div>
      <el-table
        :data="tableData"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
        row-key="authorityId"
        style="width: 100%"
      >
        <el-table-column label="角色ID" min-width="180" prop="authorityId" />
        <el-table-column
          align="left"
          label="角色名称"
          min-width="180"
          prop="authorityName"
        />
        <el-table-column align="left" label="数据权限" min-width="150">
          <template #default="scope">
            <div class="flex items-center gap-1">
              <el-dropdown
                trigger="click"
                @command="(cmd) => selectDataScope(scope.row, cmd)"
              >
                <el-button link type="primary">
                  {{ dataScopeLabel(scope.row.dataScope) }}
                  <el-icon class="ml-1"><arrow-down /></el-icon>
                </el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item
                      v-for="item in dataScopeOptions"
                      :key="item.value"
                      :command="item.value"
                      :disabled="scope.row.dataScope === item.value"
                    >
                      {{ item.label }}
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
              <el-button
                v-if="scope.row.dataScope === 5"
                icon="setting"
                type="primary"
                link
                title="配置部门集"
                @click="openDeptScopeDialog(scope.row)"
              />
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="560">
          <template #default="scope">
            <el-button
              icon="setting"
              type="primary"
              link
              @click="openDrawer(scope.row)"
              >设置权限</el-button
            >
            <el-button
              icon="user"
              type="primary"
              link
              @click="openAssignDrawer(scope.row)"
              >分配给用户</el-button
            >
            <el-button
              icon="plus"
              type="primary"
              link
              @click="addAuthority(scope.row.authorityId)"
              >新增子角色</el-button
            >
            <el-button
              icon="copy-document"
              type="primary"
              link
              @click="copyAuthorityFunc(scope.row)"
              >拷贝</el-button
            >
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editAuthority(scope.row)"
              >编辑</el-button
            >
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteAuth(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 新增角色弹窗 -->
    <el-drawer v-model="authorityFormVisible" :size="appStore.drawerSize" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">{{ authorityTitleForm }}</span>
          <div>
            <el-button @click="closeAuthorityForm">取 消</el-button>
            <el-button type="primary" @click="submitAuthorityForm"
              >确 定</el-button
            >
          </div>
        </div>
      </template>
      <el-form
        ref="authorityForm"
        :model="form"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item label="父级角色" prop="parentId">
          <el-cascader
            v-model="form.parentId"
            style="width: 100%"
            :disabled="dialogType === 'add'"
            :options="AuthorityOption"
            :props="{
              checkStrictly: true,
              label: 'authorityName',
              value: 'authorityId',
              disabled: 'disabled',
              emitPath: false
            }"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="角色ID" prop="authorityId">
          <el-input
            v-model="form.authorityId"
            :disabled="dialogType === 'edit'"
            autocomplete="off"
            maxlength="15"
          />
        </el-form-item>
        <el-form-item label="角色姓名" prop="authorityName">
          <el-input v-model="form.authorityName" autocomplete="off" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      v-if="drawer"
      v-model="drawer"
      :size="appStore.drawerSize"
      title="角色配置"
    >
      <el-tabs :before-leave="autoEnter" type="border-card">
        <el-tab-pane label="角色菜单">
          <Menus ref="menus" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
        <el-tab-pane label="角色api">
          <Apis ref="apis" :row="activeRow" @changeRow="changeRow" />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>

    <!-- 自定义部门集弹窗(数据权限第5档) -->
    <el-dialog
      v-model="deptScopeVisible"
      title="配置可见部门"
      width="420px"
      :close-on-click-modal="false"
      :show-close="false"
    >
      <el-scrollbar max-height="50vh">
        <el-tree
          ref="deptScopeTreeRef"
          :data="deptScopeTree"
          node-key="ID"
          show-checkbox
          check-strictly
          default-expand-all
          :props="{ label: 'name', children: 'children' }"
        />
      </el-scrollbar>
      <template #footer>
        <el-button @click="cancelDeptScope">取 消</el-button>
        <el-button
          type="primary"
          :loading="deptScopeSaving"
          @click="confirmDeptScope"
          >确 定</el-button
        >
      </template>
    </el-dialog>

    <!-- 分配给用户抽屉 -->
    <el-drawer
      v-model="assignDrawerVisible"
      :size="appStore.drawerSize"
      :show-close="false"
      destroy-on-close
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-base">分配用户 - {{ assignRow.authorityName }}</span>
          <div>
            <el-button @click="assignDrawerVisible = false">取 消</el-button>
            <el-button type="primary" :loading="assignSubmitting" @click="confirmAssign">确 定</el-button>
          </div>
        </div>
      </template>
      <warning-bar title="注：保存时将全量覆盖该角色的用户关联关系；若用户仅剩此一个角色，移除后其主角色保持不变" />
      <div class="gva-search-box">
        <el-form :inline="true" :model="userSearchInfo">
          <el-form-item label="用户名">
            <el-input v-model="userSearchInfo.username" placeholder="请输入用户名" />
          </el-form-item>
          <el-form-item label="昵称">
            <el-input v-model="userSearchInfo.nickName" placeholder="请输入昵称" />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="searchUserData">查 询</el-button>
            <el-button icon="refresh" @click="resetUserSearch">重 置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-table
        ref="userTableRef"
        v-loading="assignLoading"
        :data="userTableData"
        row-key="ID"
        :default-sort="{ prop: 'ID', order: 'descending' }"
        @sort-change="sortChange"
        @select="handleSelect"
        @select-all="handleSelectAll"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="ID" prop="ID" width="80" sortable="custom" />
        <el-table-column label="用户名" prop="userName" min-width="120" />
        <el-table-column label="昵称" prop="nickName" min-width="120" />
      </el-table>
      <div class="flex justify-center mt-4">
        <el-pagination
          :current-page="userSearchInfo.page"
          :page-size="userSearchInfo.pageSize"
          :page-sizes="[10, 20, 50]"
          :total="userTotal"
          layout="total, sizes, prev, pager, next"
          @current-change="handleUserPageChange"
          @size-change="handleUserSizeChange"
        />
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    getAuthorityList,
    deleteAuthority,
    createAuthority,
    updateAuthority,
    copyAuthority,
    getUsersByAuthorityId,
    setRoleUsers,
    setDataScope,
    getDataScopeDepts
  } from '@/api/authority'
  import { getUserList } from '@/api/user'
  import { getDepartmentList } from '@/api/department'

  import Menus from '@/view/superAdmin/authority/components/menus.vue'
  import Apis from '@/view/superAdmin/authority/components/apis.vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'

  import { ref, nextTick } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { useAppStore } from "@/pinia"
  import { toSQLLine } from '@/utils/stringFun'

  defineOptions({
    name: 'Authority'
  })

  const mustUint = (rule, value, callback) => {
    if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
      return callback(new Error('请输入正整数'))
    }
    return callback()
  }

  const AuthorityOption = ref([
    {
      authorityId: 0,
      authorityName: '根角色/严格模式下为当前角色'
    }
  ])
  const drawer = ref(false)
  const dialogType = ref('add')
  const activeRow = ref({})
  const appStore = useAppStore()

  const authorityTitleForm = ref('新增角色')
  const authorityFormVisible = ref(false)
  const apiDialogFlag = ref(false)
  const copyForm = ref({})

  // 数据权限档位(挂在角色上)
  const dataScopeOptions = [
    { value: 1, label: '全部数据' },
    { value: 3, label: '本部门' },
    { value: 2, label: '本部门及以下' },
    { value: 4, label: '仅本人' },
    { value: 5, label: '自定义部门' }
  ]

  const form = ref({
    authorityId: 0,
    authorityName: '',
    parentId: 0,
    dataScope: 1
  })
  const rules = ref({
    authorityId: [
      { required: true, message: '请输入角色ID', trigger: 'blur' },
      { validator: mustUint, trigger: 'blur', message: '必须为正整数' }
    ],
    authorityName: [
      { required: true, message: '请输入角色名', trigger: 'blur' }
    ],
    parentId: [{ required: true, message: '请选择父角色', trigger: 'blur' }]
  })

  const tableData = ref([])

  // 查询
  const getTableData = async () => {
    const table = await getAuthorityList()
    if (table.code === 0) {
      tableData.value = table.data
      rememberScopes(tableData.value)
    }
  }

  getTableData()

  // 记录每个角色当前档位, 用于自定义部门弹窗取消/失败时回退
  const prevScope = {}
  const rememberScopes = (rows) => {
    rows &&
      rows.forEach((row) => {
        prevScope[row.authorityId] = row.dataScope
        if (row.children && row.children.length) rememberScopes(row.children)
      })
  }

  // 数据权限: 列表内下拉即时保存; 选"自定义部门"时先弹部门选择
  const changeDataScope = async (row) => {
    if (row.dataScope === 5) {
      openDeptScopeDialog(row)
      return
    }
    const res = await setDataScope({
      authorityId: row.authorityId,
      dataScope: row.dataScope
    })
    if (res.code === 0) {
      prevScope[row.authorityId] = row.dataScope
      ElMessage.success('数据权限设置成功')
    } else {
      row.dataScope = prevScope[row.authorityId]
    }
  }

  // 档位值 -> 中文标签, 用于列表内以文字(而非常驻下拉框)展示当前数据权限
  const dataScopeLabel = (value) =>
    dataScopeOptions.find((item) => item.value === value)?.label || '未设置'

  // 点击下拉菜单切换档位: 复用 changeDataScope 的即时保存与失败回退逻辑
  const selectDataScope = (row, value) => {
    if (row.dataScope === value) return
    row.dataScope = value
    changeDataScope(row)
  }

  // 自定义部门集弹窗
  const deptScopeVisible = ref(false)
  const deptScopeRow = ref({})
  const deptScopeTree = ref([])
  const deptScopeTreeRef = ref(null)
  const deptScopeSaving = ref(false)

  const openDeptScopeDialog = async (row) => {
    deptScopeRow.value = row
    deptScopeVisible.value = true
    const [treeRes, checkedRes] = await Promise.all([
      getDepartmentList(),
      getDataScopeDepts(row.authorityId)
    ])
    if (treeRes.code === 0) {
      deptScopeTree.value = treeRes.data || []
    }
    await nextTick()
    if (checkedRes.code === 0 && deptScopeTreeRef.value) {
      deptScopeTreeRef.value.setCheckedKeys(checkedRes.data || [])
    }
  }

  const cancelDeptScope = () => {
    // 取消 → 回退档位
    deptScopeRow.value.dataScope = prevScope[deptScopeRow.value.authorityId]
    deptScopeVisible.value = false
  }

  const confirmDeptScope = async () => {
    const deptIds = deptScopeTreeRef.value
      ? deptScopeTreeRef.value.getCheckedKeys()
      : []
    if (!deptIds.length) {
      ElMessage.warning('请至少选择一个部门')
      return
    }
    deptScopeSaving.value = true
    const res = await setDataScope({
      authorityId: deptScopeRow.value.authorityId,
      dataScope: 5,
      deptIds
    })
    if (res.code === 0) {
      prevScope[deptScopeRow.value.authorityId] = 5
      ElMessage.success('数据权限设置成功')
      deptScopeVisible.value = false
    }
    deptScopeSaving.value = false
  }

  const changeRow = (key, value) => {
    activeRow.value[key] = value
  }
  const menus = ref(null)
  const apis = ref(null)
  const autoEnter = (activeName, oldActiveName) => {
    const paneArr = [menus, apis]
    if (oldActiveName) {
      if (paneArr[oldActiveName].value.needConfirm) {
        paneArr[oldActiveName].value.enterAndNext()
        paneArr[oldActiveName].value.needConfirm = false
      }
    }
  }
  // 拷贝角色
  const copyAuthorityFunc = (row) => {
    setOptions()
    authorityTitleForm.value = '拷贝角色'
    dialogType.value = 'copy'
    for (const k in form.value) {
      form.value[k] = row[k]
    }
    copyForm.value = row
    authorityFormVisible.value = true
  }
  const openDrawer = (row) => {
    drawer.value = true
    activeRow.value = row
  }
  // 删除角色
  const deleteAuth = (row) => {
    ElMessageBox.confirm('此操作将永久删除该角色, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
      .then(async () => {
        const res = await deleteAuthority({ authorityId: row.authorityId })
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '删除成功!'
          })

          getTableData()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '已取消删除'
        })
      })
  }
  // 初始化表单
  const authorityForm = ref(null)
  const initForm = () => {
    if (authorityForm.value) {
      authorityForm.value.resetFields()
    }
    form.value = {
      authorityId: 0,
      authorityName: '',
      parentId: 0,
      dataScope: 1
    }
  }
  // 关闭窗口
  const closeAuthorityForm = () => {
    initForm()
    authorityFormVisible.value = false
    apiDialogFlag.value = false
  }
  // 确定弹窗

  const submitAuthorityForm = () => {
    authorityForm.value.validate(async (valid) => {
      if (valid) {
        form.value.authorityId = Number(form.value.authorityId)
        switch (dialogType.value) {
          case 'add':
            {
              const res = await createAuthority(form.value)
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '添加成功!'
                })
                getTableData()
                closeAuthorityForm()
              }
            }
            break
          case 'edit':
            {
              const res = await updateAuthority(form.value)
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '添加成功!'
                })
                getTableData()
                closeAuthorityForm()
              }
            }
            break
          case 'copy': {
            const data = {
              authority: {
                authorityId: 0,
                authorityName: '',
                parentId: 0
              },
              oldAuthorityId: 0
            }
            data.authority.authorityId = form.value.authorityId
            data.authority.authorityName = form.value.authorityName
            data.authority.parentId = form.value.parentId
            data.authority.dataScope = form.value.dataScope
            data.oldAuthorityId = copyForm.value.authorityId
            const res = await copyAuthority(data)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '复制成功！'
              })
              getTableData()
            }
          }
        }

        initForm()
        authorityFormVisible.value = false
      }
    })
  }
  const setOptions = () => {
    AuthorityOption.value = [
      {
        authorityId: 0,
        authorityName: '根角色(严格模式下为当前用户角色)'
      }
    ]
    setAuthorityOptions(tableData.value, AuthorityOption.value, false)
  }
  const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
    AuthorityData &&
      AuthorityData.forEach((item) => {
        if (item.children && item.children.length) {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            disabled: disabled || item.authorityId === form.value.authorityId,
            children: []
          }
          setAuthorityOptions(
            item.children,
            option.children,
            disabled || item.authorityId === form.value.authorityId
          )
          optionsData.push(option)
        } else {
          const option = {
            authorityId: item.authorityId,
            authorityName: item.authorityName,
            disabled: disabled || item.authorityId === form.value.authorityId
          }
          optionsData.push(option)
        }
      })
  }
  // 增加角色
  const addAuthority = (parentId) => {
    initForm()
    authorityTitleForm.value = '新增角色'
    dialogType.value = 'add'
    form.value.parentId = parentId
    setOptions()
    authorityFormVisible.value = true
  }
  // 编辑角色
  const editAuthority = (row) => {
    setOptions()
    authorityTitleForm.value = '编辑角色'
    dialogType.value = 'edit'
    for (const key in form.value) {
      form.value[key] = row[key]
    }
    if (!form.value.dataScope) form.value.dataScope = 1
    setOptions()
    authorityForm.value && authorityForm.value.clearValidate()
    authorityFormVisible.value = true
  }

  // 分配给用户
  const assignDrawerVisible = ref(false)
  const assignRow = ref({})
  const userTableData = ref([])
  const userTotal = ref(0)
  const userSearchInfo = ref({ page: 1, pageSize: 10, username: '', nickName: '', orderKey: 'id', desc: true })
  const assignLoading = ref(false)
  const assignSubmitting = ref(false)
  const userTableRef = ref(null)

  const selectedUserIds = ref(new Set())

  const openAssignDrawer = async (row) => {
    assignRow.value = row
    userSearchInfo.value = { page: 1, pageSize: 10, username: '', nickName: '' }
    selectedUserIds.value = new Set()
    assignDrawerVisible.value = true
    const res = await getUsersByAuthorityId(row.authorityId)
    if (res.code === 0 && res.data) {
      selectedUserIds.value = new Set(res.data)
    }
    getUserData()
  }

  const getUserData = async () => {
    assignLoading.value = true
    const res = await getUserList(userSearchInfo.value)
    if (res.code === 0) {
      userTableData.value = res.data.list
      userTotal.value = res.data.total
      await nextTick()
      userTableData.value.forEach((user) => {
        userTableRef.value && userTableRef.value.toggleRowSelection(user, selectedUserIds.value.has(user.ID))
      })
    }
    assignLoading.value = false
  }

  const handleSelect = (selection, row) => {
    if (selection.some(u => u.ID === row.ID)) {
      selectedUserIds.value.add(row.ID)
    } else {
      selectedUserIds.value.delete(row.ID)
    }
  }

  const handleSelectAll = (selection) => {
    const selectedIds = new Set(selection.map(u => u.ID))
    userTableData.value.forEach((user) => {
      if (selectedIds.has(user.ID)) {
        selectedUserIds.value.add(user.ID)
      } else {
        selectedUserIds.value.delete(user.ID)
      }
    })
  }

  const sortChange = ({ prop, order }) => {
    if (prop) {
      userSearchInfo.value.orderKey = prop === 'ID' ? 'id' : toSQLLine(prop)
      userSearchInfo.value.desc = order === 'descending'
    }
    getUserData()
  }

  const searchUserData = () => {
    userSearchInfo.value.page = 1
    getUserData()
  }

  const resetUserSearch = () => {
    userSearchInfo.value = { page: 1, pageSize: 10, username: '', nickName: '' }
    getUserData()
  }

  const handleUserPageChange = (page) => {
    userSearchInfo.value.page = page
    getUserData()
  }

  const handleUserSizeChange = (size) => {
    userSearchInfo.value.pageSize = size
    userSearchInfo.value.page = 1
    getUserData()
  }

  const confirmAssign = async () => {
    assignSubmitting.value = true
    try {
      const res = await setRoleUsers({
        authorityId: assignRow.value.authorityId,
        userIds: [...selectedUserIds.value]
      })
      if (res.code === 0) {
        ElMessage({ type: 'success', message: '分配成功!' })
        assignDrawerVisible.value = false
      }
    } catch {
      ElMessage({ type: 'error', message: '分配失败，请重试' })
    }
    assignSubmitting.value = false
  }
</script>

<style lang="scss">
  .authority {
    .el-input-number {
      margin-left: 15px;
      span {
        display: none;
      }
    }
  }
</style>
