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
				<tm-text :userInteractionEnabled="false" color="grey" _class="text-align-center" :font-size="28" :label="props.textUnit[9]"></tm-text>
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
			<view class="flex flex-row flex-row-center-center" :style="[{ height: '74rpx' }]" v-for="(item, index) in _dataWeek" :key="index">
				<view :class="['opacity-5']" class="flex-1 flex-center">
					<view style="width: 62rpx" class="flex-center flex-col">
						<tm-text :font-size="24" :label="item"></tm-text>
					</view>
				</view>
				<tm-sheet
					no-level
					@click="clickWeek(item)"
					:height="66"
					:shadow="0"
					:round="10"
					_class="flex-row"
					class="flex-6"
					paren-class="flex-6"
					:text="_weekNum == item"
					:color="_weekNum == item ? _color : 'grey-4'"
					:margin="[0, 4]"
					:padding="[0, 0]"
				>
					<view
						:userInteractionEnabled="false"
						style="width: 62rpx"
						:class="[item2.isNowIn ? '' : 'opacity-5']"
						class="flex-1 flex-center"
						v-for="(item2, index2) in _data[index]"
						:key="index2"
					>
						<view style="width: 62rpx" class="flex-center flex-col" :style="[{ opacity: item2.isVaild ? '0.3' : '1' }]">
							<tm-text :font-size="28" :label="item2.date"></tm-text>
						</view>
					</view>
				</tm-sheet>
			</view>
		</view>

		<tm-button
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
 * 按周选择的日历
 * @description 可以按月，按日，按周，按季度显示
 */
import { computed, ref, Ref, watch, PropType } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import tmButton from '../tm-button/tm-button.vue'
import * as dayjs from '../../tool/dayjs/esm/index'
import isoWeek from '../../tool/dayjs/esm/plugin/isoWeek/index'
import isSameOrBefore from '../../tool/dayjs/esm/plugin/isSameOrBefore/index'
import isSameOrAfter from '../../tool/dayjs/esm/plugin/isSameOrAfter/index'
import isBetween from '../../tool/dayjs/esm/plugin/isBetween/index'
import { weekItem } from './interface'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
/**
 * 事件说明
 * v-model 绑定当前的周次时间开始和结束时间。
 * confirm 当点击确认时触发
 * click-week 点击时期上的周次时触发。
 * change 当切换年或者月的时候触发。
 */
