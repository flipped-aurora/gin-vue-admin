<template>
  <div :data-clazz="model.clazz">
    <div class="panelTitle">{{ i18n['endEvent'] }}</div>
    <div class="panelBody">
      <DefaultDetail :model="model" :onChange="onChange" :readOnly="readOnly" />
      <div class="panelRow">
        <el-checkbox
          @change="(value) => onChange('success', value)"
          :disabled="readOnly"
          :value="!!model.success"
          >是否成功</el-checkbox
        >
      </div>
      <div class="panelRow">
        <div>步骤：</div>
        <el-input
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          :value="model.step"
          placeholder="请输入步骤"
          @input="
            (value) => {
              onChange('step', value)
            }
          "
        />
      </div>
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
