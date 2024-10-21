<template>
	<view class="flex flex-col relative">
		<tm-sheet v-if="!props.hideTool" :shadow="0" :round="0" :height="88" :margin="[0, 0]" :padding="[0, 0]" _class="flex flex-col">
			<view class="flex flex-row flex-row-center-center" style="height: 88rpx">
				<view @click.stop="prevYear" class="px-16">
					<tm-icon :userInteractionEnabled="false" :font-size="22" name="tmicon-angle-double-left"></tm-icon>
				</view>
				<view @click.stop="prevMonth" class="px-16">
					<tm-icon :userInteractionEnabled="false" :font-size="22" name="tmicon-angle-left"></tm-icon>
				</view>
				<view class="px-12">
					<tm-text :userInteractionEnabled="false" _class="text-weight-b" :font-size="32" :label="_nowDate"></tm-text>
				</view>
				<view @click.stop="nextMonth" class="px-16">
					<tm-icon :userInteractionEnabled="false" :font-size="22" name="tmicon-angle-right"></tm-icon>
				</view>
				<view @click.stop="nextYear" class="px-16">
					<tm-icon :userInteractionEnabled="false" :font-size="22" name="tmicon-angle-double-right"></tm-icon>
				</view>
			</view>
			<view
				@click="nowWeekClick"
				class="absolute t-0 r--6 zIndex-10 round-12 py-4 flex flex-row flex-row-center-center"
				style="width: 90rpx; height: 88rpx"
			>
				<tm-text :userInteractionEnabled="false" color="grey" _class="text-align-center" :font-size="28" :label="props.textUnit[8]"></tm-text>
			</view>
		</tm-sheet>
		<view class="flex flex-row flex-row-center-center py-12" :style="[{ height: '74rpx' }]">
			<view class="flex-1 flex-center" v-for="(item, index) in weekStr" :key="index">
				<view style="width: 62rpx" class="flex-center flex-col">
					<tm-text :font-size="24" :label="item"></tm-text>
				</view>
			</view>
		</view>
		<view class="flex flex-col">
			<view class="flex flex-row flex-row-center-center" :style="[{ height: '94rpx' }]" v-for="(item, index) in _data" :key="index">
				<view @click="clickWeek(item2)" class="flex-1 flex flex-row flex-row-center-center" v-for="(item2, index2) in item" :key="index2">
					<tm-sheet
						:userInteractionEnabled="false"
						:height="90"
						:width="90"
						:shadow="0"
						:round="24"
						:border="item2.extra.color && isSelected(item2.dateStr) ? 1 : 0"
						_class="flex-row"
						:transprent="item2.extra.color || !isSelected(item2.dateStr)"
						:color="item2.extra.color ? item2.extra.color : isSelected(item2.dateStr) ? _color : 'white'"
						:margin="[0, 0]"
						:padding="[0, 0]"
					>
						<view
							:userInteractionEnabled="false"
							style="width: 62rpx"
							:class="[!item2.isNowIn ? 'opacity-6' : '']"
							class="flex-1 flex-center"
						>
							<view style="width: 62rpx" class="flex-center flex-col" :style="[{ opacity: item2.disabled ? '0.3' : '1' }]">
								<tm-text :font-size="28" :label="item2.date"></tm-text>
								<tm-text _class="flex-center" v-if="item2.extra.extra" :font-size="22" :label="item2.extra.extra"></tm-text>
							</view>
						</view>
					</tm-sheet>
				</view>
			</view>
		</view>

		<tm-button
			:followTheme="props.followTheme"
			v-if="!props.hideButton"
			:linear="props.linear"
			:linear-deep="props.linearDeep"
			:color="props.color"
			@click="confirm"
			block
			:label="_confirmText"
			:margin="[0, 16]"
		></tm-button>
	</view>
</template>

<script lang="ts" setup>
/**
 * 按日选择的日历
 * @description 可以多选，单选
 */
import { computed, nextTick, ref, Ref, watch, PropType } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmButton from '../tm-button/tm-button.vue'
import * as dayjs from '../../tool/dayjs/esm/index'
import isoWeek from '../../tool/dayjs/esm/plugin/isoWeek/index'
import isSameOrBefore from '../../tool/dayjs/esm/plugin/isSameOrBefore/index'
import isBetween from '../../tool/dayjs/esm/plugin/isBetween/index'
import { monthDayItem, dateItemStyle } from './interface'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
/**
 * 事件说明
 * v-model 绑定当前的周次时间开始和结束时间。
 * confirm 当点击确认时触发
 * click-day 日期被选中时触发，注意禁用的日期不会触发 。
 * change 当切换年或者月的时候触发。
 */
