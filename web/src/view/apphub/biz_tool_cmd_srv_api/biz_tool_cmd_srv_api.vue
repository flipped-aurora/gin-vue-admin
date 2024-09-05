
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
      
        <el-form-item label="应用名称（中文）" prop="appName">
         <el-input v-model="searchInfo.appName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="应用名称（英文标识）" prop="appCode">
         <el-input v-model="searchInfo.appCode" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="工具类型" prop="tool_type">
         <el-input v-model="searchInfo.tool_type" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="标题" prop="title">
         <el-input v-model="searchInfo.title" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="应用介绍" prop="desc">
         <el-input v-model="searchInfo.desc" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="分类" prop="classify">
         <el-input v-model="searchInfo.classify" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="应用版本" prop="version">
         <el-input v-model="searchInfo.version" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="收费模式" prop="mode">
            <el-select v-model="searchInfo.mode" clearable placeholder="请选择" @clear="()=>{searchInfo.mode=undefined}">
              <el-option v-for="(item,key) in price_modeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="应用标签" prop="tags">
         <el-input v-model="searchInfo.tags" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="介绍视频" prop="video">
         <el-input v-model="searchInfo.video" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="接口地址" prop="api_path">
         <el-input v-model="searchInfo.api_path" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="是否公开" prop="is_public">
            <el-select v-model="searchInfo.is_public" clearable placeholder="请选择" @clear="()=>{searchInfo.is_public=undefined}">
              <el-option v-for="(item,key) in bool_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
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
        
          <el-table-column align="left" label="应用名称（中文）" prop="appName" width="120" />
          <el-table-column align="left" label="应用名称（英文标识）" prop="appCode" width="120" />
          <el-table-column align="left" label="工具类型" prop="tool_type" width="120" />
          <el-table-column align="left" label="标题" prop="title" width="120" />
          <el-table-column align="left" label="应用介绍" prop="desc" width="120" />
          <el-table-column align="left" label="分类" prop="classify" width="120" />
          <el-table-column sortable align="left" label="应用版本" prop="version" width="120" />
        <el-table-column align="left" label="收费模式" prop="mode" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.mode,price_modeOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="后续迭代" prop="developMode" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.developMode,dev_modeOptions) }}
            </template>
        </el-table-column>
          <el-table-column align="left" label="封面地址" prop="cover" width="120" />
          <el-table-column align="left" label="应用标签" prop="tags" width="120" />
          <el-table-column align="left" label="介绍视频" prop="video" width="120" />
          <el-table-column align="left" label="状态" prop="status" width="120" />
          <el-table-column align="left" label="接口地址" prop="api_path" width="120" />
        <el-table-column align="left" label="是否公开" prop="is_public" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.is_public,bool_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateBizToolCmdSrvApiFunc(scope.row)">变更</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateBizToolCmdSrvApiVersionFunc(scope.row)">变更版本</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
              <el-button  type="primary" link icon="edit" @click="deployRecordFn(scope.row)">发布历史</el-button>
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

<!--    更新数据-->
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
            <el-form-item label="应用名称（中文）:"  prop="appName" >
              <el-input v-model="formData.appName" :clearable="true"  placeholder="请输入应用名称（中文）" />
            </el-form-item>
            <el-form-item label="应用名称（英文标识）:"  prop="appCode" >
              <el-input v-model="formData.appCode" :clearable="true"  placeholder="请输入应用名称（英文标识）" />
            </el-form-item>
