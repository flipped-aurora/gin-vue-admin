<template>
  <div>
    <div class="sticky top-0.5 z-10 bg-white">
      <el-input
        v-model="filterText"
        class="w-3/5"
        placeholder="筛选"
      />
      <el-button
        class="float-right"
        type="primary"
        @click="authApiEnter"
      >确 定</el-button>
    </div>
    <div class="tree-content">
      <el-scrollbar>
        <el-tree
          ref="apiTree"
          :data="apiTreeData"
          :default-checked-keys="apiTreeIds"
          :props="apiDefaultProps"
          default-expand-all
          highlight-current
          node-key="onlyId"
          show-checkbox
          :filter-node-method="filterNode"
          @check="nodeChange"
        />
      </el-scrollbar>
    </div>
  </div>
</template>

<script setup>
import { getAllApis } from '@/api/api'
import { UpdateCasbin, getPolicyPathByAuthorityId } from '@/api/casbin'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Apis',
})

const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  }
})

const apiDefaultProps = ref({
  children: 'children',
  label: 'description'
})
const filterText = ref('')
const apiTreeData = ref([])
const apiTreeIds = ref([])
const activeUserId = ref('')
const init = async() => {
  const res2 = await getAllApis()
  const apis = res2.data.apis

  apiTreeData.value = buildApiTree(apis)
  const res = await getPolicyPathByAuthorityId({
    authorityId: props.row.authorityId
  })
  activeUserId.value = props.row.authorityId
  apiTreeIds.value = []
  res.data.paths && res.data.paths.forEach(item => {
    apiTreeIds.value.push('p:' + item.path + 'm:' + item.method)
  })
}

init()

const needConfirm = ref(false)
const nodeChange = () => {
  needConfirm.value = true
}
// 暴露给外层使用的切换拦截统一方法
const enterAndNext = () => {
  authApiEnter()
}

// 创建api树方法
const buildApiTree = (apis) => {
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
}

// 关联关系确定
const apiTree = ref(null)
const authApiEnter = async() => {
  const checkArr = apiTree.value.getCheckedNodes(true)
  var casbinInfos = []
  checkArr && checkArr.forEach(item => {
    var casbinInfo = {
      path: item.path,
      method: item.method
    }
    casbinInfos.push(casbinInfo)
  })
  const res = await UpdateCasbin({
    authorityId: activeUserId.value,
    casbinInfos
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'api设置成功' })
  }
}

defineExpose({
  needConfirm,
  enterAndNext
})

const filterNode = (value, data) => {
  if (!value) return true
  return data.description.indexOf(value) !== -1
}
watch(filterText, (val) => {
  apiTree.value.filter(val)
})

</script>

