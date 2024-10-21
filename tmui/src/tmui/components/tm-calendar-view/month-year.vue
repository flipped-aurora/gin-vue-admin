<template>
	<view class="flex flex-col relative">
		<tm-sheet v-if="!props.hideTool" :shadow="0" :round="0" :height="88" :margin="[0, 0]" :padding="[0, 0]" _class="flex flex-col">
			<view class="flex flex-row flex-row-center-center" style="height: 88rpx">
				<view @click.stop="prevYear" class="px-16">
					<tm-icon :userInteractionEnabled="false" :font-size="22" name="tmicon-angle-double-left"></tm-icon>
				</view>
				<view class="px-12">
					<tm-text :userInteractionEnabled="false" _class="text-weight-b" :font-size="32" :label="_nowDate"></tm-text>
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
				<tm-text :userInteractionEnabled="false" color="grey" _class="text-align-center" :font-size="28" :label="props.textUnit[10]"></tm-text>
			</view>
		</tm-sheet>

		<view class="flex flex-col">
			<view v-for="(item2, index2) in _data" :key="index2">
				<view class="flex flex-row flex-row-center-center" :style="[{ height: '120rpx', flexWrap: 'wrap' }]">
					<view class="flex-4" :style="{ height: '112rpx' }" v-for="(item, index) in item2" :key="index">
						<tm-sheet
							@click="clickWeek(item)"
							:height="112"
							:shadow="0"
							:round="4"
							_class="flex-row flex-center"
							:text="_nowMonth == item.month"
							:color="_nowMonth == item.month ? _color : 'grey-4'"
							:margin="[4, 4]"
							:padding="[0, 0]"
						>
							<view
								:style="[{ opacity: item.isVaild ? '0.3' : '1' }]"
								:userInteractionEnabled="false"
								style="width: 62rpx"
								class="flex-center flex-col"
							>
								<tm-text :font-size="28" :label="item.month + props.textUnit[13]"></tm-text>
							</view>
						</tm-sheet>
					</view>
				</view>
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
 * 按月选择的日历
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
import isBetween from '../../tool/dayjs/esm/plugin/isBetween/index'
import { monthYearItem } from './interface'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
/**
 * 事件说明
 * v-model 绑定当前的周次时间开始和结束时间。
 * confirm 当点击确认时触发
 * click-month 点击时期上的月份时触发。
 * change 当切换年或者月的时候触发。
 */
const emits = defineEmits(['update:modelValue', 'confirm', 'click-month', 'change'])
const props = defineProps({
	followTheme: {
		type: Boolean,
		default: true
	},
	//默认显示的日期，自动转换为当前的周次。
	defaultValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	//当前时间
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
//当前的时间
const _value = ref(DayJs(props.defaultValue[0]).isValid() ? DayJs(props.defaultValue[0]) : DayJs())
const _start_date = computed(() => {
	let isv = DayJs(props.start).isValid()
	return isv ? DayJs(props.start) : DayJs('1980-1-1')
})
const _end_date = computed(() => {
	let isv = DayJs(props.end).isValid()
	return isv ? DayJs(props.end) : DayJs('2450-1-1')
})
const _data: Ref<Array<Array<monthYearItem>>> = ref(getDataArray())
//当前的年月
const _nowDate = computed(() => {
	return _value.value.format('YYYY-MM')
})
//当前的月
const _nowMonth = computed(() => {
	return _value.value.format('M')
})

watch(
	[() => props.modelValue, () => props.start, () => props.end],
	() => {
		_value.value = DayJs(props.modelValue[0])
		_data.value = getDataArray()
	},
	{ deep: true }
)

function nowWeekClick() {
	if (DayJs().isBetween(_start_date.value, _end_date.value, 'month', '[]') == false) {
		uni.showToast({ title: '无法选中', icon: 'none' })
		return
	}
	_value.value = DayJs()
	_data.value = getDataArray()
	updateTimes()
}
function clickWeek(wk: monthYearItem) {
	if (wk.isVaild) {
		uni.showToast({ title: '无法选中', icon: 'none' })
		return
	}
	_value.value = _value.value.month(wk.month - 1)
	updateTimes()
}

//设置当前选中的日期
function setDefault(data: Array<String | Number | Date> = []) {
	_value.value = data ? DayJs(data[0]) : DayJs(props.modelValue[0])
	_data.value = getDataArray()
}
function getDataArray() {
	let nowMonth: dayjs.Dayjs = DayJs('2000-1-1').year(_value.value.year())
	let ar: Array<monthYearItem> = []
	for (let i = 0; i < 12; i++) {
		nowMonth = nowMonth.month(i)
		ar.push({
			dateStr: nowMonth.format('YYYY-MM'),
			month: nowMonth.month() + 1,
			isVaild: !nowMonth.isBetween(_start_date.value, _end_date.value, 'month', '[]')
		})
	}
	return uni.$tm.u.splitData(ar, 3)
}

function nextYear() {
	_value.value = _value.value.add(1, 'year')
	_data.value = getDataArray()
	emits('change', _value.value.format('YYYY/MM/DD'))
}
function prevYear() {
	_value.value = _value.value.subtract(1, 'year')
	_data.value = getDataArray()
	emits('change', _value.value.format('YYYY/MM/DD'))
}
//update
function updateTimes() {
	emits('click-month', _value.value.format('YYYY-MM'))
}
function confirm() {
	emits('update:modelValue', [_value.value.format('YYYY-MM')])
	emits('confirm', [_value.value.format('YYYY-MM')])
}
defineExpose({
	setDefault: setDefault,
	nextYear: nextYear,
	prevYear: prevYear
})
</script>

<style></style>
