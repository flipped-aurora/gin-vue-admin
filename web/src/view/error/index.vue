<template>
  <div>
    <div class="w-full h-screen bg-gray-50 flex items-center justify-center">
      <div class="flex flex-col items-center text-2xl gap-4">
        <img class="w-1/3" src="../../assets/404.png">
        <p class="text-lg">{{ t('error.message1') }}</p>
        <p class="text-lg">{{ t('error.message2') }}</p>
        <p>{{ t('error.projectAddress') }}<a href="https://github.com/flipped-aurora/gin-vue-admin" target="_blank" class="text-blue-600  ">https://github.com/flipped-aurora/gin-vue-admin</a></p>
        <el-button @click="toDashboard">{{ t('error.backToHome') }}</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useUserStore } from '@/pinia/modules/user'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage

defineOptions({
  name: 'Error'
})

const userStore = useUserStore()
const router = useRouter()
const toDashboard = () => {
  router.push({ name: userStore.userInfo.authority.defaultRouter })
}
</script>
