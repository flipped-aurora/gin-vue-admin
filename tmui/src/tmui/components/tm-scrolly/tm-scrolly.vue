<template>
	<!-- #ifndef APP-NVUE -->
	<scroll-view
		class="scroyy"
		:scroll-top="scrollTop"
		@touchstart="onTouchStart"
		@touchmove="onTouchMove"
		@touchend="onTouchEnd"
		@scroll="onScroll"
		@scrolltoupper="onScrollToTop"
		@scrolltolower="onScrollToBottom"
		scroll-y
		:enable-flex="true"
		enable-back-to-top
		enhanced
		scroll-with-animation
		:bounces="false"
		:style="[props.height ? { height: props.height + 'rpx' } : '', _style]"
		:class="_class"
	>
		<view
			:class="['scroyy__track', 'scroyy__track--' + (loosing ? 'loosing' : '')]"
			:style="{ transform: `translate3d(0, ${_barHeight}rpx, 0)` }"
		>
			<view class="scroyy__tips" :class="['scroyy__track--' + (loosing ? 'loosing' : '')]" :style="{ height: _barHeight + 'rpx' }">
				<slot name="pull" :status="{ refreshStatus }">
					<view v-if="refreshStatus === 2" class="flex flex-row flex-row-center-center">
						<tm-icon :font-size="24" color="primary" name="tmicon-shuaxin" spin></tm-icon>
						<tm-text color="grey" _class="pl-16" :label="loadingTexts[refreshStatus]"></tm-text>
					</view>
					<view
						v-if="refreshStatus != -1 && refreshStatus != 2"
						class="flex flex-row flex-row-center-center srrryration"
						:style="{
							opacity: `${refreshStatus == 0 ? _barHeight / props.loadBarHeight : 1}`
						}"
					>
						<view :class="refreshStatus == 0 ? 'srrryration srrryrationOn' : 'srrryration srrryrationOf'">
							<tm-icon :font-size="24" color="primary" name="tmicon-long-arrow-down"></tm-icon>
						</view>
						<tm-text color="grey" _class="pl-16" :label="loadingTexts[refreshStatus]"></tm-text>
					</view>
				</slot>
			</view>
			<slot></slot>
			<view :class="['scroyy__track--loosing  ']" :style="{ height: (isBootRefresh ? props.loadBarHeight : 0) + 'rpx' }">
				<slot name="bottom" :status="{ isBootRefresh }">
					<view
						v-if="isBootRefresh"
						class="flex flex-row flex-row-center-center"
						:style="{ height: (isBootRefresh ? props.loadBarHeight : 0) + 'rpx' }"
					>
						<tm-icon :font-size="24" color="primary" name="tmicon-shuaxin" spin></tm-icon>
						<tm-text color="grey" _class="pl-16" label="数据加载中"></tm-text>
					</view>
				</slot>
			</view>
		</view>
	</scroll-view>
	<!-- #endif -->
	<!-- #ifdef APP-NVUE -->
	<scroll-view 
	enableBackToTop="true" 
	alwaysScrollableVertical="true" 
	scroll-y="true" 
	@scrolltolower="onScrollToBottom" 
	:style="[props.height ? { height: props.height + 'rpx' } : '', _style]" :class="_class">
		<refresh @refresh="onrefresh" @pullingdown="onpullingdown" :display="showLoading ? 'show' : 'hide'" style="width: 750rpx">
			<view :style="{ height: _barHeight / 2 + 'px' }">
				<slot name="pull" :status="{ refreshStatus }">
					<view
						class="flex flex-row flex-row-center-center"
						:style="{
							height: _barHeight / 2 + 'px',
							opacity: refreshStatus == 2 || refreshStatus == 3 ? 1 : 0
						}"
					>
						<tm-icon :font-size="24" color="primary" name="tmicon-shuaxin" spin></tm-icon>
						<tm-text color="grey" _class="pl-16" :label="loadingTexts[refreshStatus]"></tm-text>
					</view>
					<view
						v-if="refreshStatus == 0 || refreshStatus == 1"
						class="flex flex-row flex-row-center-center srrryration"
						:style="{
							opacity: `${refreshStatus == 0 ? _barHeight / nowEvt.viewHeight : 1}`,
							height: _barHeight / 2 + 'px',
							top: -(_barHeight / 2) + 'px'
						}"
					>
						<view :class="refreshStatus == 0 ? 'srrryration srrryrationOn' : 'srrryration srrryrationOf'">
							<tm-icon :font-size="24" color="primary" name="tmicon-long-arrow-down"></tm-icon>
						</view>
						<tm-text color="grey" _class="pl-16" :label="loadingTexts[refreshStatus]"></tm-text>
					</view>
				</slot>
			</view>
		</refresh>
		<slot></slot>
	</scroll-view>
	<!-- #endif -->
