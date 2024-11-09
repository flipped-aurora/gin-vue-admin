<!--
    本组件参考 arco-pro 的实现 将 ts 改为 js 写法
    https://github.com/arco-design/arco-design-pro-vue/blob/main/arco-design-pro-vite/src/views/dashboard/workplace/components/content-chart.vue
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
    @desc: 人数统计图表
!-->

<template>
  <Chart :height="height" :option="chartOption" />
</template>

<script setup>
  import Chart from '@/components/charts/index.vue'
  import useChartOption from '@/hooks/charts'
  import { graphic } from 'echarts'
  import { ref } from 'vue'
  import { storeToRefs } from 'pinia'
  import { useAppStore } from '@/pinia'
  const appStore = useAppStore()
  const { config } = storeToRefs(appStore)

  const prop = defineProps({
    height: {
      type: String,
      default: '128px'
    },
    data: {
      type: Array,
      default: () => []
    }
  })
  const graphicFactory = (side) => {
    return {
      type: 'text',
      bottom: '8',
      ...side,
      style: {
        text: '',
        textAlign: 'center',
        fill: '#4E5969',
        fontSize: 12
      }
    }
  }
  const graphicElements = ref([
    graphicFactory({ left: '5%' }),
    graphicFactory({ right: 0 })
  ])
  const { chartOption } = useChartOption(() => {
    return {
      grid: {
        left: '40',
        right: '0',
        top: '10',
        bottom: '30'
      },
      xAxis: {
        type: 'category',
        offset: 2,
        show: false,
        boundaryGap: false,
        axisLine: {
          show: false
        },
        axisTick: {
          show: false
        },
        splitLine: {
          show: false
        }
      },
      yAxis: {
        type: 'value',
        show: false,
        axisLine: {
          show: false
        },
        axisLabel: {
          show: false
        },
        splitLine: {
          show: false
        }
      },
      graphic: {
        elements: graphicElements.value
      },
      series: [
        {
          data: prop.data,
          type: 'line',
          smooth: true,
          symbolSize: 12,
          emphasis: {
            focus: 'series',
            itemStyle: {
              borderWidth: 2
            }
          },
          lineStyle: {
            width: 3,
            color: new graphic.LinearGradient(0, 0, 1, 0, [
              {
                offset: 0,
                color: `${config.value.primaryColor}32`
              },
              {
                offset: 0.5,
                color: `${config.value.primaryColor}64`
              },
              {
                offset: 1,
                color: `${config.value.primaryColor}FF`
              }
            ])
          },
          showSymbol: false,
          areaStyle: {
            opacity: 0.8,
            color: new graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: `${config.value.primaryColor}20`
              },
              {
                offset: 1,
                color: `${config.value.primaryColor}08`
              }
            ])
          }
        }
      ]
    }
  })
</script>

<style scoped lang="scss"></style>
