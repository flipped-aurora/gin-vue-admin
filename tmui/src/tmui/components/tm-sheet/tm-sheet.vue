<!--
 * @Date: 2022-03-29 12:54:41
 * @LastEditors: tmzdy
 * @Author: tmzdy
 * @Description: 文件
-->
<template>
	<view
		:hover-class="(_props.url ? ' opacity-7 ' : '  ') + _props.hoverClass"
		v-if="_blue_sheet"
		:blurEffect="blurEffect"
		@click="onClick"
		@longpress="longpress"
		@touchend="touchend"
		@touchstart="touchstart"
		@touchcancel="touchcancel"
		@mousedown="mousedown"
		@mouseup="mouseup"
		@mouseleave="mouseleave"
		:class="['flex flex-col noNvueBorder', parentClass, !isDisabledRoundAndriod ? round : '']"
		:style="[
			{
				marginLeft: margin[0] + 'rpx',
				marginTop: margin[1] + 'rpx',
				marginRight: margin[2] + 'rpx',
				marginBottom: margin[3] + 'rpx',
				paddingLeft: padding[0] + 'rpx',
				paddingTop: padding[1] + 'rpx',
				paddingRight: padding[2] + 'rpx',
				paddingBottom: padding[3] + 'rpx'
			},
			_height_real ? { height: _height + _props.unit } : '',
			_width_real ? { width: _width + _props.unit } : '',
			tmcomputed.borderCss,
			blur && store.tmStore.os == 'ios' && isNvue === true ? '' : _bgcolor,
			!transparent && _props.shadow > 0 ? tmcomputed.shadowColor : '',
			!transparent && blur ? { backdropFilter: 'blur(6px)' } : '',
			customCSSStyle
		]"
	>
		<view :class="['flex noNvueBorder flex-col flex-1', customClass]" :style="_props.contStyle">
			<slot></slot>

		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 基础容器
 * @description 提供了基础窗口布局，以代替view进行布局，可随意修改样式。
 */
import { computed, ref, provide, watch, PropType, nextTick } from 'vue'
import { custom_props} from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import useTheme from '../../tool/useFun/useTheme'
const store = useTmpiniaStore()
// 混淆props共有参数
const props = defineProps({
	...custom_props,
	parenClass: {
		type: String,
		default: ''
	},
	contStyle: {
		type: String,
		default: ''
	},
	height: {
		type: [Number],
		default: 0
	},
	width: {
		type: [Number],
		default: 0
	},
	color: {
		type: String,
		default: 'white'
	},
	transprent: {
		type: [Boolean, String],
		default: false
	},

	border: {
		type: [Number, String],
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [32]
	},
	padding: {
		type: Array as PropType<Array<number>>,
		default: () => [24]
	},
	unit: {
		type: String,
		default: 'rpx'
	},
	hoverClass: {
		type: String,
		default: 'none'
	},
	//暗下强制的背景色，
	//有时自动的背景，可能不是你想要暗黑背景，此时可以使用此参数，强制使用背景色，
	//只能是颜色值。
	darkBgColor: {
		type: String,
		default: ''
	},
	//不是同层背景，默认是同层，为false
	//如果输入框表单与tmshee在同一层下。他们的黑白暗黑背景色是相同的。为了区分这个问题，需要单独指示，以便区分背景同层。
	//主意：它只在黑和白之间的色系才起作用，其它颜色下不起作用。
	noLevel: {
		type: Boolean,
		default: false
	},
	//是否开启磨砂背景。只有是黑白灰色系才能使用。
	blur: {
		type: Boolean,
		default: false
	},
	url: {
		type: String,
		default: ''
	},
	round: {
		type: [Number,Array] as PropType<Array<number>|number>,
		default: 0
	},
})
const emits = defineEmits(['click', 'longpress', 'touchend', 'touchstart', 'touchcancel', 'mousedown', 'mouseup', 'mouseleave']);

// 设置响应式全局组件库配置表。
const tmcfg = computed(() => store.tmStore);

const {dark,isNvue,customCSSStyle,customClass,parentClass,transparent,_props,proxy,blur,
	round,margin,padding,theme
} = useTheme(computed(()=>props),tmcfg);

//计算主题
const tmcomputed = theme({text:blur.value && tmcfg.value.os == 'ios' && isNvue.value?true:null})

const _width = computed(() => props.width + padding.value[0] + padding.value[2])
const _height = computed(() => props.height + padding.value[1] + padding.value[3])
const _width_real = computed(() => props.width)
const _height_real = computed(() => props.height)

const _blue_sheet = ref(true)
const blurEffect = computed(() => {
	if (props.blur === true && dark.value) return 'dark'
	if (props.blur === true && !dark.value) return 'extralight'
	return 'none'
})
watch(
	() => dark.value,
	() => {
		// #ifdef APP-NVUE
		//在ios下。如果切换带有磨砂效果时，如果不触发发更新视图，页面是不会更改。
		if (store.tmStore.os == 'ios' && blur.value === true) {
			_blue_sheet.value = false
			nextTick(() => (_blue_sheet.value = true))
		}
		// #endif
	}
)
const _bgcolor = computed(() => {
	if (transparent.value === true) return `background-color:rgba(255,255,255,0);`
	if (props.darkBgColor !== '' && dark.value === true) {
		return `background-color:${props.darkBgColor};`
	}

	if (props.linearColor.length == 2) {
		return {
			'background-image': `linear-gradient(${tmcomputed.value.linearDirectionStr},${props.linearColor[0]},${props.linearColor[1]})`
		}
	}
	if (tmcomputed.value.gradientColor?.length == 2) {
		return tmcomputed.value.backgroundColorCss
	}

	if (_props.value.noLevel && tmcomputed.value.isBlackAndWhite === true && dark.value === true) {
		return `background-color: ${tmcomputed.value.inputcolor}`
	}
	return `background-color: ${tmcomputed.value.backgroundColor}`
})
//当前是否点按，因为uniapp的hover-class在安卓端有bug，因此采用自定事件来定义hover类。
const isLongPress = ref(false)
function longpress(e: Event) {
	isLongPress.value = true
	emits('longpress', e)
}
function touchstart(e: Event) {
	isLongPress.value = true
	emits('touchstart', e)
}
function touchend(e: Event) {
	isLongPress.value = false
	emits('touchend', e)
}
function touchcancel(e: Event) {
	isLongPress.value = false
	emits('touchcancel', e)
}
function mousedown(e: Event) {
	isLongPress.value = true
	emits('mousedown', e)
}
function mouseup(e: Event) {
	isLongPress.value = false
	emits('mouseup', e)
}
function mouseleave(e: Event) {
	isLongPress.value = false
	emits('mouseleave', e)
}

function onClick(e: Event) {
	emits('click', e)
	if (typeof props.url === 'string' && props.url) {
		uni.navigateTo({
			url: props.url,
			fail(result) {
				console.log(result)
			}
		})
	}
}


// 设置响应式主题文字色。
let textColor = computed(() => {
	return tmcomputed.value.textColor
})
provide('appTextColor', textColor)

</script>

<style scoped>
/* #ifdef H5 */
.webpc {
	cursor: pointer;
}
/* #endif */
/* #ifndef APP-NVUE */
.noNvueBorder {
	box-sizing: border-box;
}
/* #endif */
</style>
