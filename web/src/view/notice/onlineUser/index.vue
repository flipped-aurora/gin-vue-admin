<template>
  <div class="online-user-management">
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="searchInfo.username"
            placeholder="请输入用户名"
            clearable
          />
        </el-form-item>
        <el-form-item label="昵称" prop="nickName">
          <el-input
            v-model="searchInfo.nickName"
            placeholder="请输入昵称"
            clearable
          />
        </el-form-item>
        <el-form-item label="角色ID" prop="roleId">
          <el-input-number
            v-model="searchInfo.roleId"
            placeholder="请输入角色ID"
            clearable
            :min="1"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
          <el-button @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 统计卡片 -->
    <div class="m-4">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card class="">
            <div class="flex items-center justify-between">
              <div
                class="bg-green-5 rounded-full text-white text-2xl w-12 h-12 p-4 flex items-center justify-center"
              >
                <el-icon><User /></el-icon>
              </div>
              <div class="flex-1 flex items-center flex-col items-end">
                <div class="text-3xl font-bold">
                  {{ onlineStats.totalOnline || 0 }}
                </div>
                <div class="text-sm mt-2 text-center text-black/60">
                  在线用户
                </div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="">
            <div class="flex items-center justify-between">
              <div
                class="bg-blue-5 rounded-full text-white text-2xl w-12 h-12 p-4 flex items-center justify-center today"
              >
                <el-icon><Calendar /></el-icon>
              </div>
              <div class="flex-1 flex items-center flex-col">
                <div class="text-3xl font-bold">
                  {{ onlineStats.todayLogin || 0 }}
                </div>
                <div class="text-sm mt-2 text-center text-black/60">
                  今日登录
                </div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="">
            <div class="flex items-center justify-between">
              <div
                class="bg-orange-5 rounded-full text-white text-2xl w-12 h-12 p-4 flex items-center justify-center peak"
              >
                <el-icon><TrendCharts /></el-icon>
              </div>
              <div class="flex-1 flex items-center flex-col">
                <div class="text-3xl font-bold">
                  {{ onlineStats.peakOnline || 0 }}
                </div>
                <div class="text-sm mt-2 text-center text-black/60">
                  峰值在线
                </div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card class="">
            <div class="flex items-center justify-between">
              <div
                class="bg-amber rounded-full text-white text-2xl w-12 h-12 p-4 flex items-center justify-center avg"
              >
                <el-icon><DataAnalysis /></el-icon>
              </div>
              <div class="flex-1 flex items-center flex-col">
                <div class="text-3xl font-bold">
                  {{ onlineStats.avgOnline || 0 }}
                </div>
                <div class="text-sm mt-2 text-center text-black/60">
                  平均在线
                </div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="refresh" @click="getTableData"
          >刷新</el-button
        >
        <el-button
          type="danger"
          icon="close"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onKickUsers"
          >批量下线</el-button
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
        <el-table-column align="left" label="用户名" prop="username" />
        <el-table-column align="left" label="昵称" prop="nickName" />
        <el-table-column
          align="left"
          label="Socket ID"
          prop="socketId"
          width="280"
        >
          <template #default="scope">
            <el-tag size="small">{{ scope.row.socketId }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="IP地址"
          prop="ipAddress"
          width="150"
        />
        <!-- <el-table-column
          align="left"
          label="用户代理"
          prop="userAgent"
          width="200"
        >
          <template #default="scope">
            <el-tooltip :content="scope.row.userAgent" placement="top">
              <span class="user-agent-text">{{ scope.row.userAgent }}</span>
            </el-tooltip>
          </template>
        </el-table-column> -->
        <!-- <el-table-column
          align="left"
          label="登录时间"
          prop="loginTime"
          width="180"
        >
          <template #default="scope">
            <span>{{ formatDate(scope.row.loginTime) }}</span>
          </template>
        </el-table-column> -->
        <el-table-column
          align="left"
          label="最后活跃"
          prop="lastActiveTime"
          width="180"
        >
          <template #default="scope">
            <span>{{ formatDate(scope.row.lastActiveTime) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="在线时长"
          prop="onlineDuration"
          width="120"
        >
          <template #default="scope">
            <span>{{ formatDuration(scope.row.onlineDuration) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 'online' ? 'success' : 'info'">
              {{ scope.row.status === 'online' ? '在线' : '离线' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="120"
        >
          <template #default="scope">
            <el-button
              v-if="scope.row.status === 'online'"
              type="danger"
              link
              icon="close"
              size="small"
              @click="kickUser(scope.row)"
              >强制下线</el-button
            >
            <span v-else class="text-gray-400">已离线</span>
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
  </div>
</template>

<script setup>
  import {
    getOnlineUserList,
    kickUser as kickUserApi,
    getOnlineUserStats
  } from '@/api/notice'
  import { formatDate } from '@/utils/format'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { ref, onMounted, onUnmounted } from 'vue'
  import {
    User,
    Calendar,
    TrendCharts,
    DataAnalysis
  } from '@element-plus/icons-vue'

  defineOptions({
    name: 'OnlineUserManagement'
  })

  // 响应式数据
  const searchInfo = ref({
    username: '',
    nickName: '',
    roleId: null
  })

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const multipleSelection = ref([])
  const onlineStats = ref({})
  const refreshTimer = ref(null)

  // 获取列表数据
  const getTableData = async () => {
    const table = await getOnlineUserList({
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

  // 获取统计数据
  const getStatsData = async () => {
    const res = await getOnlineUserStats()
    if (res.code === 0) {
      onlineStats.value = res.data
    }
  }

  // 初始化数据
  const initData = async () => {
    await getTableData()
    await getStatsData()
  }

  initData()

  // 设置定时刷新
  const startAutoRefresh = () => {
    refreshTimer.value = setInterval(() => {
      getTableData()
      getStatsData()
    }, 30000) // 30秒刷新一次
  }

  // 停止定时刷新
  const stopAutoRefresh = () => {
    if (refreshTimer.value) {
      clearInterval(refreshTimer.value)
      refreshTimer.value = null
    }
  }

  onMounted(() => {
    startAutoRefresh()
  })

  onUnmounted(() => {
    stopAutoRefresh()
  })

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
      username: '',
      nickName: '',
      roleId: null
    }
    getTableData()
  }

  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 强制用户下线
  const kickUser = (row) => {
    ElMessageBox.confirm(`确定要强制用户 "${row.username}" 下线吗?`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const res = await kickUserApi(row.userId)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '操作成功'
        })
        getTableData()
      }
    })
  }

  // 批量强制下线
  const onKickUsers = async () => {
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要下线的用户'
      })
      return
    }

    const onlineUsers = multipleSelection.value.filter(
      (user) => user.status === 'online'
    )
    if (onlineUsers.length === 0) {
      ElMessage({
        type: 'warning',
        message: '所选用户均已离线'
      })
      return
    }

    ElMessageBox.confirm(
      `确定要强制 ${onlineUsers.length} 个用户下线吗?`,
      '提示',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    ).then(async () => {
      const promises = onlineUsers.map((user) => kickUserApi(user.userId))
      try {
        await Promise.all(promises)
        ElMessage({
          type: 'success',
          message: '批量下线成功'
        })
        getTableData()
      } catch (error) {
        ElMessage({
          type: 'error',
          message: '部分用户下线失败'
        })
      }
    })
  }

  // 格式化在线时长
  const formatDuration = (seconds) => {
    if (!seconds) return '0分钟'

    const hours = Math.floor(seconds / 3600)
    const minutes = Math.floor((seconds % 3600) / 60)

    if (hours > 0) {
      return `${hours}小时${minutes}分钟`
    } else {
      return `${minutes}分钟`
    }
  }
</script>
