<template>
  <el-menu-item
    :index="routerInfo.name"
    :style="{
          height: sideHeight
        }"
  >
  <template #title>
      <div
        v-if="!isCollapse"
        class="flex items-center"
        :style="{
          height: sideHeight
        }"
      >
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
        <span>{{ routerInfo.meta.title }}</span>
      </div>
      <template v-else>
        <el-icon v-if="routerInfo.meta.icon">
          <component :is="routerInfo.meta.icon" />
        </el-icon>
        <span>{{ routerInfo.meta.title }}</span>
      </template>
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
