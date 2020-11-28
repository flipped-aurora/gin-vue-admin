<template>
  <div>
    <p style="text-align:center;padding-bottom:4px">当前版本为0.01试用版本，可满足最简单的流程功能，未完成功能，如脚本节点，邮件节点等目前暂时注释入口，期待大家PR</p>
    <el-button
      size="small"
      style="float: right; margin-top: 6px; margin-right: 6px"
      @click="saveXML"
      >导出XML</el-button
    >
    <el-button
      size="small"
      style="float: right; margin-top: 6px; margin-right: 6px"
      @click="saveImg"
      >导出图片</el-button
    >
    <el-popover placement="bottom" width="160" v-model="visible">
      <p>确认流程无误并保存吗</p>
      <div style="text-align: right; margin: 0">
        <el-button size="mini" type="text" @click="visible = false"
          >取消</el-button
        >
        <el-button
          type="primary"
          size="mini"
          @click="save"
          :disabled="$route.query.type == 'view'"
          >确定</el-button
        >
      </div>
      <el-button
        size="small"
        slot="reference"
        style="float: right; margin-top: 6px; margin-right: 6px"
        >保存流程</el-button
      >
    </el-popover>

    <gva-wfd
      ref="wfd"
      :data="demoData"
      v-if="done"
      :height="600"
      :users="users"
      :authorities="authorities"
      :groups="groups"
      :categorys="categorys"
      :lang="lang"
      :propProcessModel="processModel"
    />
  </div>
</template>
<script>
import {
    findWorkflowProcess,
    createWorkflowProcess,
    updateWorkflowProcess
} from "@/api/workflowProcess";  //  此处请自行替换地址
import gvaWfd from "@/components/gva-wfd";
import { getUserList } from "@/api/user";
import { getAuthorityList } from "@/api/authority";
export default {
  name: "Workflow",
  components: {
    gvaWfd
  },
  data() {
    return {
      visible: false,
      lang: "zh",
      done:false,
      demoData: {
        nodes: [],
        edges: []
      },
      wkType:"create",
      users: [],
      authorities: [],
      processModel:{
          id: '',
          name: '',
          category: '',
          clazz: 'process',
          dataObjs: [],
          signalDefs: [],
          messageDefs: [],
        },
      groups: [],
      categorys: []
    };
  },
  methods: {
    async save() {
      this.visible = false;
      const obj = this.$refs["wfd"].graph.save()
      const processModel = this.$refs["wfd"].processModel
      processModel.edges = obj.edges
      processModel.nodes = JSON.parse(JSON.stringify(obj.nodes))
      processModel.nodes.map(item=>{
        if(item.assignValue){
          item.assignValue = ","+String(item.assignValue)+","
        }
      })
      if(!processModel.id){
        this.$message({
          type:"error",
          message:"流程id必填，点击空白处填写流程基本信息"
        })
        return
      }
      if(this.$route.query.type == 'edit'){
         const res = await updateWorkflowProcess(processModel)
         if(res.code == 0){
          this.$message({
            type:"success",
            message:"编辑成功"
          })
        }
      }else{
        const res = await createWorkflowProcess(processModel)
        if(res.code == 0){
          this.$message({
            type:"success",
            message:"创建成功"
          })
        }
      }
    },
    saveXML() {
      console.log(this.$refs["wfd"].graph.saveXML());
    },
    saveImg() {
      console.log(this.$refs["wfd"].graph.saveImg());
    },
    fmtAuthority(authorityList,list){
      authorityList.map(item => {
        list.push({
          id: item.authorityId,
          name: item.authorityName
        });
        if(item.children){
          this.fmtAuthority(item.children,list)
        }
      });
    }
  },
  async created() {
    const userRes = await getUserList({ page: 1, pageSize: 9999999 });
    if (userRes.code == 0) {
      userRes.data.list.map(item => {
        this.users.push({ id: item.ID, name: item.nickName });
      });
    }
    const authorityRes = await getAuthorityList({ page: 1, pageSize: 9999999 });
    if (authorityRes.code == 0) {
      this.fmtAuthority(authorityRes.data.list,this.authorities)
    }
    if(this.$route.query.id){
      const res = await findWorkflowProcess({ id: this.$route.query.id });
      this.disabled = this.$route.query.type == "view"
      if(res.code == 0){
         res.data.reworkflowProcess.nodes.map(item=>{
           if(item.assignValue){
             const watiUseArr =item.assignValue.substr(1,item.assignValue.length-2).split(",")
             if(item.assignType == 'user'){
               item.assignValue = []
               watiUseArr.map(i => {
                 item.assignValue.push(Number(i))
               })
             }else{
                item.assignValue = watiUseArr
             }
           }
         })
        this.demoData.nodes = res.data.reworkflowProcess.nodes
        delete res.data.reworkflowProcess.nodes
        this.demoData.edges = res.data.reworkflowProcess.edges
        delete res.data.reworkflowProcess.edges
        this.processModel = res.data.reworkflowProcess
      }
      this.wkType = this.$route.query.type
    }
        this.done = true

  }
};
</script>