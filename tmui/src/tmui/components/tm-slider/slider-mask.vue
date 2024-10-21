<template>
	<!-- 底部显示的标签 -->
	<view
		class="flex flex-between absolute"
		:class="[props.direction == 'vertical' ? 'flex-col ' : 'flex-row']"
		:style="[
			props.direction == 'horizontal'
				? { width: width + 'rpx', top: (props.size - 12) / 2 - 2 + 'rpx' }
				: { height: height + 'rpx', top: '0px', left: (props.size - 12) / 2 + 2 + 'rpx' }
		]"
	>
		<view
			v-if="props.direction == 'horizontal'"
			class="flex flex-row relative"
			style="width: 80rpx"
			:class="[
				index == 0 ? 'flex-row-top-start' : '',
				index == _stepArray.length - 1 ? 'flex-row-top-end' : '',
				index < _stepArray.length - 1 && index > 0 ? 'flex-row-top-center' : ''
			]"
			v-for="(item, index) in _stepArray"
			:key="index"
		>
			<TmSheet1
				:followTheme="props.followTheme"
				:color="props.color"
				:round="6"
				:margin="[0, 0]"
				:padding="[0, 0]"
				:width="12"
				:height="12"
				text
				:border="1"
			></TmSheet1>
		</view>
		<view
			v-if="props.direction != 'horizontal'"
			:translate="true"
			class="flex flex-row flex-row-center-start relative"
			style="width: 80rpx"
			v-for="(item, index) in _stepArray"
			:key="index"
		>
			<TmSheet1
				:followTheme="props.followTheme"
				:round="6"
				:margin="[0, 0]"
				:padding="[0, 0]"
				:width="12"
				:height="12"
				color="primary"
				text
				:border="1"
			></TmSheet1>
		</view>
	</view>
</template>
<script lang="ts" setup>
import tmText from '../tm-text/tm-text.vue'
import { computed } from 'vue'
import TmSheet1 from '../tm-sheet/tm-sheet.vue'
const props = defineProps({
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
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
	},
	color: {
		type: String,
		default: 'primary'
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
