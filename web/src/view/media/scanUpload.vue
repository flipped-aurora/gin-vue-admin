<template>
  <div class="flex justify-center w-full pt-2">
    <el-upload
        ref="uploadRef"
        class="h5-uploader"
        :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
        accept="image/*"
        :show-file-list="false"
        :auto-upload="false"
        :headers="{ 'x-token': token }"
        :data="{'classId': classId}"
        :on-success="handleImageSuccess"
        :on-change="handleFileChange"
    >
      <el-icon class="h5-uploader-icon"><Plus /></el-icon>
    </el-upload>
  </div>

  <div class="flex flex-col w-full h-auto p-0 pt-4">
    <!-- 左侧编辑区 -->
    <div class="flex-1 min-h-[60vh]">
      <div class="w-screen h-[calc(100vh-175px)] rounded">
        <template v-if="isCrop">
          <VueCropper
              ref="cropperRef"
              :img="imgSrc"
              mode="contain"
              outputType="jpeg"
              :autoCrop="true"
              :autoCropWidth="cropWidth"
              :autoCropHeight="cropHeight"
              :fixedBox="false"
              :fixed="fixedRatio"
              :fixedNumber="fixedNumber"
              :centerBox="true"
              :canMoveBox="true"
              :full="false"
              :maxImgSize="windowWidth"
              :original="true"
          ></VueCropper>
        </template>
        <template v-else>
          <div class="flex justify-center items-center w-full h-[calc(100vh-175px)]">
            <el-image v-if="imgSrc" :src="imgSrc" class="max-w-full max-h-full" mode="cover" />
          </div>
        </template>
      </div>
    </div>
  </div>
  <!-- 工具栏 -->
  <div class="toolbar">
    <el-button-group v-if="isCrop">
      <el-tooltip content="向左旋转">
        <el-button @click="rotate(-90)" :icon="RefreshLeft" />
      </el-tooltip>
      <el-tooltip content="向右旋转">
        <el-button @click="rotate(90)" :icon="RefreshRight" />
      </el-tooltip>
      <el-button :icon="Plus" @click="changeScale(1)"></el-button>
      <el-button :icon="Minus" @click="changeScale(-1)"></el-button>
    </el-button-group>


    <el-switch
        size="large"
        v-model="isCrop"
        inline-prompt
        active-text="裁剪"
        inactive-text="裁剪"
    />

    <el-button type="primary" @click="handleUpload" :loading="uploading"> {{ uploading ? '上传中...' : '上 传' }}
    </el-button>
  </div>


</template>

<script setup>
import { ref, getCurrentInstance, onMounted } from 'vue'
import { ElLoading, ElMessage } from 'element-plus'
import { RefreshLeft, RefreshRight, Plus, Minus } from '@element-plus/icons-vue'
import 'vue-cropper/dist/index.css'
import { VueCropper } from 'vue-cropper'
import { getBaseUrl } from '@/utils/format'
import { useRouter } from 'vue-router'
import { useUserStore } from "@/pinia";

defineOptions({
  name: 'scanUpload'
})

const classId = ref(0)
const token = ref('')
const isCrop = ref(false)

const windowWidth = ref(300)

// 获取屏幕宽度
const getWindowResize = function() {
  windowWidth.value = window.innerWidth
}

// 生命周期
onMounted(() => {
  getWindowResize()
  window.addEventListener('resize', getWindowResize)
})

const router = useRouter()
router.isReady().then(() => {
  let query = router.currentRoute.value.query
  //console.log(query)
  classId.value = query.id
  token.value = query.token
}).catch((err) => {
  console.log(err)
})

const uploadRef = ref(null)
// 响应式数据
const imgSrc = ref('')
const cropperRef = ref(null)
const { proxy } = getCurrentInstance()
const previews = ref({})
const uploading = ref(false)

// 缩放控制
const changeScale = (value) => {
  proxy.$refs.cropperRef.changeScale(value)
}

const fixedNumber = ref([1, 1])
const cropWidth = ref(300)
const cropHeight = ref(300)

const fixedRatio = ref(false)

// 文件处理
const handleFileChange = (file) => {
  const isImage = file.raw.type.includes('image')
  if (!isImage) {
    ElMessage.error('请选择图片文件')
    return
  }

  if (file.raw.size / 1024 / 1024 > 8) {
    ElMessage.error('文件大小不能超过8MB!')
    return false
  }

  const loading = ElLoading.service({
    lock: true,
    text: '请稍后',
    background: 'rgba(0, 0, 0, 0.7)',
  })

  const reader = new FileReader()
  reader.onload = (e) => {
    imgSrc.value = e.target.result
    loading.close()
  }
  reader.readAsDataURL(file.raw)
}

// 旋转控制
const rotate = (degree) => {
  if (degree === -90) {
    proxy.$refs.cropperRef.rotateLeft()
  } else {
    proxy.$refs.cropperRef.rotateRight()
  }
}

// 上传处理
const handleUpload = () => {
  uploading.value = true
  if(isCrop.value === false){
    uploadRef.value.submit()
    return true
  }
  proxy.$refs.cropperRef.getCropBlob((blob) => {
    try {
      const file = new File([blob], `${Date.now()}.jpg`, { type: 'image/jpeg' })
      uploadRef.value.clearFiles()
      uploadRef.value.handleStart(file)
      uploadRef.value.submit()

    } catch (error) {
      uploading.value = false
      ElMessage.error('上传失败: ' + error.message)
    }
  })
}

const handleImageSuccess = (res) => {
  const { data } = res
  if (data) {
    imgSrc.value = null
    uploading.value = false
    previews.value = {}
    ElMessage.success('上传成功')
  }
}

</script>

<style scoped>

/* 工具栏（固定在底部） */
.toolbar {
  @apply fixed bottom-0 m-0 rounded-none p-2.5 shadow-[0_-2px_10px_rgba(0,0,0,0.1)] z-[1000] flex justify-between w-screen bg-slate-900;

  /* 按钮组适配 */
  .el-button-group {
    @apply flex gap-2;

    .el-button {
      @apply p-2 w-10;
    }
  }
}

:deep(.vue-cropper) {
  @apply bg-transparent;
}

</style>

<style>
.h5-uploader .el-upload {
  @apply rounded cursor-pointer relative overflow-hidden;
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  transition: var(--el-transition-duration-fast);
}

.h5-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.h5-uploader-icon {
  @apply text-2xl text-gray-500 w-32 h-32 text-center;
}
</style>
