<template>
  <div class="search-component">
    <transition name="el-fade-in-linear">
      <div class="transition-box" style="display: inline-block; " v-show="show">
        <el-select
          @blur="hiddenSearch"
          @change="changeRouter"
          filterable
          placeholder="请选择"
          v-model="value"
        >
          <el-option
            :key="item.value"
            :label="item.label"
            :value="item.value"
            v-for="item in routerList"
          ></el-option>
        </el-select>
      </div>
    </transition>
    <div :style="{display:'inline-block'}" class="user-box">
      <i @click="show = !show" class="el-icon-search search-icon"></i>
    </div>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'

export default {
  name: 'searchComponent',
  data() {
    return {
      value: '',
      show: false
    }
  },
  computed: {
    ...mapGetters('router', ['routerList'])
  },
  methods: {
    changeRouter() {
      this.$router.push({ name: this.value })
      this.value = ''
    },
    hiddenSearch() {
      this.show = false
    }
  }
}
</script>
<style lang="scss">

</style>