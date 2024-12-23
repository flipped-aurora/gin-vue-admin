<template>
  <div />
</template>

<script setup>
  import { useRouter } from 'vue-router'

  defineOptions({
    name: 'Reload'
  })

  const router = useRouter()
  router.go(-1)
</script>
