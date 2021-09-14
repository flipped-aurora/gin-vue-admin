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

<script>
import { getBaseMenuTree, getMenuAuthority, addMenuAuthority } from '@/api/menu'
import {
  updateAuthority
} from '@/api/authority'
export default {
  name: 'Menus',
  props: {
    row: {
      default: function() {
        return {}
      },
      type: Object
    }
  },
  data() {
    return {
      menuTreeData: [],
      menuTreeIds: [],
      needConfirm: false,
      menuDefaultProps: {
        children: 'children',
        label: function(data) {
          return data.meta.title
        }
      }
    }
  },
  async created() {
    // 获取所有菜单树
    const res = await getBaseMenuTree()
    this.menuTreeData = res.data.menus

    const res1 = await getMenuAuthority({ authorityId: this.row.authorityId })
    const menus = res1.data.menus
    const arr = []
    menus.map(item => {
      // 防止直接选中父级造成全选
      if (!menus.some(same => same.parentId === item.menuId)) {
        arr.push(Number(item.menuId))
      }
    })
    this.menuTreeIds = arr
  },
  methods: {
    async setDefault(data) {
      const res = await updateAuthority({ authorityId: this.row.authorityId, AuthorityName: this.row.authorityName, parentId: this.row.parentId, defaultRouter: data.name })
      if (res.code === 0) {
        this.$message({ type: 'success', message: '设置成功' })
        this.$emit('changeRow', 'defaultRouter', res.data.authority.defaultRouter)
      }
    },
    nodeChange() {
      this.needConfirm = true
    },
    // 暴露给外层使用的切换拦截统一方法
    enterAndNext() {
      this.relation()
    },
    // 关联树 确认方法
    async relation() {
      const checkArr = this.$refs.menuTree.getCheckedNodes(false, true)
      const res = await addMenuAuthority({
        menus: checkArr,
        authorityId: this.row.authorityId
      })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '菜单设置成功!'
        })
      }
    }
  }
}
</script>
