<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div
    class="bg-white dark:bg-slate-900 text-gray-800 dark:text-gray-400 rounded shadow"
    :class="[customClass || '', withoutPadding ? 'p-0' : 'p-4']"
  >
    <div v-if="title" class="flex justify-between items-center">
      <div class="text-base font-bold">
        {{ title }}
      </div>
      <div v-if="showAction" class="text-sm text-active cursor-pointer">
        查看更多
      </div>
    </div>
    <div class="mt-2">
      <slot />
    </div>
  </div>
</template>

<script setup>
  defineProps({
    title: {
      type: String,
      default: ''
    },
    showAction: {
      type: Boolean,
      default: false
    },
    customClass: {
      type: String,
      default: ''
    },
    withoutPadding: {
      type: Boolean,
      default: false
    }
  })
</script>

<style scoped lang="scss"></style>
