<template>
	<view
		ref="tmRow"
		@click.stop="emits('click', $event)"
		class="flex tm-row"
		:class="['overflow ', _round, customClass, `mx-${props.margin[0]} my-${margin[1]}`]"
		:style="[
			{ flexDirection: 'row', flexWrap: 'wrap' },
			props.height ? { height: props.height + 'rpx' } : '',
			width_px_rect_rp ? { width: width_px_rect + 'px' } : '',
			{ justifyContent: justify_rp, alignItems: align_rp },
			!props.transprent ? tmcomputed.backgroundColorCss : '',
			!props.transprent ? tmcomputed.shadowColor : '',
			customCSSStyle
		]"
	>
		<slot></slot>
	</view>
</template>

<script lang="ts" setup>
/**
	 * 布局row
	 * @description 布局，必须配合tmCol，采用flex布局。必须在class上写flex-x，x=[1,12]
	 * row和col为了高性能，和兼容全平台，全部使用flex布局，因此用起来可能有点不方便。但这是值得的。
	 * 如果row设定了宽度，子col将会自动设定宽度。此时如果要切换，必须设定column列数。
	 * @example 例子中演示了col套入row时，如何让col中的row进行自动100%宽度。切换不可直接写宽度100%,请使用专用类来编写。
	 * <tm-row >
			<tm-col class="flex-1">1</tm-col>
			<tm-col class="flex-1">1</tm-col>
			<tm-col class="flex-1 " _class="flex-row flex">
				<tm-row class="flex-12 fulled">
					<tm-col v-for="item in 3" :key="item" class="flex-1">1</tm-col>
				</tm-row>
			</tm-col>
		</tm-row>
	 */
import { getCurrentInstance, computed, ref, provide, inject, onUpdated, onMounted, onUnmounted, nextTick, watch, PropType, watchEffect } from 'vue'
import { cssstyle, tmVuetify, colorThemeType } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
type justifyAlignType = 'start' | 'end' | 'center' | 'around' | 'between'
type AlignAlignType = 'start' | 'end' | 'center' | 'stretch'
const store = useTmpiniaStore()
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const props = defineProps({
	...custom_props,
	height: {
		type: [Number, String],
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	width: {
		type: [Number, String],
		default: 0
	},
	round: {
		type: [Number, String],
		default: 0
	},
	//总列数。
	column: {
		type: Number,
		default: 12
	},
	//横向排列
	justify: {
		type: String as PropType<justifyAlignType>,
		default: 'start' //'start' | 'center' | 'end' | 'around' | 'between'
	},
	//纵向排列
	align: {
		type: String as PropType<AlignAlignType>,
		default: 'center' //'start' | 'center' | 'end' | 'stretch'
	},
	color: {
		type: String,
		default: 'white'
	}
})
const emits = defineEmits(['click'])
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
//取得嵌套的col宽度
const colWidth = inject(
	'TmColWidth',
	computed(() => 0)
)

//宽度，
const width_px_rect = ref(uni.upx2px(Number(props.width)))
const width_px_rect_rp = computed(() => width_px_rect.value)
//横向对齐方式
const justifyAlign = {
	start: 'flex-start',
	end: 'flex-end',
	center: 'center',
	around: 'space-around',
	between: 'space-between'
}
const justify_rp = computed(() => justifyAlign[props.justify] || 'start')
const AlignAlign = {
	start: 'flex-start',
	end: 'flex-end',
	center: 'center',
	stretch: 'stretch'
}
const align_rp = computed(() => AlignAlign[props.align] || 'start')

/** 数组是左，上，右，下顺序。 */
const _round= computed(() => {
	if(typeof props.round == 'number') return 'round-'+props.round
	if (props.round.length == 1) return 'round-'+props.round
	if (props.round.length == 2) return `round-tl-${props.round[0]} round-tr-${props.round[1]}`
	if (props.round.length == 3) return `round-tl-${props.round[0]} round-tr-${props.round[1]} round-br-${props.round[2]} `
	if (props.round.length == 4) return `round-tl-${props.round[0]} round-tr-${props.round[1]} round-br-${props.round[2]}  round-bl-${props.round[2]}`
	return [0, 0, 0, 0]
})


function wxmpGetRect() {
	if (width_px_rect.value > 0) return
	uni.createSelectorQuery()
		.in(proxy)
		.select('.tm-row')
		.boundingClientRect()
		.exec(function (res) {
			if (res[0]?.width) {
				width_px_rect.value = res[0]?.width
			} else {
				wxmpGetRect()
			}
		})
}
function nvueGetRect() {
	if (width_px_rect.value > 0) return
	// #ifdef APP-PLUS-NVUE
	try {
		dom.getComponentRect(proxy.$refs.tmRow, function (res) {
			if (res?.size && res?.size?.width) {
				width_px_rect.value = res.size.width
			} else {
				nvueGetRect()
			}
		})
	} catch (e) {
		//TODO handle the exception
	}
	// #endif
}

onMounted(() => {
	// #ifndef APP-NVUE
	wxmpGetRect()
	// #endif
	// #ifdef APP-NVUE
	nvueGetRect()
	// #endif
})

onUpdated(() => {
	// #ifndef APP-NVUE
	wxmpGetRect()
	// #endif
	// #ifdef APP-NVUE
	nvueGetRect()
	// #endif
})

//对col子组件暴露数据。
provide('TmRowWidth', width_px_rect_rp) //微信端面使用
provide(
	'TmRowColumn',
	computed(() => props.column)
)

watchEffect(() => {
	if (colWidth.value > 0) {
		width_px_rect.value = colWidth.value
	}
})

// 设置响应式主题文字色。
let textColor = computed(() => tmcomputed.value.textColor)
provide('appTextColor', textColor)
</script>

<style></style>
