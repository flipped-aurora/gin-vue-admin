<template>
  <div
    class="rounded-xl border border-black/10 bg-white text-black/80 dark:text-slate-400 dark:bg-slate-900 dark:text-white/80"
    :class="[customClass || '', withoutPadding ? 'p-0' : 'p-4']"
  >
    <div v-if="title" class="flex justify-between items-center">
      <div class="text-sm font-semibold tracking-tight text-black dark:text-white">
        {{ title }}
      </div>
      <div
        v-if="showAction"
        class="text-xs text-black/60 dark:text-white/60 hover:text-active cursor-pointer"
      >
        查看更多
      </div>
    </div>
    <div :class="title ? 'mt-3' : ''">
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
