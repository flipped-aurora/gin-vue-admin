<template>
	<view
		class="flex flex-col overflow"
		:style="[props.height && !isDulitabs && cacheTabs.length > 0 ? { height: height + 'rpx' } : '', { width: props.width + 'rpx' }]"
	>
		<!-- 此源码有uniapp bug.如果在nvue页面编译至h5平台时，开启enable-flexr后需要里面再套层view再写flex才能真正的开flex -->
		<!-- 因此下面的内容作了条件编译分为nvue和非nvue -->
		<!-- https://ask.dcloud.net.cn/question/143230 -->
		<!-- #ifndef APP-NVUE -->
		<view
			@touchmove="onMove"
			@touchend="onEnd"
			@touchstart="onStart"
			@mousemove="onMove"
			@mouseup="onEnd"
			@mouseleave="onEnd"
			@mousedown="onStart"
			ref="tabsDom"
			:style="{
				width: props.swiper ? `${totalWidth}px` : `${props.width}rpx`,
				transform: props.swiper ? `translateX(${directoStyle}px)` : `translateX(0px)`
			}"
			v-if="_tabPos == 'bottom' && !isDulitabs"
			:class="[!isEndMove || isNvue ? 'tmTabsPane' : '']"
			class="flex flex-row flex-nowrap overflow"
		>
			<slot></slot>
		</view>
		<!-- #endif -->
		<!-- #ifdef APP-NVUE -->
		<!-- @touchmove="onMove"
		@touchend="onEnd" 
		@touchstart="onStart" -->
		<view
			@touchstart="spinNvueAni"
			ref="tabsDom"
			:style="{
				width: props.swiper ? `${totalWidth}px` : `${props.width}rpx`,
				transform: `translateX(0px)`
			}"
			v-if="_tabPos == 'bottom' && !isDulitabs && props.swiper"
			class="flex flex-row flex-nowrap overflow"
		>
			<slot></slot>
		</view>
		<!-- #endif -->

		<tm-sheet
			:darkBgColor="props.darkBgColor"
			:transprent="props.transprent"
			:color="props.color"
			:followTheme="props.followTheme"
			:dark="props.dark"
			:round="props.round"
			:shadow="props.shadow"
			:outlined="props.outlined"
			:border="props.border"
			:borderStyle="props.borderStyle"
			:borderDirection="props.borderDirection"
			:text="props.text"
			:linear="props.linear"
			:linearDeep="props.linearDeep"
			:margin="[0, 0]"
			:padding="[0, 0]"
			:height="props.itemHeight + modelStyle.border + props.gutter + 4"
			:_class="['flex-center flex-row nonvue', cstomClass]"
			:_style="props._style"
			:width="props.width"
		>
			<!-- 标题 -->
			<!-- #ifndef APP-NVUE -->
			<scroll-view
				:style="[{ width: `${props.width}rpx`, height: `${props.itemHeight + 4}rpx` }]"
				:scroll-with-animation="true"
				:scroll-into-view="_scrollToId"
				:scroll-x="true"
				:show-scrollbar="false"
				enable-flex
				class="tableHeader flex-row relative"
			>
				<view class="flex flex-row nowrap nonvue fulled" :class="[_align]" :style="[{ height: `${props.itemHeight + 4}rpx` }]">
					<view :id="tabsid + item.key" v-for="(item, index) in cacheTabs" :key="index" class="flex flex-shrink">
						<tm-sheet
							@click="changeKey(item.key)"
							:round="props.itemRound"
							:linear="props.itemLinear"
							:linearDeep="props.itemLinearDeep"
							borderDirection="bottom"
							:text="item.key === _active ? modelStyle.text : false"
							:border="item.key === _active ? modelStyle.border : 0"
							:transprent="item.key === _active ? modelStyle.transprent : true"
							:color="props.activeColor && item.key === _active ? props.activeColor : props.color"
							:width="props.itemWidth"
							_class="flex-col flex-col-center-center"
							:margin="[0, 0]"
							:padding="[0, 0]"
							:height="props.itemHeight"
						>
							<view
								:style="[props.itemWidth > 0 ? { width: (props.itemWidth) + 'rpx' } : {paddingLeft:'20rpx',paddingRight:'20rpx'}, { height: props.itemHeight + 'rpx' }]"
								class="flex flex-row flex-row-center-center relative flex-shrink"
							>
								<view class="flex flex-row flex-center">
									<tm-icon
										:userInteractionEnabled="false"
										v-if="item.icon"
										_class="pr-5"
										:color="item.key === _active ? props.activeFontColor : props.unSelectedColor"
										:font-size="item.key === _active ? props.activeFontSize : props.itemFontSize"
										:name="item.icon"
									>
									</tm-icon>
									<tm-text
										:userInteractionEnabled="false"
										:font-size="item.key === _active ? props.activeFontSize : props.itemFontSize"
										:_class="item.key === _active ? 'text-weight-b' : ''"
										:color="item.key === _active ? props.activeFontColor : props.unSelectedColor"
										:label="item.title"
									>
									</tm-text>
								</view>
								<view
									:userInteractionEnabled="false"
									v-if="!item.count && item.dot"
									:style="{
										height: 12 + 'rpx',
										top: '12rpx',
										right: '-16rpx',
										width: isNvue ? 'flex-1' : '100%'
									}"
									class="absolute t-0 r-0 flex flex-row flex-row-center-end"
								>
									<tm-badge dot :color="item.dotColor"> </tm-badge>
								</view>
								<view
									:userInteractionEnabled="false"
									v-if="item.count && !item.dot"
									:style="{
										height: props.itemHeight - 30 + 'rpx',
										top: '10rpx',
										right: props.showTabsLineAni ? 0 : -16 + 'rpx',
										width: isNvue ? 'flex-1' : '100%'
									}"
									class="absolute t-0 r-0"
								>
									<tm-badge :count="item.count" :color="item.dotColor"> </tm-badge>
								</view>
							</view>
						</tm-sheet>
					</view>
				</view>
				
				<view
					v-if="props.showTabsLineAni && props.itemWidth > 0 && props.showTabsLine"
					class="anilineBar absolute l-0"
					:style="{
						width: `${contentWidth}rpx`,
						height: '1px',
						top: `${props.itemHeight}rpx`,
						backgroundColor: props.showTabsLineAni ? (store.tmStore.dark ? '#616161' : '#ebebeb') : ''
					}"
				></view>
				<view
					v-if="props.showTabsLineAni && props.itemWidth > 0"
					class="anilineBar absolute l-0 b-0 flex flex-row flex-center"
					:style="{
						transform: `translateX(${anitLineLeft}px)`,
						top: `${props.itemHeight - 2}rpx`,
						width: _itemwidth + 'px',
						height: '4px'
					}"
				>
					<tm-sheet
						parenClass="animateAll_tabs_tmui"
						:follow-dark="false"
						:color="props.tabsLineAniColor"
						:width="widthDrag"
						unit="px"
						:height="4"
						:margin="[0, 0]"
						:padding="[0, 0]"
					></tm-sheet>
				</view>
			</scroll-view>

			<!-- #endif -->
			<!-- 标题 -->
			<!-- #ifdef APP-NVUE -->

			<scroll-view
				v-if="!props.showTabsLineAni && props.itemWidth == 0"
				:scroll-into-view="_scrollToId"
				:scroll-x="true"
				:scroll-with-animation="true"
				:show-scrollbar="false"
				enable-flex
				class="flex-row"
				:class="[_align]"
				:style="[{ width: `${props.width}rpx`, height: `${props.itemHeight + 4}rpx` }]"
			>
				<view :id="tabsid + item.key" v-for="(item, index) in cacheTabs" :key="index" class="flex flex-shrink">
					<tm-sheet
						@click="changeKey(item.key)"
						:round="props.itemRound"
						:linear="props.itemLinear"
						:linearDeep="props.itemLinearDeep"
						borderDirection="bottom"
						:text="item.key === _active ? modelStyle.text : false"
						:border="item.key === _active ? modelStyle.border : 0"
						:transprent="item.key === _active ? modelStyle.transprent : true"
						:color="props.activeColor && item.key === _active ? props.activeColor : props.color"
						:width="props.itemWidth"
						_class="flex-center flex-row"
						:margin="[0, 0]"
						:padding="[0, 0]"
						:height="props.itemHeight"
						unit="rpx"
					>
						<tm-badge :font-size="19" :dot="item.dot" :count="item.count" :color="item.dotColor">
							<view class="flex flex-row flex-center px-20" :style="{ height: props.itemHeight - 20 + 'rpx' }">
								<tm-icon
									:userInteractionEnabled="false"
									v-if="item.icon"
									_class="pr-5"
									:color="item.key === _active ? props.activeFontColor : props.unSelectedColor"
									:font-size="item.key === _active ? props.activeFontSize : props.itemFontSize"
									:name="item.icon"
								>
								</tm-icon>
								<tm-text
									:userInteractionEnabled="false"
									:font-size="item.key === _active ? props.activeFontSize : props.itemFontSize"
									:color="item.key === _active ? props.activeFontColor : props.unSelectedColor"
									:label="item.title"
								>
								</tm-text>
							</view>
						</tm-badge>
					</tm-sheet>
				</view>
				<view
					v-if="props.showTabsLineAni && props.itemWidth > 0 && props.showTabsLine"
					class="absolute l-0 b-0"
					:style="{
						width: `${contentWidth}rpx`,
						height: '1px',
						backgroundColor: props.showTabsLineAni ? (store.tmStore.dark ? '#616161' : '#ebebeb') : ''
					}"
				></view>
				<view
					v-if="props.showTabsLineAni && props.itemWidth > 0"
					class="anilineBar absolute l-0 b-0"
					:style="{ transform: `translateX(${anitLineLeft}px)` }"
				>
					<tm-sheet
						:round="10"
						:follow-dark="false"
						:width="40"
						:color="props.tabsLineAniColor"
						:height="8"
						:margin="[0, 0]"
						:padding="[0, 0]"
					></tm-sheet>
				</view>
			</scroll-view>

			<scroll-view
				v-if="props.showTabsLineAni && props.itemWidth > 0"
				:scroll-into-view="_scrollToId"
				:scroll-x="true"
				:scroll-with-animation="true"
				:show-scrollbar="false"
				enable-flex
				class="flex-row"
				:class="[_align]"
				:style="[{ width: `${_width}px`, height: `${props.itemHeight + 4}rpx` }]"
			>
				<view :id="tabsid + item.key" v-for="(item, index) in cacheTabs" :key="index">
					<tm-sheet
						@click="changeKey(item.key)"
						:round="props.itemRound"
						:linear="props.itemLinear"
						:linearDeep="props.itemLinearDeep"
						borderDirection="bottom"
						:text="item.key === _active ? modelStyle.text : false"
						:border="item.key === _active ? modelStyle.border : 0"
						:transprent="item.key === _active ? modelStyle.transprent : true"
						:color="props.activeColor && item.key === _active ? props.activeColor : props.color"
						:width="_itemwidth"
						_class="flex-center flex-row"
						:margin="[0, 0]"
						:padding="[0, 0]"
						:height="_itemheight"
						unit="px"
					>
						<tm-badge :font-size="19" :dot="item.dot" :count="item.count" :color="item.dotColor">
							<view class="flex flex-row flex-center px-20" :style="{ height: props.itemHeight - 20 + 'rpx' }">
								<tm-icon
									:userInteractionEnabled="false"
									v-if="item.icon"
									_class="pr-5"
									:color="item.key === _active ? props.activeFontColor : props.unSelectedColor"
									:font-size="item.key === _active ? props.activeFontSize : props.itemFontSize"
									:name="item.icon"
								>
								</tm-icon>
								<tm-text
									:userInteractionEnabled="false"
									:font-size="item.key === _active ? props.activeFontSize : props.itemFontSize"
									:color="item.key === _active ? props.activeFontColor : props.unSelectedColor"
									:label="item.title"
								>
								</tm-text>
							</view>
						</tm-badge>
					</tm-sheet>
				</view>
				<view
					v-if="props.showTabsLineAni && props.itemWidth > 0 && props.showTabsLine"
					class="absolute l-0 b-0"
					:style="{
						width: `${contentWidth}rpx`,
						height: '1px',
						backgroundColor: props.showTabsLineAni ? (store.tmStore.dark ? '#616161' : '#ebebeb') : ''
					}"
				></view>
				<view
					v-if="props.showTabsLineAni && props.itemWidth > 0"
					class="anilineBar absolute l-0 b-0 flex flex-row flex-center"
					:style="{
						transform: `translateX(${anitLineLeft}px)`,
						width: _itemwidth + 'px',
						height: '4px'
					}"
				>
					<tm-sheet
						ref="sliderBarDom"
						:round="10"
						:follow-dark="false"
						:width="sliderBarWidth"
						unit="px"
						:height="4"
						:color="props.tabsLineAniColor"
						:margin="[0, 0]"
						:padding="[0, 0]"
					></tm-sheet>
				</view>
			</scroll-view>
			<!-- #endif -->
		</tm-sheet>
		<!-- #ifndef APP-NVUE -->
		<view
			id="webIdTabs"
			@touchmove="onMove"
			@touchend="onEnd"
			@touchstart="onStart"
			@touchcancel="onEnd"
			@mousemove="onMove"
			@mouseup="onEnd"
			@mouseleave="onEnd"
			@mousedown="onStart"
			ref="tabsDom"
			:style="{
				width: props.swiper ? `${totalWidth}px` : `${props.width}rpx`,
				transform: props.swiper ? `translateX(${directoStyle}px)` : `translateX(0px)`
			}"
			v-if="_tabPos == 'top' && !isDulitabs"
			:class="[!isEndMove || isNvue ? 'tmTabsPane' : '']"
			class="flex flex-row flex-nowrap overflow"
		>
			<slot></slot>
		</view>
		<!-- #endif -->
		<!-- #ifdef APP-NVUE -->
		<!-- @touchmove="onMove"
		@touchend="onEnd" 
		@touchstart="onStart" -->
		<!-- @touchmove="onMove" @touchend="onEnd" @touchcancel="onEnd"  -->
		<view
			@touchstart="spinNvueAni"
			@touchmove="onMove"
			ref="tabsDom"
			:style="{
				width: props.swiper ? `${totalWidth}px` : `${props.width}rpx`,
				transform: `translateX(0px)`
			}"
			v-if="_tabPos == 'top' && !isDulitabs"
			class="flex flex-row flex-nowrap overflow"
		>
			<slot></slot>
		</view>
		<!-- #endif -->
	</view>
