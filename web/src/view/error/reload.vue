<template>
  <div />
</template>
<script>
import { useRouter } from 'vue-router'
export default {
  name: 'Reload'
}
</script>

<script setup>
const router = useRouter()
router.go(-1)
</script>
