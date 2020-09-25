<template>
    <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
    import echarts from 'echarts'
    require('echarts/theme/macarons') // echarts theme

    const animationDuration = 3000

    export default {
        name: "stackMap",
        props: {
            className: {
                type: String,
                default: 'chart'
            },
            width: {
                type: String,
                default: '100%'
            },
            height: {
                type: String,
                default: '300px'
            }
        },
        data() {
            return {
                chart: null
            }
        },
        mounted() {
            this.initChart()
            /* this.__resizeHandler = debounce(() => {
                 if (this.chart) {
                     this.chart.resize()
                 }
             }, 100)
             window.addEventListener('resize', this.__resizeHandler)*/
        },
        beforeDestroy() {
            if (!this.chart) {
                return
            }
            // window.removeEventListener('resize', this.__resizeHandler)
            this.chart.dispose()
            this.chart = null
        },
        methods: {
            initChart() {
                this.chart = echarts.init(this.$el, 'light')

                this.chart.setOption({
                    tooltip: {
                        trigger: 'axis',
                        axisPointer: {            // 坐标轴指示器，坐标轴触发有效
                            type: 'shadow'        // 默认为直线，可选为：'line' | 'shadow'
                        }
                    },
                    legend: {
                        data: ['Javascript', 'Java', 'Python', 'Ruby', 'PHP']
                    },
                    grid: {
                        left: '3%',
                        right: '4%',
                        bottom: '3%',
                        containLabel: true
                    },
                    xAxis: {
                        type: 'value',
                        axisLabel: {
                            show: true,
                            textStyle: {
                                color: 'rgb(192,192,192)',  //更改坐标轴文字颜色
                                fontSize : 12     //更改坐标轴文字大小
                            }
                        },
                        axisTick: {
                            show: false
                        },
                        axisLine:{
                            lineStyle:{
                                color:'rgb(192,192,192)' //更改坐标轴颜色
                            }
                        },
                    },
                    yAxis: {
                        type: 'category',
                        data: ['Mon', 'Tue', 'Wen', 'Thu', 'Fri', 'Sat', 'Sun'],
                        axisLabel: {
                            show: true,
                            textStyle: {
                                color: 'rgb(192,192,192)',  //更改坐标轴文字颜色
                                fontSize: 12     //更改坐标轴文字大小
                            }
                        },
                        axisTick: {
                            show: false
                        },
                        axisLine: {
                            lineStyle: {
                                color: 'rgb(192,192,192)' //更改坐标轴颜色
                            }
                        }
                    },
                    series: [
                        {
                            name: 'Javascript',
                            type: 'bar',
                            stack: '总量',
                            label: {
                                show: true,
                                position: 'insideRight'
                            },
                            data: [320, 302, 301, 334, 390, 330, 320]
                        },
                        {
                            name: 'Java',
                            type: 'bar',
                            stack: '总量',
                            label: {
                                show: true,
                                position: 'insideRight'
                            },
                            data: [120, 132, 101, 134, 90, 230, 210]
                        },
                        {
                            name: 'Python',
                            type: 'bar',
                            stack: '总量',
                            label: {
                                show: true,
                                position: 'insideRight'
                            },
                            data: [220, 182, 191, 234, 290, 330, 310]
                        },
                        {
                            name: 'Ruby',
                            type: 'bar',
                            stack: '总量',
                            label: {
                                show: true,
                                position: 'insideRight'
                            },
                            data: [150, 212, 201, 154, 190, 330, 410]
                        },
                        {
                            name: 'PHP',
                            type: 'bar',
                            stack: '总量',
                            label: {
                                show: true,
                                position: 'insideRight'
                            },
                            data: [820, 832, 901, 934, 1290, 1330, 1320]
                        }
                    ],
                    animationDuration: animationDuration
                })
            }
        }
    }
</script>

<style scoped>

</style>
