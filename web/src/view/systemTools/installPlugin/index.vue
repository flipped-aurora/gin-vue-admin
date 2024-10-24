<template>
    <div class="gva-form-box">
    <el-upload
      drag
      :action="`${getBaseUrl()}/autoCode/installPlugin`"
      :show-file-list="false"
      :on-success="handleSuccess"
      :on-error="handleSuccess"
      name="plug"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">
        {{ t('view.systemTools.installPlugin.dragOrClickUpload') }}
      </div>
      <template #tip>
        <div class="el-upload__tip">
          {{ t('view.systemTools.installPlugin.uploadPackage') }}
        </div>
      </template>
    </el-upload>
  </div>
</template>

<script setup>
import { ElMessage } from 'element-plus'
import { getBaseUrl } from '@/utils/format'

import { useI18n } from 'vue-i18n'

const { t } = useI18n()

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
