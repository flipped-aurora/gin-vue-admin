<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div class="flex items-center mx-8 gap-4">
    <el-tooltip
      class=""
      effect="dark"
      content="搜索"
      placement="bottom"
    >
      <el-icon class="w-8 h-8 shadow rounded-full border border-gray-600 cursor-pointer">
        <Search />
      </el-icon>
    </el-tooltip>

    <el-tooltip
      class=""
      effect="dark"
      content="刷新"
      placement="bottom"
    >
      <el-icon class="w-8 h-8 shadow rounded-full border border-gray-600 cursor-pointer"
               :class="showRefreshAnmite ? 'animate-spin' : ''"
               @click="toggleRefresh">
        <Refresh />
      </el-icon>
    </el-tooltip>

    <el-tooltip
      class=""
      effect="dark"
      content="切换主题"
      placement="bottom"
    >
      <el-icon v-if="theme === 'dark'" class="w-8 h-8 shadow rounded-full border border-gray-600 cursor-pointer" @click="appStore.toggleTheme(false , $event)">
        <Sunny />
      </el-icon>
      <el-icon v-else class="w-8 h-8 shadow rounded-full border border-gray-600 cursor-pointer" @click="appStore.toggleTheme(true, $event)">
        <Moon />
      </el-icon>
    </el-tooltip>

  </div>
</template>

<script setup>
import { useAppStore } from "@/pinia"
import { storeToRefs } from "pinia"
import { ref } from "vue"
import { emitter } from "@/utils/bus.js";
const appStore = useAppStore()
const { theme } = storeToRefs(appStore)
const showRefreshAnmite = ref(false)
const toggleRefresh = () =>{
  showRefreshAnmite.value = true
  emitter.emit('reload')
  setTimeout(() => {
    showRefreshAnmite.value = false
  }, 1000);
}
</script>

<style scoped lang="scss">

</style>
