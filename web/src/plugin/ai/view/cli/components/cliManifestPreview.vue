<template>
  <el-drawer :model-value="modelValue" size="760px" :show-close="false" @close="emit('update:modelValue', false)">
    <template #header>
      <div class="flex justify-between items-center w-full">
        <span class="text-lg">命令预览 - {{ cli?.command || cli?.name }}</span>
        <el-button @click="emit('update:modelValue', false)">关闭</el-button>
      </div>
    </template>
    <el-alert
      type="warning"
      :closable="false"
      show-icon
      title="开发环境能力说明"
      description="下载、编译等操作仅需在开发环境中执行。编译生成的 Skill、CLI 等产物均为独立运行版本，可直接部署至生产环境使用，无需继续依赖开发环境。"
    />
    <el-alert class="!mt-4" show-icon :title="`先执行 ${manifest.name || cli?.command || cli?.name || 'cli'} login --token <jwt> 保存登录态。`" type="info" :closable="false" />

    <el-card shadow="never" class="mt-4 border-gray-200" body-style="padding-bottom: 0;">
      <template #header>
        <span class="font-medium text-gray-800">构建与下载配置</span>
      </template>
      <el-form label-width="90px" label-position="left">
        <el-form-item label="API 地址">
          <el-input v-model="apiBaseUrl" placeholder="留空使用后台默认地址 (将编译进 CLI 或 Skill 内部)" clearable class="max-w-[400px]" />
        </el-form-item>
        <el-form-item label="目标平台">
          <el-select v-model="buildTarget.goos" class="mr-2" style="width: 140px">
            <el-option label="Windows" value="windows" />
            <el-option label="Linux" value="linux" />
            <el-option label="macOS" value="darwin" />
          </el-select>
          <el-select v-model="buildTarget.goarch" style="width: 140px">
            <el-option label="amd64" value="amd64" />
            <el-option label="arm64" value="arm64" />
          </el-select>
        </el-form-item>
        <el-form-item label="执行操作">
          <el-button type="primary" :loading="building" @click="buildAndDownload">编译 CLI 可执行程序</el-button>
          <el-button type="success" :loading="building" plain @click="downloadSkill">生成并下载 Skill 包</el-button>
          <el-button plain @click="download">获取 Manifest 文件</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <div class="flex justify-between items-center mt-6 mb-2">
      <span class="text-base font-medium text-gray-800">命令列表</span>
      <div class="text-xs text-gray-500 flex items-center gap-2">
        <span>CLI: <el-tag size="small">{{ manifest.name || cli?.command || cli?.name }}</el-tag></span>
        <span>版本: <el-tag size="small" type="info">{{ manifest.version || 'v1' }}</el-tag></span>
      </div>
    </div>

    <el-collapse v-if="manifest.commands && manifest.commands.length" class="border-t" accordion>
      <el-collapse-item v-for="item in manifest.commands" :key="item.name" :name="item.name">
        <template #title>
          <div class="flex items-center gap-3 w-full pr-4">
            <el-tag
              :type="item.method === 'GET' ? 'success' : item.method === 'POST' ? 'warning' : item.method === 'DELETE' ? 'danger' : 'primary'"
              effect="dark"
              style="width: 60px; text-align: center"
            >
              {{ item.method }}
            </el-tag>
            <span class="font-medium text-sm">{{ item.name }}</span>
            <span class="text-gray-400 text-xs ml-auto font-mono bg-gray-100 px-2 py-0.5 rounded">{{ item.path }}</span>
          </div>
        </template>
        <div class="bg-gray-50 rounded p-4 mx-2 overflow-auto text-xs font-mono text-gray-800 border">
          <pre class="m-0 leading-relaxed">{{ JSON.stringify(item, null, 2) }}</pre>
        </div>
      </el-collapse-item>
    </el-collapse>
    <el-empty v-else description="暂无命令数据" :image-size="60" />
  </el-drawer>
</template>

<script setup>
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { previewManifest, downloadManifest, buildCliBinary, downloadCliSkill } from '@/plugin/ai/api/cli'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  cli: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue'])
const manifest = ref({ commands: [] })
const buildTarget = ref({ goos: 'windows', goarch: 'amd64' })
const apiBaseUrl = ref('')
const building = ref(false)

const loadManifest = async () => {
  const res = await previewManifest({ cliId: props.cli.ID || props.cli.id })
  if (res.code === 0) {
    manifest.value = res.data || { commands: [] }
  }
}

watch(() => props.modelValue, (val) => {
  if (val && (props.cli.ID || props.cli.id)) {
    loadManifest()
  }
})

const cliBaseName = () => props.cli.command || props.cli.name || 'cli'

// 统一处理 blob 下载：后端失败时返回 JSON（含 code/msg），成功才是文件。
// 先按 content-type 判断，JSON 响应解析校验 code，非 0 视为错误并提示，不触发下载。
const saveDownload = async (res, filename, failMsg) => {
  const rawBlob = res instanceof Blob ? res : (res?.data instanceof Blob ? res.data : null)
  if (!rawBlob) {
    ElMessage.error(failMsg)
    return false
  }
  const contentType = res?.headers?.['content-type'] || ''
  let blob = rawBlob
  if (contentType.includes('application/json')) {
    const text = await rawBlob.text()
    try {
      const obj = JSON.parse(text)
      if (obj && typeof obj.code !== 'undefined' && obj.code !== 0) {
        ElMessage.error(obj.msg || failMsg)
        return false
      }
    } catch {
      // 非 JSON 文本，按文件处理
    }
    blob = new Blob([text], { type: contentType })
  }
  const url = window.URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = filename
  link.click()
  window.URL.revokeObjectURL(url)
  return true
}

const download = async () => {
  const res = await downloadManifest({ cliId: props.cli.ID || props.cli.id })
  await saveDownload(res, `${cliBaseName()}.manifest.json`, '下载 Manifest 失败')
}

const buildAndDownload = async () => {
  building.value = true
  let res
  try {
    res = await buildCliBinary({
      cliId: props.cli.ID || props.cli.id,
      goos: buildTarget.value.goos,
      goarch: buildTarget.value.goarch,
      baseUrl: apiBaseUrl.value
    })
  } catch (e) {
    building.value = false
    return
  } finally {
    building.value = false
  }
  const ext = buildTarget.value.goos === 'windows' ? '.exe' : ''
  await saveDownload(res, `${cliBaseName()}${ext}`, '编译失败，未收到二进制')
}

const downloadSkill = async () => {
  building.value = true
  let res
  try {
    res = await downloadCliSkill({
      cliId: props.cli.ID || props.cli.id,
      goos: buildTarget.value.goos,
      goarch: buildTarget.value.goarch,
      baseUrl: apiBaseUrl.value
    })
  } catch (e) {
    building.value = false
    return
  } finally {
    building.value = false
  }
  await saveDownload(res, `${cliBaseName()}-skill.zip`, '生成 Skill 失败，未收到文件')
}
</script>
