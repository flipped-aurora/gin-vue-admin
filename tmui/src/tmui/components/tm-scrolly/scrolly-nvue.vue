<template>
	<scroll-view scroll-y @scrolltolower="onScrollToBottom" :style="{ height: props.height + 'rpx' }">
		<refresh @refresh="onrefresh" @pullingdown="onpullingdown" :display="showLoading ? 'show' : 'hide'">
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
					v-if="refreshStatus != -1 && refreshStatus != 2"
					class="flex flex-row flex-row-center-center srrryration"
					:style="{
						opacity: `${refreshStatus == 0 ? _barHeight / nowEvt.viewHeight : 1}`,
						height: _barHeight / 2 + 'px'
					}"
				>
					<view :class="refreshStatus == 0 ? 'srrryration srrryrationOn' : 'srrryration srrryrationOf'">
						<tm-icon :font-size="24" color="primary" name="tmicon-long-arrow-down"></tm-icon>
					</view>
					<tm-text color="grey" _class="pl-16" :label="loadingTexts[refreshStatus]"></tm-text>
				</view>
			</slot>
		</refresh>
		<slot></slot>
	</scroll-view>
</template>

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
	if (refreshStatus.value == -1 || refreshStatus.value == 0) {
		showLoading.value = true
		nextTick(() => {
			showLoading.value = false
			emits('change', false)
		})
	}
}
/**
	 * dy: 前后两次回调滑动距离的差值
	pullingDistance: 下拉的距离
	viewHeight: refresh 组件高度
	type: “pullingdown” 常数字符串
	 */
function onpullingdown(evt: { dy: number; pullingDistance: number; viewHeight: number; type: string }) {
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
	refreshStatus.value = 3
	emits('change', false)
	emits('update:modelValue', false)
}
</script>

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
