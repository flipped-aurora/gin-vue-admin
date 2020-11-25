<template>
 <div class="big">
         <div class="card">
             <div class="car-left">
                 <div>
                    <span class="card-img"> <img :src="userInfo.headerImg"  alt="" > </span>   
                    <div class="text"><h4>早安，管理员， 请开始您一天的工作吧！</h4>
                        <p class="tips-text">
                            <i class="el-icon-sunny el-icon"></i>
                            <span>今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</span>
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
                      <el-card shadow="hover" class="grid-content">
                        <i :class="card.icon" :style="{color:card.color}"></i>
                        <p>{{card.label}}</p>  
                    </el-card>  
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
                icon:"el-icon el-icon-monitor",
                name:"user",
                color:"#ff9c6e"
            },
            {
                label:"角色管理",
                icon:"el-icon el-icon-setting",
                name:"authority",
                color:"#69c0ff"
            },
            {
                label:"菜单管理",
                icon:"el-icon el-icon-menu",
                name:"menu",
                color:"#b37feb"
            },
            {
                label:"代码生成器",
                icon:"el-icon el-icon-cpu",
                name:"autoCode",
                color:"#ffd666"
            },
            {
                label:"表单生成器",
                icon:"el-icon el-icon-document-checked",
                name:"formCreate",
                color:"#ff85c0"
            },
            {
                label:"关于我们",
                icon:"el-icon el-icon-user",
                name:"about",
                color:"#5cdbd3"
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
