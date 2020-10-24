<template>
    <div :data-clazz="model.clazz">
        <div class="panelTitle">{{model.clazz === 'exclusiveGateway' || model.clazz === 'gateway' ? i18n['exclusiveGateway']
            : model.clazz === 'parallelGateway' ? i18n['parallelGateway'] : i18n['inclusiveGateway']}}</div>
        <div class="panelBody">
            <DefaultDetail :model="model" :onChange="onChange" :readOnly="readOnly" />
        </div>
    </div>
</template>
<script>
  import DefaultDetail from "./DefaultDetail";
  export default {
    inject: ['i18n'],
    components: {
      DefaultDetail
    },
    props: {
      model: {
        type:Object,
        default: ()=>({}),
      },
      onChange: {
        type: Function,
        default: ()=>{}
      },
      readOnly:{
        type: Boolean,
        default: false,
      }
    },
  }
</script>
