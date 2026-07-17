<script setup>
import { inject, provide, ref } from 'vue'
import { PopoverRoot, PopoverTrigger, PopoverPortal, PopoverContent } from 'reka-ui'
import { FOCUS_RING } from '../utils'
import { MENU_ICON_BUTTON } from './variants'
import { visibleItems } from './shared'
import MenuItem from './MenuItem.vue'

defineOptions({ name: 'GvaMenuFlyout' })

const props = defineProps({
  node: { type: Object, required: true }
})

const ctx = inject('gvaMenuCtx')
const open = ref(false)
let timer = null
const enter = () => {
  clearTimeout(timer)
  open.value = true
}
const leave = () => {
  timer = setTimeout(() => (open.value = false), 120)
}
const children = () => visibleItems(props.node.children)
const isActive = () => ctx.active.value === props.node.name

// 飞出面板：按展开态渲染子树，且用面板内自持的展开态，
// 不回写外层侧栏的 openKeys（避免折叠飞出里的展开污染整栏）。
const flyoutOpen = ref([])
const flyoutToggle = (node) => {
  flyoutOpen.value = flyoutOpen.value.includes(node.name)
    ? flyoutOpen.value.filter((k) => k !== node.name)
    : [...flyoutOpen.value, node.name]
}
provide('gvaMenuCtx', {
  ...ctx,
  collapsed: ref(false),
  openKeys: flyoutOpen,
  toggle: flyoutToggle
})
</script>

<template>
  <PopoverRoot v-model:open="open">
    <PopoverTrigger as-child>
      <button
        type="button"
        :class="[MENU_ICON_BUTTON, FOCUS_RING, isActive() && 'text-active']"
        :style="{ height: ctx.itemHeight.value + 'px' }"
        :title="node.meta.title"
        @mouseenter="enter"
        @mouseleave="leave"
        @click="ctx.select(node.name)"
      >
        <component :is="node.meta.icon" v-if="node.meta.icon" class="h-5 w-5" />
        <span v-else class="text-[13px]">{{ node.meta.title[0] }}</span>
      </button>
    </PopoverTrigger>
    <PopoverPortal>
      <PopoverContent
        side="right"
        align="start"
        :side-offset="10"
        class="z-[3000] min-w-44 rounded-lg border border-border bg-container p-1.5 shadow-sider"
        @mouseenter="enter"
        @mouseleave="leave"
      >
        <div class="px-3 py-1.5 text-[13px] font-semibold text-base-text">
          {{ node.meta.title }}
        </div>
        <MenuItem
          v-for="child in children()"
          :key="child.name"
          :node="child"
          :depth="1"
        />
      </PopoverContent>
    </PopoverPortal>
  </PopoverRoot>
</template>
