<template>
 <div class="big">
     <!-- <div class="mid">
         <el-row :gutter="32">
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <raddar-chart />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <stackMap />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="8">
                 <div class="chart-wrapper">
                     <Sunburst/>
                 </div>
             </el-col>
         </el-row>
     </div> -->
     <!-- <div class="top">
         <div id="main" class="chart-container"></div>
     </div> -->
         <div class="card">
             <div class="car-left">
                 <div>
                    <span class="card-img"> <img :src="userInfo.headerImg"  alt="" > </span>   
                    <div class="text"><h4>早安，管理员， 请开始您一天的工作吧！</h4>
                        <p class="tips-text">
                            <i class="el-icon-sunny"></i>
                            <span>今日晴，18℃ - 28℃，出门记得防晒哦。</span>
                        </p>
                    </div>
                 </div>
             </div>
             <div class="car-right">
                 <el-row>
                    <el-col :span="8"><div class="car-item">
                        <span class="flow"><i class="el-icon-s-grid"></i></span>
                        <span>今日流量 </span>
                        <b>13260</b>
                        </div></el-col>
                    <el-col :span="8"><div class="car-item">
                       <span class="user-number">
                         <i class="el-icon-s-custom "></i>
                        </span>
                        <span>总用户 </span>
                        <b>48286</b>
                        </div></el-col>
                    <el-col :span="8"><div class="car-item ">
                      <span class="feedback">
                          <i class="el-icon-star-on"></i>
                      </span> 
                        <span>好评率 </span>
                        <b>98%</b>
                        </div></el-col>
                 </el-row>
             </div>
         </div>
         <div class="shadow">
             <el-row :gutter="20">
                 <el-col :span="4" v-for="(card,key) in toolCards" :key="key" @click.native="toTarget(card.name)">
                      <div class="grid-content">
                        <i :class="card.icon"></i>
                        <p>{{card.label}}</p>  
                    </div>  
                 </el-col>
            </el-row>
         </div>
      <div class="bottom">
         <el-row :gutter="32">
             <el-col :xs="24" :sm="24" :lg="12">
                 <div class="chart-player">
                     <musicPlayer />
                 </div>
             </el-col>
             <el-col :xs="24" :sm="24" :lg="12">
                 <div  class="chart-player">
                     <todo-list />
                 </div>
             </el-col>
         </el-row>
     </div>
 </div>

</template>

