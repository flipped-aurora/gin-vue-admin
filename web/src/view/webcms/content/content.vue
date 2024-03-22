<template>
  <div>
    <div class="gva-table-box">
        <el-row :gutter="20">
          <el-col :span="5" style="border: 1px solid #f5f7fa;" v-show="showMenuModel">
            <div class="gva-btn-list">
              <el-button type="primary">栏目列表</el-button>
              <el-checkbox v-model="menuShowAll" label="展开全部" size="large"  style="margin-left: 10px;"/>
            </div>
            <el-table
                  ref="cateMenuTable"
                  style="width: 100%"
                  tooltip-effect="dark"
                  :data="cateMenusList"
                  row-key="ID"
                  :default-expand-all="menuShowAll"
                  :expand-row-keys="[]"
                  @row-click="getSelectionRows"
                  highlight-current-row
                  >
                  <el-table-column align="center" label="层级">
                    <template v-slot="scope">
                      <el-tag v-if="scope.row.menuLevel === 0">一级</el-tag>
                      <el-tag type="success" v-else-if="scope.row.menuLevel === 1"
                        >二级</el-tag>
                      <el-tag type="warning" v-else>三级</el-tag>
                    </template>
                  </el-table-column>
                  <el-table-column align="center" label="栏目名称" prop="name" />
            </el-table>
          </el-col>
          <el-col :span="tableSpan" style="border: 1px solid #f5f7fa;">
            <div class="gva-btn-list">
              <el-button type="primary" icon="fold" @click="onshowMenuModel"></el-button>
            </div>
            <!-- 动态插件开始 -->
            <component :is="componentTag" :defaultCatid="catid"></component>
            <!-- 动态插件结束 -->
          </el-col>
        </el-row>
        
    </div>
  </div>
</template>

<script>
export default {
  name: 'Content'
}
</script>

<script setup>
import {
  getCateMenusList,
} from '@/api/cateMenus'


// 引入模型文件
import Class from '@/view/webcms/class/class.vue'
import Recruit from '@/view/webcms/recruit/recruit.vue'

import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive,shallowRef,onBeforeUnmount,markRaw } from 'vue'


// 自动化生成的字典（可能为空）以及字段
const formData = ref({

})

// =========== 表格控制部分 ===========
const cateMenuTable = ref(null)
const cateMenusList = ref([])
const menuShowAll = ref(false)
const showMenuModel = ref(true)
const tableSpan = ref(19)
const catid = ref('0')
let componentTag = shallowRef(Class)

// 选中菜单
const getSelectionRows= (row, column, event)=>{
  // console.log(row, column, event)
  // 选择栏目高亮
  cateMenuTable.value.setCurrentRow(row)
  // 设置右侧表格数据
  catid.value = row.ID.toString()
  switch(row.modeType){
    case 1: //招聘模型
    componentTag = shallowRef(Recruit)
    break;
    case 2:
    componentTag = shallowRef(Class)
    default:
    componentTag = shallowRef(Class)
  }
  
}

// 显示隐藏菜单栏
const onshowMenuModel = ()=>{
  showMenuModel.value = !showMenuModel.value
  if(!showMenuModel.value){
    tableSpan.value = 24
  }else{
    tableSpan.value = 19
  }
}

const getCateMenusListData = async() =>{
  const table = await getCateMenusList({})
  if (table.code === 0) {
    cateMenusList.value = table.data.list
  }
}
getCateMenusListData()

</script>
<style >
  .gva-table-box{
    padding: 0px;
  }
.gva-search-box{
  padding: 0px;
}
</style>
