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
      :style="{display:'inline-block',float:'right',width:'31px',textAlign:'left',fontSize:'16px',paddingTop:'2px'}"
      class="user-box"
    >
      <i :style="{cursor:'pointer'}" class="el-icon-refresh" @click="$bus.$emit('reload')" />
    </div>
    <div :style="{display:'inline-block',float:'right'}" class="user-box">
      <i :style="{cursor:'pointer'}" class="el-icon-search search-icon" @click="showSearch()" />
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: 'SearchComponent',
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
    },
    showSearch() {
      this.show = true
      this.$nextTick(() => {
        this.$refs['search-input'].focus()
      })
    }
  }
}
</script>
