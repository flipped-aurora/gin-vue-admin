<template>
  <div>
    <el-button
      size="small"
      style="float:right;margin-top:6px;margin-right:6px;"
      @click="saveXML"
    >导出XML</el-button>
    <el-button
      size="small"
      style="float:right;margin-top:6px;margin-right:6px;"
      @click="saveImg"
    >导出图片</el-button>
    <el-popover placement="bottom" width="160" v-model="visible">
      <p>确认流程无误并保存吗</p>
      <div style="text-align: right; margin: 0">
        <el-button size="mini" type="text" @click="visible = false">取消</el-button>
        <el-button type="primary" size="mini" @click="save">确定</el-button>
      </div>
      <el-button size="small" slot="reference" style="float:right;margin-top:6px;margin-right:6px;">保存流程</el-button>
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
      :processModel="processModel"
    />
  </div>
</template>
<script>
import {
    findWorkflowProcess,
    createWorkflowProcess
} from "@/api/workflowProcess";  //  此处请自行替换地址
import gvaWfd from "@/components/gva-wfd";
import { getUserList } from "@/api/user";
import { getAuthorityList } from "@/api/authority";
import edge from '../../../components/gva-wfd/shape/edge';
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
      demoData: {},
      processModel:{
          id: '',
          name: '',
          category: '',
          clazz: 'process',
          dataObjs: [],
          signalDefs: [],
          messageDefs: [],
        },
      users: [],
      authorities: [],
      groups: [
        { id: "1", name: "组1" },
        { id: "2", name: "组2" },
        { id: "3", name: "组3" }
      ],
      categorys: [
        { id: "1", name: "分类1" },
        { id: "2", name: "分类2" },
        { id: "3", name: "分类3" },
        { id: "4", 分组: "分组4" }
      ]
    };
  },
  methods: {
    async save() {
      this.visible = false;
      const obj = this.$refs["wfd"].graph.save()
      const processModel = this.processModel
      const nodeMap={}

      obj.edges.map(item=>{
        if(nodeMap[item.source]){
          nodeMap[item.source].push(item)
        }else{
          nodeMap[item.source] = [item]
        }
      })
      obj.nodes.map(item=>{
        item.edges = nodeMap[item.id]
      })
      processModel.nodes = obj.nodes
      if(!processModel.id){
        this.$message({
         type:"error",
         message:"流程ID为必填项(点击空白处录入流程基本信息)"
       })
       return
      }
     const res = await createWorkflowProcess(processModel)
     if(res.code == 0){
       this.$message({
         type:"success",
         message:"创建成功"
       })
     }
    },
    saveXML() {
      console.log(this.$refs["wfd"].graph.saveXML());
    },
    saveImg() {
      console.log(this.$refs["wfd"].graph.saveImg());
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
      authorityRes.data.list.map(item => {
        this.authorities.push({
          id: item.authorityId,
          name: item.authorityName
        });
      });
    }
    if(this.$route.query.id){
      const res = await findWorkflowProcess({ id: this.$route.query.id });
      this.disabled = this.$route.query.type == "view"
      if(res.code == 0){
        const nodes = []
        const edges = []
        res.data.reworkflowProcess.nodes.map(item=>{
          edges.push(...item.edges)
          delete item.edges
          nodes.push(item)
        })
        delete res.data.reworkflowProcess.nodes
        this.demoData = {edges,nodes}
        this.processModel = res.data.reworkflowProcess
      }
    }
        this.done = true
  }
};
</script>