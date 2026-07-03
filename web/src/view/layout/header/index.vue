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

    <div class="ml-2 flex items-center">
      <tools />
      <el-dropdown>
        <div class="flex justify-center items-center h-full w-full">
          <span
            class="cursor-pointer flex justify-center items-center text-black dark:text-gray-100"
          >
            <CustomPic />
            <span v-show="!isMobile" class="w-16">{{
              userStore.userInfo.nickName
            }}</span>
            <el-icon>
              <arrow-down />
            </el-icon>
          </span>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <span class="font-bold">
                当前角色：{{ userStore.userInfo.authority.authorityName }}
              </span>
            </el-dropdown-item>
            <template v-if="userStore.userInfo.authorities">
              <el-dropdown-item
                v-for="item in userStore.userInfo.authorities.filter(
                  (i) => i.authorityId !== userStore.userInfo.authorityId
                )"
                :key="item.authorityId"
                @click="changeUserAuth(item.authorityId)"
              >
                <span> 切换为：{{ item.authorityName }} </span>
              </el-dropdown-item>
            </template>
            <el-dropdown-item icon="avatar" @click="toPerson">
              个人信息
            </el-dropdown-item>
            <el-dropdown-item icon="reading-lamp" @click="userStore.LoginOut">
              登 出
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
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
  import { computed } from 'vue'
  import { setUserAuthority } from '@/api/user'
  import { fmtTitle } from '@/utils/fmtRouterTitle'
  import gvaAside from '@/view/layout/aside/index.vue'
  import Logo from '@/components/logo/index.vue'
  import { useLayoutMode } from '@/hooks/useLayoutMode'
  import { useSideWidth } from '@/hooks/useSideWidth'
  import { cn, FOCUS_RING } from '@/core/componentLibrary/utils'

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
  const toPerson = () => {
    router.push({ name: 'person' })
  }
  const matched = computed(() => route.meta.matched)

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

<style scoped lang="scss"></style>
