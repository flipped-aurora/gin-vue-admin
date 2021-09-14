<template>
  <div>
    <div class="clearflex" style="margin:18px">
      <el-button class="fl-right" size="mini" type="primary" @click="authDataEnter">确 定</el-button>
      <el-button class="fl-left" size="mini" type="primary" @click="all">全选</el-button>
      <el-button class="fl-left" size="mini" type="primary" @click="self">本角色</el-button>
      <el-button class="fl-left" size="mini" type="primary" @click="selfAndChildren">本角色及子角色</el-button>
    </div>
    <el-checkbox-group v-model="dataAuthorityId" @change="selectAuthority">
      <el-checkbox v-for="(item,key) in authoritys" :key="key" :label="item">{{ item.authorityName }}</el-checkbox>
    </el-checkbox-group>
    <span>
      此功能仅用于创建角色和角色的many2many关系表，具体使用还须自己结合表实现业务，详情参考示例代码（客户示例）
    </span>
  </div>
</template>

<script>
import { setDataAuthority } from '@/api/authority'
export default {
  name: 'Datas',
  props: {
    row: {
      default: function() {
        return {}
      },
      type: Object
    },
    authority: {
      default: function() {
        return []
      },
      type: Array
    }
  },
  data() {
    return {
      authoritys: [],
      dataAuthorityId: [],
      needConfirm: false
    }
  },
  created() {
    this.authoritys = []
    this.dataAuthorityId = []
    this.roundAuthority(this.authority)
    this.row.dataAuthorityId && this.row.dataAuthorityId.map(item => {
      const obj = this.authoritys && this.authoritys.filter(au => au.authorityId === item.authorityId) && this.authoritys.filter(au => au.authorityId === item.authorityId)[0]
      this.dataAuthorityId.push(obj)
    })
  },
  methods: {
    // 暴露给外层使用的切换拦截统一方法
    enterAndNext() {
      this.authDataEnter()
    },
    all() {
      this.dataAuthorityId = [...this.authoritys]
      this.$emit('changeRow', 'dataAuthorityId', this.dataAuthorityId)
      this.needConfirm = true
    },
    self() {
      this.dataAuthorityId = this.authoritys.filter(item => item.authorityId === this.row.authorityId)
      this.$emit('changeRow', 'dataAuthorityId', this.dataAuthorityId)
      this.needConfirm = true
    },
    selfAndChildren() {
      const arrBox = []
      this.getChildrenId(this.row, arrBox)
      this.dataAuthorityId = this.authoritys.filter(item => arrBox.indexOf(item.authorityId) > -1)
      this.$emit('changeRow', 'dataAuthorityId', this.dataAuthorityId)
      this.needConfirm = true
    },
    getChildrenId(row, arrBox) {
      arrBox.push(row.authorityId)
      row.children && row.children.map(item => {
        this.getChildrenId(item, arrBox)
      })
    },
    // 提交
    async authDataEnter() {
      const res = await setDataAuthority(this.row)
      if (res.code === 0) {
        this.$message({ type: 'success', message: '资源设置成功' })
      }
    },
    //   平铺角色
    roundAuthority(authoritys) {
      authoritys && authoritys.map(item => {
        const obj = {}
        obj.authorityId = item.authorityId
        obj.authorityName = item.authorityName
        this.authoritys.push(obj)
        if (item.children && item.children.length) {
          this.roundAuthority(item.children)
        }
      })
    },
    //   选择
    selectAuthority() {
      this.$emit('changeRow', 'dataAuthorityId', this.dataAuthorityId)
      this.needConfirm = true
    }
  }
}
</script>
