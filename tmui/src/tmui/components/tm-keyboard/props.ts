export const propsCutom = {
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	/** 是否随机键盘 */
	random: {
		type: Boolean,
		default: false
	},
	dark: {
		type: Boolean,
		default: false
	},
	modelValue: {
		type: String,
		default: ''
	},
	color: {
		type: String,
		default: "primary"
	},
	//是否需要显示小数点。
	decimal: {
		type: Boolean,
		default: false
	},
	// 是否显示输入内容在键盘顶部。
	showInputContent: {
		type: Boolean,
		default: false
	},
	/** 最大长度 */
	maxLength: {
		type: Number,
		default: 0
	},
	round: {
		type: Number,
		default: 2
	},
	title:{
		type:String,
		default:'安全键盘放心输入'
	}
}