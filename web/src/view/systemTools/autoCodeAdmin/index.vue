<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="goAutoCode(null)">
          {{ t("general.add") }}
        </el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" :label="t('general.createdAt')" width="180">
          <template #default="scope">
            {{
              formatDate(scope.row.CreatedAt)
            }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoCode.structName')"
          min-width="150"
          prop="structName"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoCode.structChineseName')"
          min-width="150"
          prop="structCNName"
        />
        <el-table-column
          align="left"
          :label="t('view.systemTools.autoCode.tableName')"
          min-width="150"
          prop="tableName"
        />
        <el-table-column
          align="left"
          :label="t('autoCodeAdmin.rollBackMark')"
          min-width="150"
          prop="flag"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.flag" type="danger" effect="dark">
              {{ t("autoCodeAdmin.rolledBack") }}
            </el-tag>
            <el-tag v-else type="success" effect="dark">
              {{ t("autoCodeAdmin.notRolledBack") }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" :lable="t('general.operations')" min-width="240">
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
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="openDialog(scope.row)"
              >
              {{ t("autoCodeAdmin.rollBack") }}
              </el-button>
              <el-button type="primary" link @click="goAutoCode(scope.row)">
                {{ t("autoCodeAdmin.reuse") }}
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
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.systemTools.autoCode.actionBar') }}</span>
          <div>
            <el-button
              type="primary"
              @click="runFunc"
            >
              {{ t('view.systemTools.autoPkg.generate') }}
            </el-button>
            <el-button
              type="primary"
              @click="closeFunc"
            >
              {{ t('general.cancel') }}
            </el-button>
          </div>
        </div>
      </template>
      <div class="">
        <el-form label-position="top" :model="autoFunc" label-width="80px">
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
          <el-form-item :label="t('view.systemTools.methodDescription')">
            <el-input v-model="autoFunc.funcDesc" :placeholder="t('view.systemTools.enterMethodDescription')" />
          </el-form-item>
          <el-form-item :label="t('view.systemTools.methodName')">
            <el-input v-model="autoFunc.funcName" :placeholder="t('view.systemTools.enterMethodName')" />
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
          <el-form-item :label="t('menu.routePath')">
            <el-input v-model="autoFunc.router" :placeholder="t('menu.routePath')" />
            <div>{{t('view.api.apiPath')}}: [{{ autoFunc.method }}]  /{{ autoFunc.abbreviation }}/{{ autoFunc.router }}</div>
          </el-form-item>
        </el-form>
      </div>
    </el-drawer>
  </div>
</template>

<script setup>
import { getSysHistory, rollback, delSysHistory,addFunc } from "@/api/autoCode.js";
import { useRouter } from "vue-router";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref } from "vue";
import { formatDate } from "@/utils/format";
import { useI18n } from 'vue-i18n'; // added by mohamed hassan to support multilanguage

const { t } = useI18n() // added by mohamed hassan to support multilanguage
import { toUpperCase } from "@/utils/stringFun"

defineOptions({
  name: "AutoCodeAdmin",
});

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
  funcDesc: ""
})

const addFuncBtn =  (row) => {
  const req = JSON.parse(row.request)
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
  autoFunc.value.funcDesc = t('view.systemTools.enterMethodDescription')
  funcFlag.value = true;
};

const funcFlag = ref(false);

const closeFunc = () => {
  funcFlag.value = false;
};

const runFunc = async () =>{
  // 首字母自动转换为大写
  autoFunc.value.funcName = toUpperCase(autoFunc.value.funcName)

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
  dialogFormTitle.value = t('autoCodeAdmin.rollBack') + row.structName;
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

const goAutoCode = (row) => {
  if (row) {
    router.push({
      name: "autoCodeEdit",
      params: {
        id: row.ID,
      },
    });
  } else {
    router.push({ name: "autoCode" });
  }
};
</script>