<!--            <el-form-item label="工具类型1:"  prop="tool_type" >-->
<!--              <el-input v-model="formData.tool_type" :clearable="true"  placeholder="请输入工具类型" />-->
<!--            </el-form-item>-->
            <el-form-item label="工具类型:" prop="mode">
              <el-select v-model="formData.tool_type" placeholder="请选择工具类型" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in tool_type" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>

            <el-form-item label="标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
            </el-form-item>
            <el-form-item label="应用介绍:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入应用介绍" />
            </el-form-item>
            <el-form-item label="分类:"  prop="classify" >
              <el-input v-model="formData.classify" :clearable="true"  placeholder="请输入分类" />
            </el-form-item>
            <el-form-item label="应用版本:"  prop="version" >
              <el-input v-model="formData.version" :clearable="true"  placeholder="请输入应用版本" />
            </el-form-item>
            <el-form-item label="收费模式:"  prop="mode" >
              <el-select v-model="formData.mode" placeholder="请选择收费模式" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in price_modeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="后续迭代:"  prop="developMode" >
              <el-select v-model="formData.developMode" placeholder="请选择后续迭代" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in dev_modeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="封面地址:"  prop="cover" >
              <el-input v-model="formData.cover" :clearable="true"  placeholder="请输入封面地址" />
            </el-form-item>
            <el-form-item label="应用标签:"  prop="tags" >
              <el-input v-model="formData.tags" :clearable="true"  placeholder="请输入应用标签" />
            </el-form-item>
            <el-form-item label="介绍视频:"  prop="video" >
              <el-input v-model="formData.video" :clearable="true"  placeholder="请输入介绍视频" />
            </el-form-item>
            <el-form-item label="是否公开:"  prop="is_public" >
              <el-select v-model="formData.is_public" placeholder="请选择是否公开" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in bool_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>

            <el-form-item label="上传文件:"  prop="video" >
              <UploadQiNiu oss-dir="tool" :uploadedFiles="uploadedFiles" title='请把打包后的文件压缩成zip格式上传'/>
            </el-form-item>
          </el-form>
    </el-drawer>

    <!--    变更版本-->
    <el-drawer destroy-on-close size="800" v-model="dialogFormVersionVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>


      <el-collapse v-model="activeNames" @change="handleChange">
        <el-collapse-item title="必填项" name="1">
          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="应用版本:"  prop="version" >
              <el-input v-model="formData.version" :clearable="true"  placeholder="请输入应用版本" />
            </el-form-item>
            <el-form-item  label="变更日志:"  prop="version" >
              <el-input type="textarea" v-model="formData.remark" :clearable="true"  placeholder="请输入变更日志" />
            </el-form-item>

            <el-form-item label="上传文件:"  prop="video" >
              <UploadQiNiu oss-dir="tool" :uploadedFiles="uploadedFiles" title='请把打包后的文件压缩成zip格式上传'/>
            </el-form-item>
          </el-form>
        </el-collapse-item>
        <el-collapse-item title="其他参数" name="2">
          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="应用名称（中文）:"  prop="appName" >
              <el-input v-model="formData.appName" :clearable="true"  placeholder="请输入应用名称（中文）" />
            </el-form-item>
            <el-form-item label="应用名称（英文标识）:"  prop="appCode" >
              <el-input v-model="formData.appCode" :clearable="true"  placeholder="请输入应用名称（英文标识）" />
            </el-form-item>
            <!--            <el-form-item label="工具类型1:"  prop="tool_type" >-->
            <!--              <el-input v-model="formData.tool_type" :clearable="true"  placeholder="请输入工具类型" />-->
            <!--            </el-form-item>-->
            <el-form-item label="工具类型:" prop="mode">
              <el-select v-model="formData.tool_type" placeholder="请选择工具类型" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in tool_type" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>

            <el-form-item label="标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
            </el-form-item>
            <el-form-item label="应用介绍:"  prop="desc" >
              <el-input v-model="formData.desc" :clearable="true"  placeholder="请输入应用介绍" />
            </el-form-item>
            <el-form-item label="分类:"  prop="classify" >
              <el-input v-model="formData.classify" :clearable="true"  placeholder="请输入分类" />
            </el-form-item>
