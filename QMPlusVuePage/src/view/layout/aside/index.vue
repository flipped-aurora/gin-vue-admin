<template>
  <div>
    <div class="menu-total">
        <i class="el-icon-arrow-right"></i>
    </div>
    <vue-scroll :ops="ops">
      <el-menu
        :class="['el-menu-vertical',!isCollapse&&'noCollapse']"
        :collapse="isCollapse"
        :default-active="active"
        @select="selectMenuItem"
        unique-opened
      >
        <aside-component
          :key="item.name"
          :routerInfo="item"
          v-for="item in asyncRouters[0].children"
        />
      </el-menu>
    </vue-scroll>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import AsideComponent from '@/view/layout/aside/asideComponent'
import vueScroll from 'vuescroll'
import 'vuescroll/dist/vuescroll.css'
export default {
  name: 'Aside',
  data() {
    return {
      active: '',
      isCollapse: true,
      ops: {
        bar: {disable:true}
      }
    }
  },
  methods: {
    selectMenuItem(index,indexList) {
      if (index === this.$route.name) return
      this.$router.push({ name: index })
    }
  },
  computed: {
    ...mapGetters('router', ['asyncRouters'])
  },
  components: {
    vueScroll,
    AsideComponent
  },
  created() {
    this.active = this.$route.name
    this.$bus.on('totalCollapse', () => {
      this.isCollapse = !this.isCollapse
    })
  },
  beforeDestroy() {
    this.$bus.off('totalCollapse')
  }
}
</script>

<style lang="scss">
.menu-info {
  .menu-contorl {
    line-height: 52px;
    font-size: 20px;
    display: table-cell;
    vertical-align: middle;
  }
}
</style>