<script setup>




import {defineProps} from "vue";

const props = defineProps({
  content:{
    type: String,
    default:()=>""
  }
});

</script>

<template>
<div>
  {{props.content}}
</div>
</template>

<style scoped lang="scss">

</style>
