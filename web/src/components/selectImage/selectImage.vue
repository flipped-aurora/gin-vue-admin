<template>
  <div>
    <selectComponent :rounded="rounded" v-if="!props.multiple" :model="model" @chooseItem="openChooseImg" @deleteItem="openChooseImg" />
    <div v-else class="w-full gap-4 flex flex-wrap">
      <draggable 
        v-model="model" 
        class="flex flex-wrap gap-4" 
        item-key="url" 
        ghost-class="ghost-item" 
        handle=".drag-handle"
        animation="300"
        @start="onDragStart"
        @end="onDragEnd"
      >
        <template #item="{element, index}">
          <div class="relative group">
            <div class="drag-handle absolute left-2 top-2 w-8 h-8 flex items-center justify-center cursor-move z-10 opacity-0 group-hover:opacity-100 rounded-full">
              <el-icon :size="18"><Menu /></el-icon>
            </div>
            <selectComponent :rounded="rounded" :model="element" @chooseItem="openChooseImg"
                           @deleteItem="deleteImg(index)"
            />
          </div>
        </template>
      </draggable>
      <selectComponent :rounded="rounded" v-if="model?.length < props.maxUpdateCount || props.maxUpdateCount === 0"
                       @chooseItem="openChooseImg" @deleteItem="openChooseImg"
      />
    </div>

    <el-drawer v-model="drawer" title="媒体库 | 点击“文件名”可以编辑，选择的类别即是上传的类别" :size="880">
      <div class="flex">
        <div class="w-64" style="border-right: solid 1px var(--el-border-color);">
          <el-scrollbar style="height: calc(100vh - 110px)">
            <el-tree
                :data="categories"
                node-key="id"
                :props="defaultProps"
                @node-click="handleNodeClick"
                default-expand-all
            >
              <template #default="{ node, data }">
                <div class="w-36" :class="search.classId === data.ID ? 'text-blue-500 font-bold' : ''">{{ data.name }}
                </div>
                <el-dropdown>
                  <el-icon class="ml-3 text-right" v-if="data.ID > 0"><MoreFilled /></el-icon>
                  <el-icon class="ml-3 text-right mt-1" v-else><Plus /></el-icon>
                  <template #dropdown>
                    <el-dropdown-menu>
                      <el-dropdown-item @click="addCategoryFun(data)">添加分类</el-dropdown-item>
                      <el-dropdown-item @click="editCategory(data)" v-if="data.ID > 0">编辑分类</el-dropdown-item>
                      <el-dropdown-item @click="deleteCategoryFun(data.ID)" v-if="data.ID > 0">删除分类</el-dropdown-item>
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </template>
            </el-tree>
          </el-scrollbar>
        </div>
        <div class="ml-4 w-[605px]">
          <div class="gva-btn-list gap-2">
            <el-input v-model.trim="search.keyword" class="w-96" placeholder="请输入文件名或备注" clearable />
            <el-button type="primary" icon="search" @click="onSubmit"></el-button>
          </div>
          <div class="gva-btn-list gap-2">
            <el-button @click="useSelectedImages" type="danger" :disabled="selectedImages.length === 0" :icon="ArrowLeftBold">选定</el-button>
            <upload-common :image-common="imageCommon" :classId="search.classId" @on-success="onSuccess" />
            <cropper-image :classId="search.classId" @on-success="onSuccess" />
            <QRCodeUpload :classId="search.classId" @on-success="onSuccess" />
            <upload-image :image-url="imageUrl" :file-size="2048" :max-w-h="1080" :classId="search.classId" @on-success="onSuccess" />
          </div>
          <div class="flex flex-wrap gap-4">
            <div v-for="(item,key) in picList" :key="key" class="w-40">
              <div class="w-40 h-40 border rounded overflow-hidden border-dashed border-gray-300 cursor-pointer relative group">
                <el-image :key="key" :src="getUrl(item.url)" fit="cover" class="w-full h-full relative" @click="toggleImageSelection(item)" :class="{ selected: isSelected(item) }">
                  <template #error>
                    <el-icon v-if="isVideoExt(item.url || '')" :size="32" class="absolute top-[calc(50%-16px)] left-[calc(50%-16px)]">
                      <VideoPlay />
                    </el-icon>
                    <video v-if="isVideoExt(item.url || '')"
                           class="w-full h-full object-cover"
                           muted
                           preload="metadata"
                           @click="toggleImageSelection(item)"
                           :class="{ selected: isSelected(item) }"
                    >
                      <source :src="getUrl(item.url) + '#t=1'">
                      您的浏览器不支持视频播放
                    </video>
                    <div v-else class="w-full h-full object-cover flex items-center justify-center">
                      <el-icon :size="32">
                        <icon-picture />
                      </el-icon>
                    </div>
                  </template>
                </el-image>
                <div class="absolute -right-1 top-1 w-8 h-8 group-hover:inline-block hidden" @click="deleteCheck(item)">
                  <el-icon :size="18">
                    <CloseBold />
                  </el-icon>
                </div>
              </div>
              <div class="overflow-hidden text-nowrap overflow-ellipsis text-center w-full cursor-pointer" @click="editFileNameFunc(item)">
                {{ item.name }}
              </div>
            </div>
          </div>
          <el-pagination
              :current-page="page"
              :page-size="pageSize"
              :total="total"
              class="justify-center"
              layout="total, prev, pager, next, jumper"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
          />
        </div>
      </div>
    </el-drawer>


    <!-- 添加分类弹窗 -->
    <el-dialog v-model="categoryDialogVisible" @close="closeAddCategoryDialog" width="520"
               :title="(categoryFormData.ID === 0 ? '添加' : '编辑') + '分类'"
               draggable
    >
      <el-form ref="categoryForm" :rules="rules" :model="categoryFormData" label-width="80px">
        <el-form-item label="上级分类">
          <el-tree-select
              v-model="categoryFormData.pid"
              :data="categories"
              check-strictly
              :props="defaultProps"
              :render-after-expand="false"
              style="width: 240px"
          />
        </el-form-item>
        <el-form-item label="分类名称" prop="name">
          <el-input v-model.trim="categoryFormData.name" placeholder="分类名称"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="closeAddCategoryDialog">取消</el-button>
        <el-button type="primary" @click="confirmAddCategory">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { getUrl, isVideoExt } from '@/utils/image'
