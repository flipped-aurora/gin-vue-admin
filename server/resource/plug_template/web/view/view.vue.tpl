<template>
  <div class="{{ .Snake }}">
    在此处书写页面代码
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { routerName } from '@/plugin/{{ .Snake }}/api/api.js'

const data = ref({})

defineOptions({
  name: '{{ .PlugName }}'
})

const useApi = async() =>{
    const res = await routerName(data.value)
    if(res.code === 0){
      console.log(res.data)
    }
}


</script>
<style lang="scss" scoped>
.{{ .Snake }} {

}
</style>
