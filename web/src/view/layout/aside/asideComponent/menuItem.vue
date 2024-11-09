<template>
  <el-menu-item
    :index="routerInfo.name"
    class="dark:text-slate-300 overflow-hidden"
    :style="{
      height: sideHeight
    }"
  >
    <el-icon v-if="routerInfo.meta.icon">
      <component :is="routerInfo.meta.icon" />
    </el-icon>
    <template #title>
      {{ routerInfo.meta.title }}
    </template>
  </el-menu-item>
</template>

<script setup>
  import { computed } from 'vue'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { config } = storeToRefs(appStore)

  defineOptions({
    name: 'MenuItem'
  })

  defineProps({
    routerInfo: {
      default: function () {
        return null
      },
      type: Object
    }
  })

  const sideHeight = computed(() => {
    return config.value.layout_side_item_height + 'px'
  })
</script>

<style lang="scss"></style>
