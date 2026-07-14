<template>
  <div class="gva-theme-font">
    <!-- 内置预设 -->
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">内置预设</span>
        <div class="gva-theme-divider"></div>
      </div>
      <div class="gva-theme-section-content">
        <div class="grid grid-cols-2 gap-4">
          <PresetCard
            v-for="p in builtinPresets"
            :key="p.name"
            :preset="p"
            :compatible="isPresetCompatible(p)"
            @apply="handleApply(p)"
          />
        </div>
      </div>
    </div>

    <!-- 我的预设 -->
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">我的预设</span>
        <div class="gva-theme-divider"></div>
      </div>
      <div class="gva-theme-section-content">
        <div v-if="customPresets.length" class="grid grid-cols-2 gap-4 mb-4">
          <PresetCard
            v-for="p in customPresets"
            :key="p.name"
            :preset="p"
            :compatible="isPresetCompatible(p)"
            @apply="handleApply(p)"
            @remove="handleRemove(p)"
          />
        </div>
        <div
          v-else
          class="gva-theme-card-bg text-center text-sm text-muted-foreground py-6 mb-4"
        >
          暂无自定义预设，点击下方「保存当前为预设」
        </div>
        <g-button class="w-full rounded-lg font-medium" @click="handleSaveCurrent">
          保存当前为预设
        </g-button>
      </div>
    </div>

    <!-- 导入导出 -->
    <div class="mb-10">
      <div class="gva-theme-section-header">
        <div class="gva-theme-divider"></div>
        <span class="gva-theme-section-title">导入导出</span>
        <div class="gva-theme-divider"></div>
      </div>
      <div class="gva-theme-section-content">
        <div class="gva-theme-card-bg flex gap-3">
          <g-button
            variant="outline-primary"
            class="flex-1 rounded-lg"
            @click="handleExport"
          >
            导出当前配置
          </g-button>
          <el-upload
            ref="uploadRef"
            :auto-upload="false"
            :show-file-list="false"
            accept=".json"
            class="flex-1"
            @change="handleImport"
          >
            <g-button variant="outline-success" class="w-full rounded-lg">
              导入配置
            </g-button>
          </el-upload>
        </div>
        <p class="text-xs text-muted-foreground mt-3">
          导出当前完整配置（主题 / 布局 / 顶栏 / 界面），可跨账号迁移；
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useThemeStore } from '@/pinia'
import PresetCard from '../../components/presetCard.vue'
import {
  BUILTIN_PRESETS,
  loadCustomPresets,
  addCustomPreset,
  removeCustomPreset,
  serializePreset,
  parsePreset,
  isPresetCompatible
} from '@/theme'

defineOptions({
  name: 'PresetSettings'
})

const themeStore = useThemeStore()

const builtinPresets = BUILTIN_PRESETS
const customPresets = ref(loadCustomPresets())
const uploadRef = ref()

const handleApply = (preset) => {
  if (!isPresetCompatible(preset)) {
    ElMessage.warning(`该预设需要 GVA ≥ ${preset.minMainVersion}，当前版本不支持`)
    return
  }
  themeStore.applyPreset(preset)
  ElMessage.success(`已应用预设「${preset.name}」`)
}

const handleSaveCurrent = async () => {
  try {
    const { value } = await ElMessageBox.prompt('请输入预设名称', '保存当前为预设', {
      confirmButtonText: '保存',
      cancelButtonText: '取消',
      inputPattern: /\S+/,
      inputErrorMessage: '名称不能为空'
    })
    customPresets.value = addCustomPreset(themeStore.exportPreset(value.trim()))
    ElMessage.success('预设已保存')
  } catch {
    // 用户取消
  }
}

const handleRemove = async (preset) => {
  try {
    await ElMessageBox.confirm(`确定删除预设「${preset.name}」吗？`, '删除预设', {
      confirmButtonText: '删除',
      cancelButtonText: '取消',
      type: 'warning'
    })
    customPresets.value = removeCustomPreset(preset.name)
    ElMessage.success('预设已删除')
  } catch {
    // 用户取消
  }
}

const handleExport = () => {
  const data = serializePreset(themeStore.exportPreset('gin-vue-admin-theme'))
  const blob = new Blob([data], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const link = document.createElement('a')
  link.href = url
  link.download = `gin-vue-admin-theme-${new Date().toISOString().split('T')[0]}.json`
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  URL.revokeObjectURL(url)
  ElMessage.success('配置已导出')
}

const handleImport = (file) => {
  const reader = new FileReader()
  reader.onload = (e) => {
    let preset
    try {
      preset = parsePreset(e.target.result)
    } catch {
      ElMessage.error('配置文件格式错误')
      return
    }
    // JSON 合法但不是有效预设（缺 theme / 旧格式）时 parsePreset 返回 null，需显式拦截，
    // 否则会出现「什么都没应用却提示导入成功」的误导
    if (!preset?.theme) {
      ElMessage.error('配置文件不兼容或为空')
      return
    }
    if (!isPresetCompatible(preset)) {
      ElMessage.warning(`该配置需要 GVA ≥ ${preset.minMainVersion}，当前版本不支持`)
      return
    }
    themeStore.applyPreset(preset)
    ElMessage.success('配置已导入')
  }
  reader.readAsText(file.raw)
}
</script>
