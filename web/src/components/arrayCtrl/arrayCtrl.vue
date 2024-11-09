<template>
  <div class="flex gap-2">
    <el-tag
      v-for="tag in modelValue"
      :key="tag"
      :closable="editable"
      :disable-transitions="false"
      @close="handleClose(tag)"
    >
      {{ tag }}
    </el-tag>
    <template v-if="editable">
      <el-input
        v-if="inputVisible"
        ref="InputRef"
        v-model="inputValue"
        class="w-20"
        size="small"
        @keyup.enter="handleInputConfirm"
        @blur="handleInputConfirm"
      />
      <el-button v-else class="button-new-tag" size="small" @click="showInput">
        + 新增
      </el-button>
    </template>
  </div>
</template>

<script setup>
  defineOptions({
    name: 'ArrayCtrl'
  })

  import { nextTick, ref } from 'vue'

  const inputValue = ref('')
  const inputVisible = ref(false)
  const InputRef = ref(null)

  const modelValue = defineModel()

  defineProps({
    editable: {
      type: Boolean,
      default: () => false
    }
  })

  const handleClose = (tag) => {
    modelValue.value.splice(modelValue.value.indexOf(tag), 1)
  }

  const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
      InputRef.value?.input?.focus()
    })
  }

  const handleInputConfirm = () => {
    if (inputValue.value) {
      modelValue.value.push(inputValue.value)
    }
    inputVisible.value = false
    inputValue.value = ''
  }
</script>
