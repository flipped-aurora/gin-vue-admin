<template>
	<view class="flex-1 relative" :style="{ height: props.height + 'rpx' }">
		<!-- #ifndef APP-NVUE -->
		<!-- #ifndef MP-ALIPAY -->
		<picker-view
			@pickend="emits('end')"
			@pickstart="emits('start')"
			v-if="showDom"
			:value="[colIndex]"
			@change="colchange"
			:style="[{ height: props.height + 'rpx' }]"
			:mask-style="maskStyle"
			:immediateChange="props.immediateChange"
			indicator-style="height:50px;background-image: linear-gradient(to bottom,rgba(0, 0, 0, 0.04),rgba(0, 0, 0, 0.01),rgba(0,0,0, 0.04));"
		>
			<picker-view-column>
				<view
					v-for="(item, index) in _data"
					:key="index"
					class="flex itemcel"
					:class="[index == colIndex ? '' : 'UnitemSelected', item['disabled'] ? 'opacity-5' : '']"
				>
					<TmText
						v-if="typeof item == 'string'"
						_class="text-align-center"
						:font-size="item.length > 7 ? 24 : 30"
						:dark="isDark"
						:label="item"
					></TmText>
					<TmText
						v-if="typeof item == 'object'"
						_class="text-align-center"
						:font-size="item[props.mapKey || props.dataKey].length > 7 ? 24 : 30"
						:dark="isDark"
						:label="item[props.mapKey || props.dataKey] || ''"
					></TmText>
				</view>
			</picker-view-column>
		</picker-view>
		<!-- #endif -->
		<!-- #endif -->
		<!-- #ifdef MP-ALIPAY -->
		<picker-view
			@pickend="emits('end')"
			@pickstart="emits('start')"
			v-if="showDom"
			:value="[colIndex]"
			@change="colchange"
			:style="[{ height: props.height + 'rpx' }]"
			:mask-style="maskStyle"
			:immediateChange="props.immediateChange"
		>
			<picker-view-column>
				<view
					v-for="(item, index) in _data"
					:key="index"
					:class="[index == colIndex ? '' : 'UnitemSelected', item['disabled'] ? 'opacity-5' : '']"
					:style="{
						color: store.tmStore.dark ? 'white' : 'black',
						fontSize: (item.length > 7 ? 24 : 30) + 'rpx'
					}"
				>
					{{ typeof item == 'string' ? item : item[props.mapKey || props.dataKey] || '' }}
				</view>
			</picker-view-column>
		</picker-view>
		<!-- #endif -->

		<!-- #ifdef APP-NVUE -->
		<picker-view
			ref="picker"
			@pickend="emits('end')"
			@pickstart="emits('start')"
			v-if="showDom"
			:value="[colIndex]"
			@change="colchange"
			:style="[{ height: props.height + 'rpx' }]"
			indicator-style="height:50px"
			:mask-top-style="maskStyleNvue"
			:immediateChange="props.immediateChange"
			:mask-bottom-style="maskStyleNvue"
		>
			<picker-view-column>
				<view :class="[item['disabled'] ? 'opacity-5' : '']" v-for="(item, index) in _data" :key="index" class="flex itemcel">
					<TmText
						v-if="typeof item == 'string'"
						_class="text-align-center"
						:font-size="item.length > 7 ? 24 : 30"
						:dark="isDark"
						:label="item"
					></TmText>
					<TmText
						v-if="typeof item == 'object'"
						_class="text-align-center"
						:font-size="item[props.mapKey || props.dataKey].length > 7 ? 24 : 30"
						:dark="isDark"
						:label="item[props.mapKey || props.dataKey] || ''"
					></TmText>
				</view>
			</picker-view-column>
		</picker-view>
		<view
			v-if="isDark"
			:userInteractionEnabled="false"
			class="top absolute l-0 t-0"
			:style="{ height: maskHeight + 'px', width: maskWidth + 'px' }"
		></view>
		<view
			v-if="isDark"
			:userInteractionEnabled="false"
			class="bottom absolute l-0 b-0"
			:style="{ height: maskHeight + 'px', width: maskWidth + 'px' }"
		></view>
		<!-- #endif -->
	</view>
