<template>
  <div class="notification-management">
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="标题">
          <el-input v-model="searchInfo.title" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="searchInfo.type" placeholder="请选择" clearable>
            <el-option label="系统通知" value="system" />
            <el-option label="业务通知" value="business" />
            <el-option label="警告通知" value="warning" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
            <el-option label="草稿" value="draft" />
            <el-option label="已发送" value="sent" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >删除</el-button
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
        <el-table-column align="left" label="日期" prop="CreatedAt" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.createdAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="标题" prop="title" />
        <el-table-column align="left" label="类型" prop="type" width="120">
          <template #default="scope">
            <el-tag :type="getTypeTagType(scope.row.type)">{{
              getTypeLabel(scope.row.type)
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="目标类型"
          prop="targetType"
          width="120"
        >
          <template #default="scope">
            <span>{{ getTargetTypeLabel(scope.row.targetType) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="getStatusTagType(scope.row.status)">{{
              getStatusLabel(scope.row.status)
            }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="300">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="view"
              size="small"
              @click="getDetails(scope.row)"
              >查看</el-button
            >
            <el-button
              type="primary"
              link
              icon="edit"
              size="small"
              @click="updateNotificationFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              size="small"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
            <el-button
              v-if="scope.row.status === 'draft'"
              type="success"
              link
              icon="promotion"
              size="small"
              @click="sendNotificationFunc(scope.row)"
              >发送</el-button
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
      destroy-on-close
      size="800"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            type === 'create' ? '添加' : type === 'update' ? '修改' : '查看'
          }}</span>
          <div>
            <el-button
              v-if="type !== 'create'"
              type="primary"
              @click="updateNotificationFunc(formData)"
              >变更</el-button
            >
            <el-button
              v-if="type === 'create'"
              type="primary"
              @click="createNotificationFunc"
              >添加</el-button
            >
            <el-button type="primary" @click="closeDialog">取消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="标题:" prop="title">
          <el-input
            v-model="formData.title"
            :clearable="true"
            placeholder="请输入标题"
          />
        </el-form-item>
        <el-form-item label="内容:" prop="content">
          <el-input
            v-model="formData.content"
            :clearable="true"
            placeholder="请输入内容"
            type="textarea"
            :rows="4"
          />
        </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-select
            v-model="formData.type"
            placeholder="请选择类型"
            style="width: 100%"
          >
            <el-option label="系统通知" value="system" />
            <el-option label="业务通知" value="business" />
            <el-option label="警告通知" value="warning" />
          </el-select>
        </el-form-item>
        <el-form-item label="目标类型:" prop="targetType">
          <el-select
            v-model="formData.targetType"
            placeholder="请选择目标类型"
            style="width: 100%"
            @change="onTargetTypeChange"
          >
            <el-option label="全部用户" value="all" />
            <el-option label="指定用户" value="users" />
            <el-option label="指定角色" value="roles" />
          </el-select>
        </el-form-item>
        <el-form-item
          v-if="formData.targetType === 'users'"
          label="目标用户:"
          prop="targetIds"
        >
          <el-select
            v-model="formData.targetIds"
            placeholder="请选择用户"
            multiple
            style="width: 100%"
          >
            <el-option
              v-for="user in userList"
              :key="user.ID"
              :label="user.nickName"
              :value="user.ID"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          v-if="formData.targetType === 'roles'"
          label="目标角色:"
          prop="targetIds"
        >
          <el-select
            v-model="formData.targetIds"
            placeholder="请选择角色"
            multiple
            style="width: 100%"
          >
            <el-option
              v-for="role in roleList"
              :key="role.authorityId"
              :label="role.authorityName"
              :value="role.authorityId"
            />
          </el-select>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createNotification,
    deleteNotification,
    updateNotification,
    getNotificationById,
    getNotificationList,
    publishNotification
  } from '@/api/notice'
  import { getUserList } from '@/api/user'
  import { getAuthorityList } from '@/api/authority'
  import { formatDate } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, reactive } from 'vue'

  defineOptions({
    name: 'NotificationManagement'
  })

  // 响应式数据
  const formData = ref({
    title: '',
    content: '',
    type: '',
    targetType: '',
    targetIds: []
  })

  const searchInfo = ref({
    title: '',
    type: '',
    status: ''
  })

  const type = ref('')
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const dialogFormVisible = ref(false)
  const multipleSelection = ref([])
  const userList = ref([])
  const roleList = ref([])

  // 表单验证规则
  const rule = reactive({
    title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
    content: [{ required: true, message: '请输入内容', trigger: 'blur' }],
    type: [{ required: true, message: '请选择类型', trigger: 'change' }],
    targetType: [
      { required: true, message: '请选择目标类型', trigger: 'change' }
    ]
  })

  const elFormRef = ref()

  // 获取列表数据
  const getTableData = async () => {
    const table = await getNotificationList({
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

  // 获取用户列表
  const getUserListData = async () => {
    const res = await getUserList({ page: 1, pageSize: 999 })
    if (res.code === 0) {
      userList.value = res.data.list
    }
  }

  // 获取角色列表
  const getRoleListData = async () => {
    const res = await getAuthorityList({ page: 1, pageSize: 999 })
    if (res.code === 0) {
      roleList.value = res.data.list
    }
  }

  // 初始化数据
  getUserListData()
  getRoleListData()

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 查询
  const onSubmit = () => {
    page.value = 1
    pageSize.value = 10
    getTableData()
  }

  // 重置
  const onReset = () => {
    searchInfo.value = {
      title: '',
      type: '',
      status: ''
    }
    getTableData()
  }

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
      deleteNotificationFunc(row)
    })
  }

  // 批量删除
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
      const res = await deleteNotification({ IDs })
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

  // 弹窗控制标记
  const initForm = () => {
    elFormRef.value?.resetFields()
    formData.value = {
      title: '',
      content: '',
      type: '',
      targetType: '',
      targetIds: []
    }
  }

  // 关闭弹窗
  const closeDialog = () => {
    dialogFormVisible.value = false
    initForm()
  }

  // 弹窗打开的标记
  const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
  }

  // 创建
  const createNotificationFunc = async () => {
    elFormRef.value?.validate(async (valid) => {
      if (!valid) return
      const res = await createNotification(formData.value)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建成功'
        })
        closeDialog()
        getTableData()
      }
    })
  }

  // 更新
  const updateNotificationFunc = async (row) => {
    if (type.value === 'update') {
      elFormRef.value?.validate(async (valid) => {
        if (!valid) return
        const res = await updateNotification(formData.value)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '更新成功'
          })
          closeDialog()
          getTableData()
        }
      })
    } else {
      const res = await getNotificationById(row.id)
      type.value = 'update'
      if (res.code === 0) {
        formData.value = res.data
        formData.value.targetIds = res.data.targetIds.split(',') || []
        dialogFormVisible.value = true
      }
    }
  }

  // 删除
  const deleteNotificationFunc = async (row) => {
    const res = await deleteNotification(row.id)
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

  // 发送通知
  const sendNotificationFunc = async (row) => {
    ElMessageBox.confirm('确定要发送这条通知吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'info'
    }).then(async () => {
      console.log(row)
      const res = await publishNotification(row.id)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '发送成功'
        })
        getTableData()
      }
    })
  }

  // 查看详情
  const getDetails = async (row) => {
    const res = await getNotificationById(row.id)
    if (res.code === 0) {
      formData.value = res.data
      formData.value.targetIds = res.data.targetIds.split(',') || []
      type.value = 'look'
      dialogFormVisible.value = true
    }
  }

  // 目标类型改变
  const onTargetTypeChange = () => {
    formData.value.targetIds = []
  }

  // 获取类型标签类型
  const getTypeTagType = (type) => {
    const typeMap = {
      system: 'primary',
      business: 'success',
      warning: 'warning'
    }
    return typeMap[type] || ''
  }

  // 获取类型标签
  const getTypeLabel = (type) => {
    const typeMap = {
      system: '系统通知',
      business: '业务通知',
      warning: '警告通知'
    }
    return typeMap[type] || type
  }

  // 获取目标类型标签
  const getTargetTypeLabel = (targetType) => {
    const targetTypeMap = {
      all: '全部用户',
      users: '指定用户',
      roles: '指定角色'
    }
    return targetTypeMap[targetType] || targetType
  }

  // 获取状态标签类型
  const getStatusTagType = (status) => {
    const statusMap = {
      draft: 'info',
      published: 'success'
    }
    return statusMap[status] || ''
  }

  // 获取状态标签
  const getStatusLabel = (status) => {
    const statusMap = {
      draft: '草稿',
      published: '已发送'
    }
    return statusMap[status] || status
  }
</script>

<style></style>