import { ref } from 'vue'
import { getFileList, editFileName, deleteFile } from '@/api/fileUploadAndDownload'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  ArrowLeftBold,
  CloseBold,
  Menu,
  MoreFilled,
  Picture as IconPicture,
  Plus,
  VideoPlay
} from '@element-plus/icons-vue'
import selectComponent from '@/components/selectImage/selectComponent.vue'
import { addCategory, deleteCategory, getCategoryList } from '@/api/attachmentCategory'
import CropperImage from "@/components/upload/cropper.vue";
import QRCodeUpload from "@/components/upload/QR-code.vue";
import draggable from 'vuedraggable'

const imageUrl = ref('')
const imageCommon = ref('')

const search = ref({
  keyword: null,
  classId: 0
})
const page = ref(1)
const total = ref(0)
const pageSize = ref(20)

const model = defineModel({ type: [String, Array] })

const props = defineProps({
  multiple: {
    type: Boolean,
    default: false
  },
  fileType: {
    type: String,
    default: ''
  },
  maxUpdateCount: {
    type: Number,
    default: 0
  },
  rounded: {
    type: Boolean,
    default: false
  }
})

const deleteImg = (index) => {
  model.value.splice(index, 1)
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getImageList()
}

const handleCurrentChange = (val) => {
  page.value = val
  getImageList()
}

const onSubmit = () => {
  search.value.classId = 0
  page.value = 1
  getImageList()
}

const editFileNameFunc = async(row) => {
  ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '不能为空',
    inputValue: row.name
  }).then(async({ value }) => {
    row.name = value
    const res = await editFileName(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '编辑成功!'
      })
      await getImageList()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '取消修改'
    })
  })
}

const drawer = ref(false)
const picList = ref([])

const imageTypeList = ['png', 'jpg', 'jpeg', 'gif', 'bmp', 'webp', 'svg']
const videoTypeList = ['mp4', 'avi', 'rmvb', 'rm', 'asf', 'divx', 'mpg', 'mpeg', 'mpe', 'wmv', 'mkv', 'vob']

const listObj = {
  image: imageTypeList,
  video: videoTypeList
}

const chooseImg = (url) => {
  if (props.fileType) {
    const typeSuccess = listObj[props.fileType].some(item => {
      if (url?.toLowerCase().includes(item)) {
        return true
      }
    })
    if (!typeSuccess) {
      ElMessage({
        type: 'error',
        message: '当前类型不支持使用'
      })
      return
    }
  }
  //if (props.multiple) {
  //  model.value.push(url)
  //} else {
  model.value = url
  //}
  drawer.value = false
}

