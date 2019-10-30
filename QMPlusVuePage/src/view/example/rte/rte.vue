<template>
  <div class="edit_container">
    <quill-editor
      :options="editorOption"
      @blur="onEditorBlur($event)"
      @change="onEditorChange($event)"
      @focus="onEditorFocus($event)"
      ref="myQuillEditor"
      v-model="content"
    ></quill-editor>
    <button v-on:click="saveHtml">保存</button>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      content: `<p>hello world</p>`,
      editorOption: {}
    }
  },
  computed: {
    editor() {
      return this.$refs.myQuillEditor.quill
    }
  },
  methods: {
    // onEditorReady(editor) { // 准备编辑器

    // },
    onEditorBlur() {}, // 失去焦点事件
    onEditorFocus() {}, // 获得焦点事件
    onEditorChange() {}, // 内容改变事件
    saveHtml() {} // 保存方法
  }
}
</script>

<style>

</style>