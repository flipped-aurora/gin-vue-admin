<template>
  <div
    class="flex justify-between fixed top-0 right-0 z-20 h-16 text-slate-700 dark:text-slate-300 items-center px-2"
    :style="{
      left: isVertical ? headerSideWidth + 'px' : '0px',
      background: 'var(--gva-header-bg)',
      boxShadow: 'var(--gva-header-shadow)'
    }"
  >
    <div class="flex items-center cursor-pointer flex-1">
      <!-- 移动端(<640)：展开抽屉菜单的入口，替代 Logo -->
      <button
        v-if="isMobile"
        type="button"
        aria-label="打开菜单"
        :class="cn(
          'flex h-9 w-9 cursor-pointer appearance-none items-center justify-center rounded-md bg-transparent text-base-text transition-colors hover:bg-muted',
          FOCUS_RING
        )"
        @click="appStore.toggleMobileMenu(true)"
      >
        <svg-icon icon="lucide:menu" class="h-5 w-5" />
      </button>
      <div
        v-else-if="!isVertical"
        class="flex min-w-48 cursor-pointer items-center justify-center"
        @click="router.push({ path: '/' })"
      >
        <Logo />
        <div
          class="inline-flex font-bold text-2xl ml-2"
          :class="
            (effectiveMode === 'head' || effectiveMode === 'combination') &&
            'min-w-fit'
          "
        >
          {{ $GIN_VUE_ADMIN.appName }}
        </div>
      </div>

      <el-breadcrumb
        v-show="!isMobile"
        v-if="settings.header.breadcrumb.visible && effectiveMode !== 'head' && effectiveMode !== 'combination'"
        class="ml-4"
      >
        <el-breadcrumb-item
          v-for="item in matched.slice(1, matched.length)"
          :key="item.path"
        >
          <span class="inline-flex items-center gap-1 font-bold">
            <el-icon v-if="settings.header.breadcrumb.showIcon && item.meta.icon">
              <component :is="item.meta.icon" />
            </el-icon>
            {{ fmtTitle(item.meta.title, route) }}
          </span>
        </el-breadcrumb-item>
      </el-breadcrumb>
      <gva-aside
        v-if="effectiveMode === 'head' && !isMobile"
        class="flex-1"
      />
      <gva-aside
        v-if="effectiveMode === 'combination' && !isMobile"
        mode="head"
        class="flex-1"
      />
    </div>

    <div class="ml-2 flex items-center gap-3">
      <tools />
      <!-- 0.5px 极细分隔线，区分工具区与用户区 -->
      <div
        class="h-6 w-px bg-gray-200 dark:bg-slate-600"
        role="separator"
      />
      <!-- hover 触发：鼠标移上用户胶囊即展开操作菜单，移出（含指针穿越到菜单）后自动收起 -->
      <HoverCardRoot v-model:open="hoverOpen" :open-delay="100" :close-delay="120">
        <HoverCardTrigger as-child>
          <!-- 用户信息聚合为胶囊组件：圆角 + 内边距，hover 时出现浅色反馈 -->
          <div
            class="flex items-center gap-2 px-2 py-1 rounded-lg outline-none cursor-pointer transition-colors hover:bg-muted"
          >
            <CustomPic />
            <span
              v-show="!isMobile"
              class="text-base font-medium max-w-[8rem] truncate text-black dark:text-gray-100"
            >{{ userStore.userInfo.nickName }}</span>
            <el-icon class="text-slate-400 dark:text-slate-500">
              <arrow-down />
            </el-icon>
          </div>
        </HoverCardTrigger>
        <HoverCardPortal>
          <HoverCardContent
            :side-offset="8"
            side="bottom"
            align="end"
            class="z-[3000] min-w-56 rounded-[10px] border border-border bg-container p-1.5 shadow-sider data-[state=open]:animate-[fade-in_0.1s_ease-out] data-[state=closed]:animate-[fade-out_0.1s_ease-in]"
          >
            <!-- 用户身份信息（label，非交互） -->
            <div class="gva-menu-label">
              <div class="min-w-0">
                <div class="truncate text-[15px] font-semibold text-black dark:text-gray-100">
                  {{ userStore.userInfo.nickName }}
                </div>
                <div class="truncate text-[13px] text-slate-500 dark:text-slate-400">
                  当前角色：{{ userStore.userInfo.authority.authorityName }}
                </div>
              </div>
            </div>

            <!-- 分隔线：身份信息 / 角色切换 -->
            <div class="gva-menu-sep" />

            <!-- 角色切换组 -->
            <template v-if="otherAuthorities.length">
              <div
                v-for="item in otherAuthorities"
                :key="item.authorityId"
                class="gva-menu-item"
                @click="handleSelect(() => changeUserAuth(item.authorityId))"
              >
                <svg-icon icon="lucide:repeat" class="h-4 w-4 shrink-0 text-slate-500" />
                <span>切换为：{{ item.authorityName }}</span>
              </div>
              <!-- 分隔线：角色切换 / 账户操作 -->
              <div class="gva-menu-sep" />
            </template>

            <!-- 账户操作组 -->
            <div class="gva-menu-item" @click="handleSelect(toPerson)">
              <svg-icon icon="lucide:user" class="h-4 w-4 shrink-0 text-slate-500" />
              <span>个人信息</span>
            </div>

            <!-- 分隔线：个人信息 / 登出 -->
            <div class="gva-menu-sep" />

            <!-- 退出登录（红色语义色） -->
            <div
              class="gva-menu-item gva-menu-item-danger"
              @click="handleSelect(userStore.LoginOut)"
            >
              <svg-icon icon="lucide:log-out" class="h-4 w-4 shrink-0" />
              <span class="whitespace-nowrap">登&nbsp;出</span>
            </div>
          </HoverCardContent>
        </HoverCardPortal>
      </HoverCardRoot>
    </div>
  </div>
