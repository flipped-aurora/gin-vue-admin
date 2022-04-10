<template>
  <el-drawer v-model="drawer" title="媒体库">
    <div class="gva-btn-list">
      <upload-common
          v-model:imageCommon="imageCommon"
          class="upload-btn-media-library"
          @on-success="open"
      />
      <upload-image
          v-model:imageUrl="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          class="upload-btn-media-library"
          @on-success="open"
      />
    </div>
    <div class="media">
      <el-image
        v-for="(item,key) in picList"
        :key="key"
        class="header-img-box-list"
        :src="(item.url && item.url.slice(0, 4) !== 'http')?path+item.url:item.url"
        @click="chooseImg(item.url,target,targetKey)"
      >
        <template #error>
          <div class="header-img-box-list">
            <el-icon><picture /></el-icon>
          </div>
        </template>
      </el-image>
    </div>
  </el-drawer>
</template>

<script setup>
import { ref } from 'vue'
import { getFileList } from '@/api/fileUploadAndDownload'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'

const imageUrl = ref('')
const imageCommon = ref('')

const emit = defineEmits(['enterImg'])
defineProps({
  target: {
    type: Object,
    default: null
  },
  targetKey: {
    type: String,
    default: ''
  }
})

const drawer = ref(false)
const picList = ref([])
const path = ref(import.meta.env.VITE_BASE_API)

const chooseImg = (url, target, targetKey) => {
  if (target && targetKey) {
    target[targetKey] = url
  }
  emit('enterImg', url)
  drawer.value = false
}

const open = async() => {
  const res = await getFileList({ page: 1, pageSize: 9999 })
  picList.value = res.data.list
  drawer.value = true
}

defineExpose({ open })
</script>

<style lang="scss">
.upload-btn-media-library{margin-left: 20px;}
.media{
  display:flex;
  flex-wrap:wrap;
  .header-img-box-list {
    margin-top: 20px;
    width: 120px;
    margin-left: 20px;
    height: 120px;
    border: 1px dashed #ccc;
    border-radius: 20px;
    text-align: center;
    line-height: 180px;
    cursor: pointer;
}
}

</style>
