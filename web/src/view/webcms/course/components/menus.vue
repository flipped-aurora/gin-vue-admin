<template>
  <div>
    <div class="clearfix sticky-button">
      <el-input v-model="filterText" class="fitler" placeholder="筛选" />
      <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
    </div>
    <div class="tree-content">
      <el-checkbox-group v-model="checkList">
        <el-checkbox v-for="item in tableData" :key="item.ID" :label="item.ID" style="height: 30px;">{{ item.title }}</el-checkbox>
      </el-checkbox-group>
    </div>
    {{ checkList }}
  </div>
</template>

<script setup>
import {
  getClassListById
} from '@/api/class'
import { nextTick, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const emit = defineEmits(['changeRow'])
const filterText = ref('')
const tableData = ref(null)
const checkList = ref([])
const init = async () => {
  // 根据课程id获取课时列表
  const table = await getClassListById({
    "course_id":props.row.ID
  })
  if (table.code === 0) {
    tableData.value = table.data.list
  }
}

init()
watch(filterText, (val) => {
  tableData.value.filter((item) => {
    if (!val) return true
    return item.title.indexOf(val) !== -1
  })
})

</script>

<script>

export default {
  name: 'Menus'
}
</script>

<style scoped>
@import "@/style/button.scss";
.tree-content{
  margin-top: 20px;
}
  .tree-content .el-checkbox {
      height: 40px;
      width: 29%;
      white-space:normal;
  }
</style>
