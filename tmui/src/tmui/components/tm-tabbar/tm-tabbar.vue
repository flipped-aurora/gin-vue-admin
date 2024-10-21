<template>
	
	<view>
		
		<view
		class="fixed l-0 b-0 flex tabbar"
		:style="{
			width: _width + 'px',
			height: _totalBarHeight + 'px',
			transform: `translateY(${props.bottom}rpx)`,
			zIndex: props.zIndex
		}"
	>
		<!-- 			@click="emits('click', $event)"    -->
		<view class="relative" style="top: 15px">
			<tm-sheet
				:blur="_blur"
				:color="props.color"
				parenClass="relative"
				class="relative"
				:followTheme="props.followTheme"
				:dark="props.dark"
				:round="props.round"
				:shadow="props.shadow"
				:outlined="props.outlined"
				:border="0"
				borderDirection="top"
				:text="props.text"
				:transprent="false"
				:linear="props.linear"
				:linearDeep="props.linearDeep"
				:margin="[0, 0]"
				:padding="[0, 0]"
				:height="_BarHeight"
				:width="_width"
				unit="px"
			></tm-sheet>
		</view>
		<view class="absolute flex flex-col l-0" :style="{ width: _width + 'px', height: _totalBarHeight + 'px' }">
			<view class="relative barcont flex flex-row flex-row-top-center flex-around flex-1" :style="{ width: _width + 'px' }"><slot></slot></view>
		</view>
	</view>
		<view v-if="_place" :style="{
			height: _totalBarHeight + 'px',
		}"></view>
	</view>
</template>

<script lang="ts" setup>
/**
 * 底部导航栏
 * @description 内部只能放置,tm-tabbar-item
 */
import tmSheet from '../tm-sheet/tm-sheet.vue'
import tmText from '../tm-text/tm-text.vue'
import tmIcon from '../tm-icon/tm-icon.vue'
import { custom_props } from '../../tool/lib/minxs'
import { getCurrentInstance, computed, Ref, ref, provide, inject, onUpdated, onMounted, onUnmounted, nextTick, watch, PropType } from 'vue'
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
const store = useTmpiniaStore()
const emits = defineEmits(['change', 'update:active'])
const props = defineProps({
	...custom_props,
	transprent: {
		type: [Boolean],
		default: false
	},
	color: {
		type: [String],
		default: 'white'
	},
	text: {
		type: [Boolean],
		default: false
	},
	round: {
		type: [Number],
		default: 0
	},
	shadow: {
		type: [Number],
		default: 0
	},
	//如果为0取当前窗口宽度。
	width: {
		type: [Number],
		default: 0
	},
	//底部偏移
	bottom: {
		type: [Number],
		default: 0
	},
	//是否显示安全区域
	showSafe: {
		type: [Boolean],
		default: false
	},
	//动态激活项的索引-1表示不激活任何项。
	active: {
		type: Number,
		default: -1
	},
	//是否自动选中。如果禁用，则用户通过active来切换选中值，否则，点击项目时自动选中。
	autoSelect: {
		type: Boolean,
		default: true
	},
	/** 在nvue上没有用处，请自行放到最尾部的地方。 */
	zIndex: {
		type: [Number],
		default: 200
	},
	/**是否占位，底部填充空白view */
	place: {
		type: Boolean,
		default: false
	},
})
let sys = uni.getSystemInfoSync()
const _width = computed(() => uni.upx2px(props.width) || sys?.windowWidth || 750)
const _blur = computed(() => props.blur)
const _place = computed(() => props.place)
const _activeUrl = ref('')
const _activeUid = ref('')
const tmTabbarId = 'tmTabbarId'
const _cachlist: Ref<Array<string | number>> = ref([])
const _showSafe = ref(props.showSafe)
const _activeIndex = ref(props.active)
const win_bottom = sys?.safeAreaInsets?.bottom ?? 0
if (win_bottom > 0) {
	_showSafe.value = true
}

const _totalBarHeight = computed(() => {
	if (_showSafe.value) return 90
	return 75
})
const _BarHeight = computed(() => {
	if (_showSafe.value) return 75
	return 60
})
function setNowurl(url: string, nowuid: number) {
	_activeUrl.value = url
	_activeUid.value = String(nowuid)
}
function pushKey(uid: number) {
	_cachlist.value = [...new Set([..._cachlist.value, uid])]
}
function delKey(uid: number) {
	_cachlist.value = _cachlist.value.filter((el) => el != uid)
}
defineExpose({ tmTabbarId, setNowurl, pushKey, delKey })
provide(
	'tmTabbarUrl',
	computed(() => _activeUrl.value)
)
provide(
	'tmTabbarUid',
	computed(() => _activeUid.value)
)
provide(
	'tmTabbarWidth',
	computed(() => Math.ceil(_width / _cachlist.value.length))
)
provide(
	'tmTabbarItemList',
	computed(() => _cachlist.value)
)
provide(
	'tmTabbarItemActive',
	computed(() => _activeIndex.value)
)
provide('tmTabbarItemSafe', _showSafe.value)
provide(
	'tmTabbarItemAutoSelect',
	computed(() => props.autoSelect)
)
watch(
	() => props.active,
	() => {
		if (props.active == _activeIndex.value) return
		_activeIndex.value = props.active
	}
)
watch(_activeIndex, () => {
	emits('change', _activeIndex.value)
	emits('update:active', _activeIndex.value)
})
</script>

<style>
/* #ifdef APP-NVUE */
.barcont {
	display: flex;
	flex-direction: row;
	justify-content: space-around;
	align-items: center;
}
.tabbar {
}
/* #endif */
/* #ifndef APP-NVUE */
.barcont {
	display: flex;
	flex-direction: row !important;
	justify-content: space-around !important;
	align-items: flex-start;
}
/* #endif */
</style>
