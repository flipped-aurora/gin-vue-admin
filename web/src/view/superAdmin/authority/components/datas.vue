<template>
  <div>
    <warning-bar
      title="此功能仅用于创建角色和角色的many2many关系表，具体使用还须自己结合表实现业务，详情参考示例代码（客户示例）。此功能不建议使用，建议使用插件市场【组织管理功能（点击前往）】来管理资源权限。"
      href="https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=36"
    />
    <div class="sticky top-0.5 z-10 my-4">
      <el-button class="float-left" type="primary" @click="all">全选</el-button>
      <el-button class="float-left" type="primary" @click="self"
        >本角色</el-button
      >
      <el-button class="float-left" type="primary" @click="selfAndChildren"
        >本角色及子角色</el-button
      >
      <el-button class="float-right" type="primary" @click="authDataEnter"
        >确 定</el-button
      >
    </div>
    <div class="clear-both pt-4">
      <el-checkbox-group v-model="dataAuthorityId" @change="selectAuthority">
        <el-checkbox
          v-for="(item, key) in authoritys"
          :key="key"
          :label="item"
          >{{ item.authorityName }}</el-checkbox
        >
      </el-checkbox-group>
    </div>
  </div>
</template>

<script setup>
  import { setDataAuthority } from '@/api/authority'
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'

  defineOptions({
    name: 'Datas'
  })

  const props = defineProps({
    row: {
      default: function () {
        return {}
      },
      type: Object
    },
    authority: {
      default: function () {
        return []
      },
      type: Array
    }
  })

  const authoritys = ref([])
  const needConfirm = ref(false)
  //   平铺角色
  const roundAuthority = (authoritysData) => {
    authoritysData &&
      authoritysData.forEach((item) => {
        const obj = {}
        obj.authorityId = item.authorityId
        obj.authorityName = item.authorityName
        authoritys.value.push(obj)
        if (item.children && item.children.length) {
          roundAuthority(item.children)
        }
      })
  }

  const dataAuthorityId = ref([])
  const init = () => {
    roundAuthority(props.authority)
    props.row.dataAuthorityId &&
      props.row.dataAuthorityId.forEach((item) => {
        const obj =
          authoritys.value &&
          authoritys.value.filter(
            (au) => au.authorityId === item.authorityId
          ) &&
          authoritys.value.filter(
            (au) => au.authorityId === item.authorityId
          )[0]
        dataAuthorityId.value.push(obj)
      })
  }

  init()

  // 暴露给外层使用的切换拦截统一方法
  const enterAndNext = () => {
    authDataEnter()
  }

  const emit = defineEmits(['changeRow'])
  const all = () => {
    dataAuthorityId.value = [...authoritys.value]
    emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
    needConfirm.value = true
  }
  const self = () => {
    dataAuthorityId.value = authoritys.value.filter(
      (item) => item.authorityId === props.row.authorityId
    )
    emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
    needConfirm.value = true
  }
  const selfAndChildren = () => {
    const arrBox = []
    getChildrenId(props.row, arrBox)
    dataAuthorityId.value = authoritys.value.filter(
      (item) => arrBox.indexOf(item.authorityId) > -1
    )
    emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
    needConfirm.value = true
  }
  const getChildrenId = (row, arrBox) => {
    arrBox.push(row.authorityId)
    row.children &&
      row.children.forEach((item) => {
        getChildrenId(item, arrBox)
      })
  }
  // 提交
  const authDataEnter = async () => {
    const res = await setDataAuthority(props.row)
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '资源设置成功' })
    }
  }

  //   选择
  const selectAuthority = () => {
    dataAuthorityId.value = dataAuthorityId.value.filter((item) => item)
    emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
    needConfirm.value = true
  }

  defineExpose({
    enterAndNext,
    needConfirm
  })
</script>
