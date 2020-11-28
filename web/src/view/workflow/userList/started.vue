<template>
  <div>
    <el-table
      :data="tableData"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
     <el-table-column
        label="ID"
        prop="ID"
        width="60"
      ></el-table-column>
      <el-table-column
        label="流程名称"
        prop="workflowProcess.label"
        width="150"
      ></el-table-column>
      <el-table-column
        label="发起人"
        prop="promoter.nickName"
        width="120"
      ></el-table-column>
      <el-table-column label="节点日期" width="180">
        <template slot-scope="scope">{{
          scope.row.CreatedAt | formatDate
        }}</template>
      </el-table-column>
      <el-table-column
        label="业务代码"
        prop="businessType"
        width="120"
      ></el-table-column>

      <el-table-column
        label="当前节点"
        prop="workflowNode.label"
        width="120"
      ></el-table-column>
        <el-table-column
         label="流程状态"
        width="120"
        >
            <template slot-scope="scope">
                <div>
                    {{scope.row.isActive?"进行中":"已结束"}}
                </div>
            </template>
        </el-table-column>
      <el-table-column
        label="详细介绍"
        prop="workflowProcess.description"
        min-width="200"
      ></el-table-column>
      <el-table-column label="按钮组">
        <template slot-scope="scope">
            <el-button @click="view(scope.row)">查看</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { formatTimeToStr } from '@/utils/date'
import {getMyStated} from "@/api/workflowProcess"
export default {
    data(){
        return{
            tableData:[]
        }
    },
    methods:{
        view(row){
            this.$router.push({
                    name: "workflowUse",
                    query: {
                    workflowMoveID: row.ID
                    }
            })
        }
    },
    async created(){
        const res = await getMyStated()
        if(res.code == 0){
            this.tableData = res.data.wfms
        }
    },
    filters:{   
        formatDate: function(time) {
            if (time != null && time != '') {
                var date = new Date(time)
                return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
            } else {
                return ''
            }
        }
    }
}
</script>