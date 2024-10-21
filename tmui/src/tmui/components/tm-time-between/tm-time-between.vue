<template>
	<tm-sheet :margin="[0, 0]" :padding="[0, 0]">
		<view v-if="props.quickBtn.length > 0">
			<tm-text _class="py-24  px-16" color="grey" label="快捷方式"></tm-text>
			<view class="flex flex-row flex-row-top-start flex-wrap">
				<tm-tag
					:color="props.color"
					@click="quickChangeClick(item)"
					v-for="(item, index) in props.quickBtn"
					:key="index"
					text
					outlined
					:shadow="0"
					:label="item.label"
				></tm-tag>
			</view>
		</view>
		<view class="flex flex-row flex-row-center-between mt-24 px-16 mb-32">
			<tm-text color="grey" label="自定时间范围"></tm-text>
			<tm-icon @click="clear" :font-size="24" color="grey" name="tmicon-delete"></tm-icon>
		</view>
		<view class="flex flex-row flex-row-center-center px-16">
			<tm-sheet
				text
				@click="changeFocus(0)"
				:round="24"
				class="flex-1"
				parenClass="flex-1"
				_class="flex flex-row flex-row-center-center"
				:color="focusIndex == 0 ? props.color : 'grey-4'"
				:height="76"
				outlined
				:border="4"
				:margin="[0, 0]"
				:padding="[32, 0]"
			>
				<view :userInteractionEnabled="false" style="width: 200rpx" class="flex flex-row flex-row-center-center">
					<tm-text v-if="!valuedate_start_str" color="grey" label="开始时间"></tm-text>
					<tm-text v-if="valuedate_start_str" :color="focusIndex == 0 ? props.color : 'grey'" :label="valuedate_start_str"></tm-text>
				</view>
			</tm-sheet>
			<view class="px-24">
				<tm-text :font-size="28" _class="text-weight-b" label="至"></tm-text>
			</view>
			<tm-sheet
				text
				@click="changeFocus(1)"
				:round="24"
				class="flex-1"
				parenClass="flex-1"
				_class="flex flex-row flex-row-center-center"
				:color="focusIndex == 1 ? props.color : 'grey-4'"
				:height="76"
				outlined
				:border="4"
				:margin="[0, 0]"
				:padding="[32, 0]"
			>
				<view :userInteractionEnabled="false" style="width: 200rpx" class="flex flex-row flex-row-center-center">
					<tm-text v-if="!valuedate_end_str" color="grey" label="结束时间"></tm-text>
					<tm-text v-if="valuedate_end_str" :color="focusIndex == 1 ? props.color : 'grey'" :label="valuedate_end_str"></tm-text>
				</view>
			</tm-sheet>
		</view>
		<view>
			<tm-time-view
				:showDetail="props.showDetail"
				:immediateChange="props.immediateChange"
				:start="starTtime"
				:end="endTtime"
				@change="timeChange"
				v-model="picker_focus_date"
				v-model:model-str="picker_focus_date_str"
				:defaultValue="picker_focus_date"
			></tm-time-view>
		</view>
		<view v-if="props.showBtn">
			<tm-button :color="props.color" @click="confirm" :round="24" :shadow="3" block label="确认"></tm-button>
		</view>
	</tm-sheet>
</template>

