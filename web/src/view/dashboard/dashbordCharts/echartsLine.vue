<template>
  <div class="dashbord-line-box">
    <div class="dashbord-line-title">
      访问趋势
    </div>
    <div
      ref="echart"
      class="dashbord-line"
    />
  </div>
</template>
<script>
import echarts from 'echarts'
import 'echarts/theme/macarons'

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

export default {
  name: 'Line',
  data() {
    return {
      chart: null,
    }
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
    })
  },
  beforeUnmount() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$refs.echart, 'macarons')
      this.setOptions()
    },
    setOptions() {
      this.chart.setOption({
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
    },
  },
}
</script>
<style lang="scss" scoped>
.dashbord-line-box {
  .dashbord-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }
  .dashbord-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}
</style>
