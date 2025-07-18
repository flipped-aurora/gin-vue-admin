<template>
  <error-preview v-if="showError" :error-data="errorInfo" @close="handleClose" @confirm="handleConfirm" />
</template>

<script setup>
import { ref, onUnmounted } from 'vue'
import { emitter } from '@/utils/bus'
import ErrorPreview from '@/components/errorPreview/index.vue'

const showError = ref(false)
const errorInfo = ref(null)
let cb = null

const showErrorDialog = (data) => {
  // 这玩意同时只允许存在一个
  if(showError.value) return

  errorInfo.value = data
  showError.value = true
  cb = data?.fn || null
}

const handleClose = () => {
  showError.value = false
  errorInfo.value = null
  cb = null
}

const handleConfirm = (code) => {
  cb && cb(code)
  handleClose()
}

emitter.on('show-error', showErrorDialog)

onUnmounted(() => {
  emitter.off('show-error', showErrorDialog)
})
</script>
