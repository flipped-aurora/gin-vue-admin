<template>
  <div>
    <el-button type="primary" icon="iphone" @click="createQrCode"> 扫码上传</el-button>
  </div>

  <el-dialog v-model="dialogVisible" title="扫码上传" width="320px" :show-close="false" append-to-body :close-on-click-modal="false"
             draggable
  >
    <div class="m-2">
      <vue-qr :logoSrc="logoSrc"
              :size="291"
              :margin="0"
              :autoColor="true"
              :dotScale="1"
              :text="codeUrl"
              colorDark="green"
              colorLight="white"
              ref="qrcode"
      />
    </div>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="onFinished">完成上传</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import logoSrc from '@/assets/logo.png'
import vueQr from 'vue-qr/src/packages/vue-qr.vue'
import { ref } from 'vue'
import { createScanUploadToken } from '@/api/fileUploadAndDownload'

defineOptions({
  name: 'QRCodeUpload'
})

const emit = defineEmits(['on-success'])

const props = defineProps({
  classId: {
    type: Number,
    default: 0
  }
})

const dialogVisible = ref(false)
const codeUrl = ref('')

const createQrCode = async() => {
  const res = await createScanUploadToken(props.classId)
  if (res.code !== 0) return
  const local = window.location
  const query = new URLSearchParams({ uploadToken: res.data.token })
  codeUrl.value = `${local.protocol}//${local.host}/#/scanUpload?${query.toString()}`
  dialogVisible.value = true
}

const onFinished = () => {
  dialogVisible.value = false
  codeUrl.value = ''
  emit('on-success', '')
}
</script>
