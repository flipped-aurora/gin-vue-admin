<script setup>
import { computed } from 'vue'
import { cn } from '../utils'
import ButtonTab from './ButtonTab.vue'
import ChromeTab from './ChromeTab.vue'
import SliderTab from './SliderTab.vue'
import { PAGE_TAB_MODES } from './shared'

// 单个标签页（纯展示）风格切换由 mode 决定（button / chrome / slider），三种模式各是一个无逻辑的展示子组件。
// 上层在 <g-page-tab> 上绑定的原生事件（@click / @mousedown / @contextmenu 等）会经 attribute fallthrough 透传到根元素；关闭则是显式 close 事件。
defineOptions({ name: 'PageTab' })

const props = defineProps({
  // 风格：button(描边胶囊) / chrome(浏览器标签) / slider(底部指示条)
  mode: {
    type: String,
    default: 'button',
    validator: (v) => PAGE_TAB_MODES.includes(v)
  },
  // 是否选中态（由上层根据“当前激活的 key”计算后传入，组件自身不持有激活状态）
  active: { type: Boolean, default: false },
  // 是否显示关闭按钮
  closable: { type: Boolean, default: true },
  // 透传到根：外部布局 / 覆盖用，cn() 最后合并
  class: { type: null, default: '' },
  // 各模式独立的覆盖 class
  buttonClass: { type: null, default: '' },
  chromeClass: { type: null, default: '' },
  sliderClass: { type: null, default: '' }
})

const emit = defineEmits(['close'])

const tabComponent = computed(
  () => ({ button: ButtonTab, chrome: ChromeTab, slider: SliderTab })[props.mode]
)

const modeClass = computed(
  () =>
    ({
      button: props.buttonClass,
      chrome: props.chromeClass,
      slider: props.sliderClass
    })[props.mode]
)

// 关闭按钮独立于标签本体：阻断冒泡，避免触发上层的“切换 / 中键关闭 / 右键菜单”
const handleClose = (e) => {
  e.stopPropagation()
  emit('close')
}
</script>

<template>
  <component
    :is="tabComponent"
    :active="active"
    :class="cn(modeClass, props.class)"
  >
    <template #prefix>
      <slot name="prefix" />
    </template>
    <slot />
    <template #suffix>
      <slot name="suffix">
        <button
          v-if="closable"
          type="button"
          aria-label="关闭标签"
          class="-mr-1 ml-0.5 inline-flex h-4 w-4 shrink-0 appearance-none items-center justify-center rounded-full bg-transparent text-muted-foreground transition-colors hover:text-base-text gva-tab-close-btn"
          @click.stop="handleClose"
          @mousedown.stop
          @pointerdown.stop
        >
          <svg-icon icon="lucide:x" class="h-3 w-3" />
        </button>
      </slot>
    </template>
  </component>
</template>

<style scoped>
  /* 关闭按钮 hover 圆：用 muted-foreground 淡色替代 control-track（与标签 hover 底色同色），
     亮/暗模式下都与 control-track 形成可见反差。 */
  .gva-tab-close-btn:hover {
    background-color: rgb(var(--muted-foreground-color) / 0.2);
  }
</style>
