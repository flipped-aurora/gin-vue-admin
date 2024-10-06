<script setup>
import {onMounted, reactive,ref} from 'vue';
import { useRoute } from 'vue-router';
import {findBizCloudFunction} from "@/api/biz_apphub/bizCloudFunction";
import axios from "axios";
import UploadQiNiu from "@/components/upload_oss/UploadQiNiu.vue";
import {Check, Delete, Edit, Search, Star} from "@element-plus/icons-vue";
import { formatDate  } from '@/utils/format'
import {ElMessage} from "element-plus";
const route=useRoute()
const func =reactive({detail:{}})
const funcDefine=ref("")
const fileList=ref([])
const files=ref("")
const lookUpValue=ref("")
const lookUpTitle=ref("")
const centerDialogVisible=ref(false)
const funcCall=async ()=>{


  console.log("fileList",fileList.value)
  console.log("files",files.value)
  console.log("func detail",func.detail.param)
  // return
  let bd=reactive({})
  for (let i = 0; i < func.detail.param.length; i++) {
    if (func.detail.param[i].mode==="in"){
      if (func.detail.param[i].type==="file"){
        // func.detail.param[i].files=
        bd[func.detail.param[i].code]=func.detail.param[i].files
      }else {
        bd[func.detail.param[i].code]= func.detail.param[i].value
      }
    }
  }
  if (func.detail.api_config.method==="post"){
    let res =await axios.post(func.detail.api_config.path,bd)
    if(res.data.code===0){
      for (const key in res.data.data) {
        if (res.data.data.hasOwnProperty(key)) {
          const value = res.data.data[key];
          for (let i = 0; i < func.detail.param.length; i++) {
            console.log("func.detail.param[i]",func.detail.param[i])
            if (func.detail.param[i].code===key){
              func.detail.param[i].value=value
            }
          }
        }
      }
    }else {

    }
  }else {

  }
}

onMounted(() => {

  let fn = async ()=>{
    const res=await findBizCloudFunction({ID: route.query.id})
    if (res.code===0){
      func.detail=res.data
      for (let i = 0; i < func.detail.param.length; i++) {
        if (func.detail.param[i].type==="file"){
          func.detail.param[i].files=[]
        }
      }
    }
    getFuncInfo()
  }
  fn()
})
function getFuncInfo(){
  let inp=[]
  let outp=[]
  console.log("func.detail:",func.detail)
  for (let i = 0; i < func.detail.param.length; i++) {
    let item=func.detail.param[i]
    if (item.mode==="in"){
      inp.push(item.desc+" " + item.type)
    }else{
      outp.push(item.desc+" " + item.type)
    }
  }
  funcDefine.value= func.detail.code_name+"("+inp.join(",")+")"+"->"+"("+outp.join(",")+")"
}

function getTags() {
  if(func.detail.tags){
    return func.detail.tags.split(";")
  }
  return []

}
function lookUp(v) {
  centerDialogVisible.value=true
  lookUpTitle.value=v.desc
  lookUpValue.value=v.value
}


function copy(v) {
  // centerDialogVisible.value=true
  // lookUpTitle.value=v.desc
  // lookUpValue.value=v.value

  navigator.clipboard.writeText(v.value);
  ElMessage.success("复制成功")
}

function getLookUpTitle() {
  // centerDialogVisible.value=true
  return lookUpValue.value
}
const select = ref('')
const input3 = ref('')
</script>

<template>

  <div class="common-layout">
    <el-container>
      <el-header style="height: 70px">

        <div style="display: flex;justify-content: center;align-items:start" class="mt-4">
          <div style="flex: 2;display: flex;align-items: start;">
            <el-text style="font-size: 24px" class="mx-1" type="primary">云函数</el-text>

          </div>
            <div style="flex: 5">
              <el-input
                v-model="input3"
                style="max-width: 600px"
                placeholder="请输入要搜索的内容"
                class="input-with-select"
            >
              <template #prepend>
                <el-select v-model="select" placeholder="云函数" style="width: 115px">
                  <el-option label="Restaurant" value="1" />
                  <el-option label="Order No." value="2" />
                  <el-option label="Tel" value="3" />
                </el-select>
              </template>
              <template #append>
                <el-button :icon="Search" />
              </template>
            </el-input></div>
          <div style="display: flex;flex: 2;justify-content: end;align-items:start">
            <div><el-avatar :size="30" src="https://img1.baidu.com/it/u=1090452517,2487311686&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=501" /></div>
              <div style="display: flex;justify-content: center;align-items: center;margin-left: 10px">
                <Edit style="width: 30px; height: 30px;" />
                <span>发布函数</span>
              </div>
