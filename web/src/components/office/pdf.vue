<template>
  <vue-office-pdf
    :src="pdf"
    @rendered="renderedHandler"
    @error="errorHandler"
  />
</template>
<script>
  export default {
    name: 'Pdf'
  }
</script>
<script setup>
  import { ref, watch } from 'vue'

  //引入VueOfficeDocx组件
  import VueOfficePdf from '@vue-office/pdf'
  //引入相关样式
  import '@vue-office/docx/lib/index.css'
  console.log('pdf===>')
  const props = defineProps({
    modelValue: {
      type: String,
      default: () => ''
    }
  })
  const pdf = ref(null)
  watch(
    () => props.modelValue,
    (val) => (pdf.value = val),
    { immediate: true }
  )
  const renderedHandler = () => {
    console.log('pdf 加载成功')
  }
  const errorHandler = () => {
    console.log('pdf 错误')
  }
</script>
