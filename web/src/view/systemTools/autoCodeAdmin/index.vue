<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="goAutoCode(null)">
          新增
        </el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="id" width="60" prop="ID" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">
            {{
              formatDate(scope.row.CreatedAt)
            }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="结构体名"
          min-width="150"
          prop="structName"
        />
        <el-table-column
          align="left"
          label="结构体描述"
          min-width="150"
          prop="description"
        />
        <el-table-column
          align="left"
          label="表名称"
          min-width="150"
          prop="tableName"
        />
        <el-table-column
          align="left"
          label="回滚标记"
          min-width="150"
          prop="flag"
        >
          <template #default="scope">
            <el-tag v-if="scope.row.flag" type="danger" effect="dark">
              已回滚
            </el-tag>
            <el-tag v-else type="success" effect="dark">
              未回滚
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="240">
          <template #default="scope">
            <div>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="addFuncBtn(scope.row)"
              >
                增加方法
              </el-button>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="openDialog(scope.row)"
              >
                回滚
              </el-button>
              <el-button type="primary" link @click="goAutoCode(scope.row)">
                复用
              </el-button>
              <el-button type="primary" link @click="deleteRow(scope.row)">
                删除
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
        <el-form-item label="选项：">
          <el-checkbox
            v-model="formData.deleteApi"
            label="删除接口"
          />
          <el-checkbox
            v-model="formData.deleteMenu"
            label="删除菜单"
          />
          <el-checkbox
            v-model="formData.deleteTable"
            label="删除表"
            @change="deleteTableCheck"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">
            取 消
          </el-button>
          <el-popconfirm
            title="此操作将回滚生成文件和勾选项目, 是否继续?"
            @confirm="enterDialog"
          >
            <template #reference>
              <el-button type="primary">
                确 定
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
          <span class="text-lg">操作栏</span>
          <div>
            <el-button
              type="primary"
              @click="runFunc"
            >
              生成
            </el-button>
            <el-button
              type="primary"
              @click="closeFunc"
            >
              取消
            </el-button>
          </div>
        </div>
      </template>
      <div class="">
        <el-form label-position="top" :model="autoFunc" label-width="80px">
          <el-form-item label="包名：">
            <el-input v-model="autoFunc.package" placeholder="请输入包名" disabled />
          </el-form-item>
          <el-form-item label="结构体名：">
            <el-input v-model="autoFunc.structName" placeholder="请输入结构体名" disabled />
          </el-form-item>
          <el-form-item label="前端文件名：">
            <el-input v-model="autoFunc.packageName" placeholder="请输入文件名" disabled />
          </el-form-item>
          <el-form-item label="后端文件名：">
            <el-input v-model="autoFunc.humpPackageName" placeholder="请输入文件名" disabled />
          </el-form-item>
          <el-form-item label="描述：">
            <el-input v-model="autoFunc.description" placeholder="请输入描述" disabled />
          </el-form-item>
          <el-form-item label="缩写：">
            <el-input v-model="autoFunc.abbreviation" placeholder="请输入缩写" disabled />
          </el-form-item>
          <el-form-item label="方法介绍：">
            <el-input v-model="autoFunc.funcDesc" placeholder="请输入方法介绍" />
          </el-form-item>
          <el-form-item label="方法名：">
            <el-input v-model="autoFunc.funcName" placeholder="请输入方法名" />
          </el-form-item>
          <el-form-item label="方法：">
            <el-select v-model="autoFunc.method" placeholder="请选择方法">
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
          <el-form-item label="路由path:">
            <el-input v-model="autoFunc.router" placeholder="路由path" @input="autoFunc.router = autoFunc.router.replace(/\//g, '')" />
            <div>API路径: [{{ autoFunc.method }}]  /{{ autoFunc.abbreviation }}/{{ autoFunc.router }}</div>
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
  funcDesc: "",
  isAuth:false,
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
  autoFunc.value.funcDesc = "方法介绍"
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
    ElMessage.success("增加方法成功");
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

const deleteRow = async (row) => {
  ElMessageBox.confirm("此操作将删除本历史, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const res = await delSysHistory({ id: Number(row.ID) });
    if (res.code === 0) {
      ElMessage.success("删除成功");
      getTableData();
    }
  });
};

// 打开弹窗
const openDialog = (row) => {
  dialogFormTitle.value = "回滚：" + row.structName;
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
      `此操作将删除自动创建的文件和api（会删除表！！！）, 是否继续?`,
      "提示",
      {
        closeOnClickModal: false,
        distinguishCancelAndClose: true,
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }
    )
      .then(() => {
        ElMessageBox.confirm(
          `此操作将删除自动创建的文件和api（会删除表！！！）, 请继续确认！！！`,
          "会删除表",
          {
            closeOnClickModal: false,
            distinguishCancelAndClose: true,
            confirmButtonText: "确定",
            cancelButtonText: "取消",
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
    ElMessage.success("回滚成功");
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
