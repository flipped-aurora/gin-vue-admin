# View 示例

## 这个文件负责什么

页面组件负责查询表单、表格、弹窗、抽屉和交互流程，是用户真正接触到的界面层。

## 什么时候应该这样写

- 新增后台管理页面
- 新增列表页 + 搜索 + 分页
- 新增表单弹窗或抽屉流程

## 推荐写法示例

```vue
<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">查询</el-button>
          <el-button @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <el-table :data="tableData" row-key="ID">
        <el-table-column label="ID" prop="ID" width="80" />
        <el-table-column label="名称" prop="name" />
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { getOrderList } from '@/api/order'

const searchInfo = ref({})
const tableData = ref([])

const getTableData = async () => {
  const res = await getOrderList(searchInfo.value)
  if (res.code === 0) {
    tableData.value = res.data.list
  }
}

const onSubmit = () => {
  getTableData()
}

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

getTableData()
</script>
```

## 为什么这样写

- 查询区和表格区结构清晰，符合项目后台页面习惯
- `script setup` 下把“状态、请求、交互入口”放在一起，易读
- 页面只处理展示和交互，不在这里重写公共请求逻辑

## 常见错误

- 页面里直接写大量请求封装逻辑
- 组件过大，不拆查询区、表格区、弹窗区
- 页面状态命名混乱，不区分 `searchInfo`、`tableData`、`form`

## 真实参考文件

- `web/src/view/systemTools/apiToken/index.vue`
- `web/src/view/superAdmin/api/api.vue`
