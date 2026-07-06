<template>
  <!-- 工具图标组：扁平图标，无独立描边/阴影，统一 32px 点击区域，hover 出现浅色高亮。
       el-tooltip__trigger / el-dropdown 工具类统一拉伸为 32px 高，保证所有图标垂直居中对齐。 -->
  <div class="flex items-center gap-1 mx-4 [&_.el-tooltip__trigger]:inline-flex [&_.el-tooltip__trigger]:items-center [&_.el-tooltip__trigger]:h-8 [&_.el-dropdown]:inline-flex [&_.el-dropdown]:items-center [&_.el-dropdown]:h-8">
    <el-tooltip v-if="isDev" effect="dark" content="视频教程" placement="bottom">
      <el-dropdown @command="toDoc">
        <span class="gva-tool-btn">
          <el-icon class="text-[18px] leading-none">
            <Film />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item
              v-for="item in videoList"
              :key="item.link"
              :command="item.link"
              >{{ item.title }}</el-dropdown-item
            >
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-tooltip>

    <el-tooltip
      v-if="settings.header.search.visible"
      effect="dark"
      content="搜索"
      placement="bottom"
    >
      <span class="gva-tool-btn" @click="handleCommand">
        <el-icon class="text-[18px] leading-none">
          <Search />
        </el-icon>
      </span>
    </el-tooltip>

    <el-tooltip effect="dark" content="系统设置" placement="bottom">
      <span class="gva-tool-btn" @click="toggleSetting">
        <el-icon class="text-[18px] leading-none">
          <Setting />
        </el-icon>
      </span>
    </el-tooltip>

    <el-tooltip
      v-if="settings.header.refresh.visible"
      effect="dark"
      content="刷新"
      placement="bottom"
    >
      <span class="gva-tool-btn" @click="toggleRefresh">
        <el-icon :class="showRefreshAnmite ? 'animate-spin' : ''" class="text-[18px] leading-none">
          <Refresh />
        </el-icon>
      </span>
    </el-tooltip>

    <el-tooltip effect="dark" content="切换主题" placement="bottom">
      <span
        class="gva-tool-btn"
        @click="themeStore.toggleTheme(!themeStore.isDark)"
      >
        <el-icon v-if="themeStore.isDark" class="text-[18px] leading-none">
          <Sunny />
        </el-icon>
        <el-icon v-else class="text-[18px] leading-none">
          <Moon />
        </el-icon>
      </span>
    </el-tooltip>

    <gva-setting v-model:drawer="showSettingDrawer"></gva-setting>
    <command-menu ref="command" />
  </div>
</template>

<script setup>
  import { useThemeStore } from '@/pinia'
  import { storeToRefs } from 'pinia'
  import GvaSetting from '@/view/layout/setting/index.vue'
  import { ref } from 'vue'
  import { emitter } from '@/utils/bus.js'
  import CommandMenu from '@/components/commandMenu/index.vue'
  import { toDoc } from '@/utils/doc'
  import { isDev } from '@/utils/env.js'

  const themeStore = useThemeStore()
  const { settings } = storeToRefs(themeStore)
  const showSettingDrawer = ref(false)
  const showRefreshAnmite = ref(false)
  const toggleRefresh = () => {
    showRefreshAnmite.value = true
    emitter.emit('reload')
    setTimeout(() => {
      showRefreshAnmite.value = false
    }, 1000)
  }

  const toggleSetting = () => {
    showSettingDrawer.value = true
  }

  const first = ref('')
  const command = ref()

  const handleCommand = () => {
    command.value.open()
  }
  const initPage = () => {
    // 判断当前用户的操作系统
    if (window.localStorage.getItem('osType') === 'WIN') {
      first.value = 'Ctrl'
    } else {
      first.value = '⌘'
    }
    // 当用户同时按下ctrl和k键的时候
    const handleKeyDown = (e) => {
      if (e.ctrlKey && e.key === 'k') {
        // 阻止浏览器默认事件
        e.preventDefault()
        handleCommand()
      }
    }
    window.addEventListener('keydown', handleKeyDown)
  }

  initPage()

  const videoList = [
    {
      title: '1.clone项目和安装依赖',
      link: 'https://www.bilibili.com/video/BV1jx4y1s7xx'
    },
    {
      title: '2.初始化项目',
      link: 'https://www.bilibili.com/video/BV1sr421K7sv'
    },
    {
      title: '3.开启调试工具+创建初始化包',
      link: 'https://www.bilibili.com/video/BV1iH4y1c7Na'
    },
    {
      title: '4.手动使用自动化创建功能',
      link: 'https://www.bilibili.com/video/BV1UZ421T7fV'
    },
    {
      title: '5.使用已有表格创建业务',
      link: 'https://www.bilibili.com/video/BV1NE4m1977s'
    },
    {
      title: '6.使用AI创建业务和创建数据源模式的可选项',
      link: 'https://www.bilibili.com/video/BV17i421a7DE'
    },
    {
      title: '7.创建自己的后端方法',
      link: 'https://www.bilibili.com/video/BV1Yw4m1k7fg'
    },
    {
      title: '8.新增一个前端页面',
      link: 'https://www.bilibili.com/video/BV12y411i7oE'
    },
    {
      title: '9.配置一个前端二级页面',
      link: 'https://www.bilibili.com/video/BV1ZM4m1y7i3'
    },
    {
      title: '10.配置一个前端菜单参数',
      link: 'https://www.bilibili.com/video/BV1WS42197DZ'
    },
    {
      title: '11.菜单参数实战+动态菜单标题+菜单高亮配置',
      link: 'https://www.bilibili.com/video/BV1NE4m1979c'
    },
    {
      title: '12.增加菜单可控按钮',
      link: 'https://www.bilibili.com/video/BV1Sw4m1k746'
    },
    {
      title: '14.新增客户角色和其相关配置教学',
      link: 'https://www.bilibili.com/video/BV1Ki421a7X2'
    },
    {
      title: '15.发布项目上线',
      link: 'https://www.bilibili.com/video/BV1Lx4y1s77D'
    }
  ]
</script>
