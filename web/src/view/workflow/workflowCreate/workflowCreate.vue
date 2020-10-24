<template>
  <div>
    <el-button size="small" style="float:right;margin-top:6px;margin-right:6px;" @click="saveXML">导出XML</el-button>
    <el-button size="small" style="float:right;margin-top:6px;margin-right:6px;" @click="saveImg">导出图片</el-button>
    <el-button size="small" style="float:right;margin-top:6px;margin-right:6px;" @click="save">保存流程</el-button>
    <gva-wfd ref="wfd" :data="demoData" :height="600" :users="users" :authorities="authorities" :groups="groups" :categorys="categorys" :lang="lang" />
  </div>
</template>
<script>


import gvaWfd from '@/components/gva-wfd'
import   {getUserList}     from '@/api/user'
import   {getAuthorityList}     from '@/api/authority'
export default {
  name: 'Workflow',
  components:{
    gvaWfd
  },
  data () {
    return {
      lang: "zh",
      demoData: {},
      users: [],
      authorities:[],
      groups: [{id:'1',name:'组1'},{id:'2',name:'组2'},{id:'3',name:'组3'}],
      categorys:[{id:'1',name:'分类1'},{id:'2',name:'分类2'},{id:'3',name:'分类3'},{id:'4',分组:'分组4'}]
    }
  },
  methods:{
    save(){
      console.log(this.$refs['wfd'].graph.save())
    },
    saveXML(){
      console.log(this.$refs['wfd'].graph.saveXML())
    },
    saveImg(){
      console.log(this.$refs['wfd'].graph.saveImg())
    }
  },
  async created(){
   const userRes = await getUserList({page:1,pageSize:9999999})
   if(userRes.code == 0){
     userRes.data.list.map(item=>{
       this.users.push({id:item.ID,name:item.nickName})
     })
   }
   const authorityRes = await getAuthorityList({page:1,pageSize:9999999})
   if(authorityRes.code == 0){
     authorityRes.data.list.map(item=>{
       this.authorities.push({id:item.authorityId,name:item.authorityName})
     })
   }
  },
}
</script>