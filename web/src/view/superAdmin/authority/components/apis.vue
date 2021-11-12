<template>
  <div>
    <div class="clearflex">
      <el-button class="fl-right" size="mini" type="primary" @click="authApiEnter">确 定</el-button>
    </div>
    <el-tree
      ref="apiTree"
      :data="apiTreeData"
      :default-checked-keys="apiTreeIds"
      :props="apiDefaultProps"
      default-expand-all
      highlight-current
      node-key="onlyId"
      show-checkbox
      @check="nodeChange"
    />
  </div>
</template>

<script>
import { getAllApis } from '@/api/api'
import { UpdateCasbin, getPolicyPathByAuthorityId } from '@/api/casbin'
export default {
  name: 'Apis',
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
      apiTreeData: [],
      apiTreeIds: [],
      needConfirm: false,
      apiDefaultProps: {
        children: 'children',
        label: 'description'
      }
    }
  },
  async created() {
    // 获取api并整理成树结构
    const res2 = await getAllApis()
    const apis = res2.data.apis

    this.apiTreeData = this.buildApiTree(apis)
    const res = await getPolicyPathByAuthorityId({
      authorityId: this.row.authorityId
    })
    this.activeUserId = this.row.authorityId
    this.apiTreeIds = []
    res.data.paths && res.data.paths.forEach(item => {
      this.apiTreeIds.push('p:' + item.path + 'm:' + item.method)
    })
  },
  methods: {
    nodeChange() {
      this.needConfirm = true
    },
    // 暴露给外层使用的切换拦截统一方法
    enterAndNext() {
      this.authApiEnter()
    },
    // 创建api树方法
    buildApiTree(apis) {
      const apiObj = {}
      apis &&
        apis.forEach(item => {
          item.onlyId = 'p:' + item.path + 'm:' + item.method
          if (Object.prototype.hasOwnProperty.call(apiObj, item.apiGroup)) {
            apiObj[item.apiGroup].push(item)
          } else {
            Object.assign(apiObj, { [item.apiGroup]: [item] })
          }
        })
      const apiTree = []
      for (const key in apiObj) {
        const treeNode = {
          ID: key,
          description: key + '组',
          children: apiObj[key]
        }
        apiTree.push(treeNode)
      }
      return apiTree
    },
    // 关联关系确定
    async authApiEnter() {
      const checkArr = this.$refs.apiTree.getCheckedNodes(true)
      var casbinInfos = []
      checkArr && checkArr.forEach(item => {
        var casbinInfo = {
          path: item.path,
          method: item.method
        }
        casbinInfos.push(casbinInfo)
      })
      const res = await UpdateCasbin({
        authorityId: this.activeUserId,
        casbinInfos
      })
      if (res.code === 0) {
        this.$message({ type: 'success', message: 'api设置成功' })
      }
    }
  }
}
</script>
