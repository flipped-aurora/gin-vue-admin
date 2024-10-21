<template>
	<view
		class="relative flex"
		:style="{
			width: props.width + 'rpx',
			height: props.height + (_showBar ? 32 : 0) + 'rpx'
		}"
	>
		<!-- #ifdef APP-NVUE -->
		<scroll-view
			@scroll="onScroll"
			:scroll-left="scrollLeft"
			:show-scrollbar="false"
			:scroll-x="true"
			class="flex flex-row flex-nowrap nowrap"
			:style="{ width: props.width + 'rpx', height: props.height + 'rpx' }"
		>
			<slot name="default"></slot>
		</scroll-view>
		<!-- #endif -->
		<!-- #ifndef APP-NVUE -->
		<scroll-view
			@scroll="onScroll"
			:scroll-left="scrollLeft"
			:scroll-x="true"
			class="flex flex-row flex-nowrap nowrap"
			:style="{ width: props.width + 'rpx', height: props.height + 'rpx' }"
		>
			<view class="flex-1 flex-row flex-nowrap">
				<slot name="default"></slot>
			</view>
		</scroll-view>
		<!-- #endif -->
		<view
			v-if="_showBar"
			class="absolute l-0 b-10 flex"
			:class="alignKey[props.align]"
			:style="{ width: props.width + 'rpx', height: 32 + 'rpx' }"
		>
			<view v-if="props.align == 'between'" class="flex-1 pr-40" style="width: 0px">
				<slot name="barll"></slot>
			</view>
			<view style="width: 100rpx" :class="[props.align == 'between' ? 'mr-24' : '']">
				<tm-sheet no-level :round="6" :width="100" :height="8" :color="props.bgColor" :margin="[0, 0]" :padding="[0, 0]">
					<view
						class="bar"
						:style="{
							transform: `translateX(${left}px)`
						}"
					>
						<tm-sheet :round="6" :width="50" :height="8" :color="props.color" :margin="[0, 0]" :padding="[0, 0]"></tm-sheet>
					</view>
				</tm-sheet>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
import { ref, computed, PropType } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
//居中，两边对齐。
type alignType = 'center' | 'between'
const emits = defineEmits(['change'])

const props = defineProps({
	width: {
		type: Number,
		default: 636
	},
	height: {
		type: Number,
		default: 300
	},
	/**
	 * 是否显示底部导航条
	 */
	showBar: {
		type: Boolean,
		default: true
	},
	bgColor: {
		type: String,
		default: 'grey-2'
	},
	color: {
		type: String,
		default: 'primary'
	},
	align: {
		type: String as PropType<alignType>,
		default: 'center'
	}
})
const left = ref(0)
const totalWidth = uni.upx2px(50)
const totalContWidth = uni.upx2px(props.width)
const _showBar = computed(() => props.showBar)
const alignKey = {
	center: 'flex-center',
	between: 'flex-row flex-row-center-between'
}
const onScroll = (e: Event) => {
	emits('change', e.detail)
	if (!_showBar.value) return
	let sL = e.detail.scrollLeft
	let sT = e.detail.scrollWidth
	let maxLeft = Math.abs(sT - totalContWidth)
	let nowLeft = (sL / maxLeft) * totalWidth
	if (sL <= 0) nowLeft = 0
	if (Math.abs(sL) >= maxLeft) {
		nowLeft = totalWidth
	}
	left.value = nowLeft
}
const scrollLeft = ref(0)
const scrollTo = (value: number) => {
	scrollLeft.value = value
}
defineExpose({ scrollTo })
</script>

<style scoped>
.bar {
	transition-property: transform;
	transition-delay: 0;
	transition-duration: 0.05s;
	transition-timing-function: linear;
	transform: translateX(0px);
}

/* #ifndef APP-NVUE */

::-webkit-scrollbar,
scroll-view::-webkit-scrollbar {
	display: none;
	width: 0;
	height: 0;
	color: transparent;
}

/* #endif */
</style>
