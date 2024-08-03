<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/7
!-->

<template>
  <div class="flex items-center mx-4 gap-4">
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
      <el-dropdown
        @command="handleSetLanguage"
      >
        <el-icon class="w-8 h-8 shadow rounded-full border border-gray-200 cursor-pointer border-solid" @click="appStore.toggleLang">
          <language />
        </el-icon>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item
              command="zh"
            >
              简体中文
            </el-dropdown-item>
            <el-dropdown-item
              command="en"
            >
              English
            </el-dropdown-item>
            <el-dropdown-item
              command="ar"
            >
              العربية
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
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
import Cookies from "js-cookie";
import {ElMessage} from "element-plus"; // added by mohamed hassan to support multilanguage
const i18n = useI18n() // added by mohamed hassan to support multilanguage
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const { t } = useI18n() // added by mohamed hassan to support multilanguage

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

const handleSetLanguage = (lang) => {
  // console.log('handleSetLanguage() called with value: ' + lang)
  i18n.locale.value = lang

  userStore.setLanguage(lang)

  // console.log('userStore handleSetLanguage() called with value: ' + userStore.getLanguage())

  Cookies.set('language', lang)

  // if (lang === 'ar') {
  //   console.log('Arabic language selected changing to RTL')
  //   document.querySelector('html').classList.add('is-rtl')
  // } else {
  //   console.log('Non Arabic language selected changing to LTR')
  //   document.querySelector('html').classList.add('is-ltr')
  // }

  // const htmlEl = document.querySelector('html')

  // if (this.$i18n.locale === 'ar') {
  //   console.log('change language to arabic and ltr to rtl')
  //   htmlEl.setAttribute('dir', 'rtl')
  // } else {
  //   console.log('change language to english and rtl to ltr')
  //   htmlEl.setAttribute('dir', 'ltr')
  // }

  // htmlEl.setAttribute('lang', lang)

  ElMessage({
    message: t('general.langSwitch'),
    type: 'success'
  })

  // this.$emit('handerevent')
  window.location.reload()
}

initPage();





</script>

<style scoped lang="scss">

</style>