<!--            </el-icon>-->

          </div>
        </div>

      </el-header>
      <el-container>

        <el-main>
          <el-scrollbar>
            <div>

              <el-descriptions
                  class="margin-top"
                  title="函数详情"
                  :column="3"
                  border
              >
                <template #extra>
                </template>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">
                      中文名
                    </div>
                  </template>
                  {{func.detail.cn_name}}
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">
                      函数名称
                    </div>
                  </template>
                  {{func.detail.code_name}}
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">
                      函数定义
                    </div>
                  </template>
                  {{funcDefine}}
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">
                      归属
                    </div>
                  </template>
                  <el-tag size="small">{{func.detail.classify}}</el-tag>
                </el-descriptions-item>
                <el-descriptions-item>
                  <template #label>
                    <div class="cell-item">
                      函数介绍
                    </div>
                  </template>
                  {{func.detail.content}}
                </el-descriptions-item>
              </el-descriptions>
              <h2>
                <div style="">函数名称：{{func.detail.cn_name}}</div>
              </h2>
              <el-form ref="form" reset-type="initial" style="max-width: 100%">
                <el-tabs type="border-card">
                  <el-tab-pane label="运行">
                    <span>输入参数</span>
                    <div style="padding: 24px 24px 24px 0">
                      <!--            <t-form-item v-for="(v,i) in func.detail.param" v-show="v.mode==='in'" :label="v.desc" :label-width="150">-->
                      <el-form-item v-for="(v,i) in func.detail.param" v-show="v.mode==='in'" :label="v.desc">
                        <upload-qi-niu :uploaded-files="v.files" v-if="v.type==='file'" :title="v.desc" oss-dir="cloud_func/param_in"></upload-qi-niu>
                        <el-input v-model="v.value" v-else-if="v.type!=='file'&& v.input_mode==='text_field'||v.input_mode==='line'" style="width: 80%"
                                  :autosize="{ minRows: 3, maxRows: 5 }"
                                  :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"
                                  show-word-limit
                                  type="textarea"/>
                        <p style="margin-left: 10px">
                            <el-button type="primary" @click="lookUp(v)" key="预览" text >预览</el-button>
                          <el-button type="success" @click="copy(v)" key="复制" text >复制</el-button>
                        </p>

                      </el-form-item>

                      <!--todo 上面是输入参数-->
                      <t-button @click="funcCall" style="margin: 10px" block theme="primary" variant="base">运行</t-button>
                      <span>输出结果</span>
                      <!--todo 下面是输出参数-->
                      <!--            <t-form-item v-for="(v,i) in func.detail.param" v-show="v.mode==='out'" :label="v.desc" :label-width="150">-->
                      <el-form-item v-for="(v,i) in func.detail.param" v-show="v.mode==='out'" :label="v.desc" >

                        <div v-if="v.type==='file'">
                          <span style="margin: 10px">{{v.value.path}}</span>
                          <el-link v-if="v.value.path" :href="v.value.path" target="_blank" type="primary">下载</el-link>
                        </div>
                        <el-input v-model="v.value" v-else-if="v.type!=='file'&& v.input_mode==='text_field'||v.input_mode==='line'" style="width: 80%"
                                  :autosize="{ minRows: 3, maxRows: 5 }"
                                  :placeholder="v.mock_data===''?'请输入'+v.desc:v.mock_data"
                                  show-word-limit
                                  type="textarea"/>
<!--                        <el-input-->
<!--                            v-model="v.value"-->
<!--                            v-else-if="v.type!=='file'&& v.input_mode==='text_field'"-->
<!--                            style="width: 80%"-->
<!--                            :autosize="{ minRows: 3, maxRows: 5 }"-->
<!--                            :placeholder="v.mock_data===''?''+v.desc:v.mock_data"-->
<!--                            show-word-limit-->
<!--                            type="textarea"-->
<!--                        />-->
                        <p style="margin-left: 10px">
                          <el-button type="primary" @click="lookUp(v)" key="预览" text >预览</el-button>
                          <el-button type="success" @click="copy(v)" key="复制" text >复制</el-button>
                        </p>


                      </el-form-item>
                    </div>

                  </el-tab-pane>
                  <el-tab-pane label="案例">Config</el-tab-pane>
                </el-tabs>


              </el-form>


            </div>
          </el-scrollbar>

        </el-main>
