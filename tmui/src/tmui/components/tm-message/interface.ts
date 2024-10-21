export type modelType = "load" | "error" | "info" | "warn" | "quest" | "success" | "disabled" | "wait"
export interface config {
	duration?: number, //显示的时长
	icon?: string, //显示的图标
	text?: string, //显示的文本
	color?: string, //显示的主题
	dark?: boolean, //是否暗黑
	model?: modelType,//模式见上方modelType
	mask?: boolean //是否显示黑色遮罩。
}