<script setup>
import ChromeTabBg from './ChromeTabBg.vue'

// Chrome 风格（移植自 Soybean）：弧形梯形底片 + 标签右侧短竖线。
// 相邻标签靠 -18px 交叠、靠 z-index 抬升当前 / hover 项；
// 竖线在自身 hover / 选中时隐藏，选中标签靠交叠盖住左邻竖线。
// 底片 fill=currentColor，由 .__bg 的 color 控制：选中=主色弱底、hover=更浅、未选=透明，
// 全用 rgb(var(--primary-color)/α) 真实合成色，亮/暗与任意底色上都清晰。
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

  /* 弧形底片填充：fill=currentColor 跟随这里的 color */
  .gva-tab-chrome__bg {
    color: transparent;
  }
  /* 用不透明的浅色填充（而非半透明），才能盖住交叠到下面的相邻标签内容，避免“透出”竖线/×。
     选中=浅主色(primary-50)、hover=浅灰(control-track)，亮/暗均由 token 自适配。 */
  .gva-tab-chrome:hover .gva-tab-chrome__bg {
    color: rgb(var(--control-track-color));
  }
  .gva-tab-chrome.is-active .gva-tab-chrome__bg {
    color: rgb(var(--primary-50-color));
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
