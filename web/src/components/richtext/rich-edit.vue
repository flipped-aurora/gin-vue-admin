<template>
  <div class="border border-solid border-gray-100 h-full">
    <Toolbar
      :editor="editorRef"
      :default-config="toolbarConfig"
      mode="default"
    />
    <Editor
      v-model="valueHtml"
      class="overflow-y-hidden mt-0.5"
      style="height: 18rem"
      :default-config="editorConfig"
      mode="default"
      @onCreated="handleCreated"
      @onChange="change"
    />
  </div>
</template>

<script setup>
  import '@wangeditor/editor/dist/css/style.css' // 引入 css

  const basePath = import.meta.env.VITE_BASE_API

  import { onBeforeUnmount, ref, shallowRef, watch } from 'vue'
  import { Editor, Toolbar } from '@wangeditor/editor-for-vue'

  import { ElMessage } from 'element-plus'
  import { getUrl } from '@/utils/image'

  const emits = defineEmits(['change', 'update:modelValue'])

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

  const toolbarConfig = {}
  const editorConfig = {
    placeholder: '请输入内容...',
    MENU_CONF: {}
  }
  editorConfig.MENU_CONF['uploadImage'] = {
    fieldName: 'file',
    server: basePath + '/fileUploadAndDownload/upload?noSave=1',
    customInsert(res, insertFn) {
      if (res.code === 0) {
        const urlPath = getUrl(res.data.file.url)
        insertFn(urlPath, res.data.file.name)
        return
      }
      ElMessage.error(res.msg)
    }
  }

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

  watch(
    () => props.modelValue,
    () => {
      valueHtml.value = props.modelValue
    }
  )
</script>

<style scoped lang="scss"></style>
