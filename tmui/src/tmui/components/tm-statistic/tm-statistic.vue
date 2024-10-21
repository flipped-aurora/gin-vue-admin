<template>
	<view class="flex flex-row flex-row-bottom-center">
		<slot name="prefix">
			<tm-text :followTheme="props.followTheme" :color="props.color" :font-size="props.fontSize * 0.7" :label="props.prefix"></tm-text>
		</slot>
		<slot name="default">
			<tm-text :followTheme="props.followTheme" _class="px-12" :color="props.color" :font-size="props.fontSize" :label="displayValue"></tm-text>
		</slot>
		<slot name="suffix">
			<tm-text :followTheme="props.followTheme" :color="props.color" :font-size="props.fontSize * 0.7" :label="props.suffix"></tm-text>
		</slot>
	</view>
</template>

<script lang="ts" setup>
/**
 * 数值
 * @description 主要用来显示数值，翻转动画。
 * @property {Number} startVal = [] 0,起始值
 * @property {Number} endVal = [] 0,最终值
 * @property {Number} duration = [] 3000,从起始值到结束值数字变动的时间
 * @property {Boolean} autoplay = [] true,是否自动播放
 * @property {Number} decimals = [] 0,保留的小数位数
 * @property {String} decimal = [] '.',小数点分割符号
 * @property {String} separator = [] ',',上了三位数分割的符号
 * @property {String} prefix = [] '',前缀
 * @property {String} suffix = [] '',后缀
 * @property {Boolean} isFrequent = [] false,是否隔一段时间数字跳动，这里的跳动是隔一段时间设置初始值
 * @property {Number} frequentTime = [] 5000,跳动间隔时间
 * @emits [mountedCallback,callback] 事件
 * @template prefix 前缀插槽,default数值内容插槽,suffix后缀插槽。
 * @example <tm-statistic :font-size="80" prefix="￥" suffix="元"></tm-statistic>
 */
import { requestAnimationFrame, cancelAnimationFrame } from './requestAnimationFrame'
import { computed, watch, watchEffect, onMounted, ref, onBeforeUnmount } from 'vue'
import tmText from '../tm-text/tm-text.vue'
const emits = defineEmits(['mountedCallback', 'callback'])
const props = defineProps({
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	fontSize: {
		type: Number,
		default: 30
	},
	color: {
		type: String,
		default: ''
	},
	/**
	 * @description 起始值
	 */
	startVal: {
		type: Number,
		required: false,
		default: 0
	},
	/**
	 * @description 最终值
	 */
	endVal: {
		type: Number,
		required: false,
		default: 2021
	},
	/**
	 * @description 从起始值到结束值数字变动的时间
	 */
	duration: {
		type: Number,
		required: false,
		default: 3000
	},
	/**
	 * @description 是否自动播放
	 */
	autoplay: {
		type: Boolean,
		required: false,
		default: true
	},
	/**
	 * @description 保留的小数位数
	 */
	decimals: {
		type: Number,
		required: false,
		default: 0,
		validator(value) {
			return value >= 0
		}
	},
	decimal: {
		type: String,
		required: false,
		default: '.'
	},
	/**
	 * @description 三位三位的隔开效果
	 */
	separator: {
		type: String,
		required: false,
		default: ','
	},
	/**
	 * @description 前缀
	 * @example '¥' 人民币前缀
	 */
	prefix: {
		type: String,
		required: false,
		default: ''
	},
	/**
	 * @description 后缀
	 * @example
	 */
	suffix: {
		type: String,
		required: false,
		default: ''
	},
	/**
	 * @description 是否具有连贯性
	 */
	useEasing: {
		type: Boolean,
		required: false,
		default: true
	},

	/**
	 * @description 是否隔一段时间数字跳动，这里的跳动是隔一段时间设置初始值
	 */
	isFrequent: {
		type: Boolean,
		required: false,
		default: false
	},
	/**
	 * @description 跳动间隔时间
	 */
	frequentTime: {
		type: Number,
		required: false,
		default: 5000
	}
})
let localStartVal = props.startVal
let displayValue = ref(formatNumber(props.startVal))
let printVal = null
let paused = false
let localDuration = props.duration
let startTime = null
let timestamp = null
let remaining = null
let rAF = null
let timer = null
const countDown = props.startVal > props.endVal ? true : false
watch([() => props.startVal, () => props.endVal], () => {
	start()
})
onMounted(() => {
	if (props.autoplay) {
		start()
	}
	if (props.isFrequent && props.frequentTime) {
		timer = setInterval(() => {
			start(randomNum(0, props.endVal))
		}, props.frequentTime)
	}
	emits('mountedCallback')
})

onBeforeUnmount(() => clearInterval(timer))

function easingFn(t = 0, b = 0, c = 0, d = 0) {
	let p = (c * (-Math.pow(2, (-10 * t) / d) + 1) * 1024) / 1023 + b
	return p
}
function randomNum(a, b) {
	return Math.round(Math.random() * (b - a) + a)
}
function start(startVal) {
	localStartVal = startVal || props.startVal
	startTime = null
	localDuration = props.duration
	paused = false

	rAF = requestAnimationFrame(count)
}
function pauseResume() {
	if (paused) {
		resume()
		paused = false
	} else {
		pause()
		paused = true
	}
}
function pause() {
	cancelAnimationFrame(rAF)
}
function resume() {
	startTime = null
	localDuration = +remaining
	localStartVal = +printVal
	requestAnimationFrame(count)
}
function reset() {
	startTime = null
	cancelAnimationFrame(rAF)
	displayValue.val = formatNumber(props.startVal)
}
function count(timestamp_e) {
	if (!startTime) startTime = timestamp_e
	timestamp = timestamp_e
	const progress = timestamp - startTime
	remaining = localDuration - progress

	if (props.useEasing) {
		if (countDown) {
			printVal = localStartVal - easingFn(progress, 0, localStartVal - props.endVal, localDuration) || 0
		} else {
			printVal = easingFn(progress, localStartVal, props.endVal - localStartVal, localDuration)
		}
	} else {
		if (countDown) {
			printVal = localStartVal - (localStartVal - props.endVal) * (progress / localDuration)
		} else {
			printVal = localStartVal + (props.endVal - localStartVal) * (progress / localDuration)
		}
	}
	if (countDown) {
		printVal = printVal < props.endVal ? props.endVal : printVal
	} else {
		printVal = printVal > props.endVal ? props.endVal : printVal
	}

	displayValue.value = formatNumber(printVal)
	if (progress < localDuration) {
		rAF = requestAnimationFrame(count)
	} else {
		emits('callback')
	}
}
function isNumber(val) {
	return !isNaN(parseFloat(val))
}
function formatNumber(num) {
	num = num.toFixed(props.decimals)
	num += ''
	const x = num.split('.')
	let x1 = x[0]
	const x2 = x.length > 1 ? props.decimal + x[1] : ''
	const rgx = /(\d+)(\d{3})/
	if (props.separator && !isNumber(props.separator)) {
		while (rgx.test(x1)) {
			x1 = x1.replace(rgx, '$1' + props.separator + '$2')
		}
	}
	return x1 + x2
}

// 对外暴露的方法。
defineExpose({ start: start, reset: reset, pause: pause })
</script>

<style></style>
