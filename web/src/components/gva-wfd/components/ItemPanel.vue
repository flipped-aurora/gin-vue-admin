<template>
    <div class="itemPanel" :style="{'height': height+'px'}">
        <el-collapse v-model="activeNames" >
            <el-collapse-item :title="i18n['start']" name="1">
                <img data-item="{clazz:'start',size:'50*50',label:''}"
                     :src="require('../assets/flow/start.svg')" style="width:42px;height:42px" />
                <div>{{i18n['startEvent']}}</div>
                <!-- <img data-item="{clazz:'timerStart',size:'50*50',label:''}"
                     :src="require('../assets/flow/timer-start.svg')" style="width:42px;height:42px" />
                <div>{{i18n['timerEvent']}}</div> -->
                <!-- <img data-item="{clazz:'messageStart',size:'50*50',label:''}"
                     :src="require('../assets/flow/message-start.svg')" style="width:42px;height:42px" />
                <div>{{i18n['messageEvent']}}</div>
                <img data-item="{clazz:'signalStart',size:'50*50',label:''}"
                     :src="require('../assets/flow/signal-start.svg')" style="width:42px;height:42px" />
                <div>{{i18n['signalEvent']}}</div> -->
            </el-collapse-item>
            <el-collapse-item :title="i18n['task']" name="2">
                <img :data-item="userTaskData"
                     :src="require('../assets/flow/user-task.svg')" style="width:80px;height:44px" />
                <div>{{i18n['userTask']}}</div>
                <!-- <img :data-item="scriptTaskData"
                     :src="require('../assets/flow/script-task.svg')" style="width:80px;height:44px" />
                <div>{{i18n['scriptTask']}}</div> -->
                <!-- <img :data-item="javaTaskData"
                     :src="require('../assets/flow/java-task.svg')" style="width:80px;height:44px" />
                <div>{{i18n['javaTask']}}</div> -->
                <!-- <img :data-item="mailTaskData"
                     :src="require('../assets/flow/mail-task.svg')" style="width:80px;height:44px" />
                <div>{{i18n['mailTask']}}</div>
                <img :data-item="receiveTaskData"
                     :src="require('../assets/flow/receive-task.svg')" style="width:80px;height:44px" />
                <div>{{i18n['receiveTask']}}</div> -->
            </el-collapse-item>
            <el-collapse-item :title="i18n['gateway']" name="3">
                <img data-item="{clazz:'exclusiveGateway',size:'60*60',label:''}"
                     :src="require('../assets/flow/exclusive-gateway.svg')" style="width:48px;height:48px" />
                <div>{{i18n['exclusiveGateway']}}</div>
                <img data-item="{clazz:'parallelGateway',size:'60*60',label:''}"
                     :src="require('../assets/flow/parallel-gateway.svg')" style="width:48px;height:48px" />
                <div>{{i18n['parallelGateway']}}</div>
                <img data-item="{clazz:'inclusiveGateway',size:'60*60',label:''}"
                     :src="require('../assets/flow/inclusive-gateway.svg')" style="width:48px;height:48px" />
                <div>{{i18n['inclusiveGateway']}}</div>
            </el-collapse-item>
            <!-- <el-collapse-item :title="i18n['catch']" name="4">
                <img data-item="{clazz:'timerCatch',size:'60*35',label:''}"
                     :src="require('../assets/flow/timer-catch.svg')" style="width:58px;height:38px" />
                <div>{{i18n['timerEvent']}}</div>
                <img data-item="{clazz:'messageCatch',size:'60*35',label:''}"
                     :src="require('../assets/flow/message-catch.svg')" style="width:58px;height:38px" />
                <div>{{i18n['messageEvent']}}</div>
                <img data-item="{clazz:'signalCatch',size:'60*35',label:''}"
                     :src="require('../assets/flow/signal-catch.svg')" style="width:58px;height:38px" />
                <div>{{i18n['signalEvent']}}</div>
            </el-collapse-item> -->
            <el-collapse-item :title="i18n['end']" name="5">
                <img data-item="{clazz:'end',size:'50*50',label:''}"
                     :src="require('../assets/flow/end.svg')" style="width:42px;height:42px" />
                <div>{{i18n['endEvent']}}</div>
            </el-collapse-item>
        </el-collapse>
        
    </div>
</template>
<script>
  export default {
    inject: ['i18n'],
    props: {
      height: {
        type: Number,
        default: 800,
      },
    },
    data() {
      return {
        activeNames: [],
        userTaskData: "{clazz:'userTask',size:'80*44',label:'"+this.i18n['userTask']+"'}",
        scriptTaskData: "{clazz:'scriptTask',size:'80*44',label:'"+this.i18n['scriptTask']+"'}",
        // javaTaskData: "{clazz:'javaTask',size:'80*44',label:'"+this.i18n['javaTask']+"'}",
        mailTaskData: "{clazz:'mailTask',size:'80*44',label:'"+this.i18n['mailTask']+"'}",
        receiveTaskData: "{clazz:'receiveTask',size:'80*44',label:'"+this.i18n['receiveTask']+"'}",
      };
    },
  }
</script>

<style lang="scss" >
    .itemPanel {
        float: left;
        width: 10%;
        background: #f0f2f5;
        overflow-y: auto;
        border-left: 1px solid #E9E9E9;
        border-bottom: 1px solid #E9E9E9;
        img{
            width: 92px;
            height: 96px;
            padding: 4px;
            border: 1px solid rgba(0,0,0,0);
            border-radius: 2px;
            &:hover{
                border: 1px solid #ccc;
                cursor: move;
            }
        }
        .el-collapse {
            border: 0;
            .el-collapse-item {
                > div[role=tab] > div {
                    padding-left: 10px;
                    border: 1px solid #E9E9E9;
                    border-left:0;
                }
                &:first-child{
                    > div[role=tab] > div {
                        border-top: 0;
                    }
                }
                &:last-child{
                    > div[role=tab] > div {
                        border-bottom: 1px solid #E9E9E9;
                    }
                }
                .el-collapse-item__wrap{
                    border-top: 0;
                    background: #f0f2f5;
                    text-align: center;
                }
            }
        }

    }
</style>
