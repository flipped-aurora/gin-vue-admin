<script setup>
import {
  DropdownMenuItem,
  DropdownMenuSub,
  DropdownMenuSubTrigger,
  DropdownMenuPortal,
  DropdownMenuSubContent
} from 'reka-ui'
import { visibleItems } from './shared'

defineOptions({ name: 'GvaHorizontalMoreNode' })

const props = defineProps({
  node: { type: Object, required: true }
})
const emit = defineEmits(['select'])

const hasKids = () => visibleItems(props.node.children).length > 0
const itemClass =
  'flex w-full cursor-pointer items-center gap-2 whitespace-nowrap rounded-md px-3 py-2.5 text-[14px] text-base-text outline-none transition-colors data-[highlighted]:bg-muted data-[state=open]:bg-muted'
</script>

<template>
  <DropdownMenuSub v-if="hasKids()">
    <DropdownMenuSubTrigger :class="itemClass">
      <component :is="node.meta.icon" v-if="node.meta.icon" class="h-4 w-4 shrink-0" />
      <span class="flex-1">{{ node.meta.title }}</span>
      <svg-icon icon="lucide:chevron-right" class="h-3.5 w-3.5 shrink-0" />
    </DropdownMenuSubTrigger>
    <DropdownMenuPortal>
      <DropdownMenuSubContent
        :side-offset="4"
        class="z-[3000] min-w-40 rounded-lg border border-border bg-container p-1.5 shadow-sider"
      >
        <HorizontalMoreNode
          v-for="c in visibleItems(node.children)"
          :key="c.name"
          :node="c"
          @select="(k) => emit('select', k)"
        />
      </DropdownMenuSubContent>
    </DropdownMenuPortal>
  </DropdownMenuSub>
  <DropdownMenuItem
    v-else
    :class="itemClass"
    @select="emit('select', node.name)"
  >
    <component :is="node.meta.icon" v-if="node.meta.icon" class="h-4 w-4 shrink-0" />
    {{ node.meta.title }}
  </DropdownMenuItem>
</template>
