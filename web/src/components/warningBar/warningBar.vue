<template>
  <div
    class="warning-bar"
    :class="href&&'can-click'"
    @click="open"
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
const prop = defineProps({
  title: {
    type: String,
    default: ''
  },
  href: {
    type: String,
    default: ''
  }
})

const open = () => {
  if (prop.href) {
    window.open(prop.href)
  }
}
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
.can-click{
  cursor: pointer;
}
</style>
