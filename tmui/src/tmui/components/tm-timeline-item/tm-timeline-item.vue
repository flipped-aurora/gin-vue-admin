<template>
	<view class="flex flex-col flex-1 relative">
		<view
			class="itemPos flex flex-col relative"
			:class="[_position == 'right' ? 'flex-col-top-end' : '']"
			:style="[
				_position == 'left' ? { left: '64rpx' } : '',
				_position == 'right' ? { right: '64rpx' } : '',
				_position == 'center' ? { transform: 'translateX(50%)' } : '',
				{ top: '0rpx' }
			]"
		>
			<tm-sheet
				:border="props.border"
				:borderDirection="_position == 'center' ? 'left' : _position"
				:_class="['flex flex-col flex-col-top-start']"
				class="flex-1"
				parenClass="flex-1"
				:transprent="true"
				:margin="[0, 0]"
				:padding="[0, 0]"
			>
				<view
					class="pb-24 flex-1"
					:style="[
						_position == 'left' || _position == 'center' ? { paddingLeft: _posSizeWidth + 'rpx', marginRight: '132rpx' } : '',
						_position == 'right' ? { paddingRight: _posSizeWidth + 'rpx', marginLeft: '132rpx' } : ''
					]"
				>
					<tm-text v-if="_time" _style="line-height:20rpx;" _class="pb-16" :label="_time" :fontSize="24" :color="props.color"></tm-text>
					<slot></slot>
				</view>
			</tm-sheet>
		</view>
		<view
			v-if="_position == 'center'"
			:style="[
				{
					transform: `translateX(50%)`,
					height: _posSizeWidth + 'rpx',
					left: -_posX + 'px'
				}
			]"
			class=""
		>
			<view
				class="absolute t-0 flex flex-col flex-col-center-center"
				:style="[{ transform: `translateX(${_posX / 2}px)` }, { width: _posSizeWidth + 'rpx', height: _posSizeWidth + 'rpx' }]"
			>
				<tm-sheet
					_class="flex flex-center"
					:round="10"
					:shadow="2"
					:color="props.color"
					:width="_posSize"
					:height="_posSize"
					:margin="[0, 0]"
					:padding="[0, 0]"
				>
					<tm-icon v-if="_icon" :fontSize="20" :name="_icon"></tm-icon>
				</tm-sheet>
			</view>
		</view>
		<view
			v-if="_position == 'left' || _position == 'right'"
			class="absolute t-0 flex flex-col"
			:class="[
				_position == 'left' ? 'l-0 flex-col-top-center' : '',
				_position == 'right' ? 'r-0 flex-col-top-center' : '',
				_position == 'center' ? 'flex-col-center-center' : ''
			]"
			:style="[
				_position == 'left' || _position == 'right' ? { transform: `translateX(${_position == 'left' ? _posX : -_posX}px)` } : '',
				_position == 'center' ? { transform: `translateX(${_posX / 2}px)` } : '',
				{ width: _posSizeWidth + 'rpx', height: _posSizeWidth + 'rpx' }
			]"
		>
			<tm-sheet
				_class="flex flex-center"
				:round="10"
				:shadow="0"
				:border="props.type == 'outlined' ? 3 : 0"
				:text="props.type == 'outlined' ? true : false"
				:color="props.color"
				:width="_posSize"
				:height="_posSize"
				:margin="[0, 0]"
				:padding="[0, 0]"
			>
				<tm-icon v-if="_icon" :fontSize="20" :name="_icon"></tm-icon>
			</tm-sheet>
		</view>
	</view>
</template>
<script lang="ts" setup>
/**
 * 时间轴-子组件
 * @description 内部只可放置在 tm-timeline组件中，不可单独使用。
 */
import { computed, inject } from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
const props = defineProps({
	size: {
		type: Number,
		default: 24
	},
	color: {
		type: String,
		default: 'primary'
	},
	icon: {
		type: String,
		default: ''
	},
	border: {
		type: Number,
		default: 2
	},
	//outlined,fill
	type: {
		type: String,
		default: 'outlined'
	},
	//时间，不提供默认不显示。
	time: {
		type: String,
		default: ''
	}
})
const _time = computed(() => props.time)
const _position = inject(
	'tmTimeLinePosition',
	computed(() => 'left')
)

const _posSize = computed(() => {
	return props.size
})
const _posSizeWidth = computed(() => {
	return props.size + props.size / 2
})

const _posX = computed(() => {
	let l = uni.upx2px(64 - _posSizeWidth.value / 2 + props.border)
	return l
})

const _icon = computed(() => props.icon)
</script>
<style scoped>
.itemPos {
	/* #ifndef APP-NVUE */
	min-height: 80rpx;
	/* #endif */
}
</style>