<script lang="ts" setup>
import { ref, computed, nextTick, PropType, watch } from 'vue'
import tmSheet from '@/tmui/components/tm-sheet/tm-sheet.vue'
import tmTag from '@/tmui/components/tm-tag/tm-tag.vue'
import tmText from '@/tmui/components/tm-text/tm-text.vue'
import tmIcon from '@/tmui/components/tm-icon/tm-icon.vue'
import tmButton from '@/tmui/components/tm-button/tm-button.vue'
import tmTimeView from '@/tmui/components/tm-time-view/tm-time-view.vue'
import * as dayjs from '../../tool/dayjs/esm/index'
const DayJs = dayjs.default
interface showDetail {
	year: boolean
	month: boolean
	day: boolean
	hour: boolean
	minute: boolean
	second: boolean
}
interface quickBtnType {
	label: string
	fun?: () => [string | number | Date | dayjs.Dayjs, string]
	color?: 'primary'
	//这个是内部代号，如果提供了，上面的fun可以不用提供，且无效
	// br:本日，bz:本周,by:本月,jsy:近三个月,bn:本年,d7：近7天，d30:近30天
	type?: 'br' | 'bz' | 'by' | 'jsy' | 'bn' | 'd7' | 'd30'
}
const props = defineProps({
	color: {
		type: String,
		default: 'primary'
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
	//展示格式。最终影响到modelStr输出格式的内容。
	format: {
		type: String,
		default: 'YYYY-MM-DD'
	},
	immediateChange: {
		type: Boolean,
		default: false
	},
	quickBtn: {
		type: Array as PropType<Array<quickBtnType>>,
		default: () => {
			return [
				{
					label: '本日',
					fun: null,
					color: 'primary',
					type: 'br'
				},
				{
					label: '本周',
					fun: null,
					color: 'primary',
					type: 'bz'
				},
				{
					label: '本月',
					fun: null,
					color: 'primary',
					type: 'by'
				},
				{
					label: '近三月',
					fun: null,
					color: 'primary',
					type: 'jsy'
				},
				{
					label: '本年',
					fun: null,
					color: 'primary',
					type: 'bn'
				},
				{
					label: '近7天',
					fun: null,
					color: 'primary',
					type: 'd7'
				},
				{
					label: '近30天',
					fun: null,
					color: 'primary',
					type: 'd30'
				}
			]
		}
	},
	start: {
		type: [String, Number, Date],
		default: new Date().setFullYear(new Date().getFullYear() - 1)
	},
	end: {
		type: [String, Number, Date],
		default: new Date().setFullYear(new Date().getFullYear() + 1)
	},
	defaultValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	/**
	 * 不受formart控制，始终完整显示，必须规范
	 * */
	modelValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	/**
	 * 只对外显示输出formart的格式字符串，可能不规范
	 * */
	modelStr: {
		type: Array as PropType<Array<String>>,
		default: () => []
	},
	/** 当改变日期时，是否需要同步到vmodel变量中，立即执行双向同步，如果为false日期的改变不会同步到vmodel中
	 * 然后需要你自行调用refs.getNowDate()来获取或者触发同步数据
	 * 这里设置的意义是：当你需要同步日期数据，由用户设置时，如果你设置了一个确认按钮来同步用户设置的日期，那么这里可以设置为fase，
	 * 然后你自行添加 一个确认按钮来触发refs.getNowDate()来同步数据，以执行用户的确认时间，否则不会同步。
	 */
	asyncModel: {
		type: Boolean,
		default: true
	},
	showBtn: {
		type: Boolean,
		default: true
	}
})
const emits = defineEmits(['change', 'update:modelValue', 'update:modelStr', 'confirm', 'clear'])
const valuedate_start = ref('')
const valuedate_end = ref('')
const focusIndex = ref(-1)
const picker_focus_date = ref('')
const picker_focus_date_str = ref('')
const defaultFormat = 'YYYY/MM/DD HH:mm:ss'
let tid: any = NaN
/** 开始时间 */
const starTtime = computed(() => DayJs(props.start).format(defaultFormat))
/** 结束时间 */
const endTtime = computed(() => DayJs(props.end).format(defaultFormat))

watch(
	() => props.modelValue,
	() => {
		clearTimeout(tid)
		tid = setTimeout(() => {
			setDate(props.modelValue)
		}, 100)
	},
	{ deep: true, immediate: true }
)
const valuedate_start_str = computed(() => {
	if (!valuedate_start.value) return ''
	return DayJs(valuedate_start.value).format(props.format)
})

const valuedate_end_str = computed(() => {
	if (!valuedate_end.value) return ''
	return DayJs(valuedate_end.value).format(props.format)
})
function changeFocus(index: number) {
	focusIndex.value = index
	let vstd = valuedate_start.value
	let vsed = valuedate_end.value

	if (index == 0 && !valuedate_start.value) {
		vstd = DayJs(valuedate_end.value).isValid() ? valuedate_end.value : DayJs().format(defaultFormat)
	}
	if (index == 1 && !valuedate_end.value) {
		vsed = DayJs(valuedate_start.value).isValid() ? valuedate_start.value : DayJs().format(defaultFormat)
	}

	setDate([vstd, vsed])
}

function timeChange(evt: string) {
	if (focusIndex.value == -1) focusIndex.value = 0

	if (focusIndex.value == 0) {
		valuedate_start.value = picker_focus_date.value
	} else if (focusIndex.value == 1) {
		valuedate_end.value = picker_focus_date.value
	}

	// #ifdef APP-NVUE
	updateTime()
	// #endif

	// #ifndef APP-NVUE
	nextTick(() => {
		updateTime()
	})
	// #endif
}
function updateTime() {
	checkTimeVaild()
	if (props.asyncModel) {
		emits('update:modelValue', [valuedate_start.value, valuedate_end.value])
		emits('update:modelStr', [DayJs(valuedate_start.value).format(props.format), DayJs(valuedate_end.value).format(props.format)])
	}
	nextTick(() => {
		emits('change', [valuedate_start.value, valuedate_end.value])
	})
}
//检查两个时间的正确性,就是后者要大于前者.
function checkTimeVaild() {
	if (valuedate_end.value.length == 0 || valuedate_start.value.length == 0) return
	if (focusIndex.value == 0) {
		if (DayJs(valuedate_start.value).isAfter(valuedate_end.value, 'second')) {
			valuedate_end.value = valuedate_start.value
		}
	}
	if (focusIndex.value == 1) {
		if (DayJs(valuedate_end.value).isBefore(valuedate_start.value, 'second')) {
			valuedate_start.value = valuedate_end.value
		}
	}
}

function quickChangeClick(item: quickBtnType) {
	let dstr = []
	let typesar = ['br', 'bz', 'by', 'jsy', 'bn', 'd7', 'd30']
	if (typeof item?.type === 'string' && typesar.includes(item?.type)) {
		let str = item?.type ?? ''
		if (str == 'br') {
			dstr = [DayJs().format(), DayJs().format()]
		} else if (str == 'bz') {
			let dt = DayJs()
			dstr = [dt.startOf('week').subtract(1, 'day').format(), dt.endOf('week').subtract(1, 'day').format()]
		} else if (str == 'by') {
			let dt = DayJs()
			dstr = [dt.startOf('month').format(), dt.endOf('month').format()]
		} else if (str == 'jsy') {
			let dt = DayJs()
			dstr = [dt.startOf('month').subtract(2, 'month').format(), dt.endOf('month').format()]
		} else if (str == 'bn') {
			let dt = DayJs()
			dstr = [dt.startOf('year').format(), dt.endOf('year').format()]
		} else if (str == 'd7') {
			let dt = DayJs()
			dstr = [dt.subtract(6, 'day').format(), dt.format()]
		} else if (str == 'd30') {
			let dt = DayJs()
			dstr = [dt.subtract(29, 'day').format(), dt.format()]
		}
	} else {
		if (typeof item?.fun == 'function') {
			dstr = item.fun()
			if (typeof idstr == 'function') {
				dstr = dstr()
			}
		}
	}
	if (Array.isArray(dstr)) {
		if (dstr.length !== 2) return

		valuedate_start.value = DayJs(dstr[0]).format(defaultFormat)
		valuedate_end.value = DayJs(dstr[1]).format(defaultFormat)
		if (focusIndex.value == -1) {
			focusIndex.value = 0
		}
		if (focusIndex.value == 0) {
			picker_focus_date.value = valuedate_start.value
		}
		if (focusIndex.value == 1) {
			picker_focus_date.value = valuedate_end.value
		}
	}
}

function setDate(ar: Array<String | Number | Date>) {
	if (Array.isArray(ar) && ar?.length == 2) {
		valuedate_start.value = DayJs(ar[0]).isValid() ? DayJs(ar[0]).format(defaultFormat) : DayJs(props.start).format(defaultFormat)
		valuedate_end.value = DayJs(ar[1]).isValid() ? DayJs(ar[1]).format(defaultFormat) : DayJs(props.end).format(defaultFormat)
		focusIndex.value = focusIndex.value == -1 ? 0 : focusIndex.value
		if (!DayJs(ar[focusIndex.value]).isValid()) return
		picker_focus_date.value = DayJs(ar[focusIndex.value]).format(defaultFormat)
	}
}
function confirm() {
	emits('update:modelValue', [valuedate_start.value, valuedate_end.value])
	emits('update:modelStr', [valuedate_start.value&&DayJs(valuedate_start.value).format(props.format), valuedate_end.value&&DayJs(valuedate_end.value).format(props.format)])
	emits('confirm', [valuedate_start.value, valuedate_end.value])
}
function getNowDate(): [string, string] {
	updateTime()
	return [valuedate_start.value, valuedate_end.value]
}
function clear() {
	valuedate_start.value = ''
	valuedate_end.value = ''
	emits('update:modelStr', ['', ''])
	emits('clear')
}
defineExpose({ getNowDate })
</script>

<style></style>