</template>
<script lang="ts" setup>
/**
 * 选项卡
 * @description 可以单独使用，配合tm-tabs-pane可实现卡片内容切换。
 * @example
 * <tm-tabs :width="750" default-name="6">
		<tm-tabs-pane v-for="item in 20" :name="item" :title="'未收件'+item">
			{{item}}
		</tm-tabs-pane>
	</tm-tabs>
 */
import {
	computed,
	ref,
	provide,
	watch,
	toRaw,
	nextTick,
	onMounted,
	watchEffect,
	PropType,
	getCurrentInstance,
	onUnmounted,
	ComponentInternalInstance
} from 'vue'
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmBadge from '../tm-badge/tm-badge.vue'
import { custom_props, computedClass } from '../../tool/lib/minxs'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
// #ifdef APP-NVUE || APP-PLUS-NVUE
var dom = weex.requireModule('dom')
const Binding = uni.requireNativePlugin('bindingx')
const animation = uni.requireNativePlugin('animation')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null

const bindxToken = ref(null)
const emits = defineEmits(['update:activeName', 'change', 'click'])
type alignType = 'left' | 'center' | 'around' | 'right'
const props = defineProps({
	...custom_props,
	//如果提供了，那么就不需要tm-tabs-pane，可以单独使用。
	list: {
		type: Array as PropType<Array<Tmui.tabs>>,
		default: () => []
	},
	rangKey: {
		type: String,
		default: 'title'
	},
	color: {
		type: String,
		default: 'white'
	},
	transprent: {
		type: [Boolean, String],
		default: false
	},
	width: {
		type: Number,
		default: 500
	},
	itemHeight: {
		type: Number,
		default: 80
	},
	//不设定窗口高度，在真机上有闪烁。如果设定为0将是自动高度。
	height: {
		type: Number,
		default: 1000
	},
	//内容在bar中的上下间隔。当有选项背景色时，底部为白色，这相当有用。
	gutter: {
		type: Number,
		default: 0
	},
	defaultName: {
		type: [String, Number],
		default: ''
	},
	//当前活动项。v-model:active-name
	activeName: {
		type: [String, Number],
		default: ''
	},
	//标签导航的位置，
	//top导航在上方，bottom导航在下方。
	tabPos: {
		type: String,
		default: 'top'
	},
	//项目的宽度。如果提供，每个标签是等宽度的，如果不提供自动宽度。
	itemWidth: {
		type: Number,
		default: 0
	},
	//tab选中的背景颜色。默认为空
	activeColor: {
		type: String,
		default: 'primary'
	},
	activeFontColor: {
		type: String,
		default: 'primary'
	},
	activeFontSize: {
		type: Number,
		default: 28
	},
	//选项卡样式模型
	itemModel: {
		type: String,
		default: 'text' //line底部线条，card背景颜色模式，text文本模式,textLight背景减淡模式，文字会变灰。
	},
	//默认为空即根据主题自定颜色。如果填写了将使用该颜色为未选中色。
	unSelectedColor: {
		type: String,
		default: ''
	},
	itemFontSize: {
		type: Number,
		default: 28
	},
	itemLinear: {
		type: String,
		default: ''
	},
	itemLinearDeep: {
		type: String,
		default: 'light'
	},
	itemRound: {
		type: Number,
		default: 0
	},
	/**
	 * 标题的分布方式
	 */
	align: {
		type: String as PropType<'left' | 'center' | 'around' | 'right'>,
		default: 'left' //left:左对齐,right：右对齐,center：剧中,around：剧中均分
	},
	//是否启用pane滑动切换tabs。如果关闭有助于页面更顺畅。如果启用请不要大量内容。
	swiper: {
		type: Boolean,
		default: false
	},
	//是否显示底部线条动画样式。
	showTabsLineAni: {
		type: Boolean,
		default: false
	},
	//是否显示底部线条动的底部灰色导轨
	showTabsLine: {
		type: Boolean,
		default: true
	},
	//下面活动的横线的颜色。
	tabsLineAniColor: {
		type: String,
		default: 'primary'
	},
	disAbledPull: {
		type: Boolean,
		default: true
	},
	//暗下强制的背景色，
	//有时自动的背景，可能不是你想要暗黑背景，此时可以使用此参数，强制使用背景色，
	//只能是颜色值。
	darkBgColor: {
		type: String,
		default: ''
	},
	/** 当选中某一项时,内容会往前滚动的项目数量,类似于位置让选中项始终在中间. */
	subtract: {
		type: Number,
		default: 2
	}
})
const _disAbledPull = computed(() => props.disAbledPull)
const _align = computed(() => {
	let align_list = {
		right: 'flex-row-center-end',
		left: 'flex-row-center-start',
		center: 'flex-row-center-center',
		around: 'flex-around'
	}
	let key: alignType = 'center'
	if (align_list.hasOwnProperty(props.align)) {
		key = props.align
	}
	return align_list[key]
})
const _active = ref(props.defaultName)
emits('update:activeName', _active.value)
const cstomClass = computed(() => computedClass(props))
const _scrollToId = ref('')
const modelStyle = computed(() => {
	if (props.itemModel == 'text') {
		return {
			transprent: true,
			border: 0,
			text: false
		}
	} else if (props.itemModel == 'line') {
		return {
			transprent: true,
			border: 4,
			text: false
		}
	} else if (props.itemModel == 'textLight') {
		return {
			transprent: false,
			border: 4,
			text: true
		}
	} else if (props.itemModel == 'card') {
		return {
			transprent: false,
			border: 0,
			text: false
		}
	}
	return {
		transprent: true,
		border: 0,
		text: false
	}
})
const tmTabsId = 'tmTabsId'
const _tabPos = computed(() => props.tabPos)
const cacheTabs = ref<Array<Tmui.tabs>>([])

