<template>
  <div>
    <el-upload
      v-model:file-list="fileList"
      multiple
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload?noSave=1`"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :on-remove="uploadRemove"
      :before-upload="checkFile"
      :show-file-list="true"
      :limit="limit"
      :accept="accept"
      class="upload-btn"
      :headers="{'x-token': token}"
    >
      <el-button type="primary"> 上传文件 </el-button>
    </el-upload>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { getBaseUrl } from '@/utils/format'
  import { useUserStore } from "@/pinia";

  defineOptions({
    name: 'UploadCommon'
  })

  const props = defineProps({
    limit: {
      type: Number,
      default: 3
    },
    accept: {
      type: String,
      default: ''
    },
    fileSize: {
      type: Number,
      default: 8
    }
  })

  const userStore = useUserStore()

  const token = userStore.token

  const fullscreenLoading = ref(false)

  const model = defineModel({ type: Array })

  const fileList = ref(model.value || [])

  const emits = defineEmits(['on-success', 'on-error'])

  const uploadSuccess = (res) => {
    const { data, code } = res
    if (code !== 0) {
      ElMessage({
        type: 'error',
        message: '上传失败' + res.msg
      })
      fileList.value.pop()
      return
    }
    model.value.push({
      name: data.file.name,
      url: data.file.url
    })
    emits('on-success', res)
  }

  const uploadRemove = (file) => {
    const index = model.value.findIndex(item => item.name === file.name)
    if (index > -1) {
      model.value.splice(index, 1)
      fileList.value = model.value
    }
  }

  const uploadError = (err) => {
    ElMessage({
      type: 'error',
      message: '上传失败'
    })
    fullscreenLoading.value = false
    emits('on-error', err)
  }

  const checkFile = (file) => {
    fullscreenLoading.value = true
    const isFileSize = file.size / 1024 / 1024 < props.fileSize // 文件大小
    let pass = true
    if (!isFileSize) {
      ElMessage.error(`上传文件大小不能超过 ${props.fileSize}MB`)
      fullscreenLoading.value = false
      pass = false
    }

    // 获取accept属性
    const acceptPattern = props.accept
    if (acceptPattern) {
      // 将accept字符串转换为正则表达式进行验证
      // 例如: "image/*" 需要转换为匹配所有image类型
      const pattern = new RegExp(
          acceptPattern
              .split(',')
              .map(type => type.trim().replace('*', '.*'))
              .join('|')
      )

      if (!pattern.test(file.type)) {
        ElMessage.error(`请上传${acceptPattern}格式的文件`)
        fullscreenLoading.value = false
        pass = false
      }
    }


    console.log('upload file check result: ', pass)

    return pass
  }
</script>
