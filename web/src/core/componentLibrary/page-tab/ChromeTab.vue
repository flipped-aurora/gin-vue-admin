<script setup>
import ChromeTabBg from './ChromeTabBg.vue'

// Chrome 风格（移植自 Soybean）：弧形梯形底片 + 标签右侧短竖线。
// 相邻标签靠 -18px 交叠、靠 z-index 抬升当前 / hover 项；
// 竖线在自身 hover / 选中时隐藏，选中标签靠交叠盖住左邻竖线。
// 弧形底片 fill=currentColor，由 .__bg 的 color 控制：仅选中显示（主色弱底），未选透明。
// hover 对齐真实 Chrome：非激活标签 hover 出「内嵌圆角胶囊」，不点亮整个弧形底片；
// 胶囊用全局统一的 hover 淡底 rgb(var(--primary-color)/0.08)（与 ButtonTab / el-tree 相通）。
defineOptions({ name: 'PageTabChrome' })

const props = defineProps({
  active: { type: Boolean, default: false },
  class: { type: null, default: '' }
})
</script>

<template>
  <div
    class="gva-tab-chrome relative -mr-[18px] inline-flex h-full cursor-pointer select-none items-center justify-center gap-2 whitespace-nowrap px-6 text-[13px] transition-colors duration-200"
    :class="[{ 'is-active': active }, props.class]"
  >
    <div class="gva-tab-chrome__bg pointer-events-none absolute inset-0 z-[-1] h-full w-full">
      <ChromeTabBg />
    </div>
    <!-- 非激活 hover 的内嵌胶囊（仿真实 Chrome：hover 是圆角胶囊，选中才是弧形选项卡） -->
    <div class="gva-tab-chrome__hover pointer-events-none absolute z-[-1]" />
    <slot name="prefix" />
    <slot />
    <slot name="suffix" />
    <div class="gva-tab-chrome__divider" />
  </div>
</template>

<style scoped>
  .gva-tab-chrome {
    z-index: 0;
    color: rgb(var(--base-text-color));
  }
  .gva-tab-chrome:hover {
    z-index: 9;
  }
  .gva-tab-chrome.is-active {
    z-index: 10;
    color: rgb(var(--primary-color));
  }

  /* 弧形底片填充：fill=currentColor 跟随这里的 color。仅选中显示；
     用不透明的浅主色(primary-50)填充（而非半透明），才能盖住交叠到下面的相邻标签内容，
     避免“透出”竖线/×，亮/暗由 token 自适配。 */
  .gva-tab-chrome__bg {
    color: transparent;
  }
  .gva-tab-chrome.is-active .gva-tab-chrome__bg {
    color: rgb(var(--primary-50-color));
  }

  /* 非激活 hover：内嵌圆角胶囊（仿真实 Chrome）。用全局统一 hover 淡底（主色 8% wash），
     不与弧形底片同形；相邻竖线在 hover 时已隐藏，半透明底下只有标签栏背景，不会透出杂物。 */
  .gva-tab-chrome__hover {
    inset: 5px 12px;
    border-radius: 8px;
    background: rgb(var(--primary-color) / 0.08);
    opacity: 0;
    transition: opacity 0.2s;
  }
  .gva-tab-chrome:not(.is-active):hover .gva-tab-chrome__hover {
    opacity: 1;
  }

  /* 短竖线：每个标签右侧一根；自身 hover / 选中时隐藏，
     选中标签靠 -18px 交叠 + z-index 盖住其左邻标签的竖线（与 Soybean 一致） */
  .gva-tab-chrome__divider {
    position: absolute;
    right: 7px;
    top: 50%;
    width: 1px;
    height: 14px;
    transform: translateY(-50%);
    background-color: rgb(var(--border-color));
  }
  .gva-tab-chrome:hover .gva-tab-chrome__divider,
  .gva-tab-chrome.is-active .gva-tab-chrome__divider,
  .gva-tab-chrome:has(+ .gva-tab-chrome:hover) .gva-tab-chrome__divider,
  .gva-tab-chrome:has(+ .gva-tab-chrome.is-active) .gva-tab-chrome__divider {
    opacity: 0;
  }
</style>