const isDulitabs = computed(() => props.list.length > 0)
const tabsid = 'tabs_id_' + uni.$tm.u.getUid(1) + '_'
const isNvue = ref(false)
const _itemheight = Math.ceil(uni.upx2px(props.itemHeight))
const totalWidth = computed(() => uni.upx2px(cacheTabs.value.length * props.width))
const _itemwidth = Math.ceil(uni.upx2px(props.itemWidth))
const _sliderBarwidth = Math.ceil(uni.upx2px(40))
const _width = Math.ceil(uni.upx2px(props.width))
const contentWidth = computed(() => {
	let width = (props.itemWidth ) * cacheTabs.value.length
	if (width <= props.width) {
		width = props.width
	}
	return width
})
const contentWidth_px = computed(() => {
	let width = _itemwidth * cacheTabs.value.length
	if (width <= props.width) {
		width = uni.upx2px(props.width)
	}
	return Math.ceil(width)
})

// 线滚动动画。
const anitLineLeft = ref(0)

// #ifdef APP-NVUE
isNvue.value = true
// #endif

let timerId: any = NaN
let timerId2: any = NaN
function debounce(func: Function, wait = 500, immediate = false) {
	// 清除定时器
	if (!isNaN(timerId)) clearTimeout(timerId)
	// 立即执行，此类情况一般用不到

	if (immediate) {
		var callNow = !timerId
		timerId = setTimeout(() => {
			timerId = NaN
		}, wait)

		if (callNow) typeof func === 'function' && func()
	} else {
		// 设置定时器，当最后一次操作后，timeout不会再被清除，所以在延时wait毫秒后执行func回调方法
		timerId = setTimeout(() => {
			typeof func === 'function' && func()
		}, wait)
	}
}
function debounce2(func: Function, wait = 500, immediate = false) {
	// 清除定时器
	if (!isNaN(timerId2)) clearTimeout(timerId2)
	// 立即执行，此类情况一般用不到

	if (immediate) {
		var callNow = !timerId2
		timerId2 = setTimeout(() => {
			timerId2 = NaN
		}, wait)

		if (callNow) typeof func === 'function' && func()
	} else {
		// 设置定时器，当最后一次操作后，timeout不会再被清除，所以在延时wait毫秒后执行func回调方法
		timerId2 = setTimeout(() => {
			typeof func === 'function' && func()
		}, wait)
	}
}

