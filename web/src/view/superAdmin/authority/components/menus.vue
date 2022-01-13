<template>
  <div>
    <div class="clearflex">
      <el-button class="fl-right" size="mini" type="primary" @click="relation">确 定</el-button>
    </div>
    <el-tree
      ref="menuTree"
      :data="menuTreeData"
      :default-checked-keys="menuTreeIds"
      :props="menuDefaultProps"
      default-expand-all
      highlight-current
      node-key="ID"
      show-checkbox
      @check="nodeChange"
    >
      <template #default="{ node , data }">
        <span class="custom-tree-node">
          <span>{{ node.label }}</span>
          <span>
            <el-button
              type="text"
              size="mini"
              :style="{color:row.defaultRouter === data.name?'#E6A23C':'#85ce61'}"
              :disabled="!node.checked"
              @click="() => setDefault(data)"
            >
              {{ row.defaultRouter === data.name?"首页":"设为首页" }}
            </el-button>
          </span>
        </span>
      </template>
    </el-tree>
  </div>
</template>

<script setup>
import { getBaseMenuTree, getMenuAuthority, addMenuAuthority } from '@/api/menu'
import {
  updateAuthority
} from '@/api/authority'
import { ref } from 'vue'
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

const menuTreeData = ref([])
const menuTreeIds = ref([])
const needConfirm = ref(false)
const menuDefaultProps = ref({
  children: 'children',
  label: function(data) {
    return data.meta.title
  }
})

const init = async() => {
  // 获取所有菜单树
  const res = await getBaseMenuTree()
  menuTreeData.value = res.data.menus

  const res1 = await getMenuAuthority({ authorityId: props.row.authorityId })
  const menus = res1.data.menus
  const arr = []
  menus.forEach(item => {
    // 防止直接选中父级造成全选
    if (!menus.some(same => same.parentId === item.menuId)) {
      arr.push(Number(item.menuId))
    }
  })
  menuTreeIds.value = arr
}

init()

const setDefault = async(data) => {
  const res = await updateAuthority({ authorityId: props.row.authorityId, AuthorityName: props.row.authorityName, parentId: props.row.parentId, defaultRouter: data.name })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '设置成功' })
    emit('changeRow', 'defaultRouter', res.data.authority.defaultRouter)
  }
}
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  relation()
}
// 关联树 确认方法
const menuTree = ref(null)
const relation = async() => {
  const checkArr = menuTree.value.getCheckedNodes(false, true)
  const res = await addMenuAuthority({
    menus: checkArr,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '菜单设置成功!'
    })
  }
}

defineExpose({ enterAndNext, needConfirm })

</script>

<script>

export default {
  name: 'Menus'
}
</script>