<script>
import echarts from 'echarts'
// require('echarts/theme/macarons') // echarts theme
// import RaddarChart from "./component/RaddarChart"
// import stackMap from "./component/stackMap"
// import Sunburst from "./component/Sunburst"
import musicPlayer from "./component/musicPlayer"
import TodoList from "./component/todoList"
import { mapGetters } from 'vuex'
export default {
  name: 'Dashboard',
  data() {
    return {
        toolCards:[
            {
                label:"用户管理",
                icon:"el-icon el-icon-user",
                name:"user"
            },
            {
                label:"角色管理",
                icon:"el-icon el-icon-setting",
                name:"authority"
            },
            {
                label:"菜单管理",
                icon:"el-icon el-icon-menu",
                name:"menu"
            },
            {
                label:"代码生成器",
                icon:"el-icon el-icon-cpu",
                name:"autoCode"
            },
            {
                label:"表单生成器",
                icon:"el-icon el-icon-document-checked",
                name:"formCreate"
            },
            {
                label:"关于我们",
                icon:"el-icon el-icon-user",
                name:"about"
            }

        ]
    }
  },
   computed: {
    ...mapGetters('user', ['userInfo'])
  },
    components:{
        musicPlayer,  //音乐播放器
        TodoList, //TodoList
        // RaddarChart, //雷达图
        // stackMap, //堆叠图
        // Sunburst, //旭日图
    },
    methods:{
      toTarget(name){
          this.$router.push({name})
      }
    },
//   mounted() {
//       let myChart = echarts.init(document.getElementById('main'),'macarons');
//       // let stackMap = echarts.init(document.getElementById('stackMap'));
//       let  option = {
//           legend: {},
//           tooltip: {
//               trigger: 'axis',
//               showContent: false
//           },
//           dataset: {
//               source: [
//                   ['product', '2012', '2013', '2014', '2015', '2016', '2017','2018','2019','2020'],
//                   ['Matcha Latte', 41.1, 30.4, 65.1, 53.3, 83.8, 70.0,6.4, 65.2, 82.5],
//                   ['Milk Tea', 86.5, 92.1, 85.7, 83.1, 73.4, 55.1,2, 67.1, 69.2],
//                   ['Cheese Cocoa', 24.1, 67.2, 79.5, 86.4, 65.2, 82.5,65.1, 53.3, 83.8],
//                   ['Walnut Brownie', 55.2, 67.1, 69.2, 72.4, 53.9, 39.1,86.5, 92.1, 85.7]
//               ]
//           },
//           xAxis: {
//               type: 'category',
//               axisLabel: {
//                   show: true,
//                   textStyle: {
//                       color: 'rgb(192,192,192)',  //更改坐标轴文字颜色
//                       fontSize : 14    //更改坐标轴文字大小
//                   }
//               },
//               axisTick: {
//                   show: false
//               },
//               axisLine:{
//                   lineStyle:{
//                       color:'rgb(192,192,192)' //更改坐标轴颜色
//                   }
//               },
//           },
//           yAxis: {
//               gridIndex:0,
//               axisLabel: {
//                   show: true,
//                   textStyle: {
//                       color: 'rgb(192,192,192)',  //更改坐标轴文字颜色
//                       fontSize: 14    //更改坐标轴文字大小
//                   }
//               },
//               axisTick: {
//                   show: false
//               },
//               axisLine: {
//                   lineStyle: {
//                       color: 'rgb(192,192,192)' //更改坐标轴颜色
//                   }
//               }
//           },
//           grid: {top: '55%'},
//           series: [
//               {type: 'line', smooth: true, seriesLayoutBy: 'row'},
//               {type: 'line', smooth: true, seriesLayoutBy: 'row'},
//               {type: 'line', smooth: true, seriesLayoutBy: 'row'},
//               {type: 'line', smooth: true, seriesLayoutBy: 'row'},
//               {
//                   type: 'pie',
//                   id: 'pie',
//                   radius: '30%',
//                   center: ['50%', '25%'],
//                   label: {
//                       formatter: '{b}: {@2012} ({d}%)'
//                   },
//                   encode: {
//                       itemName: 'product',
//                       value: '2012',
//                       tooltip: '2012'
//                   }
//               }
//           ]
//       };
//       //点记标记点时的动效
//       myChart.on('updateAxisPointer', function (event) {
//           var xAxisInfo = event.axesInfo[0];
//           if (xAxisInfo) {
//               var dimension = xAxisInfo.value + 1;
//               myChart.setOption({
//                   series: {
//                       id: 'pie',
//                       label: {
//                           formatter: '{b}: {@[' + dimension + ']} ({d}%)'
//                       },
//                       encode: {
//                           value: dimension,
//                           tooltip: dimension
//                       }
//                   }
//               });
//           }
//       });

//       window.addEventListener('resize',function() {myChart.resize()});
//       myChart.setOption(option);

//   }
}
</script>

<style lang="scss" scoped>
    .big{
        margin:100px 0 0 0;
        padding-top: 0;
        background-color: rgb(243,243,243);
        padding-top: 15px;
        .top{
            width: 100%;
            height: 360px;
            margin-top: 20px;
            overflow: hidden;
            .chart-container{
                position: relative;
                width: 100%;
                height: 100%;
                padding: 20px;
                background-color: #fff;
            }
        }
        .mid{
            width: 100%;
            height: 380px;
            .chart-wrapper {
                height: 340px;
                background: #fff;
                padding: 16px 16px 0;
                margin-bottom: 32px;
            }
        }
        .bottom{
            width: 100%;
            height: 300px;
            // margin: 20px 0;
            .el-row{
                margin-right: 4px !important;
            }
            .chart-player{
                width: 100%;
                height: 270px;
                padding: 10px;
                background-color: #fff;
            }
        }
    }

</style>
