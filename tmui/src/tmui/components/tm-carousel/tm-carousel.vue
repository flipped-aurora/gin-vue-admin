<template>
	<view class="flex flex-col flex-col-center-center">
		<view
			class="relative overflow"
			:class="[`round-${props.round} mx-${props.margin[0]} my-${props.margin[1]}`]"
			:style="[{ width: `${props.width}rpx`, height: `${props.height}rpx` }]"
		>
			<swiper
				:indicator-dots="false"
				:interval="props.interval"
				:duration="props.duration"
				:circular="props.circular"
				:vertical="props.vertical"
				:acceleration="props.acceleration"
				:disableProgrammaticAnimation="props.disableProgrammaticAnimation"
				:autoplay="_autoplay"
				:displayMultipleItems="props.displayMultipleItems"
				:skipHiddenItemLayout="props.skipHiddenItemLayout"
				:disableTouch="props.disableTouch"
				:touchable="props.touchable"
				@change="sliderChange"
				:class="`round-${props.round} overflow`"
				:current="_current"
				:style="[{ width: `${props.width}rpx`, height: `${props.height}rpx` }]"
			>
				<swiper-item
					:class="`round-${props.round} overflow `"
					@click="emits('click', index)"
					v-for="(item, index) in _list"
					:key="index"
					:style="[{ width: `${props.width}rpx`, height: `${props.height}rpx` }]"
				>
					<tm-image
						:model="props.imgModel"
						:round="props.round"
						:userInteractionEnabled="false"
						:showLoad="props.showLoad"
						v-if="item.type == listItemType.img"
						:src="item.url"
						:width="props.width"
						:height="props.height"
					></tm-image>
					<tm-image
						:round="props.round"
						:userInteractionEnabled="false"
						:showLoad="props.showLoad"
						v-if="item.type == listItemType.video && item.img && _currentActive != index"
						:src="item.img"
						:width="props.width"
						:height="props.height"
					></tm-image>
					<video
						:userInteractionEnabled="false"
						id="video"
						v-if="item.type == listItemType.video && _currentActive === index"
						:src="item.url"
						:style="[{ width: `${props.width}rpx`, height: `${props.height}rpx` }]"
						:autoplay="_currentActive === index"
						:class="`round-${props.round}`"
					></video>
					
					<!-- 底部文字提示语 -->
					<view v-if="item?.text"  class="absolute l-0 b-0 px-24" :style="[{ width: `${props.width}rpx`, height: `64rpx`,background:'rgba(0,0,0,0.4)','box-sizing':' border-box' }]">
						<text class="text-size-m text-white text-overflow-1" style="line-height: 64rpx;">{{item.text}}</text>
					</view>
				</swiper-item>
			</swiper>
			<!-- dot -->
			<view
				class="absolute"
				v-if="_model == 'dot' && props.indicatorDots"
				:class="[
					_dotPosition == 'bottom' || _dotPosition == 'top' ? 'flex flex-row ' : '',
					_dotPosition == 'left' || _dotPosition == 'right' ? 'flex flex-col ' : '',
					_align == 'center' ? 'flex-center' : '',
					_align == 'left' ? 'flex-row-center-start ' : '',
					_align == 'right' ? 'flex-row-center-end ' : '',
					_dotPosition == 'left' && _align == 'left' ? 'ml-12 mt-24' : '',
					_dotPosition == 'left' && _align == 'right' ? 'ml-12 pb-24' : '',
					_dotPosition == 'right' && _align == 'left' ? 'pr-12 mt-24' : '',
					_dotPosition == 'right' && _align == 'right' ? 'pr-12 pb-24' : '',

					_dotPosition == 'bottom' && _align == 'left' ? 'ml-12 mb-0' : '',
					_dotPosition == 'bottom' && _align == 'right' ? 'pr-12 mb-0' : '',
					_dotPosition == 'top' && _align == 'left' ? 'ml-12 ' : '',
					_dotPosition == 'top' && _align == 'right' ? 'pr-12 ' : ''
				]"
				:style="[
					_dotPosition == 'bottom'
						? {
								left: '0px',
								bottom: '0px',
								width: `${props.width}rpx`,
								height: `${60}rpx`
						  }
						: '',
					_dotPosition == 'top' ? { left: '0px', top: '0px', width: `${props.width}rpx`, height: `${60}rpx` } : '',
					_dotPosition == 'left' ? { left: '0px', top: '0px', width: `${60}rpx`, height: `${props.height}rpx` } : '',
					_dotPosition == 'right'
						? {
								right: '0px',
								top: '0px',
								width: `${60}rpx`,
								height: `${props.height}rpx`
						  }
						: ''
				]"
			>
				<tm-sheet
					:margin="[10, 10]"
					:follow-theme="_currentActive == index ? props.followTheme : false"
					:padding="[0, 0]"
					:round="10"
					@click="dotClick(index)"
					:color="_currentActive == index ? props.color : 'white'"
					v-for="(item, index) in _list"
					:key="index"
					:width="18"
					:height="18"
				></tm-sheet>
			</view>
			<!-- rect -->
			<view
				class="absolute"
				v-if="_model == 'rect' && props.indicatorDots"
				:class="[
					_dotPosition == 'bottom' || _dotPosition == 'top' ? 'flex flex-row ' : '',
					_dotPosition == 'left' || _dotPosition == 'right' ? 'flex flex-col ' : '',
					_align == 'center' ? 'flex-center' : '',
					_align == 'left' ? 'flex-row-center-start ' : '',
					_align == 'right' ? 'flex-row-center-end ' : '',
					_dotPosition == 'left' && _align == 'left' ? 'ml-12 mt-24' : '',
					_dotPosition == 'left' && _align == 'right' ? 'ml-12 pb-24' : '',
					_dotPosition == 'right' && _align == 'left' ? 'pr-12 mt-24' : '',
					_dotPosition == 'right' && _align == 'right' ? 'pr-12 pb-24' : '',

					_dotPosition == 'bottom' && _align == 'left' ? 'ml-12 mb-0' : '',
					_dotPosition == 'bottom' && _align == 'right' ? 'pr-12 mb-0' : '',
					_dotPosition == 'top' && _align == 'left' ? 'ml-12 ' : '',
					_dotPosition == 'top' && _align == 'right' ? 'pr-12 ' : ''
				]"
				:style="[
					_dotPosition == 'bottom'
						? {
								left: '0px',
								bottom: '0px',
								width: `${props.width}rpx`,
								height: `${60}rpx`
						  }
						: '',
					_dotPosition == 'top' ? { left: '0px', top: '0px', width: `${props.width}rpx`, height: `${60}rpx` } : '',
					_dotPosition == 'left' ? { left: '0px', top: '0px', width: `${60}rpx`, height: `${props.height}rpx` } : '',
					_dotPosition == 'right'
						? {
								right: '0px',
								top: '0px',
								width: `${60}rpx`,
								height: `${props.height}rpx`
						  }
						: ''
				]"
			>
				<tm-sheet
					:round="index == 0 || index == _list.length - 1 ? 10 : 0"
					:margin="_dotPosition == 'left' || _dotPosition == 'right' ? [10, 0] : [0, 10]"
					:padding="[0, 0]"
					@click="dotClick(index)"
					:follow-theme="_currentActive == index ? props.followTheme : false"
					:color="_currentActive == index ? props.color : props.indicatorColor"
					v-for="(item, index) in _list"
					:key="index"
					:width="_dotPosition == 'left' || _dotPosition == 'right' ? 6 : 36"
					:height="_dotPosition == 'left' || _dotPosition == 'right' ? 36 : 6"
				></tm-sheet>
			</view>
			<!-- number -->
			<view
				class="absolute"
				v-if="_model == 'number' && props.indicatorDots"
				:class="[
					_dotPosition == 'bottom' || _dotPosition == 'top' ? 'flex flex-row ' : '',
					_dotPosition == 'left' || _dotPosition == 'right' ? 'flex flex-col ' : '',
					_align == 'center' ? 'flex-center' : '',
					_align == 'left' ? 'flex-row-center-start ' : '',
					_align == 'right' ? ' flex-row-center-end ' : '',

					_dotPosition == 'left' && _align == 'left' ? 'ml-12 mt-24' : '',
					_dotPosition == 'left' && _align == 'right' ? 'ml-12 pb-24' : '',
					_dotPosition == 'right' && _align == 'left' ? 'pr-12 mt-24' : '',
					_dotPosition == 'right' && _align == 'right' ? 'pr-12 pb-24' : '',

					_dotPosition == 'bottom' && _align == 'left' ? 'ml-12 mb-0' : '',
					_dotPosition == 'bottom' && _align == 'right' ? 'pr-12 mb-0 ' : '',
					_dotPosition == 'top' && _align == 'left' ? 'ml-12 ' : '',
					_dotPosition == 'top' && _align == 'right' ? 'pr-12 ' : ''
				]"
				:style="[
					_dotPosition == 'bottom'
						? {
								left: '0px',
								bottom: '0px',
								width: `${props.width}rpx`,
								height: `${60}rpx`
						  }
						: '',
					_dotPosition == 'top' ? { left: '0px', top: '0px', width: `${props.width}rpx`, height: `${60}rpx` } : '',
					_dotPosition == 'left' ? { left: '0px', top: '0px', width: `${60}rpx`, height: `${props.height}rpx` } : '',
					_dotPosition == 'right'
						? {
								right: '0px',
								top: '0px',
								width: `${60}rpx`,
								height: `${props.height}rpx`
						  }
						: ''
				]"
			>
				<view
					class="round-10"
					:class="[
						_dotPosition == 'left' || _dotPosition == 'right' ? 'px-5 py-24 ' : '',
						_dotPosition == 'bottom' || _dotPosition == 'top' ? 'px-24 py-5 mx-32' : ''
					]"
					style="background-color: rgba(0, 0, 0, 0.4); font-size: 0px"
				>
					<text style="font-size: 22rpx; color: white">{{ _currentActive + 1 }}/{{ _list.length }}</text>
				</view>
			</view>
		</view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 图片轮播
 * @description 可以嵌入视频，图片。
 */
