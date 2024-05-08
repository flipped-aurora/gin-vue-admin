<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <VCharts
    v-if="renderChart"
    :option="options"
    :autoresize="autoResize"
    :style="{ width, height }"
  />
</template>

<script  setup>
import { ref, nextTick } from 'vue';
import VCharts from 'vue-echarts';
import { useWindowResize } from '@/hooks/use-windows-resize'

defineProps({
  options: {
    type: Object,
    default() {
      return {};
    },
  },
  autoResize: {
    type: Boolean,
    default: true,
  },
  width: {
    type: String,
    default: '100%',
  },
  height: {
    type: String,
    default: '100%',
  },
});
const renderChart = ref(false);
nextTick(() => {
  renderChart.value = true;
});
useWindowResize(()=>{
  renderChart.value = false;
  nextTick(() => {
    renderChart.value = true;
  });

})

</script>

<style scoped lang="less"></style>
