<template>
  <div class="richtext-wrapper border border-solid border-gray-100 h-full">
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
