<template>
  <div class="search-component">
    <transition name="el-fade-in-linear">
      <div v-show="show" class="transition-box" style="display: inline-block; ">
        <el-select
          ref="searchInput"
          v-model="value"
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
    </transition>
    <div
      v-if="btnShow"
      class="user-box"
    >
      <el-icon class="reload" :class="[reload ? 'reloading' : '']" @click="handleReload">
        <refresh />
      </el-icon>
    </div>
    <div
      v-if="btnShow"
      class="user-box"
    >
      <el-icon class="search-icon" @click="showSearch">
        <search />
      </el-icon>
    </div>
    <div
      v-if="btnShow"
      class="user-box"
    >
      <Screenfull class="search-icon" :style="{cursor:'pointer'}" />
    </div>
    <div
      v-if="btnShow"
      class="user-box"
    >
      <el-icon class="search-icon" @click="toService"><service /></el-icon>
    </div>
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

const value = ref('')
const changeRouter = () => {
  router.push({ name: value.value })
  value.value = ''
}

const show = ref(false)
const btnShow = ref(true)
const hiddenSearch = () => {
  show.value = false
  setTimeout(() => {
    btnShow.value = true
  }, 500)
}

const searchInput = ref(null)
const showSearch = async() => {
  btnShow.value = false
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

.reloading{
  animation:turn 0.5s linear infinite;
}
@keyframes turn {
  0%{-webkit-transform:rotate(0deg);}
  25%{-webkit-transform:rotate(90deg);}
  50%{-webkit-transform:rotate(180deg);}
  75%{-webkit-transform:rotate(270deg);}
  100%{-webkit-transform:rotate(360deg);}
}
</style>
