<template>
  <el-form
    :model="formData"
    ref="elFormRef"
    :rules="rules"
    label-position="top"
    label-width="80px"
    size="medium"
    @submit.prevent
  >
    <el-row>
      <el-col :span="24" class="grid-cell">
        <el-button @click="save" type="primary">保存</el-button>
        <el-button @click="back" type="danger">返回</el-button>
      </el-col>
    </el-row>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-row>
      <el-col :span="24" class="grid-cell">
        <el-form-item label="角色" prop="roleid">
          <el-select
            v-model="formData.roleid"
            class="full-width-input"
            clearable
            filterable
            remote
            :remote-method="roleSearch"
            :loading="loading"
          >
            <el-option
              v-for="(item, index) in roleidOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="24" class="grid-cell">
        <el-form-item label="机巢" prop="nestid">
          <el-select
            v-model="formData.nestid"
            class="full-width-input"
            clearable
            filterable
            multiple
            remote
            :remote-method="nestSearch"
            :loading="nestLoading"
          >
            <el-option
              v-for="(item, index) in nestidOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            ></el-option>
          </el-select>
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>

<script>
export default {
  name: "NestRole",
};
</script>

<script setup>
import { createNestRole, updateNestRole, findNestRole } from "@/api/nestRole";
import { getAuthorityList } from "@/api/authority";
import { getNestInfoList } from "@/api/nestInfo";

// 自动获取字典
import { getDictFunc } from "@/utils/format";
import { useRoute, useRouter } from "vue-router";
import { ElMessage } from "element-plus";
import { ref, reactive } from "vue";
const route = useRoute();
const router = useRouter();

const emit = defineEmits(['cal'])

const type = ref("");
const formData = ref({
  roleid: "",
  nestid: "",
});
// 验证规则
const rule = reactive({
  roleid: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
  nestid: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
    {
      whitespace: true,
      message: "不能只输入空格",
      trigger: ["input", "blur"],
    },
  ],
});

const elFormRef = ref();
const loading = ref(false);
const roleidOptions = ref([]);
const allRoleidOptions = ref([]);

const nestLoading = ref(false);
const nestidOptions = ref([]);
const allNestidOptions = ref([]);

const roleSearch = (query) => {
  loading.value = true;
  const tmpArr = allRoleidOptions.value.filter((item) => {
    return item.label.indexOf(query) !== -1;
  });
  roleidOptions.value = tmpArr;
  loading.value = false;
};
const nestSearch = (query) => {
  loading.value = true;
  const tmpArr = allRoleidOptions.value.filter((item) => {
    return item.label.indexOf(query) !== -1;
  });
  roleidOptions.value = tmpArr;
  loading.value = false;
};

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findNestRole({ ID: route.query.id });
    if (res.code === 0) {
      formData.value = res.data.renestrole;
      type.value = "update";
    }
  } else {
    type.value = "create";
  }
  getAuthorityList({ page: 1, pageSize: 9999 }).then((res) => {
    const data = res.data;
    for (const item of data.list) {
      allRoleidOptions.value.push({
        value: item.authorityId,
        label: item.authorityName,
      });
    }
    roleidOptions.value = allRoleidOptions.value;
    // console.log("roleid list: ",roleidOptions)
  });

  getNestInfoList({ page: 1, pageSize: 9999 }).then((res) => {
    const data = res.data;
    for (const item of data.list) {
      allNestidOptions.value.push({
        value: `${item.nestid}`,
        label: item.nestName,
      });
    }
    nestidOptions.value = allNestidOptions.value;
  });
};

init();

// const confirmDialog = () => {
//   emit("cd", formData.value)
// }
// const cancelDialog = () => {
//   emit("cal")
// }

// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) {
      return;
    }
    let res;
    formData.value.nestid = JSON.stringify(formData.value.nestid)
    console.log('val', formData.value)
    switch (type.value) {
      case "create":
        res = await createNestRole(formData.value);
        break;
      case "update":
        res = await updateNestRole(formData.value);
        break;
      default:
        res = await createNestRole(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
    }
    emit('cal')
  });
};

// 返回按钮
const back = () => {
  // router.go(-1);
  emit('cal')
};
</script>

<style lang="scss" scoped>
.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.full-width-input {
  width: 100%;
}

.el-form-item--medium {
  .el-radio {
    line-height: 36px !important;
  }

  .el-rate {
    margin-top: 8px;
  }
}

.el-form-item--small {
  .el-radio {
    line-height: 32px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--mini {
  .el-radio {
    line-height: 28px !important;
  }

  .el-rate {
    margin-top: 4px;
  }
}

.clear-fix:before,
.clear-fix:after {
  display: table;
  content: "";
}

.clear-fix:after {
  clear: both;
}

.float-right {
  float: right;
}
</style>

<style lang="scss" scoped>
div.table-container {
  table.table-layout {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;

    td.table-cell {
      display: table-cell;
      height: 36px;
      border: 1px solid #e1e2e3;
    }
  }
}

div.tab-container {
}

.label-left-align :deep(.el-form-item__label) {
  text-align: left;
}

.label-center-align :deep(.el-form-item__label) {
  text-align: center;
}

.label-right-align :deep(.el-form-item__label) {
  text-align: right;
}

.custom-label {
}

.static-content-item {
  min-height: 20px;
  display: flex;
  align-items: center;

  :deep(.el-divider--horizontal) {
    margin: 0;
  }
}
</style>
