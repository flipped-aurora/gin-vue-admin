<template>
	<view
		:render-whole="true"
		class="flex flex-row flex-row-center-center"
		:style="[
			{
				marginRight: custom_space_size[0] + 'rpx',
				marginBottom: custom_space_size[1] + 'rpx'
			}
		]"
	>
		<!-- #ifndef APP-NVUE -->
		<text
			@click="clickhandle"
			@longpress="emits('longpress', $event)"
			:class="[props.rotate ? 'ani' : '', spinComputed ? 'spin' : '', 'text-size-n d-inline-block', prefx, iconComputed, customClass]"
			:style="[{ transform: `rotate(${_rotateDeg}deg)` }, fontSizeComputed, { color: textColor }, customCSSStyle]"
			v-if="!isImg"
		></text>
		<!-- #endif  -->
		<!-- #ifdef APP-NVUE-->
		<text
			:render-whole="true"
			ref="icon"
			@click="clickhandle"
			@longpress="emits('longpress', $event)"
			:class="[props.rotate ? 'ani' : '', spinComputed ? 'spin' : '', 'text-size-n d-inline-block ', prefx, customClass]"
			:style="[{ transform: `rotate(${_rotateDeg}deg)` }, { fontFamily: prefx, color: textColor }, fontSizeComputed, customCSSStyle]"
			v-if="!isImg"
		>
			{{ iconComputed }}
		</text>
		<!-- #endif  -->
		<image
			:render-whole="true"
			@click="clickhandle"
			@longpress="emits('longpress', $event)"
			ref="icon"
			v-if="isImg"
			:src="iconComputed"
			:class="[props.rotate ? 'ani' : '', spinComputed ? 'spin' : '', customClass]"
			:style="[
				{ transform: `rotate(${_rotateDeg}deg)` },
				{
					width: (props.fontSize || 30) + props.unit,
					height: (props.fontSize || 30) + props.unit
				},
				customCSSStyle
			]"
		>
		</image>
	</view>
</template>
<script lang="ts" setup>
/**
 * 图标
 * @description 图标，提供了一个spin功能用于自动旋转图标在Nvue中使用原生动画。
 */

import { getCurrentInstance, computed, ref, provide, inject, onMounted, onUnmounted, nextTick, watch, onBeforeMount } from 'vue'
import theme from '../../tool/theme/theme'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
// #ifdef APP-NVUE || APP-PLUS-NVUE
import { tmiconFont } from './tmicon'
import { toUnicode } from 'punycode'
var domModule = weex.requireModule('dom')
const animation = uni.requireNativePlugin('animation')
// #endif
const store = useTmpiniaStore()
// 混淆props共有参数
const props = defineProps({
	...custom_props,
	fontSize: {
		type: [Number],
		default: 34
	},
	color: {
		type: String,
		default: ''
	},
	/** 自定义图标规则为:myicon-music-e617,图标前缀和字体名称相同-图标类名-图标unicode符 */
	name: {
		type: String, //图标名称。
		default: ''
	},
	spin: {
		type: [Boolean],
		default: false
	},
	unit: {
		type: String,
		default: 'rpx'
	},
	//-1表示自动
	lineHeight: {
		type: [Number],
		default: -1
	},
	rotate: {
		type: Boolean,
		default: false
	},
	rotateDeg: {
		type: Number,
		default: 0
	},
	/**
	 * 为了提高响应速度，只有开启了自定图标显示功能才会去解析用户自定义图标规则名称
	 */
	customicon: {
		type: Boolean,
		default: false
	}
})
const _rotateDeg = computed(() => props.rotateDeg)
const emits = defineEmits(['click', 'longpress'])
const proxy = getCurrentInstance()?.proxy ?? null
// 设置响应式全局组件库配置表。
const tmcfg = computed<tmVuetify>(() => store.tmStore)
//自定义样式：
const customCSSStyle = computed(() => computedStyle(props))
//自定类
const customClass = computed(() => computedClass(props))
//是否暗黑模式。
const isDark = computed(() => computedDark(props, tmcfg.value))
//计算主题
const tmcomputed = computed<cssstyle>(() => computedTheme(props, isDark.value, tmcfg.value))
// 点击文字事件。
function clickhandle(e: Event): void {
	emits('click', e)
}
//从父应用组件中获取自动文字色。
const appTextColor = inject(
	'appTextColor',
	computed(() => undefined)
)
const textColor = computed(() => {
	if (props.followTheme && store.tmStore.color) return store.tmStore.color
	let isColorHex = theme.isCssColor(props.color)
	//如果当前是自定义颜色值，直接返回。
	if (isColorHex) return props.color
	//如果定义了颜色，但不是值，而是主题色，则返回对应的主题文本色。
	if (props.color && !isColorHex) {
		let nowcolor: colorThemeType = theme.getColor(props.color)

		return nowcolor.csscolor
	}
	if (appTextColor.value) return appTextColor.value
	return 'rgba(34, 34, 34, 1.0)'
})
//图标大小。
const fontSizeComputed = computed(() => {
	// #ifdef H5
	if (props.fontSize < 24 && props.unit == 'rpx')
		return {
			transform: 'scale(0.8)',
			fontSize: (props.fontSize || 30) + props.unit,
			lineHeight: props.lineHeight > -1 ? props.lineHeight + props.unit : (props.fontSize || 30) + props.unit
		}
	// #endif
	let strc = {
		fontSize: (props.fontSize || 30) + props.unit,
		lineHeight: props.lineHeight > -1 ? props.lineHeight + props.unit : (props.fontSize || 30) + props.unit
	}
	if (props.lineHeight == 0) {
		delete strc.lineHeight
	}
	return strc
})
//当前图标是否是图片。
const isImg = computed(() => {
	if (
		props.name[0] == '.' ||
		props.name[0] == '@' ||
		props.name[0] == '/' ||
		props.name[0] == '~' ||
		props.name.substring(0, 5) == 'data:' ||
		props.name.substring(0, 4) == 'http' ||
		props.name.substring(0, 5) == 'https' ||
		props.name.substring(0, 3) == 'ftp'
	) {
		return true
	}
	return false
})
//图标前缀
const prefx = computed(() => {
	let prefix = props.name.split('-')?.[0]
	return prefix
})
//图标名称。
const iconComputed = computed(() => {
	if (isImg.value) return props.name
	// #ifdef APP-NVUE
	let name = props.name.substr(prefx.value.length + 1)
	let index = uni.$tm.tmicon.findIndex((el) => el.font == prefx.value)
	let itemIcon = uni.$tm.tmicon[index]?.fontJson.find((item, index) => {
		return item.font_class == name
	})
	try {
		if (itemIcon) {
			let ucode = '"\\u' + String(itemIcon?.unicode) + '"'
			return JSON.parse(ucode)
		} else if (props.customicon) {
			let names = name.split('-')
			if (names.length >= 2) {
				let ucode = '"\\u' + String(names[names.length - 1]) + '"'
				return JSON.parse(ucode)
			}
		}
	} catch (e) {
		return props.name
	}
	// #endif

	// #ifndef APP-NVUE
	if (props.customicon) {
		try {
			let names = props.name.split('-')
			if (/^e[0-9|a-z|A-Z]{3}/.test(names[names.length - 1])) {
				let clasName = props.name.substring(0, props.name.lastIndexOf('-'))
				return clasName
			}
		} catch (e) {
			//TODO handle the exception
		}
	}

	// #endif

	return props.name
})

