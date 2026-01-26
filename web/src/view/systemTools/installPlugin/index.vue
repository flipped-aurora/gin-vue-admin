<template>
  <div class="gva-form-box">
    <el-upload
      drag
      :action="`${getBaseUrl()}/autoCode/installPlugin`"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleSuccess"
      :headers="{'x-token': token}"
      name="plug"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">拖拽或<em>点击上传</em></div>
      <template #tip>
        <div class="el-upload__tip">请把安装包的zip拖拽至此处上传</div>
      </template>
    </el-upload>

    <!-- Plugin List Table -->
    <div style="margin-top: 20px;">
      <el-table :data="pluginList" style="width: 100%">
        <el-table-column type="expand">
            <template #default="props">
                <div style="padding: 20px;">
                    <h3>API 列表</h3>
                    <el-table :data="props.row.apis" border>
                        <el-table-column prop="path" label="路径" />
                        <el-table-column prop="method" label="方法" />
                        <el-table-column prop="description" label="描述" />
                        <el-table-column prop="apiGroup" label="APIGROUP" />
                    </el-table>
                    <h3>菜单列表</h3>
                    <el-table :data="props.row.menus" row-key="name" :tree-props="{children: 'children', hasChildren: 'hasChildren'}" border>
                        <el-table-column prop="meta.title" label="标题" />
                        <el-table-column prop="name" label="Name" />
                        <el-table-column prop="path" label="Path" />
                    </el-table>
                     <h3>字典列表</h3>
                     <el-table :data="props.row.dictionaries" border>
                         <el-table-column prop="name" label="字典名" />
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
        <el-table-column label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="delete" @click="deletePlugin(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { ElMessage } from 'element-plus'
  import { getBaseUrl } from '@/utils/format'
  import { useUserStore } from "@/pinia";
  import { getPluginList, removePlugin } from '@/api/autoCode'
  import { ElMessageBox } from 'element-plus'

  const userStore = useUserStore()
  const token = userStore.token
  const pluginList = ref([])

  const getTableData = async () => {
    const res = await getPluginList()
    if (res.code === 0) {
      pluginList.value = res.data
    }
  }

  const typeMap = {
    "server": "后端插件",
    "web": "前端插件",
    "full": "全栈插件"
  }

  const deletePlugin = (row) => {
    ElMessageBox.confirm(
    '此操作将永久删除该插件及其关联的API、菜单和字典数据, 是否继续?',
    '提示',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  )
    .then(async () => {
      const res = await removePlugin({ pluginName: row.pluginName, pluginType: row.pluginType })
      if (res.code === 0) {
        ElMessage.success('删除成功')
        getTableData()
      }
    })
    .catch(() => {
    })
  }

  onMounted(() => {
    getTableData()
  })

  const handleSuccess = (res) => {
    if (res.code === 0) {
      let msg = ``
      res.data &&
        res.data.forEach((item, index) => {
          msg += `${index + 1}.${item.msg}\n`
        })
      alert(msg)
      getTableData() // Refresh list on success
    } else {
      ElMessage.error(res.msg)
    }
  }
</script>
