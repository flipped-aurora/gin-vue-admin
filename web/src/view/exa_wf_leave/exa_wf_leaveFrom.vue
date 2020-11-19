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
        <el-button @click="start" type="primary">启动</el-button>
        <el-button @click="back" type="primary">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {
    startWorkflow,
} from "@/api/workflowProcess"; 
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
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
    ...mapGetters("user", ["userInfo"])
  },
  methods: {
    async start() {
      const res = await startWorkflow({
            business:this.formData,
            wf:{
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
   if(this.business){
     this.formData = this.business
   }
}
};
</script>

<style>
</style>