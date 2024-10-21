<template>
	<view
		:style="[
			props.direction == 'horizontal'
				? {
						width: props.width + 'px',
						height: props.size + 'rpx',
						left: props.x + 'px',
						top: -props.size + 'rpx'
				  }
				: { height: props.width + 'px', width: props.size + 'rpx', top: props.x + 'px' }
		]"
		:class="['relative', props.direction == 'horizontal' ? 'flex flex-col' : 'flex flex-row absolute']"
	>
		<tm-sheet
			:followTheme="props.followTheme"
			:round="10"
			unit="px"
			:color="props.color"
			linear="right"
			:width="props.direction == 'horizontal' ? props.width : _sizePx"
			:height="props.direction == 'horizontal' ? _sizePx : props.width"
			:margin="[0, 0]"
			:padding="[0, 0]"
		></tm-sheet>
	</view>
</template>
<script lang="ts" setup>
import { computed } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
const props = defineProps({
	//是否跟随全局主题的变换而变换
	followTheme: {
		type: [Boolean, String],
		default: true
	},
	size: {
		type: Number,
		default: 6
	},
	x: {
		type: Number,
		dfault: 0
	},
	width: {
		type: Number,
		dfault: 0
	},
	color: {
		type: String,
		default: 'primary'
	},
	/**
	 * 方向
	 * horizontal:水平,
	 * vertical:竖向。
	 */
	direction: {
		type: String,
		default: 'vertical'
	}
})
const _sizePx = computed(() => uni.upx2px(props.size))
</script>
