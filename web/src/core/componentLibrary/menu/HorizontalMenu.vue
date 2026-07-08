<script setup>
import { computed, nextTick, onBeforeUnmount, onMounted, ref, watch } from 'vue'
import {
  NavigationMenuRoot,
  NavigationMenuList,
  NavigationMenuItem,
  NavigationMenuTrigger,
  NavigationMenuContent,
  NavigationMenuLink,
  DropdownMenuRoot,
  DropdownMenuTrigger,
  DropdownMenuPortal,
  DropdownMenuContent
} from 'reka-ui'
import { cn, FOCUS_RING } from '../utils'
import { MENU_GAP, visibleItems } from './shared'
import NavMenuPanel from './NavMenuPanel.vue'
import HorizontalMoreNode from './HorizontalMoreNode.vue'

defineOptions({ name: 'GvaHorizontalMenu' })

const props = defineProps({
  items: { type: Array, default: () => [] },
  active: { type: String, default: '' },
  class: { type: null, default: '' }
})
const emit = defineEmits(['select'])

const topItems = computed(() => visibleItems(props.items))
const hasKids = (i) => visibleItems(i.children).length > 0
// 顶级项高亮：激活键等于自身或落在其子树内（父项随激活子路由高亮）
const containsKey = (node, key) => {
  const children = node.children || []
  return node.name === key || children.some((c) => containsKey(c, key))
}
const isActive = (item) => !!props.active && containsKey(item, props.active)

const triggerClass = (activeState) =>
  cn(
    'group inline-flex h-9 shrink-0 appearance-none items-center gap-1.5 rounded-md bg-transparent px-3 text-[14px] text-base-text transition-colors hover:bg-muted data-[state=open]:bg-muted',
    activeState && 'bg-[rgb(var(--primary-color)/0.12)] text-active font-medium',
    FOCUS_RING
  )

const onSelect = (key) => emit('select', key)

/* ---------- 溢出计算：可见区放不下的顶级项收进「更多」 ---------- */
const rootRef = ref(null)
const measureRef = ref(null)
const itemWidths = ref([])
const visibleCount = ref(topItems.value.length)
const MORE_BTN = 36 // 「更多」按钮固定尺寸（w-9）
// 溢出预留：更多按钮宽度 + 其与菜单列表的间距（ml，取 MENU_GAP）
const moreReserve = MORE_BTN + MENU_GAP

const measure = () => {
  if (!measureRef.value) return
  itemWidths.value = [...measureRef.value.children].map(
    (el) => el.getBoundingClientRect().width
  )
  compute()
}
const compute = () => {
  const root = rootRef.value
  if (!root || !itemWidths.value.length) return
  const avail = root.clientWidth
  const total = itemWidths.value.reduce((a, w, i) => a + w + (i ? MENU_GAP : 0), 0)
  if (total <= avail) {
    visibleCount.value = topItems.value.length
    return
  }
  let sum = 0
  let count = 0
  for (let i = 0; i < itemWidths.value.length; i++) {
    const next = sum + (i ? MENU_GAP : 0) + itemWidths.value[i]
    if (next + moreReserve > avail) break
    sum = next
    count++
  }
  visibleCount.value = Math.max(count, 0)
}

const shownItems = computed(() => topItems.value.slice(0, visibleCount.value))
const overflowItems = computed(() => topItems.value.slice(visibleCount.value))
const moreActive = computed(() => overflowItems.value.some((i) => isActive(i)))

let ro = null
onMounted(() => {
  nextTick(measure)
  // web 字体加载完成后项宽会变化，补测一次
  document.fonts?.ready?.then(() => measure())
  ro = new ResizeObserver(() => compute())
  if (rootRef.value) ro.observe(rootRef.value)
})
onBeforeUnmount(() => ro && ro.disconnect())
watch(topItems, () => nextTick(measure))
</script>

