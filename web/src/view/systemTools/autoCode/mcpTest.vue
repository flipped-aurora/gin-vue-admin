<template>
  <div class="p-2">
    <el-card class="mb-2">
      <template #header>
        <div class="flex flex-wrap items-center justify-between gap-3 font-bold">
          <span>MCP 服务器配置示例</span>
          <div class="flex flex-wrap items-center gap-2">
            <el-button
              size="small"
              type="primary"
              :loading="actionLoading.start"
              :disabled="disableStart"
              @click="handleStartMcp"
            >
              启动
            </el-button>
            <el-button
              size="small"
              :loading="actionLoading.stop"
              :disabled="disableStop"
              @click="handleStopMcp"
            >
              停用
            </el-button>
            <el-button
              size="small"
              :icon="RefreshRight"
              :loading="statusLoading"
              @click="refreshMcpOverview"
            >
              刷新
            </el-button>
            <el-tooltip content="复制配置" placement="top">
              <el-button :icon="DocumentCopy" circle @click="copyMcpConfig" />
            </el-tooltip>
          </div>
        </div>
      </template>

      <div class="mb-3 flex flex-wrap items-center gap-2">
        <el-tag :type="statusTagType">{{ statusText }}</el-tag>
        <span class="text-xs text-gray-500">{{ statusHint }}</span>
      </div>

      <div
        class="mb-3 rounded border border-dashed border-gray-200 bg-gray-50 px-3 py-2 text-xs text-gray-500"
      >
        <div>
          服务地址：
          <span class="break-all font-mono text-gray-700">{{
            mcpServiceStatus.baseURL
          }}</span>
        </div>
        <div class="mt-1">
          健康检查：
          <span class="break-all font-mono text-gray-700">{{
            mcpServiceStatus.healthURL
          }}</span>
        </div>
        <div v-if="mcpServiceStatus.startedAt" class="mt-1">
          启动时间：
          <span class="font-mono text-gray-700">{{
            formatTime(mcpServiceStatus.startedAt)
          }}</span>
        </div>
        <div v-if="mcpServiceStatus.lastError" class="mt-1">
          最近错误：
          <span class="break-all text-red-600">{{
            mcpServiceStatus.lastError
          }}</span>
        </div>
      </div>

      <pre
        class="font-mono whitespace-pre-wrap break-words rounded bg-gray-100 p-2.5 text-gray-700"
      >{{ mcpServerConfig }}</pre>
    </el-card>

    <el-empty
      v-if="!mcpTools.length && !serviceReachable && !statusLoading"
      class="mb-4"
      description="MCP 服务未启动，点击上方“启动”后可加载工具列表。"
    />

    <el-row :gutter="8">
      <el-col
        v-for="tool in mcpTools"
        :key="tool.name"
        :xs="24"
        :sm="12"
        :md="12"
        :lg="8"
      >
        <el-card class="mb-5 min-h-[150px] flex flex-col overflow-hidden">
          <template #header>
            <div class="flex items-center justify-between font-bold">
              <span>{{ tool.name }}</span>
              <el-tooltip content="测试工具" placement="top">
                <el-button
                  :icon="VideoPlay"
                  circle
                  @click="openTestDialog(tool)"
                />
              </el-tooltip>
            </div>
          </template>

          <div class="mb-1 text-sm">{{ tool.description }}</div>

          <div
            v-if="
              tool.inputSchema &&
              tool.inputSchema.properties &&
              Object.keys(tool.inputSchema.properties).length > 0
            "
            class="mt-1 max-h-[100px] overflow-y-auto rounded-b border-t border-gray-200 bg-gray-50 p-2 text-xs"
          >
            <p class="my-2 flex items-center font-semibold text-gray-700">
              <span class="mr-1">参数列表</span>
              <span class="text-xs text-gray-500"
                >({{ Object.keys(tool.inputSchema.properties).length }})</span
              >
            </p>
            <div class="space-y-2">
              <div
                v-for="(propDetails, propName) in tool.inputSchema.properties"
                :key="propName"
                class="flex flex-col rounded border border-gray-100 bg-white p-1.5 transition-colors hover:border-gray-300"
              >
                <div class="flex items-center justify-between">
                  <div class="flex items-center">
                    <span class="font-medium text-gray-800">{{ propName }}</span>
                    <span
                      v-if="
                        tool.inputSchema.required &&
                        tool.inputSchema.required.includes(propName)
                      "
                      class="ml-1 text-xs text-red-500"
                    >
                      *
                    </span>
                  </div>
                  <span
                    class="rounded bg-blue-100 px-1.5 py-0.5 text-xs text-blue-700"
                  >
                    {{ propDetails.type }}
                  </span>
                </div>
                <div
                  class="mt-0.5 line-clamp-2 text-xs text-gray-500"
                  :title="propDetails.description || '无描述'"
                >
                  {{ propDetails.description || '无描述' }}
                </div>
              </div>
            </div>
          </div>

          <div
            v-else
            class="mt-1 flex items-center justify-center rounded-b border-t border-gray-200 bg-gray-50 p-2 text-xs"
          >
            <span class="py-3 italic text-gray-500">无输入参数</span>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-dialog
      v-model="testDialogVisible"
      :title="currentTestingTool ? `${currentTestingTool.name} - 参数测试` : '参数测试'"
      width="60%"
      :before-close="handleCloseDialog"
    >
      <el-form
        v-if="currentTestingTool"
        ref="testParamsFormRef"
        :model="testParamsForm"
        label-width="120px"
        label-position="top"
        class="max-h-[calc(60vh-120px)] overflow-y-auto"
      >
        <el-form-item
          v-for="(propDetails, propName) in currentTestingTool.inputSchema.properties"
          :key="propName"
          :label="propDetails.description || propName"
          :prop="propName"
          :rules="
            currentTestingTool.inputSchema.required &&
            currentTestingTool.inputSchema.required.includes(propName)
              ? [
                  {
                    required: true,
                    message: '请输入' + (propDetails.description || propName),
                    trigger: 'blur'
                  }
                ]
              : []
          "
        >
          <el-input
            v-if="propDetails.type === 'string' && !propDetails.enum"
            v-model="testParamsForm[propName]"
            :placeholder="propDetails.description || `请输入 ${propName}`"
          />
          <el-input
            v-else-if="propDetails.type === 'number'"
            v-model.number="testParamsForm[propName]"
            type="number"
            :placeholder="propDetails.description || `请输入数字 ${propName}`"
          />
          <el-select
            v-else-if="propDetails.type === 'boolean'"
            v-model="testParamsForm[propName]"
            :placeholder="propDetails.description || '请选择'"
          >
            <el-option label="True" :value="true" />
            <el-option label="False" :value="false" />
          </el-select>
          <el-select
            v-else-if="propDetails.type === 'string' && propDetails.enum"
            v-model="testParamsForm[propName]"
            :placeholder="propDetails.description || `请选择 ${propName}`"
          >
            <el-option
              v-for="enumValue in propDetails.enum"
              :key="enumValue"
              :label="enumValue"
              :value="enumValue"
            />
          </el-select>
          <el-input
            v-else
            v-model="testParamsForm[propName]"
            type="textarea"
            :placeholder="`${propDetails.description || propName}（请输入 JSON 格式）`"
            :autosize="{ minRows: 2, maxRows: 6 }"
          />
        </el-form-item>
      </el-form>

      <div
        v-if="apiDialogResponse"
        class="mt-5 rounded border border-gray-200 bg-gray-50 p-[15px]"
      >
        <h4 class="mb-2.5 mt-0 text-base">API 返回结果:</h4>
        <div v-if="typeof apiDialogResponse === 'string'">
          <pre
            class="overflow-y-auto whitespace-pre-wrap break-words rounded bg-gray-100 p-2.5"
          >{{ apiDialogResponse }}</pre>
        </div>
        <div
          v-else-if="apiDialogResponse.type === 'image' && apiDialogResponse.content"
        >
          <el-image
            class="max-h-[300px] max-w-full"
            :src="apiDialogResponse.content"
            :preview-src-list="[apiDialogResponse.content]"
            fit="contain"
          />
        </div>
        <div
          v-else-if="apiDialogResponse.type === 'text' && apiDialogResponse.content"
        >
          <pre
            class="overflow-y-auto whitespace-pre-wrap break-words rounded bg-gray-100 p-2.5"
          >{{ apiDialogResponse.content }}</pre>
        </div>
        <div v-else>
          <pre
            class="overflow-y-auto whitespace-pre-wrap break-words rounded bg-gray-100 p-2.5"
          >{{ JSON.stringify(apiDialogResponse, null, 2) }}</pre>
        </div>
      </div>

      <template #footer>
        <span class="dialog-footer">
          <el-button @click="testDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleDialogTestTool">
            测试
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { computed, onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import {
  DocumentCopy,
  RefreshRight,
  VideoPlay
} from '@element-plus/icons-vue'
import {
  mcpList,
  mcpStart,
  mcpStatus,
  mcpStop,
  mcpTest
} from '@/api/autoCode'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'MCPTest'
})

