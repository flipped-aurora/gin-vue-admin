<template>
          <div class="router-history">
            <el-tabs v-model="activeValue" type="card" :closable="!(historys.length==1&&this.$route.name=='dashboard')" @tab-click="changeTab" @tab-remove="removeTab">
              <el-tab-pane
                v-for="item in historys"
                :key="item.name"
                :label="item.meta.title"
                :name="item.name"
              >
              </el-tab-pane>
            </el-tabs>
          </div>
</template>
<script>
export default {
    name:"HistoryComponent",
    data(){
        return{
            historys:[],
            activeValue:"dashboard"
        }
    },
    created(){
        const initHistorys = [
                {
                name:"dashboard",
                meta:{
                    title:"仪表盘"
                }
                }
            ]
        this.historys = JSON.parse(sessionStorage.getItem("historys")) || initHistorys
        this.setTab(this.$route)
    },
    methods:{
        setTab(route){
        if(!this.historys.some(item=>item.name==route.name)){
           const obj = {}
           obj.name = route.name
           obj.meta = route.meta
        this.historys.push(obj)
       }
           this.activeValue = this.$route.name
        },
        changeTab(tab){
            this.$router.push({name:tab.name})
        },
        removeTab(tab){
           const index = this.historys.findIndex(item=>item.name == tab)
           if(this.$route.name == tab){
               if(this.historys.length==1){
                   this.$router.push({name:"dashboard"})
               }else{
                    if(index<this.historys.length-1){
                        this.$router.push({name:this.historys[index+1].name})
                    }else{
                        this.$router.push({name:this.historys[index-1].name})
                    }
               }
           }
               this.historys.splice(index,1)
        }
    },
    watch:{
     $route( to ){
       this.historys = this.historys.filter(item=>!item.meta.hidden)
       this.setTab(to)
       sessionStorage.setItem("historys",JSON.stringify(this.historys))
     }
     
    }
}
</script>
<style lang="scss">
    
</style>