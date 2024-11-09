<template>
  <vue-office-docx :src="docx" @rendered="rendered" />
</template>
<script>
  export default {
    name: 'Docx'
  }
</script>
<script setup>
  import { ref, watch } from 'vue'

  // 引入VueOfficeDocx组件
  import VueOfficeDocx from '@vue-office/docx'
  // 引入相关样式
  import '@vue-office/docx/lib/index.css'

  const model = defineModel({
    type: String
  })

  const docx = ref(null)
  watch(
    () => model,
    (value) => {
      docx.value = value
    },
    { immediate: true }
  )
  const rendered = () => {}
</script>
<style lang="scss" scoped></style>
