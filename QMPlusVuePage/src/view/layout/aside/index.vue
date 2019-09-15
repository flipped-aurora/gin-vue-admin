<template>
  <el-menu
    :class="['el-menu-vertical',!isCollapse&&'noCollapse']"
    :collapse="isCollapse"
    :default-active="active"
    @select="selectMenuItem"
    unique-opened
  >
    <AsideComponent :key="item.name" :routerInfo="item" v-for="item in asyncRouters[0].children" />
  </el-menu>
</template>

<script>
import { mapGetters } from 'vuex'
import AsideComponent from '@/view/layout/aside/asideComponent'
export default {
  name: 'Aside',
  data() {
    return {
      active: '',
      isCollapse: true
    }
  },
  methods: {
    selectMenuItem(index) {
      if (index === this.$route.name) return
      this.$router.push({ name: index })
    }
  },
  computed: {
    ...mapGetters('router', ['asyncRouters'])
  },
  components: {
    AsideComponent
  },
  created() {
    this.active = this.$route.name
  }
}
</script>
