<template>
  <div class="chat-table">
    <el-button @click="test">
      test
    </el-button>
  </div>
</template>

<script setup>
import { getTable } from '@/api/chatgpt'
const test = () => {
  getTable({
    dbname:"gva",
    chat: '帮我找到角色id为888的用户',
  })
}
</script>

<style scoped>

</style>
