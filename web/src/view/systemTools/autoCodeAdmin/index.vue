<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="goAutoCode(null)">
          {{ t('general.add') }}
        </el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="ID" width="60" prop="ID" />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoCodeAdmin.structName')"
          min-width="150"
          prop="structName"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoCodeAdmin.structDesc')"
          min-width="150"
          prop="description"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.exportTemplate.tableName')"
          min-width="150"
          prop="tableName"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoCodeAdmin.rollBackMark')"
          min-width="150"
          prop="flag"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.flag" type="danger" effect="dark">
              {{ t("view.systemTools.autoCodeAdmin.rolledBack") }}
            </el-tag>
            <el-tag v-else type="success" effect="dark">
              {{ t("view.systemTools.autoCodeAdmin.notRolledBack") }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.createdAt')" width="180">
          <template #default="scope">
            {{
              formatDate(scope.row.CreatedAt)
            }}
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.operations')" min-width="240">
          <template #default="scope">
            <div>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="addFuncBtn(scope.row)"
              >
              {{ t('view.systemTools.autoPkg.addMethod') }}
              </el-button>
              <el-button type="primary" link @click="goAutoCode(scope.row,1)">
                {{ t('view.systemTools.autoCodeAdmin.addField') }}
              </el-button>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="openDialog(scope.row)"
              >
              {{ t("view.systemTools.autoCodeAdmin.rollBack") }}
              </el-button>
              <el-button type="primary" link @click="goAutoCode(scope.row)">
                {{ t("view.systemTools.autoCodeAdmin.reuse") }}
              </el-button>
              <el-button type="primary" link @click="deleteRow(scope.row)">
                {{ t("general.delete") }}
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :title="dialogFormTitle"
      :before-close="closeDialog"
      width="600px"
    >
      <el-form :inline="true" :model="formData" label-width="80px">
        <el-form-item :label="t('view.systemTools.autoPkg.options')">
          <el-checkbox
            v-model="formData.deleteApi"
            :label="t('view.systemTools.autoPkg.deleteInterface')"
          />
          <el-checkbox
            v-model="formData.deleteMenu"
            :label="t('view.systemTools.autoPkg.deleteMenu')"
          />
          <el-checkbox
            v-model="formData.deleteTable"
            :label="t('view.systemTools.autoPkg.deleteTable')"
            @change="deleteTableCheck"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">
            {{t('general.close')}}
          </el-button>
          <el-popconfirm
            :title="t('view.systemTools.autoPkg.rollbackOperationNote')"
            @confirm="enterDialog"
          >
            <template #reference>
              <el-button type="primary">
                {{t('general.confirm')}}
              </el-button>
            </template>
          </el-popconfirm>
        </div>
      </template>
    </el-dialog>


    <el-drawer
      v-model="funcFlag"
      size="60%"
      :show-close="false"
      :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.systemTools.autoCode.actionBar') }}</span>
          <div>
            <el-button
              type="primary"
              @click="runFunc"
              :loading="aiLoading"
            >
            {{ t('view.systemTools.autoPkg.generate') }}
            </el-button>
            <el-button
              type="primary"
              @click="closeFunc"
              :loading="aiLoading"
            >
            {{ t('general.cancel') }}
            </el-button>
          </div>
        </div>
      </template>
      <div class="">
        <el-form v-loading="aiLoading" label-position="top" :element-loading-text="t('view.systemTools.autoCodeAdmin.xiaoMiaoIsThinking')" :model="autoFunc" label-width="80px">
          <el-form-item :label="t('view.systemTools.autoPkg.packageName')">
            <el-input v-model="autoFunc.package" :placeholder="t('view.systemTools.autoPkg.enterPackageNameNote')" disabled />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.structName')">
            <el-input v-model="autoFunc.structName" :placeholder="t('view.systemTools.autoCode.entStructName')" disabled />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.frontendFileName')">
            <el-input v-model="autoFunc.packageName" :placeholder="t('view.systemTools.autoCode.fineNameInput')" disabled />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.backendFileName')">
            <el-input v-model="autoFunc.humpPackageName" :placeholder="t('view.systemTools.autoCode.fineNameInput')" disabled />
          </el-form-item>
          <el-form-item :label="t('general.description')">
            <el-input v-model="autoFunc.description" :placeholder="t('view.systemTools.dictionary.sysDictionary.enterDescription')" disabled />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.abbreviation')">
            <el-input v-model="autoFunc.abbreviation" :placeholder="t('view.systemTools.enterAbbreviation')" disabled />
          </el-form-item>

          
          <el-form-item label="是否AI填充：">
            <el-switch v-model="autoFunc.isAi" /> <span class="text-sm text-red-600 p-2">{{ t('view.systemTools.autoCodeAdmin.aiWritingNote') }}</span>
          </el-form-item>
          <template v-if="autoFunc.isAi">
            <el-form-item label="Ai帮写:">
              <div class="relative w-full">
                <el-input type="textarea" placeholder="AI帮写功能，输入提示信息，自动生成代码" v-model="autoFunc.prompt" :rows="5" @input="autoFunc.router = autoFunc.router.replace(/\//g, '')" />
                <el-button @click="aiAddFunc" type="primary" class="absolute right-2 bottom-2"><ai-gva />帮写</el-button>
              </div>
            </el-form-item>
            <el-form-item label="Api方法:">
              <codemirror
                v-model="autoFunc.apiFunc"
                placeholder="Code goes here..."
                :style="{ height: '300px',width:'100%' }"
                :indent-with-tab="true"
                :tab-size="2"
                :extensions=" [go(), oneDark]"
              />
            </el-form-item>
            <el-form-item label="Server方法:">
              <codemirror
                v-model="autoFunc.serverFunc"
                placeholder="Code goes here..."
                :style="{ height: '300px',width:'100%' }"
                :indent-with-tab="true"
                :tab-size="2"
                :extensions=" [go(), oneDark]"
              />
            </el-form-item>
            <el-form-item label="前端JSAPI方法:">
              <codemirror
                v-model="autoFunc.jsFunc"
                placeholder="Code goes here..."
                :style="{ height: '300px',width:'100%' }"
                :indent-with-tab="true"
                :tab-size="2"
                :extensions=" [javascript(), oneDark]"
              />
            </el-form-item>
          </template>

          <el-form-item label="方法介绍：">
            <div class="flex w-full gap-2">
              <el-input class="flex-1" v-model="autoFunc.funcDesc" placeholder="请输入方法介绍" />
              <el-button type="primary" @click="autoComplete"><ai-gva />补全</el-button>
            </div>
          </el-form-item>
          <el-form-item label="方法名：">
            <el-input @blur="autoFunc.funcName = toUpperCase(autoFunc.funcName)" v-model="autoFunc.funcName" placeholder="请输入方法名" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.method')">
            <el-select v-model="autoFunc.method" :placeholder="t('view.systemTools.selectMethod')">
              <el-option
                v-for="item in ['GET', 'POST', 'PUT', 'DELETE']"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="是否鉴权：">
            <el-switch v-model="autoFunc.isAuth" active-text="是" inactive-text="否" />
          </el-form-item>
          <el-form-item :label="t('menu.routePath')">
            <el-input v-model="autoFunc.router" :placeholder="t('menu.routePath')" @input="autoFunc.router = autoFunc.router.replace(/\//g, '')" />
            <div>{{t('view.api.apiPath')}}: [{{ autoFunc.method }}]  /{{ autoFunc.abbreviation }}/{{ autoFunc.router }}</div>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { getSysHistory, rollback, delSysHistory,addFunc,butler } from "@/api/autoCode.js";
import { useRouter } from "vue-router";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref,onMounted } from "vue";
import { formatDate } from "@/utils/format";
import { toUpperCase } from "@/utils/stringFun"
import  {useAppStore} from "@/pinia";
import { Marked } from "marked";
import { markedHighlight } from "marked-highlight";
import hljs from 'highlight.js'

