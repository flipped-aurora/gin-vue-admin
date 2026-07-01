<script setup>
// Chrome 标签的弧形底片（带圆角的梯形），移植自 Soybean 的 chrome-tab-bg：
// fill=currentColor，由外层 .__bg 的 color 控制填充色，跟随主题 / 暗色。
defineOptions({ name: 'PageTabChromeBg' })
</script>

<template>
  <svg class="h-full w-full">
    <defs>
      <symbol id="gva-chrome-geometry-left" viewBox="0 0 214 36">
        <path d="M17 0h197v36H0v-2c4.5 0 9-3.5 9-8V8c0-4.5 3.5-8 8-8z" />
      </symbol>
      <symbol id="gva-chrome-geometry-right" viewBox="0 0 214 36">
        <use xlink:href="#gva-chrome-geometry-left" />
      </symbol>
      <clipPath>
        <rect x="0" width="100%" height="100%" />
      </clipPath>
    </defs>
    <svg width="51%" height="100%">
      <use
        xlink:href="#gva-chrome-geometry-left"
        width="214"
        height="36"
        fill="currentColor"
      />
    </svg>
    <g transform="scale(-1, 1)">
      <svg x="-100%" y="0" width="51%" height="100%">
        <use
          xlink:href="#gva-chrome-geometry-right"
          width="214"
          height="36"
          fill="currentColor"
        />
      </svg>
    </g>
  </svg>
</template>
