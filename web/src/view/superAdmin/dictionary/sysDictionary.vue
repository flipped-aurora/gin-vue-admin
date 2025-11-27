<template>
  <div>
    <warning-bar
      title="获取字典且缓存方法已在前端utils/dictionary 已经封装完成 不必自己书写 使用方法查看文件内注释"
    />
    <el-splitter class="h-full">
      <el-splitter-panel size="300px" min="200px" max="800px" collapsible>
        <div
          class="flex-none bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900 rounded p-4"
        >
          <div class="flex justify-between items-center relative">
            <span class="text font-bold">字典列表</span>
            <el-input
              class="!absolute top-0 left-0 z-2 ease-in-out animate-slide-left"
              placeholder="搜索"
              v-if="showSearchInput"
              v-model="searchName"
              clearable
              :autofocus="showSearchInput"
              @clear="clearSearchInput"
              :prefix-icon="Search"
              v-click-outside="handleCloseSearchInput"
              @keydown="handleInputKeyDown"
            >
              <template #append>
                <el-button
                  :type="searchName ? 'primary' : 'info'"
                  @click="getTableData"
                  >搜索</el-button
                >
              </template>
            </el-input>
            <el-button-group class="ml-auto">
              <el-tooltip content="搜索" placement="top">
                <el-button
                  :icon="Search"
                  @click="showSearchInputHandler"
                />
              </el-tooltip>
              <el-tooltip content="导入字典" placement="top">
                <el-button
                  type="success"
                  :icon="Upload"
                  @click="openImportDialog"
                />
              </el-tooltip>
              <el-tooltip content="AI 生成字典" placement="top">
                <el-button
                  type="warning"
                  @click="openAiDialog"
                >
                  AI
                </el-button>
              </el-tooltip>
              <el-tooltip content="新建字典" placement="top">
                <el-button
                  type="primary"
                  :icon="Plus"
                  @click="openDrawer"
                />
              </el-tooltip>
            </el-button-group>
          </div>
          <el-scrollbar class="mt-4" style="height: calc(100vh - 300px)">
            <div
              v-for="dictionary in dictionaryData"
              :key="dictionary.ID"
              class="rounded flex justify-between items-center px-2 py-4 cursor-pointer mt-2 hover:bg-blue-50 dark:hover:bg-blue-900 bg-gray-50 dark:bg-gray-800 gap-4"
              :class="[
                selectID === dictionary.ID
                  ? 'text-active'
                  : 'text-slate-700 dark:text-slate-50',
                dictionary.parentID ? 'ml-4 border-l-2 border-blue-200' : ''
              ]"
              @click="toDetail(dictionary)"
            >
              <div class="max-w-[160px] truncate">
                <span
                  v-if="dictionary.parentID"
                  class="text-xs text-gray-400 mr-1"
                  >└─</span
                >
                {{ dictionary.name }}
                <span class="mr-auto text-sm">（{{ dictionary.type }}）</span>
              </div>

              <div class="min-w-[60px] flex items-center gap-2">
                <el-icon
                  class="!text-green-500"
                  @click.stop="exportDictionary(dictionary)"
                  title="导出字典"
                >
                  <Download />
                </el-icon>
                <el-icon
                  class="!text-blue-500"
                  @click.stop="updateSysDictionaryFunc(dictionary)"
                >
                  <Edit />
                </el-icon>
                <el-icon
                  class="!text-red-500"
                  @click="deleteSysDictionaryFunc(dictionary)"
                >
                  <Delete />
                </el-icon>
              </div>
            </div>
          </el-scrollbar>
        </div>
      </el-splitter-panel>
      <el-splitter-panel :min="200">
        <div
          class="flex-1 bg-white text-slate-700 dark:text-slate-400 dark:bg-slate-900"
        >
          <sysDictionaryDetail :sys-dictionary-i-d="selectID" />
        </div>
      </el-splitter-panel>
    </el-splitter>

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
        <el-form-item label="父级字典" prop="parentID">
          <el-select
            v-model="formData.parentID"
            placeholder="请选择父级字典（可选）"
            clearable
            filterable
            :style="{ width: '100%' }"
          >
            <el-option
              v-for="dict in availableParentDictionaries"
              :key="dict.ID"
              :label="`${dict.name}（${dict.type}）`"
              :value="dict.ID"
            />
          </el-select>
        </el-form-item>
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

    <!-- 导入字典抽屉 -->
    <el-drawer
      v-model="importDrawerVisible"
      :size="appStore.drawerSize"
      :show-close="false"
      :before-close="closeImportDrawer"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">导入字典JSON</span>
          <div>
            <el-button @click="closeImportDrawer"> 取 消 </el-button>
            <el-button type="primary" @click="handleImport" :loading="importing">
              确认导入
            </el-button>
          </div>
        </div>
      </template>
      
      <div class="import-drawer-content">
        <div class="mb-4">
          <el-alert
            title="请粘贴、编辑或拖拽JSON文件到下方区域"
            type="info"
            :closable="false"
            show-icon
          />
        </div>

        <!-- 拖拽上传区域 -->
        <div
          class="drag-upload-area"
          :class="{ 'is-dragging': isDragging }"
          @drop.prevent="handleDrop"
          @dragover.prevent="handleDragOver"
          @dragleave.prevent="handleDragLeave"
          @click="triggerFileInput"
        >
          <el-icon class="upload-icon"><Upload /></el-icon>
          <div class="upload-text">
            <p>将 JSON 文件拖到此处，或点击选择文件</p>
            <p class="upload-hint">也可以在下方文本框直接编辑</p>
          </div>
          <input
            ref="fileInputRef"
            type="file"
            accept=".json,application/json"
            style="display: none"
            @change="handleFileSelect"
          />
        </div>

        <div class="json-editor-container mt-4">
          <el-input
            v-model="importJsonText"
            type="textarea"
            :rows="15"
            placeholder='请输入JSON数据，例如：
{
  "name": "性别",
  "type": "gender",
  "status": true,
  "desc": "性别字典",
  "sysDictionaryDetails": [
    {
      "label": "男",
      "value": "1",
      "status": true,
      "sort": 1
    },
    {
      "label": "女",
      "value": "2",
      "status": true,
      "sort": 2
    }
  ]
}'
            class="json-textarea"
          />
        </div>

        <div class="mt-4" v-if="jsonPreviewError">
          <el-alert
            :title="jsonPreviewError"
            type="error"
            :closable="false"
            show-icon
          />
        </div>

    
      </div>
    </el-drawer>

    <!-- AI 对话框 -->
    <el-dialog
      v-model="aiDialogVisible"
      title="AI 生成字典"
      width="520px"
      :before-close="closeAiDialog"
    >
      <div class="relative">
        <el-input
          v-model="aiPrompt"
          type="textarea"
          :rows="6"
          :maxlength="2000"
          placeholder="请输入生成字典的描述，例如：生成一个用户状态字典（启用/禁用）\n支持粘贴或上传图片后识图生成。"
          resize="none"
          @keydown.ctrl.enter="handleAiGenerate"
          @paste="handlePaste"
          @focus="handleFocus"
          @blur="handleBlur"
        />

        <input
          ref="imageFileInputRef"
          type="file"
          accept="image/*"
          style="display:none"
          @change="handleImageSelect"
        />

        <div class="flex absolute right-2 bottom-2">
          <el-tooltip effect="light">
            <template #content>
              <div>粘贴或上传图片后，识别图片内容生成字典。</div>
            </template>
            <el-button type="primary" @click="eyeFunc">
                <el-icon size="18">
                <ai-gva />
              </el-icon>
              识图
            </el-button>
          </el-tooltip>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAiDialog">取 消</el-button>
          <el-button type="primary" @click="handleAiGenerate" :loading="aiGenerating">
            确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
  import {
    createSysDictionary,
    deleteSysDictionary,
    updateSysDictionary,
    findSysDictionary,
    getSysDictionaryList,
    exportSysDictionary,
    importSysDictionary
  } from '@/api/sysDictionary' // 此处请自行替换地址
  import { butler, eye } from '@/api/autoCode'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref, computed, watch } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  import sysDictionaryDetail from './sysDictionaryDetail.vue'
  import { Edit, Plus, Search, Download, Upload } from '@element-plus/icons-vue'
  import { useAppStore } from '@/pinia'

  defineOptions({
    name: 'SysDictionary'
  })

  const appStore = useAppStore()

  const selectID = ref(0)

  const formData = ref({
    name: null,
    type: null,
    status: true,
    desc: null,
    parentID: null
  })
  const searchName = ref('')
  const showSearchInput = ref(false)
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
  const availableParentDictionaries = ref([])

  // 导入相关
  const importDrawerVisible = ref(false)
  const importJsonText = ref('')
  const importing = ref(false)
  const jsonPreviewError = ref('')
  const jsonPreview = ref(null)
  const isDragging = ref(false)
  const fileInputRef = ref(null)

  // AI 相关
  const aiDialogVisible = ref(false)
  const aiPrompt = ref('')
  const aiGenerating = ref(false)

  // 图片上传/识别相关
  const imageFileInputRef = ref(null)
  const focused = ref(false)

  const handleFocus = () => {
    focused.value = true
  }
  const handleBlur = () => {
    focused.value = false
  }

  // 触发图片选择
  const triggerImageSelect = () => {
    imageFileInputRef.value?.click()
  }

  const handlePaste = (event) => {
    const items = event.clipboardData.items;
    for (let i = 0; i < items.length; i++) {
      if (items[i].type.indexOf('image') !== -1) {
        const file = items[i].getAsFile();
        const reader = new FileReader();
        reader.onload =async (e) => {
          const base64String = e.target.result;
          const res = await eye({ picture: base64String,command: 'dictEye' })
          if (res.code === 0) {
            aiPrompt.value = res.data
          }
        };
        reader.readAsDataURL(file);
      }
    }
  };

  const eyeFunc = async () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = 'image/*';

    input.onchange = (event) => {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = async (e) => {
          const base64String = e.target.result;

          const res = await eye({ picture: base64String,command: 'dictEye' })
          if (res.code === 0) {
            aiPrompt.value = res.data
          }
        };
        reader.readAsDataURL(file);
      }
    };

    input.click();
  }



  // 监听JSON文本变化，实时预览
  watch(importJsonText, (newVal) => {
    if (!newVal.trim()) {
      jsonPreview.value = null
      jsonPreviewError.value = ''
      return
    }
    try {
      jsonPreview.value = JSON.parse(newVal)
      jsonPreviewError.value = ''
    } catch (e) {
      jsonPreviewError.value = 'JSON格式错误: ' + e.message
      jsonPreview.value = null
    }
  })

  // 格式化JSON预览
  const jsonPreviewFormatted = computed(() => {
    if (!jsonPreview.value) return ''
    return JSON.stringify(jsonPreview.value, null, 2)
  })


  // 查询
  const getTableData = async () => {
    const res = await getSysDictionaryList({
      name: searchName.value.trim()
    })
    if (res.code === 0) {
      dictionaryData.value = res.data
      selectID.value = res.data[0].ID
      // 更新可选父级字典列表
      updateAvailableParentDictionaries()
    }
  }

  // 更新可选父级字典列表
  const updateAvailableParentDictionaries = () => {
    // 如果是编辑模式，排除当前字典及其子字典
    if (type.value === 'update' && formData.value.ID) {
      availableParentDictionaries.value = dictionaryData.value.filter(
        (dict) => {
          return (
            dict.ID !== formData.value.ID &&
            !isChildDictionary(dict.ID, formData.value.ID)
          )
        }
      )
    } else {
      // 创建模式，显示所有字典
      availableParentDictionaries.value = [...dictionaryData.value]
    }
  }

  // 检查是否为子字典（防止循环引用）
  const isChildDictionary = (dictId, parentId) => {
    const dict = dictionaryData.value.find((d) => d.ID === dictId)
    if (!dict || !dict.parentID) return false
    if (dict.parentID === parentId) return true
    return isChildDictionary(dict.parentID, parentId)
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
      // 更新可选父级字典列表
      updateAvailableParentDictionaries()
    }
  }
  const closeDrawer = () => {
    drawerFormVisible.value = false
    formData.value = {
      name: null,
      type: null,
      status: true,
      desc: null,
      parentID: null
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
    // 更新可选父级字典列表
    updateAvailableParentDictionaries()
  }

  const clearSearchInput = () => {
    if (!showSearchInput.value) return
    searchName.value = ''
    showSearchInput.value = false
    getTableData()
  }
  const handleCloseSearchInput = () => {
    if (!showSearchInput.value || searchName.value.trim() != '') return
    showSearchInput.value = false
  }

  const showSearchInputHandler = () => {
    showSearchInput.value = true
  }

  const handleInputKeyDown = (e) => {
    if (e.key === 'Enter' && searchName.value.trim() !== '') {
      getTableData()
    }
  }

  // 导出字典
  const exportDictionary = async (row) => {
    try {
      const res = await exportSysDictionary({ ID: row.ID })
      if (res.code === 0) {
        // 将JSON数据转换为字符串并下载
        const jsonStr = JSON.stringify(res.data, null, 2)
        const blob = new Blob([jsonStr], { type: 'application/json' })
        const url = URL.createObjectURL(blob)
        const link = document.createElement('a')
        link.href = url
        link.download = `${row.type}_${row.name}_dictionary.json`
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(url)
        ElMessage.success('导出成功')
      }
    } catch (error) {
      ElMessage.error('导出失败: ' + error.message)
    }
  }

  // 打开导入抽屉
  const openImportDialog = () => {
    importDrawerVisible.value = true
    importJsonText.value = ''
    jsonPreview.value = null
    jsonPreviewError.value = ''
    isDragging.value = false
  }

  // 关闭导入抽屉
  const closeImportDrawer = () => {
    importDrawerVisible.value = false
    importJsonText.value = ''
    jsonPreview.value = null
    jsonPreviewError.value = ''
    isDragging.value = false
  }

  // 处理拖拽进入
  const handleDragOver = (e) => {
    isDragging.value = true
  }

  // 处理拖拽离开
  const handleDragLeave = (e) => {
    isDragging.value = false
  }
  // 处理文件拖拽
  const handleDrop = (e) => {
    isDragging.value = false
    const files = e.dataTransfer.files
    if (files.length === 0) return

    const file = files[0]
    readJsonFile(file)
  }

  // 触发文件选择
  const triggerFileInput = () => {
    fileInputRef.value?.click()
  }

  // 处理文件选择
  const handleFileSelect = (e) => {
    const files = e.target.files
    if (files.length === 0) return

    const file = files[0]
    readJsonFile(file)
    
    // 清空input，以便可以重复选择同一文件
    e.target.value = ''
  }

  // 读取JSON文件
  const readJsonFile = (file) => {
    // 检查文件类型
    if (!file.name.endsWith('.json')) {
      ElMessage.warning('请上传 JSON 文件')
      return
    }

    // 读取文件内容
    const reader = new FileReader()
    reader.onload = (event) => {
      try {
        const content = event.target.result
        // 验证是否为有效的 JSON
        JSON.parse(content)
        importJsonText.value = content
        ElMessage.success('文件读取成功')
      } catch (error) {
        ElMessage.error('文件内容不是有效的 JSON 格式')
      }
    }
    reader.onerror = () => {
      ElMessage.error('文件读取失败')
    }
    reader.readAsText(file)
  }

  // 处理导入
  const handleImport = async () => {
    if (!importJsonText.value.trim()) {
      ElMessage.warning('请输入JSON数据')
      return
    }

    if (jsonPreviewError.value) {
      ElMessage.error('JSON格式错误，请检查后重试')
      return
    }

    try {
      importing.value = true
      const res = await importSysDictionary({ json: importJsonText.value })
      if (res.code === 0) {
        ElMessage.success('导入成功')
        closeImportDrawer()
        getTableData()
      }
    } catch (error) {
      ElMessage.error('导入失败: ' + error.message)
    } finally {
      importing.value = false
    }
  }

  // 打开 AI 对话框
  const openAiDialog = () => {
    aiDialogVisible.value = true
    aiPrompt.value = ''
  }

  // 关闭 AI 对话框
  const closeAiDialog = () => {
    aiDialogVisible.value = false
    aiPrompt.value = ''
  }

  // 处理 AI 生成
  const handleAiGenerate = async () => {
    if (!aiPrompt.value.trim()) {
      ElMessage.warning('请输入描述内容')
      return
    }
    try {
      aiGenerating.value = true
      const aiRes = await butler({
        prompt: aiPrompt.value,
        command: 'dict'
      })
      if (aiRes && aiRes.code === 0) {
        ElMessage.success('AI 生成成功')
        try {
          // 将 AI 返回的数据填充到导入文本框（支持字符串或对象）
          if (typeof aiRes.data === 'string') {
            importJsonText.value = aiRes.data
          } else {
            importJsonText.value = JSON.stringify(aiRes.data, null, 2)
          }
          // 清除可能的解析错误并打开导入抽屉
          jsonPreviewError.value = ''
          importDrawerVisible.value = true
          closeAiDialog()
        } catch (e) {
          ElMessage.error('处理 AI 返回结果失败: ' + (e.message || e))
        }
      } else {
        ElMessage.error(aiRes.msg || 'AI 生成失败')
      }
    } catch (err) {
      ElMessage.error('AI 调用失败: ' + (err.message || err))
    } finally {
      aiGenerating.value = false
    }
  }
</script>

<style scoped>
  .dict-box {
    height: calc(100vh - 240px);
  }

  .active {
    background-color: var(--el-color-primary) !important;
    color: #fff;
  }

  .import-drawer-content {
    padding: 0 4px;
  }

  /* 拖拽上传区域 */
  .drag-upload-area {
    border: 2px dashed #dcdfe6;
    border-radius: 8px;
    padding: 40px 20px;
    text-align: center;
    background-color: #fafafa;
    transition: all 0.3s ease;
    cursor: pointer;
  }

  .drag-upload-area:hover {
    border-color: #409eff;
    background-color: #ecf5ff;
  }

  .drag-upload-area.is-dragging {
    border-color: #409eff;
    background-color: #ecf5ff;
    transform: scale(1.02);
  }

  .upload-icon {
    font-size: 48px;
    color: #8c939d;
    margin-bottom: 16px;
  }

  .drag-upload-area.is-dragging .upload-icon {
    color: #409eff;
  }

  .upload-text {
    color: #606266;
  }

  .upload-text p {
    margin: 4px 0;
  }

  .upload-hint {
    font-size: 12px;
    color: #909399;
  }

  .json-editor-container {
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    overflow: hidden;
  }

  .json-textarea :deep(.el-textarea__inner) {
    font-family: 'Courier New', Courier, monospace;
    font-size: 13px;
    line-height: 1.5;
  }

  .json-preview {
    background-color: #f5f7fa;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    padding: 16px;
    max-height: 400px;
    overflow: auto;
  }

  .json-preview pre {
    margin: 0;
    font-family: 'Courier New', Courier, monospace;
    font-size: 13px;
    line-height: 1.5;
    white-space: pre-wrap;
    word-wrap: break-word;
  }

  .dark .drag-upload-area {
    background-color: #1d1e1f;
    border-color: #414243;
  }

  .dark .drag-upload-area:hover,
  .dark .drag-upload-area.is-dragging {
    background-color: #1a3a52;
    border-color: #409eff;
  }

  .dark .json-preview {
    background-color: #1d1e1f;
    border-color: #414243;
  }
</style>