const userStore = useUserStore()

const emptyStatus = () => ({
  state: 'stopped',
  managed: false,
  reachable: false,
  starting: false,
  baseURL: 'http://127.0.0.1:8889/mcp',
  healthURL: 'http://127.0.0.1:8889/health',
  listenAddr: ':8889',
  path: '/mcp',
  authHeader: 'x-token',
  startedAt: '',
  lastError: '',
  message: 'MCP 独立服务未启动'
})

const defaultServerConfig = {
  mcpServers: {
    gva: {
      url: 'http://127.0.0.1:8889/mcp',
      headers: {
        'x-token': ''
      }
    }
  }
}

const mcpTools = ref([])
const testDialogVisible = ref(false)
const currentTestingTool = ref(null)
const testParamsForm = reactive({})
const testParamsFormRef = ref(null)
const apiDialogResponse = ref(null)
const statusLoading = ref(false)
const actionLoading = reactive({ start: false, stop: false })
const mcpServiceStatus = ref(emptyStatus())

const buildMcpServerConfig = (config) => {
  const nextConfig = JSON.parse(
    JSON.stringify(config || defaultServerConfig)
  )
  const serverName = Object.keys(nextConfig.mcpServers || {})[0]

  if (serverName) {
    const headers = nextConfig.mcpServers[serverName].headers || {}
    const headerKeys = Object.keys(headers).length ? Object.keys(headers) : ['x-token']
    nextConfig.mcpServers[serverName].headers = Object.fromEntries(
      headerKeys.map((key) => [key, userStore.token || ''])
    )
  }

  return JSON.stringify(nextConfig, null, 2)
}

