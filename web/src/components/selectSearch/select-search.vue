
<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="80px" size="medium"
    @submit.prevent>
    <el-form-item label="select" prop="select101173">
      <el-select v-model="formData.select101173" class="full-width-input" clearable filterable remote>
        <el-option v-for="(item, index) in select101173Options" :key="index" :label="item.label"
          :value="item.value" :disabled="item.disabled"></el-option>
      </el-select>
    </el-form-item>
  </el-form>
  
</template>

<script>
  import {
    defineComponent,
    toRefs,
    reactive,
    getCurrentInstance
  }
  from 'vue'
  export default defineComponent({
    components: {},
    props: {},
    setup() {
      const state = reactive({
        formData: {
          select101173: "",
        },
        rules: {},
        select101173Options: [{
          "label": "select 1",
          "value": 1
        }, {
          "label": "select 2",
          "value": 2
        }, {
          "label": "select 3",
          "value": 3
        }],
      })
      const instance = getCurrentInstance()
      const submitForm = () => {
        instance.ctx.$refs['vForm'].validate(valid => {
          if (!valid) return
          //TODO: 提交表单
        })
      }
      const resetForm = () => {
        instance.ctx.$refs['vForm'].resetFields()
      }
      return {
        ...toRefs(state),
        submitForm,
        resetForm
      }
    }
  })
  
</script>

<style lang="scss">
  .el-input-number.full-width-input,
  .el-cascader.full-width-input {
    width: 100% !important;
  }
  
  .el-form-item--medium {
    .el-radio {
      line-height: 36px !important;
    }
    
    .el-rate {
      margin-top: 8px;
    }
  }
  
  .el-form-item--small {
    .el-radio {
      line-height: 32px !important;
    }
    
    .el-rate {
      margin-top: 6px;
    }
  }
  
  .el-form-item--mini {
    .el-radio {
      line-height: 28px !important;
    }
    
    .el-rate {
      margin-top: 4px;
    }
  }
  
  .clear-fix:before,
  .clear-fix:after {
    display: table;
    content: "";
  }
  
  .clear-fix:after {
    clear: both;
  }
  
  .float-right {
    float: right;
  }
  
</style>

<style lang="scss" scoped>
  div.table-container {
    table.table-layout {
      width: 100%;
      table-layout: fixed;
      border-collapse: collapse;

      td.table-cell {
        display: table-cell;
        height: 36px;
        border: 1px solid #e1e2e3;
      }
    }
  }
  
  div.tab-container {}
  
  .label-left-align :deep(.el-form-item__label) {
    text-align: left;
  }
  
  .label-center-align :deep(.el-form-item__label) {
    text-align: center;
  }
  
  .label-right-align :deep(.el-form-item__label) {
    text-align: right;
  }
  
  .custom-label {}
  
  .static-content-item {
    min-height: 20px;
    display: flex;
    align-items: center;

    :deep(.el-divider--horizontal) {
      margin: 0;
    }
  }
  
</style>