<template>
  <div class="grid grid-cols-2 gap-6 font-inter px-6">
    <div
      v-for="layout in layoutModes"
      :key="layout.value"
      class="bg-white dark:bg-gray-700 border-2 border-gray-200 dark:border-gray-600 rounded-xl p-6 cursor-pointer transition-all duration-150 ease-in-out hover:transform hover:-translate-y-1 hover:shadow-xl"
      :class="{
        'ring-2 ring-offset-2 ring-offset-gray-50 dark:ring-offset-gray-900 transform -translate-y-1 shadow-xl': modelValue === layout.value
      }"
      :style="modelValue === layout.value ? {
        borderColor: primaryColor,
        ringColor: primaryColor + '40'
      } : {}"
      @click="handleLayoutChange(layout.value)"
    >
      <div class="flex justify-center mb-5">
        <div
          class="w-28 h-20 bg-gray-50 dark:bg-gray-600 border border-gray-200 dark:border-gray-500 rounded-lg p-2 flex gap-1.5 shadow-inner"
          :class="layout.containerClass"
        >
          <div
            v-if="layout.showSidebar"
            class="rounded-sm"
            :class="[layout.sidebarClass]"
            :style="getSidebarStyle(layout)"
          ></div>

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
        <span class="block text-base font-semibold text-gray-900 dark:text-white mb-2" :class="{ 'text-current': modelValue === layout.value }" :style="modelValue === layout.value ? { color: primaryColor } : {}">{{ layout.label }}</span>
        <span class="block text-sm text-gray-500 dark:text-gray-400">{{ layout.description }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useAppStore } from '@/pinia'

defineOptions({
  name: 'LayoutModeCard'
})

const props = defineProps({
  modelValue: {
    type: String,
    default: 'normal'
  }
})

const emit = defineEmits(['update:modelValue'])

const appStore = useAppStore()
const { config } = storeToRefs(appStore)

const primaryColor = computed(() => config.value.primaryColor)
const lighterPrimaryColor = computed(() => {
  const hex = config.value.primaryColor.replace('#', '')
  const r = parseInt(hex.substr(0, 2), 16)
  const g = parseInt(hex.substr(2, 2), 16)
  const b = parseInt(hex.substr(4, 2), 16)
  return `rgba(${r}, ${g}, ${b}, 0.7)`
})
const lightestPrimaryColor = computed(() => {
  const hex = config.value.primaryColor.replace('#', '')
  const r = parseInt(hex.substr(0, 2), 16)
  const g = parseInt(hex.substr(2, 2), 16)
  const b = parseInt(hex.substr(4, 2), 16)
  return `rgba(${r}, ${g}, ${b}, 0.4)`
})

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
    showRightSidebar: false,
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
    showRightSidebar: false,
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
    showRightSidebar: true,
    rightSidebarClass: 'w-1/5',
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
    showRightSidebar: false,
    primaryElement: 'sidebar'
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

const getContentStyle = (layout) => {
  return { backgroundColor: lightestPrimaryColor.value, opacity: '0.5' }
}

const getRightSidebarStyle = (layout) => {
  if (layout.primaryElement === 'rightSidebar') {
    return { backgroundColor: primaryColor.value, opacity: '0.95' }
  } else if (layout.secondaryElement === 'rightSidebar') {
    return { backgroundColor: lighterPrimaryColor.value, opacity: '0.85' }
  } else {
    return { backgroundColor: lightestPrimaryColor.value, opacity: '0.6' }
  }
}

const handleLayoutChange = (layout) => {
  emit('update:modelValue', layout)
}
</script>

<style scoped>
.font-inter {
  font-family: 'Inter', sans-serif;
}

.flex-col {
  flex-direction: column;
}

.w-1\/4 {
  width: 25%;
}

.w-1\/3 {
  width: 33.333333%;
}

.w-1\/5 {
  width: 20%;
}

.h-1\/4 {
  height: 25%;
}

.h-1\/3 {
  height: 33.333333%;
}

@media (max-width: 480px) {
  .grid-cols-2 {
    grid-template-columns: repeat(1, minmax(0, 1fr));
  }
}
</style>
