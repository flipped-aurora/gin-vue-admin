<template>
  <div class="grid grid-cols-2 max-[480px]:grid-cols-1 gap-2.5 px-1 gva-theme-font">
    <div
      v-for="layout in layoutModes"
      :key="layout.value"
      class="gva-theme-layout-card"
      :style="modelValue === layout.value ? {
        borderColor: primaryColor,
        boxShadow: `0 0 0 1px ${primaryColor}`
      } : {}"
      @click="handleLayoutChange(layout.value)"
    >
      <div class="flex justify-center mb-1.5">
        <div
          class="w-full h-12 bg-gray-50 dark:bg-gray-600 border border-gray-200 dark:border-gray-500 rounded-md p-1.5 flex gap-1"
          :class="layout.containerClass"
        >
          <div
            v-if="layout.showSidebar"
            class="rounded-sm flex flex-col items-center pt-1"
            :class="[layout.sidebarClass]"
            :style="getSidebarStyle(layout)"
          >
            <!-- 通栏侧边：Logo 置顶指示 -->
            <div v-if="layout.topLogo" class="w-1.5 h-1.5 rounded-full bg-white/80"></div>
          </div>

          <div class="flex-1 flex flex-col gap-1.5">
            <div
              v-if="layout.showHeader"
              class="rounded-sm"
              :class="layout.headerClass"
              :style="getHeaderStyle(layout)"
            ></div>

            <div
              class="flex-1 rounded-sm"
              :class="layout.contentClass"
              :style="getContentStyle(layout)"
            ></div>
          </div>
        </div>
      </div>

      <div class="text-center">
        <span class="block text-[13px] font-semibold gva-theme-text-main" :style="modelValue === layout.value ? { color: primaryColor } : {}">{{ layout.label }}</span>
        <span class="block text-[11px] text-gray-400 dark:text-gray-500 mt-0.5 leading-tight">{{ layout.description }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useThemeStore } from '@/pinia'
import { addOpacityToColor } from '@/theme/color'

defineOptions({
  name: 'LayoutModeCard'
})

defineProps({
  modelValue: {
    type: String,
    default: 'normal'
  }
})

const emit = defineEmits(['update:modelValue'])

const themeStore = useThemeStore()
const { settings } = storeToRefs(themeStore)

const primaryColor = computed(() => settings.value.themeColor)
const lighterPrimaryColor = computed(() => addOpacityToColor(primaryColor.value, 0.7))
const lightestPrimaryColor = computed(() => addOpacityToColor(primaryColor.value, 0.4))

const layoutModes = [
  {
    value: 'normal',
    label: '经典布局',
    description: '左侧导航，顶部标题栏',
    containerClass: '',
    showSidebar: true,
    sidebarClass: 'w-1/4',
    showHeader: true,
    headerClass: 'h-1/4',
    contentClass: '',
    primaryElement: 'sidebar'
  },
  {
    value: 'head',
    label: '顶部导航',
    description: '水平导航栏布局',
    containerClass: 'flex-col',
    showSidebar: false,
    showHeader: true,
    headerClass: 'h-1/3',
    contentClass: '',
    primaryElement: 'header'
  },
  {
    value: 'combination',
    label: '混合布局',
    description: '多级导航组合模式',
    containerClass: '',
    showSidebar: true,
    sidebarClass: 'w-1/5',
    showHeader: true,
    headerClass: 'h-1/4',
    contentClass: '',
    primaryElement: 'header',
    secondaryElement: 'sidebar'
  },
  {
    value: 'sidebar',
    label: '侧栏常驻',
    description: '二级菜单会始终打开',
    containerClass: '',
    showSidebar: true,
    sidebarClass: 'w-1/3',
    showHeader: true,
    headerClass: 'h-1/4',
    contentClass: '',
    primaryElement: 'sidebar'
  },
  {
    value: 'vertical',
    label: '通栏侧边',
    description: '侧栏通顶，Logo 置顶',
    containerClass: '',
    showSidebar: true,
    sidebarClass: 'w-1/4',
    showHeader: true,
    headerClass: 'h-1/4',
    contentClass: '',
    primaryElement: 'sidebar',
    topLogo: true
  }
]

const getSidebarStyle = (layout) => {
  if (layout.primaryElement === 'sidebar') {
    return { backgroundColor: primaryColor.value, opacity: '0.95' }
  } else if (layout.secondaryElement === 'sidebar') {
    return { backgroundColor: lighterPrimaryColor.value, opacity: '0.85' }
  } else {
    return { backgroundColor: lightestPrimaryColor.value, opacity: '0.6' }
  }
}

const getHeaderStyle = (layout) => {
  if (layout.primaryElement === 'header') {
    return { backgroundColor: primaryColor.value, opacity: '0.95' }
  } else if (layout.secondaryElement === 'header') {
    return { backgroundColor: lighterPrimaryColor.value, opacity: '0.85' }
  } else {
    return { backgroundColor: lightestPrimaryColor.value, opacity: '0.6' }
  }
}

const getContentStyle = () => {
  return { backgroundColor: lightestPrimaryColor.value, opacity: '0.5' }
}

const handleLayoutChange = (layout) => {
  emit('update:modelValue', layout)
}
</script>