</template>

<script setup>
  import tools from './tools.vue'
  import CustomPic from '@/components/customPic/index.vue'
  import { useUserStore } from '@/pinia/modules/user'
  import { useRoute, useRouter } from 'vue-router'
  import { useAppStore, useThemeStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import { computed, ref } from 'vue'
  import { setUserAuthority } from '@/api/user'
  import { fmtTitle } from '@/utils/fmtRouterTitle'
  import gvaAside from '@/view/layout/aside/index.vue'
  import Logo from '@/components/logo/index.vue'
  import { useLayoutMode } from '@/hooks/useLayoutMode'
  import { useSideWidth } from '@/hooks/useSideWidth'
  import { cn, FOCUS_RING } from '@/core/componentLibrary/utils'
  import {
    HoverCardRoot,
    HoverCardTrigger,
    HoverCardPortal,
    HoverCardContent
  } from 'reka-ui'

  const userStore = useUserStore()
  const router = useRouter()
  const route = useRoute()
  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { device } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)
  const isMobile = computed(() => {
    return device.value === 'mobile'
  })
  const { effectiveMode } = useLayoutMode()
  // 通栏侧边布局：header 让出左侧侧栏宽度，并隐藏自身 Logo
  const isVertical = computed(
    () => effectiveMode.value === 'vertical' && !isMobile.value
  )
  const { sideWidth: headerSideWidth } = useSideWidth()
  // 除当前角色外的可切换角色
  const otherAuthorities = computed(() => {
    const list = userStore.userInfo.authorities || []
    return list.filter(
      (i) => i.authorityId !== userStore.userInfo.authorityId
    )
  })
  const toPerson = () => {
    router.push({ name: 'person' })
  }
  const matched = computed(() => route.meta.matched)

  // hover 菜单受控：点击任意菜单项后立即关闭
  const hoverOpen = ref(false)
  const handleSelect = (fn) => {
    hoverOpen.value = false
    fn && fn()
  }

  const changeUserAuth = async (id) => {
    const res = await setUserAuthority({
      authorityId: id
    })
    if (res.code === 0) {
      window.sessionStorage.setItem('needCloseAll', 'true')
      window.sessionStorage.setItem('needToHome', 'true')
      window.location.reload()
    }
  }
</script>
