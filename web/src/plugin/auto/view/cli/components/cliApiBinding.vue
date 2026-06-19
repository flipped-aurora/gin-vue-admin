<template>
  <el-drawer :model-value="modelValue" size="1240px" :show-close="false" @close="emit('update:modelValue', false)">
    <template #header>
      <div class="flex justify-between items-center w-full gap-3">
        <span class="text-lg">管理API - {{ cli?.name }}</span>
        <div class="flex items-center gap-2">
          <el-button @click="clearSelected" :disabled="selectedApiIds.length === 0">清空已选</el-button>
          <el-button @click="emit('update:modelValue', false)">取消</el-button>
          <el-button type="primary" @click="save">保存</el-button>
        </div>
      </div>
    </template>
    <div class="mb-3 text-xs text-gray-500">通过增加或移除 API，动态重新生成 CLI Manifest。</div>
    <div class="cli-api-wrapper">
      <el-transfer v-model="selectedApiIds" class="cli-api-transfer" :data="apiOptions" filterable
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

    <div v-if="selectedApiIds.length > 0" class="mt-6">
      <div class="text-base font-medium mb-2">命令定义（可自定义说明与参数）</div>
      <el-table :data="bindingRows" size="small" border>
        <el-table-column label="API" min-width="220">
          <template #default="{ row }">
            <div class="truncate" :title="row.path">{{ row.description || row.path }}</div>
            <div class="text-xs text-gray-400">{{ row.method }} {{ row.path }}</div>
          </template>
        </el-table-column>
        <el-table-column label="命令" width="160">
          <template #default="{ row }">{{ row.commandName || '—' }}</template>
        </el-table-column>
        <el-table-column label="说明" min-width="200">
          <template #default="{ row }">
            <span class="text-xs">{{ row.commandDesc || row.autoSummary || '—' }}</span>
          </template>
        </el-table-column>
        <el-table-column label="参数数" width="70" align="center">
          <template #default="{ row }">{{ row.paramCount }}</template>
        </el-table-column>
        <el-table-column label="返回数" width="70" align="center">
          <template #default="{ row }">{{ row.responseCount }}</template>
        </el-table-column>
        <el-table-column label="操作" width="150" align="center">
          <template #default="{ row }">
            <el-button link type="primary" @click="openEditor(row.apiId)">编辑</el-button>
            <el-button link type="primary" :loading="generatingIds.has(row.apiId)"
              @click="autoGenerate(row.apiId)">自动生成</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="editor.visible" title="编辑命令定义" width="820px" append-to-body>
      <el-form label-width="90px">
        <el-form-item label="命令名">
          <el-input v-model="editor.commandName" placeholder="留空使用自动派生命令名" />
        </el-form-item>
        <el-form-item label="命令说明">
          <el-input v-model="editor.commandDesc" type="textarea" :rows="2" placeholder="这条命令做什么；留空使用 API 自动说明" />
        </el-form-item>
        <el-form-item label="API简介">
          <el-input v-model="editor.apiBrief" placeholder="留空使用 API 列表的简介；填写后覆盖命令简介（summary）" />
        </el-form-item>
        <el-form-item label="参数定义">
          <div class="w-full">
            <el-table :data="editor.params" size="small" border>
              <el-table-column label="flag" width="120">
                <template #default="{ row }"><el-input v-model="row.flag" size="small" /></template>
              </el-table-column>
              <el-table-column label="字段名" width="120">
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
import { addCliApis, getCliDetail, removeCliApis, previewManifest, previewApiCommand } from '@/plugin/auto/api/cli'
import { getAllApis } from '@/api/api'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  cli: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue', 'refresh'])
const apiOptions = ref([])
const selectedApiIds = ref([])
const originalApiIds = ref([])
// apiId -> { commandName, commandDesc, paramsOverride } 用户自定义覆盖
const bindingsMap = ref({})
// apiId -> { name, summary, parameters } Swagger 自动派生（用于预填）
const manifestMap = ref({})
const editor = reactive({ visible: false, apiId: null, commandName: '', commandDesc: '', apiBrief: '', params: [], responseParams: [] })
const generatingIds = reactive(new Set())

const normalizeId = (item) => item?.apiId || item?.ApiID || item?.api?.ID || item?.Api?.ID

const buildApiOption = (api) => ({
  key: api.ID,
  label: `${api.description || ''} ${api.apiGroup || ''} ${api.path || ''} ${api.method || ''}`.trim(),
  description: api.description || '',
  apiGroup: api.apiGroup || '',
  path: api.path || '',
  method: api.method || ''
})

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
}

const parseParams = (raw) => {
  if (!raw) return []
  try {
    const v = JSON.parse(raw)
    return Array.isArray(v) ? v : []
  } catch {
    return []
  }
}

const apiOptionsMap = computed(() => {
  const m = {}
  apiOptions.value.forEach(o => { m[o.key] = o })
  return m
})