<!--            <el-form-item label="应用版本:"  prop="version" >-->
<!--              <el-input v-model="formData.version" :clearable="true"  placeholder="请输入应用版本" />-->
<!--            </el-form-item>-->
            <el-form-item label="收费模式:"  prop="mode" >
              <el-select v-model="formData.mode" placeholder="请选择收费模式" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in price_modeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="后续迭代:"  prop="developMode" >
              <el-select v-model="formData.developMode" placeholder="请选择后续迭代" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in dev_modeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="封面地址:"  prop="cover" >
              <el-input v-model="formData.cover" :clearable="true"  placeholder="请输入封面地址" />
            </el-form-item>
            <el-form-item label="应用标签:"  prop="tags" >
              <el-input v-model="formData.tags" :clearable="true"  placeholder="请输入应用标签" />
            </el-form-item>
            <el-form-item label="介绍视频:"  prop="video" >
              <el-input v-model="formData.video" :clearable="true"  placeholder="请输入介绍视频" />
            </el-form-item>
            <el-form-item label="是否公开:"  prop="is_public" >
              <el-select v-model="formData.is_public" placeholder="请选择是否公开" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in bool_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>

          </el-form>

        </el-collapse-item>

      </el-collapse>


    </el-drawer>
    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions column="1" border>
                    <el-descriptions-item label="应用名称（中文）">
                        {{ detailFrom.appName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="应用名称（英文标识）">
                        {{ detailFrom.appCode }}
                    </el-descriptions-item>
                    <el-descriptions-item label="工具类型">
                        {{ detailFrom.tool_type }}
                    </el-descriptions-item>
                    <el-descriptions-item label="标题">
                        {{ detailFrom.title }}
                    </el-descriptions-item>
                    <el-descriptions-item label="应用介绍">
                        {{ detailFrom.desc }}
                    </el-descriptions-item>
                    <el-descriptions-item label="分类">
                        {{ detailFrom.classify }}
                    </el-descriptions-item>
                    <el-descriptions-item label="应用版本">
                        {{ detailFrom.version }}
                    </el-descriptions-item>
                    <el-descriptions-item label="收费模式">
                        {{ detailFrom.mode }}
                    </el-descriptions-item>
                    <el-descriptions-item label="后续迭代">
                        {{ detailFrom.developMode }}
                    </el-descriptions-item>
                    <el-descriptions-item label="封面地址">
                        {{ detailFrom.cover }}
                    </el-descriptions-item>
                    <el-descriptions-item label="应用标签">
                        {{ detailFrom.tags }}
                    </el-descriptions-item>
                    <el-descriptions-item label="介绍视频">
                        {{ detailFrom.video }}
                    </el-descriptions-item>
                    <el-descriptions-item label="状态">
                        {{ detailFrom.status }}
                    </el-descriptions-item>
                    <el-descriptions-item label="接口地址">
                        {{ detailFrom.api_host+detailFrom.api_path }}
                    </el-descriptions-item>

              <el-descriptions-item label="验证api可用性">
                <a :href='detailFrom.api_host+detailFrom.api_path+"/ping"' target="_blank">{{ detailFrom.api_host+detailFrom.api_path+"/ping" }}</a>
              </el-descriptions-item>

              <el-descriptions-item label="查看api文档">
                <a :href='detailFrom.api_host+detailFrom.api_path+"/_docs_info_text"' target="_blank">{{ detailFrom.api_host+detailFrom.api_path+"/_docs_info_text" }}</a>
              </el-descriptions-item>

                    <el-descriptions-item label="是否公开">
                        {{ detailFrom.is_public }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>


    <el-dialog v-model="deployRecordDialogTableVisible" title="发布历史">
      <el-table :data="deployRecord">

        <!--        <el-table-column property="CreatedAt" width="250" label="发布时间" />-->

        <el-table-column property="CreatedAt" width="200" label="发布时间">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column property="appName" label="应用名称" />
        <el-table-column property="appCode" label="工具标识" />

        <el-table-column property="remark" label="变更日志" />
        <el-table-column property="version" label="版本" />
        <el-table-column property="operateUser" label="操作人" />
        <el-table-column align="left" label="操作" fixed="right">
          <template #default="scope">
            <el-button  type="primary" link @click="rollbackVersionFn(scope.row)">回滚</el-button>
            <!--            <el-button  type="primary" link @click="deployRecordFn(scope.row)">删除</el-button>-->
          </template>
        </el-table-column>

      </el-table>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  createBizToolCmdSrvApi,
  deleteBizToolCmdSrvApi,
  deleteBizToolCmdSrvApiByIds,
  updateBizToolCmdSrvApi,
  findBizToolCmdSrvApi,
  getBizToolCmdSrvApiList
} from '@/api/biz_apphub/biz_tool_cmd_srv_api'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import UploadQiNiu from "@/components/upload_oss/UploadQiNiu.vue";
import {
  getDeployList,
  getToolCmdSrvApiUpdateVersionList,
  rollbackVersionToolCmdSrvApiUpdate
} from "@/api/biz_apphub/biz_apphub";
// 上传的文件
const uploadedFiles=ref([])
defineOptions({
    name: 'BizToolCmdSrvApi'
})
const activeNames = ref(['1'])
const handleChange = (val) => {
  console.log(val)
}
// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const bool_statusOptions = ref([])
const price_modeOptions = ref([])
const tool_type = ref([])
const dev_modeOptions = ref([])
const formData = ref({
            appName: '',
            appCode: '',
            tool_type: '',
            title: '',
            desc: '',
            classify: '',
            version: '',
            remark: '',
            mode: '',
            developMode: '',
            cover: '',
            tags: '',
            video: '',
            is_public: '',
            ossPath:''
        })



// 验证规则
const rule = reactive({
               appName : [{
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
               appCode : [{
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
               tool_type : [{
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
               desc : [{
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
               version : [{
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
               mode : [{
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
               developMode : [{
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
            version: 'version',
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
  const table = await getBizToolCmdSrvApiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    price_modeOptions.value = await getDictFunc('price_mode')
    dev_modeOptions.value = await getDictFunc('dev_mode')
    tool_type.value = await getDictFunc('tool_type')
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
            deleteBizToolCmdSrvApiFunc(row)
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
      const res = await deleteBizToolCmdSrvApiByIds({ IDs })
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
const updateBizToolCmdSrvApiFunc = async(row) => {
  uploadedFiles.value=[]
    const res = await findBizToolCmdSrvApi({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}



// 更新版本
const updateBizToolCmdSrvApiVersionFunc = async(row) => {
  uploadedFiles.value=[]
  const res = await findBizToolCmdSrvApi({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVersionVisible.value = true
  }
}



// 删除行
const deleteBizToolCmdSrvApiFunc = async (row) => {
    const res = await deleteBizToolCmdSrvApi({ ID: row.ID })
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
const dialogFormVersionVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    dialogFormVersionVisible.value = false
    formData.value = {
        appName: '',
        appCode: '',
        tool_type: '',
        title: '',
        desc: '',
        classify: '',
        version: '',
        mode: '',
        developMode: '',
        cover: '',
        tags: '',
        video: '',
        is_public: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
               if (uploadedFiles.value.length>0){
                 formData.value.ossPath=uploadedFiles.value[0]
               }
              switch (type.value) {
                case 'create':
                  res = await createBizToolCmdSrvApi(formData.value)
                  break
                case 'update':
                  res = await updateBizToolCmdSrvApi(formData.value)
                  break
                default:
                  res = await createBizToolCmdSrvApi(formData.value)
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
  const res = await findBizToolCmdSrvApi({ ID: row.ID })
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
const deployRecordDialogTableVisible=ref(false)
const deployRecord=ref([])
const deployRecordFn = async (row) => {
  deployRecordDialogTableVisible.value=true
  let res =await getToolCmdSrvApiUpdateVersionList({appId:row.ID})
  console.log("res.data.list---->",res.data.list)
  if (res.code===0){
    deployRecord.value=res.data.list
  }
}

const rollbackVersionFn = async (row) => {
  ElMessage({
    type: 'warning',
    message: '回滚功能正在开发中！！'
  })
  // deployRecordDialogTableVisible.value=true
  // let res =await rollbackVersionToolCmdSrvApiUpdate({appId:row.appId+"",recordId:row.ID+""})
  // if (res.code===0){
  //   ElMessage({
  //     type: 'success',
  //     message: '回滚成功'
  //   })
  //   deployRecordDialogTableVisible.value=false
  //   getTableData()
  //   // deployRecord.value=res.data.list
  // }
  // console.log("res.data.list---->",res.data)
}
</script>

<style>

</style>