// 判断滑动方向及距离start------------------------------------------------
const _startx = ref(0)
const _starty = ref(0)
const _movex = ref(0)
const _movey = ref(0)
const _x = ref(0)
const _y = ref(0)
const directo = 'none'
const _moveX = ref(0)
const _moveY = ref(0)
const directoStyle = ref(0) //左右距离
const isEndMove = ref(true)
const maxLen = 80 //只有拖拉距离大于此值才会切换。
const activeIndex = computed(() => cacheTabs.value.findIndex((el) => el.key == _active.value))
let ctxLeft = 0
let ctxTop = 0
let timeDetail = 1 //动画时长。
let isMoveEnb = false
// 当前处于什么状态.left,right,down,up
let dirType = ref('none')

// 当前是否处于拖动中，但实际未松手，还未形成真正意义 上切换.
let isDrag = ref(false)
let sliderBarWidth = uni.upx2px(40)
let widthDrag = ref(sliderBarWidth)
// 判断滑动方向及距离end------------------------------------------------

watchEffect(() => {
	cacheTabs.value = []
	props.list.forEach((el, index) => {
		cacheTabs.value.push({
			key: el?.key ?? el?.id ?? String(index),
			title: el[props.rangKey] ?? String(index),
			icon: el?.icon ?? '',
			dot: el?.dot ?? false,
			count: el?.count ?? '',
			dotColor: el?.dotColor ?? 'red'
		})
	})
})

