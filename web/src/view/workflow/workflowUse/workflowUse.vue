<template>
    <div class="workflow-use">

    </div>
</template>
<script>
import {findWorkflowStep} from "@/api/workflowProcess.js"
export default {
    name:"WorklowUse",
    data(){
        return{
            
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
                this.node = res.data.workflow.node[0]
            }
        }
    }
}
</script>