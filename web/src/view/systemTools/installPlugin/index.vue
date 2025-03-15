<template>
  <div class="gva-form-box">
    <el-upload
      drag
      :action="`${getBaseUrl()}/autoCode/installPlugin`"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleSuccess"
      :headers="{'x-token': token}"
      name="plug"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">拖拽或<em>点击上传</em></div>
      <template #tip>
        <div class="el-upload__tip">请把安装包的zip拖拽至此处上传</div>
      </template>
    </el-upload>
  </div>
</template>

<script setup>
  import { ElMessage } from 'element-plus'
  import { getBaseUrl } from '@/utils/format'
  import { useUserStore } from "@/pinia";

  const userStore = useUserStore()

  const token = userStore.token

  const handleSuccess = (res) => {
    if (res.code === 0) {
      let msg = ``
      res.data &&
        res.data.forEach((item, index) => {
          msg += `${index + 1}.${item.msg}\n`
        })
      alert(msg)
    } else {
      ElMessage.error(res.msg)
    }
  }
</script>
