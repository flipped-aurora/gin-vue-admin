<template>
  <div>
      <div class="clearflex" style="margin:18px">
      <el-button @click="authDataEnter" class="fl-right" size="small" type="primary">确 定</el-button>
      <el-button @click="all" class="fl-left" size="small" type="primary">全选</el-button>
      <el-button @click="self" class="fl-left" size="small" type="primary">本部门</el-button>
      <el-button @click="selfAndChildren" class="fl-left" size="small" type="primary">部门及以下</el-button>
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
        dataAuthorityId:[]
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
      all(){
         this.dataAuthorityId = [...this.authoritys]
      },
      self(){
          this.dataAuthorityId = this.authoritys.filter(item=>item.ID===this.row.ID)
      },
      selfAndChildren(){
         const arrBox = []
         this.getChildrenId(this.row,arrBox)
         this.dataAuthorityId = this.authoritys.filter(item=>arrBox.indexOf(item.ID)>-1)
      },
      getChildrenId(row,arrBox){
          arrBox.push(row.ID)
          row.children&&row.children.map(item=>{
              this.getChildrenId(item,arrBox)
          })
      },
    // 提交
      async authDataEnter(){
          const res = await setDataAuthority(this.row)
          if(res.success){
              this.$message({ type: 'success', message: res.msg })
          }
      },
    //   平铺角色
      roundAuthority(authoritys){
          authoritys&&authoritys.map(item=>{
              const obj = {}
              obj.ID = item.ID
              obj.authorityName = item.authorityName
              this.authoritys.push(obj)
              if(item.children.length){
                  this.roundAuthority(item.children)
              }
          })
      },
    //   选择
      selectAuthority(){
          this.row.dataAuthorityId = this.dataAuthorityId
      }
  },
  created() {
      this.authoritys = []
      this.dataAuthorityId = []
      this.roundAuthority(this.authority)
      this.row.dataAuthorityId&&this.row.dataAuthorityId.map(item=>{
          const obj = this.authoritys&&this.authoritys.filter(au=>au.ID === item.ID)&&this.authoritys.filter(au=>au.ID === item.ID)[0]
          this.dataAuthorityId.push(obj)
      })
  }
}
</script>
<style lang="less">
</style>