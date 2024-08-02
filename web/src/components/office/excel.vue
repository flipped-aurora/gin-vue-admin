<template>
  <VueOfficeExcel :src="excel" @rendered="renderedHandler" @error="errorHandler" style="height: 100vh;width: 100vh"/>
</template>
<script>
export default {
  name: 'Excel'
}
</script>
<script setup>
//引入VueOfficeExcel组件
import VueOfficeExcel from '@vue-office/excel'
//引入相关样式
import '@vue-office/excel/lib/index.css'
import {ref, watch} from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: () => ""
  }
})
const excel = ref('')
watch(() => props.modelValue, val => excel.value = val, {immediate: true})
const renderedHandler = () => {

}
const errorHandler = () => {

}
</script>
<style>

</style>