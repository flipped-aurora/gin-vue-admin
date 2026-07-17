<script setup>
import { NavigationMenuLink } from 'reka-ui'
import { cn } from '../utils'
import { menuIndent, visibleItems } from './shared'

defineOptions({ name: 'GvaNavMenuPanel' })

const props = defineProps({
  items: { type: Array, default: () => [] },
  active: { type: String, default: '' },
  depth: { type: Number, default: 0 }
})
const emit = defineEmits(['select'])

const hasKids = (i) => visibleItems(i.children).length > 0
const linkClass = (name) =>
  cn(
    'flex w-full cursor-pointer items-center gap-2 whitespace-nowrap rounded-md py-2.5 pr-6 text-[14px] outline-none transition-colors hover:bg-muted',
    props.active === name ? 'text-active font-medium' : 'text-base-text'
  )
</script>

<template>
  <ul class="flex min-w-max flex-col gap-1">
    <li v-for="node in visibleItems(items)" :key="node.name">
      <!-- 分支：分组标题 + 递归子列表 -->
      <template v-if="hasKids(node)">
        <div
          class="whitespace-nowrap pb-1 pr-6 pt-2 text-[12px] font-semibold text-muted-foreground"
          :style="{ paddingLeft: menuIndent(depth) }"
        >
          {{ node.meta.title }}
        </div>
        <NavMenuPanel
          :items="visibleItems(node.children)"
          :active="active"
          :depth="depth + 1"
          @select="(k) => emit('select', k)"
        />
      </template>
      <!-- 叶子：NavigationMenuLink -->
      <NavigationMenuLink
        v-else
        :active="active === node.name"
        :class="linkClass(node.name)"
        :style="{ paddingLeft: menuIndent(depth) }"
        @select.prevent="emit('select', node.name)"
      >
        <component :is="node.meta.icon" v-if="node.meta.icon" class="h-4 w-4 shrink-0" />
        {{ node.meta.title }}
      </NavigationMenuLink>
    </li>
  </ul>
</template>
