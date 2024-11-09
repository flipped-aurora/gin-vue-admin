<template>
  <div
    class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 py-2 gap-4 md:gap-2 gva-container2"
  >
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32">
      <gva-chart :type="1" title="访问人数" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32 ">
      <gva-chart :type="2" title="新增客户" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32">
      <gva-chart :type="3" title="解决数量" />
    </gva-card>
    <gva-card
      title="快捷功能"
      show-action
      custom-class="col-start-1 md:col-start-3 lg:col-start-7 row-span-2 h-38"
    >
      <gva-quick-link />
    </gva-card>
    <gva-card
      title="内容数据"
      custom-class="col-span-1 md:col-span-2 md:row-start-2 lg:col-span-6 col-start-1 row-span-2"
    >
      <gva-chart :type="4" />
    </gva-card>
    <gva-card
      title="文档"
      show-action
      custom-class="md:row-start-8 md:col-start-3 lg:row-start-3 lg:col-start-7"
    >
      <gva-wiki />
    </gva-card>

    <gva-card
      title="最新更新"
      custom-class="col-span-1 md:col-span-3 row-span-2"
    >
      <gva-table />
    </gva-card>
    <gva-card
      title="最新插件"
      custom-class="col-span-1 md:col-span-3 row-span-2"
    >
      <gva-plugin-table />
    </gva-card>

    <gva-card title="公告" show-action custom-class="col-span-1 lg:col-start-7">
      <gva-notice />
    </gva-card>

    <gva-card
      without-padding
      custom-class="overflow-hidden lg:h-40 col-span-1 md:col-start-2 md:col-span-1 lg:col-start-7"
    >
      <gva-banner />
    </gva-card>
  </div>
</template>

<script setup>
  import {
    GvaPluginTable,
    GvaTable,
    GvaChart,
    GvaWiki,
    GvaNotice,
    GvaQuickLink,
    GvaCard,
    GvaBanner
  } from './components'
  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>
