<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="请假原因:">
        <el-input
          v-model="formData.cause"
          clearable
          placeholder="请输入"
        ></el-input>
      </el-form-item>

      <el-form-item label="开始时间:">
        <el-date-picker
          type="date"
          placeholder="选择日期"
          v-model="formData.startTime"
          clearable
        ></el-date-picker>
      </el-form-item>

      <el-form-item label="结束时间:">
        <el-date-picker
          type="date"
          placeholder="选择日期"
          v-model="formData.endTime"
          clearable
        ></el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button
          v-if="this.wf.clazz == 'start'"
          @click="start"
          type="primary"
          >启动</el-button
        >
        <el-button
          v-if="this.wf.clazz == 'userTask'"
          @click="complete('yes')"
          type="primary"
          >同意</el-button
        >
        <el-button
          v-if="this.wf.clazz == 'userTask'"
          @click="complete('no')"
          type="primary"
          >拒绝</el-button
        >
        <el-button
          @click="back"
          type="primary"
          >返回</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {
    startWorkflow,
    completeWorkflowMove
} from "@/api/workflowProcess"; 
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
import Axios from 'axios';
export default {
  name: "ExaWfLeave",
  mixins: [infoList],
  props:{
    business:{
       type:Object,
      default:function(){return null}
    },
    wf:{
      type:Object,
      default:function(){return{}}
    },
    workflowMoveID:{
      type:Number,
      default:0
    }
  },
  data() {
    return {
      type: "",
      formData: {
            cause:"",
            startTime:new Date(),
            endTime:new Date(),
            
      }
    };
  },
  computed:{
    canShow(){
      if(this.wf.assignType == "user"){
        if(this.wf.assginValue.indexOf(","+this.userInfo.ID+",")>0){
          return true
        }else{
          return false
        }
      }else if(this.wf.assign_type == "authority"){
        if(this.wf.assginValue.indexOf(","+this.userInfo.authorityId+",")>0){
          return true
        }else{
          return false
        }
      }
    },
    ...mapGetters("user", ["userInfo"])
  },
  methods: {
    async start() {
      const res = await startWorkflow({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessId:0,
              businessType:"leave",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"create",
              param:""
              }
          });
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"启动成功"
        })
       this.back()
      }
    },
    async complete(param){
     const res = await completeWorkflowMove({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessID:this.formData.ID,
              businessType:"leave",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"complete",
              param:param
              }
     })
     if(res.code == 0){
       this.$message({
          type:"success",
          message:"提交"
       })
       this.back()
     }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
      console.log(this.workflowMoveID)

   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
   if(this.business){
     this.formData = this.business
   }
}

};
</script>

<style>
</style>