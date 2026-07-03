<script setup>
import { computed, ref, watchEffect } from 'vue';
import { useThemeStore } from '@/pinia';
import { storeToRefs } from 'pinia';

const props = defineProps({
  // logo 尺寸，单位为:rem
  size: {
    type: Number,
    default: 2
  },
  // 侧栏深色时强制走暗色 logo：浅色主题 + 深色侧边栏场景下全局 isDark 仍为 false，
  // 通栏侧边 / 移动抽屉把此值传进来，让 logo 连同侧栏一起变深。
  dark: {
    type: Boolean,
    default: false
  }
})

const darkLogoPath = "/logo.png";  // 系统没有暗黑模式logo，如果需要暗黑模式logo请自行修改文件路径。
const lightLogoPath = "/logo.png";

const themeStore = useThemeStore();
const { isDark } = storeToRefs(themeStore);

// 有效暗色：全局暗色 或 侧栏深色（props.dark）任一成立即走暗色 logo
const effectiveDark = computed(() => props.dark || isDark.value);

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

watchEffect(async (onCleanup) => {
  // effectiveDark 快速翻转时，用 stale 标记丢弃过期的异步结果，避免旧结果覆盖新值
  let stale = false
  onCleanup(() => (stale = true))
  showTextPlaceholder.value = false; // 重置占位符状态

  // 暗色模式直接 load，可以省一次亮色的 load
  if (effectiveDark.value && await checkImageExists(darkLogoPath)) {
    if (!stale) logoSrc.value = darkLogoPath;
    return;
  }

  if (await checkImageExists(lightLogoPath)) {
    if (!stale) logoSrc.value = lightLogoPath;
    return
  }

  if (stale) return
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
        effectiveDark && logoSrc === '/logo.png',
    }" />
  <div v-else-if="showTextPlaceholder"
    class="rounded-full bg-gray-300 dark:bg-gray-600 flex items-center justify-center text-gray-700 dark:text-gray-200 font-bold text-xs"
    :style="{
      ...getSize()
    }">
    GVA
  </div>
</template>
