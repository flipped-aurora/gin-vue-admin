<template>
  <el-upload
      ref="uploadRef"
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
      accept="image/*"
      :show-file-list="false"
      :auto-upload="false"
      :data="{'classId': props.classId}"
      :on-success="handleImageSuccess"
      :on-change="handleFileChange"
      :headers="{'x-token': token}"
  >
    <el-button type="primary" icon="crop"> 裁剪上传</el-button>
  </el-upload>

  <el-dialog v-model="dialogVisible" title="图片裁剪" width="1200px" append-to-body @close="dialogVisible = false" :close-on-click-modal="false" draggable>
    <div class="flex gap-[30px] h-[600px]">
      <!-- 左侧编辑区 -->
      <div class="flex flex-col flex-1">
        <div class="flex-1 bg-[#f8f8f8] rounded-lg overflow-hidden">
          <VueCropper
              ref="cropperRef"
              :img="imgSrc"
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
              :maxImgSize="1200"
              :original="true"
              @realTime="handleRealTime"
          ></VueCropper>
        </div>

        <!-- 工具栏 -->
        <div class="mt-[20px] flex items-center p-[10px] bg-white rounded-lg shadow-[0_2px_12px_rgba(0,0,0,0.1)]">
          <el-button-group>
            <el-tooltip content="向左旋转">
              <el-button @click="rotate(-90)" :icon="RefreshLeft" />
            </el-tooltip>
            <el-tooltip content="向右旋转">
              <el-button @click="rotate(90)" :icon="RefreshRight" />
            </el-tooltip>
            <el-button :icon="Plus" @click="changeScale(1)"></el-button>
            <el-button :icon="Minus" @click="changeScale(-1)"></el-button>
          </el-button-group>


          <el-select v-model="currentRatio" placeholder="选择比例" class="w-32 ml-4" @change="onCurrentRatio">
            <el-option v-for="(item, index) in ratioOptions" :key="index" :label="item.label" :value="index" />
          </el-select>
        </div>
      </div>

      <!-- 右侧预览区 -->
      <div class="w-[340px]">
        <div class="bg-white p-5 rounded-lg shadow-[0_2px_12px_rgba(0,0,0,0.1)]">
          <div class="mb-[15px] text-gray-600">裁剪预览</div>
          <div class="bg-white p-5 rounded-lg shadow-[0_2px_12px_rgba(0,0,0,0.1)]"
               :style="{'width': previews.w + 'px', 'height': previews.h + 'px'}"
          >
            <div class="w-full h-full relative overflow-hidden">
              <img :src="previews.url" :style="previews.img" alt="" class="max-w-none absolute transition-all duration-300 ease-in-out image-render-pixelated origin-[0_0]" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleUpload" :loading="uploading"> {{ uploading ? '上传中...' : '上 传' }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
import { ElMessage } from 'element-plus'
import { RefreshLeft, RefreshRight, Plus, Minus } from '@element-plus/icons-vue'
import 'vue-cropper/dist/index.css'
import { VueCropper } from 'vue-cropper'
import { getBaseUrl } from '@/utils/format'
import { useUserStore } from "@/pinia";

defineOptions({
  name: 'CropperImage'
})

const emit = defineEmits(['on-success'])

const props = defineProps({
  classId: {
    type: Number,
    default: 0
  }
})

const uploadRef = ref(null)
// 响应式数据
const dialogVisible = ref(false)
const imgSrc = ref('')
const cropperRef = ref(null)
const { proxy } = getCurrentInstance()
const previews = ref({})
const uploading = ref(false)

// 缩放控制
const changeScale = (value) => {
  proxy.$refs.cropperRef.changeScale(value)
}

// 比例预设
const ratioOptions = ref([
  { label: '1:1', value: [1, 1] },
  { label: '16:9', value: [16, 9] },
  { label: '9:16', value: [9, 16] },
  { label: '4:3', value: [4, 3] },
  { label: '自由比例', value: [] }
])

const fixedNumber = ref([1, 1])
const cropWidth = ref(300)
const cropHeight = ref(300)

const fixedRatio = ref(false)
const currentRatio = ref(4)
const onCurrentRatio = () => {
  fixedNumber.value = ratioOptions.value[currentRatio.value].value
  switch (currentRatio.value) {
    case 0:
      cropWidth.value = 300
      cropHeight.value = 300
      fixedRatio.value = true
      break
    case 1:
      cropWidth.value = 300
      cropHeight.value = 300 * 9 / 16
      fixedRatio.value = true
      break
    case 2:
      cropWidth.value = 300 * 9 / 16
      cropHeight.value = 300
      fixedRatio.value = true
      break
    case 3:
      cropWidth.value = 300
      cropHeight.value = 300 * 3 / 4
      fixedRatio.value = true
      break
    default:
      cropWidth.value = 300
      cropHeight.value = 300
      fixedRatio.value = false
  }
}

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

  const reader = new FileReader()
  reader.onload = (e) => {
    imgSrc.value = e.target.result
    dialogVisible.value = true
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

// 实时预览
const handleRealTime = (data) => {
  previews.value = data
  //console.log(data)
}

// 上传处理
const handleUpload = () => {
  uploading.value = true
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
    setTimeout(() => {
      uploading.value = false
      dialogVisible.value = false
      previews.value = {}
      ElMessage.success('上传成功')
      emit('on-success', data.url)
    }, 1000)
  }
}

</script>

<style scoped>
:deep(.vue-cropper) {
  background: transparent;
}
</style>
