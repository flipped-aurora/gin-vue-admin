<template>
  <div>
    <div class="sticky top-0.5 z-10 bg-white pb-2">
      <div class="flex gap-2 items-center mb-2">
        <el-input v-model="filterText" class="flex-1" placeholder="筛选" />
        <el-button type="primary" @click="relation">确 定</el-button>
      </div>
      <div class="flex items-center gap-2">
        <span class="whitespace-nowrap">默认首页：</span>
        <el-select
          :model-value="row.defaultRouter"
          filterable
          placeholder="请选择默认首页"
          class="flex-1"
          @change="handleDefaultRouterChange"
        >
          <el-option
            v-for="item in menuOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </div>
    </div>
    <div class="tree-content clear-both">
      <el-scrollbar>
        <el-tree
          ref="menuTree"
          :data="menuTreeData"
          :default-checked-keys="menuTreeIds"
          :props="menuDefaultProps"
          default-expand-all
          highlight-current
          node-key="ID"
          show-checkbox
          :filter-node-method="filterNode"
          @check="nodeChange"
        >
          <template #default="{ node, data }">
            <div class="flex items-center gap-2">
              <span>{{ node.label }}</span>
                <SvgIcon v-if="row.defaultRouter === data.name" icon="ant-design:home-filled" class="inline text-lg text-active" />
              <span v-if="data.menuBtn.length">
                <el-button type="primary" link @click.stop="() => OpenBtn(data)">
                  分配按钮
                </el-button>
              </span>
            </div>
          </template>
        </el-tree>
      </el-scrollbar>
    </div>
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
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
  import {
    getBaseMenuTree,
    getMenuAuthority,
    addMenuAuthority
  } from '@/api/menu'
  import { updateAuthority } from '@/api/authority'
  import { getAuthorityBtnApi, setAuthorityBtnApi } from '@/api/authorityBtn'
  import { nextTick, ref, watch } from 'vue'
  import { ElMessage } from 'element-plus'

  defineOptions({
    name: 'Menus'
  })

  const props = defineProps({
    row: {
      default: function () {
        return {}
      },
      type: Object
    }
  })

  const emit = defineEmits(['changeRow'])
  const filterText = ref('')
  const menuTreeData = ref([])
  const menuTreeIds = ref([])
  const needConfirm = ref(false)
  const menuDefaultProps = ref({
    children: 'children',
    label: function (data) {
      return data.meta.title
    },
    disabled: function (data) {
      return props.row.defaultRouter === data.name
    }
  })

  const menuOptions = ref([])

  const flattenMenus = (menus) => {
    let result = []
    menus.forEach(item => {
      if (item.children && item.children.length > 0) {
        result = result.concat(flattenMenus(item.children))
      } else {
        if (!item.name?.startsWith('http://') && !item.name?.startsWith('https://')) {
          result.push({
            label: item.meta.title,
            value: item.name
          })
        }
      }
    })
    return result
  }

  const init = async () => {
    // 获取所有菜单树
    const res = await getBaseMenuTree()
    menuTreeData.value = res.data.menus
    menuOptions.value = flattenMenus(res.data.menus)
    const res1 = await getMenuAuthority({ authorityId: props.row.authorityId })
    const menus = res1.data.menus
    const arr = []
    menus.forEach((item) => {
      // 防止直接选中父级造成全选
      if (!menus.some((same) => same.parentId === item.menuId)) {
        arr.push(Number(item.menuId))
      }
    })
    menuTreeIds.value = arr
  }

  init()

  const setDefault = async (data) => {
    const res = await updateAuthority({
      authorityId: props.row.authorityId,
      AuthorityName: props.row.authorityName,
      parentId: props.row.parentId,
      defaultRouter: data.name
    })
    if (res.code === 0) {
      relation()
      emit('changeRow', 'defaultRouter', res.data.authority.defaultRouter)
    }
  }

  const handleDefaultRouterChange = (val) => {
    setDefault({ name: val })
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
  const relation = async () => {
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
  const OpenBtn = async (data) => {
    menuID = data.ID
    const res = await getAuthorityBtnApi({
      menuID: menuID,
      authorityId: props.row.authorityId
    })
    if (res.code === 0) {
      openDialog(data)
      await nextTick()
      if (res.data.selected) {
        res.data.selected.forEach((id) => {
          btnData.value.some((item) => {
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
  const enterDialog = async () => {
    const selected = multipleSelection.value.map((item) => item.ID)
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

  const filterNode = (value, data) => {
    if (!value) return true
    // console.log(data.mate.title)
    return data.meta.title.indexOf(value) !== -1
  }

  watch(filterText, (val) => {
    menuTree.value.filter(val)
  })
</script>
