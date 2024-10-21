<template>
	<tm-sheet
		:follow-dark="props.followDark"
		:follow-theme="props.followTheme"
		:dark="props.dark"
		:shadow="props.shadow"
		:round="props.round"
		:margin="props.margin"
		:padding="props.padding"
		:color="props.color"
	>
		<view class="flex flex-row flex-row-center-center">
			<view v-if="showArrow" @click.stop="prevWeek" class="opacity-7">
				<tm-icon :font-size="28" :userInteractionEnabled="false" name="tmicon-angle-left"></tm-icon>
			</view>
			<view class="flex-1 flex flex-row" style="width: 0px">
				<view
					@click.stop="changeDate(item.date)"
					:class="[item.date == _value ? '' : 'opacity-7']"
					class="flex-1 flex flex-col flex-col-center-center py-24"
					v-for="(item, index) in nowWeekDayArray"
					:key="index"
					:style="{backgroundColor:item.date == _value?props.itemSelectedBgColor:'rgba(255,255,255,0)'}"
				>
					<tm-text
						:followTheme="false"
						:color="item.date == _value ? props.activeColor : ''"
						:userInteractionEnabled="false"
						:font-size="23"
						:label="weekStr[index]"
					></tm-text>
					<tm-text
						:followTheme="false"
						:color="item.date == _value ? props.activeColor : ''"
						:userInteractionEnabled="false"
						:font-size="23"
						:label="item.str"
					></tm-text>
				</view>
			</view>
			<view v-if="showArrow" @click.stop="nexWeek" class="opacity-7">
				<tm-icon :font-size="28" :userInteractionEnabled="false" name="tmicon-angle-right"></tm-icon>
			</view>
		</view>
	</tm-sheet>
</template>
<script lang="ts" setup>
/**
 * 时间周
 * @description 按周显示可选的日期。
 */
import { computed, PropType, Ref, ref, watch } from 'vue'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import * as dayjs from '../../tool/dayjs/esm/index'
import isoWeek from '../../tool/dayjs/esm/plugin/isoWeek/index'
import isSameOrBefore from '../../tool/dayjs/esm/plugin/isSameOrBefore/index'
/**
 * 事件说明
 * change 点击日期时触发,
 * v-model 双向绑定当前选中的日期。
 */
const emits = defineEmits(['change', 'update:modelValue'])
const props = defineProps({
	...custom_props,
	padding: {
		type: Array as PropType<Array<number>>,
		default: [12, 0]
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: [32, 32]
	},
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	transprent: {
		type: Boolean,
		default: false
	},
	round: {
		type: Number,
		default: 3
	},
	itemSelectedBgColor:{
		type:String,
		default:""
	},
	shadow: {
		type: Number,
		default: 3
	},
	defaultValue: {
		type: [String, Date, Number],
		default: () => ''
	},
	modelValue: {
		type: [String, Date, Number],
		default: () => ''
	},
	color: {
		type: String,
		default: 'white'
	},
	activeColor: {
		type: String,
		default: 'primary'
	},
	model: {
		type: String as PropType<'week' | 'custom'>,
		default: 'week'
	},
	//只有model="custom"时，才起作用，自定显示的天数
	dayNumber: {
		type: Number,
		default: 4
	},
	showArrow: {
		type: Boolean,
		default: true
	}
})
const DayJs = dayjs.default
DayJs.extend(isoWeek)
DayJs.extend(isSameOrBefore)
const _value = ref(DayJs(props.defaultValue).isValid() ? DayJs(props.defaultValue).format('YYYY-MM-DD') : DayJs().format('YYYY-MM-DD'))
const nowWeek: Ref<Array<string | Date | number>> = ref(getweek_s_e('now', _value.value))
const nowWeekDayArray = computed(() => getAllDay(nowWeek.value[0], nowWeek.value[1]))
const weekStr = ref(['周一', '周二', '周三', '周四', '周五', '周六', '周日'])
const weekStrIndex = {
	0: '周日',
	1: '周一',
	2: '周二',
	3: '周三',
	4: '周四',
	5: '周五',
	6: '周六'
}

watch(
	() => props.modelValue,
	() => {
		if (DayJs(props.modelValue).isValid()) {
			_value.value = DayJs(props.modelValue).format('YYYY-MM-DD')
			nowWeek.value = getweek_s_e('now', _value.value)
		}
	}
)
function changeDate(date: string) {
	_value.value = date
	emits('update:modelValue', date)
	emits('change', date)
}
function nexWeek() {
	nowWeek.value = getweek_s_e('next')
}
function prevWeek() {
	nowWeek.value = getweek_s_e('prev')
}

//获取周始和结束的时间值。
function getweek_s_e(type = 'next', daytime: string | Date | number = ''): Array<string | Date | number> {
	let nowTimeDay = DayJs(_value.value)
	if (type == 'next') {
		if (props.model == 'custom') {
			// 自定义日期
			let s = nowTimeDay.add(1, 'day')
			_value.value = s.format('YYYY-MM-DD')
			return [s.format('YYYY-MM-DD'), s.add(props.dayNumber - 1, 'day').format('YYYY-MM-DD')]
		} else {
			let date = DayJs(nowWeek.value[1]).add(1, 'week')
			return [date.startOf('isoWeek').format('YYYY-MM-DD'), date.format('YYYY-MM-DD')]
		}
	} else if (type == 'prev') {
		if (props.model == 'custom') {
			// 自定义日期
			let s = nowTimeDay.subtract(1, 'day')
			_value.value = s.format('YYYY-MM-DD')
			return [s.format('YYYY-MM-DD'), s.add(props.dayNumber - 1, 'day').format('YYYY-MM-DD')]
		} else {
			let date = DayJs(nowWeek.value[0]).subtract(1, 'week')
			return [date.format('YYYY-MM-DD'), date.endOf('isoWeek').format('YYYY-MM-DD')]
		}
	} else {
		if (props.model == 'custom') {
			// 自定义日期
			if (!DayJs(daytime).isValid()) {
				return [
					DayJs().subtract(1, 'day').format('YYYY-MM-DD'),
					DayJs()
						.add(props.dayNumber - 2, 'day')
						.format('YYYY-MM-DD')
				]
			} else {
				return [
					DayJs(daytime).subtract(2, 'day').format('YYYY-MM-DD'),
					DayJs(daytime)
						.add(props.dayNumber - 3, 'day')
						.format('YYYY-MM-DD')
				]
			}
		} else {
			if (!DayJs(daytime).isValid()) {
				return [DayJs().startOf('isoWeek').format('YYYY-MM-DD'), DayJs().endOf('isoWeek').format('YYYY-MM-DD')]
			} else {
				return [DayJs(daytime).startOf('isoWeek').format('YYYY-MM-DD'), DayJs(daytime).endOf('isoWeek').format('YYYY-MM-DD')]
			}
		}
	}
}
type weekIndextype = 0 | 1 | 2 | 3 | 4 | 5 | 6
//根据开始和结束的时间获取这一段时间内的日期数组。
function getAllDay(start: string, end: string) {
	let ar = []
	let weeksType = []
	let endD = DayJs(end)
	let startd = DayJs(start)
	while (startd.isSameOrBefore(endD)) {
		ar.push({
			date: startd.format('YYYY-MM-DD'),
			str: startd.format('MM/DD')
		})
		let index: weekIndextype = <weekIndextype>startd.get('day') || 0
		if (startd.isSame(DayJs(), 'date')) {
			weeksType.push('今天')
		} else {
			weeksType.push(weekStrIndex[index])
		}

		startd = startd.add(1, 'day')
	}
	weekStr.value = weeksType
	return ar
}
</script>
