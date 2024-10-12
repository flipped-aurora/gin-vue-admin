<script setup>
import {defineProps, onMounted} from "vue";


const props = defineProps({
  title:{
    type: String,
  },
  tableData:{
    type: Object,
  },
});


onMounted(() => {
  console.log("props.tableData",props.tableData)
})

</script>

<template>
  <div>

    <el-table :data="props.tableData.data" style="width: 100%">
      <el-table-column v-for="v in props.tableData.columns" :prop="v.prop" :label="v.label" :width="v.width" />
    </el-table>

  </div>

</template>

<style scoped lang="scss">

</style>