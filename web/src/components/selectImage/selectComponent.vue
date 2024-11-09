<template>
  <div
    class="w-40 h-40 relative rounded border border-dashed border-gray-300 overflow-hidden cursor-pointer group"
  >
    <el-icon
      v-if="isVideoExt(model || '')"
      :size="32"
      class="absolute top-[calc(50%-16px)] left-[calc(50%-16px)]"
    >
      <VideoPlay />
    </el-icon>
    <video
      v-if="isVideoExt(model || '')"
      class="w-full h-full object-cover"
      muted
      preload="metadata"
    >
      <source :src="getUrl(model) + '#t=1'" />
    </video>

    <img
      v-if="model && !isVideoExt(model)"
      class="w-full h-full"
      :src="getUrl(model)"
      alt="图片"
    />
    <div
      v-if="model"
      class="left-0 top-0 hidden text-gray-600 group-hover:bg-gray-600 group-hover:bg-opacity-30 w-full h-full group-hover:flex justify-center items-center absolute z-10"
      @click="deleteItem"
    >
      <el-icon>
        <delete />
      </el-icon>
      删除
    </div>
    <div
      v-else
      class="text-gray-600 group-hover:bg-gray-400 w-full h-full flex justify-center items-center"
      @click="chooseItem"
    >
      <el-icon>
        <plus />
      </el-icon>
      上传
    </div>
  </div>
</template>
<script setup>
  import { getUrl, isVideoExt } from '@/utils/image'
  import { Delete, Plus } from '@element-plus/icons-vue'

  defineProps({
    model: {
      default: '',
      type: String
    }
  })

  const emits = defineEmits(['chooseItem', 'deleteItem'])

  const chooseItem = () => {
    emits('chooseItem')
  }

  const deleteItem = () => {
    emits('deleteItem')
  }
</script>
