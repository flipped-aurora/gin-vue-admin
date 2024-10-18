<script setup>
import {onMounted, reactive,ref} from 'vue';
import { useRoute } from 'vue-router';
import {findBizCloudFunction} from "@/api/biz_apphub/bizCloudFunction";
import axios from "axios";
import UploadQiNiu from "@/components/upload_oss/UploadQiNiu.vue";
import {ElMessage} from "element-plus";
import TableData from "@/components/funcViews/tableData.vue";
import TextView from "@/components/textView/textView.vue";
import Func_detail_right from "@/view/biz_apphub/funcHub/func_detail_right.vue";
import Func_info from "@/view/biz_apphub/funcHub/func_info.vue";
import Func_deader from "@/view/biz_apphub/funcHub/func_deader.vue";
import Func_param_desc from "@/view/biz_apphub/funcHub/func_param_desc.vue";
import Func_param_mock_data from "@/view/biz_apphub/funcHub/func_param_mock_data.vue";
const route=useRoute()
const func =reactive({detail:{}})
const funcDefine=ref("")
const fileList=ref([])
const files=ref("")
const lookUpValue=ref("")
const lookUpTitle=ref("")
const centerDialogVisible=ref(false)
const videoDialogVisible=ref(false)
const dialogTableVisible=ref(false)
const dialogTableDataList=reactive({})
const viewTableField=ref("")
const runData=ref({
  data:{}
}) //运行后的结果
const rules= reactive({})
const assetsSrc=ref("")
const assetsType=ref("")
const formRef = ref(null)


function showTable(param){
  viewTableField.value=param.code
  dialogTableDataList[param.code]=runData.data[param.code]
  dialogTableVisible.value=true
}

function getTableData(){
  return runData.data[viewTableField.value]
}


const funcCall=async ()=> {

  // let res=await formRef.value.validate()
  // if (!res){
  //   ElMessage.error("请输入参数信息")
  //   return
  // }
  console.log("param",func.detail.param)
  const startTime = performance.now();
  let bd = reactive({})
  let res;
  for (let i = 0; i < func.detail.param.length; i++) {
    if (func.detail.param[i].mode === "in") {
      if (func.detail.param[i].type === "file") {
        bd[func.detail.param[i].code] = func.detail.param[i].files
      } else {
        if(func.detail.param[i].required==="必填"){
          if (func.detail.param[i].value===""){
            ElMessage.error("请把参数信息输入完整")
            return
          }
        }
        if (func.detail.param[i].type==="muti_select"){
          func.detail.param[i].value=func.detail.param[i].checkList.join(";")
        }
        bd[func.detail.param[i].code] = func.detail.param[i].value
      }
    }
  }


  if (func.detail.api_config.method.toLowerCase() === "post") {
     res = await axios.post(func.detail.api_config.path, bd)
    if (res.data.data){
      runData.data=res.data.data
    }
    if (res.data.code === 0) {
      for (const key in res.data.data) {
        if (res.data.data.hasOwnProperty(key)) {
          const value = res.data.data[key];
          for (let i = 0; i < func.detail.param.length; i++) {
            if (func.detail.param[i].code === key) {
              func.detail.param[i].value = value
            }
          }
        }
      }
    }

  }else {
    // let res =await axios.get(func.detail.api_config.path,bd)
     res = await axios.get(func.detail.api_config.path, {
      params: bd
    })
    if (res.data.data){
      runData.data=res.data.data
    }

    const contentType = res.headers['content-type'];
    if (contentType && contentType.includes('image/')) {
      assetsType.value="image"
      assetsSrc.value=func.detail.api_config.path+"?"+new URLSearchParams(bd).toString()
      videoDialogVisible.value=true
      return
    }
    if (contentType && contentType.includes('video/')) {
      assetsType.value="video"
      assetsSrc.value=func.detail.api_config.path+"?"+new URLSearchParams(bd).toString()
      videoDialogVisible.value=true
      return
    }

    if (res.data.code === 0) {
        for (const key in res.data.data) {
        if (res.data.data.hasOwnProperty(key)) {
          const value = res.data.data[key];
          for (let i = 0; i < func.detail.param.length; i++) {
            if (func.detail.param[i].code === key) {
              func.detail.param[i].value = value
            }
          }
        }
      }
  }
}
  const endTime = performance.now();
  let cost=endTime - startTime
  run_info.run_cost=formatTime(cost)
  if (res.data.code===0){
    run_info.out_tab="res"
    run_info.run_status="运行成功"
    ElMessage.success("运行成功")
  }else {
    run_info.out_tab="res"
    run_info.run_status="运行失败，失败原因："+res.data.msg
    ElMessage.error("运行失败！")
  }
}
function formatTime(milliseconds) {
  const seconds = Math.floor(milliseconds / 1000);
  const remainingMilliseconds = milliseconds % 1000;


  if (seconds===0){
    return `${remainingMilliseconds.toFixed(2)}毫秒`
  }
  let s=remainingMilliseconds.toFixed(2)+""
  return `${seconds}.`+s.split(".")[0]+"秒";
}

