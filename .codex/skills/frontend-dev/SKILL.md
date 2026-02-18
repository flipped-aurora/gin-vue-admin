---
name: frontend-dev
description: GVA前端开发规范 - 开发Vue页面、组件时加载
allowed-tools: Bash(gh *)
context: fork
agent: Explore
---
# GVA 前端开发规范

## 核心原则

### 1. 模块化架构

**依赖链**: `页面组件 → API服务 → 后端接口`

- **API层** (`src/api/`): 封装所有后端接口调用
- **组件层** (`src/components/`): 可复用UI组件
- **页面层** (`src/view/`): 业务页面
- **状态层** (`src/pinia/`): 全局状态管理

### 2. 统一API调用

所有API调用**必须**通过 `src/api/` 封装，使用 `@/utils/request.js`：

```javascript
import service from '@/utils/request'

/**
 * 获取列表
 * @param {Object} data 查询参数
 * @param {number} data.page 页码
 * @param {number} data.pageSize 每页数量
 * @returns {Promise} 列表数据
 */
export const getXxxList = (data) => {
  return service({
    url: '/xxx/list',
    method: 'post',
    data
  })
}

export const createXxx = (data) => {
  return service({
    url: '/xxx/create',
    method: 'post',
    data
  })
}
```

### 3. Composition API

**必须**使用 Vue 3 Composition API + `<script setup>` 语法：

```vue
<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getXxxList } from '@/api/xxx'

// 响应式数据
const loading = ref(false)
const tableData = ref([])
const searchForm = reactive({
  page: 1,
  pageSize: 10,
  keyword: ''
})

// 获取数据
const fetchData = async () => {
  loading.value = true
  try {
    const res = await getXxxList(searchForm)
    if (res.code === 0) {
      tableData.value = res.data.list
    }
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>
```

## 组件规范

### Props 和 Emits 定义

```vue
<script setup>
/**
 * 通用表格组件
 * @component GvaTable
 */
const props = defineProps({
  data: {
    type: Array,
    required: true,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  },
  columns: {
    type: Array,
    required: true
  }
})

const emit = defineEmits(['refresh', 'edit', 'delete'])

const handleEdit = (row) => {
  emit('edit', row)
}
</script>
```

## Pinia 状态管理

使用 Composition API 风格：

```javascript
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useStorage } from '@vueuse/core'

export const useUserStore = defineStore('user', () => {
  // 状态
  const userInfo = ref({})
  const token = useStorage('token', '')

  // 计算属性
  const isLogin = computed(() => !!token.value)

  // 方法
  const setUserInfo = (val) => {
    userInfo.value = val
  }

  const logout = () => {
    token.value = ''
    userInfo.value = {}
  }

  return {
    userInfo,
    token,
    isLogin,
    setUserInfo,
    logout
  }
})
```

## 样式规范

### UnoCSS 原子化类

```vue
<template>
  <!-- 优先使用UnoCSS原子类 -->
  <div class="flex items-center justify-between p-4 bg-white rounded-lg shadow">
    <span class="text-lg font-bold text-gray-800">标题</span>
    <el-button type="primary">操作</el-button>
  </div>
</template>
```

### 命名规范

| 类型 | 规范 | 示例 |
|-----|------|-----|
| 文件名 | kebab-case | `user-list.vue` |
| 组件名 | PascalCase | `UserList` |
| 变量名 | camelCase | `userList` |
| 常量名 | UPPER_SNAKE | `MAX_COUNT` |

## 页面模板

```vue
<template>
  <div class="gva-container">
    <!-- 搜索区域 -->
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="名称">
          <el-input v-model="searchForm.keyword" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 表格区域 -->
    <div class="gva-table-box">
      <el-table v-loading="loading" :data="tableData">
        <el-table-column prop="name" label="名称" />
        <el-table-column prop="status" label="状态" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">编辑</el-button>
            <el-button type="danger" link @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="searchForm.page"
        v-model:page-size="searchForm.pageSize"
        :total="total"
        @current-change="fetchData"
      />
    </div>
  </div>
</template>
```

## 插件目录结构

```
src/plugin/[插件名]/
├── api/           # API接口
├── components/    # 组件
├── view/          # 页面
└── form/          # 表单
```

