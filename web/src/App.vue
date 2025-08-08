<template>
  <div
    id="app"
    class="bg-gray-50 text-slate-700 dark:text-slate-500 dark:bg-slate-800"
  >
    <el-config-provider :locale="languages[locale]">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script setup>
  import zh from 'element-plus/dist/locale/zh-cn.mjs'
  import en from 'element-plus/dist/locale/en.mjs'
  import ar from 'element-plus/dist/locale/ar.mjs'
  import ru from 'element-plus/dist/locale/ru.mjs'
  import zh_tw from 'element-plus/dist/locale/zh-tw.mjs'
  import { useAppStore } from '@/pinia'
  import { useI18n } from 'vue-i18n'

  useAppStore()
  defineOptions({
    name: 'App'
  })

  const languages = {
    'zh-cn': zh,
    'zh-TW': zh_tw,
    en: en,
    ar: ar,
    ru: ru
  }

  const { locale } = useI18n()
</script>
<style lang="scss">
  // 引入初始化样式
  #app {
    height: 100vh;
    overflow: hidden;
    font-weight: 400 !important;
  }
  .el-button {
    font-weight: 400 !important;
  }

  .gva-body-h {
    min-height: calc(100% - 3rem);
  }

  .gva-container {
    height: calc(100% - 2.5rem);
  }
  .gva-container2 {
    height: calc(100% - 4.5rem);
  }
</style>
