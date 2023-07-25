<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title">
      访问趋势
    </div>
    <div id="echarts" class="dashboard-line" />
  </div>
</template>
<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import { useWindowSize } from '@/utils/useWindowResize'

var dataAxis = []
for (var i = 1; i < 13; i++) {
  dataAxis.push(`${i}月`)
}
var data = [
  220,
  182,
  191,
  234,
  290,
  330,
  310,
  123,
  442,
  321,
  90,
  149,
]
var yMax = 500
var dataShadow = []

// eslint-disable-next-line no-redeclare
for (var i = 0; i < data.length; i++) {
  dataShadow.push(yMax)
}

let chart = null

const initChart = () => {
  nextTick(() => {
    if (chart) {
      chart?.dispose()
    }
    chart = null
    chart = echarts.init(document.querySelector('#echarts'))
    setOptions()
  })
}
const setOptions = () => {
  chart.setOption({
    grid: {
      left: '40',
      right: '20',
      top: '40',
      bottom: '20',
    },
    xAxis: {
      data: dataAxis,
      axisTick: {
        show: false,
      },
      axisLine: {
        show: false,
      },
      z: 10,
    },
    yAxis: {
      axisLine: {
        show: false,
      },
      axisTick: {
        show: false,
      },
      axisLabel: {
        textStyle: {
          color: '#999',
        },
      },
    },
    dataZoom: [
      {
        type: 'inside',
      },
    ],
    series: [
      {
        type: 'bar',
        barWidth: '40%',
        itemStyle: {
          borderRadius: [5, 5, 0, 0],
          color: '#188df0',
        },
        emphasis: {
          itemStyle: {
            color: '#188df0',
          },
        },
        data: data,
      },
    ],
  })
}
useWindowSize(() => {
  window.requestAnimationFrame(() => {
    initChart()
  })
})

onUnmounted(() => {
  if (!chart) {
    return
  }
  chart.dispose()
  chart = null
})
</script>
<style lang="scss" scoped>
.dashboard-line-box {
  .dashboard-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }

  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}
</style>
