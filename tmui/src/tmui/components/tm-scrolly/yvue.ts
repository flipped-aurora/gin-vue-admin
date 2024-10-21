import tmIcon from "../tm-icon/tm-icon.vue"
import tmText from "../tm-text/tm-text.vue"
import {
	getCurrentInstance,
	nextTick,
	onMounted,
	ref,
	Ref,
	watch
} from 'vue';
import { propsdetail } from "./propsdetail"
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(["bottom", "change", "refresh", "timeout", "update:modelValue", "update:bottomValue"])
const props = defineProps({
	...propsdetail
})
// 下拉开始的起点，主要用于计算下拉高度
const startPoint: Ref<{
	pageX: number; pageY: number
} | null> = ref(null);
const isPulling = ref(false); // 是否下拉中
const _maxBarHeight = ref(props.maxBarHeight); // 最大下拉高度，单位 rpx
// 触发刷新的下拉高度，单位rpx
// 松开时下拉高度大于这个值即会触发刷新，触发刷新后松开，会恢复到这个高度并保持，直到刷新结束
const _barHeight = ref(0)
/** 开始刷新 - 刷新成功/失败 最大间隔时间setTimeout句柄 */
let maxRefreshAnimateTimeFlag: number | null = 0;
/** 关闭动画耗时setTimeout句柄 */
let closingAnimateTimeFlag: number | null = 0;
//加载框的高度
const refreshStatus = ref(-1)
const loosing = ref(false)
const enableToRefresh = ref(true)
const scrollTop = ref(0)

/** 触底下拉刷新参数。 */
const isBootRefresh = ref(props.bottomValue)


watch(() => props.modelValue, () => {
	if (!props.modelValue) {
		if (maxRefreshAnimateTimeFlag != null) {
			clearTimeout(maxRefreshAnimateTimeFlag);
		}
		refreshStatus.value = 3;
		close();
	}
})
watch(() => props.bottomValue, () => {
	isBootRefresh.value = props.bottomValue;
})
onMounted(() => {
	clearTimeout(maxRefreshAnimateTimeFlag);
	clearTimeout(closingAnimateTimeFlag);
	nextTick(() => setDefault())
})

function setDefault() {
	if (props.defaultValue) {
		setRefreshBarHeight(props.loadBarHeight)
		refreshStatus.value = 2;
		loosing.value = true;
		isPulling.value = true;
		enableToRefresh.value = false;
		startPoint.value = null;
	}
}

function onScrollToBottom() {
	if (isBootRefresh.value) return;
	emits("update:bottomValue")
	emits("bottom")
}

function onScrollToTop() {
	enableToRefresh.value = true;
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

function onTouchStart(e: WechatMiniprogram.Component.TrivialInstance) {
	if (isPulling.value || !enableToRefresh.value) return;
	const {
		touches
	} = e;
	if (touches.length !== 1) return;
	const {
		pageX,
		pageY
	} = touches[0];

	loosing.value = false;
	startPoint.value = {
		pageX,
		pageY
	};
	isPulling.value = true;

}

function onTouchMove(e: WechatMiniprogram.Component.TrivialInstance) {
	if (!startPoint.value) return;
	const {
		touches
	} = e;

	if (touches.length !== 1) return;

	const {
		pageY
	} = touches[0];
	const offset = pageY - startPoint.value.pageY;
	const barsHeight = uni.$tm.u.torpx(offset)


	if (barsHeight > 0) {

		if (barsHeight > _maxBarHeight.value) {
			// 限高
			setRefreshBarHeight(_maxBarHeight.value);
			// this.startPoint.pageY = pageY - this.toPx(this.maxBarHeight); // 限高的同时修正起点，避免触摸点上移时无效果
		} else {
			setRefreshBarHeight(barsHeight);
		}
	}
}

function onTouchEnd(e: WechatMiniprogram.Component.TrivialInstance) {
	if (!startPoint.value) return;
	const {
		changedTouches
	} = e;
	if (changedTouches.length !== 1) return;
	const {
		pageY
	} = changedTouches[0];
	const barsHeight = uni.$tm.u.torpx(pageY - startPoint.value.pageY);
	startPoint.value = null; // 清掉起点，之后将忽略touchMove、touchEnd事件

	loosing.value = true;
	isBootRefresh.value = false;
	// 松开时高度超过阈值则触发刷新
	if (barsHeight > props.loadBarHeight) {

		_barHeight.value = props.loadBarHeight
		refreshStatus.value = 2;

		emits("change", true)
		emits("update:modelValue", true);
		emits("refresh")
		maxRefreshAnimateTimeFlag = setTimeout(() => {
			maxRefreshAnimateTimeFlag = null;

			if (refreshStatus.value === 2) {
				// 超时回调
				emits("timeout")
				close(); // 超时仍未被回调，则直接结束下拉
			}
		}, props.refreshTimeout as any) as any as number;
	} else {
		close();
	}
}

function setRefreshBarHeight(barsHeight: number) {
	if (barsHeight >= props.loadBarHeight) {
		refreshStatus.value = 1;
	} else {
		refreshStatus.value = 0;
	}
	return new Promise((resolve) => {
		_barHeight.value = barsHeight;
		nextTick(() => {
			resolve(barsHeight)
		})
	});
}

function close() {
	const animationDuration = 350;
	_barHeight.value = 0
	emits('change', false)
	emits("update:modelValue", false);
	closingAnimateTimeFlag = setTimeout(() => {
		closingAnimateTimeFlag = null;
		refreshStatus.value = -1;
		isPulling.value = false; // 退出下拉状态
		loosing.value = false;
		enableToRefresh.value = true;
	}, animationDuration) as any as number;
}