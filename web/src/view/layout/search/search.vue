<template>
  <div class="search-component">
    <transition name="el-fade-in-linear">
      <div class="transition-box" style="display: inline-block; " v-show="show">
        <el-select
          ref="search-input"
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
    <div :style="{display:'inline-block',float:'right'}" class="user-box">
      <i @click="showSearch()" class="el-icon-search search-icon"></i>
    </div>
  </div>
</template>
<script>
import { mapGetters } from "vuex";

export default {
  name: "searchComponent",
  data() {
    return {
      value: "",
      show: false
    };
  },
  computed: {
    ...mapGetters("router", ["routerList"])
  },
  methods: {
    changeRouter() {
      this.$router.push({ name: this.value });
      this.value = "";
    },
    hiddenSearch() {
      this.show = false;
    },
    showSearch() {
      this.show = true;
      this.$nextTick(()=>{
        this.$refs['search-input'].focus()
      })
    }
  }
};
</script>
<style lang="scss">
</style>