</template>
<!-- #ifndef APP-NVUE -->
<script lang="ts" setup>
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import { getCurrentInstance, nextTick, onMounted, ref, Ref, watch, computed } from 'vue'
import { propsdetail } from './propsdetail'
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['bottom', 'change', 'refresh', 'timeout', 'update:modelValue', 'update:bottomValue'])
const props = defineProps({
	...propsdetail
})
// 下拉开始的起点，主要用于计算下拉高度
const startPoint: Ref<{
	pageX: number
	pageY: number
} | null> = ref(null)
const isPulling = ref(false) // 是否下拉中
const _maxBarHeight = ref(props.maxBarHeight) // 最大下拉高度，单位 rpx
const _refresher = computed(()=>props.refresher)
// 触发刷新的下拉高度，单位rpx
// 松开时下拉高度大于这个值即会触发刷新，触发刷新后松开，会恢复到这个高度并保持，直到刷新结束
const _barHeight = ref(0)
/** 开始刷新 - 刷新成功/失败 最大间隔时间setTimeout句柄 */
let maxRefreshAnimateTimeFlag: any | null = 0
/** 关闭动画耗时setTimeout句柄 */
let closingAnimateTimeFlag: any | null = 0
//加载框的高度
const refreshStatus = ref(-1)
const loosing = ref(false)
const enableToRefresh = ref(true)
const scrollTop = ref(0)
const loadingTexts = computed(() => props.loadingTexts)
/** 触底下拉刷新参数。 */
const isBootRefresh = ref(props.bottomValue)
const _class = computed(() => props._class)
const _style = computed(() => props._style)
watch(
	() => props.modelValue,
	() => {
		if (!props.modelValue) {
			if (maxRefreshAnimateTimeFlag != null) {
				clearTimeout(maxRefreshAnimateTimeFlag)
			}
			refreshStatus.value = 3
			close()
		}
	}
)
watch(
	() => props.bottomValue,
	() => {
		isBootRefresh.value = props.bottomValue
	}
)
onMounted(() => {
	clearTimeout(maxRefreshAnimateTimeFlag)
	clearTimeout(closingAnimateTimeFlag)
	nextTick(() => setDefault())
})

function setDefault() {
	if (props.defaultValue) {
		setRefreshBarHeight(props.loadBarHeight)
		refreshStatus.value = 2
		loosing.value = true
		isPulling.value = true
		enableToRefresh.value = false
		startPoint.value = null
	}
}

function onScrollToBottom() {
	if (isBootRefresh.value) return
	emits('update:bottomValue')
	emits('bottom')
}

function onScrollToTop() {
	enableToRefresh.value = true
}

function onScroll(e: any) {
	enableToRefresh.value = e.detail?.scrollTop === 0
}

function setScrollTop(tp: number) {
	scrollTop.value = tp
}

function scrollToTop() {
	setScrollTop(0)
}

function onTouchStart(e: TouchEvent) {
	if (isPulling.value || !enableToRefresh.value || !_refresher.value) return
	const { touches } = e
	if (touches.length !== 1) return
	const { pageX, pageY } = touches[0]

	loosing.value = false
	startPoint.value = {
		pageX,
		pageY
	}
	isPulling.value = true
}

function onTouchMove(e: TouchEvent) {
	if (!startPoint.value || !_refresher.value) return
	const { touches } = e

	if (touches.length !== 1) return

	const { pageY } = touches[0]
	const offset = pageY - startPoint.value.pageY
	const barsHeight = uni.$tm.u.torpx(offset)

	if (barsHeight > 0) {
		if (barsHeight > _maxBarHeight.value) {
			// 限高
			setRefreshBarHeight(_maxBarHeight.value)
			// this.startPoint.pageY = pageY - this.toPx(this.maxBarHeight); // 限高的同时修正起点，避免触摸点上移时无效果
		} else {
			setRefreshBarHeight(barsHeight)
		}
	}
}

function onTouchEnd(e: TouchEvent) {
	if (!startPoint.value || !_refresher.value) return
	const { changedTouches } = e
	if (changedTouches.length !== 1) return
	const { pageY } = changedTouches[0]
	const barsHeight = uni.$tm.u.torpx(pageY - startPoint.value.pageY)
	startPoint.value = null // 清掉起点，之后将忽略touchMove、touchEnd事件

	loosing.value = true
	isBootRefresh.value = false
	// 松开时高度超过阈值则触发刷新
	if (barsHeight > props.loadBarHeight) {
		_barHeight.value = props.loadBarHeight
		refreshStatus.value = 2

		emits('change', true)
		emits('update:modelValue', true)
		emits('refresh')
		maxRefreshAnimateTimeFlag = setTimeout(() => {
			maxRefreshAnimateTimeFlag = null

			if (refreshStatus.value === 2) {
				// 超时回调
				emits('timeout')
				close() // 超时仍未被回调，则直接结束下拉
			}
		}, props.refreshTimeout as any) as any as number
	} else {
		close()
	}
}