const emits = defineEmits(['update:modelValue', 'confirm', 'click-day', 'change'])
const props = defineProps({
	followTheme: {
		type: Boolean,
		default: true
	},
	//默认显示的日期
	defaultValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	//当前的周时间段
	modelValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	color: {
		type: String,
		default: 'primary'
	},
	linear: {
		type: String,
		default: ''
	},
	linearDeep: {
		type: String,
		default: 'light'
	},
	//指的是：有效的可选时间，小于此时间，不允许选中。
	start: {
		type: [String, Number, Date],
		default: ''
	},
	//指的是：有效的可选时间，大于此时间，不允许选中。
	end: {
		type: [String, Number, Date],
		default: ''
	},
	//被禁用的日期数组。如果["2022-1-1","2022-5-1"]
	//被禁用的日期将无法选中。
	disabledDate: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	//是否允许多选。
	multiple: {
		type: Boolean,
		default: false
	},
	//设定单个日期的样式。
	dateStyle: {
		type: Array as PropType<Array<dateItemStyle>>,
		default: () => []
	},
	//当multiple=true时，可以选择的最大日期数量。
	max: {
		type: Number,
		default: 999
	},
	//隐藏头部操作栏
	hideTool: {
		type: Boolean,
		default: false
	},
	//隐藏底部确认按钮
	hideButton: {
		type: Boolean,
		default: false
	},
	confirmText: {
		type: String,
		default: '确认'
	},
	//周次，本日、本季、本年、本月、本周的文字请按顺序提供文本，方便定义其它语言。
	textUnit: {
		type: Array as PropType<string[]>,
		default: ['周次','一','二','三','四','五','六','日','本日','本周','本月','本季度','本年','月','第${x}季度','年']
	}
})
const _color = computed(() => {
	if (props.followTheme && store.tmStore.color) return store.tmStore.color
	return props.color
})
const _confirmText = computed(() => props.confirmText)

const DayJs = dayjs.default
DayJs.extend(isoWeek)
DayJs.extend(isSameOrBefore)
DayJs.extend(isBetween)
//当前选中的时间数组。
const _value = ref(props.defaultValue)
const weekStr = props.textUnit.slice(1,8)

//当前需要展示的年月。
const showOpenDate: Ref<dayjs.Dayjs> = ref(setShowopenDate())
//当前月历中的日期。
const _data: Ref<Array<Array<monthDayItem>>> = ref([])
//可选中开始时间
const _start_date = computed(() => {
	let isv = DayJs(props.start).isValid()
	return isv ? DayJs(props.start) : DayJs('1980-1-1')
})
//可选中的结束时间
const _end_date = computed(() => {
	let isv = DayJs(props.end).isValid()
	return isv ? DayJs(props.end) : DayJs('2450-1-1')
})

//当前展示的年月日期。
const _nowDate = computed(() => {
	return showOpenDate.value.format('YYYY-MM')
})
_data.value = getWeekOfMonthArray()

watch(
	[() => props.modelValue, () => props.dateStyle, () => props.disabledDate, () => props.start, () => props.end],
	() => {
		_value.value = props.modelValue
		showOpenDate.value = setShowopenDate()
		_data.value = getWeekOfMonthArray()
	},
	{ deep: true }
)

//设置当前需要展示的月份。
function setShowopenDate() {
	//从当前选中的第一个日期开始。如果没有就显示当前本地时间的月。

	if (_value.value.length == 0) {
		return DayJs()
	}
	let n = _value.value[0] || DayJs()
	n = typeof n == 'undefined' || n == null ? DayJs() : n
	return DayJs(n)
}

