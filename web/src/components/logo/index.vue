<script setup>
import { ref, watchEffect } from 'vue';
import { useAppStore } from '@/pinia/modules/app.js';
import { storeToRefs } from 'pinia';

const props = defineProps({
  // logo 尺寸，单位为:rem
  size: {
    type: Number,
    default: 2
  }
})

import darkLogoPath from "/public/logo.png";  // 系统没有暗黑模式logo，如果需要暗黑模式logo请自行修改文件路径。
import lightLogoPath from "/public/logo.png";

const appStore = useAppStore();
const { isDark } = storeToRefs(appStore);

const logoSrc = ref('');
const showTextPlaceholder = ref(false);

// 检查图片是否存在
function checkImageExists(url) {
  return new Promise((resolve) => {
    const tryToLoad = new Image();
    tryToLoad.onload = () => resolve(true);
    tryToLoad.onerror = () => resolve(false);
    tryToLoad.src = url;
  });
}

watchEffect(async () => {
  showTextPlaceholder.value = false; // 重置占位符状态

  // 暗色模式直接 load，可以省一次亮色的 load
  if (isDark.value && await checkImageExists(darkLogoPath)) {
    logoSrc.value = darkLogoPath;
    return;
  }

  if (await checkImageExists(lightLogoPath)) {
    logoSrc.value = lightLogoPath;
    return
  }

  // 到这里就包没有提供两种 logo 了
  showTextPlaceholder.value = true;
  console.error(
    '错误: 在公共目录中找不到logo.png（或logo-dark.png）。'
  );
  console.warn(
    '解决方案: 请在您的公共目录(/public)中放置logo.png和/或logo-dark.png文件，或确保路径正确。'
  );
});

// 直接用 16px 作为默认的基准大小
const SPACING = 16
function getSize() {
  return {
    width: `${props.size * SPACING}px`,
    height: `${props.size * SPACING}px`,
  }
}
</script>

<template>
  <img v-if="!showTextPlaceholder && logoSrc" :src="logoSrc" :alt="$GIN_VUE_ADMIN.appName" class="object-contain"
    :style="{
      ...getSize()
    }" :class="{
      'filter invert-[90%] hue-rotate-180 brightness-110':
        isDark && logoSrc === '/logo.png',
    }" />
  <div v-else-if="showTextPlaceholder"
    class="rounded-full bg-gray-300 dark:bg-gray-600 flex items-center justify-center text-gray-700 dark:text-gray-200 font-bold text-xs"
    :style="{
      ...getSize()
    }">
    GVA
  </div>
</template>
