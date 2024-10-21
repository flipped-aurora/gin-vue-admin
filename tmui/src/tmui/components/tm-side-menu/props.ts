import { ref, getCurrentInstance, nextTick, Ref, computed, PropType, watch } from "vue"
interface listItem {
	text: string,
	dotCount?: number | string
	[prop: string]: any
}
export const propsCustom = {
	height: {
		type: Number,
		default: 500
	},
	width: {
		type: Number,
		default: 750
	},
	sideWidth: {
		type: Number,
		default: 190
	},
	itemHeight: {
		type: Number,
		default: 100
	},
	fontSize: {
		type: Number,
		default: 26
	},
	active: {
		type: Number,
		default: 0
	},
	menuFontColor: {
		type: String,
		default: 'grey'
	},
	menuActiveFontColor: {
		type: String,
		default: 'primary'
	},
	menuActiveBgColor: {
		type: String,
		default: 'white'
	},
	/** 整个左侧导航的背景 */
	menuBarBgColor: {
		type: String,
		default: 'grey-4'
	},
	bodyBgColor: {
		type: String,
		default: 'white'
	},
	//只针对整体的宽和高，项目的高还是rpx.
	unit: {
		type: String,
		default: 'rpx'
	},
	list: {
		type: Array as PropType<Array<{ text?: string | number, [key: string]: any }>>,
		default: () => [],
		required: true
	},
	refresh: {
		type: Boolean,
		default: false
	},
	/**触发加载函数 */
	load: {
		type: Function as PropType<(type: 'pull' | 'bottom' | 'menu', item: listItem, index: number) => void>,
		default: () => null
	},
	/**首次是否进行load数据加载触发状态. */
	firstLoad: {
		type: Boolean,
		default: true
	},
	/**是否禁用触底刷新功能. */
	disabledBottom: {
		type: Boolean,
		default: false
	},
	/**是否禁用下拉刷新功能. */
	disabledPull: {
		type: Boolean,
		default: false
	},
	rangKey: {
		type: String,
		default: 'text'
	},
	/** 是否自定义右侧滚动区域的插槽，如果是原有的下拉，下拉的滚动组件将不显示并去除。 */
	isScroll: {
		type: Boolean,
		default: false
	}
}