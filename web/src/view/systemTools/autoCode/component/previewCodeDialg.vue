<template>
  <div class="previewCode">
    <div class="previewCodeTool">
      <p>操作栏：</p>
      <el-button @click="selectText">全选</el-button>
      <el-button @click="copy">复制</el-button>
    </div>
    <el-tabs v-model="activeName">
      <el-tab-pane :label="key" :name="key" v-for="(item, key) in previewCode" :key="key">
        <div class="tab-info" :id="key"></div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import marked from "marked";
import hljs from "highlight.js";
// import 'highlight.js/styles/atelier-cave-light.css';
import "highlight.js/styles/atelier-plateau-light.css";
export default {
  props: {
    previewCode: {
      type: Object,
      default() {
        return {};
      },
    },
  },
  data() {
    return {
      activeName: "",
    };
  },
  methods: {
    selectText() {
      const element = document.getElementById(this.activeName);
      if (document.body.createTextRange) {
        let range = document.body.createTextRange();
        range.moveToElementText(element);
        range.select();
      } else if (window.getSelection) {
        let selection = window.getSelection();
        let range = document.createRange();
        range.selectNodeContents(element);
        selection.removeAllRanges();
        selection.addRange(range);
      } else {
        alert("none");
      }
    },
    copy() {
      this.selectText();
      document.execCommand("copy");
      this.$message.success("复制成功");
    },
  },
  mounted() {
    marked.setOptions({
      renderer: new marked.Renderer(),
      highlight: function (code) {
        return hljs.highlightAuto(code).value;
      },
      pedantic: false,
      gfm: true,
      tables: true,
      breaks: false,
      sanitize: false,
      smartLists: true,
      smartypants: false,
      xhtml: false,
    });
    for (const key in this.previewCode) {
      if (this.activeName == "") {
        this.activeName = key;
      }
      document.getElementById(key).innerHTML = marked(this.previewCode[key]);
    }
  },
};
</script>

<style lang="scss">
.previewCode {
  .previewCodeTool {
    display: flex;
    align-items: center;
    padding: 5px 0;
    margin-top: -40px;
  }
  .tab-info {
    height: 50vh;
    background: #fff;
    padding: 0 20px;
    overflow-y: scroll;
  }
}
</style>
