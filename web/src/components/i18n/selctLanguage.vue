<template>
  <el-dropdown
      trigger="click"
      @command="handleSetLanguage"
  >
    <slot name="info"></slot>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
            :disabled="$i18n.locale==='en'"
            command="en"
        >
          <img
              alt="English" class="w-8 mr-1"
              src="@/assets/flags/en.svg"
          >English
        </el-dropdown-item>
        <el-dropdown-item
            :disabled="$i18n.locale==='zh'"
            command="zh"
        ><img
            alt="中文" class="w-8 mr-1"
            src="@/assets/flags/zh.svg"
        >中文
        </el-dropdown-item>
        <el-dropdown-item
            :disabled="$i18n.locale==='ar'"
            command="ar"
        ><img
            alt="العربية" class="w-8 mr-1"
            src="@/assets/flags/ar.svg"
        >العربية
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup>
import Cookies from 'js-cookie' // added by mohamed hassan to support multilanguage
import {useI18n} from 'vue-i18n'
import {ElMessage} from "element-plus"; // added by mohamed hassan to support multilanguage
const i18n = useI18n() // added by mohamed hassan to support multilanguage
const {t} = useI18n() // added by mohamed hassan to support multilanguage
import {useUserStore} from '@/pinia/modules/user'

import {emitter} from "@/utils/bus";

const userStore = useUserStore()

const toggleRefresh = () =>{
  emitter.emit('reload')
  setTimeout(() => {
  }, 300);
}
const handleSetLanguage = (lang) => {
  toggleRefresh()
  i18n.locale.value = lang
  userStore.setLanguage(lang)
  Cookies.set('language', lang)

  ElMessage({
    message: t('general.langSwitch'),
    type: 'success'
  })


}

</script>


<style lang="scss" scoped>

</style>
