<template>
  <div class="h-full gva-container2 overflow-auto bg-slate-50/60 dark:bg-slate-900">
    <div class="space-y-4 p-4 lg:p-6">
      <section
        class="relative overflow-hidden rounded-xl border border-slate-200/80 bg-white px-5 py-6 shadow-sm dark:border-slate-700 dark:from-slate-900 dark:via-slate-800 dark:to-slate-900"
      >
        
        <div class="relative flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
          <div>
            <p class="text-xs tracking-[0.2em] text-slate-500 dark:text-slate-400">DASHBOARD</p>
            <h1 class="mt-2 text-xl font-semibold text-slate-900 dark:text-slate-100 lg:text-2xl">
              欢迎回来，开始今天的Coding节奏
            </h1>
            <p class="mt-2 text-sm text-slate-600 dark:text-slate-300">
              {{ today }} · 已为你聚合核心业务数据、插件动态和系统公告
            </p>
          </div>
          <div class="flex items-center gap-2">
            <el-button type="primary" @click="goLicense">购买商业授权</el-button>
            <el-button @click="goPluginMarket">插件市场</el-button>
          </div>
        </div>
      </section>

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

      <div class="grid grid-cols-1 items-stretch gap-4 xl:grid-cols-12">
        <div class="grid grid-cols-1 gap-4 content-start xl:col-span-8 xl:h-full">
          <gva-card title="内容数据">
            <gva-chart :type="4" />
          </gva-card>

          <gva-card title="最新插件">
            <gva-plugin-table />
          </gva-card>

          <gva-card title="最新更新">
            <gva-table />
          </gva-card>
        </div>

        <div class="flex flex-col gap-4 xl:col-span-4 xl:h-full">
          <gva-card title="快捷功能" show-action custom-class="min-h-[300px]">
            <gva-quick-link />
          </gva-card>
          <gva-card title="公告" show-action custom-class="min-h-[300px]">
            <gva-notice />
          </gva-card>
          <gva-card title="文档" show-action custom-class="min-h-[120px]">
            <gva-wiki />
          </gva-card>
          <div
            class="relative min-h-[200px] flex-1 overflow-hidden rounded-lg border border-slate-200 bg-slate-900 p-5 text-white shadow-sm dark:border-slate-700"
          >
            
            <div class="relative">
              <div class="inline-flex rounded-full bg-white/10 px-3 py-1 text-xs">商业授权</div>
              <h3 class="mt-3 text-lg font-semibold">解锁完整商用支持与专属服务</h3>
              <p class="mt-2 text-sm text-slate-200/90">
                购买授权后可获得专属支持通道、插件优惠与商用合规保障，帮助团队更稳定地推进项目交付。
              </p>
              <div class="mt-4 flex flex-wrap gap-2 text-xs">
                <span class="rounded-full bg-white/10 px-2.5 py-1">专属技术支持</span>
                <span class="rounded-full bg-white/10 px-2.5 py-1">插件优惠权益</span>
                <span class="rounded-full bg-white/10 px-2.5 py-1">商用授权凭证</span>
              </div>
              <div class="mt-5 flex items-center gap-3">
                <el-button type="primary" @click="goLicense">立即购买</el-button>
                <el-button link class="!text-cyan-300" @click="goPluginMarket">查看插件市场</el-button>
              </div>
            </div>
          </div>
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
    GvaCard
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

  const goLicense = () => {
    window.open('https://plugin.gin-vue-admin.com/license', '_blank', 'noopener,noreferrer')
  }

  const goPluginMarket = () => {
    window.open('https://plugin.gin-vue-admin.com', '_blank', 'noopener,noreferrer')
  }

  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>

