<template>
  <div class="search-component">
    <div
      class="gvaIcon gvaIcon-refresh"
      :class="[reload ? 'reloading' : '']"
      @click="handleReload"
    />
    <Screenfull class="search-icon" />
    <!-- <div
      class="gvaIcon gvaIcon-customer-service"
      @click="toService"
    /> -->
    <div class="user-box">
      <el-dropdown>
        <span class="el-dropdown-link">
          {{ selectsite.company }}
          <el-icon class="el-icon--right">
            <arrow-down />
          </el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item v-for="item,index in sitelist" @click="changeSite(item.id)" :key="index">{{ item.company }}</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import {
  getWebconfig,
  changeWebconfig
} from '@/api/webconfig'
import Screenfull from '@/view/layout/screenfull/index.vue'
import { ElMessage } from 'element-plus'
import { emitter } from '@/utils/bus.js'
import { ref } from 'vue'

defineOptions({
  name: 'BtnBox',
})

const reload = ref(false)
const handleReload = () => {
  reload.value = true
  emitter.emit('reload')
  setTimeout(() => {
    reload.value = false
  }, 500)
}
// const toService = () => {
//   window.open('https://support.qq.com/product/371961')
// }
const sitelist = ref([])
const selectsite = ref({})
const getsitelist = async() => {
  const data = await getWebconfig()
  sitelist.value = data.data.list
  selectsite.value = sitelist.value.find((item)=>data.data.siteid == data.data.siteid)
}
getsitelist()

// 切换站点
const changeSite = async (siteid) => {
  const res = await changeWebconfig(siteid)
  // console.log(res)
  if (res.code === 0) {
    ElMessage.success(res.msg)
    location.reload()
  }
}

</script>
<style scoped lang="scss">

.search-component {
  @apply inline-flex overflow-hidden text-center gap-5 mr-5;
  div{
    @apply cursor-pointer;
  }
  .el-input__inner {
    @apply border-b border-solid border-gray-300;
  }
  .el-dropdown-link {
    @apply cursor-pointer;
  }
}

.reload {
  font-size: 18px;
}

.reloading{
  animation:turn 0.5s linear infinite;
}

@keyframes turn {
  0% {
    transform: rotate(0deg);
  }

  25% {
    transform: rotate(90deg);
  }

  50% {
    transform: rotate(180deg);
  }

  75% {
    transform: rotate(270deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

</style>
