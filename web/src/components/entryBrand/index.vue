<!--
  登录 / 初始化等入口页共用的品牌头部：logo + 大写字标。
  统一维护尺寸与样式，保证各入口页一致。
-->
<template>
  <div class="flex flex-col items-center">
    <Logo :size="logoSize" />
    <p
      class="mt-4 text-2xl font-bold uppercase tracking-[0.2em] text-gray-800 dark:text-white"
    >
      {{ $GIN_VUE_ADMIN.appName }}
    </p>
  </div>
</template>

<script setup>
  import Logo from '@/components/logo/index.vue'

  defineOptions({
    name: 'EntryBrand'
  })

  defineProps({
    // logo 尺寸（单位 rem，见 logo 组件）
    logoSize: {
      type: Number,
      default: 6
    }
  })
</script>
