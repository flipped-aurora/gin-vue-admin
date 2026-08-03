<template>
  <div class="gva-form-box">
    <el-upload
      drag
      :action="`${getBaseUrl()}/autoCode/installPlugin`"
      :show-file-list="false"
      :on-success="handleInstallSuccess"
      :on-error="handleInstallError"
      :headers="{ 'x-token': token }"
      name="plug"
    >
      <div class="el-icon--upload flex justify-center">
        <svg-icon icon="lucide:upload" />
      </div>
      <div class="el-upload__text">拖拽或 <em>点击上传</em></div>
      <template #tip>
        <div class="el-upload__tip">请把安装包的 zip 拖拽至此处上传</div>
      </template>
    </el-upload>

    <div class="mt-5">
      <el-table :data="pluginList" class="w-full">
        <el-table-column type="expand">
          <template #default="props">
            <div class="p-5">
              <h3>API 列表</h3>
              <el-table :data="props.row.apis" border>
                <el-table-column prop="path" label="路径" />
                <el-table-column prop="method" label="方法" />
                <el-table-column prop="description" label="描述" />
                <el-table-column prop="apiGroup" label="API Group" />
              </el-table>
              <h3>菜单列表</h3>
              <el-table
                :data="props.row.menus"
                row-key="name"
                :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
                border
              >
                <el-table-column prop="meta.title" label="标题" />
                <el-table-column prop="name" label="Name" />
                <el-table-column prop="path" label="Path" />
              </el-table>
              <h3>字典列表</h3>
              <el-table :data="props.row.dictionaries" border>
                <el-table-column prop="name" label="字典名称" />
                <el-table-column prop="type" label="字典类型" />
                <el-table-column prop="desc" label="描述" />
              </el-table>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="pluginName" label="插件名称" />
        <el-table-column prop="pluginType" label="插件类型">
          <template #default="scope">
            {{ typeMap[scope.row.pluginType] || '未知类型' }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="190">
          <template #default="scope">
            <el-button
              type="primary"
              link
              :disabled="!canInstallSubPlugin(scope.row)"
              @click="openSubPluginInstall(scope.row)"
            >
              安装子插件
            </el-button>
            <el-button type="danger" link @click="deletePlugin(scope.row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog
      v-model="subPluginInstallVisible"
      :title="`安装子插件：${parentPlugin}`"
      width="520px"
      destroy-on-close
      @closed="resetSubPluginInstall"
    >
      <el-upload
        drag
        :action="`${getBaseUrl()}/autoCode/installPlugin`"
        :data="{ parentPlugin }"
        :show-file-list="false"
        :on-success="handleSubPluginInstallSuccess"
        :on-error="handleInstallError"
        :headers="{ 'x-token': token }"
        name="plug"
      >
        <div class="el-icon--upload flex justify-center">
          <svg-icon icon="lucide:upload" />
        </div>
        <div class="el-upload__text">拖拽或 <em>点击上传</em></div>
        <template #tip>
          <div class="el-upload__tip">请把子插件安装包的 zip 拖拽至此处上传</div>
        </template>
      </el-upload>
    </el-dialog>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getBaseUrl } from '@/utils/format'
import { useUserStore } from '@/pinia'
import { getPluginList, removePlugin } from '@/api/autoCode'
import {
  canInstallSubPlugin,
  createSubPluginUploadData,
} from './subPluginInstall.js'

const userStore = useUserStore()
const token = userStore.token
const pluginList = ref([])
const parentPlugin = ref('')
const subPluginInstallVisible = ref(false)

const typeMap = {
  server: '后端插件',
  web: '前端插件',
  full: '全栈插件',
}

const getTableData = async () => {
  const res = await getPluginList()
  if (res.code === 0) {
    pluginList.value = res.data
  }
}

const openSubPluginInstall = (plugin) => {
  const uploadData = createSubPluginUploadData(plugin)
  if (!uploadData) {
    return
  }

  parentPlugin.value = uploadData.parentPlugin
  subPluginInstallVisible.value = true
}

const resetSubPluginInstall = () => {
  parentPlugin.value = ''
}

const deletePlugin = (row) => {
  ElMessageBox.confirm(
    '此操作将永久删除该插件及其关联的 API、菜单和字典数据，是否继续？',
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(async () => {
      const res = await removePlugin({
        pluginName: row.pluginName,
        pluginType: row.pluginType,
      })
      if (res.code === 0) {
        ElMessage.success('删除成功')
        getTableData()
      }
    })
    .catch(() => {})
}

const handleInstallSuccess = (res, isSubPlugin = false) => {
  if (res.code !== 0) {
    ElMessage.error(res.msg)
    return
  }

  const messages = res.data
    ?.map((item, index) => `${index + 1}. ${item.msg}`)
    .join('\n')
  ElMessage.success(messages || '安装成功')
  if (isSubPlugin) {
    subPluginInstallVisible.value = false
  }
  getTableData()
}

const handleSubPluginInstallSuccess = (res) => {
  handleInstallSuccess(res, true)
}

const handleInstallError = (error) => {
  ElMessage.error(error?.message || '安装失败，请检查安装包后重试')
}

onMounted(() => {
  getTableData()
})
</script>