function setRefreshBarHeight(barsHeight: number) {
	if(!_refresher.value) return;
	if (barsHeight >= props.loadBarHeight) {
		refreshStatus.value = 1
	} else {
		refreshStatus.value = 0
	}
	return new Promise((resolve) => {
		_barHeight.value = barsHeight
		nextTick(() => {
			resolve(barsHeight)
		})
	})
}

function close() {
	const animationDuration = 350
	_barHeight.value = 0
	emits('change', false)
	emits('update:modelValue', false)
	closingAnimateTimeFlag = setTimeout(() => {
		closingAnimateTimeFlag = null
		refreshStatus.value = -1
		isPulling.value = false // 退出下拉状态
		loosing.value = false
		enableToRefresh.value = true
	}, animationDuration) as any as number
}
</script>

<style scoped>
.scroyy {
	overflow: hidden;
	/* #ifndef APP-NVUE */
	max-height: 100%;
	/* #endif */
}

.scroyy__track {
	position: relative;
}

.scroyy__track--loosing {
	transition-property: transform, height, opacity;
	transition-timing-function: ease;
	transition-duration: 0.35s;
}

.scroyy__tips {
	position: absolute;
	color: #bbb;
	font-size: 24rpx;
	top: 0;
	width: 100%;
	transform: translateY(-100%);
	display: flex;
	justify-content: center;
	align-items: center;
	flex-direction: column;
	overflow: hidden;
}

.scroyy__text {
	margin: 16rpx 0 0;
}

.scroyy__wrap {
	position: relative;
}
</style>
<!-- #endif -->

<!-- #ifdef APP-NVUE -->

<script lang="ts" setup>
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import { getCurrentInstance, nextTick, onMounted, ref, Ref, watch } from 'vue'
import { propsdetail } from './propsdetail'
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['bottom', 'change', 'refresh', 'timeout', 'update:modelValue', 'update:bottomValue'])
const props = defineProps({
	...propsdetail
})
interface pullevent {
	dy: number
	pullingDistance: number
	viewHeight: number
	type: string
}

const showLoading = ref(false)
const refreshing = ref(false)
const isBootRefresh = ref(false)
const refreshStatus = ref(-1)
/** 开始刷新 - 刷新成功/失败 最大间隔时间setTimeout句柄 */
let maxRefreshAnimateTimeFlag: number | null = 0
const _barHeight = ref(uni.$tm.u.topx(props.loadBarHeight))
const nowEvt: Ref<pullevent> = ref({
	dy: 0,
	pullingDistance: 0,
	viewHeight: 0,
	type: ''
})

onMounted(() => {
	if (maxRefreshAnimateTimeFlag != null) {
		clearTimeout(maxRefreshAnimateTimeFlag)
	}
})
watch(
	() => props.modelValue,
	() => {
		if (!props.modelValue) {
			if (maxRefreshAnimateTimeFlag != null) {
				clearTimeout(maxRefreshAnimateTimeFlag)
			}
			close()
		}
	}
)
watch(
	() => props.bottomValue,
	() => {
		isBootRefresh.value = props.bottomValue
	}
)
function onScrollToBottom() {
	if (isBootRefresh.value) return
	emits('update:bottomValue', true)
	isBootRefresh.value = true
	emits('bottom')
}
function onrefresh() {
	if (nowEvt.value.pullingDistance >= _barHeight.value) {
		refreshStatus.value = 2
		showLoading.value = true
		emits('refresh')
		emits('change', true)
		emits('update:modelValue', true)
		maxRefreshAnimateTimeFlag = setTimeout(() => {
			maxRefreshAnimateTimeFlag = null
			if (refreshStatus.value === 2) {
				// 超时回调
				emits('timeout')
				close() // 超时仍未被回调，则直接结束下拉
			}
		}, props.refreshTimeout as any) as any as number

		return
	}
	showLoading.value = true
	nextTick(() => {
		showLoading.value = false
		emits('change', false)
	})
}
/**
	 * dy: 前后两次回调滑动距离的差值
	pullingDistance: 下拉的距离
	viewHeight: refresh 组件高度
	type: “pullingdown” 常数字符串
	 */
function onpullingdown(evt: { dy: number; pullingDistance: number; viewHeight: number; type: string }) {
	evt.pullingDistance = Math.abs(evt.pullingDistance)
	nowEvt.value = evt
	if (evt.pullingDistance <= 1) {
		refreshStatus.value = -1
	} else if (evt.pullingDistance > 1 && evt.pullingDistance < _barHeight.value) {
		refreshStatus.value = 0
	} else if (evt.pullingDistance >= _barHeight.value) {
		refreshStatus.value = 1
	}
}

function close() {
	showLoading.value = false
	refreshStatus.value = -1
	emits('change', false)
	emits('update:modelValue', false)
}
</script>

<!-- #endif -->

<style>
.srrryration {
	transition-property: transform, height, opacity;
	transition-timing-function: ease;
	transition-duration: 0.25s;
}

.srrryrationOn {
	transform: rotate(0deg);
}

.srrryrationOf {
	transform: rotate(180deg);
}
</style>
