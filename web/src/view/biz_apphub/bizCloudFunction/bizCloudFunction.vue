
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>

        <el-form-item label="函数中文名" prop="cn_name">
         <el-input v-model="searchInfo.cn_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="函数英文标识" prop="code_name">
         <el-input v-model="searchInfo.code_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="函数分类" prop="classify">
         <el-input v-model="searchInfo.classify" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="云函数执行方式" prop="exec_mode">
            <el-select v-model="searchInfo.exec_mode" clearable placeholder="请选择" @clear="()=>{searchInfo.exec_mode=undefined}">
              <el-option v-for="(item,key) in cloud_func_exec_modeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="函数标题" prop="title">
         <el-input v-model="searchInfo.title" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="函数详细介绍" prop="content">
         <el-input v-model="searchInfo.content" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="内容类型" prop="content_type">
            <el-select v-model="searchInfo.content_type" clearable placeholder="请选择" @clear="()=>{searchInfo.content_type=undefined}">
              <el-option v-for="(item,key) in ContentTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
           <el-form-item label="是否公开" prop="is_public">
            <el-select v-model="searchInfo.is_public" clearable placeholder="请选择" @clear="()=>{searchInfo.is_public=undefined}">
              <el-option v-for="(item,key) in bool_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="标签" prop="tags">
         <el-input v-model="searchInfo.tags" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="浏览量" prop="views">

            <el-input v-model.number="searchInfo.startViews" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.endViews" placeholder="最大值" />

        </el-form-item>
        <el-form-item label="执行次数" prop="exec_count">

            <el-input v-model.number="searchInfo.startExecCount" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.endExecCount" placeholder="最大值" />

        </el-form-item>
        <el-form-item label="收藏数量" prop="coll">

            <el-input v-model.number="searchInfo.startColl" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.endColl" placeholder="最大值" />

        </el-form-item>
        <el-form-item label="点赞量" prop="like">

            <el-input v-model.number="searchInfo.startLike" placeholder="最小值" />
            —
            <el-input v-model.number="searchInfo.endLike" placeholder="最大值" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

          <el-table-column align="left" label="函数中文名" prop="cn_name" width="120" />
          <el-table-column align="left" label="函数英文标识" prop="code_name" width="120" />
          <el-table-column align="left" label="函数分类" prop="classify" width="120" />
        <el-table-column align="left" label="云函数执行方式" prop="exec_mode" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.exec_mode,cloud_func_exec_modeOptions) }}
            </template>
        </el-table-column>
          <el-table-column align="left" label="函数标题" prop="title" width="120" />
          <el-table-column align="left" label="函数详细介绍" prop="content" width="120" />
        <el-table-column align="left" label="内容类型" prop="content_type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.content_type,ContentTypeOptions) }}
            </template>
        </el-table-column>
          <el-table-column label="函数参数" prop="param" width="200">
              <template #default="scope">
                  [JSON]
              </template>
          </el-table-column>
        <el-table-column align="left" label="是否公开" prop="is_public" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.is_public,bool_statusOptions) }}
            </template>
        </el-table-column>
          <el-table-column align="left" label="标签" prop="tags" width="120" />
          <el-table-column sortable align="left" label="浏览量" prop="views" width="120" />
          <el-table-column align="left" label="执行次数" prop="exec_count" width="120" />
          <el-table-column align="left" label="收藏数量" prop="coll" width="120" />
          <el-table-column align="left" label="点赞量" prop="like" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateBizCloudFunctionFunc(scope.row)">变更</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="函数中文名:"  prop="cn_name" >
              <el-input v-model="formData.cn_name" :clearable="true"  placeholder="请输入函数中文名" />
            </el-form-item>
            <el-form-item label="函数英文标识:"  prop="code_name" >
              <el-input v-model="formData.code_name" :clearable="true"  placeholder="请输入函数英文标识" />
            </el-form-item>
            <el-form-item label="函数分类:"  prop="classify" >
              <el-input v-model="formData.classify" :clearable="true"  placeholder="请输入函数分类" />
            </el-form-item>
            <el-form-item label="云函数执行方式:"  prop="exec_mode" >
              <el-select v-model="formData.exec_mode" placeholder="请选择云函数执行方式" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in cloud_func_exec_modeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="函数标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入函数标题" />
            </el-form-item>
            <el-form-item label="函数详细介绍:"  prop="content" >
              <el-input v-model="formData.content" :clearable="true"  placeholder="请输入函数详细介绍" />
            </el-form-item>
            <el-form-item label="内容类型:"  prop="content_type" >
              <el-select v-model="formData.content_type" placeholder="请选择内容类型" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in ContentTypeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="函数参数:"  prop="param" >
