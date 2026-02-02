<template>
  <el-select
    v-model="internalValue"
    :multiple="multiple"
    :placeholder="placeholder"
    :clearable="false"
    @change="handleChange"
    class="w-full"
  >
    <el-option
      v-for="item in categories"
      :key="item.ID"
      :label="item.name"
      :value="item.name"
    />
  </el-select>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { getCertCategoryList } from '../api/certCategory'

const props = defineProps({
  modelValue: {
    type: [String, Array],
    default: ''
  },
  multiple: {
    type: Boolean,
    default: false
  },
  placeholder: {
    type: String,
    default: '请选择项目'
  },
  autoSelectFirst: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

const categories = ref([])
const internalValue = ref(props.modelValue)

// 监听外部值变化
watch(() => props.modelValue, (newVal) => {
  internalValue.value = newVal
})

// 获取分类列表
const fetchCategories = async () => {
  const res = await getCertCategoryList({ page: 1, pageSize: 999 })
  if (res.code === 0) {
    categories.value = res.data.list
    // 如果没有选中值且开启了自动选中第一项，且分类列表不为空
    if (!internalValue.value && props.autoSelectFirst && categories.value.length > 0) {
      const firstValue = props.multiple ? [categories.value[0].name] : categories.value[0].name
      internalValue.value = firstValue
      emit('update:modelValue', firstValue)
      emit('change', firstValue)
    }
  }
}

// 处理值变化
const handleChange = (val) => {
  emit('update:modelValue', val)
  emit('change', val)
}

onMounted(() => {
  fetchCategories()
})
</script>
