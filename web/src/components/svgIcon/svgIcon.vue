<template>
  <template v-if="localIcon">
    <svg aria-hidden="true" width="1em" height="1em" v-bind="bindAttrs">
      <use :xlink:href="'#' + localIcon" rel="external nofollow" />
    </svg>
  </template>
  <template v-else-if="icon">
    <Icon :icon="icon" v-bind="bindAttrs" />
  </template>
</template>

<script setup>
import { computed, useAttrs } from 'vue';
import { Icon } from '@iconify/vue'

/**
 * 使用示例：
 * 本地图标（所有可用的本地图标见控制台输出）：
 * <SvgIcon localIcon="lock" class="text-red-500 text-3xl" />
 * 
 * 在线图标（相关查询网站：https://icones.js.org/ 或是：https://icon-sets.iconify.design/）：
 * <SvgIcon icon="mingcute:love-fill" class="text-red-500 text-3xl" />
 */
defineProps({
  // 通过 symbol id 使用本地注册的 svg 图标
  localIcon: {
    type: String,
    required: false,
    default: ''
  },
  // Iconify 图标名称, 例如: 'mdi:home'
  icon: {
    type: String,
    required: false,
    default: ''
  }
})
const attrs = useAttrs();

const bindAttrs = computed(() => ({
  class: (attrs.class) || '',
  style: (attrs.style) || ''
}))
</script>