onMounted(() => {

    let fn = async () => {
      const res = await findBizCloudFunction({ID: route.query.id})
      if (res.code === 0) {
        func.detail = res.data
        for (let i = 0; i < func.detail.param.length; i++) {
          if (func.detail.param[i].required==="必填"&&func.detail.param[i].mode==="in"){
            // func.detail.param[i].rules=[{
            //   required: true,
            //   message: '必填',
            //   trigger: "blur"
            // }]
          }
          if (func.detail.param[i].type === "sige_select") {
            if (func.detail.param[i].select_options===""){
              func.detail.param[i].value=func.detail.param[i].options.split(";")[0]
            }else {
              func.detail.param[i].value=func.detail.param[i].select_options
            }
          }
          if (func.detail.param[i].type === "muti_select") {
            func.detail.param[i].checkList=[]
            if (func.detail.param[i].select_options!==""){
              func.detail.param[i].checkList=func.detail.param[i].options.split(";")[0]
              func.detail.param[i].checkList= func.detail.param[i].select_options.split(";")
            }
          }
          if (func.detail.param[i].type === "file") {
            func.detail.param[i].files = []
          }
        }
      }
      getFuncInfo()
    }
    fn()
  })

function getFuncInfo() {
    let inp = []
    let outp = []
    console.log("func.detail:", func.detail)
    for (let i = 0; i < func.detail.param.length; i++) {
      let item = func.detail.param[i]
      if (item.mode === "in") {
        inp.push(item.desc + " " + item.type)
      } else {
        outp.push(item.desc + " " + item.type)
      }
    }
    funcDefine.value = func.detail.code_name + "(" + inp.join(",") + ")" + "->" + "(" + outp.join(",") + ")"
  }
function lookUp(v) {
    centerDialogVisible.value = true
    lookUpTitle.value = v.desc
    lookUpValue.value = v.value
  }

function copy(v) {
    navigator.clipboard.writeText(v.value);
    ElMessage.success("复制成功")
  }
const run_info=reactive({
  out_tab:"mock_data",
  in_tab:"run",
  run_status:"未运行",
  run_cost:""
})
const out_tab = ref('mock_data')
const in_tab = ref('run')
const run_status=ref("")

function getOpts(param){
  // param.value=param.select_options
  console.log("getOpts",param.options.split(";"))
  return param.options.split(";")
}

</script>

<template>

  <div class="common-layout">
    <el-container>
      <el-header style="height: 70px">
        <func_deader></func_deader>
      </el-header>
      <el-container>

        <el-main>
          <el-scrollbar>
            <div>
              <func_info v-show="func.detail.param" :func="func"></func_info>
              <el-form  ref="formRef" reset-type="initial" style="max-width: 100%">
                <h2>输入参数</h2>
<!--                <el-tabs type="border-card">-->
                <el-tabs v-model="run_info.in_tab" type="card">

                  <el-tab-pane name="info" label="参数介绍">
                    <func_param_desc :func="func" type="in"></func_param_desc>
                  </el-tab-pane>
                  <el-tab-pane name="mock_data" label="输入示例">
                    <func_param_mock_data :func="func" type="in"></func_param_mock_data>
                  </el-tab-pane>
                  <el-tab-pane name="run"  label="运行函数">
                    <div style="padding: 24px 24px 15px 24px">
                      <el-form-item v-for="(v,i) in func.detail.param"   v-show="v.mode==='in'" :label="v.desc" :required="v.required==='必填'">
                        <div v-if="v.type==='file'">
                          <upload-qi-niu :uploaded-files="v.files"  :title="v.desc" oss-dir="cloud_func/param_in"></upload-qi-niu>
                        </div>

                        <div style="width: 100%" v-else-if="v.type==='string'">
                          <p style="margin-left: 10px">
                          <el-input v-model="v.value"  style="width: 80%"
                                    :autosize="{ minRows: 1, maxRows: 3 }"
                                    :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"
                                    show-word-limit
                                    type="textarea"/>

                            <el-button type="primary" @click="lookUp(v)" key="预览" text >预览文本</el-button>
                            <el-button type="success" @click="copy(v)" key="复制" text >复制文本</el-button>
                          </p>
                        </div>
                        <div style="width: 80%" v-else-if="v.type==='muti_select'">
                          <el-checkbox-group v-model="v.checkList">
                            <!-- works when >=2.6.0, recommended ✔️ value not work when <2.6.0 ❌ -->
                            <el-checkbox v-for="opt in getOpts(v)" :label="opt" :value="opt" />
                            <!-- works when <2.6.0, deprecated act as value when >=3.0.0 -->
