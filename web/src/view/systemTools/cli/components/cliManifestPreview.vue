<template>
  <el-drawer :model-value="modelValue" size="760px" :show-close="false" @close="emit('update:modelValue', false)">
    <template #header>
      <div class="flex justify-between items-center w-full">
        <span class="text-lg">命令预览 - {{ cli?.command || cli?.name }}</span>
        <div class="flex items-center gap-2 flex-wrap justify-end">
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
          <el-button @click="downloadSkill">下载 Skill</el-button>
          <el-button @click="download">下载 Manifest</el-button>
          <el-button @click="emit('update:modelValue', false)">关闭</el-button>
        </div>
      </div>
    </template>
    <el-alert :title="`先执行 ${manifest.name || cli?.command || cli?.name || 'cli'} login --token <jwt> 保存登录态。`" type="info" :closable="false" />
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

const download = async () => {
  const res = await downloadManifest({ cliId: props.cli.ID || props.cli.id })
  const blob = res instanceof Blob ? res : (res?.data instanceof Blob ? res.data : null)
  if (!blob) {
    return
  }
  const url = window.URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `${props.cli.command || props.cli.name || 'cli'}.manifest.json`
  link.click()
  window.URL.revokeObjectURL(url)
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
  const blob = res instanceof Blob ? res : (res?.data instanceof Blob ? res.data : null)
  if (!blob) {
    ElMessage.error('编译失败，未收到二进制')
    return
  }
  const ext = buildTarget.value.goos === 'windows' ? '.exe' : ''
  const url = window.URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `${props.cli.command || props.cli.name || 'cli'}${ext}`
  link.click()
  window.URL.revokeObjectURL(url)
}

const downloadSkill = async () => {
  const res = await downloadCliSkill({ cliId: props.cli.ID || props.cli.id })
  const blob = res instanceof Blob ? res : (res?.data instanceof Blob ? res.data : null)
  if (!blob) {
    ElMessage.error('生成 Skill 失败，未收到文件')
    return
  }
  const url = window.URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `${props.cli.command || props.cli.name || 'cli'}-skill.zip`
  link.click()
  window.URL.revokeObjectURL(url)
}
</script>
