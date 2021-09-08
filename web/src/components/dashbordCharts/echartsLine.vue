<template>
  <div class="dashbord-line-box">
    <div class="dashbord-line-title">Gva 仓库commit 记录</div>
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
for (var i = 1; i < 21; i++) {
  dataAxis.push(`${i}号`)
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
  210,
  122,
  133,
  334,
  198,
  123,
  125,
  220,
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
  beforeDestroy() {
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
            itemStyle: {
              borderRadius: [5, 5, 0, 0],
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                { offset: 0, color: '#83bff6' },
                { offset: 0.5, color: '#188df0' },
                { offset: 1, color: '#188df0' },
              ]),
            },
            emphasis: {
              itemStyle: {
                color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                  { offset: 0, color: '#2378f7' },
                  { offset: 0.7, color: '#2378f7' },
                  { offset: 1, color: '#83bff6' },
                ]),
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
  padding: 20px;
  .dashbord-line {
    background-color: #fff;
    height: 360px;
    width: calc(100% - 40px);
  }
  .dashbord-line-title {
    color: rgb(56, 137, 206);
    font-size: 18px;
  }
}
</style>
