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
      title="开发环境能力"
      description="以下下载/编译能力仅供开发环境使用，编译过后的skill，cli可以在生产环境使用；编译类操作需等待数秒。编译下载与 Skill 包内的 CLI 已内嵌配置、可直接运行；Skill 包供 Claude Code、Codex 等 AI 助手使用。"
    />
    <el-alert class="!mt-4" :title="`先执行 ${manifest.name || cli?.command || cli?.name || 'cli'} login --token <jwt> 保存登录态。`" type="info" :closable="false" />
    <div class="flex items-center gap-2 flex-wrap mt-4">
      <el-select v-model="buildTarget.goos" size="small" style="width: 110px">
        <el-option label="Windows" value="windows" />
        <el-option label="Linux" value="linux" />
        <el-option label="macOS" value="darwin" />
      </el-select>
      <el-select v-model="buildTarget.goarch" size="small" style="width: 100px">
        <el-option label="amd64" value="amd64" />
        <el-option label="arm64" value="arm64" />
      </el-select>
      <el-button type="primary" :loading="building" @click="buildAndDownload">编译下载</el-button>
      <el-button :loading="building" @click="downloadSkill">下载 Skill</el-button>
      <el-button @click="download">下载 Manifest</el-button>
    </div>
    <div class="mt-4 text-xs text-gray-500">CLI: {{ manifest.name || cli?.command || cli?.name }} / 版本: {{ manifest.version || 'v1' }}</div>
    <el-collapse class="mt-4">
      <el-collapse-item v-for="item in manifest.commands || []" :key="item.name" :title="item.name">
        <div class="mb-2 text-sm text-gray-600">{{ item.method }} {{ item.path }}</div>
        <pre>{{ JSON.stringify(item, null, 2) }}</pre>
      </el-collapse-item>
    </el-collapse>
  </el-drawer>
</template>

<script setup>
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { previewManifest, downloadManifest, buildCliBinary, downloadCliSkill } from '@/api/system/cli'

const props = defineProps({
  modelValue: { type: Boolean, default: false },
  cli: { type: Object, default: () => ({}) }
})
const emit = defineEmits(['update:modelValue'])
const manifest = ref({ commands: [] })
const buildTarget = ref({ goos: 'windows', goarch: 'amd64' })
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
      goarch: buildTarget.value.goarch
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
      goarch: buildTarget.value.goarch
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
