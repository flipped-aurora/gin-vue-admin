<template>
	<view @click="onClick">
		<view
			:userInteractionEnabled="false"
			v-if="!_label && props.vertical"
			:style="[
				{ backgroundColor: _realColor ? tmcomputed.color : tmcomputed.border },
				props.vertical ? { width: props.border + 'rpx', height: props.height + 'rpx' } : ''
			]"
			:class="[props.vertical ? `mx-${props.margin[0]}` : `my-${props.margin[1]}`]"
		></view>
		<view :userInteractionEnabled="false" v-if="_label && !props.vertical" class="flex flex-row flex-center">
			<view
				:style="[
					tmcomputed
						? {
								backgroundColor: _realColor ? tmcomputed.color : tmcomputed.border,
								height: props.border + 'rpx'
						  }
						: ''
				]"
				:class="[
					`my-${props.margin[1]}`,
					align == 'left' ? 'flex-2' : '',
					align == 'right' ? 'flex-10' : '',
					align == 'center' ? 'flex-1' : ''
				]"
			>
			</view>
			<view v-if="_label" :class="[isDark ? 'opacity-4' : '']">
				<slot name="label">
					<tm-text
						:fontSize="props.fontSize"
						:dark="isDark"
						:followTheme="props.followTheme"
						:color="props.fontColor"
						:label="props.label"
						:_class="['mx-32']"
					></tm-text>
				</slot>
			</view>
			<view
				:style="[
					tmcomputed
						? {
								backgroundColor: _realColor ? tmcomputed.color : tmcomputed.border,
								height: props.border + 'rpx'
						  }
						: ''
				]"
				:class="[
					`my-${props.margin[1]}`,
					align == 'left' ? 'flex-10' : '',
					align == 'right' ? 'flex-2' : '',
					align == 'center' ? 'flex-1' : ''
				]"
			>
			</view>
		</view>
		<view :userInteractionEnabled="false" v-if="!_label && !props.vertical" class="flex flex-row flex-center">
			<view
				class="flex-1"
				:class="[`my-${props.margin[1]}`]"
				:style="[
					tmcomputed
						? {
								backgroundColor: _realColor ? tmcomputed.color : tmcomputed.border,
								height: props.border + 'rpx'
						  }
						: ''
				]"
			>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 分割线
 * @description 分割线，带文本标签，提供左，中，右文本标签。
 */
import { getCurrentInstance, computed, ref, provide, inject } from 'vue'
import { tmVuetify } from '../../tool/lib/interface'
import { custom_props, computedTheme, computedDark } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import tmText from '../tm-text/tm-text.vue'
const store = useTmpiniaStore()
// 混淆props共有参数
const props = defineProps({
	...custom_props,
	color: {
		type: String,
		default: 'grey-3'
	},
	fontColor: {
		type: String,
		default: 'grey-1'
	},
	fontSize: {
		type: Number,
		default: 26
	},
	vertical: {
		type: [Boolean],
		default: false
	},
	height: {
		type: [Number, String],
		default: 26
	},
	label: {
		type: String,
		default: ''
	},
	showLabel: {
		type: Boolean,
		default: false
	},
	align: {
		type: String,
		default: 'center' //left,right,center
	},
	margin: {
		type: Array,
		default: () => [16, 24]
	},
	border: {
		type: [Number],
		default: 1
	},
	//使用原始颜色为线条色，而不使用计算过的颜色值。
	realColor: {
		type: [Boolean],
		default: false
	}
})
const emits = defineEmits(['click'])
//线的方向。
const borderDir = computed(() => (props.vertical ? 'left' : 'bottom'))
const _label = computed(() => props.label || props.showLabel)
// 设置响应式全局组件库配置表。
const tmcfg = computed<tmVuetify>(() => store.tmStore)
const _realColor = computed(() => props.realColor)
const isDark = computed(() =>
	computedDark(
		{
			...props,
			borderDirection: borderDir.value
		},
		tmcfg.value
	)
)
//计算主题
const tmcomputed = computed(() =>
	computedTheme(
		{
			...props,
			borderDirection: borderDir.value
		},
		isDark.value,
		tmcfg.value
	)
)
function onClick(e: any) {
	emits('click', e)
}
</script>

<style></style>
