<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="函数中文名:" prop="cn_name">
          <el-input v-model="formData.cn_name" :clearable="true"  placeholder="请输入函数中文名" />
       </el-form-item>
        <el-form-item label="函数英文标识:" prop="code_name">
          <el-input v-model="formData.code_name" :clearable="true"  placeholder="请输入函数英文标识" />
       </el-form-item>
        <el-form-item label="函数分类:" prop="classify">
          <el-input v-model="formData.classify" :clearable="true"  placeholder="请输入函数分类" />
       </el-form-item>
        <el-form-item label="云函数执行方式:" prop="exec_mode">
           <el-select v-model="formData.exec_mode" placeholder="请选择云函数执行方式" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in cloud_func_exec_modeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="函数标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入函数标题" />
       </el-form-item>
        <el-form-item label="函数详细介绍:" prop="content">
          <el-input v-model="formData.content" :clearable="true"  placeholder="请输入函数详细介绍" />
       </el-form-item>
        <el-form-item label="内容类型:" prop="content_type">
           <el-select v-model="formData.content_type" placeholder="请选择内容类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ContentTypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="函数参数:" prop="param">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.param 后端会按照json的类型进行存取
          {{ formData.param }}
       </el-form-item>
        <el-form-item label="是否公开:" prop="is_public">
           <el-select v-model="formData.is_public" placeholder="请选择是否公开" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in bool_statusOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="接口配置:" prop="api_config">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.api_config 后端会按照json的类型进行存取
          {{ formData.api_config }}
       </el-form-item>
        <el-form-item label="JS代码:" prop="script_code">
          <el-input v-model="formData.script_code" :clearable="true"  placeholder="请输入JS代码" />
       </el-form-item>
        <el-form-item label="标签:" prop="tags">
          <el-input v-model="formData.tags" :clearable="true"  placeholder="请输入标签" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createBizCloudFunction,
  updateBizCloudFunction,
  findBizCloudFunction
} from '@/api/biz_apphub/bizCloudFunction'

defineOptions({
    name: 'BizCloudFunctionForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const bool_statusOptions = ref([])
const cloud_func_exec_modeOptions = ref([])
const ContentTypeOptions = ref([])
const formData = ref({
            cn_name: '',
            code_name: '',
            classify: '',
            exec_mode: '',
            title: '',
            content: '',
            content_type: '',
            param: {},
            is_public: '',
            api_config: {},
            script_code: '',
            tags: '',
        })
// 验证规则
const rule = reactive({
               cn_name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               code_name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               classify : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               exec_mode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               content_type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               param : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               is_public : [{
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
      const res = await findBizCloudFunction({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    bool_statusOptions.value = await getDictFunc('bool_status')
    cloud_func_exec_modeOptions.value = await getDictFunc('cloud_func_exec_mode')
    ContentTypeOptions.value = await getDictFunc('ContentType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createBizCloudFunction(formData.value)
               break
             case 'update':
               res = await updateBizCloudFunction(formData.value)
               break
             default:
               res = await createBizCloudFunction(formData.value)
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

<style>
</style>
