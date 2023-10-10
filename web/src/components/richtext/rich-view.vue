<template>
  <div class="border border-solid border-gray-100 h-full">
    <Editor
      v-model="valueHtml"
      class="overflow-y-hidden mt-0.5"
      :default-config="editorConfig"
      mode="default"
      @onCreated="handleCreated"
      @onChange="change"
    />
  </div>
</template>
<script setup>

import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, ref, shallowRef, watch } from 'vue'
import { Editor } from '@wangeditor/editor-for-vue'

import { useUserStore } from '@/pinia/modules/user'

const userStore = useUserStore()

const emits = defineEmits(['change', 'update:modelValue'])
const editorConfig = ref({
  readOnly: true
})
const change = (editor) => {
  emits('change', editor)
  emits('update:modelValue', valueHtml.value)
}

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const editorRef = shallowRef()
const valueHtml = ref('')

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleCreated = (editor) => {
  editorRef.value = editor
  valueHtml.value = props.modelValue
}

watch(() => props.modelValue, () => {
  valueHtml.value = props.modelValue
})
</script>

<style scoped lang="scss">

</style>
