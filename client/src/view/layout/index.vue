<template>
  <div>
    <h1 class="text-3xl font-bold underline flex gap-4 cursor-pointer">
      <div @click="toPath('Home')">Home</div>
      <div @click="toPath('Chat')">Chat</div>
    </h1>
    <router-view></router-view>
  </div>
</template>
<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()
const toPath = (name) => {
  router.push({ name })
}

</script>