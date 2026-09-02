<template>
  <div class="">
    <div class="flex items-center justify-between mb-2">
      <div v-if="title" class="text-sm font-semibold tracking-tight text-black dark:text-white">
        {{ title }}
      </div>
      <slot v-else name="title" />
    </div>
    <div class="w-full relative">
      <div v-if="type !== 4">
        <div class="mt-4 text-3xl font-mono text-black dark:text-white">
          <el-statistic :value="268500" />
        </div>
        <div class="mt-2 text-xs font-mono text-black/60 dark:text-white/60">
          +80% <el-icon class="align-middle"><TopRight /></el-icon>
        </div>
      </div>
      <div class="absolute top-0 right-2 w-[50%] h-20">
        <charts-people-number v-if="type === 1" :data="data[0]" height="100%" />
        <charts-people-number v-if="type === 2" :data="data[1]" height="100%" />
        <charts-people-number v-if="type === 3" :data="data[2]" height="100%" />
      </div>
      <charts-content-number v-if="type === 4" height="14rem" />
    </div>
  </div>
</template>

<script setup>
  import chartsPeopleNumber from './charts-people-numbers.vue'
  import chartsContentNumber from './charts-content-numbers.vue'
  defineProps({
    type: {
      type: Number,
      default: 1
    },
    title: {
      type: String,
      default: ''
    }
  })

  const data = [
    [12, 22, 32, 45, 32, 78, 89, 92],
    [1, 2, 43, 5, 67, 78, 89, 12],
    [12, 22, 32, 45, 32, 78, 89, 92]
  ]
</script>

<style scoped lang="scss"></style>
