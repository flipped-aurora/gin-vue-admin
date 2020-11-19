<template>
    <div class="workflow-use">
        <WorkflowInfo v-if="done" :wf="this.node" :business="business"/>
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
            node:null
        }
    },
    async created(){
        const workflowId = this.$route.query.workflowId
        const wfmId = this.$route.query.wfmId
        if(workflowId){
            const res = await findWorkflowStep({id:workflowId})
            if(res.code == 0){
                this.workflow = res.data.workflow
                this.node = res.data.workflow.nodes[0]
                this.done = true
            }
        }else if(wfmId){
            const res = await getWorkflowMoveByID({id:wfmId})
             if(res.code == 0){
               this.business =  res.data.business
               this.node = res.data.move.workflowNode
               this.done = true
            }
        }
    },
    beforeCreate(){
        this.$options.components.WorkflowInfo = ()=>import("@/"+this.node.view)
    }
}
</script>