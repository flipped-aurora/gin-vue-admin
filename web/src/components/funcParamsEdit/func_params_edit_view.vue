<script setup>

import {defineProps,watch, onMounted, ref} from "vue";
import {createBizCloudFunction, updateBizCloudFunction} from "@/api/biz_apphub/bizCloudFunction";
import {ElMessage} from "element-plus";

const props = defineProps({
  formData:{
    type: Object,
  },
  runner_name:{
    type: String,
  },
});

onMounted(() => {
  if(props.formData.english_name===""){
    props.formData.english_name=props.formData.path.replaceAll("/",".")
  }
  console.log("props.formData:",props.formData)
})


const enterDialog = async () => {

  let res = await createBizCloudFunction(props.formData)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建成功'
    })
    emit('updateShowParam', false);
  }else {
    ElMessage({
      type: 'error',
      message: res.msg
    })
  }
  // elFormRef.value?.validate( async (valid) => {
  //   if (!valid) return
  //   let res
  //   switch (type.value) {
  //     case 'create':
  //       res = await createBizCloudFunction(formData.value)
  //       break
  //     case 'update':
  //       res = await updateBizCloudFunction(formData.value)
  //       break
  //     default:
  //       res = await createBizCloudFunction(formData.value)
  //       break
  //   }
  //   if (res.code === 0) {
  //     ElMessage({
  //       type: 'success',
  //       message: '创建/更改成功'
  //     })
  //     closeDialog()
  //     getTableData()
  //   }
  // })
}
const emit = defineEmits(['updateShowParam']);
const param_in_type = ref("string")
const showParam = ref(true)
const param_out_type = ref("string")
// const http_method = ref("")
// const out_type=ref("")
const isOutStaticFile = ref(false)
const removeField = (index) => {
  if (formData.value.param[index].type==="static_file"){
    isOutStaticFile.value=false
  }
  formData.value.param.splice(index, 1);
};


const createCloudFunc =async function() {
   enterDialog()

  // showParam.value=false
}

// 监听本地响应式数据的变化，发送 update:visible 事件，由于是双向绑定，父组件的值会直接改变。这里传递的 visible 是 ref 对象，
// 可以直接传递，Vue 自动解包数据，也可使用 () => visible.value
// watch(showParam, (newVal) => {
//   // this.emit('update:showParam', newVal)
//   emit('updateShowParam', newVal);
// })
const addInField = () => {

  formData.value.param.push({
    code: '',
    desc: '',
    mode: 'in',
    mock_data: '',
    input_mode: 'text_field',
    required:"必填",
    select_options:"",
    file_size_limit:"",
    file_type_limit:"",
    number_limit:"",
    text_limit:"",
    options:"",
    type: param_in_type.value,
    value: ''
  });
};
const addOutField = () => {

  formData.value.param.push({
    code: '',
    desc: '',
    mode: 'out',
    required:"必填",
    mock_data: '',
    select_options:"",
    input_mode: 'text_field',
    options:"",
    type: param_out_type.value,
    value: ''
  });
};

const addOutStaticFile = () => {
  let flag=!isOutStaticFile.value
  isOutStaticFile.value=flag
  if (!flag){
    formData.value.param.filter(param=>{
      return param.mode!=="out_static_file"
    })
  }else {
    formData.value.param.push({
      code: '',
      desc: '',
      mode: 'out_static_file',
      mock_data: '',
      select_options:"",
      options:"",
      type: "static_file",
      value: ''
    });
  }

};
function getParamType(field) {
  if (field.type==="string"){
    return "字符串"
  }else if(field.type==="number"){
    return "数值"
  }else if(field.type==="file"){
    return "文件"
  }else if(field.type==="muti_select"){
    return "多选框"
  }else if(field.type==="sige_select"){
    return "单选框"
  }else if(field.type==="table"){
    return "表格"
  }
}
const activeNames=ref(["1","2","3"])

</script>

