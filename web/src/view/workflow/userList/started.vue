<template>
    <div>
        
    </div>
</template>

<script>
import {getMyStated} from "@/api/workflowProcess"
export default {
    async created(){
        const res = await getMyStated()
        if(res.code == 0){
            console.log(res.data)
        }
    }
}
</script>