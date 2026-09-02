<template>
  <div class="h-full gva-container2 overflow-auto bg-slate-50/60 dark:bg-slate-900">
    <div class="space-y-4 p-4 lg:p-6">
      <section
        class="relative overflow-hidden rounded-xl border border-slate-200/80 bg-white px-5 py-6 shadow-sm dark:border-slate-700 dark:bg-slate-800"
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
          <div class="flex flex-wrap items-center gap-2">
            <el-button type="primary" class="!ml-0" @click="goLicensedDemo">
              <el-icon class="mr-1"><Monitor /></el-icon>
              查看授权版演示
            </el-button>
            <el-button class="!ml-0" @click="goLicense">购买商业授权</el-button>
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
            class="relative min-h-[220px] flex-1 overflow-hidden rounded-lg border border-amber-200 bg-white p-5 shadow-sm dark:border-amber-500/40 dark:bg-slate-900"
          >
            <div class="relative flex h-full flex-col">
              <div class="flex items-start justify-between gap-3">
                <div
                  class="flex h-10 w-10 shrink-0 items-center justify-center rounded-lg bg-[var(--el-color-primary-light-9)] text-[var(--el-color-primary)] dark:bg-[var(--el-color-primary-dark-2)] dark:text-white"
                >
                  <el-icon size="20"><Monitor /></el-icon>
                </div>
                <span class="rounded bg-amber-50 px-2 py-1 text-xs font-medium text-amber-700 dark:bg-amber-400/10 dark:text-amber-300">
                  授权版演示
                </span>
              </div>
              <h3 class="mt-4 text-lg font-semibold text-slate-900 dark:text-white">
                先体验完整授权版，再决定是否购买
              </h3>
              <p class="mt-2 text-sm leading-6 text-slate-600 dark:text-slate-300">
                在线查看授权版的完整界面与业务能力，直观了解它能为团队交付带来的提升。
              </p>
              <div class="mt-4 grid grid-cols-1 gap-2 text-xs text-slate-600 sm:grid-cols-3 xl:grid-cols-1 dark:text-slate-300">
                <span class="flex items-center gap-2"><el-icon class="text-emerald-500"><CircleCheckFilled /></el-icon>完整业务体验</span>
                <span class="flex items-center gap-2"><el-icon class="text-emerald-500"><CircleCheckFilled /></el-icon>持续升级权益</span>
                <span class="flex items-center gap-2"><el-icon class="text-emerald-500"><CircleCheckFilled /></el-icon>商业授权支持</span>
              </div>
              <div class="mt-5 flex flex-wrap items-center gap-2">
                <el-button type="primary" class="!ml-0" @click="goLicensedDemo">
                  查看授权版演示
                </el-button>
                <el-button plain class="!ml-0" @click="goLicense">购买商业授权</el-button>
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
  import { COMMERCIAL_LINKS, DEMO_SITES } from '@/config/external-links'
  import { toDoc } from '@/utils/doc'

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

  const goLicensedDemo = () => {
    toDoc(DEMO_SITES.licensed.url)
  }

  const goLicense = () => {
    toDoc(COMMERCIAL_LINKS.purchase)
  }

  defineOptions({
    name: 'Dashboard'
  })
</script>

<style lang="scss" scoped></style>

