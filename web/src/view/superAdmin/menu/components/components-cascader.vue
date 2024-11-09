<template>
  <div class="flex justify-between items-center gap-2 w-full">
    <el-cascader
      v-if="pathIsSelect"
      placeholder="请选择文件路径"
      :options="pathOptions"
      v-model="activeComponent"
      filterable
      class="w-full"
      clearable
      @change="emitChange"
    />
    <el-input
      v-else
      v-model="tempPath"
      placeholder="页面:view/xxx/xx.vue 插件:plugin/xx/xx.vue"
      @change="emitChange"
    />
    <el-button @click="togglePathIsSelect"
      >{{ pathIsSelect ? '手动输入' : '快捷选择' }}
    </el-button>
  </div>
</template>

<script setup>
  import { onMounted, ref, watch } from 'vue'
  import pathInfo from '@/pathInfo.json'

  const props = defineProps({
    component: {
      type: String,
      default: ''
    }
  })

  const emits = defineEmits(['change'])

  const pathOptions = ref([])
  const tempPath = ref('')
  const activeComponent = ref([])
  const pathIsSelect = ref(true)

  const togglePathIsSelect = () => {
    if (pathIsSelect.value) {
      tempPath.value = activeComponent.value?.join('/') || ''
    } else {
      activeComponent.value = tempPath.value?.split('/') || []
    }

    pathIsSelect.value = !pathIsSelect.value
    emitChange()
  }

  function convertToCascaderOptions(data) {
    const result = []

    for (const path in data) {
      const label = data[path]
      const parts = path.split('/').filter(Boolean)

      // 如果第一个部分是 'src'，则从第二个部分开始处理
      const startIndex = parts[0] === 'src' ? 1 : 0

      let currentLevel = result

      for (let i = startIndex; i < parts.length; i++) {
        const part = parts[i]
        let node = currentLevel.find((item) => item.value === part)

        if (!node) {
          node = {
            value: part,
            label: part,
            children: []
          }
          currentLevel.push(node)
        }

        if (i === parts.length - 1) {
          // 如果是路径的最后一部分，设置标签并移除 children
          node.label = label
          delete node.children
        }

        currentLevel = node.children || []
      }
    }

    return result
  }

  watch(
    () => props.component,
    (value) => {
      initCascader(value)
    }
  )

  onMounted(() => {
    pathOptions.value = convertToCascaderOptions(pathInfo)
    initCascader(props.component)
  })

  const initCascader = (value) => {
    // 新增的时候
    if (value === '') {
      pathIsSelect.value = true
      return
    }

    // 编辑的时候，根据路径判断是选择框还是输入框
    if (pathInfo[`/src/${value}`]) {
      activeComponent.value = value.split('/').filter(Boolean)
      tempPath.value = ''
      pathIsSelect.value = true
      return
    }
    tempPath.value = value
    activeComponent.value = []
    pathIsSelect.value = false
  }

  const emitChange = () => {
    emits(
      'change',
      pathIsSelect.value ? activeComponent.value?.join('/') : tempPath.value
    )
  }
</script>

<style scoped lang="scss"></style>