//是否使图标旋转。
const spinComputed = computed(() => props.spin)
//间隙排列。
const custom_space_size = inject('custom_space_size', [0, 0])
//图标的宽度和高度
const iconWidth = computed(() => Math.ceil(props.fontSize || 34) + custom_space_size[0])
const iconHeight = computed(() => Math.ceil(props.fontSize || 34) + custom_space_size[1])

function spinNvueAni(jiaodu = 360) {
	let iconEl = proxy?.$refs['icon']
	if (!iconEl) return
	animation.transition(
		iconEl,
		{
			styles: {
				transform: `rotate(${jiaodu}deg)`,
				transformOrigin: 'center center'
			},
			duration: 2000, //ms
			timingFunction: 'linear',
			delay: 0 //ms
		},
		() => {
			nextTick(function () {
				animation.transition(
					iconEl,
					{
						styles: {
							transform: `rotate(0deg)`,
							transformOrigin: 'center center'
						},
						duration: 0, //ms
						timingFunction: 'ease',
						delay: 0 //ms
					},
					() => {
						spinNvueAni()
					}
				)
			})
		}
	)
}
watch(spinComputed, () => {
	// #ifdef APP-PLUS-NVUE
	if (spinComputed.value) {
		nextTick(function () {
			spinNvueAni()
		})
	}
	// #endif
})
onBeforeMount(() => {
	// #ifdef APP-PLUS-NVUE
	domModule.addRule('fontFace', {
		fontFamily: 'tmicon', //注意这里必须是驼峰命名，跟上面的css样式对照
		src: "url('data:font/ttf;charset=utf-8;base64," + tmiconFont + "')"
	})
	// #endif
})
onMounted(() => {
	// #ifdef APP-PLUS-NVUE
	if (spinComputed.value) {
		setTimeout(function () {
			spinNvueAni()
		}, 50)
	}
	// #endif
})
</script>

<style scoped="scoped">
.ani {
	transition-duration: 0.3s;
	transition-timing-function: ease;
	transition-property: transform;
}
/* #ifndef APP-PLUS-NVUE */
.spin {
	transform-origin: 50% 50%;
	animation: xhRote 1.2s infinite linear;
}

@keyframes xhRote {
	0% {
		transform: rotate(0deg);
	}

	100% {
		transform: rotate(360deg);
	}
}

/* #endif */
</style>
