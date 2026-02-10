<template>
  <div class="richtext-wrapper border border-solid border-gray-100 h-full z-10">
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
  import { useUserStore } from '@/pinia/modules/user'

  const emits = defineEmits(['change', 'update:modelValue'])

  const change = (editor) => {
    emits('change', editor)
    emits('update:modelValue', valueHtml.value)
  }

  const userStore = useUserStore()
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
    headers: {
      'x-token': userStore.token,
    },
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

<style scoped lang="scss">
  .richtext-wrapper {
    :deep(.w-e-text-container [data-slate-editor]) {
      h1 {
        font-size: 2em;
        font-weight: 700;
      }

      h2 {
        font-size: 1.5em;
        font-weight: 700;
      }

      h3 {
        font-size: 1.17em;
        font-weight: 700;
      }

      h4 {
        font-size: 1em;
        font-weight: 700;
      }

      h5 {
        font-size: 0.83em;
        font-weight: 700;
      }

      h6 {
        font-size: 0.67em;
        font-weight: 700;
      }

      ul,
      ol {
        margin: 1em 0;
        padding-left: 2em;
      }

      ul {
        list-style-type: disc;
      }

      ol {
        list-style-type: decimal;
      }

      li {
        margin: 0.25em 0;
      }

      a {
        color: var(--el-color-primary, #409eff);
        text-decoration: underline;
      }
    }

    :deep(.w-e-text-container [data-slate-editor] ul ul) {
      list-style-type: circle;
    }

    :deep(.w-e-text-container [data-slate-editor] ul ul ul) {
      list-style-type: square;
    }

    :deep(.w-e-text-container [data-slate-editor] ol ol) {
      list-style-type: lower-alpha;
    }

    :deep(.w-e-text-container [data-slate-editor] ol ol ol) {
      list-style-type: lower-roman;
    }
  }
</style>
