<template>
	<!-- 底部显示的标签 -->
	<view
		class="flex flex-between"
		:class="[props.direction == 'vertical' ? 'flex-col' : 'flex-row']"
		:style="[props.direction == 'horizontal' ? { width: width + 'rpx' } : { height: height + 'rpx' }]"
	>
		<view
			v-if="props.direction == 'horizontal'"
			class="flex flex-row"
			style="width: 80rpx"
			:class="[
				index == 0 ? 'flex-row-top-start' : '',
				index == _stepArray.length - 1 ? 'flex-row-top-end' : '',
				index < _stepArray.length - 1 && index > 0 ? 'flex-row-top-center' : ''
			]"
			v-for="(item, index) in _stepArray"
			:key="index"
		>
			<tmText :font-size="24" :label="item"></tmText>
		</view>
		<view
			v-if="props.direction != 'horizontal'"
			class="flex flex-row flex-row-center-start"
			style="width: 80rpx"
			v-for="(item, index) in _stepArray"
			:key="index"
		>
			<tmText :font-size="24" :label="item"></tmText>
		</view>
	</view>
</template>
<script lang="ts" setup>
import tmText from '../tm-text/tm-text.vue'
import { computed } from 'vue'
const props = defineProps({
	/**
	 * 方向
	 * horizontal:水平,
	 * vertical:竖向。
	 */
	direction: {
		type: String,
		default: 'horizontal'
	},
	width: {
		type: Number,
		default: 0
	},
	height: {
		type: Number,
		default: 0
	},
	max: {
		type: Number,
		default: 100
	},
	min: {
		type: Number,
		default: 0
	},
	step: {
		type: Number,
		default: 0
	},
	size: {
		type: Number,
		default: 32
	}
})

const _stepArray = computed(() => {
	let label = []
	if (props.step == 0) return []
	// if(props.step==1) return [props.min,props.max]
	let _step = props.max / props.step
	for (let i = 1; i <= props.step; i++) {
		label.push(i * _step)
	}
	return [props.min, ...label]
})
</script>
