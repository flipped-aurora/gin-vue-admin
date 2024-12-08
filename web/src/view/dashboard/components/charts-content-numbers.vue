<!--
    本组件参考 arco-pro 的实现 将 ts 改为 js 写法
    https://github.com/arco-design/arco-design-pro-vue/blob/main/arco-design-pro-vite/src/views/dashboard/workplace/components/content-chart.vue
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <Chart :height="height" :option="chartOption" />
</template>

<script setup>
  import Chart from '@/components/charts/index.vue'
  import useChartOption from '@/hooks/charts'
  import { graphic } from 'echarts'
  import { computed, ref } from 'vue'
  import { useAppStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  const appStore = useAppStore()
  const { config } = storeToRefs(appStore)
  defineProps({
    height: {
      type: String,
      default: '128px'
    }
  })
  const dotColor = computed(() => {
    return appStore.isDark ? '#333' : '#E5E8EF'
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
  const xAxis = ref([
    '2024-1',
    '2024-2',
    '2024-3',
    '2024-4',
    '2024-5',
    '2024-6',
    '2024-7',
    '2024-8'
  ])
  const chartsData = ref([12, 22, 32, 45, 32, 78, 89, 92])
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
        data: xAxis.value,
        boundaryGap: false,
        axisLabel: {
          color: '#4E5969',
          formatter(value, idx) {
            if (idx === 0) return ''
            if (idx === xAxis.value.length - 1) return ''
            return `${value}`
          }
        },
        axisLine: {
          show: false
        },
        axisTick: {
          show: false
        },
        splitLine: {
          show: true,
          interval: (idx) => {
            if (idx === 0) return false
            if (idx === xAxis.value.length - 1) return false
            return true
          },
          lineStyle: {
            color: dotColor.value
          }
        },
        axisPointer: {
          show: true,
          lineStyle: {
            color: `${config.value.primaryColor}FF`,
            width: 2
          }
        }
      },
      yAxis: {
        type: 'value',
        axisLine: {
          show: false
        },
        axisLabel: {
          formatter(value, idx) {
            if (idx === 0) return value
            return `${value}k`
          }
        },
        splitLine: {
          show: true,
          lineStyle: {
            type: 'dashed',
            color: dotColor.value
          }
        }
      },
      tooltip: {
        trigger: 'axis',
        formatter(params) {
          const [firstElement] = params
          return `<div>
            <p class="tooltip-title">${firstElement.axisValueLabel}</p>
            <div class="content-panel"><span>总内容量</span><span class="tooltip-value">${(
              Number(firstElement.value) * 10000
            ).toLocaleString()}</span></div>
          </div>`
        },
        className: 'echarts-tooltip-diy'
      },
      graphic: {
        elements: graphicElements.value
      },
      series: [
        {
          data: chartsData.value,
          type: 'line',
          smooth: true,
          // symbol: 'circle',
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
                color: `${config.value.primaryColor}80`
              },
              {
                offset: 0.5,
                color: `${config.value.primaryColor}92`
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