function setTabsBarLineLeft(key: string | number = '') {
	if (!props.showTabsLineAni) return
	let keybl = key || _active.value
	let index = cacheTabs.value.findIndex((el) => el.key == keybl)

	if (index > -1) {
		let leftPx = _itemwidth * index
		if (props.align == 'center') {
			leftPx = leftPx + (_width - _itemwidth * cacheTabs.value.length) / 2 - 1
		}

		anitLineLeft.value = Math.ceil(leftPx)
	}
}

function unbindKey(key: string | number) {
	let index: number = cacheTabs.value.findIndex((el) => el.key == key)
	if (index > -1) {
		cacheTabs.value.splice(index, 1)
	}
	let index2: number = cacheTabs.value.findIndex((el) => el.key == _active.value)

	if (index2 == -1 && cacheTabs.value.length > 0) {
		changeKey(cacheTabs.value[0]?.key ?? '', false, false)
	} else if (cacheTabs.value.length == 0) {
		changeKey('', false, false)
	}
}
watch(
	() => props.activeName,
	() => {
		if (props.activeName == _active.value) return
		changeKey(props.activeName, false, false)
	}
)

onMounted(() => {
	setTimeout(() => {
		_scrollToId.value = tabsid + _active.value
		nextTick(() => {
			// #ifdef APP-NVUE
			dom.getComponentRect(proxy?.$refs?.tabsDom, function (res) {
				if (res?.size) {
					ctxLeft = Math.floor(res.size.left)
					ctxTop = Math.floor(res.size.top)

					spinNvueAniEnd(0, -uni.upx2px(activeIndex.value * props.width), 1)
					_startx.value = uni.upx2px(activeIndex.value * props.width)
				}
			})
			// #endif
			setTabsBarLineLeft(props.defaultName)
		})
	}, 300)
})
watchEffect(() => {
	directoStyle.value = Math.ceil(uni.upx2px(-activeIndex.value * props.width))
	spinNvueAniEnd(0, -uni.upx2px(activeIndex.value * props.width), timeDetail)
})
watch(
	() => _active.value,
	() => {
		nextTick(() => {
			let index = cacheTabs.value.findIndex((el) => el.key == _active.value)

			if (index > -1) {
				if (typeof cacheTabs.value[index - props.subtract] !== 'undefined') {
					_scrollToId.value = tabsid + cacheTabs.value[index - props.subtract]?.key
				} else {
					_scrollToId.value = tabsid + cacheTabs.value[0]?.key
				}
			} else {
				_scrollToId.value = tabsid + _active.value
			}

			setTabsBarLineLeft()
		})
	}
)

