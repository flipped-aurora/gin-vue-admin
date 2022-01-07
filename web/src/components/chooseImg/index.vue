<template>
  <el-drawer v-model="drawer" title="媒体库">
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
