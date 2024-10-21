import type { showDetail } from "./interface"
import { PropType } from "vue"
export const propsOpts = {
	/**
	 * 这里是动态返回时间戳。这是一个标准的时间，不管showDetail是如何设置都将不影响这里的输出。
	 * 并且fomart不会影响这里的双向绑定。只会影响model-str
	 */
	modelValue: {
		type: [Number, String, Date],
		default: ''
	},
	//这里和modelValue不一样，它只代表格式化输出显示，因此这里可能并不是一个有效的时间值。
	/**
	 * 比如:format为"MM/DD",那这里就会显示12/10这样的时间格式，因此并不是一个正确的时间，
	 * 这里主要是为了方便表单上页面的显示控制输入。如果真要保存到数据库，你应该保存modelValue的值。
	 */
	modelStr: {
		type: [String],
		default: ''
	},
	defaultValue: {
		type: [Number, String, Date],
		default: ''
	},
	//禁用的部分日期，禁用的日期将不会被选中，就算滑到了该位置，也会回弹到之前的时间。
	/**
	 * 现在暂时只禁用到天，也就是一个时间到天这如果==下面的禁用日期，就会选不中。
	 */
	disabledDate: {
		type: Array as PropType<Array<Number | String | Date>>,
		default: (): Array<Number | String | Date> => []
	},
	//展示格式。最终影响到modelStr输出格式的内容。
	format: {
		type: String,
		default: "YYYY/MM/DD HH:mm:ss"
	},
	immediateChange: {
		type: Boolean,
		default: false
	},
	//需要展现的时间格式类型
	showDetail: {
		type: Object as PropType<showDetail>,
		default: () => {
			return {
				year: true,
				month: true,
				day: true,
				hour: false,
				minute: false,
				second: false
			}
		}
	},
	//日期的后缀，
	showSuffix: {
		type: Object,
		default: () => {
			return {
				year: '年',
				month: '月',
				day: '日',
				hour: '时',
				minute: '分',
				second: '秒'
			}
		}
	},
	start: {
		type: [Number, String, Date],
		default: '2008/01/01 00:00:00'
	},
	end: {
		type: [Number, String, Date],
		default: ''
	},
	height: {
		type: Number,
		default: 300
	}
}