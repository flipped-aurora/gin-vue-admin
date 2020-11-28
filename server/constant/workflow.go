package constant

//定义clazz常量

const (
	USER_TASK string = "userTask" // 审批节点
	//SCRIPT_TASK       string = "scriptTask"  //脚本节点
	//RECEIVE_TASK      string = "receiveTask"
	//MAIL_TASK         string = "mailTask" //邮件节点
	//TIMER_START       string = "timerStart" // 定时节点
	//MESSAGE_START     string = "messageStart" // 消息节点
	EXCLUSIVE_GATEWAY string = "exclusiveGateway" // 排他网关
	INCLUSIVE_GATEWAY string = "inclusiveGateway" // 包容网关
	PARELLEL_GATEWAY  string = "parallelGateway"  // 并行网关
	FLOW              string = "flow"             // 连线
	START             string = "start"            // 开始节点
	END               string = "end"              // 结束节点
	//PROCESS           string = "process" // 基础节点
)
