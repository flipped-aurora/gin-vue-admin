<template>
  <div class="workflow-use">
      <div style="padding:10px 20px">
    <el-steps :active="moves.length-1" :process-status="processStatus" finish-status="finish" align-center>
      <el-step v-for="(item, key) in moves" :key="key">
        <div slot="title">{{ item.workflowNode.label }}</div>
        <div slot="description">
          <div>节点说明:{{ item.workflowNode.description }}</div>
          <div v-if="item.operator.nickName">操作人:{{ item.operator.nickName }}</div>
          <div>操作参数:{{ item.param||'无参数' }}</div>
        </div>
      </el-step>
    </el-steps>
      </div>
    <WorkflowInfo
      v-if="done"
      :wf="node"
      :business="business"
      :move="move"
      :workflowMoveID="$route.query.workflowMoveID"
    />
  </div>
</template>
<script>
import {findWorkflowStep,getWorkflowMoveByID} from "@/api/workflowProcess.js"
export default {
    name:"WorklowUse",
    data(){
        return{
            done:false,
            business:null,
            node:null,
            moves:[],
            move:null
        }
    },
    computed:{
        processStatus(){
            const node = this.moves[this.moves.length-1]
            if(node&&node.workflowNode.clazz == "end"){
                if(node.workflowNode.success){
                    return "success"
                }else{
                    return "error"
                }
            }else{
                return "process"
            }
        }
    },
    methods:{
        createDone(){
             let path = ""
            if(this.node.view){
                path = this.node.view
            }else{
                path = this.workflow.view
            }
            this.$options.components.WorkflowInfo = ()=>import("@/"+path)
               this.done = true
        }
    },
    async created(){
        const workflowId = this.$route.query.workflowId
        const workflowMoveID = this.$route.query.workflowMoveID
        if(workflowId){
            const res = await findWorkflowStep({id:workflowId})
            if(res.code == 0){
                this.workflow = res.data.workflow
                this.node = res.data.workflow.nodes[0]
                this.createDone()
            }
        }else if(workflowMoveID){
            const res = await getWorkflowMoveByID({id:workflowMoveID})
             if(res.code == 0){
               this.business =  res.data.business
               this.workflow = res.data.move.workflowProcess
               this.node = res.data.move.workflowNode
               this.move = res.data.move
               this.moves = res.data.moves
               this.createDone()
            }
        }
    },
    beforeCreate(){
       
    }
}
</script>