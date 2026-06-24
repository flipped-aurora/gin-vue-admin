<template>
  <div
    class="flex justify-between fixed top-0 right-0 z-10 h-16 text-slate-700 dark:text-slate-300 items-center px-2"
    :style="{
      left: isVertical ? headerSideWidth + 'px' : '0px',
      background: 'var(--gva-header-bg)',
      boxShadow: 'var(--gva-header-shadow)'
    }"
  >
    <div class="flex items-center cursor-pointer flex-1">
      <div
        v-if="!isVertical"
        class="flex items-center justify-center cursor-pointer"
        :class="isMobile ? '' : 'min-w-48'"
        @click="router.push({ path: '/' })"
      >
        <Logo />
        <div
          v-if="!isMobile"
          class="inline-flex font-bold text-2xl ml-2"
          :class="
            (settings.layout.mode === 'head' ||
              settings.layout.mode === 'combination') &&
            'min-w-fit'
          "
        >
          {{ $GIN_VUE_ADMIN.appName }}
        </div>
      </div>

      <el-breadcrumb
        v-show="!isMobile"
        v-if="settings.header.breadcrumb.visible && settings.layout.mode !== 'head' && settings.layout.mode !== 'combination'"
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
        v-if="settings.layout.mode === 'head' && !isMobile"
        class="flex-1"
      />
      <gva-aside
        v-if="settings.layout.mode === 'combination' && !isMobile"
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

  const userStore = useUserStore()
  const router = useRouter()
  const route = useRoute()
  const appStore = useAppStore()
  const themeStore = useThemeStore()
  const { device, sideCollapse } = storeToRefs(appStore)
  const { settings } = storeToRefs(themeStore)
  const isMobile = computed(() => {
    return device.value === 'mobile'
  })
  // 通栏侧边布局：header 让出左侧侧栏宽度，并隐藏自身 Logo
  const isVertical = computed(
    () => settings.value.layout.mode === 'vertical' && !isMobile.value
  )
  const headerSideWidth = computed(
    () => sideCollapse.value
      ? settings.value.layout.sideCollapsedWidth
      : settings.value.layout.sideWidth
  )
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