<!--                            <el-checkbox label="Option 2 & Value 2" />-->
                          </el-checkbox-group>
                        </div>
                        <div style="width: 80%" v-else-if="v.type==='sige_select'">
                          <el-select
                              filterable
                              v-model="v.value"
                              placeholder="Select"
                              size="large"
                              style="width: 100%"
                          >
                            <el-option
                                v-for="item in getOpts(v)"
                                :key="item"
                                :label="item"
                                :value="item"/>
                          </el-select>
<!--                          <el-checkbox-group v-model="v.checkList">-->
<!--                            &lt;!&ndash; works when >=2.6.0, recommended ✔️ value not work when <2.6.0 ❌ &ndash;&gt;-->
<!--                            <el-checkbox v-for="opt in getOpts(v)" :label="opt" :value="opt" />-->
<!--                            &lt;!&ndash; works when <2.6.0, deprecated act as value when >=3.0.0 &ndash;&gt;-->
<!--                            &lt;!&ndash;                            <el-checkbox label="Option 2 & Value 2" />&ndash;&gt;-->
<!--                          </el-checkbox-group>-->
                        </div>
                      </el-form-item>

                      <!--todo 上面是输入参数-->
                    </div>
                  </el-tab-pane>
                  <el-button @click="funcCall" type="primary" style="width: 100%"  block theme="primary" variant="base">运行</el-button>
                </el-tabs>

                <h2>输出结果</h2>
                <el-tabs  v-model="run_info.out_tab" type="card">

                <el-tab-pane  label="参数介绍" name="info">
                  <func_param_desc :func="func" type="out"></func_param_desc>
                </el-tab-pane>
                <el-tab-pane label="输出示例" name="mock_data">
                  <func_param_mock_data :func="func" type="out"></func_param_mock_data>
                </el-tab-pane>
                <el-tab-pane label="输出结果" name="res">
                  <div style="padding: 24px 24px 24px 24px">
                    <p style="padding: 0px 25px 20px 25px">状态：{{run_info.run_status}} <div v-if="run_info.run_cost!==''">耗时：{{run_info.run_cost}}</div></p>
                    <el-form-item  v-for="(v,i) in func.detail.param" v-show="v.mode==='out'" :label="v.desc"  >
                      <div v-if="v.type==='file'">
                        <span style="margin: 10px">{{v.value.path}}</span>
                        <el-link v-if="v.value.path" :href="v.value.path" target="_blank" type="primary">下载</el-link>
                      </div>
                      <div style="width: 100%;" v-else-if="v.type==='string'">
                        <p>
                          <el-input  v-model="v.value"
                                     style="width: 80%"
                                     :autosize="{ minRows: 1, maxRows: 3 }"
                                     :placeholder="v.desc"
                                     show-word-limit
                                     type="textarea"/>
                            <el-button v-if="v.type==='string'" type="primary" @click="lookUp(v)" key="预览" text >预览文本</el-button>
                            <el-button v-if="v.type==='string'" type="success" @click="copy(v)" key="复制" text >复制文本</el-button>
                        </p>
                      </div>
                      <div v-else-if="v.type==='table'">
                        <el-button @click="showTable(v)" type="primary">查看表格</el-button>
                      </div>
                    </el-form-item>
                  </div>
                </el-tab-pane>
                </el-tabs>
              </el-form>
            </div>
          </el-scrollbar>

        </el-main>
<!--        右侧-->
        <el-aside width="300px">
          <func_detail_right :func="func"></func_detail_right>
        </el-aside>
      </el-container>
    </el-container>

<!--    预览文本-->
    <el-drawer direction="ltr" v-model="centerDialogVisible" :title="lookUpTitle" size="80%" center>
      <text-view :content="lookUpValue">
        {{lookUpValue}}
      </text-view>
    </el-drawer>

    <el-dialog v-model="videoDialogVisible" title="静态文件" width="80%" center>
    <video controls width="100%" v-if="assetsType==='video'" :src="assetsSrc"></video>
      <p v-else-if="assetsType==='image'" style="text-align: center"><img width="50%" :src="assetsSrc"/></p>
    </el-dialog>


    <el-dialog v-model="dialogTableVisible" title="111" width="800">
      <table-data :table-data="getTableData()"></table-data>
    </el-dialog>


      </div>
</template>

<style scoped lang="scss">
.code-container {
  display: flex;
  align-items: flex-start;
}

.code-pre {
  white-space: pre-wrap;
  word-break: break-all;
  padding-right: 10px; /* 防止行号和代码重叠 */
}

.line-numbers {
  user-select: none;
  counter-reset: line;
  padding-left: 10px;
  border-right: 1px solid #ccc;
  margin-right: 10px;
}

.line-numbers span {
  display: block;
  counter-increment: line;
}

.line-numbers span::before {
  content: counter(line);
  padding-right: 5px;
}
</style>
