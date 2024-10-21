<template>
	<view @click="_show = !_show">
		<!-- #ifdef APP-NVUE -->
		<view :eventPenetrationEnabled="true">
			<slot></slot>
		</view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<slot></slot>
		<!-- #endif -->
		<tm-drawer :disabbleScroll="true" ref="drawer" :round="props.round" :height="dHeight" @close="close" @open="open" :hideHeader="true">
			<view class="mx-16 mt-24">
				<tm-calendar-view
					:format="props.format"
					:hideButton="props.hideButton"
					:hideTool="props.hideTool"
					@update:model-value="_value = $event"
					:model-value="_value"
					@update:model-str="_strvalue = $event"
					:model-str="_strvalue"
					:default-value="_value"
					@change="change"
					@confirm="confirm"
					@click="onclick"
					:model="props.model"
					:color="props.color"
					:linear="props.linear"
					:linearDeep="props.linearDeep"
					:start="props.start"
					:end="props.end"
					:disabledDate="props.disabledDate"
					:multiple="props.multiple"
					:dateStyle="props.dateStyle"
					:max="props.max"
					:textUnit="_textUnit"
					:confirmText="_confirmText"
					ref="calendarView"
				>
				</tm-calendar-view>
			</view>
		</tm-drawer>
	</view>
</template>
<script lang="ts" setup>
/**
 * 日历(弹层式)
 * @description 用法见：tm-calendar-view组件，与其一致的用法。
 */
import { inject, computed, ref, watch, PropType, nextTick, watchEffect, onMounted } from 'vue'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import tmCalendarView from '../tm-calendar-view/tm-calendar-view.vue'
import tmDrawer from '../tm-drawer/tm-drawer.vue'
import { monthDayItem, dateItemStyle, monthYearItem, weekItem, yearItem } from '../tm-calendar-view/interface'
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

const drawer = ref<InstanceType<typeof tmDrawer> | null>(null)
const calendarView = ref<InstanceType<typeof tmCalendarView> | null>(null)

/**
 * 事件说明
 * v-model 绑定当前的时间。
 * confirm 当点击确认时触发
 * click 日期被选中时触发，注意禁用的日期不会触发 。
 * change 当切换年或者月的时候触发。
 */
const emits = defineEmits(['update:modelValue', 'update:modelStr', 'update:show', 'confirm', 'click', 'change', 'cancel', 'close', 'open'])

const props = defineProps({
	...custom_props,
	show: {
		type: Boolean,
		default: false
	},
	/**
	 * 数组
	 */
	defaultValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	modelValue: {
		type: Array as PropType<Array<String | Number | Date>>,
		default: () => []
	},
	//单向绑定输入展示日期，此字段用来在页面上展示。只向外输出。
	//功能目的：用来在页面上显示，特别是在input上绑定显示非常方便。
	//标准数据保存时，请使用modelValue保存，而不是此值。
	modelStr: {
		type: String,
		default: ''
	},
	/**
	 * 日期模式
	 * day : 单个日期选择模式（可多选，需要设置multiple=true;
	 * week :按周选择模式。
	 * month :按月选择模式。
	 * year :按年选择模式。
	 * rang :按日期范围选择模式。
	 */
	model: {
		type: String as PropType<'quarter' | 'day' | 'month' | 'year' | 'rang' | 'week'>,
		default: 'day'
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

	/** 日历组件属性 */

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
	/** 日历组件属性结束 */
	round: {
		type: Number,
		default: 12
	},
	hideButton: {
		type: Boolean,
		default: false
	},
	//隐藏头部操作栏
	hideTool: {
		type: Boolean,
		default: false
	},
	/**modelStr的格式化输出选项，不会影响value值，只对输出值有效 */
	format: {
		type: String,
		default: 'YYYY/MM/DD'
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
const sysinfo = useWindowInfo()
const _show = ref(props.show)
const isConfirm = ref(false) //是否点了确认按钮。
const _value = ref(props.defaultValue)
const _strvalue = ref(props.modelStr)
const _modelType = computed(() => props.model)
const _textUnit= computed(() => props.textUnit)
const _confirmText = computed(() => props.confirmText)
function close() {
	if (!isConfirm.value) {
		emits('cancel')
	}
	emits('close')
	emits('update:show', false)
	isConfirm.value = false
	_show.value = false
}
function open() {
	emits('open')
	emits('update:show', true)
	_show.value = true
}
watchEffect(() => {
	emits('update:modelStr', _strvalue.value)
	emits('update:modelValue', _value.value)
})

watch(
	() => props.show,
	() => {
		if (_show.value == props.show) return

		if (drawer.value) {
			if (props.show) {
				drawer.value?.open()
			} else {
				drawer.value?.close()
			}
		}
	}
)
onMounted(() => {
	if (props.show && drawer.value) {
		drawer.value?.open()
	}
})
watch(
	() => props.modelValue,
	() => {
		_value.value = props.modelValue
		// _strvalue.value = _value.value.join("~");
	},
	{ deep: true }
)
function change(e: Array<string | number>) {
	emits('change', e)
}
function onclick(e: Array<string | number>) {
	emits('click', e)
}
function confirm(e: Array<string | number>) {
	emits('confirm', e)
	drawer.value?.close()
}


const dHeight = computed(() => {
	if (_modelType.value == 'day') return 880 + sysinfo.bottomSafe
	if (_modelType.value == 'rang') return 880 + sysinfo.bottomSafe
	if (_modelType.value == 'week') return 740 + sysinfo.bottomSafe
	if (_modelType.value == 'month') return 720 + sysinfo.bottomSafe
	if (_modelType.value == 'quarter') return 480 + sysinfo.bottomSafe
	if (_modelType.value == 'year') return 620 + sysinfo.bottomSafe
	return 600 +sysinfo.bottomSafe
})
</script>
