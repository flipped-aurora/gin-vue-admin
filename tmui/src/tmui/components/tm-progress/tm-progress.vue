<template>
	<view class="flex flex-col">
		<view
			v-if="model == 'line'"
			:style="[{ width: width + 'rpx', paddingTop: '16rpx', paddingBottom: '16rpx' }]"
			class="flex relative flex flex-col overflow"
		>
			<view class="relative">
				<tm-sheet
					no-level
					:round="props.round"
					:followTheme="false"
					:dark="props.dark"
					:margin="[0, 0]"
					:width="props.width"
					:height="props.height"
					:color="props.bgColor"
					:padding="[0, 0]"
				>
				</tm-sheet>
			</view>
			<view class="absolute l-0 tmRogress overflow" :style="[{ width: activeWidth + 'rpx', top: 16 + 'rpx' }]">
				<tm-sheet
					:round="props.round"
					:followTheme="props.followTheme"
					:dark="props.dark"
					:margin="[0, 0]"
					:linear="props.linear"
					:linearDeep="props.linearDeep"
					:width="activeWidth"
					:height="props.height"
					:color="props.color"
					:padding="[0, 0]"
				>
				</tm-sheet>
			</view>
			<view
				v-if="props.showBar"
				class="absolute l-0 t-0 tmRogress flex flex-col"
				:style="[
					{
						width: activeWidth + 'rpx',
						height: props.height + 32 + 'rpx',
						'align-items': 'flex-end',
						'justify-content': 'center'
					}
				]"
			>
				<slot>
					<tm-sheet
						:linear="props.linear"
						:linearDeep="props.linearDeep"
						:followTheme="props.followTheme"
						:dark="props.dark"
						:color="props.color"
						:margin="[0, 0]"
						:padding="[12, 4]"
						:round="4"
					>
						<tm-text _class="text-size-xxs" :fontSize="22" :label="props.percent + props.percentSuffix"> </tm-text>
					</tm-sheet>
				</slot>
			</view>
		</view>
		<view
			v-if="model == 'circle'"
			:style="{
				width: `${props.width}rpx`,
				height: `${props.semicircle ? props.width / 2 : props.width}rpx`
			}"
			class="flex relative flex-col"
		>
			<!-- #ifdef APP-NVUE -->
			<view v-if="showGc">
				<gcanvas
					:id="canvasId"
					:ref="canvasId"
					class="canvas"
					:style="{
						width: `${props.width}rpx`,
						height: `${props.semicircle ? props.width / 2 : props.width}rpx`
					}"
				>
				</gcanvas>
			</view>
			<!-- #endif -->
			<!-- #ifdef MP-WEIXIN  || MP-QQ -->
			<canvas
				type="2d"
				id="canvasId"
				canvas-id="canvasId"
				class="canvas"
				:style="{
					width: `${props.width}rpx`,
					height: `${props.semicircle ? props.width / 2 : width}rpx`
				}"
			></canvas>
			<!-- #endif -->
			<!-- #ifdef MP-ALIPAY -->
			<canvas
				type="2d"
				:id="canvasId"
				:canvas-id="canvasId"
				class="canvas"
				:style="{
					width: `${props.width}rpx`,
					height: `${props.semicircle ? props.width / 2 : width}rpx`
				}"
			></canvas>
			<!-- #endif -->
			<!-- #ifndef MP-WEIXIN || MP-ALIPAY || MP-QQ || APP-NVUE -->
			<canvas
				:id="canvasId"
				:canvas-id="canvasId"
				class="canvas"
				:style="{
					width: `${props.width}rpx`,
					height: `${props.semicircle ? props.width / 2 : props.width}rpx`
				}"
			></canvas>
			<!-- #endif -->
			<!-- #ifndef APP-NVUE -->
			<cover-view
				:style="[
					{
						width: `${props.width}rpx`,
						height: `${props.semicircle ? props.width / 2 : props.width}rpx`
					},
					props.semicircle && props.semicircleFlip ? { 'justify-content': 'flex-start', 'align-items': 'center' } : '',
					props.semicircle && !props.semicircleFlip ? { 'justify-content': 'flex-end', 'align-items': 'center' } : ''
				]"
				class="relative absolute l-0 t-0 flex flex-col"
				:class="[!props.semicircle ? 'flex-center' : '']"
			>
				<cover-view v-if="props.showBar" :style="[{ fontSize: props.fontSize + 'rpx', color: isDark ? darkcolor : txtcolor }]">
					<slot name="title">
						{{ props.percent + props.percentSuffix }}
					</slot>
				</cover-view>
			</cover-view>
			<!-- #endif -->
			<!-- #ifdef APP-NVUE -->
			<cover-view
				:style="[
					{
						width: `${props.width}rpx`,
						height: `${props.semicircle ? props.width / 2 : props.width}rpx`
					},
					props.semicircle && props.semicircleFlip ? { 'justify-content': 'flex-start', 'align-items': 'center' } : '',
					props.semicircle && !props.semicircleFlip ? { 'justify-content': 'flex-end', 'align-items': 'center' } : ''
				]"
				class="relative absolute l-0 t-0 flex flex-col"
				:class="[!props.semicircle ? 'flex-center' : '']"
			>
				<slot name="title">
					<tm-text
						v-if="props.showBar"
						:color="props.color"
						:followTheme="props.followTheme"
						:dark="props.dark"
						:fontSize="props.fontSize"
						:label="props.percent + props.percentSuffix"
					>
					</tm-text>
				</slot>
			</cover-view>
			<!-- #endif -->
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 进度条
 * @description 进度条，圆形进度条，在不同的平台使用2d或者webgl方法，使得性能更加强劲。NVUE中貌似是uniapp的插件bug无法实现渐变绘制。
 * Nvue中需要在manifest.json中设置canvas模块，才能打包详见：https://github.com/dcloudio/NvueCanvasDemo
 */
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedDark, computedTheme } from '../../tool/lib/minxs'
import { getCurrentInstance, computed, ref, provide, inject, onBeforeMount, onMounted, onUnmounted, nextTick, watch, PropType } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tool from '../../tool/theme/theme'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import { getCanvas } from '../../tool/function/getCanvas'
import * as TWEEN from '../../tool/lib/tween.min.js'
// #ifdef APP-NVUE
import { enable, WeexBridge } from '../../tool/gcanvas/index.js'
// #endif
const store = useTmpiniaStore()
const emits = defineEmits(['update:percent', 'change'])
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	model: {
		type: String as PropType<'line' | 'circle'>,
		default: 'line' //line,circle
	},
	//model==circle,是否是半圆。
	semicircle: {
		type: [Boolean, String],
		default: false
	},
	//model==circle有效,半圆正常是在上方。如果反转就在下方。
	semicircleFlip: {
		type: [Boolean],
		default: false
	},
	//model==circle有效
	fontSize: {
		type: [Number, String],
		default: 28
	},
	//进度 百分比数值，不带%号。也可以使用v-model:percent
	percent: {
		type: Number,
		default: 0
	},
	//数值后缀。默认为%
	percentSuffix: {
		type: String,
		default: '%'
	},
	width: {
		type: Number,
		default: 120
	},
	height: {
		type: Number,
		default: 6
	},
	bgColor: {
		type: String,
		default: 'grey-4'
	},
	color: {
		type: String,
		default: 'primary'
	},
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	//暗黑
	dark: {
		type: [Boolean],
		default: false
	},
	linear: {
		type: [String],
		default: '' //left:右->左，right:左->右。top:下->上，bottom:上->下。
	},
	// 渐变的亮浅
	linearDeep: {
		type: [String],
		default: 'light' //light,dark,accent亮系渐变和深色渐变。
	},
	round: {
		type: [Number, String],
		default: 3
	},
	//显示数值点
	showBar: {
		type: [Boolean, String],
		default: false
	},
	disabled: {
		type: Boolean,
		default: false
	}
})
const canvasId = ref('canvasId')
// #ifndef MP-WEIXIN || MP-QQ
canvasId.value = 'tm' + String(new Date().getTime())
// #endif
let ctx: UniApp.CanvasContext
let reqId: number
let tween: any
const shadow_pr = computed(() => props.shadow * 4)
// 设置响应式全局组件库配置表。
const tmcfg = computed<tmVuetify>(() => store.tmStore)
//是否暗黑模式。
const isDark = computed(() => computedDark(props, tmcfg.value))
//计算主题
const tmcomputed = computed<cssstyle>(() => computedTheme(props, isDark.value, tmcfg.value))
const showGc = ref(false)
const sys = uni.getSystemInfoSync()
let isAndroid = false
// #ifdef APP-NVUE
if (sys.osName == 'android') {
	isAndroid = true
}
// #endif
const _bgColor = computed(
	() => computedTheme({ ...props, color: props.bgColor, followTheme: false, linear: '' }, isDark.value, tmcfg.value).backgroundColor
)
const txtcolor = tool.getColor(props.color).value
const darkcolor = tmcomputed.value.backgroundColor
const activeWidth = computed(() => {
	let pr = props.percent >= 100 ? 100 : props.percent
	let w = Math.floor((pr / 100) * props.width)
	w = w >= props.width ? props.width : w
	return w
})
const percent_rp = computed(() => {
	let pr = props.percent >= 100 ? 100 : props.percent
	return pr
})
watch(
	() => props.percent,
	(newval, oldval) => {
		if (props.disabled) return
		tween = new TWEEN.Tween({ o: oldval })
			.to({ o: newval }, 240)
			.easing(TWEEN.Easing.Linear.None)
			.onUpdate((e: any) => {
				draw(e.o)
			})
			.start()
		emits('update:percent', newval)
		emits('change', newval)
	}
)
onBeforeMount(() => {
	clearTimeout(reqId)
	reqFun()
})
onMounted(() => {
	let delay = 10
	// #ifdef APP-NVUE
	if (isAndroid) {
		delay = 250
	} else {
		delay = 100
	}
	// #endif
	// #ifdef MP
	delay = 60
	// #endif
	// #ifdef APP-VUE
	delay = 30
	// #endif
	// #ifdef APP-NVUE
	setTimeout(function () {
		showGc.value = true
		setTimeout(() => {
			getCanvas(proxy, canvasId.value, uni.upx2px(props.width), uni.upx2px(props.width))
				.then((e) => {
					ctx = e.ctx
					tween = new TWEEN.Tween({ o: 0 })
						.to({ o: props.percent }, 240)
						.easing(TWEEN.Easing.Linear.None)
						.onUpdate((e: any) => {
							draw(e.o)
						})
						.start()
				})
				.catch((error) => {
					console.error(error)
				})
		}, delay)
	}, 200)
	// #endif

	// #ifndef APP-NVUE
	setTimeout(() => {
		getCanvas(proxy, canvasId.value, uni.upx2px(props.width), uni.upx2px(props.width))
			.then((e) => {
				ctx = e.ctx
				tween = new TWEEN.Tween({ o: 0 })
					.to({ o: props.percent }, 240)
					.easing(TWEEN.Easing.Linear.None)
					.onUpdate((e: any) => {
						draw(e.o)
					})
					.start()
			})
			.catch((error) => {
				console.error(error)
			})
	}, delay)
	// #endif
})
onUnmounted(() => uni.$tm.u.cancelAnimationFrame(reqId))
function reqFun() {
	TWEEN.update()
	reqId = uni.$tm.u.requestAnimationFrame(reqFun)
}

