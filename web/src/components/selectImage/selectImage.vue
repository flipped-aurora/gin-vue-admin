<template>
  <div
    v-if="!multiple"
    class="update-image"
    :style="{
      'background-image': `url(${getUrl(modelValue)})`,
      'position': 'relative',
    }"
  >
    <el-icon
      v-if="isVideoExt(modelValue || '')"
      :size="32"
      class="video video-icon"
      style=""
    ><VideoPlay /></el-icon>
    <video
      v-if="isVideoExt(modelValue || '')"
      class="avatar video-avatar video"
      muted
      preload="metadata"
      style=""
      @click="openChooseImg"
    >
      <source :src="getUrl(modelValue) + '#t=1'">
    </video>
    <span
      v-if="modelValue"
      class="update"
      style="position: absolute;"
      @click="openChooseImg"
    >
      <el-icon>
        <delete />
      </el-icon>
      删除</span>
    <span
      v-else
      class="update text-gray-600"
      @click="openChooseImg"
    >
      <el-icon>
        <plus />
      </el-icon>
      上传</span>
  </div>
  <div
    v-else
    class="multiple-img"
  >
    <div
      v-for="(item, index) in multipleValue"
      :key="index"
      class="update-image"
      :style="{
        'background-image': `url(${getUrl(item)})`,
        'position': 'relative',
      }"
    >
      <el-icon
        v-if="isVideoExt(item || '')"
        :size="32"
        class="video video-icon"
      ><VideoPlay /></el-icon>
      <video
        v-if="isVideoExt(item || '')"
        class="avatar video-avatar video"
        muted
        preload="metadata"
        @click="deleteImg(index)"
      >
        <source :src="getUrl(item) + '#t=1'">
      </video>
      <span
        class="update"
        style="position: absolute;"
        @click="deleteImg(index)"
      >
        <el-icon>
          <delete />
        </el-icon>
        删除</span>
    </div>
    <div class="add-image">
      <span
        class="update  text-gray-600"
        @click="openChooseImg"
      >
        <el-icon>
          <Plus />
        </el-icon>
        上传</span>
    </div>
  </div>

  <el-drawer
    v-model="drawer"
    title="媒体库"
    size="650px"
  >
    <warning-bar
      title="点击“文件名/备注”可以编辑文件名或者备注内容。"
    />
    <div class="gva-btn-list">
      <upload-common
        v-model:imageCommon="imageCommon"
        class="upload-btn-media-library"
        @on-success="getImageList"
      />
      <upload-image
        v-model:imageUrl="imageUrl"
        :file-size="512"
        :max-w-h="1080"
        class="upload-btn-media-library"
        @on-success="getImageList"
      />
      <el-form
        ref="searchForm"
        :inline="true"
        :model="search"
      >
        <el-form-item label="">
          <el-input
            v-model="search.keyword"
            class="keyword"
            placeholder="请输入文件名或备注"
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="getImageList"
          >查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="media">
      <div
        v-for="(item,key) in picList"
        :key="key"
        class="media-box"
      >
        <div class="header-img-box-list">
          <el-image
            :key="key"
            :src="getUrl(item.url)"
            fit="cover"
            style="width: 100%;height: 100%;"
            @click="chooseImg(item.url)"
          >
            <template #error>
              <el-icon
                v-if="isVideoExt(item.url || '')"
                :size="32"
                class="video video-icon"
              ><VideoPlay /></el-icon>
              <video
                v-if="isVideoExt(item.url || '')"
                class="avatar video-avatar video"
                muted
                preload="metadata"
                @click="chooseImg(item.url)"
              >
                <source :src="getUrl(item.url) + '#t=1'">
                您的浏览器不支持视频播放
              </video>
              <div
                v-else
                class="header-img-box-list"
              >
                <el-icon class="lost-image">
                  <icon-picture />
                </el-icon>
              </div>
            </template>
          </el-image>
        </div>
        <div
          class="img-title"
          @click="editFileNameFunc(item)"
        >{{ item.name }}</div>
      </div>
    </div>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :total="total"
      :style="{'justify-content':'center'}"
      layout="total, prev, pager, next, jumper"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
  </el-drawer>
</template>

<script setup>

import { getUrl, isVideoExt } from '@/utils/image'
import { onMounted, ref } from 'vue'
import { getFileList, editFileName } from '@/api/fileUploadAndDownload'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Delete, FolderAdd, Plus, Picture as IconPicture } from '@element-plus/icons-vue'

