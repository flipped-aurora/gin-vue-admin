<template>
  <div class="p-5 bg-white">
    <WarningBar title="目前只支持标准插件（通过插件模板生成的标准目录插件），非标准插件请自行打包" />
    <div class="flex items-center gap-3">
      <el-input
        v-model="plugName"
        placeholder="插件模板处填写的【插件名】"
      />
      <el-button
        type="primary"
        @click="pubPlugin"
      >打包插件</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { pubPlug } from '@/api/autoCode.js'
import { ElMessage } from 'element-plus'
const plugName = ref('')

const pubPlugin = async() => {
  const res = await pubPlug({ plugName: plugName.value })
  if (res.code === 0) {
    ElMessage.success(res.msg)
  }
}

</script>

