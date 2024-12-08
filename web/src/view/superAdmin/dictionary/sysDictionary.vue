<template>
  <div>
    <warning-bar
      title="获取字典且缓存方法已在前端utils/dictionary 已经封装完成 不必自己书写 使用方法查看文件内注释"
    />
    <div class="flex gap-4 p-2">
      <div
        class="flex-none w-52 bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900 rounded p-4"
      >
        <div class="flex justify-between items-center">
          <span class="text font-bold">字典列表</span>
          <el-button type="primary" @click="openDrawer"> 新增 </el-button>
        </div>
        <el-scrollbar class="mt-4" style="height: calc(100vh - 300px)">
          <div
            v-for="dictionary in dictionaryData"
            :key="dictionary.ID"
            class="rounded flex justify-between items-center px-2 py-4 cursor-pointer mt-2 hover:bg-blue-50 dark:hover:bg-blue-900 bg-gray-50 dark:bg-gray-800 gap-4"
            :class="
              selectID === dictionary.ID
                ? 'text-active'
                : 'text-slate-700 dark:text-slate-50'
            "
            @click="toDetail(dictionary)"
          >
            <span class="max-w-[160px] truncate">{{ dictionary.name }}</span>
            <div class="min-w-[40px]">
              <el-icon
                class="text-blue-500"
                @click.stop="updateSysDictionaryFunc(dictionary)"
              >
                <Edit />
              </el-icon>
              <el-icon
                class="ml-2 text-red-500"
                @click="deleteSysDictionaryFunc(dictionary)"
              >
                <Delete />
              </el-icon>
            </div>
          </div>
        </el-scrollbar>
      </div>
      <div
        class="flex-1 bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900"
      >
        <sysDictionaryDetail :sys-dictionary-i-d="selectID" />
      </div>
    </div>
    <el-drawer
      v-model="drawerFormVisible"
      :size="appStore.drawerSize"
      :show-close="false"
      :before-close="closeDrawer"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{
            type === 'create' ? '添加字典' : '修改字典'
          }}</span>
          <div>
            <el-button @click="closeDrawer"> 取 消 </el-button>
            <el-button type="primary" @click="enterDrawer"> 确 定 </el-button>
          </div>
        </div>
      </template>
      <el-form
        ref="drawerForm"
        :model="formData"
        :rules="rules"
        label-width="110px"
      >
        <el-form-item label="字典名（中）" prop="name">
          <el-input
            v-model="formData.name"
            placeholder="请输入字典名（中）"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item label="字典名（英）" prop="type">
          <el-input
            v-model="formData.type"
            placeholder="请输入字典名（英）"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status" required>
          <el-switch
            v-model="formData.status"
            active-text="开启"
            inactive-text="停用"
          />
        </el-form-item>
        <el-form-item label="描述" prop="desc">
          <el-input
            v-model="formData.desc"
            placeholder="请输入描述"
            clearable
            :style="{ width: '100%' }"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
  import {
    createSysDictionary,
    deleteSysDictionary,
    updateSysDictionary,
    findSysDictionary,
    getSysDictionaryList
  } from '@/api/sysDictionary' // 此处请自行替换地址
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  import sysDictionaryDetail from './sysDictionaryDetail.vue'
  import { Edit } from '@element-plus/icons-vue'
  import { useAppStore } from "@/pinia";

  defineOptions({
    name: 'SysDictionary'
  })

  const appStore = useAppStore()

  const selectID = ref(0)

  const formData = ref({
    name: null,
    type: null,
    status: true,
    desc: null
  })
  const rules = ref({
    name: [
      {
        required: true,
        message: '请输入字典名（中）',
        trigger: 'blur'
      }
    ],
    type: [
      {
        required: true,
        message: '请输入字典名（英）',
        trigger: 'blur'
      }
    ],
    desc: [
      {
        required: true,
        message: '请输入描述',
        trigger: 'blur'
      }
    ]
  })

  const dictionaryData = ref([])

  // 查询
  const getTableData = async () => {
    const res = await getSysDictionaryList()
    if (res.code === 0) {
      dictionaryData.value = res.data
      selectID.value = res.data[0].ID
    }
  }

  getTableData()

  const toDetail = (row) => {
    selectID.value = row.ID
  }

  const drawerFormVisible = ref(false)
  const type = ref('')
  const updateSysDictionaryFunc = async (row) => {
    const res = await findSysDictionary({ ID: row.ID, status: row.status })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data.resysDictionary
      drawerFormVisible.value = true
    }
  }
  const closeDrawer = () => {
    drawerFormVisible.value = false
    formData.value = {
      name: null,
      type: null,
      status: true,
      desc: null
    }
  }
  const deleteSysDictionaryFunc = async (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const res = await deleteSysDictionary({ ID: row.ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        getTableData()
      }
    })
  }

  const drawerForm = ref(null)
  const enterDrawer = async () => {
    drawerForm.value.validate(async (valid) => {
      if (!valid) return
      let res
      switch (type.value) {
        case 'create':
          res = await createSysDictionary(formData.value)
          break
        case 'update':
          res = await updateSysDictionary(formData.value)
          break
        default:
          res = await createSysDictionary(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage.success('操作成功')
        closeDrawer()
        getTableData()
      }
    })
  }
  const openDrawer = () => {
    type.value = 'create'
    drawerForm.value && drawerForm.value.clearValidate()
    drawerFormVisible.value = true
  }
</script>

<style>
  .dict-box {
    height: calc(100vh - 240px);
  }

  .active {
    background-color: var(--el-color-primary) !important;
    color: #fff;
  }
</style>