<template>
    <div>
      <p>
        <el-button @click="createCloudFunc()" type="primary" style="width: 100%">提交</el-button>
      </p>
      <h2>接口配置</h2>
      <span>请求方式：</span><el-select
        v-model="props.formData.api_config.method"
        placeholder="Select"
        size="large"
        style="width: 140px"
    >
      <el-option
          key=GET
          label="GET"
          value="GET"
      />
      <el-option
          key="POST"
          label="POST"
          value="POST"
      />
    </el-select>

      <span>调用地址：</span><el-input style="width: 200px" v-model="formData.api_config.path"></el-input>
      <span>函数中文：</span><el-input style="width: 200px" v-model="formData.cn_name"></el-input>
      <span>函数英文标识：</span><el-input style="width: 200px" v-model="formData.code_name"></el-input>
      <span>分类：</span><el-input style="width: 200px" v-model="formData.classify"></el-input>
      <h2>参数配置</h2>
      <el-collapse v-model="activeNames" >
        <el-collapse-item title="输入参数" name="1">
          <div>
            <el-form  :model="formData">
              <el-form-item  v-for="(field,idx) in formData.param" v-show="field.mode==='in'" :key="field.code">
                <div  style="width: 100%">
                  <el-row style="width: 100%" :gutter="20">
                    <el-col :span="2">
                      <el-input v-model="field.code" type="textarea" placeholder="参数英文名称"></el-input>
                    </el-col>
                    <el-col :span="3">
                      <el-input v-model="field.desc" type="textarea" placeholder="参数中文名称"></el-input>
                    </el-col>
                    <el-col :span="4">
                      <el-input v-model="field.note" type="textarea" placeholder="参数注释"></el-input>
                    </el-col>
                    <el-col :span="4">
                      <el-select v-model="field.required" placeholder="是否必填" >
                        <el-option  label="必填" value="必填" />
                        <el-option  label="非必填" value="非必填" />
                      </el-select>
                    </el-col>

                    <el-col v-if="field.type==='string' || field.type==='number'" :span="6">
                      <el-input v-model="field.mock_data" type="textarea" placeholder="示例数据"></el-input>
                    </el-col>
                    <el-col v-if="field.type==='string'" :span="3">
                      <el-input v-model="field.text_limit" type="textarea" placeholder="限制字符数：例如：120，为空不限制"></el-input>
                    </el-col>
                    <el-col v-if="field.type==='number'" :span="6">
                      <el-input v-model="field.number_limit" type="textarea" placeholder="限制数值范围：例如：1-100，为空不限制"></el-input>
                    </el-col>
                    <el-col :span="6" v-if="field.type==='muti_select' || field.type==='sige_select'" >
                      <el-input v-model="field.options" type="textarea" placeholder="分号分割多个选项，例如：压缩;加水印;加密"></el-input>
                    </el-col>

                    <el-col :span="6" v-if="field.type==='muti_select'" >
                      <el-input v-model="field.select_options" type="textarea" placeholder="默认选项;多个请用分号分割;例如：压缩;加水印"></el-input>
                    </el-col>

                    <el-col :span="6" v-if="field.type==='sige_select'" >
                      <el-input v-model="field.select_options" type="textarea" placeholder="默认选项，例如：压缩"></el-input>
                    </el-col>

                    <el-col :span="6" v-if="field.type==='file'" >
                      <el-input v-model="field.file_type_limit" type="textarea" placeholder="限制文件类型，例如：jpg;png， 为空不限制"></el-input>
                    </el-col>
                    <el-col :span="6" v-if="field.type==='file'" >
                      <el-input v-model="field.file_size_limit" type="textarea" placeholder="限制文件大小(单位kb)，例如：4096，为空不限制"></el-input>
                    </el-col>

                    <el-col :span="1">
                      <span>{{getParamType(field)}}</span>
                    </el-col>

                    <el-col :span="1">
                      <el-button type="primary" @click="removeField(idx)">移除</el-button>
                    </el-col>
                  </el-row>
                </div>
              </el-form-item>
            </el-form>
          </div>
        </el-collapse-item>
        <el-collapse-item title="输出参数" name="2">
          <div>
            <p v-if="isOutStaticFile" style="text-align: center">输出静态文件</p>
            <el-form v-else  :model="formData">

              <el-form-item v-for="(field,idx) in formData.param" v-show="field.mode==='out'" :key="field.code">
                <div  style="width: 100%">
                  <el-row style="width: 100%" :gutter="20">
                    <el-col :span="2">
                      <el-input v-model="field.code" type="textarea" placeholder="参数英文名称"></el-input>
                    </el-col>
                    <el-col :span="3">
                      <el-input v-model="field.desc" type="textarea" placeholder="参数中文"></el-input>
                    </el-col>
                    <el-col :span="5">
                      <el-input v-model="field.note" type="textarea" placeholder="参数注释"></el-input>
                    </el-col>

                    <el-col :span="6" v-if="field.type==='string' || field.type==='number'" >
                      <el-input v-model="field.mock_data" type="textarea" placeholder="示例数据"></el-input>
                    </el-col>


                    <el-col v-if="field.type==='string'" :span="6">
                      <el-input v-model="field.text_limit" type="textarea" placeholder="限制字符数：例如：120，为空不限制"></el-input>
                    </el-col>

                    <el-col :span="6" v-if="field.type==='muti_select' || field.type==='sige_select'" >
                      <el-input v-model="field.options" type="textarea" placeholder="分号分割多个选项，例如：压缩;加水印;加密"></el-input>
                    </el-col>


                    <el-col :span="6" v-if="field.type==='muti_select'" >
                      <el-input v-model="field.select_options" type="textarea" placeholder="默认选项;多个请用分号分割;例如：压缩;加水印"></el-input>
                    </el-col>

                    <el-col :span="6" v-if="field.type==='sige_select'" >
                      <el-input v-model="field.select_options" type="textarea" placeholder="默认选项，例如：压缩"></el-input>
                    </el-col>

                    <el-col :span="1">
                      <span>{{getParamType(field)}}</span>
                    </el-col>
                    <el-col :span="1">
                      <el-button type="primary" @click="removeField(idx)">移除</el-button>
                    </el-col>
                  </el-row>
                </div>
              </el-form-item>
            </el-form>
          </div>
        </el-collapse-item>
        <el-collapse-item title="参数操作" name="3">
          <div>
            <div style="margin: 20px 0px">
              <p style="text-align: center">
                <el-radio-group v-model="param_in_type">
                  <el-radio value="string" size="large">文本</el-radio>
                  <el-radio value="number" size="large">数值</el-radio>
                  <el-radio value="file" size="large">文件</el-radio>
                  <el-radio value="muti_select" size="large">多选框</el-radio>
                  <el-radio value="sige_select" size="large">单选框</el-radio>
                </el-radio-group>
              </p>
              <p style="margin: 10px 10px"><el-button style="width: 100%" type="primary" @click="addInField">添加输入参数</el-button></p>

              <p style="text-align: center">
                <el-radio-group v-model="param_out_type">
                  <el-radio value="string" size="large">文本</el-radio>
                  <el-radio value="number" size="large">数值</el-radio>
                  <el-radio value="file" size="large">文件</el-radio>
                  <el-radio value="table" size="large">表格</el-radio>
                </el-radio-group>
              </p>
              <p v-if="isOutStaticFile" style="margin: 10px 10px"><el-button style="width: 100%" type="primary"  disabled  @click="addOutField">添加输出参数</el-button></p>
              <p v-else style="margin: 10px 10px"><el-button style="width: 100%" type="primary"  @click="addOutField">添加输出参数</el-button></p>
              <p style="margin: 10px 10px"><el-button style="width: 100%" type="primary" @click="addOutStaticFile">{{isOutStaticFile?"取消输出静态文件":"输出静态文件"}}</el-button></p>
            </div>
          </div>
        </el-collapse-item>
        <!--    <el-collapse-item title="Controllability" name="4">-->
        <!--      <div>-->
        <!--        Decision making: giving advices about operations is acceptable, but do-->
        <!--        not make decisions for the users;-->
        <!--      </div>-->
        <!--    </el-collapse-item>-->
      </el-collapse>




      <!--  <p style="margin: 10px 10px"><el-button style="width: 100%" type="primary" @click="setParamDialogVisible=false">确认</el-button></p>-->

    </div>
</template>

<style scoped lang="scss">

</style>