import { Codemirror } from 'vue-codemirror'
  import { javascript } from '@codemirror/lang-javascript'
  import { go } from '@codemirror/lang-go'
  import { oneDark } from '@codemirror/theme-one-dark'
  import { useI18n } from 'vue-i18n' // added by mohamed hassan to support multilingual

const { t } = useI18n() // added by mohamed hassan to support multilingual


const appStore = useAppStore()

defineOptions({
  name: "AutoCodeAdmin",
});

const aiLoading = ref(false)

const formData = ref({
  id: undefined,
  deleteApi: true,
  deleteMenu: true,
  deleteTable: false,
});

const router = useRouter();
const dialogFormVisible = ref(false);
const dialogFormTitle = ref("");

const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);

const activeInfo = ref("")

const autoFunc = ref({
  package:"",
  funcName:"",
  structName:"",
  packageName:"",
  description:"",
  abbreviation:"",
  humpPackageName:"",
  businessDB:"",
  method:"",
  funcDesc: "",
  isAuth:false,
  isAi:false,
  apiFunc:"",
  serverFunc:"",
  jsFunc:"",
})

const addFuncBtn =  (row) => {
  const req = JSON.parse(row.request)
  activeInfo.value = row.request
  autoFunc.value.package = req.package
  autoFunc.value.structName = req.structName
  autoFunc.value.packageName = req.packageName
  autoFunc.value.description = req.description
  autoFunc.value.abbreviation = req.abbreviation
  autoFunc.value.humpPackageName = req.humpPackageName
  autoFunc.value.businessDB = req.businessDB
  autoFunc.value.method = ""
  autoFunc.value.funcName = ""
  autoFunc.value.router = ""
  autoFunc.value.funcDesc = ""
  autoFunc.value.isAuth = false
  autoFunc.value.isAi = false
  autoFunc.value.apiFunc = ""
  autoFunc.value.serverFunc = ""
  autoFunc.value.jsFunc = ""
  funcFlag.value = true;
};

const funcFlag = ref(false);

const closeFunc = () => {
  funcFlag.value = false;
};