// 确实当前是单击还是移动。点击
let isMoveing = ref(false)
function webClick() {}
function onStart(event: TouchEvent) {
	if (!props.swiper) return
	isEndMove.value = true
	isMoveEnb = true
	isMoveing.value = false
	isDrag.value = true
	if (event.type.indexOf('mouse') == -1 && event.changedTouches.length == 1) {
		var touch = event.changedTouches[0]
		if (typeof touch?.pageX !== 'undefined') {
			_startx.value = touch.pageX - ctxLeft
			_starty.value = touch.pageY - ctxTop
		} else {
			_startx.value = touch.x
			_starty.value = touch.y
		}
	} else {
		_startx.value = event.pageX - event.currentTarget.offsetLeft - ctxLeft
		_starty.value = event.pageY - event.currentTarget.offsetTop - ctxTop
	}
}
function onMove(event: TouchEvent) {
	if (!props.swiper || isMoveEnb == false) return
	isMoveing.value = true
	let nowx = 0
	let nowy = 0
	if (event.type.indexOf('mouse') == -1 && event.changedTouches.length == 1) {
		var touch = event.changedTouches[0]
		if (typeof touch?.pageX !== 'undefined') {
			nowx = touch.pageX - ctxLeft
			nowy = touch.pageY - ctxTop
		} else {
			nowx = touch.x
			nowy = touch.y
		}
	} else {
		nowx = event.pageX - event.currentTarget.offsetLeft - ctxLeft
		nowy = event.pageY - event.currentTarget.offsetTop - ctxTop
	}

	_x.value = nowx - _startx.value
	_y.value = nowy - _starty.value
	setDirXy(_x.value, _y.value)
}

