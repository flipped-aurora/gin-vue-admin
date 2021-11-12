<template>
  <div class="previewCode">
    <el-tabs v-model="activeName">
      <el-tab-pane v-for="(item, key) in previewCode" :key="key" :label="key" :name="key">
        <div :id="key" class="tab-info" />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import marked from 'marked'
import hljs from 'highlight.js'
// import 'highlight.js/styles/atelier-cave-light.css';
import 'highlight.js/styles/atelier-plateau-light.css'
export default {
  props: {
    previewCode: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  data() {
    return {
      activeName: ''
    }
  },
  mounted() {
    marked.setOptions({
      renderer: new marked.Renderer(),
      highlight: function(code) {
        return hljs.highlightAuto(code).value
      },
      pedantic: false,
      gfm: true,
      tables: true,
      breaks: false,
      sanitize: false,
      smartLists: true,
      smartypants: false,
      xhtml: false
    })
    for (const key in this.previewCode) {
      if (this.activeName === '') {
        this.activeName = key
      }
      document.getElementById(key).innerHTML = marked(this.previewCode[key])
    }
  },
  methods: {
    selectText() {
      const element = document.getElementById(this.activeName)
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
    },
    copy() {
      this.selectText()
      document.execCommand('copy')
      this.$message.success('复制成功')
    }
  }
}
</script>

<style lang="scss">
.previewCode {
  .tab-info {
    height: 50vh;
    background: #fff;
    padding: 0 20px;
    overflow-y: scroll;
  }
}
</style>