const runFunc = async () =>{
  // 首字母自动转换为大写
  autoFunc.value.funcName = toUpperCase(autoFunc.value.funcName)

  if (!autoFunc.value.funcName) {
    ElMessage.error("请输入方法名")
    return
  }
  if (!autoFunc.value.method) {
    ElMessage.error("请选择方法")
    return
  }
  if (!autoFunc.value.router) {
    ElMessage.error("请输入路由")
    return
  }
  if (!autoFunc.value.funcDesc) {
    ElMessage.error("请输入方法介绍")
    return
  }

  if (autoFunc.value.isAi){
    if (!autoFunc.value.apiFunc || !autoFunc.value.serverFunc || !autoFunc.value.jsFunc) {
      ElMessage.error("请先使用AI帮写完成基础代码，如果生成失败请重新调用")
      return
    }
  }

  const res = await addFunc(autoFunc.value)
  if (res.code === 0) {
    ElMessage.success(t('view.systemTools.autoPkg.addMethodSuccess'));
    closeFunc()
  }
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getSysHistory({
    page: page.value,
    pageSize: pageSize.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

const deleteRow = async(row) => {
  ElMessageBox.confirm(t('autoCodeAdmin.deleteHistoryConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  }).then(async() => {
    const res = await delSysHistory({ id: Number(row.ID) })
    if (res.code === 0) {
      ElMessage.success(t('general.deleteSuccess'))
      getTableData()
    }
  });
};

// 打开弹窗
const openDialog = (row) => {
  dialogFormTitle.value = t('autoCodeAdmin.rollBack') + ' ' + row.structName;
  formData.value.id = row.ID;
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    id: undefined,
    deleteApi: true,
    deleteMenu: true,
    deleteTable: false,
  };
};

// 确认删除表
const deleteTableCheck = (flag) => {
  if (flag) {
    ElMessageBox.confirm(
      t('view.systemTools.autoPkg.deleteFilesNote'),
      t('general.hint'),
      {
        closeOnClickModal: false,
        distinguishCancelAndClose: true,
        confirmButtonText: t('general.confirm'),
        cancelButtonText: t('general.cancel'),
        type: "warning",
      }
    )
      .then(() => {
        ElMessageBox.confirm(
          t('view.systemTools.autoPkg.deleteFilesConfirmation'),
          t('view.systemTools.autoPkg.willDeleteTable'),
          {
            closeOnClickModal: false,
            distinguishCancelAndClose: true,
            confirmButtonText: t('general.confirm'),
            cancelButtonText: t('general.cancel'),
            type: "warning",
          }
        ).catch(() => {
          formData.value.deleteTable = false;
        });
      })
      .catch(() => {
        formData.value.deleteTable = false;
      });
  }
};

const enterDialog = async () => {
  const res = await rollback(formData.value);
  if (res.code === 0) {
    ElMessage.success(t('autoCodeAdmin.rollbackSuccess'));
    getTableData();
  }
};

const goAutoCode = (row,isAdd) => {
  if (row) {
    router.push({
      name: "autoCodeEdit",
      params: {
        id: row.ID,
      },
      query: {
        isAdd: isAdd
      },
    });
  } else {
    router.push({ name: "autoCode" });
  }
};


const aiAddFunc = async () =>{
  aiLoading.value = true
  autoFunc.value.apiFunc = ""
  autoFunc.value.serverFunc = ""
  autoFunc.value.jsFunc = ""

  if (!autoFunc.value.prompt) {
    ElMessage.error("请输入提示信息")
    return
  }

  const res = await addFunc({...autoFunc.value,isPreview:true})
  if (res.code !== 0) {
    aiLoading.value = false
    ElMessage.error(res.msg)
    return
  }

 const aiRes = await butler({
    structInfo:activeInfo.value,
    template:JSON.stringify(res.data),
    prompt: autoFunc.value.prompt,
    command: "addFunc"
  })
  aiLoading.value = false
  if (aiRes.code === 0) {
    try{
      const aiData = JSON.parse(aiRes.data)
      autoFunc.value.apiFunc = aiData.api
      autoFunc.value.serverFunc = aiData.server
      autoFunc.value.jsFunc = aiData.js
      autoFunc.value.method = aiData.method
      autoFunc.value.funcName = aiData.funcName
      const routerArr = aiData.router.split("/")
      autoFunc.value.router = routerArr[routerArr.length - 1]
      autoFunc.value.funcDesc = autoFunc.value.prompt
    } catch (e) {
      ElMessage.error("小淼忙碌，请重新调用")
    }
  }
}

const autoComplete = async () =>{
  aiLoading.value = true
  const aiRes = await butler({
    prompt: autoFunc.value.funcDesc,
    command: "autoCompleteFunc"
  })
  aiLoading.value = false
  if (aiRes.code === 0) {
    try{
      const aiData = JSON.parse(aiRes.data)
      autoFunc.value.method = aiData.method
      autoFunc.value.funcName = aiData.funcName
      autoFunc.value.router = aiData.router
      autoFunc.value.prompt = autoFunc.value.funcDesc
    } catch (e) {
      ElMessage.error("小淼开小差了，请重新调用")
    }
  }
}



</script>