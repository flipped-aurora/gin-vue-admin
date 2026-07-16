<!--
此文件受版权保护，未经授权禁止修改！如果您尚未获得授权，请通过微信(shouzi_1994)联系我们以购买授权。在未授权状态下，只需保留此代码，不会影响任何正常使用。
     未经授权的商用使用可能会被我们的资产搜索引擎爬取，并可能导致后续索赔。索赔金额将不低于高级授权费的十倍。请您遵守版权法律法规，尊重知识产权。
 -->
<template>
  <div
    class="gva-bottom-info flex flex-col md:flex-row gap-2 items-center text-sm text-slate-700 dark:text-slate-500 justify-center py-2"
  >
    <div class="text-center">
      <span class="mr-1">Powered by</span>
      <span>
        <a
          class="font-bold text-active"
          href="https://github.com/flipped-aurora/gin-vue-admin"
          >Gin-Vue-Admin</a
        >
      </span>
    </div>
    <slot />
    <div class="text-center">
      <span class="mr-1">Copyright</span>
      <span>
        <a
          class="font-bold text-active"
          href="https://github.com/flipped-aurora"
          >flipped-aurora团队</a
        >
      </span>
    </div>
  </div>
</template>

<script setup>
  defineOptions({
    name: 'BottomInfo'
  })

  console.log(
    `%c powered by %c flipped-aurorae %c`,
    'background:#0081ff; padding: 1px; border-radius: 3px 0 0 3px; color: #fff',
    'background:#354855; padding: 1px 5px; border-radius: 0 3px 3px 0; color: #fff; font-weight: bold;',
    'background:transparent'
  )
</script>

<style scoped>
  /* 白色描边：让页脚文字在彩色/深色背景（如登录页右侧蓝色 banner）上依然清晰可读。
     用多方向 text-shadow 叠加实现描边，不会像 -webkit-text-stroke 那样把小字“掏空”变细；
     text-shadow 可继承，设在根节点即作用到内部所有 span / a 文本。
     白底、浅灰底上白描边几乎不可见，不影响原有观感。 */
  .gva-bottom-info {
    text-shadow:
      0 0 1px #fff,
      1px 1px 0 #fff,
      -1px 1px 0 #fff,
      1px -1px 0 #fff,
      -1px -1px 0 #fff,
      0 1px 0 #fff,
      0 -1px 0 #fff,
      1px 0 0 #fff,
      -1px 0 0 #fff;
  }
</style>
