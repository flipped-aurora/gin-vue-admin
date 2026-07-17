<script setup>
  import { computed, ref } from 'vue'
  import { Primitive } from 'reka-ui'
  import { cn, FOCUS_RING } from '@/core/componentLibrary/utils'
  import SvgIcon from '@/components/svgIcon/svgIcon.vue'

  defineOptions({ name: 'IconButton' })

  const props = defineProps({
    // Iconify 图标名（外壳层统一走在线 lucide 集，如 'lucide:search'）
    icon: { type: String, default: '' },
    // hover / text 模式下展开显示的文案
    label: { type: String, default: '' },
    // icon：常态显示图标、hover 时morph为文案；text：始终显示文案
    mode: {
      type: String,
      default: 'icon',
      validator: (v) => ['icon', 'text'].includes(v)
    },
    disabled: { type: Boolean, default: false },
    // 透传给图标的附加 class（如刷新态的 animate-spin）
    iconClass: { type: null, default: '' },
    // 透传给 reka Primitive：默认渲染 <button>，asChild 时把行为合并到插槽根元素
    as: { type: [String, Object], default: 'button' },
    asChild: { type: Boolean, default: false },
    class: { type: null, default: '' }
  })

  const emit = defineEmits(['click'])

  const hovered = ref(false)

  // 展开条件：text 模式恒展开；无图标时只能展开显示文案；icon 模式 hover 时展开
  const expanded = computed(
    () => props.mode === 'text' || !props.icon || hovered.value
  )
</script>

<template>
  <Primitive
    :as="as"
    :as-child="asChild"
    :disabled="disabled || undefined"
    :class="
      cn(
        'group relative inline-flex h-9 min-w-9 cursor-pointer appearance-none items-center justify-center rounded-[var(--gva-radius)] bg-transparent px-2 text-muted-foreground outline-none transition-colors hover:bg-black/10 hover:text-base-text disabled:pointer-events-none disabled:opacity-50 dark:hover:bg-white/10',
        FOCUS_RING,
        props.class
      )
    "
    @mouseenter="hovered = true"
    @mouseleave="hovered = false"
    @click="(e) => emit('click', e)"
  >
    <span class="flex items-center justify-center">
      <!-- 图标列：展开时宽度收拢为 0，图标同步向左滑出（收起时反向从左滑回） -->
      <span
        v-if="icon"
        class="grid transition-[grid-template-columns] duration-300 ease-[cubic-bezier(0.4,0,0.2,1)]"
        :class="expanded && label ? 'grid-cols-[0fr]' : 'grid-cols-[1fr]'"
      >
        <span class="flex min-w-0 items-center justify-center overflow-hidden">
          <svg-icon
            :icon="icon"
            :class="
              cn(
                'text-[18px] leading-none transition-[transform,opacity] duration-300 ease-[cubic-bezier(0.4,0,0.2,1)]',
                expanded && label
                  ? '-translate-x-6 opacity-0'
                  : 'translate-x-0 opacity-100',
                iconClass
              )
            "
          />
        </span>
      </span>

      <!-- 文案列：展开时宽度撑开，文案从右侧滑入（收起时向右滑出） -->
      <span
        v-if="label"
        class="grid transition-[grid-template-columns] duration-300 ease-[cubic-bezier(0.4,0,0.2,1)]"
        :class="expanded ? 'grid-cols-[1fr]' : 'grid-cols-[0fr]'"
      >
        <span class="flex min-w-0 items-center justify-center overflow-hidden">
          <span
            class="whitespace-nowrap px-1 text-[13px] font-medium leading-none transition-[transform,opacity] duration-300 ease-[cubic-bezier(0.4,0,0.2,1)]"
            :class="expanded ? 'translate-x-0 opacity-100' : 'translate-x-6 opacity-0'"
          >
            {{ label }}
          </span>
        </span>
      </span>
    </span>
  </Primitive>
</template>