const imageUrl = ref('')
const imageCommon = ref('')

const search = ref({})
const page = ref(1)
const total = ref(0)
const pageSize = ref(20)

const props = defineProps({
  modelValue: {
    type: [String, Array],
    default: ''
  },
  multiple: {
    type: Boolean,
    default: false
  },
  fileType: {
    type: String,
    default: ''
  }
})

const multipleValue = ref([])

onMounted(() => {
  if (props.multiple) {
    multipleValue.value = props.modelValue
  }
})

const emits = defineEmits(['update:modelValue'])

const deleteImg = (index) => {
  multipleValue.value.splice(index, 1)
  emits('update:modelValue', multipleValue.value)
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getImageList()
}

const handleCurrentChange = (val) => {
  page.value = val
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
    // console.log(row)
    const res = await editFileName(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '编辑成功!',
      })
      getImageList()
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

const imageTypeList = ['png', 'jpg', 'jpge', 'gif', 'bmp', 'webp']
const videoTyteList = ['mp4', 'avi', 'rmvb', 'rm', 'asf', 'divx', 'mpg', 'mpeg', 'mpe', 'wmv', 'mkv', 'vob']

const listObj = {
  image: imageTypeList,
  video: videoTyteList
}

const chooseImg = (url) => {
  console.log(url)
  if (props.fileType) {
    const typeSuccess = listObj[props.fileType].some(item => {
      if (url.includes(item)) {
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
    multipleValue.value.push(url)
    emits('update:modelValue', multipleValue.value)
  } else {
    emits('update:modelValue', url)
  }
  drawer.value = false
}
const openChooseImg = async() => {
  if (props.modelValue && !props.multiple) {
    emits('update:modelValue', '')
    return
  }
  await getImageList()
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

</script>

<style scoped lang="scss">

.multiple-img{
  display: flex;
  gap:8px;
  width: 100%;
  flex-wrap: wrap;
}

.add-image{
  width: 120px;
  height: 120px;
  line-height: 120px;
  display: flex;
  justify-content: center;
  border-radius: 20px;
  border: 1px dashed #ccc;
  background-size: cover;
  cursor: pointer;
}

.update-image {
  cursor: pointer;
  width: 120px;
  height: 120px;
  line-height: 120px;
  display: flex;
  justify-content: center;
  border-radius: 20px;
  border: 1px dashed #ccc;
   background-repeat: no-repeat;
   background-size: cover;
   position: relative;
  &:hover {
    color: #fff;
    background: linear-gradient(
            to bottom,
            rgba(255, 255, 255, 0.15) 0%,
            rgba(0, 0, 0, 0.15) 100%
    ),
    radial-gradient(
            at top center,
            rgba(255, 255, 255, 0.4) 0%,
            rgba(0, 0, 0, 0.4) 120%
    )
    #989898;
    background-blend-mode: multiply, multiply;
    background-size: cover;
    .update {
      color: #fff;
    }
    .video {
      opacity: 0.2;
    }
  }

  .video-icon {
    position: absolute; left: calc(50% - 16px); top: calc(50% - 16px);
  }
  video {
    object-fit: cover; max-width:100%; border-radius: 20px;
  }
  .update {
    height: 120px;
    width: 120px;
    text-align: center;
    color: transparent;
    position: absolute;
  }
}

.upload-btn-media-library {
  margin-left: 20px;
}

.media {
  display: flex;
  flex-wrap: wrap;

  .media-box {
    width: 120px;
    margin-left: 20px;

    .img-title {
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      line-height: 36px;
      text-align: center;
      cursor: pointer;
    }

    .header-img-box-list {
      width: 120px;
      height: 120px;
      border: 1px dashed #ccc;
      border-radius: 8px;
      text-align: center;
      line-height: 120px;
      cursor: pointer;
      overflow: hidden;
      .el-image__inner {
        max-width: 120px;
        max-height: 120px;
        vertical-align: middle;
        width: unset;
        height: unset;
      }

      .el-image {
        position: relative;
      }
      .video-icon {
        position: absolute; left: calc(50% - 16px); top: calc(50% - 16px);
      }
      video {
        object-fit: cover; max-width:100%; min-height: 100%; border-radius: 8px;
      }
    }
  }
}
</style>
