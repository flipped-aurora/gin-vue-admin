<template>
	<view
		:hover-class="props.hoverClass"
		@click="emits('click', $event)"
		class="flex relative"
		v-if="colWidth"
		:style="[{ width: colWidth + 'px' }, props.height ? { height: props.height + 'rpx' } : '']"
	>
		<view
			:eventPenetrationEnabled="true"
			:style="[
				{
					marginLeft: _margin[0] + 'rpx',
					marginTop: _margin[1] + 'rpx',
					marginRight: _margin[2] + 'rpx',
					marginBottom: _margin[3] + 'rpx',
					borderLeft: `${props.borderGutter[0]}rpx solid ${_borderColor}`,
					borderTop: `${props.borderGutter[1]}rpx solid ${_borderColor}`,
					borderRight: `${props.borderGutter[2]}rpx solid ${_borderColor}`,
					borderBottom: `${props.borderGutter[3]}rpx solid ${_borderColor}`
				},
				!transprent && props.shadow > 0 ? tmcomputed.shadowColor : '',
				!transprent ? tmcomputed.backgroundColorCss : '',
				{ alignItems: alignComputed, justifyContent: justifyComputed },
				customCSSStyle
			]"
			:class="['flex flex-col flex-1 ', `round-${props.round}`, customClass]"
			><slot></slot
		></view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 单元格
 * @description 必须配置tmRow使用。否则报错。
 */
import { getCurrentInstance, computed, ref, provide, inject, PropType } from 'vue'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
const emits = defineEmits(['click'])
const props = defineProps({
	...custom_props,
	height: {
		type: [Number, String],
		default: 50
	},
	color: {
		type: String,
		default: 'white'
	},
	//占据的列数。
	col: {
		type: Number,
		default: 1
	},
	transprent: {
		type: [Boolean, String],
		default: false
	},
	align: {
		type: String as PropType<'start' | 'center' | 'end'>,
		default: 'center' //'start' | 'center' | 'end'
	},
	justify: {
		type: String as PropType<'start' | 'center' | 'end'>,
		default: 'center' //'start' | 'center' | 'end'
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0]
	},
	//这里因为这个组件可能会被大量使用嵌套
	//为了性能，其它属性不再使用主题，而是直接使用值，对性能来说更为重要。
	borderColor: {
		type: String,
		default: 'rgba(0,0,0,0.04)'
	},
	//表示四周的边线，这里与原有的border属性不关联，为了不与前面的属性产生混乱，
	//重新取一个属性作为值来使用。
	borderGutter: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0, 0, 0]
	},
	hoverClass: {
		type: String,
		default: ''
	}
})
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
const _borderColor = computed(() => {
	if (isDark.value) return 'rgba(255,255,255,0.02)'
	return props.borderColor || 'rgba(0,0,0,0.04)'
})
/** 数组是左，上，右，下顺序。 */
const _margin = computed(() => {
	if (props.margin.length == 1) return [props.margin[0], props.margin[0], props.margin[0], props.margin[0]]
	if (props.margin.length == 2) return [props.margin[0], props.margin[1], props.margin[0], props.margin[1]]
	if (props.margin.length == 3) return [props.margin[0], props.margin[1], props.margin[2], 0]
	if (props.margin.length == 4) return [props.margin[0], props.margin[1], props.margin[2], props.margin[3]]
	return [0, 0, 0, 0]
})

const TmRowWidth = inject(
	'TmRowWidth',
	computed(() => 0)
)

const TmRowColumn = inject(
	'TmRowColumn',
	computed(() => 0)
)

const colWidth = computed(() => {
	if (TmRowWidth.value == 0) return 0
	// 因为uni对rpx转换为px之间会有小数点的损耗，导致在Nvue平台计算时，要么多个0.要么少个0.几
	//非常的难以平均的计算，也不能直接平均或者等分计算，不然不同的屏幕会断行。
	//因此实际上的col宽度在nvue时，我们得减小个0.2左右。才能达到等分并在一行上。
	let w = Number((TmRowWidth.value / TmRowColumn.value) * props.col).toFixed(4)
	// #ifdef APP-NVUE
	w = Number((TmRowWidth.value / TmRowColumn.value) * props.col - 0.2).toFixed(4)
	// #endif
	return w
})
//这里是对嵌套的row传递父级宽度。要关掉gutter
provide(
	'TmColWidth',
	computed(() => Number(colWidth.value) - uni.upx2px(_margin.value[0] + _margin.value[2]))
)
//横向对齐方式
let justifyAlign = {
	start: 'flex-start',
	end: 'flex-end',
	center: 'center'
}
const alignComputed = computed(() => justifyAlign[props.align])
const justifyComputed = computed(() => justifyAlign[props.justify])


// 设置响应式主题文字色。
let textColor = computed(() => tmcomputed.value.textColor)
provide('appTextColor', textColor)
</script>

<style></style>
