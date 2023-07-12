<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="top" label-width="80px" size="medium"
    @submit.native.prevent>
    <el-row>
      <el-col :span="2" class="grid-cell">
        <div class="static-content-item">
          <el-button type="primary">保存</el-button>
        </div>
      </el-col>
      <el-col :span="2" class="grid-cell">
        <div class="static-content-item">
          <el-button type="danger">返回</el-button>
        </div>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="12" class="grid-cell">
        <el-form-item label="机巢ID" prop="nestid">
          <el-input v-model="formData.nestid" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12" class="grid-cell">
        <el-form-item label="机巢名称" prop="nestid">
          <el-input v-model="formData.nestName" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12" class="grid-cell">
        <el-form-item label="机巢坐标" prop="nestLocation">
          <el-input v-model="formData.nestLocation" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24" class="grid-cell">
        <el-form-item label="飞机推流地址" prop="aircraftVideoPushURL">
          <el-input v-model="formData.aircraftVideoPushURL" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24" class="grid-cell">
        <el-form-item label="飞机直播地址" prop="aircraftVideoURL">
          <el-input v-model="formData.aircraftVideoURL" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24" class="grid-cell">
        <el-form-item label="机巢监控地址" prop="nestVideoURL">
          <el-input v-model="formData.nestVideoURL" type="text" clearable></el-input>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
  
</template>

<script>
export default {
  name: 'NestInfo'
}
</script>

<script setup>
import {
  createNestInfo,
  updateNestInfo,
  findNestInfo
} from '@/api/nestInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            nestid: '',
            nestName: '',
            nestLocation: '',
            aircraftVideoPushURL: '',
            aircraftVideoURL: '',
            nestVideoURL: '',
        })
// 验证规则
const rule = reactive({
               nestid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nestLocation : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findNestInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.renestinfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createNestInfo(formData.value)
               break
             case 'update':
               res = await updateNestInfo(formData.value)
               break
             default:
               res = await createNestInfo(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

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
  
  .label-left-align ::v-deep .el-form-item__label {
    text-align: left;
  }
  
  .label-center-align ::v-deep .el-form-item__label {
    text-align: center;
  }
  
  .label-right-align ::v-deep .el-form-item__label {
    text-align: right;
  }
  
  .custom-label {}
  
  .static-content-item {
    min-height: 20px;
    display: flex;
    align-items: center;

    ::v-deep .el-divider--horizontal {
      margin: 0;
    }
  }
  
</style>
