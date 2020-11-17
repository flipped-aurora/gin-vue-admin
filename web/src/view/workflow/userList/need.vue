<template>
    <div>
        
    </div>
</template>

<script>
import {getMyNeed} from "@/api/workflowProcess"
export default {
    async created(){
        const res = await getMyNeed()
        if(res.code == 0){
            console.log(res.data)
        }
    }
}
</script>