<!--              // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.param 后端会按照json的类型进行存取-->
              <el-button @click="setParam">配置参数</el-button>
<!--              {{ formData.param }}-->
              {{getParam()}}
            </el-form-item>
            <el-form-item label="是否公开:"  prop="is_public" >
              <el-select v-model="formData.is_public" placeholder="请选择是否公开" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in bool_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="接口配置:"  prop="api_config" >
<!--              // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.api_config 后端会按照json的类型进行存取-->
              <el-button @click="setApiConfig">配置调用地址</el-button>
              {{ getApiConfig() }}

            </el-form-item>
            <el-form-item label="JS代码:"  prop="script_code" >
              <el-input v-model="formData.script_code" :clearable="true"  placeholder="请输入JS代码" />
            </el-form-item>
            <el-form-item label="标签:"  prop="tags" >
              <el-input v-model="formData.tags" :clearable="true"  placeholder="请输入标签" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions column="1" border>
                    <el-descriptions-item label="函数中文名">
                        {{ detailFrom.cn_name }}
                    </el-descriptions-item>
                    <el-descriptions-item label="函数英文标识">
                        {{ detailFrom.code_name }}
                    </el-descriptions-item>
                    <el-descriptions-item label="函数分类">
                        {{ detailFrom.classify }}
                    </el-descriptions-item>
                    <el-descriptions-item label="云函数执行方式">
                        {{ detailFrom.exec_mode }}
                    </el-descriptions-item>
                    <el-descriptions-item label="函数标题">
                        {{ detailFrom.title }}
                    </el-descriptions-item>
                    <el-descriptions-item label="函数详细介绍">
                        {{ detailFrom.content }}
                    </el-descriptions-item>
                    <el-descriptions-item label="内容类型">
                        {{ detailFrom.content_type }}
                    </el-descriptions-item>
                    <el-descriptions-item label="函数参数">
                        {{ detailFrom.param }}
                    </el-descriptions-item>
                    <el-descriptions-item label="是否公开">
                        {{ detailFrom.is_public }}
                    </el-descriptions-item>
                    <el-descriptions-item label="接口配置">
                        {{ detailFrom.api_config }}
                    </el-descriptions-item>
                    <el-descriptions-item label="JS代码">
                        {{ detailFrom.script_code }}
                    </el-descriptions-item>
                    <el-descriptions-item label="标签">
                        {{ detailFrom.tags }}
                    </el-descriptions-item>
                    <el-descriptions-item label="浏览量">
                        {{ detailFrom.views }}
                    </el-descriptions-item>
                    <el-descriptions-item label="执行次数">
                        {{ detailFrom.exec_count }}
                    </el-descriptions-item>
                    <el-descriptions-item label="收藏数量">
                        {{ detailFrom.coll }}
                    </el-descriptions-item>
                    <el-descriptions-item label="点赞量">
                        {{ detailFrom.like }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>


    <el-dialog v-model="setParamDialogVisible" title="配置函数参数">


            <el-form  :model="formData">

              <el-form-item v-for="(field,idx) in formData.param" :key="field.code">
                <el-row  :gutter="20">
                  <el-col :span="5">
                      <el-input v-model="field.code" placeholder="参数英文名称"></el-input>
                  </el-col>
                  <el-col :span="5">
                    <el-input v-model="field.desc" placeholder="参数描述"></el-input>
                  </el-col>

<!--                  <el-col :span="5">-->
<!--                    <el-input v-model="field.type" placeholder="参数类型"></el-input>-->
<!--                  </el-col>-->
<!--                  <el-col :span="5">-->
<!--                    <span>mock数据：</span><el-input v-model="field.mock_data" placeholder="mock_data"></el-input>-->
<!--                  </el-col>-->
                  <el-col :span="5">
<!--                    <span>参数类型：</span>-->
                    <el-select v-model="field.mode" placeholder="请选择模式">
                      <el-option label="输入参数" value="in"></el-option>
                       <el-option label="输出参数" value="out"></el-option>
                    </el-select>
                  </el-col>
                  <el-col :span="5">
                    <el-button type="primary" @click="removeField(idx)">移除</el-button>
                  </el-col>
                </el-row>
              </el-form-item>
            </el-form>
<!--      <el-col :span="5">-->
        <el-button type="primary" @click="addField">添加字段</el-button>
      <el-button type="primary" @click="setParamDialogVisible=false">确认</el-button>
<!--      </el-col>-->


    </el-dialog>

    <el-dialog v-model="setApiDialogVisible" title="配置函数调用地址">


      <el-form  label-width="auto" :model="formData">


<!--        <el-form :model="form" label-width="auto" style="max-width: 600px">-->
          <el-form-item label="接口地址">
            <el-input v-model="formData.api_config.path" />
          </el-form-item>
          <el-form-item label="请求方式">
            <el-select v-model="formData.api_config.method" placeholder="请求方式">
              <el-option label="post" value="post" />
              <el-option label="get" value="get" />
            </el-select>
          </el-form-item>

      </el-form>
      <!--      <el-col :span="5">-->
<!--      <el-button type="primary" @click="addField">添加字段</el-button>-->
      <el-button type="primary" @click="setApiDialogVisible=false">确认</el-button>
      <!--      </el-col>-->


    </el-dialog>
  </div>
</template>

<script setup>
import {
  createBizCloudFunction,
  deleteBizCloudFunction,
  deleteBizCloudFunctionByIds,
  updateBizCloudFunction,
  findBizCloudFunction,
  getBizCloudFunctionList
} from '@/api/biz_apphub/bizCloudFunction'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import {Paperclip} from "@element-plus/icons-vue";

defineOptions({
    name: 'BizCloudFunction'
})

const setParamDialogVisible=ref(false)
const setApiDialogVisible=ref(false)
const setParam=function (){
  setParamDialogVisible.value=true
}
const setApiConfig=function (){
  setApiDialogVisible.value=true
}

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

function getApiConfig(){
  if (!formData.value.api_config.method){
    return "未配置"
  }
  return "http "+formData.value.api_config.method+" "+formData.value.api_config.path
}

function getParam(){
  let inp=[]
  let outp=[]
  for (let i = 0; i < formData.value.param.length; i++) {
    let item=formData.value.param[i]
    if (item.mode==="in"){
      inp.push(item.code)
    }else{
      outp.push(item.code)
    }
  }
  return formData.value.code_name+"("+inp.join(",")+")"+"->"+"("+outp.join(",")+")"
}
const addField = () => {
  formData.value.param.push({
    code: '',
    desc: '',
    mode: 'in',
    mock_data: '',
    type: 'string',
    value: ''
  });
};

const removeField = (index) => {
  formData.value.param.splice(index, 1);
};

// 自动化生成的字典（可能为空）以及字段
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
            param: [],
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               code_name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               classify : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               exec_mode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               content_type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               param : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               is_public : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            views: 'views',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getBizCloudFunctionList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    bool_statusOptions.value = await getDictFunc('bool_status')
    cloud_func_exec_modeOptions.value = await getDictFunc('cloud_func_exec_mode')
    ContentTypeOptions.value = await getDictFunc('ContentType')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteBizCloudFunctionFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteBizCloudFunctionByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateBizCloudFunctionFunc = async(row) => {
    const res = await findBizCloudFunction({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBizCloudFunctionFunc = async (row) => {
    const res = await deleteBizCloudFunction({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findBizCloudFunction({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
