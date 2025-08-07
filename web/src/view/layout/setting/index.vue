<template>
  <el-drawer
    v-model="drawer"
    :title="t('system.systemConfig')"
    direction="rtl"
    :size="width"
    :show-close="false"
    class="theme-config-drawer"
  >
    <template #header>
      <div class="flex justify-between items-center">
        <span class="text-lg">{{ t('system.systemConfig') }}</span>
        <el-button type="primary" @click="resetConfig">重置配置</el-button>
      </div>
    </template>
    <div class="flex flex-col">
      <div class="mb-8">
        <Title :title="t('layout.setting.defaultTheme')"></Title>
        <div class="mt-2 text-sm p-2 flex items-center justify-center gap-2">
          <el-segmented
            v-model="config.darkMode"
            :options="options"
            size="default"
            @change="appStore.toggleDarkMode"
          />
        </div>
      </div>
      <div class="mb-8">
        <Title :title="t('layout.setting.themeColor')"></Title>
        <div class="mt-2 text-sm p-2 flex items-center gap-2 justify-center">
          <div
            v-for="item in colors"
            :key="item"
            class="w-5 h-5 rounded cursor-pointer flex items-center justify-center"
            :style="`background:${item}`"
            @click="appStore.togglePrimaryColor(item)"
          >
            <el-icon v-if="config.primaryColor === item">
              <Select />
            </el-icon>
          </div>
          <el-color-picker
            v-model="customColor"
            @change="appStore.togglePrimaryColor"
          />
        </div>
      </div>
      <div class="mb-8">
        <Title :title="t('layout.setting.interfaceDisplay')"></Title>
        <div class="mt-2 text-md p-2 flex flex-col gap-2">
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.showWaterMark') }}</div>
            <el-switch
              v-model="config.show_watermark"
              @change="appStore.toggleConfigWatermark"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.grayMode') }}</div>
            <el-switch v-model="config.grey" @change="appStore.toggleGrey" />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.colorFadeMode') }}</div>
            <el-switch
              v-model="config.weakness"
              @change="appStore.toggleWeakness"
            />
          </div>
          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.menuMode') }}</div>
            <el-segmented
              v-model="config.side_mode"
              :options="sideModes"
              size="default"
              @change="appStore.toggleSideMode"
            />
          </div>

          <div class="flex items-center justify-between">
            <div>{{ t('layout.setting.showTabs') }}</div>
            <el-switch
              v-model="config.showTabs"
              @change="appStore.toggleTabs"
            />
          </div>
          <div class="flex items-center justify-between gap-2">
            <div class="flex-shrink-0">{{ t('layout.setting.pageSwitchAnimation') }}</div>
            <el-select
              v-model="config.transition_type"
              @change="appStore.toggleTransition"
              class="w-40"
            >
              <el-option value="fade" :label="t('layout.setting.fadeInOut')" />
              <el-option value="slide" :label="t('layout.setting.slide')" />
              <el-option value="zoom" :label="t('layout.setting.zoom')" />
              <el-option value="none" :label="t('layout.setting.noAnimation')" />
            </el-select>
          </div>
        </div>
      </div>

      <div class="mb-8">
        <Title :title="t('layout.setting.layoutSizeConfig')"></Title>
        <div class="mt-2 text-md p-2 flex flex-col gap-2">
          <div class="flex items-center justify-between mb-2">
            <div>{{ t('layout.setting.sidebarExpandedWidth') }}</div>
            <el-input-number
              v-model="config.layout_side_width"
              :min="150"
              :max="400"
              :step="10"
            ></el-input-number>
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>{{ t('layout.setting.sidebarShrinkWidth') }}</div>
            <el-input-number
              v-model="config.layout_side_collapsed_width"
              :min="60"
              :max="100"
            ></el-input-number>
          </div>
          <div class="flex items-center justify-between mb-2">
            <div>{{ t('layout.setting.sidebarItemHeight') }}</div>
            <el-input-number
              v-model="config.layout_side_item_height"
              :min="30"
              :max="50"
            ></el-input-number>
          </div>
        </div>
      </div>
    </div>
  </el-drawer>
</template>

<script setup>
  import { ref, computed, watch } from 'vue'
  import { storeToRefs } from 'pinia'
  import { ElMessage } from 'element-plus'
  import { useAppStore } from '@/pinia'
  import { setSelfSetting } from '@/api/user'
  import Title from './title.vue'
  import { watch } from 'vue';
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual
  const { t } = useI18n() // added by mohamed hassan to support multilingual

  const appStore = useAppStore()
  const { config, device } = storeToRefs(appStore)


  defineOptions({
    name: 'GvaSetting'
  })

  const appStore = useAppStore()
  const { config, device } = storeToRefs(appStore)

  const activeTab = ref('appearance')

  const tabs = [
    { key: 'appearance', label: '外观' },
    { key: 'layout', label: '布局' },
    { key: 'general', label: '通用' }
  ]

  const width = computed(() => {
    return device.value === 'mobile' ? '100%' : '500px'
  })

  const drawer = defineModel('drawer', {
    default: true,
    type: Boolean
  })


  const options = ['dark', 'light', 'auto']
  const sideModes = [
    {
      label: t('layout.setting.normalMode'),
      value: 'normal'
    },
    {
      label: t('layout.setting.topMenuBarMode'),
      value: 'head'
    },
    {
      label: t('layout.setting.combinationMode'),
      value: 'combination'
    },
    {
      label: '侧边栏常驻',
      value: 'sidebar'
    }
  ]


  const saveConfig = async () => {
    const res = await setSelfSetting(config.value)
    if (res.code === 0) {
      localStorage.setItem('originSetting', JSON.stringify(config.value))
      ElMessage.success(t('layout.setting.copyConfigSuccess'))
    }
  }

  const resetConfig = () => {
    appStore.resetConfig()
  }

  watch(config, async () => {
    await saveConfig();
  }, { deep: true });
</script>

<style lang="scss" scoped>
.theme-config-drawer {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;

  ::v-deep(.el-drawer) {
    background: white;
  }

  ::v-deep(.el-drawer__header) {
    padding: 0;
    border: 0;
  }

  ::v-deep(.el-drawer__body) {
    padding: 0;
  }
}

.dark .theme-config-drawer {
  ::v-deep(.el-drawer) {
    background: #111827;
  }
}

.font-inter {
  font-family: 'Inter', sans-serif;
}

.reset-btn {
  border-radius: 0.5rem;
  font-weight: 500;
  transition: all 150ms ease-in-out;

  &:hover {
    transform: translateY(-2px);
    filter: brightness(0.9);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }
}

/* Custom scrollbar for webkit browsers */
::-webkit-scrollbar {
  width: 6px;
}

::-webkit-scrollbar-track {
  background: #f3f4f6;
  border-radius: 3px;
}

::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;

  &:hover {
    background: #9ca3af;
  }
}

.dark ::-webkit-scrollbar-track {
  background: #1f2937;
}

.dark ::-webkit-scrollbar-thumb {
  background: #4b5563;

  &:hover {
    background: #6b7280;
  }
}
</style>
