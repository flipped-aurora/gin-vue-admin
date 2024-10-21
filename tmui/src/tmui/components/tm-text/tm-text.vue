<template>
	<!-- #ifdef APP-PLUS-NVUE -->
	<text
		:render-whole="true"
		@click="emits('click', $event)"
		:selectable="selectable"
		:user-select="selectable"
		:class="[_fontSize ? '' : 'text-size-m', customClass]"
		:style="[
			props.lineHeight == 'auto' ? { lineHeight: (_fontSize ? _fontSize * 1.3 : 42) + props.unit } : {},
			props.lineHeight > 0 ? { lineHeight: props.lineHeight + props.unit } : {},
			{
				color: textColor
			},
			_fontSize ? { fontSize: _fontSize + props.unit } : '',
			customCSSStyle
		]"
		>{{ _label }}</text
	>
	<!-- #endif -->
	<!-- #ifndef APP-PLUS-NVUE -->
	
	<view :render-whole="true" class="flex text-view nv" :class="[_parentClass]">
		<text
			@click="emits('click', $event)"
			:selectable="selectable"
			:user-select="selectable"
			:class="[fontSize ? '' : 'text-size-m', customClass]"
			:style="[
				props.lineHeight == 'auto' ? { lineHeight: (_fontSize ? _fontSize * 1.3 : 42) + props.unit } : {},
				props.lineHeight > 0 ? { lineHeight: props.lineHeight + props.unit } : {},
				{
					color: textColor
				},
				_fontSize ? { fontSize: _fontSize + props.unit } : '',
				customCSSStyle
			]"
		>
			<slot>{{ _label }}</slot>
		</text>
	</view>
	<!-- #endif -->
</template>
<script lang="ts" setup>
/**
 * 文本
 * @description 自带主题和常用属性，能根据全局主题和暗黑自动切换，必须放在tmSheet下，获得更好的主题适应能力。
 */
import { computed, getCurrentInstance, inject, ref, VNode } from 'vue'
import theme from '../../tool/theme/theme'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
// 混淆props共有参数
const props = defineProps({
	...custom_props,
	label: {
		type: [String, Number],
		default: ''
	},
	fontSize: {
		type: [Number],
		default: 28
	},
	color: {
		type: String,
		default: ''
	},
	selectable: {
		type: [Boolean],
		default: false
	},
	unit: {
		type: String,
		default: 'rpx'
	},
	parentClass: {
		type: String,
		default: ''
	},
	lineHeight: {
		type: [Number, String],
		default: 'auto'
	},
	/**
	 * 微信端根据微信字号设置自动适配文字大小。
	 */
	aging:{
		type:Boolean,
		default:false
	}
})
const emits = defineEmits(['click'])

const appBaseInfo = uni.getAppBaseInfo()

const _parentClass = computed(() => props.parentClass)
// 设置响应式全局组件库配置表。
const tmcfg = computed<tmVuetify>(() => store.tmStore)
//自定义样式：
const customCSSStyle = computed(() => computedStyle(props))
//自定类
const customClass = computed(() => computedClass(props))
//是否暗黑模式。
const isDark = computed(() => computedDark(props, tmcfg.value))
//计算主题
// const tmcomputed = computed(() => computedTheme(props, isDark.value));
const _label = computed(() => props.label)
const _fontSize = computed(() => {
	let times = 1
	//#ifdef MP-WEIXIN
	if(props.aging){
		times = (Math.floor(Number(appBaseInfo?.fontSizeSetting??16) / 16 * 10)) / 10
	}
	// #endif
	return Number(props.fontSize) * (store?.tmuiConfig?.themeConfig?.globalFontSizeRatio ?? 1) * times
})


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
	if (!appTextColor) {
		if (isDark) return 'rgba(252, 252, 252, 1.0)'
		return 'rgba(34, 34, 34, 1.0)'
	}
	if (appTextColor.value) {
		return appTextColor.value
	}
	return 'rgba(34, 34, 34, 1.0)'
})
</script>
<style scoped>
/* #ifndef APP-NVUE */
.text-view {
	box-sizing: border-box;
}

.nv {
  display: inline-flex;
  vertical-align: middle;
  flex-shrink:1;
}

/* #endif */
</style>
