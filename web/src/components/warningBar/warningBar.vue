<template>
  <div class="warning-bar">
    <el-icon>
      <warning-filled />
    </el-icon>
    <span>
      {{ title }}
    </span>
  </div>
</template>
<script>
import { WarningFilled } from '@element-plus/icons'

export default {
  components: {
    [WarningFilled.name]: WarningFilled
  },
  props: {
    title: {
      type: String,
      default: ''
    }
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
    .el-icon{
        color: #ED6A0C;
    }
    margin-bottom: 22px;
    span{
        color:#5D5D5E;
        margin-left: 8px;
        line-height: 18px;
    }
}
</style>
