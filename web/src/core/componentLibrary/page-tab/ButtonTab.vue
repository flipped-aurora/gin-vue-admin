<script setup>
// 默认风格：扁平药丸；选中 / hover 用主色弱底 + 主色文字。
// 卡片样式（border/shadow）靠 html.gva-card--border / .gva-card--shadow 切换，
// 在下方全局 style 块中声明。相邻 item 间距由上层 flex gap 控制，圆角跟随 --gva-radius。
defineOptions({ name: 'PageTabButton' })

const props = defineProps({
  active: { type: Boolean, default: false },
  class: { type: null, default: '' }
})
</script>

<template>
  <div
    class="gva-tab-btn mb-1.2 relative inline-flex items-center gap-.8 h-8 cursor-pointer select-none  whitespace-nowrap rounded-[var(--gva-radius)] px-3 text-[13px] text-base-text transition-all duration-200 hover:text-active hover:bg-[rgb(var(--primary-color)/0.08)]"
    :class="[active ? 'text-active bg-[rgb(var(--primary-color)/.12)] font-medium' : '', props.class]"
  >
    <slot name="prefix" />
    <slot />
    <slot name="suffix" />
  </div>
</template>

<!-- 卡片模式：border / shadow 写在全局 style 块，由 html 上的 class 切换自动应用 -->
<style>
html.gva-card--border .gva-tab-btn {
  border: 1px solid rgba(0, 0, 0, 0.08);
}
html.gva-card--shadow .gva-tab-btn {
  border: 1px solid rgba(0, 0, 0, 0.08);
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.04), 0 4px 8px rgba(0, 0, 0, 0.04);
}
</style>
