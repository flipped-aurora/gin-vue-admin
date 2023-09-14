<template>
  <div>
    <div class="w-full h-screen bg-gray-50 flex items-center justify-center">
      <div class="flex flex-col items-center text-2xl gap-4">
        <img src="../../assets/notFound.png">
        <p>页面被神秘力量吸走了（如果您是开源版请联系我们修复）</p>
        <p style="font-size:18px;line-height:40px;">常见问题为当前此角色无当前路由，如果确定要使用本路由，请到角色管理进行分配</p>
        <p>↓</p>
        <img
          src="../../assets/qm.png"
          class="w-16 h-16 mt-20"
        >
        <el-button @click="toDashboard">返回首页</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useUserStore } from '@/pinia/modules/user'
import { useRouter } from 'vue-router'

defineOptions({
  name: 'Error'
})

const userStore = useUserStore()
const router = useRouter()
const toDashboard = () => {
  router.push({ name: userStore.userInfo.authority.defaultRouter })
}
</script>
