<template>
  <div class="flex items-center justify-between py-4 font-inter border-b border-gray-100 dark:border-gray-700 last:border-b-0">
    <div class="flex items-center gap-2">
      <span class="text-sm font-medium text-gray-900 dark:text-white">{{ label }}</span>
      <slot name="suffix"></slot>
    </div>
    <div class="flex items-center setting-controls">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useAppStore } from '@/pinia'

defineOptions({
  name: 'SettingItem'
})

defineProps({
  label: {
    type: String,
    required: true
  }
})

const appStore = useAppStore()
const { config } = storeToRefs(appStore)

const primaryColor = computed(() => config.value.primaryColor)
const primaryColorWithOpacity = computed(() => config.value.primaryColor + '40')
</script>

<style scoped>
.font-inter {
  font-family: 'Inter', sans-serif;
}

.setting-controls {
  ::v-deep(.el-switch) {
    --el-switch-on-color: v-bind(primaryColor);
    --el-switch-off-color: #d1d5db;
  }

  ::v-deep(.el-select) {
    .el-input__wrapper {
      border: 1px solid #e5e7eb;
      border-radius: 6px;
      transition: all 150ms ease-in-out;

      &:hover {
        border-color: v-bind(primaryColor);
        transform: translateY(-1px);
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
      }

      &.is-focus {
        border-color: v-bind(primaryColor);
        box-shadow: 0 0 0 2px v-bind(primaryColorWithOpacity);
      }
    }
  }

  ::v-deep(.el-input-number) {
    .el-input__wrapper {
      border: 1px solid #e5e7eb;
      border-radius: 6px;
      transition: all 150ms ease-in-out;

      &:hover {
        border-color: v-bind(primaryColor);
        transform: translateY(-1px);
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
      }

      &.is-focus {
        border-color: v-bind(primaryColor);
        box-shadow: 0 0 0 2px v-bind(primaryColorWithOpacity);
      }
    }
  }
}

.dark .setting-controls {
  ::v-deep(.el-switch) {
    --el-switch-off-color: #4b5563;
  }

  ::v-deep(.el-select) {
    .el-input__wrapper {
      border-color: #4b5563;
      background-color: #374151;

      &:hover {
        border-color: v-bind(primaryColor);
      }
    }
  }

  ::v-deep(.el-input-number) {
    .el-input__wrapper {
      border-color: #4b5563;
      background-color: #374151;

      &:hover {
        border-color: v-bind(primaryColor);
      }
    }
  }
}
</style>
