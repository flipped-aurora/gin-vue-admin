<template>
  <div style="width: 100%">
    <el-upload
      ref="uploadRef"
      multiple
      :action="`${path}/fileUploadAndDownload/upload?noSave=1`"
      :headers="{ 'x-token': userStore.token }"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :before-upload="beforeUpload"
      :on-remove="uploadRemove"
      :show-file-list="true"
      :file-list="fileList"
      accept=".zip"
      class="upload-btn"
    >
      <el-button type="primary">上传文件</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { ref, watch, onBeforeUnmount } from "vue";
import { ElMessage } from "element-plus";
import { useUserStore } from "@/pinia/modules/user";

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => [],
  },
});
const uploadRef = ref();
const path = ref(import.meta.env.VITE_BASE_API);
const userStore = useUserStore();
const fullscreenLoading = ref(false);

const fileList = ref(props.modelValue);

const emits = defineEmits(["update:modelValue"]);

watch(fileList.value, (val) => {
  console.log("watch", val);
  emits("update:modelValue", val);
});

const uploadSuccess = (res) => {
  const { data } = res;
  if (data.file) {
    fileList.value.push({
      name: data.file.name,
      url: data.file.url,
    });
    fullscreenLoading.value = false;
  }
};

const uploadError = () => {
  ElMessage({
    type: "error",
    message: "上传失败",
  });
  fullscreenLoading.value = false;
};

const beforeUpload = (rawFile) => {
  const statue =
    rawFile.name.split(".")[rawFile.name.split(".").length - 1] === "zip";
  if (!statue) {
    ElMessage({
      type: "error",
      message: "上传失败,上传文件类型不为zip",
    });
  }

  return statue;
};
const uploadRemove = (uploadFile) => {
  const index = fileList.value.findIndex(
    (item) => item.url === uploadFile.response.data.file.url
  );
  fileList.value.splice(index, 1);
};
onBeforeUnmount(() => {
  uploadRef.value.abort();
});
</script>

<script>
export default {
  name: "UploadCommon",
  methods: {},
};
</script>
