<template>
  <div class="profile-avatar relative w-[120px] h-[120px] rounded-full overflow-hidden cursor-pointer group">
    <img
      v-if="modelValue"
      class="w-full h-full object-cover"
      :src="getUrl(modelValue)"
      alt="头像"
    />
    <div
      v-else
      class="w-full h-full flex flex-col items-center justify-center bg-gray-100 dark:bg-slate-700"
    >
      <el-icon class="text-2xl text-gray-400">
        <avatar />
      </el-icon>
      <span class="mt-2 text-sm text-gray-400">点击上传</span>
    </div>

    <div 
      class="absolute inset-0 bg-black/30 opacity-0 group-hover:opacity-100 flex items-center justify-center transition-all duration-300"
      @click="chooseFile"
    >
      <div class="text-center text-white">
        <el-icon class="text-2xl"><camera-filled /></el-icon>
        <div class="text-xs mt-1">更换头像</div>
      </div>
    </div>

    <input
      ref="fileInput"
      type="file"
      class="hidden"
      accept="image/*"
      @change="handleFileChange"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { getUrl } from '@/utils/image'
import service from '@/utils/request'

defineOptions({
  name: 'ProfileAvatar'
})

const { modelValue } = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue'])

const fileInput = ref(null)

const chooseFile = () => {
  fileInput.value.click()
}

const handleFileChange = async (e) => {
  const file = e.target.files[0]
  if (!file) return

  // 验证文件类型
  if (!file.type.includes('image/')) {
    ElMessage.error('请选择图片文件')
    return
  }

  // 验证文件大小（限制为2MB）
  if (file.size > 2 * 1024 * 1024) {
    ElMessage.error('图片大小不能超过2MB')
    return
  }

  try {
    const formData = new FormData()
    formData.append('file', file)
    
    const res = await service({
      url: '/fileUploadAndDownload/upload',
      method: 'post',
      data: formData,
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    if (res.code === 0 && res.data?.file?.url) {
      emit('update:modelValue', res.data.file.url)
      ElMessage.success('头像上传成功')
    } else {
      ElMessage.error(res.msg || '上传失败')
    }
  } catch {
    ElMessage.error('头像上传失败，请重试')
  } finally {
    // 清空input，确保可以重复选择同一文件
    e.target.value = ''
  }
}
</script>

<style lang="scss" scoped>
.profile-avatar {
  img {
    @apply transition-transform duration-300;
  }

  &:hover {
    img {
      @apply scale-110;
    }
  }
}
</style> 