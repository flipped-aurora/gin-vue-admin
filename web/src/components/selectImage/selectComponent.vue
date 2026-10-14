<template>
  <div
      class="w-40 h-40 relative rounded border border-dashed border-gray-300 cursor-pointer group"
      :class="rounded ? 'rounded-full' : ''"
  >
    <div class="w-full h-full overflow-hidden" :class="rounded ? 'rounded-full' : ''">
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

      <el-image
          v-if="model && !isVideoExt(model)"
          class="w-full h-full"
          :src="imgUrl"
          :preview-src-list="srcList"
          fit="cover"
      />
      <template v-else>
      <div
          class="text-gray-600 group-hover:bg-gray-200 group-hover:opacity-60 w-full h-full flex justify-center items-center"
          @click="chooseItem"
      >
        <el-icon>
          <plus />
        </el-icon>
      </div>
      <div id="paste-zone" tabindex="0"
           class="absolute z-10 w-16 h-16 -mt-18 ml-22 text-xs flex items-center border border-dashed justify-center p-0.5 rounded border-gray-300  outline-none transition-colors focus:border-blue-400 focus:text-blue-400 text-gray-500"
           @click="(e) => e.currentTarget.focus()"
           @paste="onPaste"
      >
        先点击这里，然后粘贴(Ctrl / ⌘ + V)
      </div>
      </template>
    </div>
    <!-- 删除按钮在外层容器中 -->
    <div
        v-if="model"
        class="right-0 top-0 hidden text-gray-400 group-hover:flex justify-center items-center absolute z-10"
        @click="deleteItem"
    >
      <el-icon :size="24">
        <CircleCloseFilled />
      </el-icon>
    </div>
  </div>
</template>
<script setup>
  import { getUrl, isVideoExt } from '@/utils/image'
  import { CircleCloseFilled, Plus } from '@element-plus/icons-vue'
  import { computed } from 'vue'

  const props = defineProps({
    model: {
      default: '',
      type: String
    },
    rounded: {
      default: false,
      type: Boolean
    }
  })

  const emits = defineEmits(['chooseItem', 'deleteItem', 'pasteFile'])

  const chooseItem = () => {
    emits('chooseItem')
  }

  const deleteItem = () => {
    emits('deleteItem')
  }

  const imgUrl = computed(() => {
    return getUrl(props.model)
  })

  const srcList = computed(() => {
    return imgUrl.value ? [imgUrl.value] : []
  })

  // 粘贴图片：必须在事件回调内同步提取文件，异步后 clipboardData 会失效
  const onPaste = (e) => {
    const clipboardData = e.clipboardData
    const files = []
    if (clipboardData?.items) {
      for (let i = 0; i < clipboardData.items.length; i++) {
        const item = clipboardData.items[i]
        if (item.kind === 'file' && item.type.startsWith('image/')) {
          const file = item.getAsFile()
          if (file) files.push(file)
        }
      }
    }
    // items 为空时兜底（部分场景只有 files）
    if (!files.length && clipboardData?.files) {
      for (const file of clipboardData.files) {
        if (file.type.startsWith('image/')) files.push(file)
      }
    }
    if (!files.length) return // 非图片内容交给浏览器默认行为
    e.preventDefault()
    files.forEach((file) => emits('pasteFile', file))
  }
</script>
