<template>
  <el-dialog v-model="dialogTableVisible" :show-close="false" top="40px" :width="width" destroy-on-close :fullscreen="dialogConfig.fullscreen" append-to-body custom-class="sys-dialog" @close="handleClose">
    <template #header="{ close, titleId, titleClass }">
      <div class="sys-dialog-head">
        <h4 :id="titleId" :class="titleClass">{{ title }}</h4>
        <div>
          <el-button :icon="FullScreen" link type="warning" @click="handleCheckfullScreen" />
          <el-button :icon="Close" link  @click="close" />
        </div>
      </div>
    </template>
    <div v-if="alertTitle" class="sys-dialog-alert">
      <el-alert :title="alertTitle" :type="alertType" show-icon :closable="false" />
    </div>
    <div class="sys-dialog-body" :style="`height : ${height ? height : 'calc(100vh - 300px);'}`">
      <slot name="body" />
    </div>
    <div v-if="slotActions" class="sys-dialog-actions">
      <slot name="actions" />
    </div>
  </el-dialog>
</template>
<script setup>
import { Close, FullScreen } from '@element-plus/icons-vue'
import { reactive, ref, watch, useSlots } from 'vue'
const dialogTableVisible = ref(false)
const emit = defineEmits(['close'])

const slotActions = !!useSlots().actions
const prop = defineProps({
  value: {
    type: Boolean,
    default: false,
  },
  title: {
    type: String,
    default: 'title',
  },
  width: {
    type: [String, Number],
    default: '60%',
  },
  height: {
    type: [String],
    default: '',
  },
  alertTitle: {
    type: String,
    default: '',
  },
  alertType: {
    type: String,
    default: 'warning',
  }
})
const dialogConfig = reactive({
  fullscreen: false,
})

watch(
  () => prop.value,
  (val) => {
    dialogTableVisible.value = val
  }, {
    immediate: true,
  })

const handleClose = () => {
  dialogTableVisible.value = false
  emit('close')
}

const handleCheckfullScreen = () => {
  dialogConfig.fullscreen = !dialogConfig.fullscreen
}

</script>
<style lang="scss" scoped>

.sys-dialog.el-dialog {
  border-radius: 8px;

  &.is-fullscreen {
    border-radius: 0;
  }
  .sys-dialog-head{
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .el-dialog__header {
    height: 40px !important;

    .el-dialog__title {
      padding-left: 14px;
      position: relative;

      &::after {
        content: '';
        position: absolute;
        left: 0;
        top: 0;
        width: 4px;
        border-radius: 4px;
        height: 100%;
        background-color: var(--el-color-primary);
      }
    }
  }

  .el-dialog__body {
    padding: 10px;
  }
}

.sys-dialog {

  &-alert {
    padding-bottom: 10px;
  }

  &-body {
    height: calc(100vh - 300px);
    overflow-y: auto;
    overflow-x: hidden;
  }

  &-actions {
    padding: 10px;
    margin: 0;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    border-top: 1px solid rgba(0, 0, 0, .08);
  }
}
</style>
