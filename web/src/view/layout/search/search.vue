<template>
  <div class="search-component">
    <div v-if="show" class="transition-box">
      <el-select
        ref="searchInput"
        filterable
        placeholder="请选择"
        @blur="hiddenSearch"
        @change="changeRouter"
      >
        <el-option
          v-for="item in routerStore.routerList"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </div>
    <template v-else>
      <div
        class="user-box"
      >
        <div class="gvaIcon gvaIcon-refresh" :class="[reload ? 'reloading' : '']" @click="handleReload" />
      </div>
      <div
        class="user-box"
      >
        <div class="gvaIcon gvaIcon-search" @click="showSearch" />
      </div>
      <div
        class="user-box"
      >
        <Screenfull class="search-icon" :style="{cursor:'pointer'}" />
      </div>
      <div
        class="user-box"
      >
        <div class="service gvaIcon-customer-service" @click="toService" />
      </div>
    </template>
  </div>
</template>

<script>
export default {
  name: 'BtnBox',
}
</script>

<script setup>
import Screenfull from '@/view/layout/screenfull/index.vue'
import { emitter } from '@/utils/bus.js'
import { ref, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { useRouterStore } from '@/pinia/modules/router'

const router = useRouter()

const routerStore = useRouterStore()

const changeRouter = (e) => {
  if (e.indexOf('http:') > -1 || e.indexOf('https:') > -1) {
    window.open(e)
    return
  }
  router.push({ name: e })
}

const show = ref(false)
const hiddenSearch = async() => {
  setTimeout(() => {
    show.value = false
  }, 100)
}

const searchInput = ref(null)
const showSearch = async() => {
  show.value = true
  await nextTick()
  searchInput.value.focus()
}

const reload = ref(false)
const handleReload = () => {
  reload.value = true
  emitter.emit('reload')
  setTimeout(() => {
    reload.value = false
  }, 500)
}
const toService = () => {
  window.open('https://support.qq.com/product/371961')
}

</script>
<style scoped lang="scss">
.reload{
  font-size: 18px;
}

.transition-box {
  overflow: hidden;
  width: 160px;
  margin-right: 32px;
  text-align: center;
  ::v-deep(.el-input__wrapper){
    .el-input__inner{
      border-bottom: 1px solid var(--el-color-info-light-7);
    }
    box-shadow: none !important;
  }
  ::v-deep(.el-select .el-input .el-input__wrapper.is-focus){
    box-shadow: none !important;
  }
::v-deep(.el-select .el-input.is-focus .el-input__wrapper){
    box-shadow: none !important;
  }
}


.reloading{
  animation:turn 0.5s linear infinite;
}
@keyframes turn {
  0%{transform:rotate(0deg);}
  25%{transform:rotate(90deg);}
  50%{transform:rotate(180deg);}
  75%{transform:rotate(270deg);}
  100%{transform:rotate(360deg);}
}

.service {
  font-family: "gvaIcon",serif !important;
    font-size: 16px;
    font-style: normal;
    font-weight: 800;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }
//小屏幕不显示
@media (max-width: 750px) {
  .service {
    display: none;
  }
}
</style>
