<template>
  <div>
    <span class="text-red-600">{{ name }}</span>
    <span class="text-red-600">{{ userInfo }}</span>
    <el-button @click="loginFunc">登录</el-button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useUserStore} from "@/pinia/modules/user.js";
import {login} from '@/api/base.js'
const userStore = useUserStore();
const userInfo = userStore.userInfo
const name = ref('九尾AI')
const loginFunc = () => {
  login()
}
</script>