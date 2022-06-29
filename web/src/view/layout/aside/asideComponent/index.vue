<template>
  <component
    :is="menuComponent"
    v-if="!routerInfo.hidden"
    :is-collapse="isCollapse"
    :theme="theme"
    :router-info="routerInfo"
  >
    <template v-if="routerInfo.children&&routerInfo.children.length">
      <AsideComponent
        v-for="item in routerInfo.children"
        :key="item.name"
        :is-collapse="false"
        :router-info="item"
        :theme="theme"
      />
    </template>
  </component>
</template>

<script>
export default {
  name: 'AsideComponent'
}
</script>

<script setup>
import MenuItem from './menuItem.vue'
import AsyncSubmenu from './asyncSubmenu.vue'
import { computed } from 'vue'
const props = defineProps({
  routerInfo: {
    default: () => null,
    type: Object,
  },
  isCollapse: {
    default: () => false,
    type: Boolean
  },
  theme: {
    default: () => {},
    type: Object
  }
})

const menuComponent = computed(() => {
  if (props.routerInfo.children && props.routerInfo.children.filter(item => !item.hidden).length) {
    return AsyncSubmenu
  } else {
    return MenuItem
  }
})
</script>
