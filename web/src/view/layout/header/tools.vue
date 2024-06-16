<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div class="flex items-center mx-4 gap-4">
    <el-tooltip
      class=""
      effect="dark"
      content="搜索"
      placement="bottom"
    >
      <el-icon @click="handleCommand" class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid">
        <Search />
      </el-icon>
    </el-tooltip>

    <el-tooltip
      class=""
      effect="dark"
      content="系统设置"
      placement="bottom"
    >
      <el-icon class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid" @click="toggleSetting">
        <Setting />
      </el-icon>
    </el-tooltip>

    <el-tooltip
      class=""
      effect="dark"
      content="刷新"
      placement="bottom"
    >
      <el-icon
        class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid"
        :class="showRefreshAnmite ? 'animate-spin' : ''"
        @click="toggleRefresh"
      >
        <Refresh />
      </el-icon>
    </el-tooltip>
    <el-tooltip
      class=""
      effect="dark"
      content="切换主题"
      placement="bottom"
      :disabled="appStore.theme === 'auto'"
    >
      <el-icon v-if="appStore.theme === 'dark'" class="w-8 h-8 shadow rounded-full border border-gray-600 cursor-pointer border-solid" @click="appStore.toggleTheme(false )">
        <Sunny />
      </el-icon>
      <el-icon v-else class="w-8 h-8 shadow rounded-full border border-gray-200 cursor-pointer border-solid" @click="appStore.toggleTheme(true)">
        <Moon />
      </el-icon>
    </el-tooltip>

    <gva-setting v-model:drawer="showSettingDrawer"></gva-setting>
    <command-menu ref="command" />
  </div>
</template>

<script setup>

import { useAppStore } from "@/pinia"
import GvaSetting from "@/view/layout/setting/index.vue"
import { ref } from "vue"
import { emitter } from "@/utils/bus.js";
import CommandMenu from "@/components/commandMenu/index.vue";

const appStore = useAppStore()
const showSettingDrawer = ref(false)
const showRefreshAnmite = ref(false)
const toggleRefresh = () =>{
  showRefreshAnmite.value = true
  emitter.emit('reload')
  setTimeout(() => {
    showRefreshAnmite.value = false
  }, 1000);
}

const toggleSetting = () => {
  showSettingDrawer.value = true
}


const first = ref("");
const command = ref();

const handleCommand = () => {
  command.value.open();
};
const initPage = () => {
  // 判断当前用户的操作系统
  if (window.localStorage.getItem("osType") === "WIN") {
    first.value = "Ctrl";
  } else {
    first.value = "⌘";
  }
  // 当用户同时按下ctrl和k键的时候
  const handleKeyDown = (e) => {
    if (e.ctrlKey && e.key === "k") {
      // 阻止浏览器默认事件
      e.preventDefault();
      handleCommand();
    }
  };
  window.addEventListener("keydown", handleKeyDown);
};

initPage();




</script>

<style scoped lang="scss">

</style>
