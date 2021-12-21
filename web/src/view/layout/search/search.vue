<template>
  <div class="search-component">
    <transition name="el-fade-in-linear">
      <div v-show="show" class="transition-box" style="display: inline-block; ">
        <el-select
          ref="search-input"
          v-model="value"
          filterable
          placeholder="请选择"
          @blur="hiddenSearch"
          @change="changeRouter"
        >
          <el-option
            v-for="item in routerList"
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
      <el-icon class="search-icon" @click="showSearch()">
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
import { mapGetters } from 'vuex'
import { emitter } from '@/utils/bus.js'
import Screenfull from '@/view/layout/screenfull/index.vue'

export default {
  name: 'SearchComponent',
  components: {
    Screenfull
  },
  data() {
    return {
      value: '',
      show: false,
      btnShow: true,
      reload: false
    }
  },
  computed: {
    ...mapGetters('router', ['routerList']),
  },

  methods: {
    changeRouter() {
      this.$router.push({ name: this.value })
      this.value = ''
    },
    hiddenSearch() {
      this.show = false
      setTimeout(() => {
        this.btnShow = true
      }, 500)
    },
    showSearch() {
      this.btnShow = false
      this.show = true
      this.$nextTick(() => {
        this.$refs['search-input'].focus()
      })
    },
    handleReload() {
      this.reload = true
      emitter.emit('reload')
      setTimeout(() => {
        this.reload = false
      }, 500)
    },
    toService() {
      window.open('https://support.qq.com/product/371961')
    }
  }
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
