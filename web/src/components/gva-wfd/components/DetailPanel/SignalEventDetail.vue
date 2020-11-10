<template>
  <div :data-clazz="model.clazz">
    <div class="panelTitle">{{ i18n['signalEvent'] }}</div>
    <div class="panelBody">
      <DefaultDetail :model="model" :onChange="onChange" :readOnly="readOnly" />
      <div class="panelRow">
        <div>{{ i18n['signalEvent.signal'] }}：</div>
        <el-select
          style="width: 90%; font-size: 12px"
          :placeholder="i18n['signalEvent.signal']"
          :value="model.signal"
          :disabled="readOnly"
          @change="
            (e) => {
              onChange('signal', e)
            }
          "
        >
          <el-option
            v-for="signal in signalDefs"
            :key="signal.id"
            :label="signal.name"
            :value="signal.id"
          />
        </el-select>
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
      },
      signalDefs: {
        type: Array,
        default: ()=>([]),
      },
    },
  }
</script>
