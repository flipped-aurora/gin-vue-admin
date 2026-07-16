<template>
  <Chart :height="height" :option="chartOption" />
</template>

<script setup>
  import Chart from '@/components/charts/index.vue'
  import useChartOption from '@/hooks/charts'
  import { graphic } from 'echarts'
  import { ref } from 'vue'
  import { storeToRefs } from 'pinia'
  import { useThemeStore } from '@/pinia'
  import { addOpacityToColor } from '@/theme/color'
  const themeStore = useThemeStore()
  const { settings } = storeToRefs(themeStore)

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
        fill: themeStore.isDark ? '#FFFFFF' : '#000000',
        fontSize: 12
      }
    }
  }
  const graphicElements = ref([
    graphicFactory({ left: '5%' }),
    graphicFactory({ right: 0 })
  ])
  const primaryColor = (opacity) => addOpacityToColor(settings.value.themeColor, opacity)
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
                color: primaryColor(0.2)
              },
              {
                offset: 0.5,
                color: primaryColor(0.39)
              },
              {
                offset: 1,
                color: primaryColor(1)
              }
            ])
          },
          showSymbol: false,
          areaStyle: {
            opacity: 0.8,
            color: new graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: primaryColor(0.13)
              },
              {
                offset: 1,
                color: primaryColor(0.03)
              }
            ])
          }
        }
      ]
    }
  })
</script>

<style scoped lang="scss"></style>