<!--        右侧-->
        <el-aside width="300px">

          <div style="padding: 20px 5px 5px 0px">
            <div>
              <el-row>
                <el-col :span="6"><div class="grid-content ep-bg-purple" />
                  <el-avatar :size="50" src="https://img1.baidu.com/it/u=1465664392,2808406094&fm=253&fmt=auto&app=120&f=JPEG?w=800&h=800" />
                </el-col>
                <el-col :span="18">
                  <div>
                    <span style="font-size: 18px">就爱自娱自乐</span>
                    <!--                  <p>出生在中国，爱好写代码</p>-->
                    <p style="margin: 6px 0px">
                      <el-tooltip
                          class="box-item"
                          effect="dark"
                          content="出生在中国，爱好写代码，一个00后的独立开发者"
                          placement="top-start"
                      >
                        <el-text truncated size="small" class="mx-1">出生在中国，爱好写代码，一个00后的独立开发者</el-text>
                      </el-tooltip>
                    </p>
                  </div>
                </el-col>
              </el-row>

              <el-row>
                <el-col :span="6"><div class="grid-content ep-bg-purple" />
                </el-col>
                <el-col :span="9">
                  <el-button size="small" type="primary">+ 关注 988</el-button>
                </el-col>
                <el-col :span="9">
                  <el-button size="small" type="primary">打赏项目</el-button>
                </el-col>
              </el-row>
            </div>

            <el-divider border-style="double" />
            <h2>
              {{func.detail.cn_name}}
            </h2>
            <el-row>
              <el-col :span="6">
                <div>{{func.detail.classify}}</div>
              </el-col>
              <el-col :span="6">
                <div style="display: flex;align-items: center;">
                  <View  style="height: 16px"/>
                  <span style="margin-left: 2px">1200</span>
                </div>
              </el-col>

              <el-col :span="12"><div>
<!--                CreatedAt-->
                {{ formatDate(func.detail.CreatedAt) }}
              </div></el-col>
            </el-row>
            <h3>函数简介</h3>
            <el-text size="small" class="mx-1">{{func.detail.content}}</el-text>

<!--            <span>{{func.detail.content}}</span>-->
            <h3>相关Tag</h3>
            <div><el-tag effect="plain" style="margin: 3px" :key="i" v-for="(v,i) in getTags()">{{v}}</el-tag></div>
          </div>

          <el-row style="margin-top: 10px">
            <el-col :span="6">
              <div style="display: flex;align-items: center;">
                <el-tooltip
                    class="box-item"
                    effect="dark"
                    content="300星标"
                    placement="top-start"
                >
                <Star style="height: 30px"/>
                </el-tooltip>
                <span style="margin-left: 3px">200</span>
              </div>
            </el-col>
            <el-col :span="6">

              <div style="display: flex;align-items: center;">
<!--                <Coin style="height: 30px"/><span style="margin-left: 3px">300</span>-->
                <el-tooltip
                    class="box-item"
                    effect="dark"
                    content="300打赏"
                    placement="top-start"
                >
                  <Coin style="height: 30px"/>
                </el-tooltip>
                <span style="margin-left: 3px">200</span>
              </div>

            </el-col>

            <el-col :span="6"><div>

              <div style="display: flex;align-items: center;">
<!--                <el-icon><Cpu /></el-icon>-->
<!--                <CollectionTag style="height: 30px"/><span style="margin-left: 3px">300</span>-->

                <el-tooltip
                    class="box-item"
                    effect="dark"
                    content="300收藏"
                    placement="top-start"
                >
                  <CollectionTag style="height: 30px"/>
                </el-tooltip>
                <span style="margin-left: 3px">200</span>
              </div>
            </div>
            </el-col>

            <el-col :span="6"><div>

              <div style="display: flex;align-items: center;">
<!--                <Cpu style="height: 30px"/><span style="margin-left: 3px">300</span>-->

                <el-tooltip
                    class="box-item"
                    effect="dark"
                    content="300次运行"
                    placement="top-start"
                >
                  <Cpu style="height: 30px"/>
                </el-tooltip>
                <span style="margin-left: 3px">200</span>
              </div>
            </div>
            </el-col>

          </el-row>
        </el-aside>
      </el-container>
    </el-container>

<!--    预览文本-->
    <el-drawer direction="ltr" v-model="centerDialogVisible" :title="lookUpTitle" size="80%" center>
      <el-scrollbar height="98%">
        <div class="code-container">
          <pre id="code-pre" class="code-pre">{{lookUpValue}}</pre>
        </div>

      </el-scrollbar>
    </el-drawer>

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
