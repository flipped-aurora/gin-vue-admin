export interface actionItem {
	text?: string,//操作按钮的文本,
	width?: number,//按钮宽度
	color?: string,//主题色
	icon?: string,
	[prop: string]: any,
}