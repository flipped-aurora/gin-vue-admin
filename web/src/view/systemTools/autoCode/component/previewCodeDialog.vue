<template>
  <el-tabs
    v-model="activeName"
    tab-position="left"
    class="h-[calc(100vh-110px)]"
  >
    <el-tab-pane
      v-for="(item, key) in useCode"
      :key="key"
      :label="key"
      :name="key"
    >
      <div :id="key" class="h-[calc(100vh-110px)] px-5 overflow-y-scroll"></div>
    </el-tab-pane>
  </el-tabs>
</template>

<script setup>
  import { Marked } from 'marked'
  import { markedHighlight } from 'marked-highlight'
  import hljs from 'highlight.js'
  import { ElMessage } from 'element-plus'
  import { onMounted, ref, watchEffect } from 'vue'
  import { useAppStore } from '@/pinia'

  const appStore = useAppStore()

  const useCode = ref({})

  const createKey = [
    'enter.go',
    'gorm_biz.go',
    'router_biz.go',
    'api',
    'router',
    'initialize',
    'gen.go'
  ]

  onMounted(() => {
    const isDarkMode = appStore.config.darkMode === 'dark'
    if (isDarkMode) {
      import('highlight.js/styles/atom-one-dark.css')
    } else {
      import('highlight.js/styles/atom-one-light.css')
    }
  })

  const props = defineProps({
    previewCode: {
      type: Object,
      default() {
        return {}
      }
    },
    isAdd: {
      type: Boolean,
      default: false
    }
  })

  watchEffect(() => {
    for (const key in props.previewCode) {
      if (
        props.isAdd &&
        createKey.some((createKeyItem) => key.includes(createKeyItem))
      ) {
        continue
      }
      useCode.value[key] = props.previewCode[key]
    }
  })

  const activeName = ref('')
  onMounted(() => {
    const marked = new Marked(
      markedHighlight({
        langPrefix: 'hljs language-',
        highlight(code, lang) {
          const language = hljs.getLanguage(lang) ? lang : 'plaintext'
          if (lang === 'vue') {
            return hljs.highlight(code, { language: 'html' }).value
          }
          return hljs.highlight(code, { language }).value
        }
      })
    )
    for (const key in useCode.value) {
      if (activeName.value === '') {
        activeName.value = key
      }
      document.getElementById(key).innerHTML = marked.parse(useCode.value[key])
    }
  })

  const selectText = () => {
    const element = document.getElementById(activeName.value)
    if (document.body.createTextRange) {
      const range = document.body.createTextRange()
      range.moveToElementText(element)
      range.select()
    } else if (window.getSelection) {
      const selection = window.getSelection()
      const range = document.createRange()
      range.selectNodeContents(element)
      selection.removeAllRanges()
      selection.addRange(range)
    } else {
      alert('none')
    }
  }
  const copy = () => {
    selectText()
    document.execCommand('copy')
    ElMessage.success('复制成功')
  }

  defineExpose({ copy, selectText })
</script>
