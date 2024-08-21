<template>
  <div>
    <el-upload
        class="upload-demo"
        drag
        multiple
        :before-upload="handleBeforeUpload"
        :auto-upload="true"
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">{{title}} <em>点击上传</em></div>
      <div v-for="(v,i) in uploadedFiles">{{ v }}</div>
<!--      <div v-if="uploadedFiles.length>0">{{ uploadedFiles[0] }}</div>-->
    </el-upload>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';
import { UploadFilled } from '@element-plus/icons-vue';
import * as qiniu from 'qiniu-js';
import { onMounted, ref } from 'vue';
import {getUploadToken} from "@/api/biz_apphub/oss";

const props = defineProps({
  uploadedFiles: {
    type: Array,
    default: () => []
  },
  title:{
    type: String,
    default:()=>"请上传文件"
  }
});

const fileList = ref(props.uploadedFiles);

const handleBeforeUpload = async function (file) {
  const token = await getUploadToken();
  const config = {
    useCdnDomain: false,
    region: qiniu.region.z2
  };

  let ossPath = 'web/beiluo/' + new Date().getTime() + '/' + file.name;
  const observable = qiniu.upload(file, ossPath, token, null, config);

  observable.subscribe({
    next: (res) => {
      console.log('上传进度：', res.total.percent);
    },
    error: (err) => {
      console.error('上传失败：', err);
    },
    complete: (res) => {
      fileList.value.push(res['key']);
      console.log('上传成功：', res);
    }
  });

  return false;
};

// const getUploadToken = async function () {
//   let res = await axios.get('/webapi/api/v1/upload/token');
//   console.log(res.data.data.token);
//   return res.data.data.token;
// };

onMounted(() => {
});

function info() {
  console.log(fileList.value);
}
</script>