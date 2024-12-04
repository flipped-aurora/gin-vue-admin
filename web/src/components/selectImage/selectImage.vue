<template>
  <div>
    <selectComponent
      v-if="!props.multiple"
      :model="model"
      @chooseItem="openChooseImg"
      @deleteItem="openChooseImg"
    />
    <div v-else class="w-full gap-4 flex flex-wrap">
      <selectComponent
        v-for="(item, index) in model"
        :key="index"
        :model="item"
        @chooseItem="openChooseImg"
        @deleteItem="deleteImg(index)"
      />
      <selectComponent
        v-if="
          model?.length < props.maxUpdateCount || props.maxUpdateCount === 0
        "
        @chooseItem="openChooseImg"
        @deleteItem="openChooseImg"
      />
    </div>

    <el-drawer v-model="drawer" title="媒体库" :size="appStore.drawerSize">
      <warning-bar title="点击“文件名/备注”可以编辑文件名或者备注内容。" />
      <div class="gva-btn-list gap-2">
        <upload-common :image-common="imageCommon" @on-success="getImageList" />
        <upload-image
          :image-url="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          @on-success="getImageList"
        />
        <el-input
          v-model="search.keyword"
          class="keyword"
          placeholder="请输入文件名或备注"
        />
        <el-button type="primary" icon="search" @click="getImageList">
          查询
        </el-button>
      </div>
      <div class="flex flex-wrap gap-4">
        <div v-for="(item, key) in picList" :key="key" class="w-40">
          <div
            class="w-40 h-40 border rounded overflow-hidden border-dashed border-gray-300 cursor-pointer relative group"
          >
            <el-image
              :key="key"
              :src="getUrl(item.url)"
              fit="cover"
              class="w-full h-full relative"
              @click="chooseImg(item.url)"
            >
              <template #error>
                <el-icon
                  v-if="isVideoExt(item.url || '')"
                  :size="32"
                  class="absolute top-[calc(50%-16px)] left-[calc(50%-16px)]"
                >
                  <VideoPlay />
                </el-icon>
                <video
                  v-if="isVideoExt(item.url || '')"
                  class="w-full h-full object-cover"
                  muted
                  preload="metadata"
                  @click="chooseImg(item.url)"
                >
                  <source :src="getUrl(item.url) + '#t=1'" />
                  您的浏览器不支持视频播放
                </video>
                <div
                  v-else
                  class="w-full h-full object-cover flex items-center justify-center"
                >
                  <el-icon :size="32">
                    <icon-picture />
                  </el-icon>
                </div>
              </template>
            </el-image>
            <div
              class="absolute -right-1 top-1 w-8 h-8 group-hover:inline-block hidden"
              @click="deleteCheck(item)"
            >
              <el-icon :size="16"><CircleClose /></el-icon>
            </div>
          </div>
          <div
            class="overflow-hidden text-nowrap overflow-ellipsis text-center w-full"
            @click="editFileNameFunc(item)"
          >
            {{ item.name }}
          </div>
        </div>
      </div>
      <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :total="total"
        :style="{ 'justify-content': 'center' }"
        layout="total, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </el-drawer>
  </div>
</template>

<script setup>
  import { getUrl, isVideoExt } from '@/utils/image'
  import { ref } from 'vue'
  import {
    getFileList,
    editFileName,
    deleteFile
  } from '@/api/fileUploadAndDownload'
  import UploadImage from '@/components/upload/image.vue'
  import UploadCommon from '@/components/upload/common.vue'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { Picture as IconPicture } from '@element-plus/icons-vue'
  import selectComponent from '@/components/selectImage/selectComponent.vue'
  import { useAppStore } from "@/pinia";

  const appStore = useAppStore()

  const imageUrl = ref('')
  const imageCommon = ref('')

  const search = ref({})
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

  const editFileNameFunc = async (row) => {
    ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      inputPattern: /\S/,
      inputErrorMessage: '不能为空',
      inputValue: row.name
    })
      .then(async ({ value }) => {
        row.name = value
        const res = await editFileName(row)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '编辑成功!'
          })
          getImageList()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消修改'
        })
      })
  }

  const drawer = ref(false)
  const picList = ref([])

  const imageTypeList = ['png', 'jpg', 'jpeg', 'gif', 'bmp', 'webp']
  const videoTyteList = [
    'mp4',
    'avi',
    'rmvb',
    'rm',
    'asf',
    'divx',
    'mpg',
    'mpeg',
    'mpe',
    'wmv',
    'mkv',
    'vob'
  ]

  const listObj = {
    image: imageTypeList,
    video: videoTyteList
  }

  const chooseImg = (url) => {
    if (props.fileType) {
      const typeSuccess = listObj[props.fileType].some((item) => {
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
    if (props.multiple) {
      model.value.push(url)
    } else {
      model.value = url
    }
    drawer.value = false
  }

  const openChooseImg = async () => {
    if (model.value && !props.multiple) {
      model.value = ''
      return
    }
    await getImageList()
    drawer.value = true
  }

  const getImageList = async () => {
    const res = await getFileList({
      page: page.value,
      pageSize: pageSize.value,
      ...search.value
    })
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
    })
      .then(async () => {
        const res = await deleteFile(item)
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '删除成功!'
          })
          getImageList()
        }
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '已取消删除'
        })
      })
  }
</script>
