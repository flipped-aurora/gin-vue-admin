<template>
  <div class="h-full flex flex-col overflow-hidden">
    <div
      v-if="!targetId"
      class="flex-1 flex items-center justify-center text-gray-400"
    >
      请选择左侧{{ kind === 'department' ? '部门' : '岗位' }}后分配成员
    </div>
    <template v-else>
      <div class="flex justify-between items-center flex-none mb-3">
        <span class="font-bold text-base truncate"
          >「{{ targetName }}」成员</span
        >
        <el-button type="primary" :loading="saving" @click="save"
          >保存成员</el-button
        >
      </div>
      <div class="gva-search-box !mb-2 flex-none">
        <el-form :inline="true" :model="search">
          <el-form-item label="用户名">
            <el-input
              v-model="search.username"
              placeholder="用户名"
              @keyup.enter="query"
            />
          </el-form-item>
          <el-form-item label="昵称">
            <el-input
              v-model="search.nickName"
              placeholder="昵称"
              @keyup.enter="query"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="query"
              >查询</el-button
            >
            <el-button icon="refresh" @click="reset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="flex-1 min-h-0">
        <el-table
          ref="tableRef"
          v-loading="loading"
          :data="users"
          row-key="ID"
          height="100%"
          @select="onSelect"
          @select-all="onSelectAll"
        >
          <el-table-column type="selection" width="55" reserve-selection />
          <el-table-column label="ID" prop="ID" width="70" />
          <el-table-column label="用户名" prop="userName" min-width="120" />
          <el-table-column label="昵称" prop="nickName" min-width="120" />
        </el-table>
      </div>
      <div class="flex justify-end mt-2 flex-none">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 20, 50]"
          :total="total"
          layout="total, sizes, prev, pager, next"
          @current-change="onPage"
          @size-change="onSize"
        />
      </div>
    </template>
  </div>
</template>

<script setup>
  import { ref, watch, nextTick } from 'vue'
  import { ElMessage } from 'element-plus'
  import { getUserList } from '@/api/user'
  import { getDepartmentUsers, setDepartmentUsers } from '@/api/department'
  import { getPositionUsers, setPositionUsers } from '@/api/position'

  const props = defineProps({
    kind: { type: String, required: true }, // 'department' | 'position'
    targetId: { type: Number, default: 0 },
    targetName: { type: String, default: '' }
  })

  const tableRef = ref(null)
  const users = ref([])
  const page = ref(1)
  const pageSize = ref(10)
  const total = ref(0)
  const loading = ref(false)
  const saving = ref(false)
  const search = ref({ username: '', nickName: '' })
  const selected = ref(new Set())

  const getMembers = () =>
    props.kind === 'department'
      ? getDepartmentUsers({ deptId: props.targetId })
      : getPositionUsers({ positionId: props.targetId })

  const setMembers = () =>
    props.kind === 'department'
      ? setDepartmentUsers({
          deptId: props.targetId,
          userIds: [...selected.value]
        })
      : setPositionUsers({
          positionId: props.targetId,
          userIds: [...selected.value]
        })

  const syncChecks = () => {
    users.value.forEach((u) => {
      tableRef.value &&
        tableRef.value.toggleRowSelection(u, selected.value.has(u.ID))
    })
  }

  const loadUsers = async () => {
    loading.value = true
    const res = await getUserList({
      page: page.value,
      pageSize: pageSize.value,
      ...search.value
    })
    if (res.code === 0) {
      users.value = res.data.list || []
      total.value = res.data.total
      await nextTick()
      syncChecks()
    }
    loading.value = false
  }

  const load = async () => {
    if (!props.targetId) return
    selected.value = new Set()
    const res = await getMembers()
    if (res.code === 0 && res.data) selected.value = new Set(res.data)
    page.value = 1
    await loadUsers()
  }

  watch(() => props.targetId, load, { immediate: true })

  const onSelect = (selection, row) => {
    if (selection.some((u) => u.ID === row.ID)) selected.value.add(row.ID)
    else selected.value.delete(row.ID)
  }
  const onSelectAll = (selection) => {
    const ids = new Set(selection.map((u) => u.ID))
    users.value.forEach((u) => {
      if (ids.has(u.ID)) selected.value.add(u.ID)
      else selected.value.delete(u.ID)
    })
  }
  const onPage = (p) => {
    page.value = p
    loadUsers()
  }
  const onSize = (s) => {
    pageSize.value = s
    page.value = 1
    loadUsers()
  }
  const query = () => {
    page.value = 1
    loadUsers()
  }
  const reset = () => {
    search.value = { username: '', nickName: '' }
    page.value = 1
    loadUsers()
  }

  const save = async () => {
    saving.value = true
    const res = await setMembers()
    if (res.code === 0) ElMessage.success('成员保存成功')
    saving.value = false
  }
</script>