const bindingRows = computed(() => selectedApiIds.value.map(id => {
  const opt = apiOptionsMap.value[id] || {}
  const local = bindingsMap.value[id] || {}
  const auto = manifestMap.value[id] || {}
  const params = local.paramsOverride ? parseParams(local.paramsOverride) : (auto.parameters || [])
  const responseParams = local.responseOverride ? parseParams(local.responseOverride) : (auto.response || [])
  return {
    apiId: id,
    description: local.apiBrief || opt.description,
    path: opt.path,
    method: opt.method,
    commandName: local.commandName || auto.name || '',
    commandDesc: local.commandDesc || '',
    autoSummary: auto.summary || '',
    paramCount: params.length,
    responseCount: responseParams.length
  }
}))

const loadData = async () => {
  const [apiRes, detailRes, manifestRes] = await Promise.all([
    getAllApis({}),
    getCliDetail({ id: props.cli.ID || props.cli.id }),
    previewManifest({ cliId: props.cli.ID || props.cli.id })
  ])
  if (apiRes.code === 0) {
    apiOptions.value = (apiRes.data.apis || []).map(buildApiOption)
  }
  let ids = []
  if (detailRes.code === 0) {
    const map = {}
    ids = []
    ;(detailRes.data.bindings || []).forEach(b => {
      const id = normalizeId(b)
      if (!id) return
      ids.push(id)
      map[id] = {
        commandName: b.commandName || '',
        commandDesc: b.commandDesc || '',
        apiBrief: b.apiBrief || '',
        paramsOverride: b.paramsOverride || '',
        responseOverride: b.responseOverride || ''
      }
    })
    originalApiIds.value = [...ids]
    bindingsMap.value = map
  }
  if (manifestRes.code === 0) {
    const map = {}
    ;(manifestRes.data.commands || []).forEach(c => {
      const id = c.source && c.source.apiId
      if (id) {
        map[id] = { name: c.name, summary: c.summary, parameters: c.parameters || [], response: c.response || [] }
      }
    })
    manifestMap.value = map
  }
  selectedApiIds.value = ids
}

watch(() => props.modelValue, (val) => {
  if (val && (props.cli.ID || props.cli.id)) {
    loadData()
  }
})

const autoGenerate = async (apiId) => {
  const cliId = props.cli.ID || props.cli.id
  if (!cliId || !apiId || generatingIds.has(apiId)) return
  generatingIds.add(apiId)
  try {
    const res = await previewApiCommand({ apiId, cliId })
    if (res.code === 0) {
      manifestMap.value = {
        ...manifestMap.value,
        [apiId]: {
          name: res.data.name,
          summary: res.data.summary,
          parameters: res.data.parameters || [],
          response: res.data.response || []
        }
      }
      ElMessage.success('已生成该 API 的命令定义')
    }
  } finally {
    generatingIds.delete(apiId)
  }
}

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

const addEditorParam = () => {
  editor.params.push({ flag: '', field: '', location: 'query', type: 'string', required: false, description: '' })
}

const addEditorResponseParam = () => {
  editor.responseParams.push({ name: '', description: '' })
}

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
  ElMessage.success('已更新命令定义，点击"保存"后生效')
}

const save = async () => {
  const cliId = props.cli.ID || props.cli.id
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
        enabled: true,
        sort: index
      }
    })
    const addRes = await addCliApis({ cliId, bindings })
    if (addRes.code !== 0) {
      return
    }
  }
  if (removed.length > 0) {
    const removeRes = await removeCliApis({ cliId, apiIds: removed })
    if (removeRes.code !== 0) {
      return
    }
  }
  ElMessage.success('保存成功')
  emit('update:modelValue', false)
  emit('refresh')
}
</script>

<style scoped>
.cli-api-wrapper {
  display: flex;
  justify-content: center;
  width: 100%;
}

.cli-api-wrapper :deep(.el-transfer) {
  display: flex;
  align-items: center;
  justify-content: center;
}

.cli-api-wrapper :deep(.el-transfer-panel) {
  width: 480px !important;
}

.cli-api-wrapper :deep(.el-transfer-panel__body) {
  height: 560px !important;
}

.cli-api-wrapper :deep(.el-transfer-panel__list) {
  height: 100% !important;
}

.cli-api-wrapper :deep(.el-transfer-panel__item) {
  width: 100% !important;
  height: auto !important;
  min-height: 60px !important;
  display: flex !important;
  align-items: flex-start !important;
  padding: 12px 16px !important;
}

.cli-api-wrapper :deep(.el-transfer-panel__item .el-checkbox__input) {
  position: static !important;
  margin-top: 3px !important;
}

.cli-api-wrapper :deep(.el-transfer-panel__item .el-checkbox__label) {
  display: block !important;
  flex: 1 !important;
  width: calc(100% - 24px) !important;
  padding-left: 8px !important;
  white-space: normal !important;
  overflow: hidden !important;
  line-height: normal !important;
}
</style>
