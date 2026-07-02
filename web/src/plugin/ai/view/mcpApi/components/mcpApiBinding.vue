<template>
  <el-drawer :model-value="modelValue" size="1240px" :show-close="false" @close="emit('update:modelValue', false)">
    <template #header>
      <div class="flex justify-between items-center w-full gap-3">
        <div class="flex items-center gap-4">
          <span class="text-lg">管理API - {{ mcp?.name }}</span>
          <el-radio-group v-model="step">
            <el-radio-button :label="1">① 选择API</el-radio-button>
            <el-radio-button :label="2">② 已绑定API · {{ bindingRows.length }}</el-radio-button>
          </el-radio-group>
        </div>
        <div class="flex items-center">
          <el-button @click="clearSelected" :disabled="selectedApiIds.length === 0">清空已选</el-button>
          <el-button @click="emit('update:modelValue', false)">取消</el-button>
          <el-button type="primary" :loading="saving" @click="save">保存</el-button>
        </div>
      </div>
    </template>

    <!-- ① 选择 API -->
    <div v-show="step === 1">
      <div class="mb-3 text-xs text-gray-500">通过穿梭框为 MCP 增加或移除 API，保存后即时生效。</div>
      <div class="mcp-api-wrapper">
        <el-transfer v-model="selectedApiIds" class="mcp-api-transfer" :data="apiOptions" filterable
          :filter-method="filterApiOption" filter-placeholder="请输入API名称/分组/路径/请求方法" :titles="['可选API', '已选API']"
          :button-texts="['移除', '选中']">
          <template #default="{ option }">
            <div class="flex flex-col gap-1 w-full overflow-hidden">
              <div class="font-medium text-gray-800 dark:text-gray-200 truncate"
                :title="option.description || option.path">
                {{ option.description || option.path }}
              </div>
              <div class="text-xs text-gray-500 dark:text-gray-400 truncate"
                :title="`${option.apiGroup || '未分组'} · ${option.method} · ${option.path}`">
                <span>{{ option.apiGroup || '未分组' }}</span>
                <span class="mx-1">·</span>
                <span
                  :class="{ 'text-green-600 dark:text-green-500': option.method === 'GET', 'text-blue-600 dark:text-blue-500': option.method === 'POST', 'text-orange-600 dark:text-orange-500': option.method === 'PUT', 'text-red-600 dark:text-red-500': option.method === 'DELETE' }">
                  {{ option.method }}
                </span>
                <span class="mx-1">·</span>
                <span class="text-gray-400">{{ option.path }}</span>
              </div>
            </div>
          </template>
        </el-transfer>
      </div>
      <div v-if="selectedApiIds.length > 0" class="flex justify-end mt-3">
        <el-button type="primary" plain @click="step = 2">下一步：已绑定API（{{ selectedApiIds.length }}）→</el-button>
      </div>
    </div>

    <!-- ② 已绑定 API -->
    <div v-show="step === 2">
      <div v-if="bindingRows.length > 0">
        <div class="text-base font-medium mb-2">已绑定 API（可启用/禁用、删除）</div>
        <el-table :data="bindingRows" size="small" border>
          <el-table-column label="API名称" min-width="160">
            <template #default="{ row }">{{ row.description || row.path || row.commandName || '—' }}</template>
          </el-table-column>
          <el-table-column label="工具名" width="140">
            <template #default="{ row }">{{ row.commandName || '—' }}</template>
          </el-table-column>
          <el-table-column label="说明" min-width="200">
            <template #default="{ row }">
              <span class="text-xs">{{ row.commandDesc || row.autoSummary || '—' }}</span>
              <div v-if="row.path" class="text-xs text-gray-400">{{ row.method }} {{ row.path }}</div>
            </template>
          </el-table-column>
          <el-table-column label="参数数" width="70" align="center">
            <template #default="{ row }">{{ row.paramCount }}</template>
          </el-table-column>
          <el-table-column label="返回数" width="70" align="center">
            <template #default="{ row }">{{ row.responseCount }}</template>
          </el-table-column>
          <el-table-column label="启用" width="80" align="center">
            <template #default="{ row }">
              <el-switch v-model="row.enabled" @change="toggleEnabled(row)" />
            </template>
          </el-table-column>
          <el-table-column label="操作" width="190" align="center">
            <template #default="{ row }">
              <el-button link type="primary" @click="openEditor(row.apiId)">编辑</el-button>
              <el-button link type="primary" :loading="generatingIds.has(row.apiId)"
                @click="autoGenerate(row.apiId)">自动生成</el-button>
              <el-button link type="danger" @click="removeRow(row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <el-empty v-else description="还没有绑定 API，请先回到「① 选择API」选择" />
    </div>

    <el-dialog v-model="editor.visible" title="编辑工具定义" width="820px" append-to-body>
      <el-form label-width="90px">
        <el-form-item label="工具名">
          <el-input v-model="editor.commandName" placeholder="留空使用自动派生工具名" />
        </el-form-item>
        <el-form-item label="工具说明">
          <el-input v-model="editor.commandDesc" type="textarea" :rows="2" placeholder="这个工具做什么；留空使用 API 自动说明" />
        </el-form-item>
        <el-form-item label="API简介">
          <el-input v-model="editor.apiBrief" placeholder="留空使用 API 列表的简介；填写后覆盖工具简介（summary）" />
        </el-form-item>
        <el-form-item label="参数定义">
          <div class="w-full">
            <el-table :data="editor.params" size="small" border>
              <el-table-column label="字段名" min-width="140">
                <template #default="{ row }"><el-input v-model="row.field" size="small" /></template>
              </el-table-column>
              <el-table-column label="位置" width="100">
                <template #default="{ row }">
                  <el-select v-model="row.location" size="small">
                    <el-option label="query" value="query" />
                    <el-option label="body" value="body" />
                    <el-option label="path" value="path" />
                    <el-option label="header" value="header" />
                  </el-select>
                </template>
              </el-table-column>
              <el-table-column label="类型" width="90">
                <template #default="{ row }"><el-input v-model="row.type" size="small" /></template>
              </el-table-column>
              <el-table-column label="必填" width="55" align="center">
                <template #default="{ row }"><el-checkbox v-model="row.required" /></template>
              </el-table-column>
              <el-table-column label="说明" min-width="150">
                <template #default="{ row }"><el-input v-model="row.description" size="small" /></template>
              </el-table-column>
              <el-table-column label="操作" width="55" align="center">
                <template #default="{ $index }">
                  <el-button link type="danger" @click="editor.params.splice($index, 1)">删</el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-button class="mt-2" size="small" @click="addEditorParam">+ 新增参数</el-button>
          </div>
        </el-form-item>
        <el-form-item label="返回参数">
          <div class="w-full">
            <el-table :data="editor.responseParams" size="small" border>
              <el-table-column label="字段路径" min-width="200">
                <template #default="{ row }"><el-input v-model="row.name" size="small" placeholder="如 user.name" /></template>
              </el-table-column>
              <el-table-column label="说明" min-width="200">
                <template #default="{ row }"><el-input v-model="row.description" size="small" /></template>
              </el-table-column>
              <el-table-column label="操作" width="55" align="center">
                <template #default="{ $index }">
                  <el-button link type="danger" @click="editor.responseParams.splice($index, 1)">删</el-button>
                </template>
              </el-table-column>
            </el-table>
            <el-button class="mt-2" size="small" @click="addEditorResponseParam">+ 新增返回字段</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="editor.visible = false">取消</el-button>
        <el-button type="primary" @click="saveEditor">保存</el-button>
      </template>
    </el-dialog>
  </el-drawer>
