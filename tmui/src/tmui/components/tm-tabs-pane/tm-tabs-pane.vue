<template>
	<tm-sheet
		v-if="tabsSwiper || (!tabsSwiper && isShowRender)"
		:transprent="props.transprent"
		:color="props.color"
		:followTheme="props.followTheme"
		:followDark="props.followDark"
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
		:_style="props._style"
		:_class="props._class"
		:eventPenetrationEnabled="true"
		:margin="[0, 0]"
		:padding="[0, 0]"
		:width="_width"
		:height="_height"
	>
		<view
			v-if="sc_top < -30 && reFresh != 0 && !tabsSwiperDisAbledPull"
			:style="{
				top: (reFresh == 2 ? -refreshJuli / 2 : -sc_top - 30) + 'px',
				width: _width + 'rpx'
			}"
			class="zIndex-17 absolute l-0 flex flex-row flex-row-center-center"
		>
			<view class="pr-32">
				<tm-icon v-if="sc_top > refreshJuli && reFresh != 2" name="tmicon-long-arrow-down"></tm-icon>
				<tm-icon v-if="sc_top <= refreshJuli && reFresh != 2" name="tmicon-long-arrow-up"></tm-icon>
				<tm-icon v-if="reFresh == 2" spin name="tmicon-shuaxin"></tm-icon>
			</view>
			<view>
				<tm-text v-if="sc_top > refreshJuli" _class="text-align-center" label="下拉刷新"></tm-text>
				<tm-text v-if="sc_top <= refreshJuli" _class="text-align-center" label="松开立即刷新"></tm-text>
				<tm-text _class="text-align-center" label="更新于今日8:30"></tm-text>
			</view>
		</view>
		<scroll-view
			@scrolltolower="onScrollBootom"
			@scroll="onScroll"
			:scroll-y="_height ? true : false"
			enable-flex
			class="flex-col"
			:style="[{ width: _width + 'rpx' }, _height ? { height: _height + 'rpx' } : '']"
		>
			<view
				@touchStart="onScrollStart"
				@touchend="onScrollEnd"
				v-if="isShowRender"
				:style="{ transform: `translateY(${reFresh == 2 ? -refreshJuli : 0}px)` }"
				class="flex contentx"
			>
				<slot></slot>
			</view>
		</scroll-view>
	</tm-sheet>
</template>

