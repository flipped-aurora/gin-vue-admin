<template>
  <div
    class="h-full gva-container2 overflow-auto bg-white text-black dark:bg-slate-800 dark:text-white"
  >
    <div class="p-4 lg:p-6">
      <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 xl:grid-cols-3">
        <gva-card>
          <gva-chart :type="1" title="访问人数" />
        </gva-card>
        <gva-card>
          <gva-chart :type="2" title="新增客户" />
        </gva-card>
        <gva-card>
          <gva-chart :type="3" title="解决数量" />
        </gva-card>
      </div>

      <div class="mt-4 grid grid-cols-1 gap-4 xl:grid-cols-12 items-start">
        <div class="grid grid-cols-1 gap-4 xl:col-span-8 self-start content-start">
          <gva-card title="内容数据">
            <gva-chart :type="4" />
          </gva-card>

          <div class="grid grid-cols-1 gap-4">
            <gva-card title="最新插件">
              <gva-plugin-table />
            </gva-card>
          </div>

          <div class="grid grid-cols-1 gap-4">
            <gva-card title="最新更新">
              <gva-table />
            </gva-card>
          </div>
        </div>

        <div class="grid grid-cols-1 gap-4 xl:col-span-4 self-start content-start">
          <gva-card title="快捷功能" show-action>
            <gva-quick-link />
          </gva-card>
          <gva-card title="公告" show-action>
            <gva-notice />
          </gva-card>
          <gva-card title="文档" show-action>
            <gva-wiki />
          </gva-card>
          <gva-card
            without-padding
            custom-class="overflow-hidden"
          >
            <gva-banner />
          </gva-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { computed } from 'vue'
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

  const today = computed(() => {
    try {
      const d = new Date()
      return d.toLocaleDateString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit'
      })
    } catch (e) {
      return new Date().toISOString().slice(0, 10)
    }
  })
  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>
