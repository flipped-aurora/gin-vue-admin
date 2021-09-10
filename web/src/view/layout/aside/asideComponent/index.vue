<template>
  <component :is="menuComponent" v-if="!routerInfo.hidden" :router-info="routerInfo">
    <template v-if="routerInfo.children&&routerInfo.children.length">
      <AsideComponent v-for="item in routerInfo.children" :key="item.name" :router-info="item" />
    </template>
  </component>
</template>

<script>
import MenuItem from './menuItem.vue'
import AsyncSubmenu from './asyncSubmenu.vue'

export default {
  name: 'AsideComponent',
  components: {
    MenuItem,
    AsyncSubmenu
  },
  props: {
    routerInfo: {
      default: function() {
        return null
      },
      type: Object
    }
  },
  computed: {
    menuComponent() {
      if (this.routerInfo.children && this.routerInfo.children.filter(item => !item.hidden).length) {
        return 'AsyncSubmenu'
      } else {
        return 'MenuItem'
      }
    }
  }
}
</script>
