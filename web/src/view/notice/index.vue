<template>
  <div class="notice-center">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>通知中心</span>
        </div>
      </template>
      <div class="notice-content">
        <p>欢迎使用通知中心！</p>
        <p>这里是通知中心的主页面，您可以通过左侧菜单访问具体功能：</p>
        <ul>
          <li>通知管理：创建、编辑、发送通知</li>
          <li>在线用户：查看在线用户，管理用户连接</li>
        </ul>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  console.log('Notice Center mounted')
})
</script>

<style scoped>
.notice-center {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.notice-content {
  padding: 20px 0;
}

.notice-content ul {
  margin-top: 15px;
  padding-left: 20px;
}

.notice-content li {
  margin-bottom: 8px;
  line-height: 1.6;
}
</style>