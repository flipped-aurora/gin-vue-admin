<template>
	<view ref="tmspin" :class="[customClass]" :style="[customCSSStyle]" class="flex relative flex-col">
		<view class="relative zIndex-1">
			<slot></slot>
		</view>
		<view
			v-if="loadingComputed"
			@click.stop="clickMask"
			class="absolute zIndex-10 flex-center blur"
			:class="[css_is_nvue ? '' : ' fulled-height   ']"
			:style="[
				css_is_nvue ? '' : { width: '100%' },
				{ backgroundColor: bgColorComputed },
				css_is_nvue
					? {
							width: (css_nvue_size[0] || props.fontSize) + 'px',
							height: (css_nvue_size[1] || props.fontSize) + 'px'
					  }
					: ''
			]"
		>
			<view :userInteractionEnabled="false" class="pa-10 flex-col flex-col-center-center">
				<tm-icon
					spin
					:fontSize="props.size"
					:dark="isDark"
					:color="_color"
					:followDark="followDark"
					:followTheme="props.followTheme"
					:name="props.icon"
				></tm-icon>
				<tm-text
					:followTheme="props.followTheme"
					:dark="isDark"
					:followDark="followDark"
					:color="_color"
					_class="mt-16"
					:label="props.tip"
				></tm-text>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 局部加载
 * @description 放置在其里面的所有内容，加载状态时，会在其中所有内容上添加遮罩模糊加载页面，并且不可操作。
 */

import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'

import { getCurrentInstance, computed, ref, provide, inject, onMounted, onUnmounted, onUpdated, nextTick, watch } from 'vue'
import { cssstyle, tmVuetify } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedClass, computedStyle, computedDark } from '../../tool/lib/minxs'

import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()

// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
// 混淆props共有参数
const props = defineProps({
	...custom_props,
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	icon: {
		type: String,
		default: 'tmicon-shuaxin'
	},
	size: {
		type: Number,
		default: 46
	},
	bgColor: {
		type: String,
		default: 'rgba(255,255,255,0.9)'
	},
	tip: {
		type: [String],
		defalut: ''
	},
	load: {
		type: [Boolean, String],
		default: false
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
// const tmcomputed = computed<cssstyle>(() => computedTheme(props, isDark.value));
// 点击文字事件。
function clickhandle(e: Event): void {
	emits('click', e)
}
const css_is_nvue = ref(true)
// #ifndef APP-NVUE
css_is_nvue.value = false
// #endif
const css_nvue_size = ref([0, 0])
const bgColorComputed = computed(() => (isDark.value ? 'rgba(0,0,0,0.9)' : props.bgColor))
const loadingComputed = computed(() => props.load)
const _color = computed(() => props.color)
onUpdated(() => {
	// #ifdef APP-PLUS-NVUE
	nvuegetClientRect()
	// #endif
})
onMounted(() => {
	// #ifdef APP-PLUS-NVUE
	nvuegetClientRect()
	// #endif
})
function nvuegetClientRect() {
	// #ifdef APP-PLUS-NVUE
	nextTick(function () {
		dom.getComponentRect(proxy.$refs.tmspin, function (res) {
			if (res?.size) {
				css_nvue_size.value = [res.size.width, res.size.height]
				if (res.size.height == 0) {
					nvuegetClientRect()
				}
			}
		})
	})
	// #endif
}
function clickMask(e) {
	// #ifdef APP-PLUS-NVUE
	e.stopPropagation()
	// #endif
	emits('click', e)
}
</script>

<style></style>
