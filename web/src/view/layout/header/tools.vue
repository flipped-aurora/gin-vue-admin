<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div class="flex items-center mx-4 gap-4">
    <el-tooltip
        class=""
        effect="dark"
        :content="t('layout.tools.videoTutorial')"
        placement="bottom"
    >
      <el-dropdown @command="toDoc">
      <el-icon class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid">
        <Film />
      </el-icon>

      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item v-for="item in videoList" :key="item.link" :command="item.link">{{ item.title }}</el-dropdown-item>
        </el-dropdown-menu>
      </template>
      </el-dropdown>
    </el-tooltip>

    <el-tooltip
      class=""
      effect="dark"
      :content="t('layout.tools.search')"
      placement="bottom"
    >
      <el-icon class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid" @click="handleCommand">
        <Search />
      </el-icon>
    </el-tooltip>

    <el-tooltip
      class=""
      effect="dark"
      :content="t('layout.tools.systemSettings')"
      placement="bottom"
    >
      <el-icon class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid" @click="toggleSetting">
        <Setting />
      </el-icon>
    </el-tooltip>

    <el-tooltip
      class=""
      effect="dark"
      :content="t('layout.tools.refresh')"
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
      :content="t('layout.tools.switchTheme')"
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

    <el-tooltip
      class=""
      effect="dark"
      :content="t('layout.tools.changeLanguage')"
      placement="bottom"
      :disabled="appStore.theme === 'auto'"
    >
      <SelectLang @success="changeSuccess">
          <el-icon  class="w-8 h-8 shadow rounded-full border border-gray-200 dark:border-gray-600 cursor-pointer border-solid" @click="appStore.toggleLang">
            <language />
          </el-icon>
      </SelectLang>
    </el-tooltip>

    <gva-setting v-model:drawer="showSettingDrawer" />
    <command-menu ref="command" />
  </div>
</template>

<script setup>

import { useAppStore } from "@/pinia"
import GvaSetting from "@/view/layout/setting/index.vue"
import { ref } from "vue"
import { emitter } from "@/utils/bus.js";
import CommandMenu from "@/components/commandMenu/index.vue";
import { useI18n } from 'vue-i18n'
const { t } = useI18n() // added by mohamed hassan to support multilanguage

import SelectLang from '@/components/i18n/selectLanguage.vue'

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

const changeSuccess = () => {
  window.location.reload();
}


const videoList = [
  {
    title: t('view.layout.header.tools.video01'),
    link:"https://www.bilibili.com/video/BV1jx4y1s7xx",
  },
  {
    title: t('view.layout.header.tools.video02'),
    link:"https://www.bilibili.com/video/BV1sr421K7sv",
  },
  {
    title: t('view.layout.header.tools.video03'),
    link:"https://www.bilibili.com/video/BV1iH4y1c7Na",
  },
  {
    title: t('view.layout.header.tools.video04'),
    link:"https://www.bilibili.com/video/BV1UZ421T7fV",
  },
  {
    title: t('view.layout.header.tools.video05'),
    link:"https://www.bilibili.com/video/BV1NE4m1977s",
  },
  {
    title: t('view.layout.header.tools.video06'),
    link:"https://www.bilibili.com/video/BV17i421a7DE",
  },
  {
    title: t('view.layout.header.tools.video07'),
    link:"https://www.bilibili.com/video/BV1Yw4m1k7fg",
  },
  {
    title: t('view.layout.header.tools.video08'),
    link:"https://www.bilibili.com/video/BV12y411i7oE",
  },
  {
    title: t('view.layout.header.tools.video09'),
    link:"https://www.bilibili.com/video/BV1ZM4m1y7i3",
  },
  {
    title: t('view.layout.header.tools.video10'),
    link:"https://www.bilibili.com/video/BV1WS42197DZ",
  },
  {
    title: t('view.layout.header.tools.video11'),
    link:"https://www.bilibili.com/video/BV1NE4m1979c",
  },
  {
    title: t('view.layout.header.tools.video12'),
    link:"https://www.bilibili.com/video/BV1Sw4m1k746",
  },
  {
    title: t('view.layout.header.tools.video13'),
    link:"https://www.bilibili.com/video/BV1Ki421a7X2",
  },
  {
    title: t('view.layout.header.tools.video14'),
    link:"https://www.bilibili.com/video/BV1Lx4y1s77D",
  }
]


</script>

<style scoped lang="scss">

</style>
