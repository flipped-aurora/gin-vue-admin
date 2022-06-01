<template>
  <div>
    <div class="clearflex">
      <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
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
              type="primary" link
              size="small"
              :style="{color:row.defaultRouter === data.name?'#E6A23C':'#85ce61'}"
              :disabled="!node.checked"
              @click="() => setDefault(data)"
            >
              {{ row.defaultRouter === data.name?"首页":"设为首页" }}
            </el-button>
          </span>
          <span v-if="data.menuBtn.length">
            <el-button
              type="primary" link
              size="small"
              @click="() => OpenBtn(data)"
            >
              分配按钮
            </el-button>
          </span>
        </span>
      </template>
    </el-tree>

    <el-dialog v-model="btnVisible" title="分配按钮" destroy-on-close>
      <el-table
        ref="btnTableRef"
        :data="btnData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column label="按钮名称" prop="name" />
        <el-table-column label="按钮备注" prop="desc" />
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { getBaseMenuTree, getMenuAuthority, addMenuAuthority } from '@/api/menu'
import {
  updateAuthority
} from '@/api/authority'
import { getAuthorityBtnApi, setAuthorityBtnApi } from '@/api/authorityBtn'
import { nextTick, ref } from 'vue'
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

const btnVisible = ref(false)

const btnData = ref([])
const multipleSelection = ref([])
const btnTableRef = ref()
let menuID = ''
const OpenBtn = async(data) => {
  menuID = data.ID
  const res = await getAuthorityBtnApi({ menuID: menuID, authorityId: props.row.authorityId })
  if (res.code === 0) {
    openDialog(data)
    await nextTick()
    if (res.data.selected) {
      res.data.selected.forEach(id => {
        btnData.value.some(item => {
          if (item.ID === id) {
            btnTableRef.value.toggleRowSelection(item, true)
          }
        })
      })
    }
  }
}

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const openDialog = (data) => {
  btnVisible.value = true
  btnData.value = data.menuBtn
}

const closeDialog = () => {
  btnVisible.value = false
}
const enterDialog = async() => {
  const selected = multipleSelection.value.map(item => item.ID)
  const res = await setAuthorityBtnApi({
    menuID,
    selected,
    authorityId: props.row.authorityId
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '设置成功' })
    btnVisible.value = false
  }
}

</script>

<script>

export default {
  name: 'Menus'
}
</script>

<style lang="scss" scope>
.custom-tree-node{
  span+span{
    margin-left: 12px;
  }
}
</style>