</template>

<script setup>
import { ref, reactive, computed, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { addMcpApis, getMcpDetail, removeMcpApis, previewMcpManifest, previewApiCommand } from '@/plugin/ai/api/mcpApi'
import { getAllApis } from '@/api/api'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  mcp: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue', 'refresh'])

const step = ref(1) // 1=选择API，2=已绑定API
const saving = ref(false)
const apiOptions = ref([])
const selectedApiIds = ref([]) // 穿梭框当前勾选的 apiId
const originalApiIds = ref([]) // 加载时已绑定的 apiId，用于计算 diff
// 已绑定详情：apiId -> { apiId, commandName, commandDesc, apiBrief, paramsOverride, responseOverride, description, enabled, path, method }
const bindingsMap = ref({})
// apiId -> { name, summary, parameters, response } Swagger 自动派生（从 previewMcpManifest，用于预填）
const manifestMap = ref({})
const editor = reactive({ visible: false, apiId: null, commandName: '', commandDesc: '', apiBrief: '', params: [], responseParams: [] })
const generatingIds = reactive(new Set())

const mcpId = computed(() => props.mcp?.ID || props.mcp?.id)

const normalizeId = (item) => item?.apiId || item?.ApiID || item?.api?.ID || item?.Api?.ID

// 构造穿梭框选项：description（第一行）+ 分组·METHOD·path（第二行）
const buildApiOption = (api) => ({
  key: api.ID,
  label: `${api.description || ''} ${api.apiGroup || ''} ${api.path || ''} ${api.method || ''}`.trim(),
  description: api.description || '',
  apiGroup: api.apiGroup || '',
  path: api.path || '',
  method: api.method || ''
})

// 穿梭框过滤：支持按描述/分组/路径/方法搜索
const filterApiOption = (query, option) => {
  const keyword = query.trim().toLowerCase()
  if (!keyword) {
    return true
  }
  return [option.description, option.apiGroup, option.path, option.method]
    .filter(Boolean)
    .some(value => value.toLowerCase().includes(keyword))
}

const clearSelected = () => {
  selectedApiIds.value = []
  bindingsMap.value = {}
}

// apiId -> option 映射，用于在已绑定表格里补充 path/method
const apiOptionsMap = computed(() => {
  const m = {}
  apiOptions.value.forEach(o => { m[o.key] = o })
  return m
})

const parseParams = (raw) => {
  if (!raw) return []
  try {
    const v = JSON.parse(raw)
    return Array.isArray(v) ? v : []
  } catch {
    return []
  }
}

// 已绑定 API 行：合并"加载时已绑定的详情" + "穿梭框当前选中的"。
// 这样用户在第一步新选中 API 后，第二步也能立即看到（未保存的行标记为待保存）。
const bindingRows = computed(() => {
  const rows = []
  const seen = new Set()
  // 先放已绑定详情（含 enabled/commandName 等持久状态）
  Object.values(bindingsMap.value).forEach(b => {
    const opt = apiOptionsMap.value[b.apiId] || {}
    const auto = manifestMap.value[b.apiId] || {}
    const params = b.paramsOverride ? parseParams(b.paramsOverride) : (auto.parameters || [])
    const responseParams = b.responseOverride ? parseParams(b.responseOverride) : (auto.response || [])
    rows.push({
      ...b,
      path: b.path || opt.path || '',
      method: b.method || opt.method || '',
      commandName: b.commandName || auto.name || '',
      commandDesc: b.commandDesc || '',
      autoSummary: auto.summary || '',
      paramCount: params.length,
      responseCount: responseParams.length
    })
    seen.add(b.apiId)
  })
  // 再补穿梭框选中但尚未保存的（新勾选的）
  selectedApiIds.value.forEach(apiId => {
    if (seen.has(apiId)) return
    const opt = apiOptionsMap.value[apiId] || {}
    const auto = manifestMap.value[apiId] || {}
    const params = auto.parameters || []
    const responseParams = auto.response || []
    rows.push({
      apiId,
      commandName: auto.name || '',
      commandDesc: '',
      apiBrief: '',
      description: opt.description || '',
      enabled: true,
      path: opt.path || '',
      method: opt.method || '',
      autoSummary: auto.summary || '',
      paramCount: params.length,
      responseCount: responseParams.length
    })
    seen.add(apiId)
  })
  return rows
})

const loadData = async () => {
  const [apiRes, detailRes, manifestRes] = await Promise.all([
    getAllApis({}),
    getMcpDetail({ id: mcpId.value }),
    previewMcpManifest({ mcpId: mcpId.value })
  ])
  if (apiRes.code === 0) {
    apiOptions.value = (apiRes.data.apis || []).map(buildApiOption)
  }
  const ids = []
  const map = {}
  if (detailRes.code === 0) {
    ;(detailRes.data.bindings || []).forEach(b => {
      const id = normalizeId(b)
      if (!id) return
      ids.push(id)
      map[id] = {
        apiId: id,
        commandName: b.commandName || '',
        commandDesc: b.commandDesc || '',
        apiBrief: b.apiBrief || '',
        paramsOverride: b.paramsOverride || '',
        responseOverride: b.responseOverride || '',
        description: b.description || '',
        enabled: b.enabled !== false,
        path: b.path || '',
        method: b.method || ''
      }
    })
  }
  originalApiIds.value = [...ids]
  bindingsMap.value = map
  // 从 previewMcpManifest 结果构建 manifestMap，按 c.source.apiId 索引
  if (manifestRes.code === 0) {
    const m = {}
    ;(manifestRes.data.commands || []).forEach(c => {
      const id = c.source && c.source.apiId
      if (id) {
        m[id] = { name: c.name, summary: c.summary, parameters: c.parameters || [], response: c.response || [] }
      }
    })
    manifestMap.value = m
  }
  selectedApiIds.value = ids
}

watch(() => props.modelValue, (val) => {
  if (val && mcpId.value) {
    step.value = 1
    loadData()
  }
})

// 打开编辑弹窗：优先用 bindingsMap（用户覆盖），否则用 manifestMap（Swagger 自动派生）
const openEditor = (apiId) => {
  const local = bindingsMap.value[apiId] || {}
  const auto = manifestMap.value[apiId] || {}
  let params = local.paramsOverride ? parseParams(local.paramsOverride) : []
  if (params.length === 0) {
    params = (auto.parameters || []).map(p => ({ ...p }))
  }
  let responseParams = local.responseOverride ? parseParams(local.responseOverride) : []
  if (responseParams.length === 0) {
    responseParams = (auto.response || []).map(p => ({ ...p }))
  }
  editor.visible = true
  editor.apiId = apiId
  editor.commandName = local.commandName || auto.name || ''
  editor.commandDesc = local.commandDesc || auto.summary || ''
  editor.apiBrief = local.apiBrief || ''
  editor.params = params
  editor.responseParams = responseParams
}

// 自动生成：调 previewApiCommand 实时按 apiId 拉 Swagger 派生数据，刷新 manifestMap（仿 CLI）
const autoGenerate = async (apiId) => {
  if (!apiId || generatingIds.has(apiId)) return
  generatingIds.add(apiId)
  try {
    const res = await previewApiCommand({ apiId })
    if (res.code === 0) {
      const cmd = res.data
      manifestMap.value = {
        ...manifestMap.value,
        [apiId]: {
          name: cmd.name,
          summary: cmd.summary,
          description: cmd.description,
          commandDesc: cmd.commandDesc,
          parameters: cmd.parameters || [],
          response: cmd.response || []
        }
      }
      ElMessage.success('已生成该 API 的能力定义')
    } else {
      ElMessage.error(res.msg || '生成失败')
    }
  } finally {
    generatingIds.delete(apiId)
  }
}

const addEditorParam = () => {
  editor.params.push({ field: '', location: 'query', type: 'string', required: false, description: '' })
}

const addEditorResponseParam = () => {
  editor.responseParams.push({ name: '', description: '' })
}

// 保存编辑结果到 bindingsMap（尚未落库，点"保存"后整体生效）
const saveEditor = () => {
  const { apiId, commandName, commandDesc, apiBrief, params, responseParams } = editor
  bindingsMap.value = {
    ...bindingsMap.value,
    [apiId]: {
      ...(bindingsMap.value[apiId] || {}),
      commandName: commandName || '',
      commandDesc: commandDesc || '',
      apiBrief: apiBrief || '',
      paramsOverride: params.length > 0 ? JSON.stringify(params) : '',
      responseOverride: responseParams.length > 0 ? JSON.stringify(responseParams) : ''
    }
  }
  editor.visible = false
  ElMessage.success('已更新工具定义，点击"保存"后生效')
}

// 保存：新增选中、移除取消的
const save = async () => {
  if (!mcpId.value) return
  saving.value = true
  try {
    const removed = originalApiIds.value.filter(id => !selectedApiIds.value.includes(id))
    if (selectedApiIds.value.length > 0) {
      const bindings = selectedApiIds.value.map((apiId, index) => {
        const b = bindingsMap.value[apiId] || {}
        return {
          apiId,
          commandName: b.commandName || '',
          commandDesc: b.commandDesc || '',
          apiBrief: b.apiBrief || '',
          paramsOverride: b.paramsOverride || '',
          responseOverride: b.responseOverride || '',
          enabled: b.enabled !== false,
          sort: index
        }
      })
      const addRes = await addMcpApis({ mcpId: mcpId.value, bindings })
      if (addRes.code !== 0) {
        return
      }
    }
    if (removed.length > 0) {
      const removeRes = await removeMcpApis({ mcpId: mcpId.value, apiIds: removed })
      if (removeRes.code !== 0) {
        return
      }
    }
    ElMessage.success('保存成功')
    emit('refresh')
    // 保存后刷新详情，保持表格与最新绑定一致
    await loadData()
    emit('update:modelValue', false)
  } finally {
    saving.value = false
  }
}

// 启用/禁用：通过 addMcpApis 以单条 upsert 更新 enabled 状态
const toggleEnabled = async (row) => {
  if (!mcpId.value || !row.apiId) return
  const addRes = await addMcpApis({
    mcpId: mcpId.value,
    bindings: [{ apiId: row.apiId, enabled: row.enabled }]
  })
  if (addRes.code === 0) {
    ElMessage.success(row.enabled ? '已启用' : '已禁用')
    emit('refresh')
  } else {
    // 失败回滚开关
    row.enabled = !row.enabled
  }
}

// 删除单条绑定：从穿梭框与详情中移除，并立即调用 removeMcpApis
const removeRow = async (row) => {
  if (!mcpId.value || !row.apiId) return
  const removeRes = await removeMcpApis({ mcpId: mcpId.value, apiIds: [row.apiId] })
  if (removeRes.code === 0) {
    ElMessage.success('已删除')
    selectedApiIds.value = selectedApiIds.value.filter(id => id !== row.apiId)
    originalApiIds.value = originalApiIds.value.filter(id => id !== row.apiId)
    const next = { ...bindingsMap.value }
    delete next[row.apiId]
    bindingsMap.value = next
    emit('refresh')
  }
}
</script>

<style scoped>
.mcp-api-wrapper {
  display: flex;
  justify-content: center;
  width: 100%;
}

.mcp-api-wrapper :deep(.el-transfer) {
  display: flex;
  align-items: center;
  justify-content: center;
}

.mcp-api-wrapper :deep(.el-transfer-panel) {
  width: 480px !important;
}

.mcp-api-wrapper :deep(.el-transfer-panel__body) {
  height: calc(100vh - 260px) !important;
  min-height: 300px;
}

.mcp-api-wrapper :deep(.el-transfer-panel__list) {
  height: 100% !important;
}

.mcp-api-wrapper :deep(.el-transfer-panel__item) {
  width: 100% !important;
  height: auto !important;
  min-height: 60px !important;
  display: flex !important;
  align-items: flex-start !important;
  padding: 12px 16px !important;
}

.mcp-api-wrapper :deep(.el-transfer-panel__item .el-checkbox__input) {
  position: static !important;
  margin-top: 3px !important;
}

.mcp-api-wrapper :deep(.el-transfer-panel__item .el-checkbox__label) {
  display: block !important;
  flex: 1 !important;
  width: calc(100% - 24px) !important;
  padding-left: 8px !important;
  white-space: normal !important;
  overflow: hidden !important;
  line-height: normal !important;
}
</style>
