<template>
	<view class="flex flex-col">
		<view v-if="props.model == 'line'" class="ma-32">
			<view v-for="item in props.rows" :key="item" class="flex flex-col">
				<tm-skeleton-line :height="props.height"></tm-skeleton-line>
			</view>
			<view class="flex flex-row">
				<view class="flex-2">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
				<view class="flex-2 mx-24">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
				<view class="flex-2 mr-24">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
				<view class="flex-4">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
			</view>
		</view>

		<view v-if="props.model == 'rect'" class="ma-32">
			<view v-for="item in props.rows" :key="item" class="flex flex-row">
				<view class="flex-2">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
				<view class="flex-4 mx-24">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
				<view class="flex-2 mr-24">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
				<view class="flex-2">
					<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
				</view>
			</view>
		</view>

		<view v-if="props.model == 'card'" class="ma-32">
			<view class="flex flex-row">
				<view class="flex-1">
					<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				</view>
				<view class="flex-10 mx-24">
					<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				</view>
				<view class="flex-1">
					<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				</view>
			</view>
			<view class="">
				<tm-skeleton-line :height="props.height * 4"></tm-skeleton-line>
			</view>
			<view class="flex flex-row">
				<view class="flex-2">
					<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				</view>
				<view class="flex-4 mx-24">
					<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				</view>
				<view class="flex-2 mr-24">
					<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				</view>
				<view class="flex-2">
					<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				</view>
			</view>
		</view>
		<view v-if="props.model == 'chat'" class="flex flex-row ma-32">
			<view class="flex-2" :style="[{ height: props.height * 2 + 'rpx' }]">
				<tm-skeleton-line :height="props.height * 2"></tm-skeleton-line>
			</view>
			<view class="flex-8 mx-24">
				<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				<tm-skeleton-line :height="props.height"></tm-skeleton-line>
				<tm-skeleton-line :height="props.height"></tm-skeleton-line>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 骨架屏
 * @description 骨架加载状态屏，用来展示数据前的一种加载状态。预设了：line ,rect,card,chat四种类型，如果想要更好的自定义请使用tm-skeleton-line组件自行组合。
 */
import tmSkeletonLine from '../tm-skeleton-line/tm-skeleton-line.vue'
const props = defineProps({
	height: {
		type: Number,
		default: 60
	},
	rows: {
		type: Number,
		default: 3
	},
	model: {
		type: String,
		default: 'line' //line ,rect,card,chat,
	}
})
</script>

<style></style>
