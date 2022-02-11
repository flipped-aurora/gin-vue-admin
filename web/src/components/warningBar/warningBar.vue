<template>
  <div
    class="warning-bar"
  >
    <el-icon>
      <warning-filled />
    </el-icon>
    <span>
      {{ title }}
    </span>
  </div>
</template>
<script setup>
import { WarningFilled } from '@element-plus/icons-vue'
defineProps({
  title: {
    type: String,
    default: ''
  },
})
</script>
<style lang="scss" scoped>
.warning-bar{
    background-color: #FFF5ED;
    font-size: 14px;
    padding: 6px 14px;
    display: flex;
    align-items: center;
    border-radius: 2px;
    .el-icon{
      font-size: 18px;
      color: #ED6A0C;
    }
    margin-bottom: 12px;
    span{
        line-height: 22px;
        color:#F67207;
        margin-left: 8px;
    }
}
</style>