const mcpServerConfig = ref(buildMcpServerConfig(defaultServerConfig))

const serviceReachable = computed(() => Boolean(mcpServiceStatus.value.reachable))
const disableStart = computed(
  () =>
    statusLoading.value ||
    actionLoading.stop ||
    actionLoading.start ||
    mcpServiceStatus.value.starting ||
    serviceReachable.value
)
const disableStop = computed(
  () =>
    statusLoading.value ||
    actionLoading.start ||
    actionLoading.stop ||
    !mcpServiceStatus.value.managed
)

const statusTagType = computed(() => {
  switch (mcpServiceStatus.value.state) {
    case 'running':
      return 'success'
    case 'starting':
      return 'warning'
    case 'external':
      return 'primary'
    case 'error':
      return 'danger'
    default:
      return 'info'
  }
})

const statusText = computed(() => {
  switch (mcpServiceStatus.value.state) {
    case 'running':
      return '页面托管运行中'
    case 'starting':
      return '启动中'
    case 'external':
      return '外部服务运行中'
    case 'error':
      return '启动异常'
    default:
      return '未启动'
  }
})

const statusHint = computed(
  () => mcpServiceStatus.value.message || 'MCP 独立服务未启动'
)

const applyMcpOverview = (payload = {}) => {
  if (payload.status) {
    mcpServiceStatus.value = {
      ...emptyStatus(),
      ...payload.status
    }
  }
  if (payload.mcpServerConfig) {
    mcpServerConfig.value = buildMcpServerConfig(payload.mcpServerConfig)
  }
}

const fetchMcpTools = async ({ silent = false } = {}) => {
  if (!serviceReachable.value) {
    mcpTools.value = []
    return
  }

  const res = await mcpList()
  if (res.code === 0 && res.data?.list?.tools) {
    applyMcpOverview(res.data)
    mcpTools.value = res.data.list.tools
    return
  }

  mcpTools.value = []
  await refreshMcpOverview({ silent: true, loadTools: false })
  if (!silent) {
    ElMessage.error(res.msg || '获取工具列表失败')
  }
}

