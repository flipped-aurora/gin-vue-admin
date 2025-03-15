<template>
  <div class="title relative my-2">
    <div class="flex-shrink-0 text-center text-xl text-gray-600">
      {{ title }}
    </div>
  </div>
</template>

<script setup>
  defineOptions({
    name: 'layoutSettingTitle'
  })

  defineProps({
    title: String
  })
</script>

<style scoped>
  .title {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 4rem;
  }

  .title::before,
  .title::after {
    content: '';
    height: 1px;
    width: 100%;
    background-color: #e3e3e3;
  }
</style>
