<template>
  <div>
      <div class="clearflex" style="margin:18px">
      <el-button @click="authDataEnter" class="fl-right" size="small" type="primary">确 定</el-button>
      <el-button @click="all" class="fl-left" size="small" type="primary">全选</el-button>
      <el-button @click="self" class="fl-left" size="small" type="primary">本角色</el-button>
      <el-button @click="selfAndChildren" class="fl-left" size="small" type="primary">本角色及子角色</el-button>
    </div>
     <el-checkbox-group v-model="dataAuthorityId" @change="selectAuthority">
        <el-checkbox v-for="(item,key) in authoritys" :label="item" :key="key">{{item.authorityName}}</el-checkbox>
    </el-checkbox-group>
  </div>
</template>
<script>
import {setDataAuthority} from '@/api/authority'
export default {
  name: 'Datas',
  data() {
    return {
        authoritys:[],
        dataAuthorityId:[],
        needConfirm:false
    }
  },
  props: {
    row: {
      default: function() {
        return {}
      },
      type: Object
    },
    authority: {
      default: function() {
        return {}
      },
      type: Array
    }
  },
  methods:{
    // 暴露给外层使用的切换拦截统一方法
      enterAndNext(){
        this.authDataEnter()
      },
      all(){
         this.dataAuthorityId = [...this.authoritys]
         this.row.dataAuthorityId = this.dataAuthorityId
         this.needConfirm = true
      },
      self(){
          this.dataAuthorityId = this.authoritys.filter(item=>item.authorityId===this.row.authorityId)
          this.row.dataAuthorityId = this.dataAuthorityId
          this.needConfirm = true
      },
      selfAndChildren(){
         const arrBox = []
         this.getChildrenId(this.row,arrBox)
         this.dataAuthorityId = this.authoritys.filter(item=>arrBox.indexOf(item.authorityId)>-1)
         this.row.dataAuthorityId = this.dataAuthorityId
         this.needConfirm = true
      },
      getChildrenId(row,arrBox){
          arrBox.push(row.authorityId)
          row.children&&row.children.map(item=>{
              this.getChildrenId(item,arrBox)
          })
      },
    // 提交
      async authDataEnter(){
          const res = await setDataAuthority(this.row)
          if(res.code == 0){
              this.$message({ type: 'success', message: "资源设置成功" })
          }
      },
    //   平铺角色
      roundAuthority(authoritys){
          authoritys&&authoritys.map(item=>{
              const obj = {}
              obj.authorityId = item.authorityId
              obj.authorityName = item.authorityName
              this.authoritys.push(obj)
              if(item.children&&item.children.length){
                  this.roundAuthority(item.children)
              }
          })
      },
    //   选择
      selectAuthority(){
          this.row.dataAuthorityId = this.dataAuthorityId
          this.needConfirm = true
      }
  },
  created() {
      this.authoritys = []
      this.dataAuthorityId = []
      this.roundAuthority(this.authority)
      this.row.dataAuthorityId&&this.row.dataAuthorityId.map(item=>{
          const obj = this.authoritys&&this.authoritys.filter(au=>au.authorityId === item.authorityId)&&this.authoritys.filter(au=>au.authorityId === item.authorityId)[0]
          this.dataAuthorityId.push(obj)
      })
  }
}
</script>
<style lang="less">
</style>