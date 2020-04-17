<template>
    <div :class="className" :style="{height:height,width:width}" />
</template>

<script>
    import echarts from 'echarts'
    require('echarts/theme/macarons') // echarts theme

    const animationDuration = 3000

    export default {
        name: "Sunburst",
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
                chart: null,
                datas:[
                    {
                    name: 'Flora',
                    itemStyle: {
                        color: '#da0d68'
                    },
                    children: [{
                        name: 'Black Tea',
                        value: 1,
                        itemStyle: {
                            color: '#975e6d'
                        }
                    }, {
                        name: 'Floral',
                        itemStyle: {
                            color: '#e0719c'
                        },
                        children: [{
                            name: 'Chamomile',
                            value: 1,
                            itemStyle: {
                                color: '#f99e1c'
                            }
                        }, {
                            name: 'Rose',
                            value: 1,
                            itemStyle: {
                                color: '#ef5a78'
                            }
                        }, {
                            name: 'Jasmine',
                            value: 1,
                            itemStyle: {
                                color: '#f7f1bd'
                            }
                        }]
                    }]
                },
                    {
                    name: 'Fruity',
                    itemStyle: {
                        color: '#da1d23'
                    },
                    children: [{
                        name: 'Berry',
                        itemStyle: {
                            color: '#dd4c51'
                        },
                        children: [{
                            name: 'Blackberry',
                            value: 1,
                            itemStyle: {
                                color: '#3e0317'
                            }
                        }, {
                            name: 'Raspberry',
                            value: 1,
                            itemStyle: {
                                color: '#e62969'
                            }
                        }, {
                            name: 'Blueberry',
                            value: 1,
                            itemStyle: {
                                color: '#6569b0'
                            }
                        }, {
                            name: 'Strawberry',
                            value: 1,
                            itemStyle: {
                                color: '#ef2d36'
                            }
                        }]
                    }, {
                        name: 'Dried Fruit',
                        itemStyle: {
                            color: '#c94a44'
                        },
                        children: [{
                            name: 'Raisin',
                            value: 1,
                            itemStyle: {
                                color: '#b53b54'
                            }
                        }, {
                            name: 'Prune',
                            value: 1,
                            itemStyle: {
                                color: '#a5446f'
                            }
                        }]
                    }, {
                        name: 'Other Fruit',
                        itemStyle: {
                            color: '#dd4c51'
                        },
                        children: [{
                            name: 'Coconut',
                            value: 1,
                            itemStyle: {
                                color: '#f2684b'
                            }
                        }, {
                            name: 'Cherry',
                            value: 1,
                            itemStyle: {
                                color: '#e73451'
                            }
                        }, {
                            name: 'Pomegranate',
                            value: 1,
                            itemStyle: {
                                color: '#e65656'
                            }
                        }, {
                            name: 'Pineapple',
                            value: 1,
                            itemStyle: {
                                color: '#f89a1c'
                            }
                        }, {
                            name: 'Grape',
                            value: 1,
                            itemStyle: {
                                color: '#aeb92c'
                            }
                        }, {
                            name: 'Apple',
                            value: 1,
                            itemStyle: {
                                color: '#4eb849'
                            }
                        }, {
                            name: 'Peach',
                            value: 1,
                            itemStyle: {
                                color: '#f68a5c'
                            }
                        }, {
                            name: 'Pear',
                            value: 1,
                            itemStyle: {
                                color: '#baa635'
                            }
                        }]
                    }, {
                        name: 'Citrus Fruit',
                        itemStyle: {
                            color: '#f7a128'
                        },
                        children: [{
                            name: 'Grapefruit',
                            value: 1,
                            itemStyle: {
                                color: '#f26355'
                            }
                        }, {
                            name: 'Orange',
                            value: 1,
                            itemStyle: {
                                color: '#e2631e'
                            }
                        }, {
                            name: 'Lemon',
                            value: 1,
                            itemStyle: {
                                color: '#fde404'
                            }
                        }, {
                            name: 'Lime',
                            value: 1,
                            itemStyle: {
                                color: '#7eb138'
                            }
                        }]
                    }]
                },
                    {
                    name: 'Other',
                    itemStyle: {
                        color: '#0aa3b5'
                    },
                    children: [{
                        name: 'Papery/Musty',
                        itemStyle: {
                            color: '#9db2b7'
                        },
                        children: [{
                            name: 'Stale',
                            value: 1,
                            itemStyle: {
                                color: '#8b8c90'
                            }
                        }, {
                            name: 'Cardboard',
                            value: 1,
                            itemStyle: {
                                color: '#beb276'
                            }
                        }, {
                            name: 'Papery',
                            value: 1,
                            itemStyle: {
                                color: '#fefef4'
                            }
                        }, {
                            name: 'Woody',
                            value: 1,
                            itemStyle: {
                                color: '#744e03'
                            }
                        }, {
                            name: 'Moldy/Damp',
                            value: 1,
                            itemStyle: {
                                color: '#a3a36f'
                            }
                        }, {
                            name: 'Musty/Dusty',
                            value: 1,
                            itemStyle: {
                                color: '#c9b583'
                            }
                        }, {
                            name: 'Musty/Earthy',
                            value: 1,
                            itemStyle: {
                                color: '#978847'
                            }
                        }, {
                            name: 'Animalic',
                            value: 1,
                            itemStyle: {
                                color: '#9d977f'
                            }
                        }, {
                            name: 'Meaty Brothy',
                            value: 1,
                            itemStyle: {
                                color: '#cc7b6a'
                            }
                        }, {
                            name: 'Phenolic',
                            value: 1,
                            itemStyle: {
                                color: '#db646a'
                            }
                        }]
                    }, {
                        name: 'Chemical',
                        itemStyle: {
                            color: '#76c0cb'
                        },
                        children: [{
                            name: 'Bitter',
                            value: 1,
                            itemStyle: {
                                color: '#80a89d'
                            }
                        }, {
                            name: 'Salty',
                            value: 1,
                            itemStyle: {
                                color: '#def2fd'
                            }
                        }, {
                            name: 'Medicinal',
                            value: 1,
                            itemStyle: {
                                color: '#7a9bae'
                            }
                        }, {
                            name: 'Petroleum',
                            value: 1,
                            itemStyle: {
                                color: '#039fb8'
                            }
                        }, {
                            name: 'Skunky',
                            value: 1,
                            itemStyle: {
                                color: '#5e777b'
                            }
                        }, {
                            name: 'Rubber',
                            value: 1,
                            itemStyle: {
                                color: '#120c0c'
                            }
                        }]
                    }]
                },
                ]
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
                this.chart = echarts.init(this.$el, 'macarons')

                this.chart.setOption({

                        series: {
                            type: 'sunburst',
                            highlightPolicy: 'ancestor',
                            data: this.datas,
                            radius: [0, '95%'],
                            sort: null,
                            levels: [{}, {
                                r0: '15%',
                                r: '35%',
                                itemStyle: {
                                    borderWidth: 2
                                },
                                label: {
                                    rotate: 'tangential'
                                }
                            }, {
                                r0: '35%',
                                r: '70%',
                                label: {
                                    align: 'right'
                                }
                            }, {
                                r0: '70%',
                                r: '72%',
                                label: {
                                    position: 'outside',
                                    padding: 3,
                                    silent: false
                                },
                                itemStyle: {
                                    borderWidth: 3
                                }
                            }]
                        },
                animationDuration: animationDuration
                })
            }
        }
    }
</script>

<style scoped>

</style>
