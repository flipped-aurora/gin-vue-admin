<script setup>
import { ref, onBeforeUnmount } from 'vue'
import { DropdownMenuRoot, DropdownMenuTrigger } from 'reka-ui'
import DropdownMenuContent from './DropdownMenuContent.vue'
import DropdownMenuItem from './DropdownMenuItem.vue'
import { DROPDOWN_MENU_TRIGGERS } from './index'

defineOptions({ name: 'UiDropdownMenu' })

const props = defineProps({
  // 便捷模式菜单项：{ label, value?, danger?, disabled?, ...自定义字段 }，
  // 选中时整个 item 原样从 @select 抛出；需要完全自定义面板时用 #content 插槽替代
  items: { type: Array, default: () => [] },
  trigger: {
    type: String,
    default: 'click',
    validator: (v) => DROPDOWN_MENU_TRIGGERS.includes(v)
  },
  side: { type: String, default: 'bottom' },
  align: { type: String, default: 'center' },
  sideOffset: { type: Number, default: 6 },
  // 指向触发器的小箭头，不需要时显式关闭
  arrow: { type: Boolean, default: true },
  // 透传给浮层面板的附加 class
  contentClass: { type: null, default: '' }
})

const emit = defineEmits(['select'])

// 展开态受控：click 模式完全交给 reka（trigger 点击 / Esc / 点击外部），
// hover 模式在此基础上叠加「悬停展开、移出延迟收起」；两者共用同一个 open。
const open = ref(false)

// 移出后的收起延迟：给指针从触发器滑向面板留出走位时间，进入面板即取消
let closeTimer = null
const cancelClose = () => {
  if (closeTimer) {
    clearTimeout(closeTimer)
    closeTimer = null
  }
}
const onHoverEnter = () => {
  if (props.trigger !== 'hover') return
  cancelClose()
  open.value = true
}
const onHoverLeave = () => {
  if (props.trigger !== 'hover') return
  cancelClose()
  closeTimer = setTimeout(() => {
    open.value = false
  }, 120)
}
onBeforeUnmount(cancelClose)

// hover 模式关闭时不把焦点还给触发器：reka 的 closeAutoFocus 会程序化聚焦 trigger，
// Chrome 会从面板传播 :focus-visible，留下一圈并非键盘操作引起的 focus ring
const onCloseAutoFocus = (e) => {
  if (props.trigger === 'hover') e.preventDefault()
}
</script>

<template>
  <!-- hover 模式必须非 modal：modal 会给 body 设 pointer-events:none，
       触发器立刻收到 pointerleave → 关闭 → 事件恢复 → pointerenter → 重开，造成闪烁死循环 -->
  <DropdownMenuRoot v-model:open="open" :modal="trigger !== 'hover'">
    <DropdownMenuTrigger
      as-child
      @pointerenter="onHoverEnter"
      @pointerleave="onHoverLeave"
    >
      <!-- 触发器：任意元素 / 组件，reka 通过 asChild 把展开行为合并到插槽根元素 -->
      <slot />
    </DropdownMenuTrigger>

    <DropdownMenuContent
      :side="side"
      :align="align"
      :side-offset="sideOffset"
      :arrow="arrow"
      :class="contentClass"
      @pointerenter="onHoverEnter"
      @pointerleave="onHoverLeave"
      @close-auto-focus="onCloseAutoFocus"
    >
      <slot name="content">
        <DropdownMenuItem
          v-for="(item, index) in items"
          :key="item.value ?? index"
          :disabled="item.disabled"
          :danger="item.danger"
          @select="emit('select', item)"
        >
          {{ item.label }}
        </DropdownMenuItem>
      </slot>
    </DropdownMenuContent>
  </DropdownMenuRoot>
</template>
