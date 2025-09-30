<template>
  <div class="user-notification-center">
    <!-- 页面标题 -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="Refresh" @click="getTableData">刷新</el-button>
        <el-button 
          type="success" 
          icon="Check" 
          :disabled="!multipleSelection.length"
          @click="batchMarkAsRead"
        >
          批量标记已读
        </el-button>
        <el-button 
          type="danger" 
          icon="Delete" 
          :disabled="!multipleSelection.length"
          @click="batchDelete"
        >
          批量删除
        </el-button>
      </div>

      <!-- 搜索区域 -->
      <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
          <el-form-item label="通知类型">
            <el-select v-model="searchInfo.type" placeholder="请选择通知类型" clearable>
              <el-option label="公告" value="announcement" />
              <el-option label="通知" value="notification" />
              <el-option label="警告" value="warning" />
              <el-option label="违规" value="violation" />
              <el-option label="信息" value="info" />
            </el-select>
          </el-form-item>
          <el-form-item label="阅读状态">
            <el-select v-model="searchInfo.isRead" placeholder="请选择阅读状态" clearable>
              <el-option label="未读" :value="false" />
              <el-option label="已读" :value="true" />
            </el-select>
          </el-form-item>
          <el-form-item label="优先级">
            <el-select v-model="searchInfo.priority" placeholder="请选择优先级" clearable>
              <el-option label="低" value="low" />
              <el-option label="中" value="medium" />
              <el-option label="高" value="high" />
              <el-option label="紧急" value="urgent" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 统计信息 -->
      <div class="notification-stats">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-card class="stats-card">
              <div class="stats-item">
                <div class="stats-number">{{ stats.totalCount }}</div>
                <div class="stats-label">总通知数</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stats-card unread">
              <div class="stats-item">
                <div class="stats-number">{{ stats.unreadCount }}</div>
                <div class="stats-label">未读通知</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stats-card read">
              <div class="stats-item">
                <div class="stats-number">{{ stats.readCount }}</div>
                <div class="stats-label">已读通知</div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card class="stats-card">
              <div class="stats-item">
                <div class="stats-number">{{ unreadPercentage }}%</div>
                <div class="stats-label">未读率</div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>

      <!-- 通知列表 -->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @row-click="handleRowClick"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.createdAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="标题" prop="title" show-overflow-tooltip />
        <el-table-column align="left" label="类型" prop="type" width="100">
          <template #default="scope">
            <el-tag :type="getTypeTagType(scope.row.type)">{{ getTypeLabel(scope.row.type) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="优先级" prop="priority" width="100">
          <template #default="scope">
            <el-tag :type="getPriorityTagType(scope.row.priority)">{{ getPriorityLabel(scope.row.priority) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="阅读状态" prop="isRead" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.isRead ? 'success' : 'warning'">
              {{ scope.row.isRead ? '已读' : '未读' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="阅读时间" prop="readTime" width="180">
          <template #default="scope">
            <span>{{ scope.row.readTime ? formatDate(scope.row.readTime) : '-' }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="240">
          <template #default="scope">
            <el-button type="primary" link icon="view" size="small" @click="viewNotification(scope.row)">查看</el-button>
            <el-button 
              v-if="!scope.row.isRead" 
              type="success" 
              link 
              icon="check" 
              size="small" 
              @click="markAsRead(scope.row)"
            >
              标记已读
            </el-button>
            <el-button type="danger" link icon="delete" size="small" @click="deleteNotificationHandler(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
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

    <!-- 通知详情弹窗 -->
    <el-dialog v-model="detailVisible" title="通知详情" width="60%" :before-close="closeDetail">
      <div v-if="currentNotification" class="notification-detail">
        <div class="detail-header">
          <h3>{{ currentNotification.title }}</h3>
          <div class="detail-meta">
            <el-tag :type="getTypeTagType(currentNotification.type)">{{ getTypeLabel(currentNotification.type) }}</el-tag>
            <el-tag :type="getPriorityTagType(currentNotification.priority)">{{ getPriorityLabel(currentNotification.priority) }}</el-tag>
            <el-tag :type="currentNotification.isRead ? 'success' : 'warning'">
              {{ currentNotification.isRead ? '已读' : '未读' }}
            </el-tag>
          </div>
        </div>
        <div class="detail-content">
          <div v-html="currentNotification.content"></div>
        </div>
        <div class="detail-footer">
          <p><strong>发送时间：</strong>{{ formatDate(currentNotification.createdAt) }}</p>
          <p v-if="currentNotification.readTime"><strong>阅读时间：</strong>{{ formatDate(currentNotification.readTime) }}</p>
        </div>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button v-if="!currentNotification?.isRead" type="success" @click="markAsReadInDetail">标记已读</el-button>
          <el-button type="danger" @click="deleteInDetail">删除</el-button>
          <el-button @click="closeDetail">关闭</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import { getUserNotificationList, markNotificationAsRead, deleteUserNotification } from '@/api/notice'

defineOptions({
  name: 'UserNotificationCenter'
})

// 响应式数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const multipleSelection = ref([])
const detailVisible = ref(false)
const currentNotification = ref(null)

// 搜索条件
const searchInfo = reactive({
  type: '',
  isRead: null,
  priority: ''
})

// 统计信息
const stats = reactive({
  totalCount: 0,
  unreadCount: 0,
  readCount: 0
})

// 计算未读率
const unreadPercentage = computed(() => {
  if (stats.totalCount === 0) return 0
  return Math.round((stats.unreadCount / stats.totalCount) * 100)
})

// 获取表格数据
const getTableData = async () => {
  const table = await getUserNotificationList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    
    // 更新统计信息
    stats.totalCount = table.data.total
    stats.unreadCount = table.data.list.filter(item => !item.isRead).length
    stats.readCount = table.data.list.filter(item => item.isRead).length
  }
}

// 分页相关
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 多选相关
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 搜索相关
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

const onReset = () => {
  searchInfo.type = ''
  searchInfo.isRead = null
  searchInfo.priority = ''
  getTableData()
}

// 查看通知详情
const viewNotification = (row) => {
  currentNotification.value = row
  detailVisible.value = true
  
  // 如果是未读通知，自动标记为已读
  if (!row.isRead) {
    markAsRead(row, false)
  }
}

// 关闭详情弹窗
const closeDetail = () => {
  detailVisible.value = false
  currentNotification.value = null
}

// 行点击事件
const handleRowClick = (row) => {
  viewNotification(row)
}

// 标记已读
const markAsRead = async (row, showMessage = true) => {
  const res = await markNotificationAsRead({
    notificationIds: [row.notificationId]
  })
  if (res.code === 0) {
    row.isRead = true
    row.readTime = new Date().toISOString()
    if (showMessage) {
      ElMessage.success('标记已读成功')
    }
    getTableData()
  }
}

// 在详情中标记已读
const markAsReadInDetail = async () => {
  await markAsRead(currentNotification.value)
  currentNotification.value.isRead = true
  currentNotification.value.readTime = new Date().toISOString()
}

// 批量标记已读
const batchMarkAsRead = async () => {
  const notificationIds = multipleSelection.value.map(item => item.notificationId)
  const res = await markNotificationAsRead({
    notificationIds
  })
  if (res.code === 0) {
    ElMessage.success('批量标记已读成功')
    getTableData()
  }
}

// 删除通知
const deleteNotificationHandler = async (row) => {
  ElMessageBox.confirm('此操作将永久删除该通知, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteUserNotification(row.notificationId)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 在详情中删除
const deleteInDetail = async () => {
  ElMessageBox.confirm('此操作将永久删除该通知, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteUserNotification(currentNotification.value.notificationId)
    if (res.code === 0) {
      ElMessage.success('删除成功')
      closeDetail()
      getTableData()
    }
  })
}

// 批量删除
const batchDelete = async () => {
  ElMessageBox.confirm('此操作将永久删除选中的通知, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const promises = multipleSelection.value.map(item => deleteUserNotification(item.notificationId))
    await Promise.all(promises)
    ElMessage.success('批量删除成功')
    getTableData()
  })
}

// 获取类型标签样式
const getTypeTagType = (type) => {
  const typeMap = {
    announcement: '',
    notification: 'success',
    warning: 'warning',
    violation: 'danger',
    info: 'info'
  }
  return typeMap[type] || ''
}

// 获取类型标签文本
const getTypeLabel = (type) => {
  const typeMap = {
    announcement: '公告',
    notification: '通知',
    warning: '警告',
    violation: '违规',
    info: '信息'
  }
  return typeMap[type] || type
}

// 获取优先级标签样式
const getPriorityTagType = (priority) => {
  const priorityMap = {
    low: 'info',
    medium: '',
    high: 'warning',
    urgent: 'danger'
  }
  return priorityMap[priority] || ''
}

// 获取优先级标签文本
const getPriorityLabel = (priority) => {
  const priorityMap = {
    low: '低',
    medium: '中',
    high: '高',
    urgent: '紧急'
  }
  return priorityMap[priority] || priority
}

// 组件挂载时获取数据
onMounted(() => {
  getTableData()
})
</script>

<style lang="scss" scoped>
.user-notification-center {
  .notification-stats {
    margin-bottom: 20px;
    
    .stats-card {
      text-align: center;
      cursor: pointer;
      transition: all 0.3s;
      
      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
      }
      
      &.unread {
        border-left: 4px solid #f56c6c;
      }
      
      &.read {
        border-left: 4px solid #67c23a;
      }
      
      .stats-item {
        .stats-number {
          font-size: 28px;
          font-weight: bold;
          color: #303133;
          margin-bottom: 8px;
        }
        
        .stats-label {
          font-size: 14px;
          color: #909399;
        }
      }
    }
  }
  
  .notification-detail {
    .detail-header {
      border-bottom: 1px solid #ebeef5;
      padding-bottom: 16px;
      margin-bottom: 20px;
      
      h3 {
        margin: 0 0 12px 0;
        color: #303133;
      }
      
      .detail-meta {
        .el-tag {
          margin-right: 8px;
        }
      }
    }
    
    .detail-content {
      line-height: 1.6;
      color: #606266;
      margin-bottom: 20px;
      min-height: 100px;
    }
    
    .detail-footer {
      border-top: 1px solid #ebeef5;
      padding-top: 16px;
      color: #909399;
      font-size: 14px;
      
      p {
        margin: 8px 0;
      }
    }
  }
}
</style>