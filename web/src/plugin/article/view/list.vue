<template>
  <div class="article-list p-4">
    <el-card shadow="never" class="mb-4">
      <div class="flex items-center justify-between flex-wrap gap-2">
        <div class="text-base font-semibold">文章列表</div>
        <el-button type="primary" @click="openEdit()">新增文章</el-button>
      </div>
    </el-card>

    <el-card shadow="never" class="mb-4">
      <el-form :inline="true" :model="search" class="flex flex-wrap items-center gap-2">
        <el-form-item label="标题">
          <el-input v-model="search.title" placeholder="搜索标题" clearable @change="load" />
        </el-form-item>
        <el-form-item label="语言">
          <el-select v-model="search.lang" placeholder="全部" clearable style="width: 100px" @change="load">
            <el-option label="中文" value="zh" />
            <el-option label="英文" value="en" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="search.status" placeholder="全部" clearable style="width: 100px" @change="load">
            <el-option label="发布" :value="1" />
            <el-option label="草稿" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="分类">
          <el-select v-model="search.categoryID" placeholder="全部" clearable style="width: 140px" @change="load">
            <el-option v-for="c in categories" :key="c.ID" :label="c.name" :value="c.ID" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="load">查询</el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="never">
      <el-table :data="list" v-loading="loading" stripe>
        <el-table-column prop="ID" label="ID" width="70" />
        <el-table-column prop="title" label="标题" min-width="200" show-overflow-tooltip />
        <el-table-column prop="lang" label="语言" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.lang === 'zh'" type="success" size="small">中文</el-tag>
            <el-tag v-else-if="row.lang === 'en'" type="primary" size="small">英文</el-tag>
            <el-tag v-else size="small">{{ row.lang }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="80">
          <template #default="{ row }">
            <el-tag v-if="row.status === 1" type="success" size="small">发布</el-tag>
            <el-tag v-else type="info" size="small">草稿</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="readCount" label="阅读" width="80" align="right" />
        <el-table-column prop="nickName" label="作者" width="120" />
        <el-table-column prop="CreatedAt" label="创建时间" width="170">
          <template #default="{ row }">{{ formatDate(row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" width="220" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" @click="openRead(row)">阅读</el-button>
            <el-button link type="primary" @click="openEdit(row)">编辑</el-button>
            <el-button link type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        class="mt-3"
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[10, 20, 50]"
        layout="total, sizes, prev, pager, next"
        @size-change="load"
        @current-change="load"
      />
    </el-card>

    <!-- 编辑抽屉 -->
    <el-drawer v-model="drawerVisible" :title="isEdit ? '编辑文章' : '新增文章'" size="65%" destroy-on-close>
      <el-form :model="form" label-width="70px">
        <el-form-item label="标题" required>
          <el-input v-model="form.title" placeholder="请输入文章标题" />
        </el-form-item>
        <el-row :gutter="16">
          <el-col :span="8">
            <el-form-item label="语言">
              <el-radio-group v-model="form.lang">
                <el-radio label="zh">中文</el-radio>
                <el-radio label="en">英文</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="状态">
              <el-radio-group v-model="form.status">
                <el-radio :label="1">发布</el-radio>
                <el-radio :label="0">草稿</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="分类">
              <el-select v-model="form.categoryID" placeholder="选择分类" clearable style="width: 100%">
                <el-option v-for="c in categories" :key="c.ID" :label="c.name" :value="c.ID" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="封面">
          <el-input v-model="form.cover" placeholder="封面图URL (可选)" />
        </el-form-item>
        <el-form-item label="摘要">
          <el-input v-model="form.summary" type="textarea" :rows="2" placeholder="文章摘要 (可选)" />
        </el-form-item>
        <el-form-item label="内容" required>
          <rich-edit v-model="form.content" style="height: 400px" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="flex justify-end gap-2">
          <el-button @click="drawerVisible = false">取消</el-button>
          <el-button type="primary" :loading="submitting" @click="handleSubmit">保存</el-button>
        </div>
      </template>
    </el-drawer>

    <!-- 阅读弹窗 -->
    <el-dialog v-model="readVisible" :title="currentArticle.title" width="800px" destroy-on-close class="article-read-dialog">
      <div class="article-read-body">
        <div class="article-meta flex items-center gap-4 text-sm text-gray-500 mb-4">
          <span v-if="currentArticle.lang === 'zh'" class="lang-tag">中文</span>
          <span v-else-if="currentArticle.lang === 'en'" class="lang-tag">英文</span>
          <span>作者: {{ currentArticle.nickName }}</span>
          <span>阅读: {{ currentArticle.readCount }}</span>
          <span>{{ formatDate(currentArticle.CreatedAt) }}</span>
        </div>
        <div v-if="currentArticle.cover" class="mb-4">
          <el-image :src="currentArticle.cover" fit="cover" style="width: 100%; max-height: 300px; border-radius: 8px;" />
        </div>
        <div v-if="currentArticle.summary" class="article-summary mb-4 p-3 bg-gray-50 rounded text-gray-600 italic">
          {{ currentArticle.summary }}
        </div>
        <div class="article-content" v-html="currentArticle.content" />
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import RichEdit from '@/components/richtext/rich-edit.vue'
import {
  getArticleList,
  createArticle,
  updateArticle,
  deleteArticle,
  findArticle,
  readArticle,
  getAllCategories
} from '@/plugin/article/api/article'

const list = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const categories = ref([])

const search = ref({ title: '', lang: '', status: null, categoryID: null })

const drawerVisible = ref(false)
const isEdit = ref(false)
const submitting = ref(false)
const form = ref({
  title: '',
  summary: '',
  content: '',
  cover: '',
  lang: 'zh',
  status: 1,
  categoryID: null
})

const readVisible = ref(false)
const currentArticle = ref({})

const loadCategories = async () => {
  const res = await getAllCategories()
  if (res.code === 0) categories.value = res.data || []
}

const load = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      pageSize: pageSize.value
    }
    if (search.value.title) params.title = search.value.title
    if (search.value.lang) params.lang = search.value.lang
    if (search.value.status !== null && search.value.status !== '' && search.value.status !== undefined) {
      params.status = search.value.status
    }
    if (search.value.categoryID !== null && search.value.categoryID !== '' && search.value.categoryID !== undefined) {
      params.categoryID = search.value.categoryID
    }
    const res = await getArticleList(params)
    if (res.code === 0) {
      list.value = res.data.list || []
      total.value = res.data.total || 0
    }
  } finally {
    loading.value = false
  }
}