function onEnd(event: TouchEvent) {
	if (!props.swiper || !isMoveEnb || !isMoveing.value) return
	isEndMove.value = false
	isMoveing.value = false
	debounce2(
		() => {
			setDirXy(_x.value, _y.value, true)
			isDrag.value = false
		},
		250,
		true
	)
	isMoveEnb = false
}

function setDirXy(x: number, y: number, isEnd = false) {
	const oldindex = activeIndex.value
	let nowLeft = uni.upx2px(activeIndex.value * props.width)

	// 这里防止抖动引起的方向错误
	debounce(
		() => {
			if (x > 0 && Math.abs(x) > Math.abs(y)) {
				dirType.value = 'right'
			} else if (x < 0 && Math.abs(x) > Math.abs(y)) {
				dirType.value = 'left'
			} else if (y > 0 && Math.abs(y) > Math.abs(x)) {
				dirType.value = 'down'
			} else if (y < 0 && Math.abs(y) > Math.abs(x)) {
				dirType.value = 'up'
			} else {
				dirType.value = 'none'
			}
		},
		300,
		true
	)

	if (dirType.value == 'right') {
		//第一和最后一张，不需要滑动。
		if (activeIndex.value == 0) return
		directoStyle.value = x - nowLeft
		// #ifdef H5
		let sx = Math.abs(sliderBarWidth) * 1.05
		sx = Math.min(sx, sliderBarWidth)
		sx = Math.max(sx, _itemwidth)
		widthDrag.value = sx
		// #endif
		if (isEnd) {
			setRightDirRight()
			widthDrag.value = sliderBarWidth
		}
	} else if (dirType.value == 'left') {
		//第一和最后一张，不需要滑动。
		if (activeIndex.value == cacheTabs.value.length - 1) return
		directoStyle.value = x - nowLeft
		// #ifdef H5
		let sx = Math.abs(_x.value) * 1.0002
		sx = Math.min(sx, sliderBarWidth)
		sx = Math.max(sx, _itemwidth)
		widthDrag.value = sx
		// #endif
		if (isEnd) {
			setLeftDirLeft()
			widthDrag.value = sliderBarWidth
		}
	}

	function setRightDirRight() {
		if (x < maxLen || activeIndex.value <= 0) {
			directoStyle.value = -nowLeft
		} else {
			_active.value = cacheTabs.value[activeIndex.value - 1]?.key ?? -1
			changeKey(_active.value, false)
		}
	}

	function setLeftDirLeft() {
		if (Math.abs(x) < maxLen || activeIndex.value >= cacheTabs.value.length - 1) {
			directoStyle.value = -nowLeft
		} else {
			_active.value = cacheTabs.value[activeIndex.value + 1]?.key ?? -1
			changeKey(_active.value, false)
		}
	}
}

function getEl(el: HTMLElement) {
	if (typeof el === 'string' || typeof el === 'number') return el
	// @ts-nocheck
	if (WXEnvironment) {
		return el.ref
	} else {
		return el instanceof HTMLElement ? el : el?.$el
	}
}
onUnmounted(() => {
	// #ifdef APP-PLUS-NVUE
	if (bindxToken.value) {
		Binding.unbind({
			token: bindxToken.value,
			eventType: 'timing'
		})
	}
	// #endif
})

function setDirXyNvue(x: number, y: number, dirX = 'none') {
	const oldindex = activeIndex.value
	let nowLeft = uni.upx2px(activeIndex.value * props.width)
	let maxLeft = uni.upx2px((cacheTabs.value.length - 1) * props.width)
	dirType.value = dirX
	// _startx.value 为当前tabs相对左边的距离。 ，x为偏移量.

	if (dirType.value == 'right') {
		//第一张回弹
		if (activeIndex.value == 0) {
			_startx.value = 0
			spinNvueAniEnd(0, 0, 250)
			return
		}
		// 小于此阀值，不产生切换。
		if (Math.abs(x) < maxLen) {
			_startx.value = nowLeft
			spinNvueAniEnd(-nowLeft, 0, 250)
		} else {
			// 进入上一个
			_active.value = cacheTabs.value[activeIndex.value - 1]?.key ?? -1
			nowLeft = uni.upx2px(activeIndex.value * props.width)
			_startx.value = nowLeft
			spinNvueAniEnd(-nowLeft, 0, 250)
			changeKey(_active.value, false)
		}
	} else if (dirType.value == 'left') {
		//最后一张回弹
		if (activeIndex.value == cacheTabs.value.length - 1) {
			_startx.value = maxLeft
			spinNvueAniEnd(-maxLeft, 0, 250)
			return
		}
		// 小于此阀值，不产生切换。
		if (Math.abs(x) < maxLen) {
			_startx.value = nowLeft
			spinNvueAniEnd(-nowLeft, 0, 250)
		} else {
			// 进入下一个
			_active.value = cacheTabs.value[activeIndex.value + 1]?.key ?? -1
			nowLeft = uni.upx2px(activeIndex.value * props.width)
			_startx.value = nowLeft
			spinNvueAniEnd(-nowLeft, 0, 250)

			changeKey(_active.value, false)
		}
	}
}