const emits = defineEmits(['update:modelValue', 'confirm', 'click-week', 'change'])
const props = defineProps({
	followTheme: {
		type: Boolean,
		default: true
	},
	//默认显示的日期，自动转换为当前的周次。只会取第一个时间判断当前周次。
	defaultValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	//当前的周时间段，本年度周一至周日的时间范围段。
	modelValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	color: {
		type: String,
		default: 'primary'
	},
	linear: {
		type: String as PropType<'left' | 'right' | 'bottom' | 'top' | ''>,
		default: ''
	},
	linearDeep: {
		type: String as PropType<'accent' | 'dark' | 'light'>,
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
DayJs.extend(isSameOrAfter)
DayJs.extend(isBetween)

//当前的时间
const _value = ref(DayJs(props.defaultValue[0]).isValid() ? DayJs(props.defaultValue[0]) : DayJs())
//当前的周次。
const _weekNum = ref(getNowWeek(_value.value))
const weekStr = props.textUnit.slice(0,8)
const _data: Ref<Array<Array<weekItem>>> = ref([])
const _dataWeek: Ref<Array<number>> = ref([]) //本月的周次。
const _start_date = computed(() => {
	let isv = DayJs(props.start).isValid()
	return isv ? DayJs(props.start) : DayJs('1980-1-1')
})
const _end_date = computed(() => {
	let isv = DayJs(props.end).isValid()
	return isv ? DayJs(props.end) : DayJs('2450-1-1')
})

//当前的年月日期。
const _nowDate = computed(() => {
	return _value.value.format('YYYY-MM')
})
_data.value = getWeekOfMonthArray()

watch(
	() => props.modelValue,
	() => {
		if (!Array.isArray(props.modelValue)) return
		let date_str = props.modelValue[0] ?? ''
		_value.value = DayJs(date_str).isValid() ? DayJs(date_str) : DayJs()
		_weekNum.value = getNowWeek(_value.value)
		_data.value = getWeekOfMonthArray()
	},
	{ deep: true }
)
//获取当前周次。根据当前日期。
function getNowWeek(str: string | dayjs.Dayjs | number | Date = '') {
	if (DayJs(str).isValid()) {
		return DayJs(str).isoWeek()
	} else {
		return DayJs().isoWeek()
	}
}
function nowWeekClick() {
	_value.value = DayJs()
	_data.value = getWeekOfMonthArray()
	let wk = getNowWeek(_value.value)
	if (!canSelected(wk)) {
		uni.showToast({ title: '无法选中', icon: 'none' })
		return
	}
	_weekNum.value = wk

	updateTimes()
}
function clickWeek(wk: number) {
	if (!canSelected(wk)) {
		uni.showToast({ title: '无法选中', icon: 'none' })
		return
	}
	_weekNum.value = wk
	updateTimes()
}
function canSelected(num: number) {
	// 能否选中周次。
	let index = _dataWeek.value.findIndex((el) => el == num)
	let item = _data.value[index]
	let ar = item.filter((el) => !el.isVaild)

	return ar.length > 0
}
function nextYear() {
	_value.value = _value.value.add(1, 'year')
	_data.value = getWeekOfMonthArray()
	emits('change', _value.value.format('YYYY/MM/DD'))
}
function nextMonth() {
	_value.value = _value.value.add(1, 'month')
	_data.value = getWeekOfMonthArray()
	emits('change', _value.value.format('YYYY/MM/DD'))
}
function prevMonth() {
	_value.value = _value.value.subtract(1, 'month')
	_data.value = getWeekOfMonthArray()
	emits('change', _value.value.format('YYYY/MM/DD'))
}
function prevYear() {
	_value.value = _value.value.subtract(1, 'year')
	_data.value = getWeekOfMonthArray()
	emits('change', _value.value.format('YYYY/MM/DD'))
}

//设置当前选中的日期范围，参数为周次中的某一个日期即可，会自动转换成以该日期为周的范围段。
function setDefault(data: Array<String | Number | Date> = []) {
	let date_str = data[0]
	_value.value = DayJs(date_str).isValid() ? DayJs(date_str) : DayJs()
	_weekNum.value = getNowWeek(_value.value)
	_data.value = getWeekOfMonthArray()
}

//获取本月按周分组的时间段。
function getWeekOfMonthArray() {
	let nowMonth: dayjs.Dayjs = DayJs(_value.value)
	let startStatickDay = nowMonth.startOf('month').format('YYYY/MM/DD')
	let endStatickDay = nowMonth.endOf('month').format('YYYY/MM/DD')
	let startd = DayJs(startStatickDay)

	let arOfmonth: Array<number> = [] //本月的周次。
	let ar: Array<weekItem> = []
	while (startd.isSameOrBefore(endStatickDay)) {
		ar.push({
			dateStr: startd.format('YYYY/MM/DD'),
			date: startd.date() < 10 ? '0' + startd.date() : startd.date(),
			week: startd.isoWeek(),
			day: startd.isoWeekday(),
			isNowIn: isInNowMonth(nowMonth, startd),
			isVaild: !startd.isBetween(_start_date.value, _end_date.value, 'day', '[]')
		})
		arOfmonth.push(startd.isoWeek())
		startd = startd.add(1, 'day')
	}
	arOfmonth = [...new Set(arOfmonth)]
	_dataWeek.value = arOfmonth
	let dArray: Array<Array<weekItem>> = []
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
	//补齐前尾的7天时间。
	if (dArray[0].length !== 7) {
		let item = dArray[0][dArray[0].length - 1]
		let start_of = DayJs(item.dateStr).isoWeek(item.week).subtract(6, 'day')
		let end_of = DayJs(item.dateStr).isoWeek(item.week)
		let pr = []
		let startd = DayJs(start_of)
		while (startd.isSameOrBefore(end_of)) {
			pr.push({
				dateStr: startd.format('YYYY/MM/DD'),
				date: startd.date() < 10 ? '0' + startd.date() : startd.date(),
				week: startd.isoWeek(),
				day: startd.isoWeekday(),
				isNowIn: isInNowMonth(nowMonth, startd),
				isVaild: !startd.isBetween(_start_date.value, _end_date.value, 'day', '[]')
			})
			startd = startd.add(1, 'day')
		}
		dArray[0] = pr
	}
	if (dArray[dArray.length - 1].length !== 7) {
		let item = dArray[dArray.length - 1][0]

		let start_of = DayJs(item.dateStr).isoWeek(item.week)
		let end_of = DayJs(item.dateStr).isoWeek(item.week).add(6, 'day')
		let pr = []
		let startd = DayJs(start_of)
		while (startd.isSameOrBefore(end_of)) {
			pr.push({
				dateStr: startd.format('YYYY/MM/DD'),
				date: startd.date() < 10 ? '0' + startd.date() : startd.date(),
				week: startd.isoWeek(),
				day: startd.isoWeekday(),
				isNowIn: isInNowMonth(nowMonth, startd),
				isVaild: !startd.isBetween(_start_date.value, _end_date.value, 'day', '[]')
			})
			startd = startd.add(1, 'day')
		}
		dArray[dArray.length - 1] = pr
	}
	return dArray
}
// 检测now日期是否在某个日期的月份中.
function isInNowMonth(date: string | dayjs.Dayjs | number = '', now: string | dayjs.Dayjs | number = '') {
	let startStatickDay = DayJs(date).startOf('month').format('YYYY/MM/DD')
	let endStatickDay = DayJs(date).endOf('month').format('YYYY/MM/DD')
	return DayJs(now).isBetween(startStatickDay, endStatickDay, 'day', '[]')
}

//update
function updateTimes() {
	let index = _dataWeek.value.findIndex((el) => el == _weekNum.value)
	let item = _data.value[index]
	let start = item[0].dateStr
	let end = item[item.length - 1].dateStr
	emits('click-week', [start, end])
}
function confirm() {
	let index = _dataWeek.value.findIndex((el) => el == _weekNum.value)
	let item = [..._data.value[index]]
	let start = item[0].dateStr
	// 最开始的日期如果在start后就可以直接读取
	for (let i = 0; i < item.length; i++) {
		if (DayJs(item[i].dateStr).isSameOrAfter(props.start, 'date')) {
			start = item[i].dateStr
			break
		}
	}
	let end = item[item.length - 1].dateStr
	item = item.reverse()
	console.log(item)
	for (let i = 0; i < item.length; i++) {
		if (DayJs(item[i].dateStr).isSameOrBefore(props.end, 'date')) {
			end = item[i].dateStr
			break
		}
	}

	emits('update:modelValue', [start, end])
	emits('confirm', [start, end])
}
defineExpose({
	setDefault: setDefault,
	nextYear: nextYear,
	nextMonth: nextMonth,
	prevYear: prevYear,
	prevMonth: prevMonth
})
</script>

<style></style>
