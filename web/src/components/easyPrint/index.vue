<template>
    <div>
      <el-button @click="onprint" >打印</el-button>
      <div v-html="previewDiv ? previewDiv.html() : ''" />
    </div>
</template>
<script lang='ts' setup>
import { ref, watch } from 'vue'
import {ElMessage as message} from 'element-plus'
import ticketPanel from './panel';
import converResDataToPrintData from './adapter';
import * as print from 'vue-plugin-hiprint'
console.log(print, 'ppp')
const { hiprint } = print

hiprint.init();
const props = defineProps<{
  hall:any
  info: any
  printList: any
}>()

const emits = defineEmits(['handleAfterPrint'])
const hiprintTemplate = new hiprint.PrintTemplate({ template: ticketPanel });

  // printData 打印接受的数据 
const printData = ref([])
const previewDiv = ref(null)
const setPrintData = (val)=>printData.value = val
const setPreviewDiv = (val)=>previewDiv.value = val
const handleAfterPrint=(list)=>{
  emits('handleAfterPrint', list)
}
const onprint=()=> {
  if (!props.hall || !props.info || !props.printList?.length) {
    message.info('好像还没有选择数据');
    return;
  }
  hiprintTemplate.printByHtml(previewDiv.value);
  console.log('props.printList',props.printList)
  handleAfterPrint(props.printList);
}

watch(()=>printData.value,()=>{
  const previewDom = hiprintTemplate.getSimpleHtml(printData.value);
  setPreviewDiv(previewDom);
})
watch(()=>props.printList,()=>{
  setPrintData(converResDataToPrintData(props.printList,props.info));
})
</script>
<style lang='scss' scoped></style>