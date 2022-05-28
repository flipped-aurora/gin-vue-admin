<template>
  <el-sub-menu ref="subMenu" :index="routerInfo.name">
    <template #title>
      <el-icon v-if="routerInfo.meta.icon">
        <component :is="routerInfo.meta.icon" />
      </el-icon>
      <span>{{ routerInfo.meta.title }}</span>
    </template>
    <slot />
  </el-sub-menu>
</template>

<script>
export default {
  name: 'AsyncSubmenu',
}
</script>

<script setup>
import { ref, watch } from 'vue'
const props = defineProps({
  routerInfo: {
    default: function() {
      return null
    },
    type: Object
  },
  isCollapse: {
    default: function() {
      return false
    },
    type: Boolean
  },
  theme: {
    default: function() {
      return {}
    },
    type: Object
  }
})

// const activeBackground = ref(props.theme.activeBackground)
// const activeText = ref(props.theme.activeText)
const normalText = ref(props.theme.normalText)
// const hoverBackground = ref(props.theme.hoverBackground)
// const hoverText = ref(props.theme.hoverText)

watch(() => props.theme, () => {
  // activeBackground.value = props.theme.activeBackground
  // activeText.value = props.theme.activeText
  normalText.value = props.theme.normalText
  // hoverBackground.value = props.theme.hoverBackground
  // hoverText.value = props.theme.hoverText
})

</script>

<style lang="scss" scoped>
.el-sub-menu{
  ::v-deep(.el-sub-menu__title){
      color: v-bind(normalText);
  }
}

</style>
