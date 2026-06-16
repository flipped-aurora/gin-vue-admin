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
  </el-drawer>
</template>

<script setup>
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { addCliApis, getCliDetail, removeCliApis } from '@/api/system/cli'
import { getAllApis } from '@/api/api'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  cli: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue', 'refresh'])
const apiOptions = ref([])
const selectedApiIds = ref([])
const originalApiIds = ref([])

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

const loadData = async () => {
  const [apiRes, detailRes] = await Promise.all([
    getAllApis({}),
    getCliDetail({ id: props.cli.ID || props.cli.id })
  ])
  if (apiRes.code === 0) {
    apiOptions.value = (apiRes.data.apis || []).map(buildApiOption)
  }
  if (detailRes.code === 0) {
    const ids = (detailRes.data.bindings || []).map(normalizeId).filter(Boolean)
    selectedApiIds.value = ids
    originalApiIds.value = [...ids]
  }
}

watch(() => props.modelValue, (val) => {
  if (val && (props.cli.ID || props.cli.id)) {
    loadData()
  }
})

const save = async () => {
  const cliId = props.cli.ID || props.cli.id
  const added = selectedApiIds.value.filter(id => !originalApiIds.value.includes(id))
  const removed = originalApiIds.value.filter(id => !selectedApiIds.value.includes(id))
  if (added.length > 0) {
    const addRes = await addCliApis({ cliId, bindings: added.map((apiId, index) => ({ apiId, enabled: true, sort: index })) })
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
