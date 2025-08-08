<template>
  <fc-designer
    ref="designer"
    :config="config"
    :locale="languages[locale]"
    height="calc(100vh - 160px)"
  />
</template>

<script setup>
  import { ref } from 'vue'

  import FcDesigner from '@form-create/designer'

  import en from './locale/en.js' // 导入英文语言包
  import zh from './locale/zh_cn.js' // 导入中文语言包
  import zh_tw from './locale/zh_tw.js' // 导入中文语言包
  import ar from './locale/ar.js'
  import ru from './locale/ru.js'
  import { useI18n } from 'vue-i18n' // 导入中文语言包

  const { locale } = useI18n()

  console.log(locale.value)

  defineOptions({
    name: 'FormGenerator'
  })

  const languages = {
    'zh-cn': zh,
    'zh-TW': zh_tw,
    en: en,
    ar: ar,
    ru: ru
  }

  const designer = ref(null)

  const config = {
    fieldReadonly: false
  }
</script>