const refreshMcpOverview = async ({ silent = false, loadTools = true } = {}) => {
  statusLoading.value = true
  try {
    const res = await mcpStatus()
    if (res.code !== 0) {
      if (!silent) {
        ElMessage.error(res.msg || '获取 MCP 状态失败')
      }
      return
    }

    applyMcpOverview(res.data)
    if (mcpServiceStatus.value.reachable && loadTools) {
      await fetchMcpTools({ silent: true })
    } else if (!mcpServiceStatus.value.reachable) {
      mcpTools.value = []
    }
  } finally {
    statusLoading.value = false
  }
}

onMounted(() => {
  refreshMcpOverview({ silent: true, loadTools: true })
})

const handleStartMcp = async () => {
  actionLoading.start = true
  try {
    const res = await mcpStart()
    if (res.data) {
      applyMcpOverview(res.data)
    }
    if (res.code !== 0) {
      ElMessage.error(res.msg || '启动 MCP 失败')
      return
    }
    ElMessage.success(res.msg || 'MCP 独立服务已启动')
    await refreshMcpOverview({ silent: true, loadTools: true })
  } finally {
    actionLoading.start = false
  }
}

const handleStopMcp = async () => {
  actionLoading.stop = true
  try {
    const res = await mcpStop()
    if (res.data) {
      applyMcpOverview(res.data)
    }
    if (res.code !== 0) {
      ElMessage.error(res.msg || '停用 MCP 失败')
      return
    }
    ElMessage.success(res.msg || 'MCP 独立服务已停用')
    mcpTools.value = []
    await refreshMcpOverview({ silent: true, loadTools: false })
  } finally {
    actionLoading.stop = false
  }
}

const copyMcpConfig = async () => {
  try {
    await navigator.clipboard.writeText(mcpServerConfig.value)
    ElMessage.success('配置已复制')
  } catch (error) {
    ElMessage.error(`复制失败: ${error}`)
  }
}

const openTestDialog = (tool) => {
  currentTestingTool.value = tool
  apiDialogResponse.value = null

  for (const key in testParamsForm) {
    delete testParamsForm[key]
  }

  if (tool.inputSchema?.properties) {
    Object.keys(tool.inputSchema.properties).forEach((propName) => {
      const propDetails = tool.inputSchema.properties[propName]
      if (typeof propDetails.default !== 'undefined') {
        testParamsForm[propName] = propDetails.default
      } else if (propDetails.type === 'boolean') {
        testParamsForm[propName] = false
      } else if (propDetails.type === 'number') {
        testParamsForm[propName] = null
      } else {
        testParamsForm[propName] = ''
      }
    })
  }

  testDialogVisible.value = true
  testParamsFormRef.value?.clearValidate()
}

const handleCloseDialog = (done) => {
  apiDialogResponse.value = null
  done()
}

const normalizeToolPayload = (tool, payload) => {
  if (!tool?.inputSchema?.properties) return payload

  Object.keys(tool.inputSchema.properties).forEach((propName) => {
    const propDetails = tool.inputSchema.properties[propName]
    const value = payload[propName]

    if (
      (propDetails.type === 'object' || propDetails.type === 'array') &&
      value &&
      typeof value === 'string'
    ) {
      payload[propName] = JSON.parse(value)
    }
  })

  return payload
}

const handleDialogTestTool = async () => {
  if (!currentTestingTool.value) {
    ElMessage.warning('没有选中的测试工具')
    return
  }

  if (!serviceReachable.value) {
    ElMessage.warning('MCP 服务未启动，请先启动服务')
    return
  }

  try {
    if (testParamsFormRef.value) {
      await testParamsFormRef.value.validate()
    }

    const payload = normalizeToolPayload(currentTestingTool.value, {
      ...testParamsForm
    })

    const res = await mcpTest({
      name: currentTestingTool.value.name,
      arguments: payload
    })

    if (res.code !== 0) {
      ElMessage.error(res.msg || '工具调用失败')
      return
    }

    apiDialogResponse.value = res.data
    ElMessage.success('API 调用成功')
  } catch (error) {
    if (error?.message) {
      ElMessage.error(error.message)
    }
  }
}

const formatTime = (value) => {
  const date = new Date(value)
  return Number.isNaN(date.getTime()) ? String(value || '-') : date.toLocaleString()
}
</script>
