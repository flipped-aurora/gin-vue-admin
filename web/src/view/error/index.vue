<template>
  <div>
    <div class="w-full h-screen bg-gray-50 flex items-center justify-center">
      <div class="flex flex-col items-center text-2xl gap-4">
        <img class="w-1/3" src="../../assets/404.png" />
        <p class="text-lg">页面被神秘力量吸走了</p>
        <p class="text-lg">
          常见问题为当前此角色无当前路由，如果确定要使用本路由，请到角色管理进行分配
        </p>
        <p>
          项目地址：<a
            href="https://github.com/flipped-aurora/gin-vue-admin"
            target="_blank"
            class="text-blue-600"
            >https://github.com/flipped-aurora/gin-vue-admin</a
          >
        </p>
        <el-button @click="toDashboard">返回首页</el-button>
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
  const toDashboard = () => {
    try {
      router.push({ name: userStore.userInfo.authority.defaultRouter })
    } catch (error) {
        emitter.emit('show-error', {
        code: '401',
        message: "检测到其他用户修改了路由权限，请重新登录",
        fn: () => {
          userStore.ClearStorage()
          router.push({ name: 'Login', replace: true })
        }
      })
    }
  }
</script>