function spinNvueAni() {
	if (!props.swiper) return
	// #ifdef APP-NVUE
	if (!proxy?.$refs?.tabsDom) return

	let icon = getEl(proxy?.$refs?.tabsDom)
	let sliderBarDom = getEl(proxy?.$refs?.sliderBarDom)
	let icon_bind = Binding.bind(
		{
			anchor: icon,
			eventType: 'pan',
			props: [
				{
					element: icon,
					property: 'transform.translateX',
					expression: `abs(y)>abs(x)?${-_startx.value}:x+` + (0 - _startx.value)
				}
			]
		},
		function (res: any) {
			if (res.state == 'end') {
				isMoveing.value = false
				if (Math.abs(res.deltaY) > 80) {
					let nowLeft = uni.upx2px(activeIndex.value * props.width)
					spinNvueAniEnd(-nowLeft, 0, 0)
					return
				}
				const startx = _startx.value
				_startx.value -= res.deltaX
				_starty.value += res.deltaY
				let lx = 'left'
				if (res.deltaX > 0) {
					lx = 'right'
				}
				setDirXyNvue(res.deltaX, res.deltaY, lx)
			} else if (res.state == 'start') {
				isMoveing.value = true
			}
		}
	)
	// #endif
}
function spinNvueAniEnd(start: number, end: number, time = timeDetail) {
	if (!props.swiper) return
	// #ifdef APP-NVUE
	if (!proxy?.$refs?.tabsDom) return
	animation.transition(
		proxy?.$refs?.tabsDom,
		{
			styles: {
				transform: `translateX(${start + end}px)`,
				transformOrigin: 'center center'
			},
			duration: time, //ms
			timingFunction: 'ease',
			delay: 0 //ms
		},
		() => {}
	)

	// #endif
}

function pushKey(o: Tmui.tabs) {
	let index: number = cacheTabs.value.findIndex((el) => el.key === o.key)
	if (index > -1) {
		cacheTabs.value.splice(index, 1, {
			...cacheTabs.value[0],
			...o
		})
	} else {
		cacheTabs.value.push(o)
	}

	if (_active.value == '') {
		changeKey(cacheTabs.value[0]?.key ?? -1, false)
	}
}

function changeKey(key: string | number, isclick = true, isNomarlChange = true) {
	isEndMove.value = true
	_active.value = key
	// #ifdef APP-NVUE
	_startx.value = uni.upx2px(activeIndex.value * props.width)
	// #endif
	timeDetail = 1
	emits('update:activeName', toRaw(_active.value))
	if (isNomarlChange) {
		nextTick(() => {
			emits('change', key)
		})
	}
	if (isclick) {
		nextTick(() => {
			emits('click', key)
		})
	}
}

function setTitle(o: Tmui.tabs) {
	let index: number = cacheTabs.value.findIndex((el) => el.key == o.key)
	if (index > -1) {
		cacheTabs.value.splice(index, 1, o)
	}
}
provide(
	'tabsActiveName',
	computed(() => _active.value)
)
provide('tabsActiveactiveIndex', activeIndex)
provide(
	'tabsActiveCacheTabse',
	computed(() => cacheTabs.value)
)

provide(
	'tabsWidth',
	computed(() => props.width)
)
provide(
	'tabsheight',
	computed(() => {
		if (!props.height) return 0
		return props.height - props.itemHeight - props.gutter
	})
)
provide(
	'tabsSwiper',
	computed(() => props.swiper)
)
provide(
	'tabsSwiperIsMoveing',
	computed(() => isMoveing.value)
)
provide(
	'tabsSwiperDisAbledPull',
	computed(() => props.disAbledPull)
)

defineExpose({
	pushKey: pushKey,
	changeKey: changeKey,
	unbindKey,
	setTitle: setTitle,
	tmTabsId
})
</script>
<style>
.animateAll_tabs_tmui {
	transition-delay: 0;
	transition-timing-function: ease;
	transition-property: transform, width;
	transition-duration: 0.3s;
}
</style>
<style scoped>
	.nonvue{
		/* #ifndef APP-NVUE */
		box-sizing: border-box;
		/* #endif */
	}
.tmTabsPane {
	transition-delay: 0;
	transition-timing-function: ease;
	transition-property: transform;
	transition-duration: 0.3s;
}

.anilineBar {
	transition-delay: 0;
	transition-timing-function: ease;
	transition-property: transform, width;
	transition-duration: 0.3s;
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
