<script setup>
import {
  DropdownMenuArrow,
  DropdownMenuContent,
  DropdownMenuPortal
} from 'reka-ui'
import { cn } from '../utils'

defineOptions({ name: 'UiDropdownMenuContent', inheritAttrs: false })

// 面板：bg-container 浮层底 + shadow-sider，z-[3000] 保证盖过 EP 浮层；
// 进出场按官方文档用 --reka-dropdown-menu-content-transform-origin 做缩放淡入 / 淡出，
// 从触发器方向弹出。inheritAttrs 关闭后用 v-bind="$attrs" 把外部事件
// （如 hover 触发的 pointerenter/leave）落到真正的浮层元素上，而不是不渲染节点的 Portal。
const props = defineProps({
  side: { type: String, default: 'bottom' },
  align: { type: String, default: 'center' },
  sideOffset: { type: Number, default: 6 },
  // 指向触发器的小箭头（官方 anatomy 的 DropdownMenuArrow），不需要时显式关闭
  arrow: { type: Boolean, default: true },
  class: { type: null, default: '' }
})
</script>

<template>
  <DropdownMenuPortal>
    <DropdownMenuContent
      v-bind="$attrs"
      :side="side"
      :align="align"
      :side-offset="sideOffset"
      :class="cn(
        'z-[3000] min-w-40 rounded-[10px] border border-border bg-container p-1.5 text-base-text shadow-sider',
        'origin-[var(--reka-dropdown-menu-content-transform-origin)] data-[state=open]:animate-[popper-in_0.2s_cubic-bezier(0.16,1,0.3,1)] data-[state=closed]:animate-[popper-out_0.15s_ease-in]',
        props.class
      )"
    >
      <slot />
      <!-- 箭头与面板一体化，两步处理（本地坐标经 reka 各朝向的旋转后方向恒定）：
           1. -translate-y-px 向面板内嵌 1px，盖住面板 border 从箭头底部穿过的缝隙；
           2. drop-shadow 沿三角形轮廓给两条斜边描 1px border 色（底边嵌进面板不受影响），
              否则 fill 与面板同色时箭头在同色背景上不可见 -->
      <DropdownMenuArrow
        v-if="arrow"
        :width="14"
        :height="7"
        class="-translate-y-px fill-container [filter:drop-shadow(0_1px_0_rgb(var(--border-color)))]"
      />
    </DropdownMenuContent>
  </DropdownMenuPortal>
</template>