<template>
  <div ref="rootRef" :class="cn('relative w-full min-w-0', props.class)">
    <NavigationMenuRoot class="relative flex items-center">
      <NavigationMenuList class="flex items-center" :style="{ gap: MENU_GAP + 'px' }">
        <NavigationMenuItem
          v-for="item in shownItems"
          :key="item.name"
          :value="item.name"
          class="relative"
        >
          <template v-if="hasKids(item)">
            <NavigationMenuTrigger :class="triggerClass(isActive(item))">
              <component :is="item.meta.icon" v-if="item.meta.icon" class="h-4 w-4" />
              {{ item.meta.title }}
              <svg-icon
                icon="lucide:chevron-down"
                class="h-3.5 w-3.5 transition-transform duration-200 group-data-[state=open]:rotate-180"
              />
            </NavigationMenuTrigger>
            <NavigationMenuContent
              class="gva-nav-content absolute left-0 top-full z-[3000] mt-1.5 rounded-lg border border-border bg-container p-1.5 shadow-sider"
            >
              <NavMenuPanel
                :items="visibleItems(item.children)"
                :active="active"
                @select="onSelect"
              />
            </NavigationMenuContent>
          </template>
          <NavigationMenuLink
            v-else
            :active="isActive(item)"
            :class="triggerClass(isActive(item))"
            @select.prevent="onSelect(item.name)"
          >
            <component :is="item.meta.icon" v-if="item.meta.icon" class="h-4 w-4" />
            {{ item.meta.title }}
          </NavigationMenuLink>
        </NavigationMenuItem>
      </NavigationMenuList>

      <!-- 更多：溢出的顶级项级联展开（DropdownMenuSub） -->
      <DropdownMenuRoot v-if="overflowItems.length">
        <DropdownMenuTrigger
          aria-label="更多"
          :class="cn(triggerClass(moreActive), 'ml-2 w-9 justify-center px-0')"
        >
          <svg-icon icon="lucide:ellipsis" class="h-4 w-4" />
        </DropdownMenuTrigger>
        <DropdownMenuPortal>
          <DropdownMenuContent
            :side-offset="6"
            align="end"
            class="z-[3000] min-w-40 rounded-lg border border-border bg-container p-1.5 shadow-sider"
          >
            <HorizontalMoreNode
              v-for="item in overflowItems"
              :key="item.name"
              :node="item"
              @select="onSelect"
            />
          </DropdownMenuContent>
        </DropdownMenuPortal>
      </DropdownMenuRoot>
    </NavigationMenuRoot>

    <!-- 隐藏测量行：按自然宽度渲染全部顶级项，供溢出计算用 -->
    <div
      class="pointer-events-none absolute left-0 top-0 h-0 w-0 overflow-hidden"
      aria-hidden="true"
    >
      <div ref="measureRef" class="flex items-center" :style="{ gap: MENU_GAP + 'px' }">
        <span
          v-for="item in topItems"
          :key="item.name"
          class="inline-flex h-9 shrink-0 items-center gap-1.5 whitespace-nowrap rounded-md px-3 text-[14px]"
        >
          <component :is="item.meta.icon" v-if="item.meta.icon" class="h-4 w-4" />
          {{ item.meta.title }}
          <svg-icon v-if="hasKids(item)" icon="lucide:chevron-down" class="h-3.5 w-3.5" />
        </span>
      </div>
    </div>
  </div>
</template>

<style>
  /* 下拉展开 / 关闭过渡：下滑 + 淡入淡出 */
  .gva-nav-content[data-state='open'] {
    animation: gva-nav-in 220ms ease;
  }
  .gva-nav-content[data-state='closed'] {
    animation: gva-nav-out 160ms ease forwards;
  }
  @keyframes gva-nav-in {
    from {
      opacity: 0;
      transform: translateY(-6px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
  @keyframes gva-nav-out {
    from {
      opacity: 1;
      transform: translateY(0);
    }
    to {
      opacity: 0;
      transform: translateY(-6px);
    }
  }
</style>
