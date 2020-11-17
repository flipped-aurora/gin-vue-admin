<template>
    <div class="workflow-use">

        <WorkflowInfo v-if="done" :wf="this.node"/>
    </div>
</template>
<script>
import {findWorkflowStep} from "@/api/workflowProcess.js"
export default {
    name:"WorklowUse",
    data(){
        return{
            done:false
        }
    },
    async created(){
        const workflowId = this.$route.query.workflowId
        const actionId = this.$route.query.actionId
        const businessId = this.$route.query.businessId
        if(workflowId){
            const res = await findWorkflowStep({id:workflowId})
            if(res.code == 0){
                this.workflow = res.data.workflow
                this.node = res.data.workflow.nodes[0]
                this.done = true
            }
        }
    },
    beforeCreate(){
        this.$options.components.WorkflowInfo = ()=>import("@/"+this.node.view)
    }
}
</script>