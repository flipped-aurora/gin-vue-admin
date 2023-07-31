<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          />
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
              >删除</el-button
            >
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="名称" prop="name" width="120" />
        <el-table-column align="left" label="拍摄时间" width="180">
          <template #default="scope">{{
            formatDate(scope.row.photographyCreatetime)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="上传人"
          prop="uploadBy"
          width="120"
        />
        <el-table-column align="left" label="类型" prop="type" width="120">
          <template #default="scope">
            {{ setTableType(scope.row.type) }}
          </template>
        </el-table-column>
        <el-table-column label="航拍文件" width="200">
          <template #default="scope">
            <div class="file-list">
              <el-tag
                v-for="file in scope.row.aerialPhotographyFile"
                :key="file.uid"
                >{{ file.name }}</el-tag
              >
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">
            {{ setTableStatus(scope.row.status) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="是否加载"
          prop="loadOrNot"
          width="120"
        >
          <template #default="scope">
            {{ setTableLoadOrNot(scope.row.loadOrNot) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="机巢"
          prop="nestIds"
          width="200"
          show-overflow-tooltip
        >
          <template #default="scope">
            {{ findNestInfoByNestId(scope.row.nestIds) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateAerialPhotographyResultFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close
    >
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="90px"
      >
        <el-form-item label="名称:" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="拍摄时间:" prop="photographyCreatetime">
          <el-date-picker
            v-model="formData.photographyCreatetime"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="上传人:" prop="uploadBy">
          <el-input
            v-model="formData.uploadBy"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="类型:" prop="type">
          <el-select
            v-model="formData.type"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in formDataTypeList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="航拍文件:" prop="aerialPhotographyFile">
          <SelectFile v-model="formData.aerialPhotographyFile" />
        </el-form-item>
        <el-form-item label="状态:" prop="status" style="display: none">
          <el-input
            v-model.number="formData.status"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="坐标:" prop="position" style="display: none">
          <el-input
            v-model="formData.position"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="是否加载:" prop="loadOrNot">
          <el-select
            v-model="formData.loadOrNot"
            class="m-2"
            placeholder="Select"
            size="small"
          >
            <el-option
              v-for="item in formDataLoadOrNotList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="机巢:" prop="nestIds">
          <el-select
            v-model="formData.nestIds"
            multiple
            collapse-tags
            collapse-tags-tooltip
            clearable
            placeholder="Select"
          >
            <el-option
              v-for="item in nestinfoAll"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "AerialPhotographyResult",
};
</script>

<script setup>
import {
  createAerialPhotographyResult,
  deleteAerialPhotographyResult,
  deleteAerialPhotographyResultByIds,
  updateAerialPhotographyResult,
  findAerialPhotographyResult,
  getAerialPhotographyResultList,
  findNestInfoByNestIds,
} from "@/api/aerialPhotographyResult";
import { getUrl } from "@/utils/image";
// 文件选择组件
import SelectFile from "@/components/selectFile/selectFile.vue";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";
import { async } from "ezuikit-js";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: "",
  photographyCreatetime: new Date(),
  uploadBy: "",
  type: 0,
  status: 0,
  position: "",
  aerialPhotographyFile: [],
  loadOrNot: 0,
  nestIds: "",
});
const formDataTypeList = [
  { value: 0, label: "高清正射" },
  { value: 1, label: "三维模型" },
];
const formDataLoadOrNotList = [
  { value: 0, label: "加载" },
  { value: 1, label: "不加载" },
];
// 验证规则
const rule = reactive({
  aerialPhotographyFile: [
    { required: true, message: "请上传航拍文件", trigger: "blur" },
  ],
  nestIds: [{ required: true, message: "请选择机巢", trigger: "blur" }],
});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

let nestinfoAll = [];
const getNestInfoByNestIds = async () => {
  nestinfoAll = await findNestInfoByNestIds();
  nestinfoAll = nestinfoAll.data.renestinfoList;
  nestinfoAll = nestinfoAll.map((item) => ({
    label: item.nestName,
    value: item.nestid,
  }));
};
getNestInfoByNestIds();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getAerialPhotographyResultList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// 根据机槽id显示机槽名称
const findNestInfoByNestId = (ids) => {
  const idlist = ids.split(",");
  return idlist
    .map((item) => {
      console.log(nestinfoAll.find((val) => val.value === item));
      return nestinfoAll.find((val) => val.value === item)
        ? nestinfoAll.find((val) => val.value === item).label
        : "";
    })
    .join(",");
};
const setTableLoadOrNot = (loadOrNot) => {
  if (loadOrNot) {
    return "不加载";
  } else {
    return "加载";
  }
};
const setTableType = (type) => {
  if (type) {
    return "三维模型";
  } else {
    return "高清正射";
  }
};
const setTableStatus = (status) => {
  switch (status) {
    case 0:
      return "上传中";
    case 1:
      return "解压中";
    case 2:
      return "已完成";
    case 3:
      return "异常";
  }
};
// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteAerialPhotographyResultFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deleteAerialPhotographyResultByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateAerialPhotographyResultFunc = async (row) => {
  const res = await findAerialPhotographyResult({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.reALPhotographyResult;
    formData.value.nestIds = formData.value.nestIds.split(",");

    console.log(formData);
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteAerialPhotographyResultFunc = async (row) => {
  const res = await deleteAerialPhotographyResult({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    name: "",
    photographyCreatetime: new Date(),
    uploadBy: "",
    type: 0,
    status: 0,
    position: "",
    loadOrNot: 0,
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    formData.value.nestIds = formData.value.nestIds.join(",");
    switch (type.value) {
      case "create":
        res = await createAerialPhotographyResult(formData.value);
        break;
      case "update":
        res = await updateAerialPhotographyResult(formData.value);
        break;
      default:
        res = await createAerialPhotographyResult(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};

const downloadFile = (url) => {
  window.open(getUrl(url), "_blank");
};
</script>

<style>
.file-list {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}
</style>