function draw(bl: number = 0) {
	if (!ctx) return
	let c: any = tmcomputed.value
	const width = uni.upx2px(props.width)
	let x = width / 2
	let y = width / 2
	const borderWidth = uni.upx2px(props.height)
	const radius = x - borderWidth
	const Pi = Math.PI / 180
	let startAngle = -Pi * 90
	let endAngle = Math.PI * 2
	if (props.semicircle) {
		startAngle = -Pi * 90 * 2
		endAngle = 0
		if (props.semicircleFlip) {
			startAngle = Pi * 180
			endAngle = 0
		}
	}
	//圆环背景色
	let bgColor = _bgColor.value || '#f5f5f5'
	//圆环进度条的颜色
	let activeColor = tool.getColor(props.color).csscolor || '#ff0000'

	ctx.clearRect(0, 0, width, width)
	//先绘制背景圆;
	ctx.lineWidth = borderWidth
	ctx.strokeStyle = bgColor
	ctx.lineCap = 'round'
	ctx.beginPath()
	if (props.semicircle) {
		// 半圆
		if (props.semicircleFlip) {
			y = borderWidth / 2
			ctx.arc(x, y, radius, 0, Pi * 180, false)
		} else {
			y = y - borderWidth / 2
			ctx.arc(x, y, radius, Pi * 180, 0, false)
		}
	} else {
		ctx.arc(x, y, radius, -Math.PI / 2, 1.5 * Math.PI)
	}
	ctx.stroke()

	// 绘制圆形进度环
	const progress = Number((bl / 100).toFixed(6)) // 进度比例，0-1之间的数值
	ctx.beginPath()
	if (props.semicircle) {
		// 半圆
		if (props.semicircleFlip) {
			ctx.arc(x, y, radius, startAngle * (1 - progress), startAngle, false)
		} else {
			ctx.arc(x, y, radius, startAngle, startAngle * (1 - progress), false)
		}
	} else {
		if (props.semicircleFlip) {
			ctx.arc(x, y, radius, startAngle + endAngle * (1 - progress), startAngle + endAngle, false)
		} else {
			ctx.arc(x, y, radius, startAngle, startAngle + endAngle * progress, false)
		}
	}
	if (ctx.setLineWidth) {
		ctx.setLineWidth(borderWidth)
	} else {
		ctx.lineWidth = borderWidth
	}

	// #ifndef APP-NVUE
	//如果是渐变
	if (props.linear) {
		let gradient: any = ctx.createLinearGradient(borderWidth / 2, borderWidth / 2, x * 2 - borderWidth, y * 2 - borderWidth)
		if (props.semicircle) {
			gradient = ctx.createLinearGradient(borderWidth / 2, borderWidth / 2, x * 2 - borderWidth, y)
		}
		gradient.addColorStop(0, c.gradientColor[0])
		gradient.addColorStop(1, c.gradientColor[1])
		// ctx.setStrokeStyle(gradient)
		if (ctx.setStrokeStyle) {
			ctx.setStrokeStyle(gradient)
			ctx.setShadow(c.gradientColor[0])
		} else {
			ctx.strokeStyle = gradient
			ctx.shadowColor = c.gradientColor[0]
		}
	} else {
		if (ctx.setStrokeStyle) {
			ctx.setStrokeStyle(activeColor)
			ctx.setShadow(activeColor)
		} else {
			ctx.strokeStyle = activeColor
			ctx.shadowColor = activeColor
		}
	}

	// #endif
	// #ifdef APP-NVUE
	if (ctx.setStrokeStyle) {
		ctx.setStrokeStyle(activeColor)
	} else {
		ctx.strokeStyle = activeColor
	}
	// #endif

	ctx.stroke()
	if (ctx?.draw) {
		ctx.draw()
	}
}
</script>
<style>
.tmRogress {
	transition-duration: 0.34s;
	transition-timing-function: linear;
	transition-property: width;
}

/* #ifndef APP-NVUE */
cover-view {
	background: transparent;
}

/* #endif */
</style>