<script lang="ts" setup>
/**
 * 选项卡面板
 * @description 不可以单独使用，必须放置在tm-tabs组件中使用。
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmTabs from '../tm-tabs/tm-tabs.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import tmText from '../tm-text/tm-text.vue'
import { tabsobj } from '../tm-tabs/interface'
import { Ref, ref, computed, watch, onUnmounted, inject, ComputedRef, getCurrentInstance, ComponentInternalInstance, PropType, onMounted } from 'vue'
import { custom_props } from '../../tool/lib/minxs'
const tabs = ref<InstanceType<typeof tmTabs> | null>(null)
const proxy = getCurrentInstance()?.proxy ?? null
const props = defineProps({
	...custom_props,
	transprent: {
		type: [Boolean, String],
		default: false
	},
	color: {
		type: String,
		default: 'white'
	},
	//是否显示红点，与count不能同时出现
	dot: {
		type: Boolean,
		default: false
	},
	count: {
		type: [String, Number],
		default: ''
	},
	//红点颜色
	dotColor: {
		type: String,
		default: 'red'
	},
	name: {
		type: [String, Number],
		default: '',
		required: true
	},
	title: {
		type: String,
		default: ''
	},
	icon: {
		type: String,
		default: ''
	},
	pullFun: {
		type: [Function] as PropType<(type: 'top' | 'bottom') => boolean>,
		default: () => {
			return (type: 'top' | 'bottom') => {
				return true
			}
		}
	}
})

//---scroll下拉刷新---------
//0未初始化/松手可能未被触发，没有达到刷新的条件（即被中止了），1下拉手势中还未放开手。2松开手触发下拉刷新，3复位刷新完成
type reFreshType = 0 | 1 | 2 | 3
const sc_top = ref(0)
const refreshJuli = -100 //达到120才能被刷新
const reFresh: Ref<reFreshType> = ref(0)
const isyesResh = ref(false) //是否达到了规定刷新条件。
const isUpToogle = ref(true) //手指是否放开。

//---------------

const _pname = computed(() => String(props.name))
//父级方法。
let parent: any = proxy?.$parent
while (parent) {
	if (parent?.tmTabsId == 'tmTabsId' || !parent) {
		break
	} else {
		parent = parent?.$parent ?? null
	}
}

if (typeof _pname.value != 'undefined' && _pname.value != '') {
	parent?.pushKey({
		key: _pname.value,
		title: props.title,
		icon: props.icon,
		dot: props.dot,
		count: props.count,
		dotColor: props.dotColor
	})
}

const _width = inject(
	'tabsWidth',
	computed(() => 0)
)
const _height = inject(
	'tabsheight',
	computed(() => 0)
)
const tabsActiveName: ComputedRef<string | number | undefined> = inject(
	'tabsActiveName',
	computed(() => undefined)
)
const tabsActiveCacheTabse = inject(
	'tabsActiveCacheTabse',
	computed<Array<tabsobj>>(() => {
		return []
	})
)
const tabsSwiper = inject(
	'tabsSwiper',
	computed(() => false)
)
const tabsSwiperIsMoveing = inject(
	'tabsSwiperIsMoveing',
	computed(() => false)
)
const tabsSwiperDisAbledPull = inject(
	'tabsSwiperDisAbledPull',
	computed(() => true)
)

const activeIndex = inject(
	'tabsActiveactiveIndex',
	computed(() => 0)
)
// const selfIndex = computed(()=>tabsActiveCacheTabse.value.findIndex(el=>el.key==tabsActiveName.value))
const selfIndex = computed(() => tabsActiveCacheTabse.value.findIndex((el) => el.key == _pname.value))

const isShowRender = computed(() => {
	if (tabsSwiper.value == false) {
		return selfIndex.value == activeIndex.value
	}
	return selfIndex.value >= activeIndex.value - 1 && selfIndex.value <= activeIndex.value + 1
})

watch([() => props.title, () => props.icon, () => props.dot, () => props.dotColor, () => props.count], () => {
	parent?.setTitle({
		key: _pname.value,
		title: props.title,
		icon: props.icon,
		dot: props.dot,
		count: props.count,
		dotColor: props.dotColor
	})
})
onUnmounted(() => {
	parent?.unbindKey(_pname.value)
})

//=====
function onScroll(e: any) {
	// 切换页面中不允许下拉。
	if (tabsSwiperIsMoveing.value) return
	if (reFresh.value == 2 && isUpToogle.value) return
	if (sc_top.value <= refreshJuli) {
		isyesResh.value = true
	} else {
		isyesResh.value = false
	}

	sc_top.value = e.detail.scrollTop
	reFresh.value = 1
}
function onScrollStart() {
	// 切换页面中不允许下拉。
	if (tabsSwiperIsMoveing.value) return
	isUpToogle.value = false
}
async function onScrollEnd() {
	// 切换页面中不允许下拉。
	if (tabsSwiperIsMoveing.value) return
	isUpToogle.value = true
	if (reFresh.value == 2) return
	if (isyesResh.value) {
		let p = await funPull('top')
		if (p) {
			reFresh.value = 0
		} else {
			isUpToogle.value = false
		}
	} else {
		reFresh.value = 0
	}
}
async function onScrollBootom() {
	// 切换页面中不允许下拉。
	if (tabsSwiperIsMoveing.value) return
	isUpToogle.value = true
	if (reFresh.value == 2) return
	if (isyesResh.value) {
		let p = await funPull('bottom')
		if (p) {
			reFresh.value = 0
		} else {
			isUpToogle.value = false
		}
	} else {
		reFresh.value = 0
	}
}
async function funPull(type: 'top' | 'bottom') {
	if (typeof props.pullFun === 'function') {
		uni.showLoading({
			title: '...',
			mask: true
		})
		let p = await props.pullFun(type)
		if (typeof p === 'function') {
			p = await p(type)
		}
		uni.hideLoading()
		return p
	}
}
</script>

<style scoped>
.contentx {
	transition-duration: 0.5s;
	transition-timing-function: ease;
	transition-delay: 0s;
	transform: translateY(0px);
	transition-property: transform, top;
}
</style>