</template>
<script lang="ts" setup>
import { useTmpiniaStore } from '../../tool/lib/tmpinia'
import { computed, PropType, Ref, onUpdated, watchEffect, ref, getCurrentInstance, nextTick, onMounted, watch, toRaw } from 'vue'
import TmText from '../tm-text/tm-text.vue'
import { columnsItem } from './interface'
// #ifdef APP-PLUS-NVUE
const dom = uni.requireNativePlugin('dom')
// #endif
const emits = defineEmits(['change', 'end', 'start'])
const proxy = getCurrentInstance()?.proxy ?? null
const store = useTmpiniaStore()
const props = defineProps({
	followTheme: {
		type: [Boolean],
		default: true
	},
	col: {
		type: Number,
		default: 0
	},
	//禁用的部分日期，禁用的日期将不会被选中，就算滑到了该位置，也会回弹到之前的时间。
	disabled: {
		type: Boolean,
		default: false
	},
	height: {
		type: Number,
		default: 600
	},
	data: {
		type: Array as PropType<Array<columnsItem>>,
		default: () => [],
		required: true
	},
	//当columns项目中的data数据为对象时的key取值字段。
	dataKey: {
		type: String,
		default: 'text'
	},
	//当columns项目中的data数据为对象时的key取值字段。兼容上方dataKey,因为微信dataKey与本字段重名，无法设置。
	mapKey: {
		type: String,
		default: 'text'
	},
	immediateChange: {
		type: Boolean,
		default: true
	}
})
const isDark = computed(() => store.tmStore.dark)
const _data = computed(() => props.data)
const colIndex = ref(0)
const showDom = ref(false)
const maskHeight = computed(() => {
	return (uni.upx2px(props.height) - 34) / 2
})

const maskWidth = ref(0)
const maskStyle = computed(() => {
	let str_white =
		'background-image:linear-gradient(rgba(255,255,255,0.95),rgba(255,255,255,0.6)),linear-gradient(rgba(255,255,255,0.6),rgba(255,255,255,0.95))'
	let str_black =
		'background-image:linear-gradient(rgba(17, 17, 17, 1.0),rgba(106, 106, 106, 0.2)),linear-gradient(rgba(106, 106, 106, 0.2),rgba(17, 17, 17, 1.0))'

	// #ifdef APP-NVUE
	str_black = 'background-image: linear-gradient(to bottom,rgba(30, 30, 30, 0.9),rgba(104, 104, 104, 0.6))'
	// #endif
	if (!isDark.value) {
		return str_white
	}
	return str_black
})
const maskStyleNvue = computed(() => {
	let str_white = 'background:linear-gradient(rgba(255,255,255,0.95),rgba(255,255,255,0.6))'
	let str_black = 'background:linear-gradient(0deg,rgba(0,0,0,0.4),rgba(0,0,0,0),rgba(0,0,0,0.4))'

	return isDark ? str_black : str_white
})

onMounted(() => {
	//在微信小程序因为有渲染等待30ms，为了兼容统一全部等待30ms
	showDom.value = true
	nvuegetClientRect()
	setTimeout(function () {
		colIndex.value = props.col
	}, 50)
})
onUpdated(() => nvuegetClientRect())
watch(
	() => props.col,
	() => {
		nextTick(() => {
			colIndex.value = props.col
		})
	}
)
function colchange(e) {
	colIndex.value = e.detail.value[0]
	emits('change', colIndex.value)
}

function nvuegetClientRect() {
	nextTick(function () {
		// #ifdef APP-PLUS-NVUE
		dom.getComponentRect(proxy.$refs.picker, function (res) {
			if (res?.size) {
				maskWidth.value = res.size.width

				if (res.size.width == 0) {
					nvuegetClientRect()
				}
			}
		})
		// #endif
	})
}
</script>
<style scoped>
.top {
	background-image: linear-gradient(to bottom, rgba(17, 17, 17, 1), rgba(36, 36, 36, 0.6));
}
.bottom {
	background-image: linear-gradient(to top, rgba(17, 17, 17, 1), rgba(36, 36, 36, 0.6));
}
.itemcel {
	justify-content: center;
	height: 50px;
	align-items: center;
}
.itemSelected {
	background-image: linear-gradient(to bottom, rgba(0, 0, 0, 0.04), rgba(0, 0, 0, 0.01), rgba(0, 0, 0, 0.04));
}
.UnitemSelected {
	background-image: rgba(0, 0, 0, 0);
}
</style>