const resetSearch = () => {
  search.value = { title: '', lang: '', status: null, categoryID: null }
  page.value = 1
  load()
}

const openEdit = async (row) => {
  await loadCategories()
  if (row) {
    isEdit.value = true
    const res = await findArticle({ ID: row.ID })
    if (res.code === 0) {
      form.value = { ...res.data }
    }
  } else {
    isEdit.value = false
    form.value = {
      title: '',
      summary: '',
      content: '',
      cover: '',
      lang: 'zh',
      status: 1,
      categoryID: null
    }
  }
  drawerVisible.value = true
}

const handleSubmit = async () => {
  if (!form.value.title) {
    ElMessage.warning('请输入标题')
    return
  }
  if (!form.value.content) {
    ElMessage.warning('请输入内容')
    return
  }
  submitting.value = true
  try {
    const fn = isEdit.value ? updateArticle : createArticle
    const res = await fn(form.value)
    if (res.code === 0) {
      ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
      drawerVisible.value = false
      load()
    } else {
      ElMessage.error(res.msg || '操作失败')
    }
  } finally {
    submitting.value = false
  }
}

const handleDelete = (row) => {
  ElMessageBox.confirm('确认删除该文章?', '提示', { type: 'warning' }).then(async () => {
    const res = await deleteArticle({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      load()
    } else {
      ElMessage.error(res.msg || '删除失败')
    }
  })
}

const openRead = async (row) => {
  const res = await readArticle({ ID: row.ID })
  if (res.code === 0) {
    currentArticle.value = res.data
    readVisible.value = true
  } else {
    ElMessage.error(res.msg || '读取失败')
  }
}

onMounted(() => {
  load()
  loadCategories()
})
</script>

<style scoped>
.article-read-dialog :deep(.el-dialog__body) {
  padding-top: 0;
}
.article-read-body {
  max-height: 70vh;
  overflow-y: auto;
}
.article-meta .lang-tag {
  padding: 2px 8px;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);
  font-size: 12px;
}
.article-summary {
  border-left: 3px solid var(--el-color-primary);
}
.article-content :deep(img) {
  max-width: 100%;
  border-radius: 6px;
}
.article-content :deep(p) {
  line-height: 1.8;
  margin: 0.8em 0;
}
.article-content :deep(h1),
.article-content :deep(h2),
.article-content :deep(h3) {
  margin-top: 1.2em;
  margin-bottom: 0.6em;
}
</style>
