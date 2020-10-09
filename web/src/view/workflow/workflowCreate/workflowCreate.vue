<template>
  <div>
    <el-button size="small" style="float:right;margin-top:6px;margin-right:6px;" @click="saveXML">导出XML</el-button>
    <el-button size="small" style="float:right;margin-top:6px;margin-right:6px;" @click="saveImg">导出图片</el-button>
    <el-button size="small" style="float:right;margin-top:6px;margin-right:6px;" @click="save">保存流程</el-button>
    <wfd-gva ref="wfd" :data="demoData" :height="600" :users="users" :groups="groups" :categorys="categorys" :lang="lang" />
  </div>
</template>
<script>


import wfdGva from 'wfd-gva'
export default {
  name: 'Workflow',
  components:{
    wfdGva
  },
  data () {
    return {
      lang: "zh",
      demoData: {},
      users: [{id:'1',name:'审批人1'},{id:'2',name:'审批人2'},{id:'3',name:'审批人3'}],
      groups: [{id:'1',name:'组1'},{id:'2',name:'组2'},{id:'3',name:'组3'}],
      categorys:[{id:'1',name:'分类1'},{id:'2',name:'分类2'},{id:'3',name:'分类3'}]
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
  }
}
</script>