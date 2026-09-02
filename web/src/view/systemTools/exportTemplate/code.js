export const getCode = (templateID) => {
  return `<template>
  <!-- 导出组件 -->
  <ExportExcel templateId="${templateID}" :condition="condition" :limit="limit" :offset="offset" :order="order" />

  <!-- 导入组件 handleSuccess为导入成功后的回调函数 -->
  <ImportExcel templateId="${templateID}" @on-success="handleSuccess" />

  <!-- 导出模板 -->
  <ExportTemplate templateId="${templateID}" />
</template>

<script setup>
import { ref } from 'vue';
// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue';
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue';
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue';

const condition = ref({}); // 查询条件
const limit = ref(10); // 最大条数限制
const offset = ref(0); // 偏移量
const order = ref('id desc'); // 排序条件

const handleSuccess = (res) => {
  console.log(res);
  // 导入成功的回调函数
};
</script>`
}
