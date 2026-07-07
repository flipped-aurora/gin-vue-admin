<script setup>
// Chrome 标签的弧形底片（带圆角的梯形），移植自 Soybean 的 chrome-tab-bg：
// fill=currentColor，由外层 .__bg 的 color 控制填充色，跟随主题 / 暗色。
//
// 高度自适配：底片纵向拉伸铺满整个标签高度，底边始终贴住标签栏底部（Chrome 观感靠“标签底与内容区无缝相接”）。
// symbol 用 preserveAspectRatio="none" 只做纵向缩放；横向仍靠固定 214 宽的 <use> 保持圆角形状不变形（Soybean 的“固定圆角+中段拉伸”技巧）。
// 注意：不要把 <use> 的 height 写成固定像素（如 36），否则标签实际高度一旦不等于该值，底片就会顶对齐、底部空出缝隙、标签“飘”起来贴不到底。
defineOptions({ name: 'PageTabChromeBg' })
</script>

<template>
  <svg class="h-full w-full">
    <defs>
      <symbol
        id="gva-chrome-geometry-left"
        viewBox="0 0 214 36"
        preserveAspectRatio="none"
      >
        <path d="M17 0h197v36H0v-2c4.5 0 9-3.5 9-8V8c0-4.5 3.5-8 8-8z" />
      </symbol>
      <symbol
        id="gva-chrome-geometry-right"
        viewBox="0 0 214 36"
        preserveAspectRatio="none"
      >
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
        height="100%"
        fill="currentColor"
      />
    </svg>
    <g transform="scale(-1, 1)">
      <svg x="-100%" y="0" width="51%" height="100%">
        <use
          xlink:href="#gva-chrome-geometry-right"
          width="214"
          height="100%"
          fill="currentColor"
        />
      </svg>
    </g>
  </svg>
</template>