import { ref, PropType, computed, getCurrentInstance, nextTick } from 'vue'

import tmImage from '../tm-image/tm-image.vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import { listItem, listItemType, listItemTypeStr } from './interface'
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['change', 'click'])
const props = defineProps({
	imgModel: {
		type: String,
		default: 'scaleToFill'
	},
	followTheme: {
		type: Boolean,
		default: true
	},
	color: {
		type: String,
		default: 'primary'
	},
	indicatorColor: {
		type: String,
		default: 'white'
	},
	width: {
		type: Number,
		default: 750
	},
	height: {
		type: Number,
		default: 500
	},
	round: {
		type: Number,
		default: 0
	},
	margin: {
		type: Array as PropType<Array<number>>,
		default: () => [0, 0]
	},
	/**
	 * 轮播图的数据列表
	 * 可以是图片地址数组，也可以是对象数组（需要提供rangKey字段，默认url）
	 */
	list: {
		type: Array as PropType<Array<string | listItem>>,
		default: () => []
	},
	/**
	 * 当list为对象时，需要提供key名称以获取图片地址。
	 */
	rangKey: {
		type: String,
		default: 'url'
	},
	defaultValue: {
		type: Number,
		default: 0
	},
	/**
	 * 指示点的位置
	 * top,left,right,bottom
	 */
	dotPosition: {
		type: String as PropType<'top' | 'left' | 'right' | 'bottom'>,
		default: 'bottom'
	},
	/**
	 * 内容居对齐，左，中，右
	 * left,center,right
	 **/
	align: {
		type: String as PropType<'left' | 'right' | 'center'>,
		default: 'center'
	},
	/**
	 * 指示点类型 dot,number,rect
	 */
	model: {
		type: String as PropType<'dot' | 'number' | 'rect'>,
		default: 'number'
	},
	/**间隔时间 */
	interval: {
		type: Number,
		default: 5000
	},
	/**动画时间 */
	duration: {
		type: Number,
		default: 500
	},
	circular: {
		type: Boolean,
		default: true
	},
	/**是否垂直轮播 */
	vertical: {
		type: Boolean,
		default: false
	},
	acceleration: {
		type: Boolean,
		default: false
	},
	disableProgrammaticAnimation: {
		type: Boolean,
		default: false
	},
	autoplay: {
		type: Boolean,
		default: true
	},
	displayMultipleItems: {
		type: Number,
		default: 1
	},
	skipHiddenItemLayout: {
		type: Boolean,
		default: false
	},
	disableTouch: {
		type: Boolean,
		default: false
	},
	touchable: {
		type: Boolean,
		default: false
	},
	indicatorDots: {
		type: Boolean,
		default: true
	},
	//是否显示加载动画
	showLoad: {
		type: Boolean,
		default: true
	}
})
const _list = computed(() => {
	let l: Array<listItem> = []
	props.list.forEach((el) => {
		if (typeof el == 'string') {
			l.push({
				url: el,
				type: listItemType.img,
				text:"",
			})
		} else if (typeof el === 'object') {
			l.push({
				url: el[props.rangKey],
				type: el?.type ?? listItemType.img,
				text:el?.text??"",
				img: el?.img ?? '',
				...el
			})
		}
	})
	return l
})

const _current = ref(props.defaultValue || 0)
const _currentActive = ref(props.defaultValue || 0)
const _model = computed(() => props.model)
const _dotPosition = computed(() => props.dotPosition)
const _align = computed(() => props.align)
const _autoplay = computed(() => props.autoplay)
function sliderChange(e: any) {
	if (!_autoplay.value) {
		_current.value = e?.detail?.current
	}
	_currentActive.value = e?.detail?.current
	nextTick(() => {
		emits('change', _currentActive.value)
	})
}
function dotClick(index: number) {
	_currentActive.value = index
	if (!_autoplay.value) {
		_current.value = index
	}
}
</script>

<style></style>
