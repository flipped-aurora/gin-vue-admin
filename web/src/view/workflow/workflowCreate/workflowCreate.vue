<template>
  <div>
    <el-button
      size="small"
      style="float:right;margin-top:6px;margin-right:6px;"
      @click="saveXML"
    >导出XML</el-button>
    <el-button
      size="small"
      style="float:right;margin-top:6px;margin-right:6px;"
      @click="saveImg"
    >导出图片</el-button>
    <el-popover placement="bottom" width="160" v-model="visible">
      <p>确认流程无误并保存吗</p>
      <div style="text-align: right; margin: 0">
        <el-button size="mini" type="text" @click="visible = false">取消</el-button>
        <el-button type="primary" size="mini" @click="save">确定</el-button>
      </div>
      <el-button size="small" slot="reference" style="float:right;margin-top:6px;margin-right:6px;">保存流程</el-button>
    </el-popover>

    <gva-wfd
      ref="wfd"
      :data="demoData"
      :height="600"
      :users="users"
      :authorities="authorities"
      :groups="groups"
      :categorys="categorys"
      :lang="lang"
    />
  </div>
</template>
<script>
import {
    findWorkflowProcess
} from "@/api/workflowProcess";  //  此处请自行替换地址
import gvaWfd from "@/components/gva-wfd";
import { getUserList } from "@/api/user";
import { getAuthorityList } from "@/api/authority";
export default {
  name: "Workflow",
  components: {
    gvaWfd
  },
  data() {
    return {
      visible: false,
      lang: "zh",
      demoData: {
        nodes: [
          {
            clazz: "start",
            label: "发起请假",
            type: "start-node",
            shape: "start-node",
            x: 110,
            y: 195,
            id: "start1603681292875",
            style: {}
          },
          {
            clazz: "parallelGateway",
            label: "会签",
            type: "parallel-gateway-node",
            shape: "parallel-gateway-node",
            x: 228,
            y: 195,
            id: "parallelGateway1603681296419",
            style: {}
          },
          {
            clazz: "userTask",
            label: "审批人1",
            type: "user-task-node",
            shape: "user-task-node",
            x: 372,
            y: 84,
            id: "userTask1603681299962",
            style: {},
            assignValue: 1,
            assignType: "user"
          },
          {
            clazz: "userTask",
            label: "审批人2",
            type: "user-task-node",
            shape: "user-task-node",
            x: 370,
            y: 321,
            id: "userTask1603681302372",
            style: {},
            assignValue: 2,
            assignType: "user"
          },
          {
            clazz: "parallelGateway",
            label: "会签结果检测",
            type: "parallel-gateway-node",
            shape: "parallel-gateway-node",
            x: 519,
            y: 195,
            id: "parallelGateway1603681338222",
            style: {}
          },
          {
            clazz: "end",
            label: "请假失败",
            type: "end-node",
            shape: "end-node",
            x: 704,
            y: 317,
            id: "end1603681358043",
            style: {}
          },
          {
            clazz: "end",
            label: "请假成功",
            type: "end-node",
            shape: "end-node",
            x: 706.5,
            y: 55.5,
            id: "end1603681360882",
            style: {}
          }
        ],
        edges: [
          {
            id: "flow1603681320738",
            clazz: "flow",
            source: "parallelGateway1603681296419",
            target: "userTask1603681299962",
            sourceAnchor: 0,
            targetAnchor: 3,
            shape: "flow-polyline-round",
            style: {},
            startPoint: { x: 228, y: 169, index: 0 },
            endPoint: { x: 321.5, y: 84, index: 3 }
          },
          {
            id: "flow1603681321969",
            clazz: "flow",
            source: "parallelGateway1603681296419",
            target: "userTask1603681302372",
            sourceAnchor: 2,
            targetAnchor: 3,
            shape: "flow-polyline-round",
            style: {},
            startPoint: { x: 228, y: 221, index: 2 },
            endPoint: { x: 319.5, y: 321, index: 3 }
          },
          {
            id: "flow1603681323274",
            clazz: "flow",
            source: "start1603681292875",
            target: "parallelGateway1603681296419",
            sourceAnchor: 1,
            targetAnchor: 3,
            shape: "flow-polyline-round",
            style: {},
            startPoint: { x: 138, y: 195, index: 1 },
            endPoint: { x: 202, y: 195, index: 3 },
            label: "发起",
            conditionExpression: "complete"
          },
          {
            id: "flow1603681341777",
            clazz: "flow",
            source: "userTask1603681299962",
            target: "parallelGateway1603681338222",
            sourceAnchor: 1,
            targetAnchor: 3,
            shape: "flow-polyline-round",
            style: {},
            startPoint: { x: 422.5, y: 84, index: 1 },
            endPoint: { x: 493, y: 195, index: 3 }
          },
          {
            id: "flow1603681343425",
            clazz: "flow",
            source: "userTask1603681302372",
            target: "parallelGateway1603681338222",
            sourceAnchor: 1,
            targetAnchor: 3,
            shape: "flow-polyline-round",
            style: {},
            startPoint: { x: 420.5, y: 321, index: 1 },
            endPoint: { x: 493, y: 195, index: 3 }
          },
          {
            id: "flow1603681362913",
            clazz: "flow",
            source: "parallelGateway1603681338222",
            target: "end1603681360882",
            sourceAnchor: 0,
            targetAnchor: 2,
            shape: "flow-polyline-round",
            style: {},
            startPoint: { x: 519, y: 169, index: 0 },
            endPoint: { x: 678.5, y: 55.5, index: 2 },
            conditionExpression: "complete",
            label: "所有人同意"
          },
          {
            id: "flow1603681392729",
            clazz: "flow",
            source: "parallelGateway1603681338222",
            target: "end1603681358043",
            sourceAnchor: 2,
            targetAnchor: 2,
            shape: "flow-polyline-round",
            style: {},
            startPoint: { x: 519, y: 221, index: 2 },
            endPoint: { x: 676, y: 317, index: 2 },
            conditionExpression: "reject",
            label: "任何一人拒绝"
          }
        ],
        combos: [],
        groups: []
      },
      users: [],
      authorities: [],
      groups: [
        { id: "1", name: "组1" },
        { id: "2", name: "组2" },
        { id: "3", name: "组3" }
      ],
      categorys: [
        { id: "1", name: "分类1" },
        { id: "2", name: "分类2" },
        { id: "3", name: "分类3" },
        { id: "4", 分组: "分组4" }
      ]
    };
  },
  methods: {
    save() {
      this.visible = false;
      console.log(this.$refs["wfd"].graph.save());
      console.log(this.$refs["wfd"].processModel);
      console.log(JSON.stringify(this.$refs["wfd"].graph.save()));
    },
    saveXML() {
      console.log(this.$refs["wfd"].graph.saveXML());
    },
    saveImg() {
      console.log(this.$refs["wfd"].graph.saveImg());
    }
  },
  async created() {
    const userRes = await getUserList({ page: 1, pageSize: 9999999 });
    if (userRes.code == 0) {
      userRes.data.list.map(item => {
        this.users.push({ id: item.ID, name: item.nickName });
      });
    }
    const authorityRes = await getAuthorityList({ page: 1, pageSize: 9999999 });
    if (authorityRes.code == 0) {
      authorityRes.data.list.map(item => {
        this.authorities.push({
          id: item.authorityId,
          name: item.authorityName
        });
      });
    }
    if(this.$route.params.ID){
      const res = await findWorkflowProcess({ ID: this.$route.params.ID });
      console.log(res.data)
    }
  }
};
</script>