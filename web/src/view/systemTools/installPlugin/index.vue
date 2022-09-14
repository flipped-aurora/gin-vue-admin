<template>
  <div>
    <el-upload
      class="upload-demo"
      drag
      :action="`${path}/autoCode/installPlugin`"
      :headers="{'x-token':userStore.token}"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleSuccess"
      name="plug"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        拖拽或<em>点击上传</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          请把安装包的zip拖拽至此处上传
        </div>
      </template>
    </el-upload>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import { ElMessage } from 'element-plus'
const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_API)

const handleSuccess = (res) => {
  if (res.code === 0) {
    let msg = ``
    res.data && res.data.forEach((item, index) => {
      msg += `${index + 1}.${item.msg}\n`
    })
    alert(msg)
  } else {
    ElMessage.error(res.msg)
  }
}
</script>
