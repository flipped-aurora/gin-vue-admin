<template>
  <div class="w-full min-h-screen bg-white flex items-center justify-center">
    <div class="flex flex-col items-center text-center px-4">
      <!-- 指南针图标 -->
      <svg
        class="w-24 h-24 mb-10"
        viewBox="0 0 96 96"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <circle cx="48" cy="48" r="45" fill="#ffffff" stroke="#475569" stroke-width="2" />
        <!-- 四个方向刻度(短、带间距,保持通透) -->
        <path
          d="M48 6V15 M48 81V90 M6 48H15 M81 48H90"
          stroke="#475569"
          stroke-width="2"
          stroke-linecap="round"
        />
        <!-- 指针下半(浅色) -->
        <path
          d="M40.5 48 L55.5 48 L48 79 Z"
          fill="#e5e7eb"
          stroke="#94a3b8"
          stroke-width="1"
          stroke-linejoin="round"
        />
        <!-- 指针上半(蓝色) -->
        <path d="M48 17 L55.5 48 L40.5 48 Z" fill="#2563eb" stroke-linejoin="round" />
        <!-- 中心轴点 -->
        <circle cx="48" cy="48" r="2.4" fill="#ffffff" stroke="#475569" stroke-width="1" />
      </svg>

      <!-- 404 -->
      <p class="text-base font-semibold tracking-[0.35em] text-gray-400 mb-5 pl-[0.35em]">
        404
      </p>

      <!-- 主标题 -->
      <h1 class="text-5xl font-bold text-gray-900 mb-5">页面未找到</h1>

      <!-- 副标题 -->
      <p class="text-base text-gray-500 mb-10">抱歉,你访问的页面不存在或已被移动。</p>

      <!-- 操作按钮 -->
      <div class="flex items-center gap-4">
        <!-- 返回首页 -->
        <button
          class="inline-flex items-center gap-2 px-6 py-3 rounded-lg text-base font-medium text-white bg-blue-600 hover:bg-blue-700 transition-colors"
          @click="toDashboard"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
            <polyline points="9 22 9 12 15 12 15 22" />
          </svg>
          返回首页
        </button>

        <!-- 退出登录 -->
        <button
          class="inline-flex items-center gap-2 px-6 py-3 rounded-lg text-base font-medium text-gray-700 bg-white border border-gray-200 hover:bg-gray-50 transition-colors"
          @click="logout"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
            <polyline points="16 17 21 12 16 7" />
            <line x1="21" y1="12" x2="9" y2="12" />
          </svg>
          退出登录
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { useUserStore } from '@/pinia/modules/user'
  import { useRouter } from 'vue-router'
  import { emitter } from '@/utils/bus'

  defineOptions({
    name: 'Error'
  })

  const userStore = useUserStore()
  const router = useRouter()

  // 返回首页:跳转到当前用户的默认路由
  const toDashboard = () => {
    try {
      router.push({ name: userStore.userInfo.authority.defaultRouter })
    } catch (error) {
      emitter.emit('show-error', {
        code: '401',
        message: '检测到其他用户修改了路由权限,请重新登录',
        fn: () => {
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        }
      })
    }
  }

  // 退出登录:复用导航栏的登出逻辑(拉黑 token、清理缓存、跳转登录页并刷新)
  const logout = () => {
    userStore.LoginOut()
  }
</script>