<template>
          <div class="router-history">
            <el-tabs v-model="activeValue" type="card" :closable="!(historys.length==1&&this.$route.name=='dashbord')" @tab-click="changeTab" @tab-remove="removeTab">
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
            historys:[
                {
                name:"dashbord",
                meta:{
                    title:"仪表盘"
                }
                }
            ],
            activeValue:"dashbord"
        }
    },
    methods:{
        changeTab(tab){
            this.$router.push({name:tab.name})
        },
        removeTab(tab){
           const index = this.historys.findIndex(item=>item.name == tab)
           if(this.$route.name == tab){
               if(this.historys.length==1){
                   this.$router.push({name:"dashbord"})
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
       if(!this.historys.some(item=>item.name==to.name)){
        this.historys.push(to)
       }
       this.activeValue = to.name
     }
    }
}
</script>
<style lang="scss">
    
</style>