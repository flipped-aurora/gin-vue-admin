<template>
  <div
    class="gva-card bg-container text-base-text"
    :class="[customClass || '', withoutPadding ? 'p-0' : 'p-4']"
  >
    <div v-if="title" class="flex justify-between items-center">
      <div class="text-sm font-semibold tracking-tight text-base-text">
        {{ title }}
      </div>
      <div
        v-if="showAction"
        class="text-xs text-muted-foreground hover:text-active cursor-pointer"
      >
        更多
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