function nowWeekClick() {
	if (isDisabledDate(DayJs())) {
		uni.showToast({ title: '无法选中', icon: 'none' })
		return
	}
	selected(DayJs().format('YYYY/MM/DD'))
	showOpenDate.value = DayJs()
	_data.value = getWeekOfMonthArray()
	emits('click-day', DayJs().format('YYYY/MM/DD'))
}
function clickWeek(wk: monthDayItem) {
	if (wk.disabled) {
		uni.showToast({ title: '无法选中', icon: 'none' })
		return
	}
	selected(wk.dateStr)
	emits('click-day', wk.dateStr)
}
//选中日期。相同需要减去。重复需要去除。
function selected(item: string | dayjs.Dayjs | number) {
	let fr = _value.value.filter((el) => DayJs(el).isSame(item))
	if (!props.multiple) {
		_value.value = [DayJs(item).format('YYYY/MM/DD')]
		return
	}

	if (fr.length > 0) {
		_value.value = _value.value.filter((el) => !DayJs(el).isSame(item))
	} else {
		if (_value.value.length >= props.max) {
			uni.showToast({ title: '只可选择' + props.max + '天', icon: 'none' })
			return
		}
		_value.value.push(DayJs(item).format('YYYY/MM/DD'))
	}
}
function nextYear() {
	showOpenDate.value = showOpenDate.value.add(1, 'year')
	let dys = getWeekOfMonthArray()
	emits('change', showOpenDate.value.format('YYYY/MM/DD'))
	nextTick(() => {
		_data.value = [...dys]
	})
}
function nextMonth() {
	showOpenDate.value = showOpenDate.value.add(1, 'month')
	let dys = getWeekOfMonthArray()
	emits('change', showOpenDate.value.format('YYYY/MM/DD'))
	nextTick(() => {
		_data.value = [...dys]
	})
}
function prevMonth() {
	showOpenDate.value = showOpenDate.value.subtract(1, 'month')
	let dys = getWeekOfMonthArray()
	emits('change', showOpenDate.value.format('YYYY/MM/DD'))
	nextTick(() => {
		_data.value = [...dys]
	})
}
function prevYear() {
	showOpenDate.value = showOpenDate.value.subtract(1, 'year')
	let dys = getWeekOfMonthArray()
	emits('change', showOpenDate.value.format('YYYY/MM/DD'))
	nextTick(() => {
		_data.value = [...dys]
	})
}
//设置当前选中的日期范围，参数必须为有效的时间数组。
function setDefault(data: Array<String | Number | Date> = []) {
	_value.value = data.length > 0 ? data : props.modelValue
	showOpenDate.value = setShowopenDate()
	_data.value = getWeekOfMonthArray()
}
//获取本月按周分组的时间段。
function getWeekOfMonthArray() {
	let nowMonth: dayjs.Dayjs = showOpenDate.value.date(1)
	let startStatickDay = nowMonth.startOf('month')
	let endStatickDay = nowMonth.endOf('month')

	///当前月份有多少天。
	let nowMonthDayNum = nowMonth.daysInMonth()
	//第一个日期需要补齐1-7天的。
	let startOfday = startStatickDay.isoWeekday() - 1

	startStatickDay = nowMonth.subtract(Math.abs(startOfday), 'day')
	//最后一个日期也要补齐1-7天。
	let endOfday = 7 - endStatickDay.isoWeekday()
	if (endOfday > 0) {
		endStatickDay = nowMonth.date(nowMonthDayNum).add(Math.abs(endOfday), 'day')
	}

	let startd = DayJs(startStatickDay)
	let arOfmonth: Array<number> = [] //本月的周次。
	let ar: Array<monthDayItem> = []
	function setAr() {
		let dy = props.dateStyle.map((el) => {
			el.date = DayJs(el.date).format('YYYY/MM/DD')
			return el
		})
		let dyObj = {}
		dy.forEach((el) => {
			dyObj[el.date] = el
		})
		let dySet = Object.keys(dyObj)
		while (startd.isSameOrBefore(endStatickDay)) {
			let idate = startd.format('YYYY/MM/DD')
			let ext = dySet.includes(idate) ? dyObj[idate] : {}
			ar.push({
				dateStr: idate,
				date: startd.date() < 10 ? '0' + startd.date() : startd.date(),
				day: startd.isoWeekday(),
				week: startd.isoWeek(),
				isNowIn: isInNowMonth(nowMonth, startd),
				disabled: isDisabledDate(startd),
				extra: {
					date: idate,
					text: false,
					color: '',
					extra: '',
					...ext
				}
			})
			arOfmonth.push(startd.isoWeek())
			startd = startd.add(1, 'day')
		}
	}
	setAr()
	// 再检查当前展示时段内有没有满足6*7 42天没有也要补上剩余的天数。
	if (ar.length < 42) {
		let chaJi = 42 - ar.length
		endStatickDay = endStatickDay.add(chaJi, 'day')
		setAr()
	}

	arOfmonth = [...new Set(arOfmonth)]
	let dArray: Array<Array<monthDayItem>> = []
	let index = 0
	dArray.push([])
	ar.forEach((el) => {
		if (el.week == arOfmonth[index]) {
			dArray[index].push(el)
		} else {
			index += 1
			dArray.push([])
			dArray[index].push(el)
		}
	})
	return dArray
}
// 检测now日期是否在某个日期的月份中.
function isInNowMonth(date: string | dayjs.Dayjs | number = '', now: string | dayjs.Dayjs | number = '') {
	let startStatickDay = DayJs(date).startOf('month').format('YYYY/MM/DD')
	let endStatickDay = DayJs(date).endOf('month').format('YYYY/MM/DD')
	return DayJs(now).isBetween(startStatickDay, endStatickDay, 'day', '[]')
}
//检测某个日期，是否在禁用中。
function isDisabledDate(date: string | dayjs.Dayjs | number = '') {
	let valdate = DayJs(date)
	let isds = false
	isds = !valdate.isBetween(_start_date.value, _end_date.value, 'day', '[]')
	for (let i = 0; i < props.disabledDate.length; i++) {
		let item = props.disabledDate[i]
		if (DayJs(item).isSame(valdate)) {
			isds = true
			break
		}
	}
	return isds
}
//检测 某个日期是否被选中。
function isSelected(date: string | dayjs.Dayjs | number = '') {
	let fr = _value.value.filter((el) => DayJs(el).isSame(date))

	return fr.length > 0
}

function confirm() {
	let ar = _value.value.map((el) => DayJs(el).format('YYYY/MM/DD'))
	emits('update:modelValue', ar)
	emits('confirm', ar)
}

//对外方法。
defineExpose({
	setDefault: setDefault,
	nextYear: nextYear,
	nextMonth: nextMonth,
	prevYear: prevYear,
	prevMonth: prevMonth
})
</script>

<style></style>