const openChooseImg = async() => {
  if (model.value && !props.multiple) {
    model.value = ''
    return
  }
  await getImageList()
  await fetchCategories()
  drawer.value = true
}

const getImageList = async() => {
  const res = await getFileList({ page: page.value, pageSize: pageSize.value, ...search.value })
  if (res.code === 0) {
    picList.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}

const deleteCheck = (item) => {
  ElMessageBox.confirm('是否删除该文件', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const res = await deleteFile(item)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功!'
      })
      await getImageList()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '已取消删除'
    })
  })
}

const defaultProps = {
  children: 'children',
  label: 'name',
  value: 'ID'
}

const categories = ref([])
const fetchCategories = async() => {
  const res = await getCategoryList()
  let data = {
    name: '全部分类',
    ID: 0,
    pid: 0,
    children:[]
  }
  if (res.code === 0) {
    categories.value = res.data || []
    categories.value.unshift(data)
  }
}

const handleNodeClick = (node) => {
  search.value.keyword = null
  search.value.classId = node.ID
  page.value = 1
  getImageList()
}

const onSuccess = () => {
  search.value.keyword = null
  page.value = 1
  getImageList()
}

const categoryDialogVisible = ref(false)
const categoryFormData = ref({
  ID: 0,
  pid: 0,
  name: ''
})

const categoryForm = ref(null)
const rules = ref({
  name: [
    { required: true, message: '请输入分类名称', trigger: 'blur' },
    { max: 20, message: '最多20位字符', trigger: 'blur' }
  ]
})

const addCategoryFun = (category) => {
  categoryDialogVisible.value = true
  categoryFormData.value.ID = 0
  categoryFormData.value.pid = category.ID
}

const editCategory = (category) => {
  categoryFormData.value = {
    ID: category.ID,
    pid: category.pid,
    name: category.name
  }
  categoryDialogVisible.value = true
}

const deleteCategoryFun = async(id) => {
  const res = await deleteCategory({ id: id })
  if (res.code === 0) {
    ElMessage.success({ type: 'success', message: '删除成功' })
    await fetchCategories()
  }
}

const confirmAddCategory = async() => {
  categoryForm.value.validate(async valid => {
    if (valid) {
      const res = await addCategory(categoryFormData.value)
      if (res.code === 0) {
        ElMessage({ type: 'success', message: '操作成功' })
        await fetchCategories()
        closeAddCategoryDialog()
      }
    }
  })
}

const closeAddCategoryDialog = () => {
  categoryDialogVisible.value = false
  categoryFormData.value = {
    ID: 0,
    pid: 0,
    name: ''
  }
}

const selectedImages = ref([])

const toggleImageSelection = (item) => {
  if (props.multiple === false) {
    chooseImg(item.url)
    return
  }
  const index = selectedImages.value.findIndex(img => img.ID === item.ID)
  if (index > -1) {
    selectedImages.value.splice(index, 1)
  } else {
    selectedImages.value.push(item)
  }
}

const isSelected = (item) => {
  return selectedImages.value.some(img => img.ID === item.ID)
}

const useSelectedImages = () => {
  selectedImages.value.forEach((item) => {
    model.value.push(item.url)
  })
  drawer.value = false
  selectedImages.value = []
}

const onDragStart = () => {
  // 拖拽开始时的处理
  document.body.style.cursor = 'grabbing'
}

const onDragEnd = () => {
  // 拖拽结束时的处理
  document.body.style.cursor = 'default'
  // 确保model是数组类型
  if (!Array.isArray(model.value)) {
    model.value = []
  }
}

</script>
<style scoped>
.selected {
  border: 3px solid #409eff;
}

.selected:before {
  content: "";
  position: absolute;
  left: 0;
  top: 0;
  border: 10px solid #409eff;
}

.selected:after {
  content: "";
  width: 9px;
  height: 14px;
  position: absolute;
  left: 6px;
  top: 0;
  border: 3px solid #fff;
  border-top-color: transparent;
  border-left-color: transparent;
  transform: rotate(45deg);
}

.ghost-item {
  opacity: 0.5;
  background: #c8ebfb;
  border: 1px dashed #409eff;
}

.drag-handle {
  background-color: rgba(64, 158, 255, 0.1);
  border-radius: 4px;
  transition: opacity 0.3s;
}

.drag-handle:hover {
  background-color: rgba(64, 158, 255, 0.2);
}
</style>
