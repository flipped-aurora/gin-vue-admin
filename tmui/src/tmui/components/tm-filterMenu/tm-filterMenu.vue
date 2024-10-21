<template>
	<view class="relative" ref="filterMenu" id="filterMenu" :style="{ width: props.width + 'rpx' }">
		<!-- 静态菜单 -->
		<tm-sheet
			:color="props.color"
			:margin="[0, 0]"
			:padding="[0, 0]"
			:height="props.height"
			:width="props.width"
			_class="flex flex-row flex-row-center-center"
		>
			<view
				v-if="isNaN(activeIndex) || isButton"
				@click="onClick(item, index)"
				:style="{ height: props.height + 'rpx', width: itemwidth + 'rpx' }"
				class="flex flex-row flex-row-center-center"
				v-for="(item, index) in cachList"
				:key="index"
			>
				<tm-text
					:userInteractionEnabled="false"
					:label="item.text"
					:font-size="item.fontSize"
					:color="activeIndex == index ? item.fontColor : store.tmStore.dark ? '' : item.unFontColor"
					_class="pr-12"
				></tm-text>
				<tm-icon
					v-if="item.icon && item.unIcon"
					:userInteractionEnabled="false"
					:font-size="item.fontSize - 4"
					:name="activeIndex == index ? item.icon : item.unIcon"
					:color="activeIndex == index ? item.fontColor : store.tmStore.dark ? '' : item.unFontColor"
				></tm-icon>
			</view>
		</tm-sheet>

		<!-- 背景遮罩,透明背景 -->
		<view
			@click="close"
			@touchmove.stop="stopEvent"
			class="bg fixed"
			:style="{
				left: (isNaN(activeIndex) || isButton ? '-10000' : 0) + 'px',
				top: 0 + 'px',
				height: sysinfo.height + 'px',
				width: sysinfo.width + 'px',
				zIndex:402
			}"
		></view>
		<!-- #ifdef MP-WEIXIN || MP-ALIPAY-->
		<root-portal>
		<!-- #endif -->
		<!-- #ifdef H5 -->
		<teleport to="body" >
		<!-- #endif -->
		<!-- 动态菜单及内容的背景 -->
		<view
			@click="close"
			@touchmove.stop="stopEvent"
			:class="[isNaN(activeIndex) ? 'bgAnioff' : 'bgAni']"
			class="bgContent fixed"
			:style="{
				left: (isNaN(activeIndex) || isButton ? '-10000' : (props.fixed?0:el_left)) + 'px',
				top: _fixed ?  '0px' : el_top  + 'px',
				width: props.width + 'rpx',
				height: sysinfo.height + el_top + 'px',
				zIndex:402
			}"
		></view>
		<!-- 动态菜单及内容 -->
		<view
			@click="close"
			@touchmove.stop="stopEvent"
			class="content fixed"
			:style="{
				left: (isNaN(activeIndex) || isButton ? '-10000' : (props.fixed?0:el_left)) + 'px',
				top: _fixed ? 0 + 'px' : el_top  + 'px',
				width: props.width + 'rpx',
				height: sysinfo.height + el_top + 'px',
				zIndex:402
			}"
		>
			<tm-sheet
				:shadow="1"
				:color="props.color"
				:margin="[0, 0]"
				:padding="[0, 0]"
				:height="props.height"
				:width="props.width"
				_class="flex flex-row flex-row-center-center"
			>
				<view
					@click.stop="onClick(item, index)"
					:style="{ height: props.height + 'rpx', width: itemwidth + 'rpx' }"
					class="flex flex-row flex-row-center-center"
					v-for="(item, index) in cachList"
					:key="index"
				>
					<tm-text
						:userInteractionEnabled="false"
						:label="item.text"
						:font-size="item.fontSize"
						:color="activeIndex == index ? item.fontColor : store.tmStore.dark ? '' : item.unFontColor"
						_class="pr-12"
					></tm-text>
					<tm-icon
						v-if="item.icon && item.unIcon"
						:userInteractionEnabled="false"
						:font-size="item.fontSize - 4"
						:name="activeIndex == index ? item.icon : item.unIcon"
						:color="activeIndex == index ? item.fontColor : store.tmStore.dark ? '' : item.unFontColor"
					></tm-icon>
				</view>
			</tm-sheet>
			<slot name="default"></slot>
		</view>
		<!-- #ifdef MP-WEIXIN || MP-ALIPAY-->
		</root-portal>
		<!-- #endif -->
		<!-- #ifdef H5 -->
		</teleport>
		<!-- #endif -->

	</view>
</template>
<script lang="ts" setup>
import tmSheet from '@/tmui/components/tm-sheet/tm-sheet.vue'
import tmText from '@/tmui/components/tm-text/tm-text.vue'
import tmIcon from '@/tmui/components/tm-icon/tm-icon.vue'
import { inject, computed, getCurrentInstance, nextTick, ref, onMounted, Ref, provide, toRaw, watch } from 'vue'
import { FilterMenuType } from './interface'
import { useTmpiniaStore } from '@/tmui/tool/lib/tmpinia'
import { useWindowInfo } from '../../tool/useFun/useWindowInfo'

// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
const animation = uni.requireNativePlugin('animation')
// #endif
const proxy = getCurrentInstance()?.proxy ?? null
const emits = defineEmits(['click', 'close', 'cancel', 'confirm', 'onButton', 'update:modelValue'])
const store = useTmpiniaStore()

const props = defineProps({
	width: {
		type: Number,
		default: 750
	},
	height: {
		type: Number,
		default: 70
	},
	color: {
		type: String,
		default: 'white'
	},
	activeColor: {
		type: String,
		default: 'primary'
	},
	modelValue: {
		type: Number,
		default: NaN
	},
	fixed: {
		type: Boolean,
		default: false
	}
})
const sysinfo = useWindowInfo()
//元件的位置
const el_left = ref(0)
const el_top = ref(0)

//缓存的节点数据.
const cachList: Ref<Array<FilterMenuType>> = ref([])
const itemwidth = computed(() => Math.ceil(props.width / (cachList.value.length || 1)))
const activeIndex = ref(props.modelValue)
const _fixed = computed(() => props.fixed)
const isButton = computed(() => {
	if (isNaN(activeIndex.value)) return false
	let el = cachList.value[activeIndex.value]
	if (!el) return true
	return el.isButton
})
onMounted(() => {
	getDomRect()
})
function getDomRect() {
	// #ifdef APP-NVUE
	nextTick(function () {
		dom.getComponentRect(proxy.$refs.filterMenu, function (res: any) {
			if (!res?.size) {
				getDomRect()
				return
			}
			el_left.value = res?.size.left
			el_top.value = res?.size.top
		})
	})
	// #endif
	// #ifndef APP-NVUE
	nextTick(() => {
		uni.createSelectorQuery()
			.in(proxy)
			.select('#filterMenu')
			.boundingClientRect()
			.exec(
				(
					result: Array<{
						id: string
						left: number
						height: number
						right: number
						top: number
					}>
				) => {
					if (!result[0]) return
					let el = result[0]
					el_left.value = el.left
					el_top.value = el.top
				}
			)
	})
	// #endif
}
watch(
	() => props.modelValue,
	() => {
		activeIndex.value = props.modelValue
	}
)
function onClick(item: FilterMenuType, index: number) {
	if (activeIndex.value === index) {
		if (isButton.value) {
			emits('onButton', false, index)
		}
		activeIndex.value = NaN
		emits('update:modelValue', NaN)
		emits('cancel')
		return
	}
	activeIndex.value = index
	emits('click', item, index)

	if (isButton.value) {
		emits('onButton', true, index)
	}

	emits('update:modelValue', index)
}
function close() {
	activeIndex.value = NaN

	emits('close')
	emits('update:modelValue', NaN)
	emits('cancel')
}
function pushKey(arg: FilterMenuType) {
	let index = cachList.value.findIndex((el) => el.id == arg.id)
	if (index > -1) {
		cachList.value[index] = { ...arg }
	} else {
		cachList.value.push({ ...arg })
	}
}
function delKey(id: string | number) {
	let index = cachList.value.findIndex((el) => el.id == id)
	if (index > -1) {
		cachList.value.splice(index, 1)
	}
}
function stopEvent(e: any) {
	e?.preventDefault()
	e?.stopPropagation()
}
function filterConfirm() {
	activeIndex.value = NaN
	emits('update:modelValue', NaN)
	emits('confirm', toRaw(cachList.value[activeIndex.value]), toRaw(activeIndex.value))
}
function filterCancel() {
	activeIndex.value = NaN
	emits('update:modelValue', NaN)
	emits('cancel')
}
provide(
	'activeIndex',
	computed(() => activeIndex.value)
)
provide(
	'AllList',
	computed(() => cachList.value)
)
provide(
	'FilterMenuMaxHeight',
	computed(() => sysinfo.height + el_top.value)
)
defineExpose({
	FilterMenu: 'FilterMenu',
	pushKey,
	delKey,
	close,
	filterCancel,
	filterConfirm
})
</script>

<style scoped>
.bg {
	background-color: rgba(122, 106, 106, 0);
	/* #ifndef APP-NVUE */
	/* backdrop-filter: blur(0px); */
	z-index: 50;
	/* #endif */
}

.bgContent {
	background-color: rgba(0, 0, 0, 0.35);
	/* #ifndef APP-NVUE */
	/* backdrop-filter: blur(5px); */
	z-index: 51;
	/* #endif */
	/* #ifdef APP-NVUE */
	opacity: 0.4;
	transition-delay: 100;
	transition-timing-function: linear;
	transition-property: transform;
	/* #endif */
}
.bgAni {
	/* #ifndef APP-NVUE */
	animation: bgcolorani 0.35s linear;
	/* #endif */
	/* #ifdef APP-NVUE */
	opacity: 1;
	/* #endif */
}
.bgAnioff {
	/* #ifdef APP-NVUE */
	opacity: 0.4;
	/* #endif */
}
.content {
	/* #ifndef APP-NVUE */
	z-index: 52;
	/* #endif */
}

/* #ifndef APP-NVUE */
@keyframes bgcolorani {
	0% {
		background-color: rgba(0, 0, 0, 0);
		/* backdrop-filter: blur(0px); */
	}

	100% {
		background-color: rgba(0, 0, 0, 0.35);
		/* backdrop-filter: blur(5px); */
	}
}

/* #endif */
</style>
