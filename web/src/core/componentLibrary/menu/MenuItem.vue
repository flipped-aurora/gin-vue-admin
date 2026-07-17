<script setup>
import { computed, inject } from 'vue'
import { CollapsibleRoot, CollapsibleTrigger, CollapsibleContent } from 'reka-ui'
import { cn, FOCUS_RING } from '../utils'
import { menuItemVariants } from './variants'
import { menuIndent, visibleItems } from './shared'
import MenuFlyout from './MenuFlyout.vue'

defineOptions({ name: 'GvaMenuItem' })

const props = defineProps({
  node: { type: Object, required: true },
  depth: { type: Number, default: 0 }
})

// 由 Menu.vue provide 的运行时上下文（全为 ref + 回调）
const ctx = inject('gvaMenuCtx')

const children = computed(() => visibleItems(props.node.children))
const hasChildren = computed(() => children.value.length > 0)
const isActive = computed(() => ctx.active.value === props.node.name)
const isOpen = computed(() => ctx.openKeys.value.includes(props.node.name))
// group 风格：顶层分支渲染成「分组标题」
const isGroupHeader = computed(
  () => ctx.theme.value === 'group' && props.depth === 0 && hasChildren.value
)
// 折叠图标栏：折叠态 + 顶层 + 有子级 → 飞出
const isFlyout = computed(
  () => ctx.collapsed.value && props.depth === 0 && hasChildren.value
)
// 折叠态顶层叶子：只显图标
const iconOnly = computed(() => ctx.collapsed.value && props.depth === 0)

const rowStyle = computed(() => ({
  height: ctx.itemHeight.value + 'px',
  // design 全宽、按层级右移；其余留基础左内边距（折叠图标态无内边距）
  paddingLeft: iconOnly.value ? '0px' : menuIndent(props.depth, ctx.theme.value)
}))

const rowClass = computed(() =>
  cn(
    menuItemVariants({
      theme: ctx.theme.value,
      active: isActive.value,
      role: isGroupHeader.value ? 'header' : 'item'
    }),
    iconOnly.value ? 'justify-center px-0' : 'px-3',
    FOCUS_RING
  )
)

const onLeaf = () => ctx.select(props.node.name)
const onToggle = () => ctx.toggle(props.node)
</script>

<template>
  <!-- 折叠飞出 -->
  <MenuFlyout v-if="isFlyout" :node="node" />

  <!-- 分支（含 group 分组标题）：Collapsible -->
  <CollapsibleRoot v-else-if="hasChildren" :open="isOpen" @update:open="onToggle">
    <CollapsibleTrigger :class="cn(rowClass, 'text-left')" :style="rowStyle">
      <component
        :is="node.meta.icon"
        v-if="node.meta.icon && !isGroupHeader"
        class="h-[18px] w-[18px] shrink-0"
      />
      <span class="flex-1 truncate">{{ node.meta.title }}</span>
      <svg-icon
        icon="lucide:chevron-down"
        class="h-4 w-4 shrink-0 transition-transform duration-200"
        :class="isOpen ? 'rotate-180' : ''"
      />
    </CollapsibleTrigger>
    <CollapsibleContent class="gva-collapsible-content">
      <MenuItem
        v-for="child in children"
        :key="child.name"
        :node="child"
        :depth="depth + 1"
      />
    </CollapsibleContent>
  </CollapsibleRoot>

  <!-- 叶子 -->
  <button
    v-else
    type="button"
    :class="rowClass"
    :style="rowStyle"
    :title="iconOnly ? node.meta.title : null"
    @click="onLeaf"
  >
    <component
      :is="node.meta.icon"
      v-if="node.meta.icon"
      class="h-[18px] w-[18px] shrink-0"
    />
    <span v-else-if="iconOnly" class="text-[13px]">{{ node.meta.title[0] }}</span>
    <span v-if="!iconOnly" class="flex-1 truncate text-left">{{
      node.meta.title
    }}</span>
  </button>
</template>

<style>
  /* 展开 / 收起过渡：高度取 reka 提供的 --reka-collapsible-content-height */
  .gva-collapsible-content {
    overflow: hidden;
  }
  .gva-collapsible-content[data-state='open'] {
    animation: gva-collapse-down 200ms ease-out;
  }
  .gva-collapsible-content[data-state='closed'] {
    animation: gva-collapse-up 200ms ease-out;
  }
  @keyframes gva-collapse-down {
    from {
      height: 0;
    }
    to {
      height: var(--reka-collapsible-content-height);
    }
  }
  @keyframes gva-collapse-up {
    from {
      height: var(--reka-collapsible-content-height);
    }
    to {
      height: 0;
    }
  }
</style>
