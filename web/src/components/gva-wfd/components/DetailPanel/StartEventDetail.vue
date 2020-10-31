<template>
    <div :data-clazz="model.clazz">
        <div class="panelTitle">{{i18n['startEvent']}}</div>
        <div class="panelBody">
            <DefaultDetail :model="model" :onChange="onChange" :readOnly="readOnly" />
        </div>
         <div class="panelRow">
                <div>视图路径（以view开头）</div>
                <el-input style="width:90%; font-size:12px"
                          :disabled="readOnly"
                          type="view"
                          :value="model.view"
                          @input="(value) => {onChange('view', value)}" />
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
