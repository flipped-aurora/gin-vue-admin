import { PropType } from "vue"
import { actionItem } from "./interface"
export const defaultProps = {
	width: {
		type: Number,
		default: 750
	},
	height: {
		type: Number,
		default: 88
	},

	disabled: {
		type: Boolean,
		default: false
	},
	transprent: {
		type: Boolean,
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	round: {
		type: Number,
		default: 0
	},
	//当前打开的状态，可以使用v-model:open-status
	openStatus: {
		type: Boolean,
		default: false
	},
	/** 单位 rpx, */
	leftWidth: {
		type: Number,
		default: 120
	},
	/** 单位 rpx, */
	rightWidth: {
		type: Number,
		default: 220
	},
}