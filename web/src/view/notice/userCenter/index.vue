<template>
  <div class="user-notification-center">
    <!-- 页面头部 -->
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="通知类型">
          <el-select v-model="searchInfo.type" placeholder="请选择" clearable>
            <el-option label="系统通知" value="system" />
            <el-option label="业务通知" value="business" />
            <el-option label="警告通知" value="warning" />
          </el-select>
        </el-form-item>
        <el-form-item label="阅读状态">
          <el-select v-model="searchInfo.isRead" placeholder="请选择" clearable>
            <el-option label="未读" :value="false" />
            <el-option label="已读" :value="true" />
          </el-select>
        </el-form-item>
        <el-form-item label="优先级">
          <el-select
            v-model="searchInfo.priority"
            placeholder="请选择"
            clearable
          >
            <el-option label="低" value="low" />
            <el-option label="普通" value="normal" />
            <el-option label="高" value="high" />
            <el-option label="紧急" value="urgent" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 通知列表 -->
    <div class="gva-table-box">
      <div class="gva-btn-list mb-4">
        <el-button
          type="success"
          icon="check"
          :disabled="selectedNotifications.length === 0"
          @click="batchMarkAsRead"
        >
          批量已读
        </el-button>
        <el-button
          type="danger"
          icon="delete"
          :disabled="selectedNotifications.length === 0"
          @click="batchDeleteNotifications"
        >
          批量删除
        </el-button>
      </div>

      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        @selection-change="handleSelectionChange"
        @row-click="handleRowClick"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column prop="title" label="标题" min-width="200">
          <template #default="scope">
            <div
              class="notification-title"
              :class="{ unread: !scope.row.isRead }"
            >
              <el-icon v-if="!scope.row.isRead" class="unread-dot">
                <CircleCheck />
              </el-icon>
              {{ scope.row.title }}
            </div>
          </template>
        </el-table-column>

        <el-table-column prop="type" label="类型" width="120">
          <template #default="scope">
            <el-tag :type="getTypeTagType(scope.row.type)" size="small">
              {{ getTypeLabel(scope.row.type) }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column prop="priority" label="优先级" width="100">
          <template #default="scope">
            <el-tag :type="getPriorityTagType(scope.row.priority)" size="small">
              {{ getPriorityLabel(scope.row.priority) }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column prop="isRead" label="状态" width="80">
          <template #default="scope">
            <el-tag
              :type="scope.row.isRead ? 'success' : 'warning'"
              size="small"
            >
              {{ scope.row.isRead ? '已读' : '未读' }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column prop="createdAt" label="发送时间" width="180">
          <template #default="scope">
            {{ formatDate(scope.row.createdAt) }}
          </template>
        </el-table-column>

        <el-table-column prop="readTime" label="阅读时间" width="180">
          <template #default="scope">
            {{ scope.row.readTime ? formatDate(scope.row.readTime) : '-' }}
          </template>
        </el-table-column>

        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button
              size="small"
              type="primary"
              @click.stop="viewNotification(scope.row)"
            >
              查看
            </el-button>
            <el-button
              v-if="!scope.row.isRead"
              size="small"
              type="success"
              @click.stop="markAsRead(scope.row)"
            >
              标记已读
            </el-button>
            <el-button
              size="small"
              type="danger"
              @click.stop="deleteNotificationHandler(scope.row)"
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
          :page-sizes="[10, 25, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 通知详情弹窗 -->
    <el-dialog
      v-model="detailDialogVisible"
      title="通知详情"
      width="60%"
      :before-close="closeDetailDialog"
    >
      <div v-if="currentNotification" class="notification-detail">
        <div class="detail-header">
          <h3>{{ currentNotification.title }}</h3>
          <div class="detail-meta">
            <el-tag
              :type="getTypeTagType(currentNotification.type)"
              size="small"
            >
              {{ getTypeLabel(currentNotification.type) }}
            </el-tag>
            <el-tag
              :type="getPriorityTagType(currentNotification.priority)"
              size="small"
              class="ml-2"
            >
              {{ getPriorityLabel(currentNotification.priority) }}
            </el-tag>
            <span class="detail-time ml-4">
              发送时间：{{ formatDate(currentNotification.createdAt) }}
            </span>
          </div>
        </div>

        <div class="detail-content">
          <div class="content-text" v-html="currentNotification.content"></div>
        </div>
      </div>

      <template #footer>
        <span class="dialog-footer">
          <el-button
            v-if="currentNotification && !currentNotification.isRead"
            type="success"
            @click="markAsReadAndClose"
          >
            标记已读
          </el-button>
          <el-button type="danger" @click="deleteAndClose"> 删除 </el-button>
          <el-button @click="closeDetailDialog">关闭</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
  import { ref, reactive, onMounted, nextTick } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { CircleCheck } from '@element-plus/icons-vue'
  import {
    getUserNotificationList,
    markNotificationAsRead,
    getUnreadNotificationCount,
    deleteNotification as deleteNotificationApi
  } from '@/api/notice.js'
  import { formatTimeToStr } from '@/utils/date'

  defineOptions({
    name: 'UserNotificationCenter'
  })

  // 响应式数据
  const page = ref(1)
  const pageSize = ref(10)
  const total = ref(0)
  const tableData = ref([])
  const loading = ref(false)
  const selectedNotifications = ref([])
  const detailDialogVisible = ref(false)
  const currentNotification = ref(null)

  // 搜索条件
  const searchInfo = reactive({
    type: '',
    isRead: null,
    priority: ''
  })

  // 统计信息
  const stats = reactive({
    total: 0,
    unread: 0,
    read: 0
  })

  // 获取通知列表
  const getTableData = async () => {
    loading.value = true
    try {
      const params = {
        page: page.value,
        pageSize: pageSize.value,
        ...searchInfo
      }

      const res = await getUserNotificationList(params)
      if (res.code === 0) {
        tableData.value = res.data.list || []
        total.value = res.data.total || 0
      }
    } catch (error) {
      console.error('获取通知列表失败:', error)
      ElMessage.error('获取通知列表失败')
    } finally {
      loading.value = false
    }
  }

  // 获取统计信息
  const getStats = async () => {
    try {
      const res = await getUnreadNotificationCount()
      if (res.code === 0) {
        stats.total = res.data.total || 0
        stats.unread = res.data.unread || 0
        stats.read = res.data.read || 0
      }
    } catch (error) {
      console.error('获取统计信息失败:', error)
    }
  }

  // 搜索
  const onSubmit = () => {
    page.value = 1
    getTableData()
  }

  // 重置搜索
  const onReset = () => {
    searchInfo.type = ''
    searchInfo.isRead = null
    searchInfo.priority = ''
    page.value = 1
    getTableData()
  }

  // 分页处理
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  const handleSizeChange = (val) => {
    pageSize.value = val
    page.value = 1
    getTableData()
  }

  // 选择处理
  const handleSelectionChange = (val) => {
    selectedNotifications.value = val
  }

  // 行点击处理
  const handleRowClick = (row) => {
    viewNotification(row)
  }

  // 查看通知详情
  const viewNotification = (row) => {
    currentNotification.value = row
    detailDialogVisible.value = true

    // 如果是未读通知，自动标记为已读
    if (!row.isRead) {
      nextTick(() => {
        markAsRead(row, false)
      })
    }
  }

  // 关闭详情弹窗
  const closeDetailDialog = () => {
    detailDialogVisible.value = false
    currentNotification.value = null
  }

  // 标记单个通知为已读
  const markAsRead = async (row, showMessage = true) => {
    try {
      const res = await markNotificationAsRead({
        notificationIds: [row.notificationId]
      })

      if (res.code === 0) {
        row.isRead = true
        row.readTime = new Date()

        if (showMessage) {
          ElMessage.success('标记已读成功')
        }

        // 更新统计信息
        getStats()
      }
    } catch (error) {
      console.error('标记已读失败:', error)
      ElMessage.error('标记已读失败')
    }
  }

  // 标记已读并关闭弹窗
  const markAsReadAndClose = async () => {
    if (currentNotification.value) {
      await markAsRead(currentNotification.value)
      closeDetailDialog()
    }
  }

  // 批量标记已读
  const batchMarkAsRead = async () => {
    if (selectedNotifications.value.length === 0) {
      ElMessage.warning('请选择要标记的通知')
      return
    }

    const unreadNotifications = selectedNotifications.value.filter(
      (item) => !item.isRead
    )
    if (unreadNotifications.length === 0) {
      ElMessage.warning('所选通知都已是已读状态')
      return
    }

    try {
      const notificationIds = unreadNotifications.map(
        (item) => item.notificationId
      )
      const res = await markNotificationAsRead({ notificationIds })

      if (res.code === 0) {
        // 更新本地数据
        unreadNotifications.forEach((item) => {
          item.isRead = true
          item.readTime = new Date()
        })

        ElMessage.success(`成功标记 ${unreadNotifications.length} 条通知为已读`)
        getStats()
      }
    } catch (error) {
      console.error('批量标记已读失败:', error)
      ElMessage.error('批量标记已读失败')
    }
  }

  // 删除单个通知
  const deleteNotificationHandler = async (row) => {
    ElMessageBox.confirm(
      '确定要删除这条通知吗？如果此通知只针对您一人，删除后将无法恢复。',
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    ).then(async () => {
      try {
        const res = await deleteNotificationApi(row.notificationId)
        if (res.code === 0) {
          ElMessage.success('删除成功')
          getTableData()
          getStats()
        }
      } catch (error) {
        console.error('删除通知失败:', error)
        ElMessage.error('删除通知失败')
      }
    })
  }

  // 删除并关闭弹窗
  const deleteAndClose = async () => {
    if (currentNotification.value) {
      await deleteNotificationHandler(currentNotification.value)
      closeDetailDialog()
    }
  }

  // 批量删除通知
  const batchDeleteNotifications = async () => {
    if (selectedNotifications.value.length === 0) {
      ElMessage.warning('请选择要删除的通知')
      return
    }

    ElMessageBox.confirm(
      `确定要删除选中的 ${selectedNotifications.value.length} 条通知吗？如果通知只针对您一人，删除后将无法恢复。`,
      '批量删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    ).then(async () => {
      try {
        const deletePromises = selectedNotifications.value.map((item) =>
          deleteNotificationApi(item.notificationId)
        )

        await Promise.all(deletePromises)
        ElMessage.success(
          `成功删除 ${selectedNotifications.value.length} 条通知`
        )
        getTableData()
        getStats()
      } catch (error) {
        console.error('批量删除失败:', error)
        ElMessage.error('批量删除失败')
      }
    })
  }

  // 工具函数
  const formatDate = (date) => {
    if (!date) return '-'
    return formatTimeToStr(date, 'yyyy-mm-dd hh:MM:ss')
  }

  const getTypeLabel = (type) => {
    const typeMap = {
      system: '系统通知',
      business: '业务通知',
      warning: '警告通知'
    }
    return typeMap[type] || type
  }

  const getTypeTagType = (type) => {
    const typeMap = {
      system: 'info',
      business: 'success',
      warning: 'warning'
    }
    return typeMap[type] || 'info'
  }

  const getPriorityLabel = (priority) => {
    const priorityMap = {
      low: '低',
      normal: '普通',
      high: '高',
      urgent: '紧急'
    }
    return priorityMap[priority] || priority
  }

  const getPriorityTagType = (priority) => {
    const priorityMap = {
      low: 'info',
      normal: 'success',
      high: 'warning',
      urgent: 'danger'
    }
    return priorityMap[priority] || 'info'
  }

  // 初始化
  onMounted(() => {
    getTableData()
    getStats()
  })
</script>

<style scoped>
  .user-notification-center {
    padding: 20px;
  }

  .notification-stats {
    margin-bottom: 20px;
  }

  .stats-card {
    text-align: center;
    cursor: pointer;
    transition: all 0.3s;
  }

  .stats-card:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }

  .stats-card.unread {
    border-color: #f56c6c;
  }

  .stats-card.read {
    border-color: #67c23a;
  }

  .stats-content {
    padding: 20px;
  }

  .stats-number {
    font-size: 28px;
    font-weight: bold;
    color: #409eff;
    margin-bottom: 8px;
  }

  .stats-card.unread .stats-number {
    color: #f56c6c;
  }

  .stats-card.read .stats-number {
    color: #67c23a;
  }

  .stats-label {
    font-size: 14px;
    color: #909399;
  }

  .notification-title {
    display: flex;
    align-items: center;
    font-weight: normal;
  }

  .notification-title.unread {
    font-weight: bold;
    color: #303133;
  }

  .unread-dot {
    color: #f56c6c;
    margin-right: 8px;
    font-size: 8px;
  }

  .notification-detail {
    max-height: 60vh;
    overflow-y: auto;
  }

  .detail-header {
    margin-bottom: 20px;
    padding-bottom: 15px;
    border-bottom: 1px solid #ebeef5;
  }

  .detail-header h3 {
    margin: 0 0 10px 0;
    font-size: 18px;
    color: #303133;
  }

  .detail-meta {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 8px;
  }

  .detail-time {
    color: #909399;
    font-size: 12px;
  }

  .detail-content {
    line-height: 1.6;
  }

  .content-text {
    color: #606266;
    white-space: pre-wrap;
    word-break: break-word;
  }

  .gva-pagination {
    display: flex;
    justify-content: flex-end;
    margin-top: 20px;
  }

  .ml-2 {
    margin-left: 8px;
  }

  .ml-4 {
    margin-left: 16px;
  }

  .mb-4 {
    margin-bottom: 16px;
  }
</style>